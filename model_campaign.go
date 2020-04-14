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

// Campaign
type Campaign struct {
	// Unique ID for this entity.
	Id int32 `json:"id"`
	// The exact moment this entity was created.
	Created time.Time `json:"created"`
	// The ID of the application that owns this entity.
	ApplicationId int32 `json:"applicationId"`
	// The ID of the account that owns this entity.
	UserId int32 `json:"userId"`
	// A friendly name for this campaign.
	Name string `json:"name"`
	// A detailed description of the campaign.
	Description string `json:"description"`
	// Datetime when the campaign will become active.
	StartTime *time.Time `json:"startTime,omitempty"`
	// Datetime when the campaign will become in-active.
	EndTime *time.Time `json:"endTime,omitempty"`
	// Arbitrary properties associated with this campaign
	Attributes *map[string]interface{} `json:"attributes,omitempty"`
	// A disabled or archived campaign is not evaluated for rules or coupons.
	State string `json:"state"`
	// ID of Ruleset this campaign applies on customer session evaluation.
	ActiveRulesetId *int32 `json:"activeRulesetId,omitempty"`
	// A list of tags for the campaign.
	Tags []string `json:"tags"`
	// A list of features for the campaign.
	Features         []string               `json:"features"`
	CouponSettings   *CodeGeneratorSettings `json:"couponSettings,omitempty"`
	ReferralSettings *CodeGeneratorSettings `json:"referralSettings,omitempty"`
	// The set of limits that will operate for this campaign
	Limits []LimitConfig `json:"limits"`
	// Number of coupons redeemed in the campaign.
	CouponRedemptionCount *int32 `json:"couponRedemptionCount,omitempty"`
	// Number of referral codes redeemed in the campaign.
	ReferralRedemptionCount *int32 `json:"referralRedemptionCount,omitempty"`
	// Total amount of discounts redeemed in the campaign.
	DiscountCount *int32 `json:"discountCount,omitempty"`
	// Total number of times discounts were redeemed in this campaign.
	DiscountEffectCount *int32 `json:"discountEffectCount,omitempty"`
	// Total number of coupons created by rules in this campaign.
	CouponCreationCount *int32 `json:"couponCreationCount,omitempty"`
	// Timestamp of the most recent event received by this campaign.
	LastActivity *time.Time `json:"lastActivity,omitempty"`
	// Timestamp of the most recent update to the campaign or any of its elements.
	Updated *time.Time `json:"updated,omitempty"`
	// Name of the user who created this campaign if available.
	CreatedBy *string `json:"createdBy,omitempty"`
	// Name of the user who last updated this campaign if available.
	UpdatedBy *string `json:"updatedBy,omitempty"`
}

// GetId returns the Id field value
func (o *Campaign) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// SetId sets field value
func (o *Campaign) SetId(v int32) {
	o.Id = v
}

// GetCreated returns the Created field value
func (o *Campaign) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// SetCreated sets field value
func (o *Campaign) SetCreated(v time.Time) {
	o.Created = v
}

// GetApplicationId returns the ApplicationId field value
func (o *Campaign) GetApplicationId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ApplicationId
}

// SetApplicationId sets field value
func (o *Campaign) SetApplicationId(v int32) {
	o.ApplicationId = v
}

// GetUserId returns the UserId field value
func (o *Campaign) GetUserId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UserId
}

// SetUserId sets field value
func (o *Campaign) SetUserId(v int32) {
	o.UserId = v
}

