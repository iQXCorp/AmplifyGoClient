# Order

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Primary key | [optional] [default to null]
**ContactRefId** | **string** | Contact reference id | [optional] [default to null]
**FinancialStatus** | **string** | Financial status of the order | [optional] [default to null]
**IqxOrder** | **string** | iqx internal order id | [optional] [default to null]
**LineItems** | [**[]OrderLineItems**](Order_line_items.md) | The individual line items of the order | [optional] [default to null]
**ProcessedAt** | [**time.Time**](time.Time.md) | The Order was processed at | [optional] [default to null]
**RefId** | **string** | The Order reference number | [optional] [default to null]
**SubtotalPrice** | **float32** | The Order&#39;s sub total | [optional] [default to null]
**TotalPrice** | **float32** | The Order&#39;s total | [optional] [default to null]
**TotalShipping** | **float32** | The Order&#39;s shipping cost | [optional] [default to null]
**TotalTax** | **float32** | The Order&#39;s tax amount | [optional] [default to null]
**UserId** | **string** | Connected user id | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


