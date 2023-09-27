package provider

// This file was auto-generated by tools/gen_tasks.js

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccResourceWorkflowTaskInviteToSlackChannelPagerduty(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceWorkflowTaskInviteToSlackChannelPagerduty,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("rootly_workflow_incident.foo", "name", "test-workflow"),
				),
			},
			{
				Config: testAccResourceWorkflowTaskInviteToSlackChannelPagerdutyUpdate,
			},
		},
	})
}

const testAccResourceWorkflowTaskInviteToSlackChannelPagerduty = `
resource "rootly_workflow_incident" "foo" {
  	name = "test-workflow"
	trigger_params {
		triggers = ["incident_updated"]
	}
}

resource "rootly_workflow_task_invite_to_slack_channel_pagerduty" "foo" {
	workflow_id = rootly_workflow_incident.foo.id
	task_params {
		
	}
}
`

const testAccResourceWorkflowTaskInviteToSlackChannelPagerdutyUpdate = `
resource "rootly_workflow_incident" "foo" {
  	name = "test-workflow"
	trigger_params {
		triggers = ["incident_updated"]
	}
}

resource "rootly_workflow_task_invite_to_slack_channel_pagerduty" "foo" {
	workflow_id = rootly_workflow_incident.foo.id
	task_params {
		
	}
}
`
