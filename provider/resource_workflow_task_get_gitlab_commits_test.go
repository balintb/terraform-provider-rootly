package provider

// This file was auto-generated by tools/gen_tasks.js

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccResourceWorkflowTaskGetGitlabCommits(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceWorkflowTaskGetGitlabCommits,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("rootly_workflow_incident.foo", "name", "test-workflow"),
				),
			},
			{
				Config: testAccResourceWorkflowTaskGetGitlabCommitsUpdate,
			},
		},
	})
}

const testAccResourceWorkflowTaskGetGitlabCommits = `
resource "rootly_workflow_incident" "foo" {
  	name = "test-workflow"
	trigger_params {
		triggers = ["incident_updated"]
	}
}

resource "rootly_workflow_task_get_gitlab_commits" "foo" {
	workflow_id = rootly_workflow_incident.foo.id
	task_params {
		branch = "test"
past_duration = "1 hour"
	}
}
`

const testAccResourceWorkflowTaskGetGitlabCommitsUpdate = `
resource "rootly_workflow_incident" "foo" {
  	name = "test-workflow"
	trigger_params {
		triggers = ["incident_updated"]
	}
}

resource "rootly_workflow_task_get_gitlab_commits" "foo" {
	workflow_id = rootly_workflow_incident.foo.id
	task_params {
		branch = "test"
past_duration = "1 hour"
	}
}
`
