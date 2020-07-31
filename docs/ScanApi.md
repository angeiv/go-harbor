# {{classname}}

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RepositoriesRepoNameTagsTagScanGet**](ScanApi.md#RepositoriesRepoNameTagsTagScanGet) | **Get** /repositories/{repo_name}/tags/{tag}/scan | Get the scan report
[**RepositoriesRepoNameTagsTagScanPost**](ScanApi.md#RepositoriesRepoNameTagsTagScanPost) | **Post** /repositories/{repo_name}/tags/{tag}/scan | Scan the image.
[**RepositoriesRepoNameTagsTagScanUuidLogGet**](ScanApi.md#RepositoriesRepoNameTagsTagScanUuidLogGet) | **Get** /repositories/{repo_name}/tags/{tag}/scan/{uuid}/log | Get scan log
[**ScansAllMetricsGet**](ScanApi.md#ScansAllMetricsGet) | **Get** /scans/all/metrics | Get the metrics of the latest scan all process
[**ScansScheduleMetricsGet**](ScanApi.md#ScansScheduleMetricsGet) | **Get** /scans/schedule/metrics | Get the metrics of the latest scheduled scan all process

# **RepositoriesRepoNameTagsTagScanGet**
> Report RepositoriesRepoNameTagsTagScanGet(ctx, repoName, tag, optional)
Get the scan report

Retrieve the scan report for the artifact identified by the repo_name and tag. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoName** | **string**| Repository name | 
  **tag** | **string**| Tag name | 
 **optional** | ***ScanApiRepositoriesRepoNameTagsTagScanGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ScanApiRepositoriesRepoNameTagsTagScanGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accept** | **optional.String**| Mimetype in header. e.g: \&quot;application/vnd.scanner.adapter.vuln.report.harbor+json; version&#x3D;1.0\&quot;  | 

### Return type

[**Report**](Report.md)

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

