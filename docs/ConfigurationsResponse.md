# ConfigurationsResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthMode** | [***StringConfigItem**](StringConfigItem.md) |  | [optional] [default to null]
**CountPerProject** | [***IntegerConfigItem**](IntegerConfigItem.md) |  | [optional] [default to null]
**EmailFrom** | [***StringConfigItem**](StringConfigItem.md) |  | [optional] [default to null]
**EmailHost** | [***StringConfigItem**](StringConfigItem.md) |  | [optional] [default to null]
**EmailPort** | [***IntegerConfigItem**](IntegerConfigItem.md) |  | [optional] [default to null]
**EmailIdentity** | [***StringConfigItem**](StringConfigItem.md) |  | [optional] [default to null]
**EmailUsername** | [***StringConfigItem**](StringConfigItem.md) |  | [optional] [default to null]
**EmailSsl** | [***BoolConfigItem**](BoolConfigItem.md) |  | [optional] [default to null]
**EmailInsecure** | [***BoolConfigItem**](BoolConfigItem.md) |  | [optional] [default to null]
**LdapUrl** | [***StringConfigItem**](StringConfigItem.md) |  | [optional] [default to null]
**LdapBaseDn** | [***StringConfigItem**](StringConfigItem.md) |  | [optional] [default to null]
**LdapFilter** | [***StringConfigItem**](StringConfigItem.md) |  | [optional] [default to null]
**LdapScope** | **int32** | 0-LDAP_SCOPE_BASE, 1-LDAP_SCOPE_ONELEVEL, 2-LDAP_SCOPE_SUBTREE | [optional] [default to null]
**LdapUid** | [***StringConfigItem**](StringConfigItem.md) |  | [optional] [default to null]
**LdapSearchDn** | **string** | The DN of the user to do the search. | [optional] [default to null]
**LdapTimeout** | [***IntegerConfigItem**](IntegerConfigItem.md) |  | [optional] [default to null]
**LdapGroupAttributeName** | [***StringConfigItem**](StringConfigItem.md) |  | [optional] [default to null]
**LdapGroupBaseDn** | [***StringConfigItem**](StringConfigItem.md) |  | [optional] [default to null]
**LdapGroupSearchFilter** | [***StringConfigItem**](StringConfigItem.md) |  | [optional] [default to null]
**LdapGroupSearchScope** | [***IntegerConfigItem**](IntegerConfigItem.md) |  | [optional] [default to null]
**LdapGroupAdminDn** | [***StringConfigItem**](StringConfigItem.md) |  | [optional] [default to null]
**OidcClientId** | [***StringConfigItem**](StringConfigItem.md) |  | [optional] [default to null]
**OidcEndpoint** | [***StringConfigItem**](StringConfigItem.md) |  | [optional] [default to null]
**OidcName** | [***StringConfigItem**](StringConfigItem.md) |  | [optional] [default to null]
**OidcScope** | [***StringConfigItem**](StringConfigItem.md) |  | [optional] [default to null]
**OidcVerifyCert** | [***BoolConfigItem**](BoolConfigItem.md) |  | [optional] [default to null]
**ProjectCreationRestriction** | [***StringConfigItem**](StringConfigItem.md) |  | [optional] [default to null]
**QuotaPerProjectEnable** | [***BoolConfigItem**](BoolConfigItem.md) |  | [optional] [default to null]
**ReadOnly** | [***BoolConfigItem**](BoolConfigItem.md) |  | [optional] [default to null]
**SelfRegistration** | [***BoolConfigItem**](BoolConfigItem.md) |  | [optional] [default to null]
**StoragePerProject** | [***IntegerConfigItem**](IntegerConfigItem.md) |  | [optional] [default to null]
**TokenExpiration** | [***IntegerConfigItem**](IntegerConfigItem.md) |  | [optional] [default to null]
**VerifyRemoteCert** | [***BoolConfigItem**](BoolConfigItem.md) |  | [optional] [default to null]
**ScanAllPolicy** | [***ConfigurationsScanAllPolicy**](Configurations_scan_all_policy.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

