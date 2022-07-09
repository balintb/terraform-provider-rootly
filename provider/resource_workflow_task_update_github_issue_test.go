package provider

// This file was auto-generated by tools/gen_tasks.js

import (
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceWorkflowTaskUpdateGithubIssue(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() {
			testAccPreCheck(t)
			time.Sleep(1 * time.Second)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceWorkflowTaskUpdateGithubIssue,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("rootly_workflow_incident.foo", "name", "test-workflow"),
				),
			},
			{
				Config: testAccResourceWorkflowTaskUpdateGithubIssueUpdate,
			},
		},
	})
}

const testAccResourceWorkflowTaskUpdateGithubIssue = `
resource "rootly_workflow_incident" "foo" {
  name = "test-workflow"
	trigger_params {
		triggers = ["incident_updated"]
	}
}

resource "rootly_workflow_task_update_github_issue" "foo" {
	workflow_id = rootly_workflow_incident.foo.id
	task_params {
		issue_id = "test"
completion = {
					id = "foo"
					name = "bar"
				}
	}
}
`

const testAccResourceWorkflowTaskUpdateGithubIssueUpdate = `
resource "rootly_workflow_incident" "foo" {
  name = "test-workflow"
	trigger_params {
		triggers = ["incident_updated"]
	}
}

resource "rootly_workflow_task_update_github_issue" "foo" {
	workflow_id = rootly_workflow_incident.foo.id
	task_params {
		issue_id = "test"
completion = {
					id = "foo"
					name = "bar"
				}
	}
}
`
