# CloudEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spec** | Pointer to **string** | Cloud Event Open Telemetry specification | [optional] [readonly] 
**Source** | Pointer to **string** | Cloud Event source | [optional] 
**Id** | Pointer to **string** | Cloud Event identifier | [optional] 
**Type** | Pointer to **string** | Equinix supported event type | [optional] 
**Subject** | Pointer to **string** | Cloud Event subject | [optional] 
**Dataschema** | Pointer to **string** | Cloud Event dataschema reference | [optional] 
**Datacontenttype** | Pointer to **string** | Cloud Event data content type | [optional] 
**Severitynumber** | Pointer to **string** | Cloud Event severity number | [optional] 
**Severitytext** | Pointer to **string** | Cloud Event severity text | [optional] 
**Equinixorganization** | Pointer to **string** | Equinix organization identifier | [optional] 
**Equinixproject** | Pointer to **string** | Equinix project identifier | [optional] 
**Authtype** | Pointer to **string** | Cloud Event auth type | [optional] 
**Authid** | Pointer to **string** | Cloud Event user identifier | [optional] 
**Traceparent** | Pointer to **string** | Cloud Event traceparent | [optional] 
**Tracestate** | Pointer to **string** | Cloud Event tracestate | [optional] 
**Data** | Pointer to [**CloudEventData**](CloudEventData.md) |  | [optional] 

## Methods

### NewCloudEvent

`func NewCloudEvent() *CloudEvent`

NewCloudEvent instantiates a new CloudEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudEventWithDefaults

`func NewCloudEventWithDefaults() *CloudEvent`

NewCloudEventWithDefaults instantiates a new CloudEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpec

`func (o *CloudEvent) GetSpec() string`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *CloudEvent) GetSpecOk() (*string, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *CloudEvent) SetSpec(v string)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *CloudEvent) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetSource

