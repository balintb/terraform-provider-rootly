package provider

// This file was auto-generated by tools/gen_tasks.js

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccResourceWorkflowTaskCreateDropboxPaperPage(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceWorkflowTaskCreateDropboxPaperPage,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("rootly_workflow_incident.foo", "name", "test-workflow"),
				),
			},
			{
				Config: testAccResourceWorkflowTaskCreateDropboxPaperPageUpdate,
			},
		},
	})
}

const testAccResourceWorkflowTaskCreateDropboxPaperPage = `
resource "rootly_workflow_incident" "foo" {
  	name = "test-workflow"
	trigger_params {
		triggers = ["incident_updated"]
	}
}

resource "rootly_workflow_task_create_dropbox_paper_page" "foo" {
	workflow_id = rootly_workflow_incident.foo.id
	task_params {
		title = "test"
	}
}
`

const testAccResourceWorkflowTaskCreateDropboxPaperPageUpdate = `
resource "rootly_workflow_incident" "foo" {
  	name = "test-workflow"
	trigger_params {
		triggers = ["incident_updated"]
	}
}

resource "rootly_workflow_task_create_dropbox_paper_page" "foo" {
	workflow_id = rootly_workflow_incident.foo.id
	task_params {
		title = "test"
	}
}
`
