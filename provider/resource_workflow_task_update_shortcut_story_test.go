package provider

// This file was auto-generated by tools/gen_tasks.js

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccResourceWorkflowTaskUpdateShortcutStory(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceWorkflowTaskUpdateShortcutStory,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("rootly_workflow_incident.foo", "name", "test-workflow"),
				),
			},
			{
				Config: testAccResourceWorkflowTaskUpdateShortcutStoryUpdate,
			},
		},
	})
}

const testAccResourceWorkflowTaskUpdateShortcutStory = `
resource "rootly_workflow_incident" "foo" {
  	name = "test-workflow"
	trigger_params {
		triggers = ["incident_updated"]
	}
}

resource "rootly_workflow_task_update_shortcut_story" "foo" {
	workflow_id = rootly_workflow_incident.foo.id
	task_params {
		story_id = "test"
archivation = {
					id = "foo"
					name = "bar"
				}
	}
}
`

const testAccResourceWorkflowTaskUpdateShortcutStoryUpdate = `
resource "rootly_workflow_incident" "foo" {
  	name = "test-workflow"
	trigger_params {
		triggers = ["incident_updated"]
	}
}

resource "rootly_workflow_task_update_shortcut_story" "foo" {
	workflow_id = rootly_workflow_incident.foo.id
	task_params {
		story_id = "test"
archivation = {
					id = "foo"
					name = "bar"
				}
	}
}
`
