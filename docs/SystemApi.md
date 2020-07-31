# {{classname}}

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SystemCVEWhitelistGet**](SystemApi.md#SystemCVEWhitelistGet) | **Get** /system/CVEWhitelist | Get the system level whitelist of CVE.
[**SystemCVEWhitelistPut**](SystemApi.md#SystemCVEWhitelistPut) | **Put** /system/CVEWhitelist | Update the system level whitelist of CVE.
[**SystemOidcPingPost**](SystemApi.md#SystemOidcPingPost) | **Post** /system/oidc/ping | Test the OIDC endpoint.

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
 **optional** | ***SystemApiSystemCVEWhitelistPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemApiSystemCVEWhitelistPutOpts struct
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

