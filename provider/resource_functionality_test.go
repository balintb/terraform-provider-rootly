package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceFunctionality(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceFunctionality,
			},
		},
	})
}

const testAccResourceFunctionality = `
resource "rootly_functionality" "test" {
	name = "test"
}
`
