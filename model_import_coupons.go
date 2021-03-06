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

// ImportCoupons struct for ImportCoupons
type ImportCoupons struct {
	Coupons string `json:"coupons"`
}

// GetCoupons returns the Coupons field value
func (o *ImportCoupons) GetCoupons() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Coupons
}

// SetCoupons sets field value
func (o *ImportCoupons) SetCoupons(v string) {
	o.Coupons = v
}

type NullableImportCoupons struct {
	Value        ImportCoupons
	ExplicitNull bool
}

func (v NullableImportCoupons) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableImportCoupons) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