`func (o *CloudEvent) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CloudEvent) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CloudEvent) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *CloudEvent) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetId

`func (o *CloudEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudEvent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *CloudEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CloudEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CloudEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CloudEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSubject

`func (o *CloudEvent) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *CloudEvent) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *CloudEvent) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *CloudEvent) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetDataschema

`func (o *CloudEvent) GetDataschema() string`

GetDataschema returns the Dataschema field if non-nil, zero value otherwise.

### GetDataschemaOk

`func (o *CloudEvent) GetDataschemaOk() (*string, bool)`

GetDataschemaOk returns a tuple with the Dataschema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataschema

`func (o *CloudEvent) SetDataschema(v string)`

SetDataschema sets Dataschema field to given value.

### HasDataschema

`func (o *CloudEvent) HasDataschema() bool`

HasDataschema returns a boolean if a field has been set.

### GetDatacontenttype

`func (o *CloudEvent) GetDatacontenttype() string`

GetDatacontenttype returns the Datacontenttype field if non-nil, zero value otherwise.

### GetDatacontenttypeOk

`func (o *CloudEvent) GetDatacontenttypeOk() (*string, bool)`

GetDatacontenttypeOk returns a tuple with the Datacontenttype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacontenttype

`func (o *CloudEvent) SetDatacontenttype(v string)`

SetDatacontenttype sets Datacontenttype field to given value.

### HasDatacontenttype

`func (o *CloudEvent) HasDatacontenttype() bool`

HasDatacontenttype returns a boolean if a field has been set.

### GetSeveritynumber

`func (o *CloudEvent) GetSeveritynumber() string`

GetSeveritynumber returns the Severitynumber field if non-nil, zero value otherwise.

### GetSeveritynumberOk

`func (o *CloudEvent) GetSeveritynumberOk() (*string, bool)`

GetSeveritynumberOk returns a tuple with the Severitynumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeveritynumber

`func (o *CloudEvent) SetSeveritynumber(v string)`

SetSeveritynumber sets Severitynumber field to given value.

### HasSeveritynumber

`func (o *CloudEvent) HasSeveritynumber() bool`

HasSeveritynumber returns a boolean if a field has been set.

### GetSeveritytext

`func (o *CloudEvent) GetSeveritytext() string`

GetSeveritytext returns the Severitytext field if non-nil, zero value otherwise.

### GetSeveritytextOk

`func (o *CloudEvent) GetSeveritytextOk() (*string, bool)`

GetSeveritytextOk returns a tuple with the Severitytext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeveritytext

`func (o *CloudEvent) SetSeveritytext(v string)`

SetSeveritytext sets Severitytext field to given value.

### HasSeveritytext

`func (o *CloudEvent) HasSeveritytext() bool`

HasSeveritytext returns a boolean if a field has been set.

### GetEquinixorganization

`func (o *CloudEvent) GetEquinixorganization() string`

GetEquinixorganization returns the Equinixorganization field if non-nil, zero value otherwise.

### GetEquinixorganizationOk

`func (o *CloudEvent) GetEquinixorganizationOk() (*string, bool)`

GetEquinixorganizationOk returns a tuple with the Equinixorganization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquinixorganization

`func (o *CloudEvent) SetEquinixorganization(v string)`

SetEquinixorganization sets Equinixorganization field to given value.

### HasEquinixorganization

`func (o *CloudEvent) HasEquinixorganization() bool`

HasEquinixorganization returns a boolean if a field has been set.

### GetEquinixproject

`func (o *CloudEvent) GetEquinixproject() string`

GetEquinixproject returns the Equinixproject field if non-nil, zero value otherwise.

### GetEquinixprojectOk

`func (o *CloudEvent) GetEquinixprojectOk() (*string, bool)`

GetEquinixprojectOk returns a tuple with the Equinixproject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquinixproject

`func (o *CloudEvent) SetEquinixproject(v string)`

SetEquinixproject sets Equinixproject field to given value.

### HasEquinixproject

`func (o *CloudEvent) HasEquinixproject() bool`

HasEquinixproject returns a boolean if a field has been set.

### GetAuthtype

`func (o *CloudEvent) GetAuthtype() string`

GetAuthtype returns the Authtype field if non-nil, zero value otherwise.

### GetAuthtypeOk

`func (o *CloudEvent) GetAuthtypeOk() (*string, bool)`

GetAuthtypeOk returns a tuple with the Authtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthtype

`func (o *CloudEvent) SetAuthtype(v string)`

SetAuthtype sets Authtype field to given value.

### HasAuthtype

`func (o *CloudEvent) HasAuthtype() bool`

HasAuthtype returns a boolean if a field has been set.

### GetAuthid

`func (o *CloudEvent) GetAuthid() string`

GetAuthid returns the Authid field if non-nil, zero value otherwise.

### GetAuthidOk

`func (o *CloudEvent) GetAuthidOk() (*string, bool)`

GetAuthidOk returns a tuple with the Authid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthid

`func (o *CloudEvent) SetAuthid(v string)`

SetAuthid sets Authid field to given value.

### HasAuthid

`func (o *CloudEvent) HasAuthid() bool`

HasAuthid returns a boolean if a field has been set.

### GetTraceparent

`func (o *CloudEvent) GetTraceparent() string`

GetTraceparent returns the Traceparent field if non-nil, zero value otherwise.

### GetTraceparentOk

`func (o *CloudEvent) GetTraceparentOk() (*string, bool)`

GetTraceparentOk returns a tuple with the Traceparent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceparent

`func (o *CloudEvent) SetTraceparent(v string)`

SetTraceparent sets Traceparent field to given value.

### HasTraceparent

`func (o *CloudEvent) HasTraceparent() bool`

HasTraceparent returns a boolean if a field has been set.

### GetTracestate

`func (o *CloudEvent) GetTracestate() string`

GetTracestate returns the Tracestate field if non-nil, zero value otherwise.

### GetTracestateOk

`func (o *CloudEvent) GetTracestateOk() (*string, bool)`

GetTracestateOk returns a tuple with the Tracestate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTracestate

`func (o *CloudEvent) SetTracestate(v string)`

SetTracestate sets Tracestate field to given value.

### HasTracestate

`func (o *CloudEvent) HasTracestate() bool`

HasTracestate returns a boolean if a field has been set.

### GetData

`func (o *CloudEvent) GetData() CloudEventData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CloudEvent) GetDataOk() (*CloudEventData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CloudEvent) SetData(v CloudEventData)`

SetData sets Data field to given value.

### HasData

`func (o *CloudEvent) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