// GetName returns the Name field value
func (o *Campaign) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// SetName sets field value
func (o *Campaign) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *Campaign) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// SetDescription sets field value
func (o *Campaign) SetDescription(v string) {
	o.Description = v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *Campaign) GetStartTime() time.Time {
	if o == nil || o.StartTime == nil {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetStartTimeOk() (time.Time, bool) {
	if o == nil || o.StartTime == nil {
		var ret time.Time
		return ret, false
	}
	return *o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *Campaign) HasStartTime() bool {
	if o != nil && o.StartTime != nil {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *Campaign) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *Campaign) GetEndTime() time.Time {
	if o == nil || o.EndTime == nil {
		var ret time.Time
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetEndTimeOk() (time.Time, bool) {
	if o == nil || o.EndTime == nil {
		var ret time.Time
		return ret, false
	}
	return *o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *Campaign) HasEndTime() bool {
	if o != nil && o.EndTime != nil {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given time.Time and assigns it to the EndTime field.
func (o *Campaign) SetEndTime(v time.Time) {
	o.EndTime = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *Campaign) GetAttributes() map[string]interface{} {
	if o == nil || o.Attributes == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || o.Attributes == nil {
		var ret map[string]interface{}
		return ret, false
	}
	return *o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *Campaign) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *Campaign) SetAttributes(v map[string]interface{}) {
	o.Attributes = &v
}

// GetState returns the State field value
func (o *Campaign) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// SetState sets field value
func (o *Campaign) SetState(v string) {
	o.State = v
}

// GetActiveRulesetId returns the ActiveRulesetId field value if set, zero value otherwise.
func (o *Campaign) GetActiveRulesetId() int32 {
	if o == nil || o.ActiveRulesetId == nil {
		var ret int32
		return ret
	}
	return *o.ActiveRulesetId
}

// GetActiveRulesetIdOk returns a tuple with the ActiveRulesetId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetActiveRulesetIdOk() (int32, bool) {
	if o == nil || o.ActiveRulesetId == nil {
		var ret int32
		return ret, false
	}
	return *o.ActiveRulesetId, true
}

// HasActiveRulesetId returns a boolean if a field has been set.
func (o *Campaign) HasActiveRulesetId() bool {
	if o != nil && o.ActiveRulesetId != nil {
		return true
	}

	return false
}

// SetActiveRulesetId gets a reference to the given int32 and assigns it to the ActiveRulesetId field.
func (o *Campaign) SetActiveRulesetId(v int32) {
	o.ActiveRulesetId = &v
}

// GetTags returns the Tags field value
func (o *Campaign) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tags
}

// SetTags sets field value
func (o *Campaign) SetTags(v []string) {
	o.Tags = v
}

// GetFeatures returns the Features field value
func (o *Campaign) GetFeatures() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Features
}

// SetFeatures sets field value
func (o *Campaign) SetFeatures(v []string) {
	o.Features = v
}

// GetCouponSettings returns the CouponSettings field value if set, zero value otherwise.
func (o *Campaign) GetCouponSettings() CodeGeneratorSettings {
	if o == nil || o.CouponSettings == nil {
		var ret CodeGeneratorSettings
		return ret
	}
	return *o.CouponSettings
}

// GetCouponSettingsOk returns a tuple with the CouponSettings field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetCouponSettingsOk() (CodeGeneratorSettings, bool) {
	if o == nil || o.CouponSettings == nil {
		var ret CodeGeneratorSettings
		return ret, false
	}
	return *o.CouponSettings, true
}

// HasCouponSettings returns a boolean if a field has been set.
func (o *Campaign) HasCouponSettings() bool {
	if o != nil && o.CouponSettings != nil {
		return true
	}

	return false
}

// SetCouponSettings gets a reference to the given CodeGeneratorSettings and assigns it to the CouponSettings field.
func (o *Campaign) SetCouponSettings(v CodeGeneratorSettings) {
	o.CouponSettings = &v
}

// GetReferralSettings returns the ReferralSettings field value if set, zero value otherwise.
func (o *Campaign) GetReferralSettings() CodeGeneratorSettings {
	if o == nil || o.ReferralSettings == nil {
		var ret CodeGeneratorSettings
		return ret
	}
	return *o.ReferralSettings
}

// GetReferralSettingsOk returns a tuple with the ReferralSettings field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetReferralSettingsOk() (CodeGeneratorSettings, bool) {
	if o == nil || o.ReferralSettings == nil {
		var ret CodeGeneratorSettings
		return ret, false
	}
	return *o.ReferralSettings, true
}

