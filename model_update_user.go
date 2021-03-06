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

// UpdateUser struct for UpdateUser
type UpdateUser struct {
	// Your name.
	Name *string `json:"name,omitempty"`
	// a blob of acl json
	Policy *string `json:"policy,omitempty"`
	// New state (\"deactivated\" or \"active\") for the user. Only usable by admins for the user.
	State *string `json:"state,omitempty"`
	// Update
	Roles                                *[]int32                `json:"roles,omitempty"`
	ApplicationNotificationSubscriptions *map[string]interface{} `json:"applicationNotificationSubscriptions,omitempty"`
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateUser) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUser) GetNameOk() (string, bool) {
	if o == nil || o.Name == nil {
		var ret string
		return ret, false
	}
	return *o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateUser) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateUser) SetName(v string) {
	o.Name = &v
}

// GetPolicy returns the Policy field value if set, zero value otherwise.
func (o *UpdateUser) GetPolicy() string {
	if o == nil || o.Policy == nil {
		var ret string
		return ret
	}
	return *o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUser) GetPolicyOk() (string, bool) {
	if o == nil || o.Policy == nil {
		var ret string
		return ret, false
	}
	return *o.Policy, true
}

// HasPolicy returns a boolean if a field has been set.
func (o *UpdateUser) HasPolicy() bool {
	if o != nil && o.Policy != nil {
		return true
	}

	return false
}

// SetPolicy gets a reference to the given string and assigns it to the Policy field.
func (o *UpdateUser) SetPolicy(v string) {
	o.Policy = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *UpdateUser) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUser) GetStateOk() (string, bool) {
	if o == nil || o.State == nil {
		var ret string
		return ret, false
	}
	return *o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *UpdateUser) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *UpdateUser) SetState(v string) {
	o.State = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *UpdateUser) GetRoles() []int32 {
	if o == nil || o.Roles == nil {
		var ret []int32
		return ret
	}
	return *o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUser) GetRolesOk() ([]int32, bool) {
	if o == nil || o.Roles == nil {
		var ret []int32
		return ret, false
	}
	return *o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *UpdateUser) HasRoles() bool {
	if o != nil && o.Roles != nil {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []int32 and assigns it to the Roles field.
func (o *UpdateUser) SetRoles(v []int32) {
	o.Roles = &v
}

// GetApplicationNotificationSubscriptions returns the ApplicationNotificationSubscriptions field value if set, zero value otherwise.
func (o *UpdateUser) GetApplicationNotificationSubscriptions() map[string]interface{} {
	if o == nil || o.ApplicationNotificationSubscriptions == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.ApplicationNotificationSubscriptions
}

// GetApplicationNotificationSubscriptionsOk returns a tuple with the ApplicationNotificationSubscriptions field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUser) GetApplicationNotificationSubscriptionsOk() (map[string]interface{}, bool) {
	if o == nil || o.ApplicationNotificationSubscriptions == nil {
		var ret map[string]interface{}
		return ret, false
	}
	return *o.ApplicationNotificationSubscriptions, true
}

// HasApplicationNotificationSubscriptions returns a boolean if a field has been set.
func (o *UpdateUser) HasApplicationNotificationSubscriptions() bool {
	if o != nil && o.ApplicationNotificationSubscriptions != nil {
		return true
	}

	return false
}

// SetApplicationNotificationSubscriptions gets a reference to the given map[string]interface{} and assigns it to the ApplicationNotificationSubscriptions field.
func (o *UpdateUser) SetApplicationNotificationSubscriptions(v map[string]interface{}) {
	o.ApplicationNotificationSubscriptions = &v
}

type NullableUpdateUser struct {
	Value        UpdateUser
	ExplicitNull bool
}

func (v NullableUpdateUser) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableUpdateUser) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
