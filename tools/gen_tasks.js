const fs = require("fs")
const path = require('path')
const schema = require(path.resolve(process.argv[2]))

Object.keys(schema).forEach((key) => {
	const task_name = key.replace("_task_params", "")
	const task_schema = schema[key]
	const task_name_camel = task_name.split("_").map((p) => `${p[0].toUpperCase()}${p.slice(1)}`).join('')
	fs.writeFileSync(`./provider/resource_workflow_task_${task_name}.go`, genResourceFile(task_name, task_schema))
	fs.writeFileSync(`./provider/resource_workflow_task_${task_name}_test.go`, genResourceTestFile(task_name, task_schema))
})

function genResourceFile(task_name, task_schema) {
	const task_name_camel = task_name.split("_").map((p) => `${p[0].toUpperCase()}${p.slice(1)}`).join('')
return `package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/rootlyhq/terraform-provider-rootly/client"
)

func resourceWorkflowTask${task_name_camel}() *schema.Resource {
	return &schema.Resource{
		Description: "Manages workflow ${task_name} task.",

		CreateContext: resourceWorkflowTask${task_name_camel}Create,
		ReadContext:   resourceWorkflowTask${task_name_camel}Read,
		UpdateContext: resourceWorkflowTask${task_name_camel}Update,
		DeleteContext: resourceWorkflowTask${task_name_camel}Delete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"workflow_id": {
				Description:  "The ID of the parent workflow",
				Type:         schema.TypeString,
				Required:     true,
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
							Default: "${task_name}",
							ValidateFunc: validation.StringInSlice([]string{
								"${task_name}",
							}, false),
						},
${Object.keys(task_schema.properties).filter((k) => k !== "task_type").map((key) => genTaskSchemaProperty(key, task_schema.properties[key], task_schema.required)).join('\n')}
					},
				},
			},
		},
	}
}

func resourceWorkflowTask${task_name_camel}Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)

	workflowId := d.Get("workflow_id").(string)
	taskParams := d.Get("task_params").([]interface{})[0].(map[string]interface{})

	tflog.Trace(ctx, fmt.Sprintf("Creating workflow task: %s", workflowId))

	s := &client.WorkflowTask{
		WorkflowId: workflowId,
		TaskParams: taskParams,
	}

	res, err := c.CreateWorkflowTask(s)
	if err != nil {
		return diag.Errorf("Error creating workflow task: %s", err.Error())
	}

	d.SetId(res.ID)
	tflog.Trace(ctx, fmt.Sprintf("created an workflow task resource: %v (%s)", workflowId, d.Id()))

	return resourceWorkflowTask${task_name_camel}Read(ctx, d, meta)
}

func resourceWorkflowTask${task_name_camel}Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)
	tflog.Trace(ctx, fmt.Sprintf("Reading workflow task: %s", d.Id()))

	res, err := c.GetWorkflowTask(d.Id())
	if err != nil {
		// In the case of a NotFoundError, it means the resource may have been removed upstream
		// We just remove it from the state.
		if _, ok := err.(client.NotFoundError); ok && !d.IsNewResource() {
			tflog.Warn(ctx, fmt.Sprintf("WorkflowTask${task_name_camel} (%s) not found, removing from state", d.Id()))
			d.SetId("")
			return nil
		}

		return diag.Errorf("Error reading workflow task: %s", d.Id())
	}

	d.Set("workflow_id", res.WorkflowId)
	tps := make([]interface{}, 1, 1)
	tps[0] = res.TaskParams
	d.Set("task_params", tps)

	return nil
}

func resourceWorkflowTask${task_name_camel}Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)
	tflog.Trace(ctx, fmt.Sprintf("Updating workflow task: %s", d.Id()))

	workflowId := d.Get("workflow_id").(string)
	taskParams := d.Get("task_params").([]interface{})[0].(map[string]interface{})

	s := &client.WorkflowTask{
		WorkflowId: workflowId,
		TaskParams: taskParams,
	}

	tflog.Debug(ctx, fmt.Sprintf("adding value: %#v", s))
	_, err := c.UpdateWorkflowTask(d.Id(), s)
	if err != nil {
		return diag.Errorf("Error updating workflow task: %s", err.Error())
	}

	return resourceWorkflowTask${task_name_camel}Read(ctx, d, meta)
}

func resourceWorkflowTask${task_name_camel}Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)
	tflog.Trace(ctx, fmt.Sprintf("Deleting workflow task: %s", d.Id()))

	err := c.DeleteWorkflowTask(d.Id())
	if err != nil {
		// In the case of a NotFoundError, it means the resource may have been removed upstream.
		// We just remove it from the state.
		if _, ok := err.(client.NotFoundError); ok && !d.IsNewResource() {
			tflog.Warn(ctx, fmt.Sprintf("WorkflowTask${task_name_camel} (%s) not found, removing from state", d.Id()))
			d.SetId("")
			return nil
		}
		return diag.Errorf("Error deleting workflow task: %s", err.Error())
	}

	d.SetId("")

	return nil
}
`
}

