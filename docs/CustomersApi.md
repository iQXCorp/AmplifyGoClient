# \CustomersApi

All URIs are relative to *https://public-api.iqxamplify.com/V1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrUpdateCustomer**](CustomersApi.md#CreateOrUpdateCustomer) | **Post** /customers/ | Create or Update a customer
[**FindCustomers**](CustomersApi.md#FindCustomers) | **Get** /customers/ | Retrieve all customers
[**FindCustomersById**](CustomersApi.md#FindCustomersById) | **Get** /customers/{id} | Retrieve a single customer by id


# **CreateOrUpdateCustomer**
> Customer CreateOrUpdateCustomer($customer)

Create or Update a customer

Create or Update a customer


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customer** | [**Customer**](Customer.md)| Add or update customer details | 

### Return type

[**Customer**](Customer.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindCustomers**
> Customers FindCustomers($page)

Retrieve all customers

Retrieve all customers


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32**| The page number ro return | [optional] 

### Return type

[**Customers**](Customers.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindCustomersById**
> Customer FindCustomersById($id)

Retrieve a single customer by id

Retrieve a single customer by id


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the customer to get | 

### Return type

[**Customer**](Customer.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

