/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 1.10.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package harbor

// The item in CVE whitelist
type CveWhitelistItem struct {
	// The ID of the CVE, such as \"CVE-2019-10164\"
	CveId string `json:"cve_id,omitempty"`
}
