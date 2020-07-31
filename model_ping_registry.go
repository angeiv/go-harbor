/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 1.10.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package harbor

type PingRegistry struct {
	// The ID of the registry
	Id int32 `json:"id,omitempty"`
	// Type of the registry, e.g. 'harbor'.
	Type_ string `json:"type,omitempty"`
	// The registry address URL string.
	Url string `json:"url,omitempty"`
	// Credential type of the registry, e.g. 'basic'.
	CredentialType string `json:"credential_type,omitempty"`
	// The registry access key.
	AccessKey string `json:"access_key,omitempty"`
	// The registry access secret.
	AccessSecret string `json:"access_secret,omitempty"`
	// Whether or not the certificate will be verified when Harbor tries to access the server.
	Insecure bool `json:"insecure,omitempty"`
}
