/* 
 * iQX Amplify API Reference
 *
 * Welcome to the iQX Amplify API reference. This is a live example of how you can use [iQX Amplify](http://app.iqxamplify.com/) to integrate your own e-commerce store or ticketing engine.  The iQX Amplify API is organized around the [REST](http://en.wikipedia.org/wiki/Representational_State_Transfer) mothodology, and it uses resource-oriented URLs, and common HTTP response codes to indicate API errors. All requests are authenticated using an `api-key` which can be obtained from your developer dashboard within iQX Amplify.
 *
 * OpenAPI spec version: 1.0.0
 * Contact: development@iqxcorp.com
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

// Product variants for a variable type product
type OrderLineItems struct {

	// The qunatity of the line item
	FulfillableQuantity float32 `json:"fulfillable_quantity,omitempty"`

	// The price of the line item
	Price float32 `json:"price,omitempty"`

	// The Product's reference id
	ProductRefId string `json:"product_ref_id,omitempty"`

	// The order's reference id
	RefId string `json:"ref_id,omitempty"`

	// The order price for theline item
	OrderPrice float32 `json:"order_price,omitempty"`

	// Product's sku
	Sku string `json:"sku,omitempty"`

	// Product's name
	Name string `json:"name,omitempty"`

	// Product's title
	Title string `json:"title,omitempty"`

	// Number of items
	Quantity float32 `json:"quantity,omitempty"`

	// Weight in grams
	Grams float32 `json:"grams,omitempty"`

	// If the order requires shipping
	RequiresShipping bool `json:"requires_shipping,omitempty"`
}