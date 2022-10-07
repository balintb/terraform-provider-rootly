package provider

// This file was auto-generated by tools/gen_tasks.js

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/rootlyhq/terraform-provider-rootly/client"
)

func resourceWorkflowTaskCreateJiraIssue() *schema.Resource {
	return &schema.Resource{
		Description: "Manages workflow create_jira_issue task.",

		CreateContext: resourceWorkflowTaskCreateJiraIssueCreate,
		ReadContext:   resourceWorkflowTaskCreateJiraIssueRead,
		UpdateContext: resourceWorkflowTaskCreateJiraIssueUpdate,
		DeleteContext: resourceWorkflowTaskCreateJiraIssueDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"workflow_id": {
				Description:  "The ID of the parent workflow",
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
			},
			"position": {
				Description:  "The position of the workflow task (1 being top of list)",
				Type:         schema.TypeInt,
				Optional:     true,
				Computed:     true,
			},
			"task_params": {
				Description: "The parameters for this workflow task.",
				Type: schema.TypeList,
				Required: true,
				MinItems: 1,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"task_type": &schema.Schema{
							Type: schema.TypeString,
							Optional: true,
							Default: "create_jira_issue",
							ValidateFunc: validation.StringInSlice([]string{
								"create_jira_issue",
							}, false),
						},
						"title": &schema.Schema{
							Description: "The issue title.",
							Type: schema.TypeString,
							Required: true,
						},
						"description": &schema.Schema{
							Description: "The issue description.",
							Type: schema.TypeString,
							Optional: true,
						},
						"labels": &schema.Schema{
							Description: "The issue labels.",
							Type: schema.TypeString,
							Optional: true,
						},
						"assign_user_email": &schema.Schema{
							Description: "The assigned user's email.",
							Type: schema.TypeString,
							Optional: true,
						},
						"reporter_user_email": &schema.Schema{
							Description: "The reporter user's email.",
							Type: schema.TypeString,
							Optional: true,
						},
						"project_key": &schema.Schema{
							Description: "The project key.",
							Type: schema.TypeString,
							Required: true,
						},
						"due_date": &schema.Schema{
							Description: "The due date.",
							Type: schema.TypeString,
							Optional: true,
						},
						"issue_type": &schema.Schema{
							Description: "Map must contain two fields, `id` and `name`. The issue type id and display name.",
							Type: schema.TypeMap,
							Required: true,
						},
						"priority": &schema.Schema{
							Description: "Map must contain two fields, `id` and `name`. The priority id and display name.",
							Type: schema.TypeMap,
							Optional: true,
						},
						"status": &schema.Schema{
							Description: "Map must contain two fields, `id` and `name`. The status id and display name.",
							Type: schema.TypeMap,
							Optional: true,
						},
						"custom_fields_mapping": &schema.Schema{
							Description: "Custom field mappings. Can contain liquid markup and need to be valid JSON.",
							Type: schema.TypeString,
							Optional: true,
							Default: "{}",
						},
						"update_payload": &schema.Schema{
							Description: "Update payload. Can contain liquid markup and need to be valid JSON.",
							Type: schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func resourceWorkflowTaskCreateJiraIssueCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)

	workflowId := d.Get("workflow_id").(string)
	position := d.Get("position").(int)
	taskParams := d.Get("task_params").([]interface{})[0].(map[string]interface{})

	tflog.Trace(ctx, fmt.Sprintf("Creating workflow task: %s", workflowId))

	s := &client.WorkflowTask{
		WorkflowId: workflowId,
		Position: position,
		TaskParams: taskParams,
	}

	res, err := c.CreateWorkflowTask(s)
	if err != nil {
		return diag.Errorf("Error creating workflow task: %s", err.Error())
	}

	d.SetId(res.ID)
	tflog.Trace(ctx, fmt.Sprintf("created an workflow task resource: %v (%s)", workflowId, d.Id()))

	return resourceWorkflowTaskCreateJiraIssueRead(ctx, d, meta)
}

func resourceWorkflowTaskCreateJiraIssueRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)
	tflog.Trace(ctx, fmt.Sprintf("Reading workflow task: %s", d.Id()))

	res, err := c.GetWorkflowTask(d.Id())
	if err != nil {
		// In the case of a NotFoundError, it means the resource may have been removed upstream
		// We just remove it from the state.
		if _, ok := err.(client.NotFoundError); ok && !d.IsNewResource() {
			tflog.Warn(ctx, fmt.Sprintf("WorkflowTaskCreateJiraIssue (%s) not found, removing from state", d.Id()))
			d.SetId("")
			return nil
		}

		return diag.Errorf("Error reading workflow task: %s", d.Id())
	}

	d.Set("workflow_id", res.WorkflowId)
	d.Set("position", res.Position)
	tps := make([]interface{}, 1, 1)
	tps[0] = res.TaskParams
	d.Set("task_params", tps)

	return nil
}

func resourceWorkflowTaskCreateJiraIssueUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)
	tflog.Trace(ctx, fmt.Sprintf("Updating workflow task: %s", d.Id()))

	workflowId := d.Get("workflow_id").(string)
	position := d.Get("position").(int)
	taskParams := d.Get("task_params").([]interface{})[0].(map[string]interface{})

	s := &client.WorkflowTask{
		WorkflowId: workflowId,
		Position: position,
		TaskParams: taskParams,
	}

	tflog.Debug(ctx, fmt.Sprintf("adding value: %#v", s))
	_, err := c.UpdateWorkflowTask(d.Id(), s)
	if err != nil {
		return diag.Errorf("Error updating workflow task: %s", err.Error())
	}

	return resourceWorkflowTaskCreateJiraIssueRead(ctx, d, meta)
}

func resourceWorkflowTaskCreateJiraIssueDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)
	tflog.Trace(ctx, fmt.Sprintf("Deleting workflow task: %s", d.Id()))

	err := c.DeleteWorkflowTask(d.Id())
	if err != nil {
		// In the case of a NotFoundError, it means the resource may have been removed upstream.
		// We just remove it from the state.
		if _, ok := err.(client.NotFoundError); ok && !d.IsNewResource() {
			tflog.Warn(ctx, fmt.Sprintf("WorkflowTaskCreateJiraIssue (%s) not found, removing from state", d.Id()))
			d.SetId("")
			return nil
		}
		return diag.Errorf("Error deleting workflow task: %s", err.Error())
	}

	d.SetId("")

	return nil
}
