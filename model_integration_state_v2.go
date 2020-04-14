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

// IntegrationStateV2 Contains all entities that might interest Talon.One integrations. This is the response type returned by the V2 PUT customer_session endpoint
type IntegrationStateV2 struct {
	CustomerSession    *CustomerSessionV2 `json:"customerSession,omitempty"`
	CustomerProfile    *CustomerProfile   `json:"customerProfile,omitempty"`
	Event              *Event             `json:"event,omitempty"`
	Loyalty            *Loyalty           `json:"loyalty,omitempty"`
	Referral           *Referral          `json:"referral,omitempty"`
	Coupons            *[]Coupon          `json:"coupons,omitempty"`
	TriggeredCampaigns *[]Campaign        `json:"triggeredCampaigns,omitempty"`
	Effects            []Effect           `json:"effects"`
	CreatedCoupons     []Coupon           `json:"createdCoupons"`
	CreatedReferrals   []Referral         `json:"createdReferrals"`
}

// GetCustomerSession returns the CustomerSession field value if set, zero value otherwise.
func (o *IntegrationStateV2) GetCustomerSession() CustomerSessionV2 {
	if o == nil || o.CustomerSession == nil {
		var ret CustomerSessionV2
		return ret
	}
	return *o.CustomerSession
}

// GetCustomerSessionOk returns a tuple with the CustomerSession field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationStateV2) GetCustomerSessionOk() (CustomerSessionV2, bool) {
	if o == nil || o.CustomerSession == nil {
		var ret CustomerSessionV2
		return ret, false
	}
	return *o.CustomerSession, true
}

// HasCustomerSession returns a boolean if a field has been set.
func (o *IntegrationStateV2) HasCustomerSession() bool {
	if o != nil && o.CustomerSession != nil {
		return true
	}

	return false
}

// SetCustomerSession gets a reference to the given CustomerSessionV2 and assigns it to the CustomerSession field.
func (o *IntegrationStateV2) SetCustomerSession(v CustomerSessionV2) {
	o.CustomerSession = &v
}

// GetCustomerProfile returns the CustomerProfile field value if set, zero value otherwise.
func (o *IntegrationStateV2) GetCustomerProfile() CustomerProfile {
	if o == nil || o.CustomerProfile == nil {
		var ret CustomerProfile
		return ret
	}
	return *o.CustomerProfile
}

// GetCustomerProfileOk returns a tuple with the CustomerProfile field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationStateV2) GetCustomerProfileOk() (CustomerProfile, bool) {
	if o == nil || o.CustomerProfile == nil {
		var ret CustomerProfile
		return ret, false
	}
	return *o.CustomerProfile, true
}

// HasCustomerProfile returns a boolean if a field has been set.
func (o *IntegrationStateV2) HasCustomerProfile() bool {
	if o != nil && o.CustomerProfile != nil {
		return true
	}

	return false
}

// SetCustomerProfile gets a reference to the given CustomerProfile and assigns it to the CustomerProfile field.
func (o *IntegrationStateV2) SetCustomerProfile(v CustomerProfile) {
	o.CustomerProfile = &v
}

// GetEvent returns the Event field value if set, zero value otherwise.
func (o *IntegrationStateV2) GetEvent() Event {
	if o == nil || o.Event == nil {
		var ret Event
		return ret
	}
	return *o.Event
}

// GetEventOk returns a tuple with the Event field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationStateV2) GetEventOk() (Event, bool) {
	if o == nil || o.Event == nil {
		var ret Event
		return ret, false
	}
	return *o.Event, true
}

// HasEvent returns a boolean if a field has been set.
func (o *IntegrationStateV2) HasEvent() bool {
	if o != nil && o.Event != nil {
		return true
	}

	return false
}

// SetEvent gets a reference to the given Event and assigns it to the Event field.
func (o *IntegrationStateV2) SetEvent(v Event) {
	o.Event = &v
}

// GetLoyalty returns the Loyalty field value if set, zero value otherwise.
func (o *IntegrationStateV2) GetLoyalty() Loyalty {
	if o == nil || o.Loyalty == nil {
		var ret Loyalty
		return ret
	}
	return *o.Loyalty
}

// GetLoyaltyOk returns a tuple with the Loyalty field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationStateV2) GetLoyaltyOk() (Loyalty, bool) {
	if o == nil || o.Loyalty == nil {
		var ret Loyalty
		return ret, false
	}
	return *o.Loyalty, true
}

