package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/rootlyhq/terraform-provider-rootly/client"
)

func resourcePlaybookTask() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePlaybookTaskCreate,
		ReadContext:   resourcePlaybookTaskRead,
		UpdateContext: resourcePlaybookTaskUpdate,
		DeleteContext: resourcePlaybookTaskDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{

			"playbook_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Required:    false,
				Optional:    true,
				ForceNew:    true,
				Description: "",
			},

			"task": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    false,
				Required:    true,
				Optional:    false,
				ForceNew:    false,
				Description: "The task of the incident task",
			},

			"description": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Required:    false,
				Optional:    true,
				ForceNew:    false,
				Description: "The description of incident task",
			},
		},
	}
}

func resourcePlaybookTaskCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)

	tflog.Trace(ctx, fmt.Sprintf("Creating PlaybookTask"))

	s := &client.PlaybookTask{}

	if value, ok := d.GetOkExists("playbook_id"); ok {
		s.PlaybookId = value.(string)
	}
	if value, ok := d.GetOkExists("task"); ok {
		s.Task = value.(string)
	}
	if value, ok := d.GetOkExists("description"); ok {
		s.Description = value.(string)
	}

	res, err := c.CreatePlaybookTask(s)
	if err != nil {
		return diag.Errorf("Error creating playbook_task: %s", err.Error())
	}

	d.SetId(res.ID)
	tflog.Trace(ctx, fmt.Sprintf("created a playbook_task resource: %s", d.Id()))

	return resourcePlaybookTaskRead(ctx, d, meta)
}

func resourcePlaybookTaskRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)
	tflog.Trace(ctx, fmt.Sprintf("Reading PlaybookTask: %s", d.Id()))

	item, err := c.GetPlaybookTask(d.Id())
	if err != nil {
		// In the case of a NotFoundError, it means the resource may have been removed upstream
		// We just remove it from the state.
		if _, ok := err.(client.NotFoundError); ok && !d.IsNewResource() {
			tflog.Warn(ctx, fmt.Sprintf("PlaybookTask (%s) not found, removing from state", d.Id()))
			d.SetId("")
			return nil
		}

		return diag.Errorf("Error reading playbook_task: %s", d.Id())
	}

	d.Set("playbook_id", item.PlaybookId)
	d.Set("task", item.Task)
	d.Set("description", item.Description)

	return nil
}

func resourcePlaybookTaskUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)
	tflog.Trace(ctx, fmt.Sprintf("Updating PlaybookTask: %s", d.Id()))

	s := &client.PlaybookTask{}

	if d.HasChange("playbook_id") {
		s.PlaybookId = d.Get("playbook_id").(string)
	}
	if d.HasChange("task") {
		s.Task = d.Get("task").(string)
	}
	if d.HasChange("description") {
		s.Description = d.Get("description").(string)
	}

	_, err := c.UpdatePlaybookTask(d.Id(), s)
	if err != nil {
		return diag.Errorf("Error updating playbook_task: %s", err.Error())
	}

	return resourcePlaybookTaskRead(ctx, d, meta)
}

func resourcePlaybookTaskDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)
	tflog.Trace(ctx, fmt.Sprintf("Deleting PlaybookTask: %s", d.Id()))

	err := c.DeletePlaybookTask(d.Id())
	if err != nil {
		// In the case of a NotFoundError, it means the resource may have been removed upstream.
		// We just remove it from the state.
		if _, ok := err.(client.NotFoundError); ok && !d.IsNewResource() {
			tflog.Warn(ctx, fmt.Sprintf("PlaybookTask (%s) not found, removing from state", d.Id()))
			d.SetId("")
			return nil
		}
		return diag.Errorf("Error deleting playbook_task: %s", err.Error())
	}

	d.SetId("")

	return nil
}
