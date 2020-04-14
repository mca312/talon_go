/*
 * Talon.One API
 *
 * The Talon.One API is used to manage applications and campaigns, as well as to integrate with your application. The operations in the _Integration API_ section are used to integrate with our platform, while the other operations are used to manage applications and campaigns.  ### Where is the API?  The API is available at the same hostname as these docs. For example, if you are reading this page at `https://mycompany.talon.one/docs/api/`, the URL for the [updateCustomerProfile][] operation is `https://mycompany.talon.one/v1/customer_profiles/id`  [updateCustomerProfile]: #operation--v1-customer_profiles--integrationId--put
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package talon

import (
	"bytes"
	"encoding/json"
)

// AttributesMandatory Arbitrary settings associated with attributes.
type AttributesMandatory struct {
	// List of mandatory attributes for campaigns.
	Campaigns *[]string `json:"campaigns,omitempty"`
	// List of mandatory attributes for campaigns.
	Coupons *[]string `json:"coupons,omitempty"`
}

// GetCampaigns returns the Campaigns field value if set, zero value otherwise.
func (o *AttributesMandatory) GetCampaigns() []string {
	if o == nil || o.Campaigns == nil {
		var ret []string
		return ret
	}
	return *o.Campaigns
}

// GetCampaignsOk returns a tuple with the Campaigns field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *AttributesMandatory) GetCampaignsOk() ([]string, bool) {
	if o == nil || o.Campaigns == nil {
		var ret []string
		return ret, false
	}
	return *o.Campaigns, true
}

// HasCampaigns returns a boolean if a field has been set.
func (o *AttributesMandatory) HasCampaigns() bool {
	if o != nil && o.Campaigns != nil {
		return true
	}

	return false
}

// SetCampaigns gets a reference to the given []string and assigns it to the Campaigns field.
func (o *AttributesMandatory) SetCampaigns(v []string) {
	o.Campaigns = &v
}

// GetCoupons returns the Coupons field value if set, zero value otherwise.
func (o *AttributesMandatory) GetCoupons() []string {
	if o == nil || o.Coupons == nil {
		var ret []string
		return ret
	}
	return *o.Coupons
}

// GetCouponsOk returns a tuple with the Coupons field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *AttributesMandatory) GetCouponsOk() ([]string, bool) {
	if o == nil || o.Coupons == nil {
		var ret []string
		return ret, false
	}
	return *o.Coupons, true
}

// HasCoupons returns a boolean if a field has been set.
func (o *AttributesMandatory) HasCoupons() bool {
	if o != nil && o.Coupons != nil {
		return true
	}

	return false
}

// SetCoupons gets a reference to the given []string and assigns it to the Coupons field.
func (o *AttributesMandatory) SetCoupons(v []string) {
	o.Coupons = &v
}

type NullableAttributesMandatory struct {
	Value        AttributesMandatory
	ExplicitNull bool
}

func (v NullableAttributesMandatory) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableAttributesMandatory) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}