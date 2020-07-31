/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 1.10.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package harbor

type ProjectMetadata struct {
	// The public status of the project. The valid values are \"true\", \"false\".
	Public string `json:"public,omitempty"`
	// Whether content trust is enabled or not. If it is enabled, user can't pull unsigned images from this project. The valid values are \"true\", \"false\".
	EnableContentTrust string `json:"enable_content_trust,omitempty"`
	// Whether prevent the vulnerable images from running. The valid values are \"true\", \"false\".
	PreventVul string `json:"prevent_vul,omitempty"`
	// If the vulnerability is high than severity defined here, the images can't be pulled. The valid values are \"none\", \"low\", \"medium\", \"high\", \"critical\".
	Severity string `json:"severity,omitempty"`
	// Whether scan images automatically when pushing. The valid values are \"true\", \"false\".
	AutoScan string `json:"auto_scan,omitempty"`
	// Whether this project reuse the system level CVE whitelist as the whitelist of its own.  The valid values are \"true\", \"false\". If it is set to \"true\" the actual whitelist associate with this project, if any, will be ignored.
	ReuseSysCveWhitelist string `json:"reuse_sys_cve_whitelist,omitempty"`
}
