---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "gocd_cluster_profile Data Source - terraform-provider-gocd"
subcategory: ""
description: |-
  
---

# gocd_cluster_profile (Data Source)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `profile_id` (String) The identifier of the cluster profile.

### Optional

- `etag` (String) Etag used to track the cluster profile
- `plugin_id` (String) The plugin identifier of the cluster profile.
- `properties` (Block List) The list of configuration properties that represent the configuration of this profile. (see [below for nested schema](#nestedblock--properties))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--properties"></a>
### Nested Schema for `properties`

Required:

- `key` (String) the name of the property key.

Optional:

- `encrypted_value` (String) The encrypted value of the property
- `is_secure` (Boolean) Specify whether the given property is secure or not. If true and encrypted_value is not specified, GoCD will store the value in encrypted format.
- `value` (String) The value of the property


