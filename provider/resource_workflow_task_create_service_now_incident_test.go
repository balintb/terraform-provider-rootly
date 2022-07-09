package provider

// This file was auto-generated by tools/gen_tasks.js

import (
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceWorkflowTaskCreateServiceNowIncident(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() {
			testAccPreCheck(t)
			time.Sleep(1 * time.Second)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceWorkflowTaskCreateServiceNowIncident,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("rootly_workflow_incident.foo", "name", "test-workflow"),
				),
			},
			{
				Config: testAccResourceWorkflowTaskCreateServiceNowIncidentUpdate,
			},
		},
	})
}

const testAccResourceWorkflowTaskCreateServiceNowIncident = `
resource "rootly_workflow_incident" "foo" {
  name = "test-workflow"
	trigger_params {
		triggers = ["incident_updated"]
	}
}

resource "rootly_workflow_task_create_service_now_incident" "foo" {
	workflow_id = rootly_workflow_incident.foo.id
	task_params {
		title = "test"
	}
}
`

const testAccResourceWorkflowTaskCreateServiceNowIncidentUpdate = `
resource "rootly_workflow_incident" "foo" {
  name = "test-workflow"
	trigger_params {
		triggers = ["incident_updated"]
	}
}

resource "rootly_workflow_task_create_service_now_incident" "foo" {
	workflow_id = rootly_workflow_incident.foo.id
	task_params {
		title = "test"
	}
}
`
