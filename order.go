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

import (
	"time"
)

type Order struct {

	// Primary key
	Id string `json:"id,omitempty"`

	// Contact reference id
	ContactRefId string `json:"contact_ref_id,omitempty"`

	// Financial status of the order
	FinancialStatus string `json:"financial_status,omitempty"`

	// iqx internal order id
	IqxOrder string `json:"iqx_order,omitempty"`

	// The individual line items of the order
	LineItems []OrderLineItems `json:"line_items,omitempty"`

	// The Order was processed at
	ProcessedAt time.Time `json:"processed_at,omitempty"`

	// The Order reference number
	RefId string `json:"ref_id,omitempty"`

	// The Order's sub total
	SubtotalPrice float32 `json:"subtotal_price,omitempty"`

	// The Order's total
	TotalPrice float32 `json:"total_price,omitempty"`

	// The Order's shipping cost
	TotalShipping float32 `json:"total_shipping,omitempty"`

	// The Order's tax amount
	TotalTax float32 `json:"total_tax,omitempty"`

	// Connected user id
	UserId string `json:"user_id,omitempty"`
}
