# CustomerProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IntegrationId** | Pointer to **string** | The integration ID for this entity sent to and used in the Talon.One system. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The exact moment this entity was created. | 
**Attributes** | Pointer to [**map[string]interface{}**](.md) | Arbitrary properties associated with this item | 
**AccountId** | Pointer to **int32** | The ID of the Talon.One account that owns this profile. | 
**ClosedSessions** | Pointer to **int32** | The total amount of closed sessions by a customer. A closed session is a successful purchase. | 
**TotalSales** | Pointer to **float32** | Sum of all purchases made by this customer | 
**LoyaltyMemberships** | Pointer to [**[]LoyaltyMembership**](LoyaltyMembership.md) | A list of loyalty programs joined by the customer | [optional] 
**AudienceMemberships** | Pointer to [**[]AudienceMembership**](AudienceMembership.md) | A list of audiences the customer belongs to | [optional] 
**LastActivity** | Pointer to [**time.Time**](time.Time.md) | Timestamp of the most recent event received from this customer | 

## Methods

### GetIntegrationId

`func (o *CustomerProfile) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *CustomerProfile) GetIntegrationIdOk() (string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIntegrationId

`func (o *CustomerProfile) HasIntegrationId() bool`

HasIntegrationId returns a boolean if a field has been set.

### SetIntegrationId

`func (o *CustomerProfile) SetIntegrationId(v string)`

SetIntegrationId gets a reference to the given string and assigns it to the IntegrationId field.

### GetCreated

`func (o *CustomerProfile) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CustomerProfile) GetCreatedOk() (time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *CustomerProfile) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *CustomerProfile) SetCreated(v time.Time)`

SetCreated gets a reference to the given time.Time and assigns it to the Created field.

### GetAttributes

`func (o *CustomerProfile) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CustomerProfile) GetAttributesOk() (map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAttributes

`func (o *CustomerProfile) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributes

`func (o *CustomerProfile) SetAttributes(v map[string]interface{})`

SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.

### GetAccountId

`func (o *CustomerProfile) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CustomerProfile) GetAccountIdOk() (int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccountId

`func (o *CustomerProfile) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountId

`func (o *CustomerProfile) SetAccountId(v int32)`

SetAccountId gets a reference to the given int32 and assigns it to the AccountId field.

### GetClosedSessions

`func (o *CustomerProfile) GetClosedSessions() int32`

GetClosedSessions returns the ClosedSessions field if non-nil, zero value otherwise.

### GetClosedSessionsOk

`func (o *CustomerProfile) GetClosedSessionsOk() (int32, bool)`

GetClosedSessionsOk returns a tuple with the ClosedSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasClosedSessions

`func (o *CustomerProfile) HasClosedSessions() bool`

HasClosedSessions returns a boolean if a field has been set.

### SetClosedSessions

`func (o *CustomerProfile) SetClosedSessions(v int32)`

SetClosedSessions gets a reference to the given int32 and assigns it to the ClosedSessions field.

### GetTotalSales

`func (o *CustomerProfile) GetTotalSales() float32`

GetTotalSales returns the TotalSales field if non-nil, zero value otherwise.

### GetTotalSalesOk

`func (o *CustomerProfile) GetTotalSalesOk() (float32, bool)`

GetTotalSalesOk returns a tuple with the TotalSales field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotalSales

`func (o *CustomerProfile) HasTotalSales() bool`

HasTotalSales returns a boolean if a field has been set.

### SetTotalSales

`func (o *CustomerProfile) SetTotalSales(v float32)`

SetTotalSales gets a reference to the given float32 and assigns it to the TotalSales field.

### GetLoyaltyMemberships

`func (o *CustomerProfile) GetLoyaltyMemberships() []LoyaltyMembership`

GetLoyaltyMemberships returns the LoyaltyMemberships field if non-nil, zero value otherwise.

### GetLoyaltyMembershipsOk

`func (o *CustomerProfile) GetLoyaltyMembershipsOk() ([]LoyaltyMembership, bool)`

GetLoyaltyMembershipsOk returns a tuple with the LoyaltyMemberships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLoyaltyMemberships

`func (o *CustomerProfile) HasLoyaltyMemberships() bool`

HasLoyaltyMemberships returns a boolean if a field has been set.

### SetLoyaltyMemberships

`func (o *CustomerProfile) SetLoyaltyMemberships(v []LoyaltyMembership)`

SetLoyaltyMemberships gets a reference to the given []LoyaltyMembership and assigns it to the LoyaltyMemberships field.

### GetAudienceMemberships

`func (o *CustomerProfile) GetAudienceMemberships() []AudienceMembership`

GetAudienceMemberships returns the AudienceMemberships field if non-nil, zero value otherwise.

### GetAudienceMembershipsOk

`func (o *CustomerProfile) GetAudienceMembershipsOk() ([]AudienceMembership, bool)`

GetAudienceMembershipsOk returns a tuple with the AudienceMemberships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAudienceMemberships

`func (o *CustomerProfile) HasAudienceMemberships() bool`

HasAudienceMemberships returns a boolean if a field has been set.

### SetAudienceMemberships

`func (o *CustomerProfile) SetAudienceMemberships(v []AudienceMembership)`

SetAudienceMemberships gets a reference to the given []AudienceMembership and assigns it to the AudienceMemberships field.

### GetLastActivity

`func (o *CustomerProfile) GetLastActivity() time.Time`

GetLastActivity returns the LastActivity field if non-nil, zero value otherwise.

### GetLastActivityOk

`func (o *CustomerProfile) GetLastActivityOk() (time.Time, bool)`

GetLastActivityOk returns a tuple with the LastActivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastActivity

`func (o *CustomerProfile) HasLastActivity() bool`

HasLastActivity returns a boolean if a field has been set.

### SetLastActivity

`func (o *CustomerProfile) SetLastActivity(v time.Time)`

SetLastActivity gets a reference to the given time.Time and assigns it to the LastActivity field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


