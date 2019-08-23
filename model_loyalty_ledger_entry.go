/*
 * Talon.One API
 *
 * The Talon.One API is used to manage applications and campaigns, as well as to integrate with your application. The operations in the _Integration API_ section are used to integrate with our platform, while the other operations are used to manage applications and campaigns.  ### Where is the API?  The API is available at the same hostname as these docs. For example, if you are reading this page at `https://mycompany.talon.one/docs/api/`, the URL for the [updateCustomerProfile][] operation is `https://mycompany.talon.one/v1/customer_profiles/id`  [updateCustomerProfile]: #operation--v1-customer_profiles--integrationId--put 
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package talon

import (
	"time"
)

// A single row of the ledger, describing one addition or deduction.
type LoyaltyLedgerEntry struct {
	Created time.Time `json:"created"`
	ProgramID int32 `json:"programID"`
	CustomerProfileID string `json:"customerProfileID"`
	CustomerSessionID string `json:"customerSessionID,omitempty"`
	EventID int32 `json:"eventID,omitempty"`
	Type_ string `json:"type"`
	Amount float32 `json:"amount"`
	ExpiryDate time.Time `json:"expiryDate,omitempty"`
	// A name referencing the condition or effect that added this entry, or the specific name provided in an API call.
	Name string `json:"name"`
	// This specifies if we are adding loyalty points to the main ledger or a subledger
	SubLedgerID string `json:"subLedgerID"`
}
