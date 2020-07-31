/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 1.10.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package github.com/angeiv/harbor

type RobotAccountAccess struct {
	// the resource of harbor
	Resource string `json:"resource,omitempty"`
	// the action to resource that perdefined in harbor rbac
	Action string `json:"action,omitempty"`
}