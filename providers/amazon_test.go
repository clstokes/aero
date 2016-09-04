package providers

import (
	"github.com/clstokes/aero/structs"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	DIR_FIXTURES = "test-fixtures/amazon/"
)

var TEST_DATA = map[string]string{
	structs.KEY_ADDRESS_PRIVATE: "172.31.62.58",
	structs.KEY_ADDRESS_PUBLIC:  "52.207.213.249",
	structs.KEY_INSTANCE_NAME:   "i-0849c4a43b144116d",
	structs.KEY_PROVIDER:        "amazon",
	structs.KEY_REGION:          "us-east-1",
	structs.KEY_ZONE:            "us-east-1d",
}

func TestName(t *testing.T) {
	server, amazon := getServerAndProvider()
	defer server.Close()

	value := amazon.Name()
	if value != structs.NAME_AMAZON {
		t.Fatalf("Name from provider [%s] did not match constant [%s].", value, structs.NAME_AMAZON)
	}
}

func TestIsCurrentProvider(t *testing.T) {
	server, amazon := getServerAndProvider()
	defer server.Close()

	if !amazon.IsCurrentProvider() {
		t.Fatalf("Not set as current provider.")
	}
}

func TestRead(t *testing.T) {
	server, amazon := getServerAndProvider()
	defer server.Close()

	for key, testValue := range TEST_DATA {
		value, err := amazon.Read(key)
		if err != nil {
			t.Fatalf("err: %s", err)
		}

		if value != string(testValue) {
			t.Fatalf("Value from provider [%s] did not match test data [%s].", value, testValue)
		}
	}

}

func getServerAndProvider() (*httptest.Server, structs.Provider) {
	server := httptest.NewServer(http.FileServer(http.Dir(DIR_FIXTURES)))
	mapping := structs.ProviderMapping{
		MetadataAddress: server.URL,
	}
	return server, InitAmazon(mapping)
}
