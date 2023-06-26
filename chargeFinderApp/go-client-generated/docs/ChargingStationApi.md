# {{classname}}

All URIs are relative to *https://petstore3.swagger.io/api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAvailableStations**](ChargingStationApi.md#GetAvailableStations) | **Get** /availableStations | 
[**OccupiedStations**](ChargingStationApi.md#OccupiedStations) | **Get** /occupiedStations | 

# **GetAvailableStations**
> []ChargingStation GetAvailableStations(ctx, )


Multiple status values can be provided with comma separated strings

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

# **OccupiedStations**
> []ChargingStation OccupiedStations(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]ChargingStation**](ChargingStation.md)

### Authorization

[petstore_auth](../README.md#petstore_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

