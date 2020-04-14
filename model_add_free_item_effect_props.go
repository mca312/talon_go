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

// AddFreeItemEffectProps The properties specific to the \"addFreeItem\" effect. This gets triggered whenever a validated rule contained an \"add free item\" effect.
type AddFreeItemEffectProps struct {
	// SKU of the item that needs to be added
	Sku string `json:"sku"`
	// The name/description of the effect
	Name string `json:"name"`
}

// GetSku returns the Sku field value
func (o *AddFreeItemEffectProps) GetSku() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sku
}

// SetSku sets field value
func (o *AddFreeItemEffectProps) SetSku(v string) {
	o.Sku = v
}

// GetName returns the Name field value
func (o *AddFreeItemEffectProps) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// SetName sets field value
func (o *AddFreeItemEffectProps) SetName(v string) {
	o.Name = v
}

type NullableAddFreeItemEffectProps struct {
	Value        AddFreeItemEffectProps
	ExplicitNull bool
}

func (v NullableAddFreeItemEffectProps) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableAddFreeItemEffectProps) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
