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

// Product variants for a variable type product
type ProductVariants struct {

	Attributes []ProductAttributes `json:"attributes,omitempty"`

	// The date that the product variant was created
	Created time.Time `json:"created,omitempty"`

	// The URL to the variants image
	ImageSourceUrl string `json:"image_source_url,omitempty"`

	// If the variant is currently in stock
	InStock bool `json:"in_stock,omitempty"`

	// If inventory management is used for the variant
	InventoryManagement string `json:"inventory_management,omitempty"`

	// Inventory quantity for the variant
	InventoryQuantity string `json:"inventory_quantity,omitempty"`

	// The price of the variant
	Price float32 `json:"price,omitempty"`

	// The price compare for sale items
	PriceCompare float32 `json:"price_compare,omitempty"`

	// The Products's external reference Id
	ProductRefId string `json:"product_ref_id,omitempty"`

	// he variant's external reference Id
	RefId string `json:"ref_id,omitempty"`

	// Variant's sku
	Sku string `json:"sku,omitempty"`

	// If the variant is taxable
	Taxable bool `json:"taxable,omitempty"`

	// The title of tha variant
	Title string `json:"title,omitempty"`

	// The date that the product variant was updated
	Updated time.Time `json:"updated,omitempty"`

	// If the variant is visible
	Visible bool `json:"visible,omitempty"`

	// The weight of the product Variant
	Weight float32 `json:"weight,omitempty"`
}
