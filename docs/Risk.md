# Risk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransferAllRemainingBalance** | **bool** | transfer all the remaining balance of the wallet. | 
**Ton** | **int64** |  | 
**Jettons** | [**[]JettonQuantity**](JettonQuantity.md) |  | 
**Nfts** | [**[]NftItem**](NftItem.md) |  | 

## Methods

### NewRisk

`func NewRisk(transferAllRemainingBalance bool, ton int64, jettons []JettonQuantity, nfts []NftItem, ) *Risk`

NewRisk instantiates a new Risk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskWithDefaults

`func NewRiskWithDefaults() *Risk`

NewRiskWithDefaults instantiates a new Risk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransferAllRemainingBalance

`func (o *Risk) GetTransferAllRemainingBalance() bool`

GetTransferAllRemainingBalance returns the TransferAllRemainingBalance field if non-nil, zero value otherwise.

### GetTransferAllRemainingBalanceOk

`func (o *Risk) GetTransferAllRemainingBalanceOk() (*bool, bool)`

GetTransferAllRemainingBalanceOk returns a tuple with the TransferAllRemainingBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferAllRemainingBalance

`func (o *Risk) SetTransferAllRemainingBalance(v bool)`

SetTransferAllRemainingBalance sets TransferAllRemainingBalance field to given value.


### GetTon

`func (o *Risk) GetTon() int64`

GetTon returns the Ton field if non-nil, zero value otherwise.

### GetTonOk

`func (o *Risk) GetTonOk() (*int64, bool)`

GetTonOk returns a tuple with the Ton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTon

`func (o *Risk) SetTon(v int64)`

SetTon sets Ton field to given value.


### GetJettons

`func (o *Risk) GetJettons() []JettonQuantity`

GetJettons returns the Jettons field if non-nil, zero value otherwise.

### GetJettonsOk

`func (o *Risk) GetJettonsOk() (*[]JettonQuantity, bool)`

GetJettonsOk returns a tuple with the Jettons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJettons

`func (o *Risk) SetJettons(v []JettonQuantity)`

SetJettons sets Jettons field to given value.


### GetNfts

`func (o *Risk) GetNfts() []NftItem`

GetNfts returns the Nfts field if non-nil, zero value otherwise.

### GetNftsOk

`func (o *Risk) GetNftsOk() (*[]NftItem, bool)`

GetNftsOk returns a tuple with the Nfts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfts

`func (o *Risk) SetNfts(v []NftItem)`

SetNfts sets Nfts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


