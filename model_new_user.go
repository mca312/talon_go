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

// NewUser
type NewUser struct {
	// The email address associated with your account.
	Email string `json:"email"`
	// The password for your account.
	Password string `json:"password"`
	// Your name.
	Name        *string `json:"name,omitempty"`
	InviteToken string  `json:"inviteToken"`
}

// GetEmail returns the Email field value
func (o *NewUser) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// SetEmail sets field value
func (o *NewUser) SetEmail(v string) {
	o.Email = v
}

// GetPassword returns the Password field value
func (o *NewUser) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// SetPassword sets field value
func (o *NewUser) SetPassword(v string) {
	o.Password = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NewUser) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewUser) GetNameOk() (string, bool) {
	if o == nil || o.Name == nil {
		var ret string
		return ret, false
	}
	return *o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NewUser) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NewUser) SetName(v string) {
	o.Name = &v
}

// GetInviteToken returns the InviteToken field value
func (o *NewUser) GetInviteToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InviteToken
}

// SetInviteToken sets field value
func (o *NewUser) SetInviteToken(v string) {
	o.InviteToken = v
}

type NullableNewUser struct {
	Value        NewUser
	ExplicitNull bool
}

func (v NullableNewUser) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableNewUser) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
