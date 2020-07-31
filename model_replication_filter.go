/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 1.10.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package github.com/angeiv/harbor

type ReplicationFilter struct {
	// The replication policy filter type.
	Type_ string `json:"type,omitempty"`
	// The value of replication policy filter.
	Value string `json:"value,omitempty"`
}
