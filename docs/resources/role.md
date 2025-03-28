---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "gocd_role Resource - terraform-provider-gocd"
subcategory: ""
description: |-

---

# gocd_role (Resource)
Creates role in GoCD with all below passed parameters by interacting with GoCD roles [api](https://api.gocd.org/current/#roles).

## Example Usage
```terraform
# For role type gocd
resource "gocd_role" "sample" {
    name  = "sample"
    type  = "gocd"
    users = ["nikhil"]
    policy = [{
        "permission" : "allow",
        "action" : "administer",
        "type" : "*",
        "resource" : "*"
    }]
}

# For role type plugin
resource "gocd_role" "sample_ldap" {
    name           = "sample-ldap"
    type           = "plugin"
    auth_config_id = "ldap-config"
    policy = [{
        "permission" : "allow",
        "action" : "administer",
        "type" : "*",
        "resource" : "*"
    }]
    properties {
        key   = "UserGroupMembershipAttribute"
        value = "testing"
    }
    properties {
        key   = "GroupIdentifiers"
        value = "CN=opts,OU=Groups,OU=TESTCOM,DC=TESTCOM,DC=COM"
    }
    properties {
        key   = "GroupSearchBases"
        value = "OU=Groups,OU=TESTCOM,DC=TESTCOM,DC=COM"
    }
    properties {
        key   = "GroupMembershipFilter"
        value = "(&(member={dn})(cn=opts))"
    }
}
```


<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) The name of the role.
- `policy` (List of Map of String) Policy is fine-grained permissions attached to the users belonging to the current role.
- `type` (String) Type of the role. Use GoCD to create core role and plugin to create plugin role.

### Optional

- `auth_config_id` (String) The authorization configuration identifier.
- `etag` (String) Etag used to track the role.
- `properties` (Block Set) The list of configuration properties that represent the configuration of the profile. (see [below for nested schema](#nestedblock--properties)).
- `system_admin` (Boolean) Enable if the role should be set as admin.
- `users` (List of String) The list of users belongs to the role.

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--properties"></a>
### Nested Schema for `properties`

Required:

- `key` (String) the name of the property key.

Optional:

- `encrypted_value` (String) The encrypted value of the property.
- `is_secure` (Boolean) Specify whether the given property is secure or not. If true and encrypted_value is not specified, GoCD will store the value in encrypted format.
- `value` (String) The value of the property.


