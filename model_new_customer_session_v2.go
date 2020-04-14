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

// NewCustomerSessionV2
type NewCustomerSessionV2 struct {
	// ID of the customers profile as used within this Talon.One account. May be omitted or set to the empty string if the customer does not yet have a known profile ID.
	ProfileId *string `json:"profileId,omitempty"`
	// Any coupon codes entered.
	CouponCodes *[]string `json:"couponCodes,omitempty"`
	// Any referral code entered.
	ReferralCode *string `json:"referralCode,omitempty"`
	// Indicates the current state of the session. All sessions must start in the \"open\" state, after which valid transitions are...  1. open -> closed 2. open -> cancelled 3. closed -> cancelled
	State *string `json:"state,omitempty"`
	// All items the customer will be purchasing in this session
	CartItems *[]CartItem `json:"cartItems,omitempty"`
	// Any costs associated with the session that can not be explicitly attributed to cart items. Examples include shipping costs and service fees.
	AdditionalCosts *map[string]AdditionalCost `json:"additionalCosts,omitempty"`
	// Identifiers for the customer, this can be used for limits on values such as device ID.
	Identifiers *[]string `json:"identifiers,omitempty"`
	// A key-value map of the sessions attributes. The potentially valid attributes are configured in your accounts developer settings.
	Attributes *map[string]interface{} `json:"attributes,omitempty"`
}

// GetProfileId returns the ProfileId field value if set, zero value otherwise.
func (o *NewCustomerSessionV2) GetProfileId() string {
	if o == nil || o.ProfileId == nil {
		var ret string
		return ret
	}
	return *o.ProfileId
}

// GetProfileIdOk returns a tuple with the ProfileId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewCustomerSessionV2) GetProfileIdOk() (string, bool) {
	if o == nil || o.ProfileId == nil {
		var ret string
		return ret, false
	}
	return *o.ProfileId, true
}

// HasProfileId returns a boolean if a field has been set.
func (o *NewCustomerSessionV2) HasProfileId() bool {
	if o != nil && o.ProfileId != nil {
		return true
	}

	return false
}

// SetProfileId gets a reference to the given string and assigns it to the ProfileId field.
func (o *NewCustomerSessionV2) SetProfileId(v string) {
	o.ProfileId = &v
}

// GetCouponCodes returns the CouponCodes field value if set, zero value otherwise.
func (o *NewCustomerSessionV2) GetCouponCodes() []string {
	if o == nil || o.CouponCodes == nil {
		var ret []string
		return ret
	}
	return *o.CouponCodes
}

// GetCouponCodesOk returns a tuple with the CouponCodes field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewCustomerSessionV2) GetCouponCodesOk() ([]string, bool) {
	if o == nil || o.CouponCodes == nil {
		var ret []string
		return ret, false
	}
	return *o.CouponCodes, true
}

// HasCouponCodes returns a boolean if a field has been set.
func (o *NewCustomerSessionV2) HasCouponCodes() bool {
	if o != nil && o.CouponCodes != nil {
		return true
	}

	return false
}

// SetCouponCodes gets a reference to the given []string and assigns it to the CouponCodes field.
func (o *NewCustomerSessionV2) SetCouponCodes(v []string) {
	o.CouponCodes = &v
}

// GetReferralCode returns the ReferralCode field value if set, zero value otherwise.
func (o *NewCustomerSessionV2) GetReferralCode() string {
	if o == nil || o.ReferralCode == nil {
		var ret string
		return ret
	}
	return *o.ReferralCode
}

// GetReferralCodeOk returns a tuple with the ReferralCode field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewCustomerSessionV2) GetReferralCodeOk() (string, bool) {
	if o == nil || o.ReferralCode == nil {
		var ret string
		return ret, false
	}
	return *o.ReferralCode, true
}

// HasReferralCode returns a boolean if a field has been set.
func (o *NewCustomerSessionV2) HasReferralCode() bool {
	if o != nil && o.ReferralCode != nil {
		return true
	}

	return false
}

