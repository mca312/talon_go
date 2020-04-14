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

// ReferralRejectionReason Holds a reference to the campaign, the referral and the reason for which that referral was rejected. Should only be present when there is a 'rejectReferral' effect.
type ReferralRejectionReason struct {
	CampaignId int32  `json:"campaignId"`
	ReferralId int32  `json:"referralId"`
	Reason     string `json:"reason"`
}

// GetCampaignId returns the CampaignId field value
func (o *ReferralRejectionReason) GetCampaignId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CampaignId
}

// SetCampaignId sets field value
func (o *ReferralRejectionReason) SetCampaignId(v int32) {
	o.CampaignId = v
}

// GetReferralId returns the ReferralId field value
func (o *ReferralRejectionReason) GetReferralId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ReferralId
}

// SetReferralId sets field value
func (o *ReferralRejectionReason) SetReferralId(v int32) {
	o.ReferralId = v
}

// GetReason returns the Reason field value
func (o *ReferralRejectionReason) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// SetReason sets field value
func (o *ReferralRejectionReason) SetReason(v string) {
	o.Reason = v
}

type NullableReferralRejectionReason struct {
	Value        ReferralRejectionReason
	ExplicitNull bool
}

func (v NullableReferralRejectionReason) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableReferralRejectionReason) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
