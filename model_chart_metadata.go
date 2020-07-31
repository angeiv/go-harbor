/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 1.10.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package harbor

// The metadata of chart version
type ChartMetadata struct {
	// The name of the chart
	Name string `json:"name"`
	// The URL to the relevant project page
	Home string `json:"home,omitempty"`
	// The URL to the source code of chart
	Sources []string `json:"sources,omitempty"`
	// A SemVer 2 version of chart
	Version string `json:"version"`
	// A one-sentence description of chart
	Description string `json:"description,omitempty"`
	// A list of string keywords
	Keywords []string `json:"keywords,omitempty"`
	// The name of template engine
	Engine string `json:"engine"`
	// The URL to an icon file
	Icon string `json:"icon"`
	// The API version of this chart
	ApiVersion string `json:"apiVersion"`
	// The version of the application enclosed in the chart
	AppVersion string `json:"appVersion"`
	// Whether or not this chart is deprecated
	Deprecated bool `json:"deprecated,omitempty"`
}
