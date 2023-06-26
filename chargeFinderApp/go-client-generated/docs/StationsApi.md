# {{classname}}

All URIs are relative to *https://petstore3.swagger.io/api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddChargingStation**](StationsApi.md#AddChargingStation) | **Post** /chargingStations | add charging station
[**GetLists**](StationsApi.md#GetLists) | **Get** /chargingStations | 

# **AddChargingStation**
> ChargingStation AddChargingStation(ctx, optional)
add charging station

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StationsApiAddChargingStationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StationsApiAddChargingStationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ChargingStation**](ChargingStation.md)| Created user object | 

### Return type

[**ChargingStation**](ChargingStation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLists**
> []ChargingStation GetLists(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]ChargingStation**](ChargingStation.md)

### Authorization

[petstore_auth](../README.md#petstore_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

