/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 1.10.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package github.com/angeiv/harbor

// The harbor native report format
type Report struct {
	// Time of generating this report
	GeneratedAt string `json:"generated_at,omitempty"`
	// A standard scale for measuring the severity of a vulnerability.
	Severity string `json:"severity,omitempty"`
	Vulnerabilities []VulnerabilityItem `json:"vulnerabilities,omitempty"`
	Scanner *Scanner `json:"scanner,omitempty"`
}
