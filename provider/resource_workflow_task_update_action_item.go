package provider

// This file was auto-generated by tools/gen_tasks.js

import (
	"context"
	"fmt"

	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/rootlyhq/terraform-provider-rootly/client"
	"github.com/rootlyhq/terraform-provider-rootly/tools"
)

func resourceWorkflowTaskUpdateActionItem() *schema.Resource {
	return &schema.Resource{
		Description: "Manages workflow update_action_item task.",

		CreateContext: resourceWorkflowTaskUpdateActionItemCreate,
		ReadContext:   resourceWorkflowTaskUpdateActionItemRead,
		UpdateContext: resourceWorkflowTaskUpdateActionItemUpdate,
		DeleteContext: resourceWorkflowTaskUpdateActionItemDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"workflow_id": {
				Description: "The ID of the parent workflow",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"name": {
				Description: "Name of the workflow task",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"position": {
				Description: "The position of the workflow task (1 being top of list)",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"skip_on_failure": {
				Description: "Skip workflow task if any failures",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
			"enabled": {
				Description: "Enable/disable this workflow task",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
			},
			"task_params": {
				Description: "The parameters for this workflow task.",
				Type:        schema.TypeList,
				Required:    true,
				MinItems:    1,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"task_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Default:  "update_action_item",
							ValidateFunc: validation.StringInSlice([]string{
								"update_action_item",
							}, false),
						},
						"query_value": &schema.Schema{
							Description: "Value that attribute_to_query_by to uses to match against",
							Type:        schema.TypeString,
							Required:    true,
						},
						"attribute_to_query_by": &schema.Schema{
							Description: "Attribute of the action item to match against. Value must be one of `id`, `jira_issue_id`, `asana_task_id`, `shortcut_task_id`, `linear_issue_id`, `zendesk_ticket_id`, `trello_card_id`, `airtable_record_id`, `shortcut_story_id`, `github_issue_id`, `freshservice_ticket_id`, `freshservice_task_id`.",
							Type:        schema.TypeString,
							Required:    true,
							ValidateFunc: validation.StringInSlice([]string{
								"id",
								"jira_issue_id",
								"asana_task_id",
								"shortcut_task_id",
								"linear_issue_id",
								"zendesk_ticket_id",
								"trello_card_id",
								"airtable_record_id",
								"shortcut_story_id",
								"github_issue_id",
								"freshservice_ticket_id",
								"freshservice_task_id",
							}, false),
						},
						"summary": &schema.Schema{
							Description: "Brief description of the action item",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"assigned_to_user_id": &schema.Schema{
							Description: "[DEPRECATED] Use assigned_to_user attribute instead. The user id this action item is assigned to",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"assigned_to_user": &schema.Schema{
							Description: "Map must contain two fields, `id` and `name`.  The user this action item is assigned to",
							Type:        schema.TypeMap,
							Optional:    true,
						},
						"group_ids": &schema.Schema{
							Description: "",
							Type:        schema.TypeList,
							Optional:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"description": &schema.Schema{
							Description: "The action item description.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"priority": &schema.Schema{
							Description: "The action item priority. Value must be one of `high`, `medium`, `low`.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     nil,
							ValidateFunc: validation.StringInSlice([]string{
								"high",
								"medium",
								"low",
							}, false),
						},
						"status": &schema.Schema{
							Description: "The action item status. Value must be one of `open`, `in_progress`, `cancelled`, `done`.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     nil,
							ValidateFunc: validation.StringInSlice([]string{
								"open",
								"in_progress",
								"cancelled",
								"done",
							}, false),
						},
						"custom_fields_mapping": &schema.Schema{
							Description: "Custom field mappings. Can contain liquid markup and need to be valid JSON",
							Type:        schema.TypeString,
							Optional:    true,
							DiffSuppressFunc: func(k, old string, new string, d *schema.ResourceData) bool {
								var oldJSONAsInterface, newJSONAsInterface interface{}

								if err := json.Unmarshal([]byte(old), &oldJSONAsInterface); err != nil {
									return false
								}

								if err := json.Unmarshal([]byte(new), &newJSONAsInterface); err != nil {
									return false
								}

								return reflect.DeepEqual(oldJSONAsInterface, newJSONAsInterface)
							},
							Default: "{}",
						},
						"post_to_incident_timeline": &schema.Schema{
							Description: "",
							Type:        schema.TypeBool,
							Optional:    true,
						},
					},
				},
			},
		},
	}
}

func resourceWorkflowTaskUpdateActionItemCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)

	workflowId := d.Get("workflow_id").(string)
	name := d.Get("name").(string)
	position := d.Get("position").(int)
	skipOnFailure := tools.Bool(d.Get("skip_on_failure").(bool))
	enabled := tools.Bool(d.Get("enabled").(bool))
	taskParams := d.Get("task_params").([]interface{})[0].(map[string]interface{})

	tflog.Trace(ctx, fmt.Sprintf("Creating workflow task: %s", workflowId))

	s := &client.WorkflowTask{
		WorkflowId:    workflowId,
		Name:          name,
		Position:      position,
		SkipOnFailure: skipOnFailure,
		Enabled:       enabled,
		TaskParams:    taskParams,
	}

	res, err := c.CreateWorkflowTask(s)
	if err != nil {
		return diag.Errorf("Error creating workflow task: %s", err.Error())
	}

	d.SetId(res.ID)
	tflog.Trace(ctx, fmt.Sprintf("created an workflow task resource: %v (%s)", workflowId, d.Id()))

	return resourceWorkflowTaskUpdateActionItemRead(ctx, d, meta)
}

func resourceWorkflowTaskUpdateActionItemRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)
	tflog.Trace(ctx, fmt.Sprintf("Reading workflow task: %s", d.Id()))

	res, err := c.GetWorkflowTask(d.Id())
	if err != nil {
		// In the case of a NotFoundError, it means the resource may have been removed upstream
		// We just remove it from the state.
		if _, ok := err.(client.NotFoundError); ok && !d.IsNewResource() {
			tflog.Warn(ctx, fmt.Sprintf("WorkflowTaskUpdateActionItem (%s) not found, removing from state", d.Id()))
			d.SetId("")
			return nil
		}

		return diag.Errorf("Error reading workflow task: %s", d.Id())
	}

	d.Set("workflow_id", res.WorkflowId)
	d.Set("name", res.Name)
	d.Set("position", res.Position)
	d.Set("skip_on_failure", res.SkipOnFailure)
	d.Set("enabled", res.Enabled)
	tps := make([]interface{}, 1, 1)
	tps[0] = res.TaskParams
	d.Set("task_params", tps)

	return nil
}

func resourceWorkflowTaskUpdateActionItemUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)
	tflog.Trace(ctx, fmt.Sprintf("Updating workflow task: %s", d.Id()))

	workflowId := d.Get("workflow_id").(string)
	name := d.Get("name").(string)
	position := d.Get("position").(int)
	skipOnFailure := tools.Bool(d.Get("skip_on_failure").(bool))
	enabled := tools.Bool(d.Get("enabled").(bool))
	taskParams := d.Get("task_params").([]interface{})[0].(map[string]interface{})

	s := &client.WorkflowTask{
		WorkflowId:    workflowId,
		Name:          name,
		Position:      position,
		SkipOnFailure: skipOnFailure,
		Enabled:       enabled,
		TaskParams:    taskParams,
	}

	tflog.Debug(ctx, fmt.Sprintf("adding value: %#v", s))
	_, err := c.UpdateWorkflowTask(d.Id(), s)
	if err != nil {
		return diag.Errorf("Error updating workflow task: %s", err.Error())
	}

	return resourceWorkflowTaskUpdateActionItemRead(ctx, d, meta)
}

func resourceWorkflowTaskUpdateActionItemDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)
	tflog.Trace(ctx, fmt.Sprintf("Deleting workflow task: %s", d.Id()))

	err := c.DeleteWorkflowTask(d.Id())
	if err != nil {
		// In the case of a NotFoundError, it means the resource may have been removed upstream.
		// We just remove it from the state.
		if _, ok := err.(client.NotFoundError); ok && !d.IsNewResource() {
			tflog.Warn(ctx, fmt.Sprintf("WorkflowTaskUpdateActionItem (%s) not found, removing from state", d.Id()))
			d.SetId("")
			return nil
		}
		return diag.Errorf("Error deleting workflow task: %s", err.Error())
	}

	d.SetId("")

	return nil
}
