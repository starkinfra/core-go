package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/starkbank/ecdsa-go/v2/ellipticcurve/curve"
	"github.com/starkbank/ecdsa-go/v2/ellipticcurve/privatekey"
	"github.com/starkinfra/core-go/starkcore/user/project"
	"github.com/starkinfra/core-go/starkcore/utils/hosts"
	"github.com/starkinfra/core-go/starkcore/utils/rest"
)

var patchMultiResource = map[string]string{"name": "Card"}

type recordedRequest struct {
	Method string
	Url    string
	Body   string
	Header http.Header
}

// fakeTransport replaces only the outgoing HTTP call: signing, headers and payload building run for real.
type fakeTransport struct {
	status   int
	body     string
	recorded []recordedRequest
}

func (f *fakeTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	body, _ := io.ReadAll(req.Body)
	f.recorded = append(f.recorded, recordedRequest{Method: req.Method, Url: req.URL.String(), Body: string(body), Header: req.Header})
	return &http.Response{
		StatusCode: f.status,
		Body:       io.NopCloser(strings.NewReader(f.body)),
		Header:     http.Header{},
		Request:    req,
	}, nil
}

func useFakeApi(t *testing.T, status int, body string) *fakeTransport {
	fake := &fakeTransport{status: status, body: body}
	original := http.DefaultTransport
	http.DefaultTransport = fake
	t.Cleanup(func() { http.DefaultTransport = original })
	return fake
}

func patchMultiUser() project.Project {
	return project.Project{
		Id:          "5656565656565656",
		PrivateKey:  privatekey.New(curve.Secp256k1).ToPem(),
		Environment: "sandbox",
	}
}

func patchMulti(entities interface{}, query map[string]interface{}) ([]byte, error) {
	response, err := rest.PatchMulti("0.0.0", hosts.Infra, "v2", "en-US", 15, patchMultiUser(), patchMultiResource, entities, query)
	if err.Errors != nil {
		return nil, &errorsAsError{err.Errors[0].Code}
	}
	return response, nil
}

type errorsAsError struct{ code string }

func (e *errorsAsError) Error() string { return e.code }

func TestPatchMultiSendsPatchWithPluralPayloadAndParsesResponse(t *testing.T) {
	fake := useFakeApi(t, 200, `{"cards": [{"id": "1", "holderName": "Tony Stark"}, {"id": "2", "holderName": "Tony Stark"}]}`)

	response, err := patchMulti([]map[string]interface{}{
		{"id": "1", "holderName": "Tony Stark", "holderTaxId": "012.345.678-90"},
		{"id": "2", "holderName": "Tony Stark", "holderTaxId": "012.345.678-90"},
	}, nil)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if len(fake.recorded) != 1 {
		t.Fatalf("expected one request, got %d", len(fake.recorded))
	}
	sent := fake.recorded[0]
	if sent.Method != "PATCH" {
		t.Errorf("method should be PATCH, got %s", sent.Method)
	}
	if !strings.HasSuffix(sent.Url, "/v2/card") {
		t.Errorf("unexpected url %s", sent.Url)
	}
	if sent.Header.Get("Access-Signature") == "" {
		t.Errorf("request was not signed")
	}
	var body map[string][]map[string]interface{}
	if jsonError := json.Unmarshal([]byte(sent.Body), &body); jsonError != nil {
		t.Fatalf("body is not json: %s", sent.Body)
	}
	if len(body["cards"]) != 2 || body["cards"][0]["id"] != "1" || body["cards"][1]["holderTaxId"] != "012.345.678-90" {
		t.Errorf("unexpected payload %s", sent.Body)
	}

	var cards []map[string]interface{}
	if jsonError := json.Unmarshal(response, &cards); jsonError != nil {
		t.Fatalf("response is not the entity list: %s", response)
	}
	if len(cards) != 2 || cards[0]["id"] != "1" || cards[1]["id"] != "2" {
		t.Errorf("unexpected parsed response %s", response)
	}
}

func TestPatchMultiDoesNotSendUnsetAttributes(t *testing.T) {
	fake := useFakeApi(t, 200, `{"cards": [{"id": "1"}]}`)

	_, err := patchMulti([]map[string]interface{}{{"id": "1", "tags": []string{"x"}, "holderName": nil}}, nil)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if strings.Contains(fake.recorded[0].Body, "holderName") {
		t.Errorf("nil attributes must not be sent: %s", fake.recorded[0].Body)
	}
}

func TestPatchMultiSendsQueryInTheUrl(t *testing.T) {
	fake := useFakeApi(t, 200, `{"cards": [{"id": "1"}]}`)

	_, err := patchMulti([]map[string]interface{}{{"id": "1"}}, map[string]interface{}{"expand": []string{"rules"}})
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if !strings.HasSuffix(fake.recorded[0].Url, "/v2/card?expand=rules") {
		t.Errorf("query not in the url: %s", fake.recorded[0].Url)
	}
}

func TestPatchMultiReturnsTheApiErrorUntouched(t *testing.T) {
	useFakeApi(t, 400, `{"errors": [{"code": "invalidHolder", "message": "must all be provided together"}]}`)

	_, err := patchMulti([]map[string]interface{}{{"id": "1", "holderName": "Tony Stark"}}, nil)

	if err == nil || err.Error() != "invalidHolder" {
		t.Errorf("expected the invalidHolder error, got %v", err)
	}
}
