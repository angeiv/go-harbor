/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 1.10.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package github.com/angeiv/harbor

// The quota object
type Quota struct {
	// ID of the quota
	Id int32 `json:"id,omitempty"`
	Ref *ModelMap `json:"ref,omitempty"`
	Hard *ModelMap `json:"hard,omitempty"`
	Used *ModelMap `json:"used,omitempty"`
	// the creation time of the quota
	CreationTime string `json:"creation_time,omitempty"`
	// the update time of the quota
	UpdateTime string `json:"update_time,omitempty"`
}
