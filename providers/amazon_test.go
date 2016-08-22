package providers

import (
	"github.com/clstokes/aero/structs"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRead(t *testing.T) {
	server := httptest.NewServer(http.FileServer(http.Dir("test-fixtures/amazon")))
	defer server.Close()

	baseURL := server.URL
	amazon := getProvider(baseURL)

	instance, err := amazon.Read(structs.KEY_INSTANCE_NAME)
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	// TODO: Test all metadata keys.
	instanceFromFile, err := ioutil.ReadFile("test-fixtures/amazon/latest/meta-data/instance-id")
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if instance != string(instanceFromFile) {
		t.Fatalf("Instance value [%s] incorrect.", instance)
	}
}

func getProvider(url string) structs.Provider {
	mapping := structs.ProviderMapping{
		MetadataAddress: url,
	}
	return InitAmazon(mapping)
}
