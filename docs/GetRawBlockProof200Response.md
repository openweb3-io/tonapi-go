# GetRawBlockProof200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Complete** | **bool** |  | 
**From** | [**BlockRaw**](BlockRaw.md) |  | 
**To** | [**BlockRaw**](BlockRaw.md) |  | 
**Steps** | [**[]GetRawBlockProof200ResponseStepsInner**](GetRawBlockProof200ResponseStepsInner.md) |  | 

## Methods

### NewGetRawBlockProof200Response

`func NewGetRawBlockProof200Response(complete bool, from BlockRaw, to BlockRaw, steps []GetRawBlockProof200ResponseStepsInner, ) *GetRawBlockProof200Response`

NewGetRawBlockProof200Response instantiates a new GetRawBlockProof200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRawBlockProof200ResponseWithDefaults

`func NewGetRawBlockProof200ResponseWithDefaults() *GetRawBlockProof200Response`

NewGetRawBlockProof200ResponseWithDefaults instantiates a new GetRawBlockProof200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComplete

`func (o *GetRawBlockProof200Response) GetComplete() bool`

GetComplete returns the Complete field if non-nil, zero value otherwise.

### GetCompleteOk

`func (o *GetRawBlockProof200Response) GetCompleteOk() (*bool, bool)`

GetCompleteOk returns a tuple with the Complete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplete

`func (o *GetRawBlockProof200Response) SetComplete(v bool)`

SetComplete sets Complete field to given value.


### GetFrom

`func (o *GetRawBlockProof200Response) GetFrom() BlockRaw`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *GetRawBlockProof200Response) GetFromOk() (*BlockRaw, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *GetRawBlockProof200Response) SetFrom(v BlockRaw)`

SetFrom sets From field to given value.


### GetTo

`func (o *GetRawBlockProof200Response) GetTo() BlockRaw`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *GetRawBlockProof200Response) GetToOk() (*BlockRaw, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *GetRawBlockProof200Response) SetTo(v BlockRaw)`

SetTo sets To field to given value.


### GetSteps

`func (o *GetRawBlockProof200Response) GetSteps() []GetRawBlockProof200ResponseStepsInner`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *GetRawBlockProof200Response) GetStepsOk() (*[]GetRawBlockProof200ResponseStepsInner, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *GetRawBlockProof200Response) SetSteps(v []GetRawBlockProof200ResponseStepsInner)`

SetSteps sets Steps field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


