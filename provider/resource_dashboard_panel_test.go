package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceDashboardPanel(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceDashboardPanel,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("rootly_dashboard_panel.foo", "name", "test"),
				),
			},
		},
	})
}

const testAccResourceDashboardPanel = `
resource "rootly_dashboard" "foo" {
  name = "mydashboard_panel"
}

resource "rootly_dashboard_panel" "foo" {
  dashboard_id = rootly_dashboard.foo.id
	name = "test"
	params {
		display = "line_chart"
		datasets {
			collection = "incidents"
			filter {
				operation = "and"
				rules {
					operation = "and"
					condition = "="
					key = "status"
					value = "started"
				}
			}
			group_by = "severity"
			aggregate {
				cumulative = false
				key = "results"
				operation = "count"
			}
		}
	}
}
`
