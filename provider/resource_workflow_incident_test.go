package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceWorkflowIncident(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceWorkflowIncident,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("rootly_workflow_incident.foo3", "name", "test-incident-workflow3"),
				),
			},
			{
				Config: testAccResourceWorkflowIncidentUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("rootly_workflow_incident.foo3", "name", "test-incident-workflow3"),
				),
			},
		},
	})
}

const testAccResourceWorkflowIncident = `
resource "rootly_workflow_incident" "foo1" {
  name = "test-incident-workflow1"
	trigger_params {
		triggers = ["incident_updated"]
	}
}
resource "rootly_workflow_incident" "foo2" {
  name = "test-incident-workflow2"
	trigger_params {
		triggers = ["incident_updated"]
	}
	depends_on = [rootly_workflow_incident.foo1]
}
resource "rootly_workflow_incident" "foo3" {
  name = "test-incident-workflow3"
	trigger_params {
		triggers = ["incident_updated"]
	}
	depends_on =[rootly_workflow_incident.foo2]
}
`

const testAccResourceWorkflowIncidentUpdate = `
resource "rootly_workflow_incident" "foo3" {
  name = "test-incident-workflow3"
	trigger_params {
		triggers = ["incident_updated"]
	}
}
`
