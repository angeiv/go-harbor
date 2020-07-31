/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 1.10.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package github.com/angeiv/harbor

// The parameters of the policy, the values are dependant on the type of the policy.
type ConfigurationsScanAllPolicyParameter struct {
	// The offset in seconds of UTC 0 o'clock, only valid when the policy type is \"daily\"
	DailyTime int32 `json:"daily_time,omitempty"`
}
