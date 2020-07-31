# CveWhitelist

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | ID of the whitelist | [optional] [default to null]
**ProjectId** | **int32** | ID of the project which the whitelist belongs to.  For system level whitelist this attribute is zero. | [optional] [default to null]
**ExpiresAt** | **int32** | the time for expiration of the whitelist, in the form of seconds since epoch.  This is an optional attribute, if it&#x27;s not set the CVE whitelist does not expire. | [optional] [default to null]
**Items** | [**[]CveWhitelistItem**](CVEWhitelistItem.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

