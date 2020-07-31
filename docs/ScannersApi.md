# {{classname}}

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProjectsProjectIdScannerCandidatesGet**](ScannersApi.md#ProjectsProjectIdScannerCandidatesGet) | **Get** /projects/{project_id}/scanner/candidates | Get scanner registration candidates for configurating project level scanner
[**ProjectsProjectIdScannerGet**](ScannersApi.md#ProjectsProjectIdScannerGet) | **Get** /projects/{project_id}/scanner | Get project level scanner
[**ProjectsProjectIdScannerPut**](ScannersApi.md#ProjectsProjectIdScannerPut) | **Put** /projects/{project_id}/scanner | Configure scanner for the specified project
[**ScannersGet**](ScannersApi.md#ScannersGet) | **Get** /scanners | List scanner registrations
[**ScannersPingPost**](ScannersApi.md#ScannersPingPost) | **Post** /scanners/ping | Tests scanner registration settings
[**ScannersPost**](ScannersApi.md#ScannersPost) | **Post** /scanners | Create a scanner registration
[**ScannersRegistrationIdDelete**](ScannersApi.md#ScannersRegistrationIdDelete) | **Delete** /scanners/{registration_id} | Delete a scanner registration
[**ScannersRegistrationIdGet**](ScannersApi.md#ScannersRegistrationIdGet) | **Get** /scanners/{registration_id} | Get a scanner registration details
[**ScannersRegistrationIdMetadataGet**](ScannersApi.md#ScannersRegistrationIdMetadataGet) | **Get** /scanners/{registration_id}/metadata | Get the metadata of the specified scanner registration
[**ScannersRegistrationIdPatch**](ScannersApi.md#ScannersRegistrationIdPatch) | **Patch** /scanners/{registration_id} | Set system default scanner registration
[**ScannersRegistrationIdPut**](ScannersApi.md#ScannersRegistrationIdPut) | **Put** /scanners/{registration_id} | Update a scanner registration

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

# **ProjectsProjectIdScannerPut**
> ProjectsProjectIdScannerPut(ctx, body, projectId)
Configure scanner for the specified project

Set one of the system configured scanner registration as the indepndent scanner of the specified project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ProjectScanner**](ProjectScanner.md)|  | 
  **projectId** | **int64**| The project identifier. | 

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

# **ScannersPost**
> ScannersPost(ctx, body)
Create a scanner registration

Creats a new scanner registration with the given data. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ScannerRegistrationReq**](ScannerRegistrationReq.md)| A scanner registration to be created. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ScannersRegistrationIdDelete**
> ScannerRegistration ScannersRegistrationIdDelete(ctx, registrationId)
Delete a scanner registration

Deletes the specified scanner registration. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **registrationId** | **string**| The scanner registration identifier. | 

### Return type

[**ScannerRegistration**](ScannerRegistration.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

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

# **ScannersRegistrationIdPatch**
> ScannersRegistrationIdPatch(ctx, body, registrationId)
Set system default scanner registration

Set the specified scanner registration as the system default one. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IsDefault**](IsDefault.md)|  | 
  **registrationId** | **string**| The scanner registration identifier. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ScannersRegistrationIdPut**
> ScannersRegistrationIdPut(ctx, body, registrationId)
Update a scanner registration

Updates the specified scanner registration. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ScannerRegistrationReq**](ScannerRegistrationReq.md)| A scanner registraiton to be updated. | 
  **registrationId** | **string**| The scanner registration identifier. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