// HasReferralSettings returns a boolean if a field has been set.
func (o *Campaign) HasReferralSettings() bool {
	if o != nil && o.ReferralSettings != nil {
		return true
	}

	return false
}

// SetReferralSettings gets a reference to the given CodeGeneratorSettings and assigns it to the ReferralSettings field.
func (o *Campaign) SetReferralSettings(v CodeGeneratorSettings) {
	o.ReferralSettings = &v
}

// GetLimits returns the Limits field value
func (o *Campaign) GetLimits() []LimitConfig {
	if o == nil {
		var ret []LimitConfig
		return ret
	}

	return o.Limits
}

// SetLimits sets field value
func (o *Campaign) SetLimits(v []LimitConfig) {
	o.Limits = v
}

// GetCouponRedemptionCount returns the CouponRedemptionCount field value if set, zero value otherwise.
func (o *Campaign) GetCouponRedemptionCount() int32 {
	if o == nil || o.CouponRedemptionCount == nil {
		var ret int32
		return ret
	}
	return *o.CouponRedemptionCount
}

// GetCouponRedemptionCountOk returns a tuple with the CouponRedemptionCount field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetCouponRedemptionCountOk() (int32, bool) {
	if o == nil || o.CouponRedemptionCount == nil {
		var ret int32
		return ret, false
	}
	return *o.CouponRedemptionCount, true
}

// HasCouponRedemptionCount returns a boolean if a field has been set.
func (o *Campaign) HasCouponRedemptionCount() bool {
	if o != nil && o.CouponRedemptionCount != nil {
		return true
	}

	return false
}

// SetCouponRedemptionCount gets a reference to the given int32 and assigns it to the CouponRedemptionCount field.
func (o *Campaign) SetCouponRedemptionCount(v int32) {
	o.CouponRedemptionCount = &v
}

// GetReferralRedemptionCount returns the ReferralRedemptionCount field value if set, zero value otherwise.
func (o *Campaign) GetReferralRedemptionCount() int32 {
	if o == nil || o.ReferralRedemptionCount == nil {
		var ret int32
		return ret
	}
	return *o.ReferralRedemptionCount
}

// GetReferralRedemptionCountOk returns a tuple with the ReferralRedemptionCount field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetReferralRedemptionCountOk() (int32, bool) {
	if o == nil || o.ReferralRedemptionCount == nil {
		var ret int32
		return ret, false
	}
	return *o.ReferralRedemptionCount, true
}

// HasReferralRedemptionCount returns a boolean if a field has been set.
func (o *Campaign) HasReferralRedemptionCount() bool {
	if o != nil && o.ReferralRedemptionCount != nil {
		return true
	}

	return false
}

// SetReferralRedemptionCount gets a reference to the given int32 and assigns it to the ReferralRedemptionCount field.
func (o *Campaign) SetReferralRedemptionCount(v int32) {
	o.ReferralRedemptionCount = &v
}

// GetDiscountCount returns the DiscountCount field value if set, zero value otherwise.
func (o *Campaign) GetDiscountCount() int32 {
	if o == nil || o.DiscountCount == nil {
		var ret int32
		return ret
	}
	return *o.DiscountCount
}

// GetDiscountCountOk returns a tuple with the DiscountCount field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetDiscountCountOk() (int32, bool) {
	if o == nil || o.DiscountCount == nil {
		var ret int32
		return ret, false
	}
	return *o.DiscountCount, true
}

// HasDiscountCount returns a boolean if a field has been set.
func (o *Campaign) HasDiscountCount() bool {
	if o != nil && o.DiscountCount != nil {
		return true
	}

	return false
}

// SetDiscountCount gets a reference to the given int32 and assigns it to the DiscountCount field.
func (o *Campaign) SetDiscountCount(v int32) {
	o.DiscountCount = &v
}

// GetDiscountEffectCount returns the DiscountEffectCount field value if set, zero value otherwise.
func (o *Campaign) GetDiscountEffectCount() int32 {
	if o == nil || o.DiscountEffectCount == nil {
		var ret int32
		return ret
	}
	return *o.DiscountEffectCount
}

