# \OrdersApi

All URIs are relative to *https://public-api.iqxamplify.com/V1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrUpdateOrder**](OrdersApi.md#CreateOrUpdateOrder) | **Post** /orders/ | Create or Update a order
[**FindOrders**](OrdersApi.md#FindOrders) | **Get** /orders/ | Retrieve all orders
[**FindOrdersById**](OrdersApi.md#FindOrdersById) | **Get** /orders/{id} | Retrieve a single order by id


# **CreateOrUpdateOrder**
> Order CreateOrUpdateOrder($order)

Create or Update a order

Create or Update a order


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **order** | [**Order**](Order.md)| Add or update orders details | 

### Return type

[**Order**](Order.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindOrders**
> Orders FindOrders($page)

Retrieve all orders

Retrieve all orders


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32**| The page number ro return | [optional] 

### Return type

[**Orders**](Orders.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindOrdersById**
> Order FindOrdersById($id)

Retrieve a single order by id

Retrieve a single order by id


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the order to get | 

### Return type

[**Order**](Order.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

