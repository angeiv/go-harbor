# Artifact

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The ID of the artifact | [optional] [default to null]
**Type_** | **string** | The type of the artifact, e.g. image, chart, etc | [optional] [default to null]
**MediaType** | **string** | The media type of the artifact | [optional] [default to null]
**ManifestMediaType** | **string** | The manifest media type of the artifact | [optional] [default to null]
**ProjectId** | **int64** | The ID of the project that the artifact belongs to | [optional] [default to null]
**RepositoryId** | **int64** | The ID of the repository that the artifact belongs to | [optional] [default to null]
**Digest** | **string** | The digest of the artifact | [optional] [default to null]
**Size** | **int64** | The size of the artifact | [optional] [default to null]
**PushTime** | [**time.Time**](time.Time.md) | The push time of the artifact | [optional] [default to null]
**PullTime** | [**time.Time**](time.Time.md) | The latest pull time of the artifact | [optional] [default to null]
**ExtraAttrs** | [***ModelMap**](map.md) |  | [optional] [default to null]
**Annotations** | [***ModelMap**](map.md) |  | [optional] [default to null]
**References** | [**[]Reference**](Reference.md) |  | [optional] [default to null]
**Tags** | [**[]Tag**](Tag.md) |  | [optional] [default to null]
**AdditionLinks** | [***ModelMap**](map.md) |  | [optional] [default to null]
**Labels** | [**[]Label**](Label.md) |  | [optional] [default to null]
**ScanOverview** | [***ModelMap**](map.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

