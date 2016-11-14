package providers

import (
	"github.com/clstokes/aero/structs"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	GOOGLE_DIR_FIXTURES = "test-fixtures/google/"
)

var GOOGLE_TEST_DATA = map[string]string{
	structs.KEY_ADDRESS_PRIVATE: "10.240.0.2",
	structs.KEY_ADDRESS_PUBLIC:  "104.155.153.127",
	structs.KEY_INSTANCE_NAME:   "main",
	structs.KEY_PROVIDER:        "google",
	structs.KEY_REGION:          "us-central1",
	structs.KEY_ZONE:            "us-central1-a",
}

func TestGoogleName(t *testing.T) {
	server, provider := getGoogleServerAndProvider()
	defer server.Close()

	value := provider.Name()
	testValue := structs.NAME_GOOGLE
	if value != testValue {
		t.Fatalf("Name from provider [%s] did not match constant [%s].", value, testValue)
	}
}

func TestGoogleIsCurrentProvider(t *testing.T) {
	server, provider := getGoogleServerAndProvider()
	defer server.Close()

	if !provider.IsCurrentProvider() {
		t.Fatalf("Not set as current provider.")
	}
}

func TestGoogleReadAll(t *testing.T) {
	server, provider := getGoogleServerAndProvider()
	defer server.Close()

	for key, testValue := range GOOGLE_TEST_DATA {
		value, err := provider.Read(key)
		if err != nil {
			t.Fatalf("err: %s", err)
		}

		if value != string(testValue) {
			t.Fatalf("Value from provider [%s] did not match test data [%s].", value, testValue)
		}
	}

}

func getGoogleServerAndProvider() (*httptest.Server, structs.Provider) {
	server := httptest.NewServer(http.FileServer(http.Dir(GOOGLE_DIR_FIXTURES)))
	mapping := structs.ProviderMapping{
		MetadataAddress: server.URL,
	}
	return server, InitGoogle(mapping)
}
