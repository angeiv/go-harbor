# {{classname}}

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QuotasIdGet**](QuotaApi.md#QuotasIdGet) | **Get** /quotas/{id} | Get the specified quota
[**QuotasIdPut**](QuotaApi.md#QuotasIdPut) | **Put** /quotas/{id} | Update the specified quota

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