// SetReferralCode gets a reference to the given string and assigns it to the ReferralCode field.
func (o *NewCustomerSessionV2) SetReferralCode(v string) {
	o.ReferralCode = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *NewCustomerSessionV2) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewCustomerSessionV2) GetStateOk() (string, bool) {
	if o == nil || o.State == nil {
		var ret string
		return ret, false
	}
	return *o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *NewCustomerSessionV2) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *NewCustomerSessionV2) SetState(v string) {
	o.State = &v
}

// GetCartItems returns the CartItems field value if set, zero value otherwise.
func (o *NewCustomerSessionV2) GetCartItems() []CartItem {
	if o == nil || o.CartItems == nil {
		var ret []CartItem
		return ret
	}
	return *o.CartItems
}

// GetCartItemsOk returns a tuple with the CartItems field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewCustomerSessionV2) GetCartItemsOk() ([]CartItem, bool) {
	if o == nil || o.CartItems == nil {
		var ret []CartItem
		return ret, false
	}
	return *o.CartItems, true
}

// HasCartItems returns a boolean if a field has been set.
func (o *NewCustomerSessionV2) HasCartItems() bool {
	if o != nil && o.CartItems != nil {
		return true
	}

	return false
}

// SetCartItems gets a reference to the given []CartItem and assigns it to the CartItems field.
func (o *NewCustomerSessionV2) SetCartItems(v []CartItem) {
	o.CartItems = &v
}

// GetAdditionalCosts returns the AdditionalCosts field value if set, zero value otherwise.
func (o *NewCustomerSessionV2) GetAdditionalCosts() map[string]AdditionalCost {
	if o == nil || o.AdditionalCosts == nil {
		var ret map[string]AdditionalCost
		return ret
	}
	return *o.AdditionalCosts
}

// GetAdditionalCostsOk returns a tuple with the AdditionalCosts field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewCustomerSessionV2) GetAdditionalCostsOk() (map[string]AdditionalCost, bool) {
	if o == nil || o.AdditionalCosts == nil {
		var ret map[string]AdditionalCost
		return ret, false
	}
	return *o.AdditionalCosts, true
}

// HasAdditionalCosts returns a boolean if a field has been set.
func (o *NewCustomerSessionV2) HasAdditionalCosts() bool {
	if o != nil && o.AdditionalCosts != nil {
		return true
	}

	return false
}

// SetAdditionalCosts gets a reference to the given map[string]AdditionalCost and assigns it to the AdditionalCosts field.
func (o *NewCustomerSessionV2) SetAdditionalCosts(v map[string]AdditionalCost) {
	o.AdditionalCosts = &v
}

// GetIdentifiers returns the Identifiers field value if set, zero value otherwise.
func (o *NewCustomerSessionV2) GetIdentifiers() []string {
	if o == nil || o.Identifiers == nil {
		var ret []string
		return ret
	}
	return *o.Identifiers
}

// GetIdentifiersOk returns a tuple with the Identifiers field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewCustomerSessionV2) GetIdentifiersOk() ([]string, bool) {
	if o == nil || o.Identifiers == nil {
		var ret []string
		return ret, false
	}
	return *o.Identifiers, true
}

// HasIdentifiers returns a boolean if a field has been set.
func (o *NewCustomerSessionV2) HasIdentifiers() bool {
	if o != nil && o.Identifiers != nil {
		return true
	}

	return false
}

// SetIdentifiers gets a reference to the given []string and assigns it to the Identifiers field.
func (o *NewCustomerSessionV2) SetIdentifiers(v []string) {
	o.Identifiers = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *NewCustomerSessionV2) GetAttributes() map[string]interface{} {
	if o == nil || o.Attributes == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewCustomerSessionV2) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || o.Attributes == nil {
		var ret map[string]interface{}
		return ret, false
	}
	return *o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *NewCustomerSessionV2) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *NewCustomerSessionV2) SetAttributes(v map[string]interface{}) {
	o.Attributes = &v
}

type NullableNewCustomerSessionV2 struct {
	Value        NewCustomerSessionV2
	ExplicitNull bool
}

func (v NullableNewCustomerSessionV2) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableNewCustomerSessionV2) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}