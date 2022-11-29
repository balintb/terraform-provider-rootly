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

func resourceWorkflowTaskUpdatePagerdutyIncident() *schema.Resource {
	return &schema.Resource{
		Description: "Manages workflow update_pagerduty_incident task.",

		CreateContext: resourceWorkflowTaskUpdatePagerdutyIncidentCreate,
		ReadContext:   resourceWorkflowTaskUpdatePagerdutyIncidentRead,
		UpdateContext: resourceWorkflowTaskUpdatePagerdutyIncidentUpdate,
		DeleteContext: resourceWorkflowTaskUpdatePagerdutyIncidentDelete,
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
							Default: "update_pagerduty_incident",
							ValidateFunc: validation.StringInSlice([]string{
								"update_pagerduty_incident",
							}, false),
						},
						"pagerduty_incident_id": &schema.Schema{
							Description: "Pagerduty incident id",
							Type: schema.TypeString,
							Required: true,
						},
						"title": &schema.Schema{
							Description: "Title to update to",
							Type: schema.TypeString,
							Optional: true,
						},
						"status": &schema.Schema{
							Description: ". Value must be one of `resolved`, `acknowledged`, `auto`.",
							Type: schema.TypeString,
							Optional: true,
							Default: nil,
							ValidateFunc: validation.StringInSlice([]string{
								"resolved",
"acknowledged",
"auto",
							}, false),
						},
						"resolution": &schema.Schema{
							Description: "A message outlining the incident's resolution in Pagerduty",
							Type: schema.TypeString,
							Optional: true,
						},
						"escalation_level": &schema.Schema{
							Description: "Escalation level of policy attached to incident",
							Type: schema.TypeInt,
							Optional: true,
							Default: nil,
						},
						"urgency": &schema.Schema{
							Description: "Pagerduty incident urgency, selecting auto will let Rootly auto map our incident severity.. Value must be one of `high`, `low`, `auto`.",
							Type: schema.TypeString,
							Optional: true,
							Default: nil,
							ValidateFunc: validation.StringInSlice([]string{
								"high",
"low",
"auto",
							}, false),
						},
					},
				},
			},
		},
	}
}

func resourceWorkflowTaskUpdatePagerdutyIncidentCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	return resourceWorkflowTaskUpdatePagerdutyIncidentRead(ctx, d, meta)
}

func resourceWorkflowTaskUpdatePagerdutyIncidentRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)
	tflog.Trace(ctx, fmt.Sprintf("Reading workflow task: %s", d.Id()))

	res, err := c.GetWorkflowTask(d.Id())
	if err != nil {
		// In the case of a NotFoundError, it means the resource may have been removed upstream
		// We just remove it from the state.
		if _, ok := err.(client.NotFoundError); ok && !d.IsNewResource() {
			tflog.Warn(ctx, fmt.Sprintf("WorkflowTaskUpdatePagerdutyIncident (%s) not found, removing from state", d.Id()))
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

func resourceWorkflowTaskUpdatePagerdutyIncidentUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	return resourceWorkflowTaskUpdatePagerdutyIncidentRead(ctx, d, meta)
}

func resourceWorkflowTaskUpdatePagerdutyIncidentDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)
	tflog.Trace(ctx, fmt.Sprintf("Deleting workflow task: %s", d.Id()))

	err := c.DeleteWorkflowTask(d.Id())
	if err != nil {
		// In the case of a NotFoundError, it means the resource may have been removed upstream.
		// We just remove it from the state.
		if _, ok := err.(client.NotFoundError); ok && !d.IsNewResource() {
			tflog.Warn(ctx, fmt.Sprintf("WorkflowTaskUpdatePagerdutyIncident (%s) not found, removing from state", d.Id()))
			d.SetId("")
			return nil
		}
		return diag.Errorf("Error deleting workflow task: %s", err.Error())
	}

	d.SetId("")

	return nil
}
