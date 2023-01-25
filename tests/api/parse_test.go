package api

import (
	"fmt"
	"github.com/starkinfra/core-go/starkcore/utils/hosts"
	"github.com/starkinfra/core-go/starkcore/utils/parse"
	"github.com/starkinfra/core-go/tests/utils"
	User "github.com/starkinfra/core-go/tests/utils/user"
	"testing"
)

func TestRightParseAndVerify(t *testing.T) {
	uuid := "21f174ab942843eb90837a5c3135dfd6"
	validSignature := "MEYCIQC+Ks0M54DPLEbHIi0JrMiWbBFMRETe/U2vy3gTiid3rAIhANMmOaxT03nx2bsdo+vg6EMhWGzdphh90uBH9PY2gJdd"
	parsed := parse.ParseAndVerify(
		uuid,
		validSignature,
		utils.SdkVersion,
		utils.ApiVersion,
		utils.Language,
		utils.Timeout,
		hosts.Infra,
		User.ExampleProjectBank,
		"",
	)
	fmt.Println(parsed)
}

func TestWrongParseAndVerify(t *testing.T) {
	uuid := "21f174ab942843eb90837a5c3135dfd6"
	invalidSignature := "MEUCIQDOpo1j+V40DNZK2URL2786UQK/8mDXon9ayEd8U0/l7AIgYXtIZJBTs8zCRR3vmted6Ehz/qfw1GRut/eYyvf1yOk="
	parsed := parse.ParseAndVerify(
		uuid,
		invalidSignature,
		utils.SdkVersion,
		utils.ApiVersion,
		utils.Language,
		utils.Timeout,
		hosts.Infra,
		User.ExampleProjectBank,
		"",
	)
	fmt.Println(parsed)
}

func TestRightVerify(t *testing.T) {
	uuid := "21f174ab942843eb90837a5c3135dfd6"
	validSignature := "MEYCIQC+Ks0M54DPLEbHIi0JrMiWbBFMRETe/U2vy3gTiid3rAIhANMmOaxT03nx2bsdo+vg6EMhWGzdphh90uBH9PY2gJdd"
	parsed := parse.ParseAndVerify(
		uuid,
		validSignature,
		utils.SdkVersion,
		utils.ApiVersion,
		utils.Language,
		utils.Timeout,
		hosts.Infra,
		User.ExampleProjectBank,
		"",
	)
	fmt.Println(parsed)
}

func TestWrongVerify(t *testing.T) {
	uuid := "21f174ab942843eb90837a5c3135dfd6"
	invalidSignature := "MEUCIQDOpo1j+V40DNZK2URL2786UQK/8mDXon9ayEd8U0/l7AIgYXtIZJBTs8zCRR3vmted6Ehz/qfw1GRut/eYyvf1yOk="
	parsed := parse.ParseAndVerify(
		uuid,
		invalidSignature,
		utils.SdkVersion,
		utils.ApiVersion,
		utils.Language,
		utils.Timeout,
		hosts.Infra,
		User.ExampleProjectBank,
		"",
	)
	fmt.Println(parsed)
}

func TestMalformed(t *testing.T) {
	content := "{\"event\": {\"created\": \"2021-04-26T20:16:51.866857+00:00\", \"id\": \"5415223380934656\", \"log\": {\"created\": \"2021-04-26T20:16:50.927706+00:00\", \"errors\": [], \"id\": \"4687457496858624\", \"invoice\": {\"amount\": 256, \"brcode\": \"00020101021226890014br.gov.bcb.pix2567invoice-h.sandbox.starkbank.com/v2/afdf94b770b0458a8440a335daf77c4c5204000053039865802BR5915Stark Bank S.A.6009Sao Paulo62070503***6304CC32\", \"created\": \"2021-04-26T20:16:50.886319+00:00\", \"descriptions\": [{\"key\": \"Field1\", \"value\": \"Something\"}], \"discountAmount\": 0, \"discounts\": [{\"due\": \"2021-05-07T09:43:15+00:00\", \"percentage\": 10.0}], \"due\": \"2021-05-09T19:11:39+00:00\", \"expiration\": 123456789, \"fee\": 0, \"fine\": 2.5, \"fineAmount\": 0, \"id\": \"5941925571985408\", \"interest\": 1.3, \"interestAmount\": 0, \"link\": \"https://cdottori.sandbox.starkbank.com/invoicelink/afdf94b770b0458a8440a335daf77c4c\", \"name\": \"Oscar Cartwright\", \"nominalAmount\": 256, \"pdf\": \"https://invoice-h.sandbox.starkbank.com/pdf/afdf94b770b0458a8440a335daf77c4c\", \"status\": \"created\", \"tags\": [\"war supply\", \"invoice #1234\"], \"taxId\": \"337.451.076-08\", \"transactionIds\": [], \"updated\": \"2021-04-26T20:16:51.442989+00:00\"}, \"type\": \"created\"}, \"subscription\": \"invoice\", \"workspaceId\": \"5078376503050240\"}}"
	invalidSignature := "signature mal formed"
	parsed := parse.ParseAndVerify(
		content,
		invalidSignature,
		utils.SdkVersion,
		utils.ApiVersion,
		utils.Language,
		utils.Timeout,
		hosts.Infra,
		User.ExampleProjectBank,
		"",
	)
	fmt.Println(parsed)
}
