# Configurations

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthMode** | **string** | The auth mode of current system, such as \&quot;db_auth\&quot;, \&quot;ldap_auth\&quot; | [optional] [default to null]
**CountPerProject** | **string** | The default count quota for the new created projects. | [optional] [default to null]
**EmailFrom** | **string** | The sender name for Email notification. | [optional] [default to null]
**EmailHost** | **string** | The hostname of SMTP server that sends Email notification. | [optional] [default to null]
**EmailPort** | **int32** | The port of SMTP server. | [optional] [default to null]
**EmailIdentity** | **string** | By default it&#x27;s empty so the email_username is picked. | [optional] [default to null]
**EmailUsername** | **string** | The username for authenticate against SMTP server. | [optional] [default to null]
**EmailSsl** | **bool** | When it&#x27;s set to true the system will access Email server via TLS by default.  If it&#x27;s set to false, it still will handle \&quot;STARTTLS\&quot; from server side. | [optional] [default to null]
**EmailInsecure** | **bool** | Whether or not the certificate will be verified when Harbor tries to access the email server. | [optional] [default to null]
**LdapUrl** | **string** | The URL of LDAP server. | [optional] [default to null]
**LdapBaseDn** | **string** | The Base DN for LDAP binding. | [optional] [default to null]
**LdapFilter** | **string** | The filter for LDAP binding. | [optional] [default to null]
**LdapScope** | **int32** | 0-LDAP_SCOPE_BASE, 1-LDAP_SCOPE_ONELEVEL, 2-LDAP_SCOPE_SUBTREE | [optional] [default to null]
**LdapUid** | **string** | The attribute which is used as identity for the LDAP binding, such as \&quot;CN\&quot; or \&quot;SAMAccountname\&quot; | [optional] [default to null]
**LdapSearchDn** | **string** | The DN of the user to do the search. | [optional] [default to null]
**LdapTimeout** | **int32** | timeout in seconds for connection to LDAP server. | [optional] [default to null]
**LdapGroupAttributeName** | **string** | The attribute which is used as identity of the LDAP group, default is cn. | [optional] [default to null]
**LdapGroupBaseDn** | **string** | The base DN to search LDAP group. | [optional] [default to null]
**LdapGroupSearchFilter** | **string** | The filter to search the ldap group. | [optional] [default to null]
**LdapGroupSearchScope** | **int32** | The scope to search ldap. &#x27;0-LDAP_SCOPE_BASE, 1-LDAP_SCOPE_ONELEVEL, 2-LDAP_SCOPE_SUBTREE&#x27; | [optional] [default to null]
**LdapGroupAdminDn** | **string** | Specify the ldap group which have the same privilege with Harbor admin. | [optional] [default to null]
**OidcClientId** | **string** | The client id of the OIDC. | [optional] [default to null]
**OidcClientSecret** | **string** | The client secret of the OIDC. | [optional] [default to null]
**OidcEndpoint** | **string** | The URL of an OIDC-complaint server, must start with &#x27;https://&#x27;. | [optional] [default to null]
**OidcName** | **string** | The name of the OIDC provider. | [optional] [default to null]
**OidcScope** | **string** | The scope sent to OIDC server during authentication, should be separated by comma. It has to contain “openid”, and “offline_access”. If you are using google, please remove “offline_access” from this field. | [optional] [default to null]
**OidcVerifyCert** | **bool** | Whether verify your OIDC server certificate, disable it if your OIDC server is hosted via self-hosted certificate. | [optional] [default to null]
**ProjectCreationRestriction** | **string** | This attribute restricts what users have the permission to create project.  It can be \&quot;everyone\&quot; or \&quot;adminonly\&quot;. | [optional] [default to null]
**QuotaPerProjectEnable** | **bool** | This attribute indicates whether quota per project enabled in harbor | [optional] [default to null]
**ReadOnly** | **bool** | &#x27;docker push&#x27; is prohibited by Harbor if you set it to true.    | [optional] [default to null]
**SelfRegistration** | **bool** | Whether the Harbor instance supports self-registration.  If it&#x27;s set to false, admin need to add user to the instance. | [optional] [default to null]
**StoragePerProject** | **string** | The default storage quota for the new created projects. | [optional] [default to null]
**TokenExpiration** | **int32** | The expiration time of the token for internal Registry, in minutes. | [optional] [default to null]
**VerifyRemoteCert** | **bool** | Whether or not the certificate will be verified when Harbor tries to access a remote Harbor instance for replication. | [optional] [default to null]
**ScanAllPolicy** | [***ConfigurationsScanAllPolicy**](Configurations_scan_all_policy.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

