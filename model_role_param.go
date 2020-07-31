/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 1.10.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package github.com/angeiv/harbor

type RoleParam struct {
	// Role ID for updating project role member.
	Roles []int32 `json:"roles,omitempty"`
	// Username relevant to a project role member.
	Username string `json:"username,omitempty"`
}
