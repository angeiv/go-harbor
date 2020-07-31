/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 1.10.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package harbor

type LdapImportUsers struct {
	// selected uid list
	LdapUidList []string `json:"ldap_uid_list,omitempty"`
}