function genTaskSchemaProperty(property_name, property_schema, required_props) {
	const isRequired = required_props && required_props.indexOf(property_name) !== -1
	let a = `						"${property_name}": &schema.Schema{
							Description: "${property_schema.description || ""}",
							Type: ${genTaskSchemaPropertyType(property_schema.type)},
							${isRequired ? 'Required' : 'Optional'}: true,`
	if (property_schema.enum) {
		if (!isRequired) {
			a = `${a}
							Default: "${property_schema.enum[0]}",`
			}
		a = `${a}
							ValidateFunc: validation.StringInSlice([]string{
								${property_schema.enum.map((k) => `"${k}",`).join('\n')}
							}, false),`
	}
	if (property_schema.type === "array") {
		if (property_schema.items.type === "string") {
			a = `${a}
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},`
		} else if (property_schema.items.type === "object") {
			a = `${a}
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type: schema.TypeString,
										Required: true,
									},
									"name": &schema.Schema{
										Type: schema.TypeString,
										Required: true,
									},
								},
							},`
		}
	}
return `${a}
						},`
}

function genTaskSchemaPropertyType(type_) {
	switch (type_) {
		case "string":
			return "schema.TypeString"
		case "number":
			return "schema.TypeInt"
		case "array":
			return "schema.TypeList"
		case "object":
			return "schema.TypeMap"
		case "boolean":
			return "schema.TypeBool"
		default:
			return "schema.TypeString"
	}
}

function genResourceTestFile(task_name, task_schema) {
	const task_name_camel = task_name.split("_").map((p) => `${p[0].toUpperCase()}${p.slice(1)}`).join('')
return `package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceWorkflowTask${task_name_camel}(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceWorkflowTask${task_name_camel},
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("rootly_workflow_incident.foo", "name", "test-workflow"),
				),
			},
			{
				Config: testAccResourceWorkflowTask${task_name_camel}Update,
			},
		},
	})
}

const testAccResourceWorkflowTask${task_name_camel} = \`
resource "random_id" "workflow" {
	byte_length = 8
}

resource "rootly_workflow_incident" "foo" {
  name = "test-\${random_id.workflow.hex}"
	trigger_params {
		triggers = ["incident_updated"]
	}
}

resource "rootly_workflow_task_${task_name}" "foo" {
	workflow_id = rootly_workflow_incident.foo.id
	task_params {
		${genTestParams(task_name, task_schema).join("\n")}
	}
}
\`

const testAccResourceWorkflowTask${task_name_camel}Update = \`
resource "random_id" "workflow" {
  byte_length = 8
}

resource "rootly_workflow_incident" "foo" {
  name = "test-\${random_id.workflow.hex}"
	trigger_params {
		triggers = ["incident_updated"]
	}
}

resource "rootly_workflow_task_${task_name}" "foo" {
	workflow_id = rootly_workflow_incident.foo.id
	task_params {
		${genTestParams(task_name, task_schema).join("\n")}
	}
}
\`
`
}

function genTestParams(task_name, task_schema) {
	return (task_schema.required || []).map((key) => {
		console.log(task_name, key)
		let val = task_schema.properties[key].enum ? task_schema.properties[key].enum[0] : "test"
		switch (task_schema.properties[key].type) {
			case "boolean":
				return `${key} = false`
			case "string":
				return `${key} = "${val}"`
			case "number":
				return `${key} = 1`
			case "array":
				if (task_schema.properties[key].items.type === "object" && task_schema.properties[key].items.properties.id) {
					return `${key} {
						id = "foo"
						name = "bar"
					}`
				}
				return `${key} = ["foo"]`
			case "object":
				return `${key} = {
					id = "foo"
					name = "bar"
				}`
		}
	})
}
