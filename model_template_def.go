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

// TemplateDef
type TemplateDef struct {
	// Unique ID for this entity.
	Id int32 `json:"id"`
	// The exact moment this entity was created.
	Created time.Time `json:"created"`
	// The ID of the application that owns this entity.
	ApplicationId int32 `json:"applicationId"`
	// Campaigner-friendly name for the template that will be shown in the rule editor.
	Title string `json:"title"`
	// A short description of the template that will be shown in the rule editor.
	Description string `json:"description"`
	// Extended help text for the template.
	Help string `json:"help"`
	// Used for grouping templates in the rule editor sidebar.
	Category string `json:"category"`
	// A Talang expression that contains variable bindings referring to args.
	Expr []map[string]interface{} `json:"expr"`
	// An array of argument definitions.
	Args []TemplateArgDef `json:"args"`
	// A flag to control exposure in Rule Builder.
	Expose *bool `json:"expose,omitempty"`
	// The template name used in Talang.
	Name string `json:"name"`
}

// GetId returns the Id field value
func (o *TemplateDef) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// SetId sets field value
func (o *TemplateDef) SetId(v int32) {
	o.Id = v
}

// GetCreated returns the Created field value
func (o *TemplateDef) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// SetCreated sets field value
func (o *TemplateDef) SetCreated(v time.Time) {
	o.Created = v
}

// GetApplicationId returns the ApplicationId field value
func (o *TemplateDef) GetApplicationId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ApplicationId
}

// SetApplicationId sets field value
func (o *TemplateDef) SetApplicationId(v int32) {
	o.ApplicationId = v
}

// GetTitle returns the Title field value
func (o *TemplateDef) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// SetTitle sets field value
func (o *TemplateDef) SetTitle(v string) {
	o.Title = v
}

// GetDescription returns the Description field value
func (o *TemplateDef) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// SetDescription sets field value
func (o *TemplateDef) SetDescription(v string) {
	o.Description = v
}

// GetHelp returns the Help field value
func (o *TemplateDef) GetHelp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Help
}

// SetHelp sets field value
func (o *TemplateDef) SetHelp(v string) {
	o.Help = v
}

// GetCategory returns the Category field value
func (o *TemplateDef) GetCategory() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Category
}

// SetCategory sets field value
func (o *TemplateDef) SetCategory(v string) {
	o.Category = v
}

// GetExpr returns the Expr field value
func (o *TemplateDef) GetExpr() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}

	return o.Expr
}

// SetExpr sets field value
func (o *TemplateDef) SetExpr(v []map[string]interface{}) {
	o.Expr = v
}

// GetArgs returns the Args field value
func (o *TemplateDef) GetArgs() []TemplateArgDef {
	if o == nil {
		var ret []TemplateArgDef
		return ret
	}

	return o.Args
}

// SetArgs sets field value
func (o *TemplateDef) SetArgs(v []TemplateArgDef) {
	o.Args = v
}

// GetExpose returns the Expose field value if set, zero value otherwise.
func (o *TemplateDef) GetExpose() bool {
	if o == nil || o.Expose == nil {
		var ret bool
		return ret
	}
	return *o.Expose
}

// GetExposeOk returns a tuple with the Expose field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TemplateDef) GetExposeOk() (bool, bool) {
	if o == nil || o.Expose == nil {
		var ret bool
		return ret, false
	}
	return *o.Expose, true
}

// HasExpose returns a boolean if a field has been set.
func (o *TemplateDef) HasExpose() bool {
	if o != nil && o.Expose != nil {
		return true
	}

	return false
}

// SetExpose gets a reference to the given bool and assigns it to the Expose field.
func (o *TemplateDef) SetExpose(v bool) {
	o.Expose = &v
}

// GetName returns the Name field value
func (o *TemplateDef) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// SetName sets field value
func (o *TemplateDef) SetName(v string) {
	o.Name = v
}

type NullableTemplateDef struct {
	Value        TemplateDef
	ExplicitNull bool
}

func (v NullableTemplateDef) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableTemplateDef) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
