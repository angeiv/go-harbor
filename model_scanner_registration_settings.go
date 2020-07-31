/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 1.10.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package github.com/angeiv/harbor

type ScannerRegistrationSettings struct {
	// The name of this registration
	Name string `json:"name,omitempty"`
	// A base URL of the scanner adapter.
	Url string `json:"url,omitempty"`
	// Specify what authentication approach is adopted for the HTTP communications. Supported types Basic\", \"Bearer\" and api key header \"X-ScannerAdapter-API-Key\" 
	Auth string `json:"auth,omitempty"`
	// An optional value of the HTTP Authorization header sent with each request to the Scanner Adapter API. 
	AccessCredential string `json:"access_credential,omitempty"`
}
