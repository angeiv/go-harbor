# {{classname}}

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RetentionsIdExecutionsEidPatch**](RetentionApi.md#RetentionsIdExecutionsEidPatch) | **Patch** /retentions/{id}/executions/{eid} | Stop a Retention job
[**RetentionsIdExecutionsEidTasksGet**](RetentionApi.md#RetentionsIdExecutionsEidTasksGet) | **Get** /retentions/{id}/executions/{eid}/tasks | Get Retention job tasks
[**RetentionsIdExecutionsEidTasksTidGet**](RetentionApi.md#RetentionsIdExecutionsEidTasksTidGet) | **Get** /retentions/{id}/executions/{eid}/tasks/{tid} | Get Retention job task log
[**RetentionsIdExecutionsGet**](RetentionApi.md#RetentionsIdExecutionsGet) | **Get** /retentions/{id}/executions | Get a Retention job
[**RetentionsIdExecutionsPost**](RetentionApi.md#RetentionsIdExecutionsPost) | **Post** /retentions/{id}/executions | Trigger a Retention job
[**RetentionsIdGet**](RetentionApi.md#RetentionsIdGet) | **Get** /retentions/{id} | Get Retention Policy
[**RetentionsMetadatasGet**](RetentionApi.md#RetentionsMetadatasGet) | **Get** /retentions/metadatas | Get Retention Metadatas
[**RetentionsPost**](RetentionApi.md#RetentionsPost) | **Post** /retentions | Create Retention Policy

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

