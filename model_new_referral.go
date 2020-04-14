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
	"time"
)

// NewReferral struct for NewReferral
type NewReferral struct {
	// ID of the campaign from which the referral received the referral code.
	CampaignId int32 `json:"campaignId"`
	// The Integration Id of the Advocate's Profile
	AdvocateProfileIntegrationId string `json:"advocateProfileIntegrationId"`
	// An optional Integration ID of the Friend's Profile
	FriendProfileIntegrationId *string `json:"friendProfileIntegrationId,omitempty"`
	// Timestamp at which point the referral code becomes valid.
	StartDate *time.Time `json:"startDate,omitempty"`
	// Expiry date of the referral code. Referral never expires if this is omitted, zero, or negative.
	ExpiryDate *time.Time `json:"expiryDate,omitempty"`
}

// GetCampaignId returns the CampaignId field value
func (o *NewReferral) GetCampaignId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CampaignId
}

// SetCampaignId sets field value
func (o *NewReferral) SetCampaignId(v int32) {
	o.CampaignId = v
}

// GetAdvocateProfileIntegrationId returns the AdvocateProfileIntegrationId field value
func (o *NewReferral) GetAdvocateProfileIntegrationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdvocateProfileIntegrationId
}

// SetAdvocateProfileIntegrationId sets field value
func (o *NewReferral) SetAdvocateProfileIntegrationId(v string) {
	o.AdvocateProfileIntegrationId = v
}

// GetFriendProfileIntegrationId returns the FriendProfileIntegrationId field value if set, zero value otherwise.
func (o *NewReferral) GetFriendProfileIntegrationId() string {
	if o == nil || o.FriendProfileIntegrationId == nil {
		var ret string
		return ret
	}
	return *o.FriendProfileIntegrationId
}

// GetFriendProfileIntegrationIdOk returns a tuple with the FriendProfileIntegrationId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewReferral) GetFriendProfileIntegrationIdOk() (string, bool) {
	if o == nil || o.FriendProfileIntegrationId == nil {
		var ret string
		return ret, false
	}
	return *o.FriendProfileIntegrationId, true
}

// HasFriendProfileIntegrationId returns a boolean if a field has been set.
func (o *NewReferral) HasFriendProfileIntegrationId() bool {
	if o != nil && o.FriendProfileIntegrationId != nil {
		return true
	}

	return false
}

// SetFriendProfileIntegrationId gets a reference to the given string and assigns it to the FriendProfileIntegrationId field.
func (o *NewReferral) SetFriendProfileIntegrationId(v string) {
	o.FriendProfileIntegrationId = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *NewReferral) GetStartDate() time.Time {
	if o == nil || o.StartDate == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewReferral) GetStartDateOk() (time.Time, bool) {
	if o == nil || o.StartDate == nil {
		var ret time.Time
		return ret, false
	}
	return *o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *NewReferral) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *NewReferral) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetExpiryDate returns the ExpiryDate field value if set, zero value otherwise.
func (o *NewReferral) GetExpiryDate() time.Time {
	if o == nil || o.ExpiryDate == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiryDate
}

// GetExpiryDateOk returns a tuple with the ExpiryDate field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewReferral) GetExpiryDateOk() (time.Time, bool) {
	if o == nil || o.ExpiryDate == nil {
		var ret time.Time
		return ret, false
	}
	return *o.ExpiryDate, true
}

// HasExpiryDate returns a boolean if a field has been set.
func (o *NewReferral) HasExpiryDate() bool {
	if o != nil && o.ExpiryDate != nil {
		return true
	}

	return false
}

// SetExpiryDate gets a reference to the given time.Time and assigns it to the ExpiryDate field.
func (o *NewReferral) SetExpiryDate(v time.Time) {
	o.ExpiryDate = &v
}

type NullableNewReferral struct {
	Value        NewReferral
	ExplicitNull bool
}

func (v NullableNewReferral) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableNewReferral) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
