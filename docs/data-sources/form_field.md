---
page_title: "Data Source rootly_form_field - terraform-provider-rootly"
subcategory:
description: |-
    
---

# Data Source (rootly_form_field)



## Example Usage

```terraform
data "rootly_form_field" "my-form-field" {
  slug = "my-form-field"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `created_at` (Map of String) Filter by date range using 'lt' and 'gt'.
- `enabled` (Boolean)
- `kind` (String)
- `name` (String)
- `slug` (String)

### Read-Only

- `id` (String) The ID of this resource.
