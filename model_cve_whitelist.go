/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 1.10.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package github.com/angeiv/harbor

// The CVE Whitelist for system or project
type CveWhitelist struct {
	// ID of the whitelist
	Id int32 `json:"id,omitempty"`
	// ID of the project which the whitelist belongs to.  For system level whitelist this attribute is zero.
	ProjectId int32 `json:"project_id,omitempty"`
	// the time for expiration of the whitelist, in the form of seconds since epoch.  This is an optional attribute, if it's not set the CVE whitelist does not expire.
	ExpiresAt int32 `json:"expires_at,omitempty"`
	Items []CveWhitelistItem `json:"items,omitempty"`
}
