package provider

// This file was auto-generated by tools/gen_tasks.js

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccResourceWorkflowTaskHttpClient(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceWorkflowTaskHttpClient,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("rootly_workflow_incident.foo", "name", "test-workflow"),
				),
			},
			{
				Config: testAccResourceWorkflowTaskHttpClientUpdate,
			},
		},
	})
}

const testAccResourceWorkflowTaskHttpClient = `
resource "rootly_workflow_incident" "foo" {
  	name = "test-workflow"
	trigger_params {
		triggers = ["incident_updated"]
	}
}

resource "rootly_workflow_task_http_client" "foo" {
	workflow_id = rootly_workflow_incident.foo.id
	task_params {
		url = "https://example.com/foo.json"
succeed_on_status = "200"
	}
}
`

const testAccResourceWorkflowTaskHttpClientUpdate = `
resource "rootly_workflow_incident" "foo" {
  	name = "test-workflow"
	trigger_params {
		triggers = ["incident_updated"]
	}
}

resource "rootly_workflow_task_http_client" "foo" {
	workflow_id = rootly_workflow_incident.foo.id
	task_params {
		url = "https://example.com/foo.json"
succeed_on_status = "200"
	}
}
`
