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

// CouponCreatedEffectProps The properties specific to the \"couponCreated\" effect. This gets triggered whenever a validated rule contained a \"create coupon\" effect, and a coupon was created for a customer. See \"createdCoupons\" on the response for all details of this coupon.
type CouponCreatedEffectProps struct {
	// The coupon code that was created
	Value string `json:"value"`
	// The integration identifier of the customer for whom this coupon was created
	ProfileId string `json:"profileId"`
}

// GetValue returns the Value field value
func (o *CouponCreatedEffectProps) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// SetValue sets field value
func (o *CouponCreatedEffectProps) SetValue(v string) {
	o.Value = v
}

// GetProfileId returns the ProfileId field value
func (o *CouponCreatedEffectProps) GetProfileId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProfileId
}

// SetProfileId sets field value
func (o *CouponCreatedEffectProps) SetProfileId(v string) {
	o.ProfileId = v
}

type NullableCouponCreatedEffectProps struct {
	Value        CouponCreatedEffectProps
	ExplicitNull bool
}

func (v NullableCouponCreatedEffectProps) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableCouponCreatedEffectProps) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
