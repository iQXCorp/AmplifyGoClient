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

type Customer struct {

	// Primary key
	Id string `json:"id,omitempty"`

	// Address Line 1
	Address1 string `json:"address1,omitempty"`

	// Address Line 2
	Address2 string `json:"address2,omitempty"`

	// The Customer's Area Code
	AreaCode string `json:"area_code,omitempty"`

	// The Customer's City
	City string `json:"city,omitempty"`

	// The Customer's Company
	Company string `json:"company,omitempty"`

	// The Customer's Country
	Country string `json:"country,omitempty"`

	// The Customer's Email Address
	Email string `json:"email,omitempty"`

	// The Customer's First Name
	FirstName string `json:"first_name,omitempty"`

	// The Customer's Last Name
	LastName string `json:"last_name,omitempty"`

	// Date customer was activated
	ActivationDate time.Time `json:"activation_date,omitempty"`

	// Date last modified
	ModifiedDate time.Time `json:"modified_date,omitempty"`

	// Total number of orders
	OrdersCount float32 `json:"orders_count,omitempty"`

	// The Customer's Phone Number
	Phone string `json:"phone,omitempty"`

	// The Customer's Province
	Province string `json:"province,omitempty"`

	// The Customer's external reference Id
	RefId string `json:"ref_id,omitempty"`

	// The date that the Customer signed up
	SignedUpAt time.Time `json:"signed_up_at,omitempty"`

	// The Customer's DST Timezone offset
	TimeZoneDstOffset float32 `json:"time_zone_dst_offset,omitempty"`

	// The customer's Timezone
	TimeZoneId string `json:"time_zone_id,omitempty"`

	// The Customer's Timezone offset
	TimeZoneRawOffset float32 `json:"time_zone_raw_offset,omitempty"`

	// Customer's total spend to date
	TotalSpent float32 `json:"total_spent,omitempty"`

	// Connected user id
	UserId string `json:"user_id,omitempty"`

	// The Customer has a valid phone number
	ValidPhoneNumber bool `json:"valid_phone_number,omitempty"`

	// The Customer has been validated and whitelisted
	Verified bool `json:"verified,omitempty"`

	// The customer's Zipcode
	Zip string `json:"zip,omitempty"`
}
