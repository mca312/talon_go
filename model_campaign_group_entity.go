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

// CampaignGroupEntity struct for CampaignGroupEntity
type CampaignGroupEntity struct {
	// The IDs of the campaign groups that own this entity.
	CampaignGroups *[]int32 `json:"campaignGroups,omitempty"`
}

// GetCampaignGroups returns the CampaignGroups field value if set, zero value otherwise.
func (o *CampaignGroupEntity) GetCampaignGroups() []int32 {
	if o == nil || o.CampaignGroups == nil {
		var ret []int32
		return ret
	}
	return *o.CampaignGroups
}

// GetCampaignGroupsOk returns a tuple with the CampaignGroups field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CampaignGroupEntity) GetCampaignGroupsOk() ([]int32, bool) {
	if o == nil || o.CampaignGroups == nil {
		var ret []int32
		return ret, false
	}
	return *o.CampaignGroups, true
}

// HasCampaignGroups returns a boolean if a field has been set.
func (o *CampaignGroupEntity) HasCampaignGroups() bool {
	if o != nil && o.CampaignGroups != nil {
		return true
	}

	return false
}

// SetCampaignGroups gets a reference to the given []int32 and assigns it to the CampaignGroups field.
func (o *CampaignGroupEntity) SetCampaignGroups(v []int32) {
	o.CampaignGroups = &v
}

type NullableCampaignGroupEntity struct {
	Value        CampaignGroupEntity
	ExplicitNull bool
}

func (v NullableCampaignGroupEntity) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableCampaignGroupEntity) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
