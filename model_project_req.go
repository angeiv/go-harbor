/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 1.10.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package harbor

type ProjectReq struct {
	// The name of the project.
	ProjectName string `json:"project_name,omitempty"`
	Metadata *ProjectMetadata `json:"metadata,omitempty"`
	CveWhitelist *CveWhitelist `json:"cve_whitelist,omitempty"`
	// The count quota of the project.
	CountLimit int64 `json:"count_limit,omitempty"`
	// The storage quota of the project.
	StorageLimit int64 `json:"storage_limit,omitempty"`
}