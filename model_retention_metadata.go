/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 1.10.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package harbor

// the tag retention metadata
type RetentionMetadata struct {
	// templates
	Templates []RetentionRuleMetadata `json:"templates,omitempty"`
	// supported scope selectors
	ScopeSelectors []RetentionSelectorMetadata `json:"scope_selectors,omitempty"`
	// supported tag selectors
	TagSelectors []RetentionSelectorMetadata `json:"tag_selectors,omitempty"`
}