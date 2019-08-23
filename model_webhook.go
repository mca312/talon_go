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

// 
type Webhook struct {
	// Unique ID for this entity.
	Id int32 `json:"id"`
	// The exact moment this entity was created.
	Created time.Time `json:"created"`
	// The exact moment this entity was last modified.
	Modified time.Time `json:"modified"`
	// The IDs of the applications that are related to this entity. The IDs of the applications that are related to this entity.
	ApplicationIds []int32 `json:"applicationIds"`
	// Friendly title for this webhook
	Title string `json:"title"`
	// API method for this webhook
	Verb string `json:"verb"`
	// API url (supports templating using parameters) for this webhook
	Url string `json:"url"`
	// List of API HTTP headers for this webhook
	Headers []string `json:"headers"`
	// API payload (supports templating using parameters) for this webhook
	Payload string `json:"payload,omitempty"`
	// Array of template argument definitions
	Params []TemplateArgDef `json:"params"`
	// Enables or disables webhook from showing in rule builder
	Enabled bool `json:"enabled"`
	// array of rulesets where webhook is used
	UsedAt []string `json:"usedAt"`
}