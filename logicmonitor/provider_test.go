package logicmonitor

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"logicmonitor": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("LM_API_ID"); v == "" {
		t.Fatal("LM_API_ID must be set for acceptance tests")
	}
	if v := os.Getenv("LM_API_KEY"); v == "" {
		t.Fatal("LM_API_KEY must be set for acceptance tests")
	}
	if v := os.Getenv("LM_COMPANY"); v == "" {
		t.Fatal("LM_COMPANY must be set for acceptance tests")
	}
}
