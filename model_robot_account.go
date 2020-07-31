/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 1.10.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package harbor

// The object of robot account
type RobotAccount struct {
	// The id of robot account
	Id int32 `json:"id,omitempty"`
	// The name of robot account
	Name string `json:"name,omitempty"`
	// The description of robot account
	Description string `json:"description,omitempty"`
	// The expiration of robot account (in seconds)
	ExpiresAt int32 `json:"expires_at,omitempty"`
	// The project id of robot account
	ProjectId int32 `json:"project_id,omitempty"`
	// The robot account is disable or enable
	Disabled bool `json:"disabled,omitempty"`
	// The creation time of the robot account
	CreationTime string `json:"creation_time,omitempty"`
	// The update time of the robot account
	UpdateTime string `json:"update_time,omitempty"`
}