// GetDiscountEffectCountOk returns a tuple with the DiscountEffectCount field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetDiscountEffectCountOk() (int32, bool) {
	if o == nil || o.DiscountEffectCount == nil {
		var ret int32
		return ret, false
	}
	return *o.DiscountEffectCount, true
}

// HasDiscountEffectCount returns a boolean if a field has been set.
func (o *Campaign) HasDiscountEffectCount() bool {
	if o != nil && o.DiscountEffectCount != nil {
		return true
	}

	return false
}

// SetDiscountEffectCount gets a reference to the given int32 and assigns it to the DiscountEffectCount field.
func (o *Campaign) SetDiscountEffectCount(v int32) {
	o.DiscountEffectCount = &v
}

// GetCouponCreationCount returns the CouponCreationCount field value if set, zero value otherwise.
func (o *Campaign) GetCouponCreationCount() int32 {
	if o == nil || o.CouponCreationCount == nil {
		var ret int32
		return ret
	}
	return *o.CouponCreationCount
}

// GetCouponCreationCountOk returns a tuple with the CouponCreationCount field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetCouponCreationCountOk() (int32, bool) {
	if o == nil || o.CouponCreationCount == nil {
		var ret int32
		return ret, false
	}
	return *o.CouponCreationCount, true
}

// HasCouponCreationCount returns a boolean if a field has been set.
func (o *Campaign) HasCouponCreationCount() bool {
	if o != nil && o.CouponCreationCount != nil {
		return true
	}

	return false
}

// SetCouponCreationCount gets a reference to the given int32 and assigns it to the CouponCreationCount field.
func (o *Campaign) SetCouponCreationCount(v int32) {
	o.CouponCreationCount = &v
}

// GetLastActivity returns the LastActivity field value if set, zero value otherwise.
func (o *Campaign) GetLastActivity() time.Time {
	if o == nil || o.LastActivity == nil {
		var ret time.Time
		return ret
	}
	return *o.LastActivity
}

// GetLastActivityOk returns a tuple with the LastActivity field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetLastActivityOk() (time.Time, bool) {
	if o == nil || o.LastActivity == nil {
		var ret time.Time
		return ret, false
	}
	return *o.LastActivity, true
}

// HasLastActivity returns a boolean if a field has been set.
func (o *Campaign) HasLastActivity() bool {
	if o != nil && o.LastActivity != nil {
		return true
	}

	return false
}

// SetLastActivity gets a reference to the given time.Time and assigns it to the LastActivity field.
func (o *Campaign) SetLastActivity(v time.Time) {
	o.LastActivity = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *Campaign) GetUpdated() time.Time {
	if o == nil || o.Updated == nil {
		var ret time.Time
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetUpdatedOk() (time.Time, bool) {
	if o == nil || o.Updated == nil {
		var ret time.Time
		return ret, false
	}
	return *o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *Campaign) HasUpdated() bool {
	if o != nil && o.Updated != nil {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given time.Time and assigns it to the Updated field.
func (o *Campaign) SetUpdated(v time.Time) {
	o.Updated = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *Campaign) GetCreatedBy() string {
	if o == nil || o.CreatedBy == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetCreatedByOk() (string, bool) {
	if o == nil || o.CreatedBy == nil {
		var ret string
		return ret, false
	}
	return *o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *Campaign) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *Campaign) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *Campaign) GetUpdatedBy() string {
	if o == nil || o.UpdatedBy == nil {
		var ret string
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetUpdatedByOk() (string, bool) {
	if o == nil || o.UpdatedBy == nil {
		var ret string
		return ret, false
	}
	return *o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *Campaign) HasUpdatedBy() bool {
	if o != nil && o.UpdatedBy != nil {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given string and assigns it to the UpdatedBy field.
func (o *Campaign) SetUpdatedBy(v string) {
	o.UpdatedBy = &v
}

type NullableCampaign struct {
	Value        Campaign
	ExplicitNull bool
}

func (v NullableCampaign) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableCampaign) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
