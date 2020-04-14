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

// MiscUpdateUserLatestFeature struct for MiscUpdateUserLatestFeature
type MiscUpdateUserLatestFeature struct {
	LatestFeature *string `json:"latestFeature,omitempty"`
}

// GetLatestFeature returns the LatestFeature field value if set, zero value otherwise.
func (o *MiscUpdateUserLatestFeature) GetLatestFeature() string {
	if o == nil || o.LatestFeature == nil {
		var ret string
		return ret
	}
	return *o.LatestFeature
}

// GetLatestFeatureOk returns a tuple with the LatestFeature field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MiscUpdateUserLatestFeature) GetLatestFeatureOk() (string, bool) {
	if o == nil || o.LatestFeature == nil {
		var ret string
		return ret, false
	}
	return *o.LatestFeature, true
}

// HasLatestFeature returns a boolean if a field has been set.
func (o *MiscUpdateUserLatestFeature) HasLatestFeature() bool {
	if o != nil && o.LatestFeature != nil {
		return true
	}

	return false
}

// SetLatestFeature gets a reference to the given string and assigns it to the LatestFeature field.
func (o *MiscUpdateUserLatestFeature) SetLatestFeature(v string) {
	o.LatestFeature = &v
}

type NullableMiscUpdateUserLatestFeature struct {
	Value        MiscUpdateUserLatestFeature
	ExplicitNull bool
}

func (v NullableMiscUpdateUserLatestFeature) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableMiscUpdateUserLatestFeature) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
