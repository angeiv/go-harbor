# {{classname}}

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChartrepoChartsPost**](ProductsApi.md#ChartrepoChartsPost) | **Post** /chartrepo/charts | Upload a chart file to the defult &#x27;library&#x27; project.
[**ChartrepoHealthGet**](ProductsApi.md#ChartrepoHealthGet) | **Get** /chartrepo/health | Check the health of chart repository service.
[**ChartrepoRepoChartsGet**](ProductsApi.md#ChartrepoRepoChartsGet) | **Get** /chartrepo/{repo}/charts | Get all the charts under the specified project
[**ChartrepoRepoChartsNameDelete**](ProductsApi.md#ChartrepoRepoChartsNameDelete) | **Delete** /chartrepo/{repo}/charts/{name} | Delete all the versions of the specified chart
[**ChartrepoRepoChartsNameGet**](ProductsApi.md#ChartrepoRepoChartsNameGet) | **Get** /chartrepo/{repo}/charts/{name} | Get all the versions of the specified chart
[**ChartrepoRepoChartsNameVersionDelete**](ProductsApi.md#ChartrepoRepoChartsNameVersionDelete) | **Delete** /chartrepo/{repo}/charts/{name}/{version} | Delete the specified chart version
[**ChartrepoRepoChartsNameVersionGet**](ProductsApi.md#ChartrepoRepoChartsNameVersionGet) | **Get** /chartrepo/{repo}/charts/{name}/{version} | Get the specified chart version
[**ChartrepoRepoChartsNameVersionLabelsGet**](ProductsApi.md#ChartrepoRepoChartsNameVersionLabelsGet) | **Get** /chartrepo/{repo}/charts/{name}/{version}/labels | Return the attahced labels of chart.
[**ChartrepoRepoChartsNameVersionLabelsIdDelete**](ProductsApi.md#ChartrepoRepoChartsNameVersionLabelsIdDelete) | **Delete** /chartrepo/{repo}/charts/{name}/{version}/labels/{id} | Remove label from chart.
[**ChartrepoRepoChartsNameVersionLabelsPost**](ProductsApi.md#ChartrepoRepoChartsNameVersionLabelsPost) | **Post** /chartrepo/{repo}/charts/{name}/{version}/labels | Mark label to chart.
[**ChartrepoRepoChartsPost**](ProductsApi.md#ChartrepoRepoChartsPost) | **Post** /chartrepo/{repo}/charts | Upload a chart file to the specified project.
[**ChartrepoRepoProvPost**](ProductsApi.md#ChartrepoRepoProvPost) | **Post** /chartrepo/{repo}/prov | Upload a provance file to the specified project.
[**ConfigurationsGet**](ProductsApi.md#ConfigurationsGet) | **Get** /configurations | Get system configurations.
[**ConfigurationsPut**](ProductsApi.md#ConfigurationsPut) | **Put** /configurations | Modify system configurations.
[**EmailPingPost**](ProductsApi.md#EmailPingPost) | **Post** /email/ping | Test connection and authentication with email server.
[**HealthGet**](ProductsApi.md#HealthGet) | **Get** /health | Health check API
[**InternalSwitchquotaPut**](ProductsApi.md#InternalSwitchquotaPut) | **Put** /internal/switchquota | Enable or disable quota.
[**InternalSyncquotaPost**](ProductsApi.md#InternalSyncquotaPost) | **Post** /internal/syncquota | Sync quota from registry/chart to DB.
[**InternalSyncregistryPost**](ProductsApi.md#InternalSyncregistryPost) | **Post** /internal/syncregistry | Sync repositories from registry to DB.
[**LabelsGet**](ProductsApi.md#LabelsGet) | **Get** /labels | List labels according to the query strings.
[**LabelsIdDelete**](ProductsApi.md#LabelsIdDelete) | **Delete** /labels/{id} | Delete the label specified by ID.
[**LabelsIdGet**](ProductsApi.md#LabelsIdGet) | **Get** /labels/{id} | Get the label specified by ID.
[**LabelsIdPut**](ProductsApi.md#LabelsIdPut) | **Put** /labels/{id} | Update the label properties.
[**LabelsIdResourcesGet**](ProductsApi.md#LabelsIdResourcesGet) | **Get** /labels/{id}/resources | Get the resources that the label is referenced by.
[**LabelsPost**](ProductsApi.md#LabelsPost) | **Post** /labels | Post creates a label
[**LdapGroupsSearchGet**](ProductsApi.md#LdapGroupsSearchGet) | **Get** /ldap/groups/search | Search available ldap groups.
[**LdapPingPost**](ProductsApi.md#LdapPingPost) | **Post** /ldap/ping | Ping available ldap service.
[**LdapUsersImportPost**](ProductsApi.md#LdapUsersImportPost) | **Post** /ldap/users/import | Import selected available ldap users.
[**LdapUsersSearchGet**](ProductsApi.md#LdapUsersSearchGet) | **Get** /ldap/users/search | Search available ldap users.
[**LogsGet**](ProductsApi.md#LogsGet) | **Get** /logs | Get recent logs of the projects which the user is a member of
[**ProjectsGet**](ProductsApi.md#ProjectsGet) | **Get** /projects | List projects
[**ProjectsHead**](ProductsApi.md#ProjectsHead) | **Head** /projects | Check if the project name user provided already exists.
[**ProjectsPost**](ProductsApi.md#ProjectsPost) | **Post** /projects | Create a new project.
[**ProjectsProjectIdDelete**](ProductsApi.md#ProjectsProjectIdDelete) | **Delete** /projects/{project_id} | Delete project by projectID
[**ProjectsProjectIdGet**](ProductsApi.md#ProjectsProjectIdGet) | **Get** /projects/{project_id} | Return specific project detail information
[**ProjectsProjectIdImmutabletagrulesGet**](ProductsApi.md#ProjectsProjectIdImmutabletagrulesGet) | **Get** /projects/{project_id}/immutabletagrules | List all immutable tag rules of current project
[**ProjectsProjectIdImmutabletagrulesIdDelete**](ProductsApi.md#ProjectsProjectIdImmutabletagrulesIdDelete) | **Delete** /projects/{project_id}/immutabletagrules/{id} | Delete the immutable tag rule.
[**ProjectsProjectIdImmutabletagrulesIdPut**](ProductsApi.md#ProjectsProjectIdImmutabletagrulesIdPut) | **Put** /projects/{project_id}/immutabletagrules/{id} | Update the immutable tag rule or enable or disable the rule
[**ProjectsProjectIdImmutabletagrulesPost**](ProductsApi.md#ProjectsProjectIdImmutabletagrulesPost) | **Post** /projects/{project_id}/immutabletagrules | Add an immutable tag rule to current project
[**ProjectsProjectIdLogsGet**](ProductsApi.md#ProjectsProjectIdLogsGet) | **Get** /projects/{project_id}/logs | Get access logs accompany with a relevant project.
[**ProjectsProjectIdMembersGet**](ProductsApi.md#ProjectsProjectIdMembersGet) | **Get** /projects/{project_id}/members | Get all project member information
[**ProjectsProjectIdMembersMidDelete**](ProductsApi.md#ProjectsProjectIdMembersMidDelete) | **Delete** /projects/{project_id}/members/{mid} | Delete project member
[**ProjectsProjectIdMembersMidGet**](ProductsApi.md#ProjectsProjectIdMembersMidGet) | **Get** /projects/{project_id}/members/{mid} | Get the project member information
[**ProjectsProjectIdMembersMidPut**](ProductsApi.md#ProjectsProjectIdMembersMidPut) | **Put** /projects/{project_id}/members/{mid} | Update project member
[**ProjectsProjectIdMembersPost**](ProductsApi.md#ProjectsProjectIdMembersPost) | **Post** /projects/{project_id}/members | Create project member
[**ProjectsProjectIdMetadatasGet**](ProductsApi.md#ProjectsProjectIdMetadatasGet) | **Get** /projects/{project_id}/metadatas | Get project metadata.
[**ProjectsProjectIdMetadatasMetaNameDelete**](ProductsApi.md#ProjectsProjectIdMetadatasMetaNameDelete) | **Delete** /projects/{project_id}/metadatas/{meta_name} | Delete metadata of a project
[**ProjectsProjectIdMetadatasMetaNameGet**](ProductsApi.md#ProjectsProjectIdMetadatasMetaNameGet) | **Get** /projects/{project_id}/metadatas/{meta_name} | Get project metadata
[**ProjectsProjectIdMetadatasMetaNamePut**](ProductsApi.md#ProjectsProjectIdMetadatasMetaNamePut) | **Put** /projects/{project_id}/metadatas/{meta_name} | Update metadata of a project.
[**ProjectsProjectIdMetadatasPost**](ProductsApi.md#ProjectsProjectIdMetadatasPost) | **Post** /projects/{project_id}/metadatas | Add metadata for the project.
[**ProjectsProjectIdPut**](ProductsApi.md#ProjectsProjectIdPut) | **Put** /projects/{project_id} | Update properties for a selected project.
[**ProjectsProjectIdRobotsGet**](ProductsApi.md#ProjectsProjectIdRobotsGet) | **Get** /projects/{project_id}/robots | Get all robot accounts of specified project
[**ProjectsProjectIdRobotsPost**](ProductsApi.md#ProjectsProjectIdRobotsPost) | **Post** /projects/{project_id}/robots | Create a robot account for project
[**ProjectsProjectIdRobotsRobotIdDelete**](ProductsApi.md#ProjectsProjectIdRobotsRobotIdDelete) | **Delete** /projects/{project_id}/robots/{robot_id} | Delete the specified robot account
[**ProjectsProjectIdRobotsRobotIdGet**](ProductsApi.md#ProjectsProjectIdRobotsRobotIdGet) | **Get** /projects/{project_id}/robots/{robot_id} | Return the infor of the specified robot account.
[**ProjectsProjectIdRobotsRobotIdPut**](ProductsApi.md#ProjectsProjectIdRobotsRobotIdPut) | **Put** /projects/{project_id}/robots/{robot_id} | Update status of robot account.
[**ProjectsProjectIdScannerCandidatesGet**](ProductsApi.md#ProjectsProjectIdScannerCandidatesGet) | **Get** /projects/{project_id}/scanner/candidates | Get scanner registration candidates for configurating project level scanner
[**ProjectsProjectIdScannerGet**](ProductsApi.md#ProjectsProjectIdScannerGet) | **Get** /projects/{project_id}/scanner | Get project level scanner
[**ProjectsProjectIdSummaryGet**](ProductsApi.md#ProjectsProjectIdSummaryGet) | **Get** /projects/{project_id}/summary | Get summary of the project.
[**ProjectsProjectIdWebhookJobsGet**](ProductsApi.md#ProjectsProjectIdWebhookJobsGet) | **Get** /projects/{project_id}/webhook/jobs | List project webhook jobs
[**ProjectsProjectIdWebhookLasttriggerGet**](ProductsApi.md#ProjectsProjectIdWebhookLasttriggerGet) | **Get** /projects/{project_id}/webhook/lasttrigger | Get project webhook policy last trigger info
[**ProjectsProjectIdWebhookPoliciesGet**](ProductsApi.md#ProjectsProjectIdWebhookPoliciesGet) | **Get** /projects/{project_id}/webhook/policies | List project webhook policies.
[**ProjectsProjectIdWebhookPoliciesPolicyIdDelete**](ProductsApi.md#ProjectsProjectIdWebhookPoliciesPolicyIdDelete) | **Delete** /projects/{project_id}/webhook/policies/{policy_id} | Delete webhook policy of a project
[**ProjectsProjectIdWebhookPoliciesPolicyIdGet**](ProductsApi.md#ProjectsProjectIdWebhookPoliciesPolicyIdGet) | **Get** /projects/{project_id}/webhook/policies/{policy_id} | Get project webhook policy
[**ProjectsProjectIdWebhookPoliciesPolicyIdPut**](ProductsApi.md#ProjectsProjectIdWebhookPoliciesPolicyIdPut) | **Put** /projects/{project_id}/webhook/policies/{policy_id} | Update webhook policy of a project.
[**ProjectsProjectIdWebhookPoliciesPost**](ProductsApi.md#ProjectsProjectIdWebhookPoliciesPost) | **Post** /projects/{project_id}/webhook/policies | Create project webhook policy.
[**ProjectsProjectIdWebhookPoliciesTestPost**](ProductsApi.md#ProjectsProjectIdWebhookPoliciesTestPost) | **Post** /projects/{project_id}/webhook/policies/test | Test project webhook connection
[**QuotasGet**](ProductsApi.md#QuotasGet) | **Get** /quotas | List quotas
[**QuotasIdGet**](ProductsApi.md#QuotasIdGet) | **Get** /quotas/{id} | Get the specified quota
[**QuotasIdPut**](ProductsApi.md#QuotasIdPut) | **Put** /quotas/{id} | Update the specified quota
[**RegistriesGet**](ProductsApi.md#RegistriesGet) | **Get** /registries | List registries.
[**RegistriesIdDelete**](ProductsApi.md#RegistriesIdDelete) | **Delete** /registries/{id} | Delete specific registry.
[**RegistriesIdGet**](ProductsApi.md#RegistriesIdGet) | **Get** /registries/{id} | Get registry.
[**RegistriesIdInfoGet**](ProductsApi.md#RegistriesIdInfoGet) | **Get** /registries/{id}/info | Get registry info.
[**RegistriesIdNamespaceGet**](ProductsApi.md#RegistriesIdNamespaceGet) | **Get** /registries/{id}/namespace | List namespaces of registry
[**RegistriesIdPut**](ProductsApi.md#RegistriesIdPut) | **Put** /registries/{id} | Update a given registry.
[**RegistriesPingPost**](ProductsApi.md#RegistriesPingPost) | **Post** /registries/ping | Ping status of a registry.
[**RegistriesPost**](ProductsApi.md#RegistriesPost) | **Post** /registries | Create a new registry.
[**ReplicationAdaptersGet**](ProductsApi.md#ReplicationAdaptersGet) | **Get** /replication/adapters | List supported adapters.
[**ReplicationExecutionsGet**](ProductsApi.md#ReplicationExecutionsGet) | **Get** /replication/executions | List replication executions.
[**ReplicationExecutionsIdGet**](ProductsApi.md#ReplicationExecutionsIdGet) | **Get** /replication/executions/{id} | Get the execution of the replication.
[**ReplicationExecutionsIdPut**](ProductsApi.md#ReplicationExecutionsIdPut) | **Put** /replication/executions/{id} | Stop the execution of the replication.
[**ReplicationExecutionsIdTasksGet**](ProductsApi.md#ReplicationExecutionsIdTasksGet) | **Get** /replication/executions/{id}/tasks | Get the task list of one execution.
[**ReplicationExecutionsIdTasksTaskIdLogGet**](ProductsApi.md#ReplicationExecutionsIdTasksTaskIdLogGet) | **Get** /replication/executions/{id}/tasks/{task_id}/log | Get the log of one task.
[**ReplicationExecutionsPost**](ProductsApi.md#ReplicationExecutionsPost) | **Post** /replication/executions | Start one execution of the replication.
[**ReplicationPoliciesGet**](ProductsApi.md#ReplicationPoliciesGet) | **Get** /replication/policies | List replication policies
[**ReplicationPoliciesIdDelete**](ProductsApi.md#ReplicationPoliciesIdDelete) | **Delete** /replication/policies/{id} | Delete the replication policy specified by ID.
[**ReplicationPoliciesIdGet**](ProductsApi.md#ReplicationPoliciesIdGet) | **Get** /replication/policies/{id} | Get replication policy.
[**ReplicationPoliciesIdPut**](ProductsApi.md#ReplicationPoliciesIdPut) | **Put** /replication/policies/{id} | Update the replication policy
[**ReplicationPoliciesPost**](ProductsApi.md#ReplicationPoliciesPost) | **Post** /replication/policies | Create a replication policy
[**RepositoriesGet**](ProductsApi.md#RepositoriesGet) | **Get** /repositories | Get repositories accompany with relevant project and repo name.
[**RepositoriesRepoNameDelete**](ProductsApi.md#RepositoriesRepoNameDelete) | **Delete** /repositories/{repo_name} | Delete a repository.
[**RepositoriesRepoNameLabelsGet**](ProductsApi.md#RepositoriesRepoNameLabelsGet) | **Get** /repositories/{repo_name}/labels | Get labels of a repository.
[**RepositoriesRepoNameLabelsLabelIdDelete**](ProductsApi.md#RepositoriesRepoNameLabelsLabelIdDelete) | **Delete** /repositories/{repo_name}/labels/{label_id} | Delete label from the repository.
[**RepositoriesRepoNameLabelsPost**](ProductsApi.md#RepositoriesRepoNameLabelsPost) | **Post** /repositories/{repo_name}/labels | Add a label to the repository.
[**RepositoriesRepoNamePut**](ProductsApi.md#RepositoriesRepoNamePut) | **Put** /repositories/{repo_name} | Update description of the repository.
[**RepositoriesRepoNameSignaturesGet**](ProductsApi.md#RepositoriesRepoNameSignaturesGet) | **Get** /repositories/{repo_name}/signatures | Get signature information of a repository
[**RepositoriesRepoNameTagsGet**](ProductsApi.md#RepositoriesRepoNameTagsGet) | **Get** /repositories/{repo_name}/tags | Get tags of a relevant repository.
[**RepositoriesRepoNameTagsPost**](ProductsApi.md#RepositoriesRepoNameTagsPost) | **Post** /repositories/{repo_name}/tags | Retag an image
[**RepositoriesRepoNameTagsTagDelete**](ProductsApi.md#RepositoriesRepoNameTagsTagDelete) | **Delete** /repositories/{repo_name}/tags/{tag} | Delete a tag in a repository.
[**RepositoriesRepoNameTagsTagGet**](ProductsApi.md#RepositoriesRepoNameTagsTagGet) | **Get** /repositories/{repo_name}/tags/{tag} | Get the tag of the repository.
[**RepositoriesRepoNameTagsTagLabelsGet**](ProductsApi.md#RepositoriesRepoNameTagsTagLabelsGet) | **Get** /repositories/{repo_name}/tags/{tag}/labels | Get labels of an image.
[**RepositoriesRepoNameTagsTagLabelsLabelIdDelete**](ProductsApi.md#RepositoriesRepoNameTagsTagLabelsLabelIdDelete) | **Delete** /repositories/{repo_name}/tags/{tag}/labels/{label_id} | Delete label from the image.
[**RepositoriesRepoNameTagsTagLabelsPost**](ProductsApi.md#RepositoriesRepoNameTagsTagLabelsPost) | **Post** /repositories/{repo_name}/tags/{tag}/labels | Add a label to image.
[**RepositoriesRepoNameTagsTagManifestGet**](ProductsApi.md#RepositoriesRepoNameTagsTagManifestGet) | **Get** /repositories/{repo_name}/tags/{tag}/manifest | Get manifests of a relevant repository.
[**RepositoriesRepoNameTagsTagScanPost**](ProductsApi.md#RepositoriesRepoNameTagsTagScanPost) | **Post** /repositories/{repo_name}/tags/{tag}/scan | Scan the image.
[**RepositoriesRepoNameTagsTagScanUuidLogGet**](ProductsApi.md#RepositoriesRepoNameTagsTagScanUuidLogGet) | **Get** /repositories/{repo_name}/tags/{tag}/scan/{uuid}/log | Get scan log
[**RepositoriesTopGet**](ProductsApi.md#RepositoriesTopGet) | **Get** /repositories/top | Get public repositories which are accessed most.
[**RetentionsIdExecutionsEidPatch**](ProductsApi.md#RetentionsIdExecutionsEidPatch) | **Patch** /retentions/{id}/executions/{eid} | Stop a Retention job
[**RetentionsIdExecutionsEidTasksGet**](ProductsApi.md#RetentionsIdExecutionsEidTasksGet) | **Get** /retentions/{id}/executions/{eid}/tasks | Get Retention job tasks
[**RetentionsIdExecutionsEidTasksTidGet**](ProductsApi.md#RetentionsIdExecutionsEidTasksTidGet) | **Get** /retentions/{id}/executions/{eid}/tasks/{tid} | Get Retention job task log
[**RetentionsIdExecutionsGet**](ProductsApi.md#RetentionsIdExecutionsGet) | **Get** /retentions/{id}/executions | Get a Retention job
[**RetentionsIdExecutionsPost**](ProductsApi.md#RetentionsIdExecutionsPost) | **Post** /retentions/{id}/executions | Trigger a Retention job
[**RetentionsIdGet**](ProductsApi.md#RetentionsIdGet) | **Get** /retentions/{id} | Get Retention Policy
[**RetentionsIdPut**](ProductsApi.md#RetentionsIdPut) | **Put** /retentions/{id} | Update Retention Policy
[**RetentionsMetadatasGet**](ProductsApi.md#RetentionsMetadatasGet) | **Get** /retentions/metadatas | Get Retention Metadatas
[**RetentionsPost**](ProductsApi.md#RetentionsPost) | **Post** /retentions | Create Retention Policy
[**ScannersGet**](ProductsApi.md#ScannersGet) | **Get** /scanners | List scanner registrations
[**ScannersPingPost**](ProductsApi.md#ScannersPingPost) | **Post** /scanners/ping | Tests scanner registration settings
[**ScannersRegistrationIdGet**](ProductsApi.md#ScannersRegistrationIdGet) | **Get** /scanners/{registration_id} | Get a scanner registration details
[**ScannersRegistrationIdMetadataGet**](ProductsApi.md#ScannersRegistrationIdMetadataGet) | **Get** /scanners/{registration_id}/metadata | Get the metadata of the specified scanner registration
[**ScansAllMetricsGet**](ProductsApi.md#ScansAllMetricsGet) | **Get** /scans/all/metrics | Get the metrics of the latest scan all process
[**ScansScheduleMetricsGet**](ProductsApi.md#ScansScheduleMetricsGet) | **Get** /scans/schedule/metrics | Get the metrics of the latest scheduled scan all process
[**SearchGet**](ProductsApi.md#SearchGet) | **Get** /search | Search for projects, repositories and helm charts
[**StatisticsGet**](ProductsApi.md#StatisticsGet) | **Get** /statistics | Get projects number and repositories number relevant to the user
[**SystemCVEWhitelistGet**](ProductsApi.md#SystemCVEWhitelistGet) | **Get** /system/CVEWhitelist | Get the system level whitelist of CVE.
[**SystemCVEWhitelistPut**](ProductsApi.md#SystemCVEWhitelistPut) | **Put** /system/CVEWhitelist | Update the system level whitelist of CVE.
[**SystemGcGet**](ProductsApi.md#SystemGcGet) | **Get** /system/gc | Get gc results.
[**SystemGcIdGet**](ProductsApi.md#SystemGcIdGet) | **Get** /system/gc/{id} | Get gc status.
[**SystemGcIdLogGet**](ProductsApi.md#SystemGcIdLogGet) | **Get** /system/gc/{id}/log | Get gc job log.
[**SystemGcScheduleGet**](ProductsApi.md#SystemGcScheduleGet) | **Get** /system/gc/schedule | Get gc&#x27;s schedule.
[**SystemGcSchedulePost**](ProductsApi.md#SystemGcSchedulePost) | **Post** /system/gc/schedule | Create a gc schedule.
[**SystemGcSchedulePut**](ProductsApi.md#SystemGcSchedulePut) | **Put** /system/gc/schedule | Update gc&#x27;s schedule.
[**SystemOidcPingPost**](ProductsApi.md#SystemOidcPingPost) | **Post** /system/oidc/ping | Test the OIDC endpoint.
[**SystemScanAllScheduleGet**](ProductsApi.md#SystemScanAllScheduleGet) | **Get** /system/scanAll/schedule | Get scan_all&#x27;s schedule.
[**SystemScanAllSchedulePost**](ProductsApi.md#SystemScanAllSchedulePost) | **Post** /system/scanAll/schedule | Create a schedule or a manual trigger for the scan all job.
[**SystemScanAllSchedulePut**](ProductsApi.md#SystemScanAllSchedulePut) | **Put** /system/scanAll/schedule | Update scan all&#x27;s schedule.
[**SysteminfoGet**](ProductsApi.md#SysteminfoGet) | **Get** /systeminfo | Get general system info
[**SysteminfoGetcertGet**](ProductsApi.md#SysteminfoGetcertGet) | **Get** /systeminfo/getcert | Get default root certificate.
[**SysteminfoVolumesGet**](ProductsApi.md#SysteminfoVolumesGet) | **Get** /systeminfo/volumes | Get system volume info (total/free size).
[**UsergroupsGet**](ProductsApi.md#UsergroupsGet) | **Get** /usergroups | Get all user groups information
[**UsergroupsGroupIdDelete**](ProductsApi.md#UsergroupsGroupIdDelete) | **Delete** /usergroups/{group_id} | Delete user group
[**UsergroupsGroupIdGet**](ProductsApi.md#UsergroupsGroupIdGet) | **Get** /usergroups/{group_id} | Get user group information
[**UsergroupsGroupIdPut**](ProductsApi.md#UsergroupsGroupIdPut) | **Put** /usergroups/{group_id} | Update group information
[**UsergroupsPost**](ProductsApi.md#UsergroupsPost) | **Post** /usergroups | Create user group
[**UsersCurrentGet**](ProductsApi.md#UsersCurrentGet) | **Get** /users/current | Get current user info.
[**UsersCurrentPermissionsGet**](ProductsApi.md#UsersCurrentPermissionsGet) | **Get** /users/current/permissions | Get current user permissions.
[**UsersGet**](ProductsApi.md#UsersGet) | **Get** /users | Get registered users of Harbor.
[**UsersPost**](ProductsApi.md#UsersPost) | **Post** /users | Creates a new user account.
[**UsersSearchGet**](ProductsApi.md#UsersSearchGet) | **Get** /users/search | Search users by username
[**UsersUserIdCliSecretPut**](ProductsApi.md#UsersUserIdCliSecretPut) | **Put** /users/{user_id}/cli_secret | Set CLI secret for a user.
[**UsersUserIdDelete**](ProductsApi.md#UsersUserIdDelete) | **Delete** /users/{user_id} | Mark a registered user as be removed.
[**UsersUserIdGet**](ProductsApi.md#UsersUserIdGet) | **Get** /users/{user_id} | Get a user&#x27;s profile.
[**UsersUserIdPasswordPut**](ProductsApi.md#UsersUserIdPasswordPut) | **Put** /users/{user_id}/password | Change the password on a user that already exists.
[**UsersUserIdPut**](ProductsApi.md#UsersUserIdPut) | **Put** /users/{user_id} | Update a registered user to change his profile.
[**UsersUserIdSysadminPut**](ProductsApi.md#UsersUserIdSysadminPut) | **Put** /users/{user_id}/sysadmin | Update a registered user to change to be an administrator of Harbor.

# **ChartrepoChartsPost**
> ChartrepoChartsPost(ctx, chart, prov)
Upload a chart file to the defult 'library' project.

Upload a chart file to the default 'library' project. Uploading together with the prov file at the same time is also supported.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **chart** | **string****string**|  | 
  **prov** | **string****string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChartrepoHealthGet**
> InlineResponse200 ChartrepoHealthGet(ctx, )
Check the health of chart repository service.

Check the health of chart repository service.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChartrepoRepoChartsGet**
> []ChartInfoEntry ChartrepoRepoChartsGet(ctx, repo)
Get all the charts under the specified project

Get all the charts under the specified project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repo** | **string**| The project name | 

### Return type

[**[]ChartInfoEntry**](ChartInfoEntry.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChartrepoRepoChartsNameDelete**
> ChartrepoRepoChartsNameDelete(ctx, repo, name)
Delete all the versions of the specified chart

Delete all the versions of the specified chart

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repo** | **string**| The project name | 
  **name** | **string**| The chart name | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChartrepoRepoChartsNameGet**
> Array ChartrepoRepoChartsNameGet(ctx, repo, name)
Get all the versions of the specified chart

Get all the versions of the specified chart

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repo** | **string**| The project name | 
  **name** | **string**| The chart name | 

### Return type

[**Array**](array.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChartrepoRepoChartsNameVersionDelete**
> ChartrepoRepoChartsNameVersionDelete(ctx, repo, name, version)
Delete the specified chart version

Delete the specified chart version

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repo** | **string**| The project name | 
  **name** | **string**| The chart name | 
  **version** | **string**| The chart version | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChartrepoRepoChartsNameVersionGet**
> ChartVersionDetails ChartrepoRepoChartsNameVersionGet(ctx, repo, name, version)
Get the specified chart version

Get the specified chart version

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repo** | **string**| The project name | 
  **name** | **string**| The chart name | 
  **version** | **string**| The chart version | 

### Return type

[**ChartVersionDetails**](ChartVersionDetails.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChartrepoRepoChartsNameVersionLabelsGet**
> ChartrepoRepoChartsNameVersionLabelsGet(ctx, repo, name, version)
Return the attahced labels of chart.

Return the attahced labels of the specified chart version.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repo** | **string**| The project name | 
  **name** | **string**| The chart name | 
  **version** | **string**| The chart version | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChartrepoRepoChartsNameVersionLabelsIdDelete**
> ChartrepoRepoChartsNameVersionLabelsIdDelete(ctx, repo, name, version, id)
Remove label from chart.

Remove label from the specified chart version.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repo** | **string**| The project name | 
  **name** | **string**| The chart name | 
  **version** | **string**| The chart version | 
  **id** | **int32**| The label ID | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChartrepoRepoChartsNameVersionLabelsPost**
> ChartrepoRepoChartsNameVersionLabelsPost(ctx, body, repo, name, version)
Mark label to chart.

Mark label to the specified chart version.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Label**](Label.md)| The label being marked to the chart version | 
  **repo** | **string**| The project name | 
  **name** | **string**| The chart name | 
  **version** | **string**| The chart version | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChartrepoRepoChartsPost**
> ChartrepoRepoChartsPost(ctx, chart, prov, repo)
Upload a chart file to the specified project.

Upload a chart file to the specified project. With this API, the corresponding provance file can be uploaded together with chart file at once.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **chart** | **string****string**|  | 
  **prov** | **string****string**|  | 
  **repo** | **string**| The project name | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChartrepoRepoProvPost**
> ChartrepoRepoProvPost(ctx, prov, repo)
Upload a provance file to the specified project.

Upload a provance file to the specified project. The provance file should be targeted for an existing chart file.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **prov** | **string****string**|  | 
  **repo** | **string**| The project name | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConfigurationsGet**
> ConfigurationsResponse ConfigurationsGet(ctx, )
Get system configurations.

This endpoint is for retrieving system configurations that only provides for admin user. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ConfigurationsResponse**](ConfigurationsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConfigurationsPut**
> ConfigurationsPut(ctx, body)
Modify system configurations.

This endpoint is for modifying system configurations that only provides for admin user. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Configurations**](Configurations.md)| The configuration map can contain a subset of the attributes of the schema, which are to be updated. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EmailPingPost**
> EmailPingPost(ctx, optional)
Test connection and authentication with email server.

Test connection and authentication with email server. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProductsApiEmailPingPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductsApiEmailPingPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of EmailServerSetting**](EmailServerSetting.md)| Email server settings, if some of the settings are not assigned, they will be read from system configuration. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HealthGet**
> OverallHealthStatus HealthGet(ctx, )
Health check API

The endpoint returns the health stauts of the system. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**OverallHealthStatus**](OverallHealthStatus.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InternalSwitchquotaPut**
> InternalSwitchquotaPut(ctx, body)
Enable or disable quota.

This endpoint is for enable/disable quota. When quota is disabled, no resource require/release in image/chart push and delete. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**QuotaSwitcher**](QuotaSwitcher.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InternalSyncquotaPost**
> InternalSyncquotaPost(ctx, )
Sync quota from registry/chart to DB.

This endpoint is for syncing quota usage of registry/chart with database. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InternalSyncregistryPost**
> InternalSyncregistryPost(ctx, )
Sync repositories from registry to DB.

This endpoint is for syncing all repositories of registry with database. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LabelsGet**
> []Label LabelsGet(ctx, scope, optional)
List labels according to the query strings.

This endpoint let user list labels by name, scope and project_id 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **scope** | **string**| The label scope. Valid values are g and p. g for global labels and p for project labels. | 
 **optional** | ***ProductsApiLabelsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductsApiLabelsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| The label name. | 
 **projectId** | **optional.Int64**| Relevant project ID, required when scope is p. | 
 **page** | **optional.Int32**| The page nubmer. | 
 **pageSize** | **optional.Int32**| The size of per page. | 

### Return type

[**[]Label**](Label.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LabelsIdDelete**
> LabelsIdDelete(ctx, id)
Delete the label specified by ID.

Delete the label specified by ID. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| Label ID | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LabelsIdGet**
> Label LabelsIdGet(ctx, id)
Get the label specified by ID.

This endpoint let user get the label by specific ID. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| Label ID | 

### Return type

[**Label**](Label.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LabelsIdPut**
> LabelsIdPut(ctx, body, id)
Update the label properties.

This endpoint let user update label properties. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Label**](Label.md)| The updated label json object. | 
  **id** | **int64**| Label ID | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LabelsIdResourcesGet**
> Resource LabelsIdResourcesGet(ctx, id)
Get the resources that the label is referenced by.

This endpoint let user get the resources that the label is referenced by. Only the replication policies are returned for now. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| Label ID | 

### Return type

[**Resource**](Resource.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LabelsPost**
> LabelsPost(ctx, body)
Post creates a label

This endpoint let user creates a label. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Label**](Label.md)| The json object of label. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LdapGroupsSearchGet**
> []UserGroup LdapGroupsSearchGet(ctx, optional)
Search available ldap groups.

This endpoint searches the available ldap groups based on related configuration parameters. support to search by groupname or groupdn. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProductsApiLdapGroupsSearchGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductsApiLdapGroupsSearchGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupname** | **optional.String**| Ldap group name | 
 **groupdn** | **optional.String**| The LDAP group DN | 

### Return type

[**[]UserGroup**](UserGroup.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LdapPingPost**
> LdapPingPost(ctx, optional)
Ping available ldap service.

This endpoint ping the available ldap service for test related configuration parameters. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProductsApiLdapPingPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductsApiLdapPingPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of LdapConf**](LdapConf.md)| ldap configuration. support input ldap service configuration. If it&#x27;s a empty request, will load current configuration from the system. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LdapUsersImportPost**
> LdapUsersImportPost(ctx, body)
Import selected available ldap users.

This endpoint adds the selected available ldap users to harbor based on related configuration parameters from the system. System will try to guess the user email address and realname, add to harbor user information. If have errors when import user, will return the list of importing failed uid and the failed reason. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LdapImportUsers**](LdapImportUsers.md)| The uid listed for importing. This list will check users validity of ldap service based on configuration from the system. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LdapUsersSearchGet**
> []LdapUsers LdapUsersSearchGet(ctx, optional)
Search available ldap users.

This endpoint searches the available ldap users based on related configuration parameters. Support searched by input ladp configuration, load configuration from the system and specific filter. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProductsApiLdapUsersSearchGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductsApiLdapUsersSearchGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **optional.String**| Registered user ID | 

### Return type

[**[]LdapUsers**](LdapUsers.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LogsGet**
> []AccessLog LogsGet(ctx, optional)
Get recent logs of the projects which the user is a member of

This endpoint let user see the recent operation logs of the projects which he is member of 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProductsApiLogsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductsApiLogsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **optional.String**| Username of the operator. | 
 **repository** | **optional.String**| The name of repository | 
 **tag** | **optional.String**| The name of tag | 
 **operation** | **optional.String**| The operation | 
 **beginTimestamp** | **optional.String**| The begin timestamp | 
 **endTimestamp** | **optional.String**| The end timestamp | 
 **page** | **optional.Int32**| The page number, default is 1. | 
 **pageSize** | **optional.Int32**| The size of per page, default is 10, maximum is 100. | 

### Return type

[**[]AccessLog**](AccessLog.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsGet**
> []Project ProjectsGet(ctx, optional)
List projects

This endpoint returns all projects created by Harbor, and can be filtered by project name. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProductsApiProjectsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductsApiProjectsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| The name of project. | 
 **public** | **optional.Bool**| The project is public or private. | 
 **owner** | **optional.String**| The name of project owner. | 
 **page** | **optional.Int32**| The page number, default is 1. | 
 **pageSize** | **optional.Int32**| The size of per page, default is 10, maximum is 100. | 

### Return type

[**[]Project**](Project.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsHead**
> ProjectsHead(ctx, projectName)
Check if the project name user provided already exists.

This endpoint is used to check if the project name user provided already exist. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectName** | **string**| Project name for checking exists. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsPost**
> ProjectsPost(ctx, body)
Create a new project.

This endpoint is for user to create a new project. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ProjectReq**](ProjectReq.md)| New created project. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectIdDelete**
> ProjectsProjectIdDelete(ctx, projectId)
Delete project by projectID

This endpoint is aimed to delete project by project ID. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| Project ID of project which will be deleted. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectIdGet**
> Project ProjectsProjectIdGet(ctx, projectId)
Return specific project detail information

This endpoint returns specific project information by project ID. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| Project ID for filtering results. | 

### Return type

[**Project**](Project.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectIdImmutabletagrulesGet**
> []ImmutableTagRule ProjectsProjectIdImmutabletagrulesGet(ctx, projectId)
List all immutable tag rules of current project

This endpoint returns the immutable tag rules of a project 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| Relevant project ID. | 

### Return type

[**[]ImmutableTagRule**](ImmutableTagRule.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectIdImmutabletagrulesIdDelete**
> ProjectsProjectIdImmutabletagrulesIdDelete(ctx, projectId, id)
Delete the immutable tag rule.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| Relevant project ID. | 
  **id** | **int64**| Immutable tag rule ID. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectIdImmutabletagrulesIdPut**
> ProjectsProjectIdImmutabletagrulesIdPut(ctx, projectId, id, optional)
Update the immutable tag rule or enable or disable the rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| Relevant project ID. | 
  **id** | **int64**| Immutable tag rule ID. | 
 **optional** | ***ProductsApiProjectsProjectIdImmutabletagrulesIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductsApiProjectsProjectIdImmutabletagrulesIdPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ImmutableTagRule**](ImmutableTagRule.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectIdImmutabletagrulesPost**
> ProjectsProjectIdImmutabletagrulesPost(ctx, projectId, optional)
Add an immutable tag rule to current project

This endpoint add an immutable tag rule to the project 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| Relevant project ID. | 
 **optional** | ***ProductsApiProjectsProjectIdImmutabletagrulesPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductsApiProjectsProjectIdImmutabletagrulesPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ImmutableTagRule**](ImmutableTagRule.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectIdLogsGet**
> []AccessLog ProjectsProjectIdLogsGet(ctx, projectId, optional)
Get access logs accompany with a relevant project.

This endpoint let user search access logs filtered by operations and date time ranges. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| Relevant project ID | 
 **optional** | ***ProductsApiProjectsProjectIdLogsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductsApiProjectsProjectIdLogsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **username** | **optional.String**| Username of the operator. | 
 **repository** | **optional.String**| The name of repository | 
 **tag** | **optional.String**| The name of tag | 
 **operation** | **optional.String**| The operation | 
 **beginTimestamp** | **optional.String**| The begin timestamp | 
 **endTimestamp** | **optional.String**| The end timestamp | 
 **page** | **optional.Int32**| The page number, default is 1. | 
 **pageSize** | **optional.Int32**| The size of per page, default is 10, maximum is 100. | 

### Return type

[**[]AccessLog**](AccessLog.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectIdMembersGet**
> []ProjectMemberEntity ProjectsProjectIdMembersGet(ctx, projectId, optional)
Get all project member information

Get all project member information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| Relevant project ID. | 
 **optional** | ***ProductsApiProjectsProjectIdMembersGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductsApiProjectsProjectIdMembersGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **entityname** | **optional.String**| The entity name to search. | 

### Return type

[**[]ProjectMemberEntity**](ProjectMemberEntity.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectIdMembersMidDelete**
> ProjectsProjectIdMembersMidDelete(ctx, projectId, mid)
Delete project member

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| Relevant project ID. | 
  **mid** | **int64**| Member ID. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectIdMembersMidGet**
> ProjectMemberEntity ProjectsProjectIdMembersMidGet(ctx, projectId, mid)
Get the project member information

Get the project member information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| Relevant project ID. | 
  **mid** | **int64**| The member ID | 

### Return type

[**ProjectMemberEntity**](ProjectMemberEntity.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectIdMembersMidPut**
> ProjectsProjectIdMembersMidPut(ctx, projectId, mid, optional)
Update project member

Update project member relationship

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| Relevant project ID. | 
  **mid** | **int64**| Member ID. | 
 **optional** | ***ProductsApiProjectsProjectIdMembersMidPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductsApiProjectsProjectIdMembersMidPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of RoleRequest**](RoleRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectIdMembersPost**
> ProjectsProjectIdMembersPost(ctx, projectId, optional)
Create project member

Create project member relationship, the member can be one of the user_member and group_member,  The user_member need to specify user_id or username. If the user already exist in harbor DB, specify the user_id,  If does not exist in harbor DB, it will SearchAndOnBoard the user. The group_member need to specify id or ldap_group_dn. If the group already exist in harbor DB. specify the user group's id,  If does not exist, it will SearchAndOnBoard the group. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| Relevant project ID. | 
 **optional** | ***ProductsApiProjectsProjectIdMembersPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductsApiProjectsProjectIdMembersPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ProjectMember**](ProjectMember.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectIdMetadatasGet**
> ProjectMetadata ProjectsProjectIdMetadatasGet(ctx, projectId)
Get project metadata.

This endpoint returns metadata of the project specified by project ID. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| The ID of project. | 

### Return type

[**ProjectMetadata**](ProjectMetadata.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectIdMetadatasMetaNameDelete**
> ProjectsProjectIdMetadatasMetaNameDelete(ctx, projectId, metaName)
Delete metadata of a project

This endpoint is aimed to delete metadata of a project. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| The ID of project. | 
  **metaName** | **string**| The name of metadat. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectIdMetadatasMetaNameGet**
> ProjectMetadata ProjectsProjectIdMetadatasMetaNameGet(ctx, projectId, metaName)
Get project metadata

This endpoint returns specified metadata of a project. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| Project ID for filtering results. | 
  **metaName** | **string**| The name of metadat. | 

### Return type

[**ProjectMetadata**](ProjectMetadata.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectIdMetadatasMetaNamePut**
> ProjectsProjectIdMetadatasMetaNamePut(ctx, projectId, metaName)
Update metadata of a project.

This endpoint is aimed to update the metadata of a project. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| The ID of project. | 
  **metaName** | **string**| The name of metadat. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectIdMetadatasPost**
> ProjectsProjectIdMetadatasPost(ctx, body, projectId)
Add metadata for the project.

This endpoint is aimed to add metadata of a project. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ProjectMetadata**](ProjectMetadata.md)| The metadata of project. | 
  **projectId** | **int64**| Selected project ID. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectIdPut**
> ProjectsProjectIdPut(ctx, body, projectId)
Update properties for a selected project.

This endpoint is aimed to update the properties of a project. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ProjectReq**](ProjectReq.md)| Updates of project. | 
  **projectId** | **int64**| Selected project ID. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectIdRobotsGet**
> []RobotAccount ProjectsProjectIdRobotsGet(ctx, projectId)
Get all robot accounts of specified project

Get all robot accounts of specified project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| Relevant project ID. | 

### Return type

[**[]RobotAccount**](RobotAccount.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectIdRobotsPost**
> RobotAccountPostRep ProjectsProjectIdRobotsPost(ctx, body, projectId)
Create a robot account for project

Create a robot account for project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RobotAccountCreate**](RobotAccountCreate.md)| Request body of creating a robot account. | 
  **projectId** | **int64**| Relevant project ID. | 

### Return type

[**RobotAccountPostRep**](RobotAccountPostRep.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectIdRobotsRobotIdDelete**
> ProjectsProjectIdRobotsRobotIdDelete(ctx, projectId, robotId)
Delete the specified robot account

Delete the specified robot account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| Relevant project ID. | 
  **robotId** | **int64**| The ID of robot account. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectIdRobotsRobotIdGet**
> RobotAccount ProjectsProjectIdRobotsRobotIdGet(ctx, projectId, robotId)
Return the infor of the specified robot account.

Return the infor of the specified robot account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| Relevant project ID. | 
  **robotId** | **int64**| The ID of robot account. | 

### Return type

[**RobotAccount**](RobotAccount.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectIdRobotsRobotIdPut**
> ProjectsProjectIdRobotsRobotIdPut(ctx, body, projectId, robotId)
Update status of robot account.

Used to disable/enable a specified robot account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RobotAccountUpdate**](RobotAccountUpdate.md)| Request body of enable/disable a robot account. | 
  **projectId** | **int64**| Relevant project ID. | 
  **robotId** | **int64**| The ID of robot account. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectIdScannerCandidatesGet**
> []ScannerRegistration ProjectsProjectIdScannerCandidatesGet(ctx, projectId)
Get scanner registration candidates for configurating project level scanner

Retrieve the system configured scanner registrations as candidates of setting project level scanner. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| The project identifier. | 

### Return type

[**[]ScannerRegistration**](ScannerRegistration.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectIdScannerGet**
> ScannerRegistration ProjectsProjectIdScannerGet(ctx, projectId)
Get project level scanner

Get the scanner registration of the specified project. If no scanner registration is configured for the specified project, the system default scanner registration will be returned.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| The project identifier. | 

### Return type

[**ScannerRegistration**](ScannerRegistration.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectIdSummaryGet**
> ProjectSummary ProjectsProjectIdSummaryGet(ctx, projectId)
Get summary of the project.

Get summary of the project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| Relevant project ID | 

### Return type

[**ProjectSummary**](ProjectSummary.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectIdWebhookJobsGet**
> []WebhookJob ProjectsProjectIdWebhookJobsGet(ctx, projectId, policyId)
List project webhook jobs

This endpoint returns webhook jobs of a project. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| Relevant project ID. | 
  **policyId** | **int64**| The policy ID. | 

### Return type

[**[]WebhookJob**](WebhookJob.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectIdWebhookLasttriggerGet**
> []WebhookLastTrigger ProjectsProjectIdWebhookLasttriggerGet(ctx, projectId)
Get project webhook policy last trigger info

This endpoint returns last trigger information of project webhook policy. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| Relevant project ID. | 

### Return type

[**[]WebhookLastTrigger**](WebhookLastTrigger.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectIdWebhookPoliciesGet**
> []WebhookPolicy ProjectsProjectIdWebhookPoliciesGet(ctx, projectId)
List project webhook policies.

This endpoint returns webhook policies of a project. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| Relevant project ID. | 

### Return type

[**[]WebhookPolicy**](WebhookPolicy.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectIdWebhookPoliciesPolicyIdDelete**
> ProjectsProjectIdWebhookPoliciesPolicyIdDelete(ctx, projectId, policyId)
Delete webhook policy of a project

This endpoint is aimed to delete webhookpolicy of a project. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| Relevant project ID. | 
  **policyId** | **int64**| The id of webhook policy. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectIdWebhookPoliciesPolicyIdGet**
> WebhookPolicy ProjectsProjectIdWebhookPoliciesPolicyIdGet(ctx, projectId, policyId)
Get project webhook policy

This endpoint returns specified webhook policy of a project. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| Relevant project ID. | 
  **policyId** | **int64**| The id of webhook policy. | 

### Return type

[**WebhookPolicy**](WebhookPolicy.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectIdWebhookPoliciesPolicyIdPut**
> ProjectsProjectIdWebhookPoliciesPolicyIdPut(ctx, body, projectId, policyId)
Update webhook policy of a project.

This endpoint is aimed to update the webhook policy of a project. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**WebhookPolicy**](WebhookPolicy.md)| All properties needed except &quot;id&quot;, &quot;project_id&quot;, &quot;creation_time&quot;, &quot;update_time&quot;. | 
  **projectId** | **int64**| Relevant project ID. | 
  **policyId** | **int64**| The id of webhook policy. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectIdWebhookPoliciesPost**
> ProjectsProjectIdWebhookPoliciesPost(ctx, body, projectId)
Create project webhook policy.

This endpoint create a webhook policy if the project does not have one. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**WebhookPolicy**](WebhookPolicy.md)| Properties &quot;targets&quot; and &quot;event_types&quot; needed. | 
  **projectId** | **int64**| Relevant project ID | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectIdWebhookPoliciesTestPost**
> ProjectsProjectIdWebhookPoliciesTestPost(ctx, body, projectId)
Test project webhook connection

This endpoint tests webhook connection of a project. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**WebhookPolicy**](WebhookPolicy.md)| Only property &quot;targets&quot; needed. | 
  **projectId** | **int64**| Relevant project ID. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QuotasGet**
> []Quota QuotasGet(ctx, optional)
List quotas

List quotas

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProductsApiQuotasGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductsApiQuotasGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reference** | **optional.String**| The reference type of quota. | 
 **referenceId** | **optional.String**| The reference id of quota. | 
 **sort** | **optional.String**| Sort method, valid values include: &#x27;hard.resource_name&#x27;, &#x27;-hard.resource_name&#x27;, &#x27;used.resource_name&#x27;, &#x27;-used.resource_name&#x27;. Here &#x27;-&#x27; stands for descending order, resource_name should be the real resource name of the quota.  | 
 **page** | **optional.Int32**| The page number, default is 1. | 
 **pageSize** | **optional.Int32**| The size of per page, default is 10, maximum is 100. | 

### Return type

[**[]Quota**](Quota.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QuotasIdGet**
> Quota QuotasIdGet(ctx, id)
Get the specified quota

Get the specified quota

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Quota ID | 

### Return type

[**Quota**](Quota.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QuotasIdPut**
> QuotasIdPut(ctx, body, id)
Update the specified quota

Update hard limits of the specified quota

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**QuotaUpdateReq**](QuotaUpdateReq.md)| The new hard limits for the quota | 
  **id** | **int32**| Quota ID | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegistriesGet**
> []Registry RegistriesGet(ctx, optional)
List registries.

This endpoint let user list filtered registries by name, if name is nil, list returns all registries. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProductsApiRegistriesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductsApiRegistriesGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| Registry&#x27;s name. | 

### Return type

[**[]Registry**](Registry.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegistriesIdDelete**
> RegistriesIdDelete(ctx, id)
Delete specific registry.

This endpoint is for to delete specific registry. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| The registry&#x27;s ID. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegistriesIdGet**
> Registry RegistriesIdGet(ctx, id)
Get registry.

This endpoint is for get specific registry.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| The registry ID. | 

### Return type

[**Registry**](Registry.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegistriesIdInfoGet**
> RegistryInfo RegistriesIdInfoGet(ctx, id)
Get registry info.

Get the info of one specific registry.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| The registry ID. | 

### Return type

[**RegistryInfo**](RegistryInfo.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegistriesIdNamespaceGet**
> []Namespace RegistriesIdNamespaceGet(ctx, id, optional)
List namespaces of registry

This endpoint let user list namespaces of registry according to query. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The registry ID. | 
 **optional** | ***ProductsApiRegistriesIdNamespaceGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductsApiRegistriesIdNamespaceGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| The name of namespace. | 

### Return type

[**[]Namespace**](Namespace.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegistriesIdPut**
> RegistriesIdPut(ctx, body, id)
Update a given registry.

This endpoint is for update a given registry. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PutRegistry**](PutRegistry.md)| Updates registry. | 
  **id** | **int64**| The registry&#x27;s ID. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegistriesPingPost**
> RegistriesPingPost(ctx, body)
Ping status of a registry.

This endpoint checks status of a registry, the registry can be given by ID or URL (together with credential) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Registry**](Registry.md)| Registry to ping. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegistriesPost**
> RegistriesPost(ctx, body)
Create a new registry.

This endpoint is for user to create a new registry. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Registry**](Registry.md)| New created registry. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplicationAdaptersGet**
> []string ReplicationAdaptersGet(ctx, )
List supported adapters.

This endpoint let user list supported adapters. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

**[]string**

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplicationExecutionsGet**
> []ReplicationExecution ReplicationExecutionsGet(ctx, optional)
List replication executions.

This endpoint let user list replication executions. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProductsApiReplicationExecutionsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductsApiReplicationExecutionsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **policyId** | **optional.Int32**| The policy ID. | 
 **status** | **optional.String**| The execution status. | 
 **trigger** | **optional.String**| The trigger mode. | 
 **page** | **optional.Int32**| The page. | 
 **pageSize** | **optional.Int32**| The page size. | 

### Return type

[**[]ReplicationExecution**](ReplicationExecution.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplicationExecutionsIdGet**
> ReplicationExecution ReplicationExecutionsIdGet(ctx, id)
Get the execution of the replication.

This endpoint is for user to get one execution of the replication. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| The execution ID. | 

### Return type

[**ReplicationExecution**](ReplicationExecution.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplicationExecutionsIdPut**
> ReplicationExecutionsIdPut(ctx, id)
Stop the execution of the replication.

This endpoint is for user to stop one execution of the replication. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| The execution ID. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplicationExecutionsIdTasksGet**
> []ReplicationTask ReplicationExecutionsIdTasksGet(ctx, id)
Get the task list of one execution.

This endpoint is for user to get the task list of one execution. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| The execution ID. | 

### Return type

[**[]ReplicationTask**](ReplicationTask.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplicationExecutionsIdTasksTaskIdLogGet**
> ReplicationExecutionsIdTasksTaskIdLogGet(ctx, id, taskId)
Get the log of one task.

This endpoint is for user to get the log of one task. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| The execution ID. | 
  **taskId** | **int64**| The task ID. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplicationExecutionsPost**
> ReplicationExecutionsPost(ctx, body)
Start one execution of the replication.

This endpoint is for user to start one execution of the replication. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ReplicationExecution**](ReplicationExecution.md)| The execution that needs to be started, only the property &quot;policy_id&quot; is needed. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplicationPoliciesGet**
> []ReplicationPolicy ReplicationPoliciesGet(ctx, optional)
List replication policies

This endpoint let user list replication policies 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProductsApiReplicationPoliciesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductsApiReplicationPoliciesGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| The replication policy name. | 
 **page** | **optional.Int32**| The page nubmer. | 
 **pageSize** | **optional.Int32**| The size of per page. | 

### Return type

[**[]ReplicationPolicy**](ReplicationPolicy.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplicationPoliciesIdDelete**
> ReplicationPoliciesIdDelete(ctx, id)
Delete the replication policy specified by ID.

Delete the replication policy specified by ID. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| Replication policy ID | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplicationPoliciesIdGet**
> ReplicationPolicy ReplicationPoliciesIdGet(ctx, id)
Get replication policy.

This endpoint let user get replication policy by specific ID. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| policy ID | 

### Return type

[**ReplicationPolicy**](ReplicationPolicy.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplicationPoliciesIdPut**
> ReplicationPoliciesIdPut(ctx, body, id)
Update the replication policy

This endpoint let user update policy. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ReplicationPolicy**](ReplicationPolicy.md)| The replication policy model. | 
  **id** | **int64**| policy ID | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplicationPoliciesPost**
> ReplicationPoliciesPost(ctx, body)
Create a replication policy

This endpoint let user create a replication policy 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ReplicationPolicy**](ReplicationPolicy.md)| The policy model. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesGet**
> []Repository RepositoriesGet(ctx, projectId, optional)
Get repositories accompany with relevant project and repo name.

This endpoint lets user search repositories accompanying with relevant project ID and repo name. Repositories can be sorted by repo name, creation_time, update_time in either ascending or descending order. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int32**| Relevant project ID. | 
 **optional** | ***ProductsApiRepositoriesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductsApiRepositoriesGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **q** | **optional.String**| Repo name for filtering results. | 
 **sort** | **optional.String**| Sort method, valid values include: &#x27;name&#x27;, &#x27;-name&#x27;, &#x27;creation_time&#x27;, &#x27;-creation_time&#x27;, &#x27;update_time&#x27;, &#x27;-update_time&#x27;. Here &#x27;-&#x27; stands for descending order.  | 
 **labelId** | **optional.Int32**| The ID of label used to filter the result. | 
 **page** | **optional.Int32**| The page number, default is 1. | 
 **pageSize** | **optional.Int32**| The size of per page, default is 10, maximum is 100. | 

### Return type

[**[]Repository**](Repository.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesRepoNameDelete**
> RepositoriesRepoNameDelete(ctx, repoName)
Delete a repository.

This endpoint let user delete a repository with name. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoName** | **string**| The name of repository which will be deleted. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesRepoNameLabelsGet**
> []Label RepositoriesRepoNameLabelsGet(ctx, repoName)
Get labels of a repository.

Get labels of a repository specified by the repo_name. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoName** | **string**| The name of repository. | 

### Return type

[**[]Label**](Label.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesRepoNameLabelsLabelIdDelete**
> RepositoriesRepoNameLabelsLabelIdDelete(ctx, repoName, labelId)
Delete label from the repository.

Delete the label from the repository specified by the repo_name. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoName** | **string**| The name of repository. | 
  **labelId** | **int32**| The ID of label. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesRepoNameLabelsPost**
> RepositoriesRepoNameLabelsPost(ctx, body, repoName)
Add a label to the repository.

Add a label to the repository. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Label**](Label.md)| Only the ID property is required. | 
  **repoName** | **string**| The name of repository. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesRepoNamePut**
> RepositoriesRepoNamePut(ctx, body, repoName)
Update description of the repository.

This endpoint is used to update description of the repository. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RepositoryDescription**](RepositoryDescription.md)| The description of the repository. | 
  **repoName** | **string**| The name of repository which will be deleted. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesRepoNameSignaturesGet**
> []RepoSignature RepositoriesRepoNameSignaturesGet(ctx, repoName)
Get signature information of a repository

This endpoint aims to retrieve signature information of a repository, the data is from the nested notary instance of Harbor. If the repository does not have any signature information in notary, this API will return an empty list with response code 200, instead of 404 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoName** | **string**| repository name. | 

### Return type

[**[]RepoSignature**](RepoSignature.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesRepoNameTagsGet**
> []DetailedTag RepositoriesRepoNameTagsGet(ctx, repoName, optional)
Get tags of a relevant repository.

This endpoint aims to retrieve tags from a relevant repository. If deployed with Notary, the signature property of response represents whether the image is singed or not. If the property is null, the image is unsigned. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoName** | **string**| Relevant repository name. | 
 **optional** | ***ProductsApiRepositoriesRepoNameTagsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductsApiRepositoriesRepoNameTagsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **labelId** | **optional.String**| A label ID. | 
 **detail** | **optional.Bool**| Bool value indicating whether return detailed information of the tag, such as vulnerability scan info, if set to false, only tag name is returned. | 

### Return type

[**[]DetailedTag**](DetailedTag.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesRepoNameTagsPost**
> RepositoriesRepoNameTagsPost(ctx, body, repoName)
Retag an image

This endpoint tags an existing image with another tag in this repo, source images can be in different repos or projects. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RetagReq**](RetagReq.md)| Request to give source image and target tag. | 
  **repoName** | **string**| Relevant repository name. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesRepoNameTagsTagDelete**
> RepositoriesRepoNameTagsTagDelete(ctx, repoName, tag)
Delete a tag in a repository.

This endpoint let user delete tags with repo name and tag. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoName** | **string**| The name of repository which will be deleted. | 
  **tag** | **string**| Tag of a repository. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesRepoNameTagsTagGet**
> DetailedTag RepositoriesRepoNameTagsTagGet(ctx, repoName, tag)
Get the tag of the repository.

This endpoint aims to retrieve the tag of the repository. If deployed with Notary, the signature property of response represents whether the image is singed or not. If the property is null, the image is unsigned. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoName** | **string**| Relevant repository name. | 
  **tag** | **string**| Tag of the repository. | 

### Return type

[**DetailedTag**](DetailedTag.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesRepoNameTagsTagLabelsGet**
> []Label RepositoriesRepoNameTagsTagLabelsGet(ctx, repoName, tag)
Get labels of an image.

Get labels of an image specified by the repo_name and tag. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoName** | **string**| The name of repository. | 
  **tag** | **string**| The tag of the image. | 

### Return type

[**[]Label**](Label.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesRepoNameTagsTagLabelsLabelIdDelete**
> RepositoriesRepoNameTagsTagLabelsLabelIdDelete(ctx, repoName, tag, labelId)
Delete label from the image.

Delete the label from the image specified by the repo_name and tag. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoName** | **string**| The name of repository. | 
  **tag** | **string**| The tag of the image. | 
  **labelId** | **int32**| The ID of label. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesRepoNameTagsTagLabelsPost**
> RepositoriesRepoNameTagsTagLabelsPost(ctx, body, repoName, tag)
Add a label to image.

Add a label to the image. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Label**](Label.md)| Only the ID property is required. | 
  **repoName** | **string**| The name of repository. | 
  **tag** | **string**| The tag of the image. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesRepoNameTagsTagManifestGet**
> Manifest RepositoriesRepoNameTagsTagManifestGet(ctx, repoName, tag, optional)
Get manifests of a relevant repository.

This endpoint aims to retreive manifests from a relevant repository. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoName** | **string**| Repository name | 
  **tag** | **string**| Tag name | 
 **optional** | ***ProductsApiRepositoriesRepoNameTagsTagManifestGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductsApiRepositoriesRepoNameTagsTagManifestGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **version** | **optional.String**| The version of manifest, valid value are \&quot;v1\&quot; and \&quot;v2\&quot;, default is \&quot;v2\&quot; | 

### Return type

[**Manifest**](Manifest.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesRepoNameTagsTagScanPost**
> RepositoriesRepoNameTagsTagScanPost(ctx, repoName, tag)
Scan the image.

Trigger a scan targeting the artifact identified by the repo_name and tag. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoName** | **string**| Repository name | 
  **tag** | **string**| Tag name | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesRepoNameTagsTagScanUuidLogGet**
> string RepositoriesRepoNameTagsTagScanUuidLogGet(ctx, repoName, tag, uuid)
Get scan log

Get the log text stream for the given artifact and scan action.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoName** | **string**| Repository name | 
  **tag** | **string**| Tag name | 
  **uuid** | **string**| the scan unique identifier | 

### Return type

**string**

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesTopGet**
> []Repository RepositoriesTopGet(ctx, optional)
Get public repositories which are accessed most.

This endpoint aims to let users see the most popular public repositories 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProductsApiRepositoriesTopGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductsApiRepositoriesTopGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**| The number of the requested public repositories, default is 10 if not provided. | 

### Return type

[**[]Repository**](Repository.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetentionsIdExecutionsEidPatch**
> RetentionsIdExecutionsEidPatch(ctx, body, id, eid)
Stop a Retention job

Stop a Retention job, only support \"stop\" action now.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body6**](Body6.md)| The action, only support &quot;stop&quot; now. | 
  **id** | **int64**| Retention ID. | 
  **eid** | **int64**| Retention execution ID. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetentionsIdExecutionsEidTasksGet**
> []RetentionExecutionTask RetentionsIdExecutionsEidTasksGet(ctx, id, eid)
Get Retention job tasks

Get Retention job tasks, each repository as a task.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| Retention ID. | 
  **eid** | **int64**| Retention execution ID. | 

### Return type

[**[]RetentionExecutionTask**](RetentionExecutionTask.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetentionsIdExecutionsEidTasksTidGet**
> string RetentionsIdExecutionsEidTasksTidGet(ctx, id, eid, tid)
Get Retention job task log

Get Retention job task log, tags ratain or deletion detail will be shown in a table.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| Retention ID. | 
  **eid** | **int64**| Retention execution ID. | 
  **tid** | **int64**| Retention execution ID. | 

### Return type

**string**

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetentionsIdExecutionsGet**
> []RetentionExecution RetentionsIdExecutionsGet(ctx, id)
Get a Retention job

Get a Retention job, job status may be delayed before job service schedule it up.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| Retention ID. | 

### Return type

[**[]RetentionExecution**](RetentionExecution.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetentionsIdExecutionsPost**
> RetentionsIdExecutionsPost(ctx, body, id)
Trigger a Retention job

Trigger a Retention job, if dry_run is True, nothing would be deleted actually.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body5**](Body5.md)|  | 
  **id** | **int64**| Retention ID. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetentionsIdGet**
> RetentionPolicy RetentionsIdGet(ctx, id)
Get Retention Policy

Get Retention Policy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| Retention ID. | 

### Return type

[**RetentionPolicy**](RetentionPolicy.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetentionsIdPut**
> RetentionsIdPut(ctx, body, id)
Update Retention Policy

Update Retention Policy, you can reference metadatas API for the policy model. You can check project metadatas to find whether a retention policy is already binded. This method should only be called when retention policy has already binded to project. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RetentionPolicy**](RetentionPolicy.md)|  | 
  **id** | **int64**| Retention ID. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetentionsMetadatasGet**
> RetentionMetadata RetentionsMetadatasGet(ctx, )
Get Retention Metadatas

Get Retention Metadatas.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**RetentionMetadata**](RetentionMetadata.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetentionsPost**
> RetentionsPost(ctx, body)
Create Retention Policy

Create Retention Policy, you can reference metadatas API for the policy model. You can check project metadatas to find whether a retention policy is already binded. This method should only be called when no retention policy binded to project yet. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RetentionPolicy**](RetentionPolicy.md)| Create Retention Policy successfully. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ScannersGet**
> []ScannerRegistration ScannersGet(ctx, )
List scanner registrations

Returns a list of currently configured scanner registrations. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]ScannerRegistration**](ScannerRegistration.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ScannersPingPost**
> ScannersPingPost(ctx, body)
Tests scanner registration settings

Pings scanner adapter to test endpoint URL and authorization settings. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ScannerRegistrationSettings**](ScannerRegistrationSettings.md)| A scanner registration settings to be tested. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ScannersRegistrationIdGet**
> ScannerRegistration ScannersRegistrationIdGet(ctx, registrationId)
Get a scanner registration details

Retruns the details of the specified scanner registration. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **registrationId** | **string**| The scanner registration identifer. | 

### Return type

[**ScannerRegistration**](ScannerRegistration.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ScannersRegistrationIdMetadataGet**
> ScannerAdapterMetadata ScannersRegistrationIdMetadataGet(ctx, registrationId)
Get the metadata of the specified scanner registration

Get the metadata of the specified scanner registration, including the capabilities and customzied properties. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **registrationId** | **string**| The scanner registration identifier. | 

### Return type

[**ScannerAdapterMetadata**](ScannerAdapterMetadata.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ScansAllMetricsGet**
> Stats ScansAllMetricsGet(ctx, )
Get the metrics of the latest scan all process

Get the metrics of the latest scan all process

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Stats**](Stats.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ScansScheduleMetricsGet**
> Stats ScansScheduleMetricsGet(ctx, )
Get the metrics of the latest scheduled scan all process

Get the metrics of the latest scheduled scan all process

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Stats**](Stats.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchGet**
> []Search SearchGet(ctx, q)
Search for projects, repositories and helm charts

The Search endpoint returns information about the projects ,repositories  and helm charts offered at public status or related to the current logged in user. The response includes the project, repository list and charts in a proper display order. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **q** | **string**| Search parameter for project and repository name. | 

### Return type

[**[]Search**](Search.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StatisticsGet**
> StatisticMap StatisticsGet(ctx, )
Get projects number and repositories number relevant to the user

This endpoint is aimed to statistic all of the projects number and repositories number relevant to the logined user, also the public projects number and repositories number. If the user is admin, he can also get total projects number and total repositories number. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**StatisticMap**](StatisticMap.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SystemCVEWhitelistGet**
> CveWhitelist SystemCVEWhitelistGet(ctx, )
Get the system level whitelist of CVE.

Get the system level whitelist of CVE.  This API can be called by all authenticated users.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**CveWhitelist**](CVEWhitelist.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SystemCVEWhitelistPut**
> SystemCVEWhitelistPut(ctx, optional)
Update the system level whitelist of CVE.

This API overwrites the system level whitelist of CVE with the list in request body.  Only system Admin has permission to call this API.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProductsApiSystemCVEWhitelistPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductsApiSystemCVEWhitelistPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CveWhitelist**](CveWhitelist.md)| The whitelist with new content | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SystemGcGet**
> []GcResult SystemGcGet(ctx, )
Get gc results.

This endpoint let user get latest ten gc results.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]GcResult**](GCResult.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SystemGcIdGet**
> GcResult SystemGcIdGet(ctx, id)
Get gc status.

This endpoint let user get gc status filtered by specific ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| Relevant job ID | 

### Return type

[**GcResult**](GCResult.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SystemGcIdLogGet**
> string SystemGcIdLogGet(ctx, id)
Get gc job log.

This endpoint let user get gc job logs filtered by specific ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| Relevant job ID | 

### Return type

**string**

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SystemGcScheduleGet**
> AdminJobSchedule SystemGcScheduleGet(ctx, )
Get gc's schedule.

This endpoint is for get schedule of gc job.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**AdminJobSchedule**](AdminJobSchedule.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SystemGcSchedulePost**
> SystemGcSchedulePost(ctx, body)
Create a gc schedule.

This endpoint is for update gc schedule. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AdminJobSchedule**](AdminJobSchedule.md)| Updates of gc&#x27;s schedule. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SystemGcSchedulePut**
> SystemGcSchedulePut(ctx, body)
Update gc's schedule.

This endpoint is for update gc schedule. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AdminJobSchedule**](AdminJobSchedule.md)| Updates of gc&#x27;s schedule. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SystemOidcPingPost**
> SystemOidcPingPost(ctx, body)
Test the OIDC endpoint.

Test the OIDC endpoint, the setting of the endpoint is provided in the request.  This API can only be called by system admin.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body4**](Body4.md)| Request body for OIDC endpoint to be tested. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SystemScanAllScheduleGet**
> AdminJobSchedule SystemScanAllScheduleGet(ctx, )
Get scan_all's schedule.

This endpoint is for getting a schedule for the scan all job, which scans all of images in Harbor.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**AdminJobSchedule**](AdminJobSchedule.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SystemScanAllSchedulePost**
> SystemScanAllSchedulePost(ctx, body)
Create a schedule or a manual trigger for the scan all job.

This endpoint is for creating a schedule or a manual trigger for the scan all job, which scans all of images in Harbor. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AdminJobSchedule**](AdminJobSchedule.md)| Create a schedule or a manual trigger for the scan all job. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SystemScanAllSchedulePut**
> SystemScanAllSchedulePut(ctx, body)
Update scan all's schedule.

This endpoint is for updating the schedule of scan all job, which scans all of images in Harbor. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AdminJobSchedule**](AdminJobSchedule.md)| Updates the schedule of scan all job, which scans all of images in Harbor. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SysteminfoGet**
> GeneralInfo SysteminfoGet(ctx, )
Get general system info

This API is for retrieving general system info, this can be called by anonymous request. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GeneralInfo**](GeneralInfo.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SysteminfoGetcertGet**
> SysteminfoGetcertGet(ctx, )
Get default root certificate.

This endpoint is for downloading a default root certificate. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SysteminfoVolumesGet**
> SystemInfo SysteminfoVolumesGet(ctx, )
Get system volume info (total/free size).

This endpoint is for retrieving system volume info that only provides for admin user. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SystemInfo**](SystemInfo.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsergroupsGet**
> []UserGroup UsergroupsGet(ctx, )
Get all user groups information

Get all user groups information

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]UserGroup**](UserGroup.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsergroupsGroupIdDelete**
> UsergroupsGroupIdDelete(ctx, groupId)
Delete user group

Delete user group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsergroupsGroupIdGet**
> UserGroup UsergroupsGroupIdGet(ctx, groupId)
Get user group information

Get user group information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **int64**| Group ID | 

### Return type

[**UserGroup**](UserGroup.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsergroupsGroupIdPut**
> UsergroupsGroupIdPut(ctx, groupId, optional)
Update group information

Update user group information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **int64**| Group ID | 
 **optional** | ***ProductsApiUsergroupsGroupIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductsApiUsergroupsGroupIdPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UserGroup**](UserGroup.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsergroupsPost**
> UsergroupsPost(ctx, optional)
Create user group

Create user group information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProductsApiUsergroupsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductsApiUsergroupsPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of UserGroup**](UserGroup.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersCurrentGet**
> User UsersCurrentGet(ctx, )
Get current user info.

This endpoint is to get the current user information. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**User**](User.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersCurrentPermissionsGet**
> []Permission UsersCurrentPermissionsGet(ctx, optional)
Get current user permissions.

This endpoint is to get the current user permissions. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProductsApiUsersCurrentPermissionsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductsApiUsersCurrentPermissionsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scope** | **optional.String**| Get permissions of the scope | 
 **relative** | **optional.Bool**| If true, the resources in the response are relative to the scope, eg for resource &#x27;/project/1/repository&#x27; if relative is &#x27;true&#x27; then the resource in response will be &#x27;repository&#x27;.  | 

### Return type

[**[]Permission**](Permission.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersGet**
> []User UsersGet(ctx, optional)
Get registered users of Harbor.

This endpoint is for user to search registered users, support for filtering results with username.Notice, by now this operation is only for administrator. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProductsApiUsersGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductsApiUsersGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **optional.String**| Username for filtering results. | 
 **email** | **optional.String**| Email for filtering results. | 
 **page** | **optional.Int32**| The page number, default is 1. | 
 **pageSize** | **optional.Int32**| The size of per page. | 

### Return type

[**[]User**](User.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersPost**
> UsersPost(ctx, body)
Creates a new user account.

This endpoint is to create a user if the user does not already exist. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**User**](User.md)| New created user. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersSearchGet**
> []UserSearch UsersSearchGet(ctx, username, optional)
Search users by username

This endpoint is to search the users by username. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**| Username for filtering results. | 
 **optional** | ***ProductsApiUsersSearchGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductsApiUsersSearchGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| The page number, default is 1. | 
 **pageSize** | **optional.Int32**| The size of per page. | 

### Return type

[**[]UserSearch**](UserSearch.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersUserIdCliSecretPut**
> UsersUserIdCliSecretPut(ctx, body, userId)
Set CLI secret for a user.

This endpoint let user generate a new CLI secret for himself.  This API only works when auth mode is set to 'OIDC'. Once this API returns with successful status, the old secret will be invalid, as there will be only one CLI secret for a user. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body**](Body.md)| JSON object that includes the new secret | 
  **userId** | **int32**| User ID | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersUserIdDelete**
> UsersUserIdDelete(ctx, userId)
Mark a registered user as be removed.

This endpoint let administrator of Harbor mark a registered user as be removed.It actually won't be deleted from DB. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **int32**| User ID for marking as to be removed. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersUserIdGet**
> User UsersUserIdGet(ctx, userId)
Get a user's profile.

Get user's profile with user id. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **int32**| Registered user ID | 

### Return type

[**User**](User.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersUserIdPasswordPut**
> UsersUserIdPasswordPut(ctx, body, userId)
Change the password on a user that already exists.

This endpoint is for user to update password. Users with the admin role can change any user's password. Guest users can change only their own password. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Password**](Password.md)| Password to be updated, the attribute &#x27;old_password&#x27; is optional when the API is called by the system administrator. | 
  **userId** | **int32**| Registered user ID. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersUserIdPut**
> UsersUserIdPut(ctx, body, userId)
Update a registered user to change his profile.

This endpoint let a registered user change his profile. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UserProfile**](UserProfile.md)| Only email, realname and comment can be modified. | 
  **userId** | **int32**| Registered user ID | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersUserIdSysadminPut**
> UsersUserIdSysadminPut(ctx, body, userId)
Update a registered user to change to be an administrator of Harbor.

This endpoint let a registered user change to be an administrator of Harbor. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**HasAdminRole**](HasAdminRole.md)| Toggle a user to admin or not. | 
  **userId** | **int32**| Registered user ID | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