// HasLoyalty returns a boolean if a field has been set.
func (o *IntegrationStateV2) HasLoyalty() bool {
	if o != nil && o.Loyalty != nil {
		return true
	}

	return false
}

// SetLoyalty gets a reference to the given Loyalty and assigns it to the Loyalty field.
func (o *IntegrationStateV2) SetLoyalty(v Loyalty) {
	o.Loyalty = &v
}

// GetReferral returns the Referral field value if set, zero value otherwise.
func (o *IntegrationStateV2) GetReferral() Referral {
	if o == nil || o.Referral == nil {
		var ret Referral
		return ret
	}
	return *o.Referral
}

// GetReferralOk returns a tuple with the Referral field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationStateV2) GetReferralOk() (Referral, bool) {
	if o == nil || o.Referral == nil {
		var ret Referral
		return ret, false
	}
	return *o.Referral, true
}

// HasReferral returns a boolean if a field has been set.
func (o *IntegrationStateV2) HasReferral() bool {
	if o != nil && o.Referral != nil {
		return true
	}

	return false
}

// SetReferral gets a reference to the given Referral and assigns it to the Referral field.
func (o *IntegrationStateV2) SetReferral(v Referral) {
	o.Referral = &v
}

// GetCoupons returns the Coupons field value if set, zero value otherwise.
func (o *IntegrationStateV2) GetCoupons() []Coupon {
	if o == nil || o.Coupons == nil {
		var ret []Coupon
		return ret
	}
	return *o.Coupons
}

// GetCouponsOk returns a tuple with the Coupons field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationStateV2) GetCouponsOk() ([]Coupon, bool) {
	if o == nil || o.Coupons == nil {
		var ret []Coupon
		return ret, false
	}
	return *o.Coupons, true
}

// HasCoupons returns a boolean if a field has been set.
func (o *IntegrationStateV2) HasCoupons() bool {
	if o != nil && o.Coupons != nil {
		return true
	}

	return false
}

// SetCoupons gets a reference to the given []Coupon and assigns it to the Coupons field.
func (o *IntegrationStateV2) SetCoupons(v []Coupon) {
	o.Coupons = &v
}

// GetTriggeredCampaigns returns the TriggeredCampaigns field value if set, zero value otherwise.
func (o *IntegrationStateV2) GetTriggeredCampaigns() []Campaign {
	if o == nil || o.TriggeredCampaigns == nil {
		var ret []Campaign
		return ret
	}
	return *o.TriggeredCampaigns
}

// GetTriggeredCampaignsOk returns a tuple with the TriggeredCampaigns field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationStateV2) GetTriggeredCampaignsOk() ([]Campaign, bool) {
	if o == nil || o.TriggeredCampaigns == nil {
		var ret []Campaign
		return ret, false
	}
	return *o.TriggeredCampaigns, true
}

// HasTriggeredCampaigns returns a boolean if a field has been set.
func (o *IntegrationStateV2) HasTriggeredCampaigns() bool {
	if o != nil && o.TriggeredCampaigns != nil {
		return true
	}

	return false
}

// SetTriggeredCampaigns gets a reference to the given []Campaign and assigns it to the TriggeredCampaigns field.
func (o *IntegrationStateV2) SetTriggeredCampaigns(v []Campaign) {
	o.TriggeredCampaigns = &v
}

// GetEffects returns the Effects field value
func (o *IntegrationStateV2) GetEffects() []Effect {
	if o == nil {
		var ret []Effect
		return ret
	}

	return o.Effects
}

// SetEffects sets field value
func (o *IntegrationStateV2) SetEffects(v []Effect) {
	o.Effects = v
}

// GetCreatedCoupons returns the CreatedCoupons field value
func (o *IntegrationStateV2) GetCreatedCoupons() []Coupon {
	if o == nil {
		var ret []Coupon
		return ret
	}

	return o.CreatedCoupons
}

// SetCreatedCoupons sets field value
func (o *IntegrationStateV2) SetCreatedCoupons(v []Coupon) {
	o.CreatedCoupons = v
}

// GetCreatedReferrals returns the CreatedReferrals field value
func (o *IntegrationStateV2) GetCreatedReferrals() []Referral {
	if o == nil {
		var ret []Referral
		return ret
	}

	return o.CreatedReferrals
}

// SetCreatedReferrals sets field value
func (o *IntegrationStateV2) SetCreatedReferrals(v []Referral) {
	o.CreatedReferrals = v
}

type NullableIntegrationStateV2 struct {
	Value        IntegrationStateV2
	ExplicitNull bool
}

func (v NullableIntegrationStateV2) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableIntegrationStateV2) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}