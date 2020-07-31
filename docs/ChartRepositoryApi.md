# {{classname}}

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChartrepoChartsPost**](ChartRepositoryApi.md#ChartrepoChartsPost) | **Post** /chartrepo/charts | Upload a chart file to the defult &#x27;library&#x27; project.
[**ChartrepoHealthGet**](ChartRepositoryApi.md#ChartrepoHealthGet) | **Get** /chartrepo/health | Check the health of chart repository service.
[**ChartrepoRepoChartsGet**](ChartRepositoryApi.md#ChartrepoRepoChartsGet) | **Get** /chartrepo/{repo}/charts | Get all the charts under the specified project
[**ChartrepoRepoChartsNameDelete**](ChartRepositoryApi.md#ChartrepoRepoChartsNameDelete) | **Delete** /chartrepo/{repo}/charts/{name} | Delete all the versions of the specified chart
[**ChartrepoRepoChartsNameGet**](ChartRepositoryApi.md#ChartrepoRepoChartsNameGet) | **Get** /chartrepo/{repo}/charts/{name} | Get all the versions of the specified chart
[**ChartrepoRepoChartsNameVersionDelete**](ChartRepositoryApi.md#ChartrepoRepoChartsNameVersionDelete) | **Delete** /chartrepo/{repo}/charts/{name}/{version} | Delete the specified chart version
[**ChartrepoRepoChartsNameVersionGet**](ChartRepositoryApi.md#ChartrepoRepoChartsNameVersionGet) | **Get** /chartrepo/{repo}/charts/{name}/{version} | Get the specified chart version
[**ChartrepoRepoChartsNameVersionLabelsGet**](ChartRepositoryApi.md#ChartrepoRepoChartsNameVersionLabelsGet) | **Get** /chartrepo/{repo}/charts/{name}/{version}/labels | Return the attahced labels of chart.
[**ChartrepoRepoChartsNameVersionLabelsIdDelete**](ChartRepositoryApi.md#ChartrepoRepoChartsNameVersionLabelsIdDelete) | **Delete** /chartrepo/{repo}/charts/{name}/{version}/labels/{id} | Remove label from chart.
[**ChartrepoRepoChartsNameVersionLabelsPost**](ChartRepositoryApi.md#ChartrepoRepoChartsNameVersionLabelsPost) | **Post** /chartrepo/{repo}/charts/{name}/{version}/labels | Mark label to chart.
[**ChartrepoRepoChartsPost**](ChartRepositoryApi.md#ChartrepoRepoChartsPost) | **Post** /chartrepo/{repo}/charts | Upload a chart file to the specified project.
[**ChartrepoRepoProvPost**](ChartRepositoryApi.md#ChartrepoRepoProvPost) | **Post** /chartrepo/{repo}/prov | Upload a provance file to the specified project.

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

