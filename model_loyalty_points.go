/*
 * Talon.One API
 *
 * The Talon.One API is used to manage applications and campaigns, as well as to integrate with your application. The operations in the _Integration API_ section are used to integrate with our platform, while the other operations are used to manage applications and campaigns.  ### Where is the API?  The API is available at the same hostname as these docs. For example, if you are reading this page at `https://mycompany.talon.one/docs/api/`, the URL for the [updateCustomerProfile][] operation is `https://mycompany.talon.one/v1/customer_profiles/id`  [updateCustomerProfile]: #operation--v1-customer_profiles--integrationId--put 
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package talon

// Points to add or deduct
type LoyaltyPoints struct {
	// Amount of loyalty points
	Points float32 `json:"points"`
	// Allows to specify a name for the addition or deduction
	Name string `json:"name,omitempty"`
	// Indicates the duration after which the added loyalty points should expire. The format is a number followed by one letter indicating the unit, like '1h' or '40m' or '30d'.
	ExpiryDuration string `json:"expiryDuration,omitempty"`
	// This specifies if we are adding loyalty points to the main ledger or a subledger
	SubLedgerID string `json:"subLedgerID,omitempty"`
}
