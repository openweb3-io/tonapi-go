# Action

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Status** | **string** |  | 
**TonTransfer** | Pointer to [**TonTransferAction**](TonTransferAction.md) |  | [optional] 
**ContractDeploy** | Pointer to [**ContractDeployAction**](ContractDeployAction.md) |  | [optional] 
**JettonTransfer** | Pointer to [**JettonTransferAction**](JettonTransferAction.md) |  | [optional] 
**JettonBurn** | Pointer to [**JettonBurnAction**](JettonBurnAction.md) |  | [optional] 
**JettonMint** | Pointer to [**JettonMintAction**](JettonMintAction.md) |  | [optional] 
**NftItemTransfer** | Pointer to [**NftItemTransferAction**](NftItemTransferAction.md) |  | [optional] 
**Subscribe** | Pointer to [**SubscriptionAction**](SubscriptionAction.md) |  | [optional] 
**UnSubscribe** | Pointer to [**UnSubscriptionAction**](UnSubscriptionAction.md) |  | [optional] 
**AuctionBid** | Pointer to [**AuctionBidAction**](AuctionBidAction.md) |  | [optional] 
**NftPurchase** | Pointer to [**NftPurchaseAction**](NftPurchaseAction.md) |  | [optional] 
**DepositStake** | Pointer to [**DepositStakeAction**](DepositStakeAction.md) |  | [optional] 
**WithdrawStake** | Pointer to [**WithdrawStakeAction**](WithdrawStakeAction.md) |  | [optional] 
**WithdrawStakeRequest** | Pointer to [**WithdrawStakeRequestAction**](WithdrawStakeRequestAction.md) |  | [optional] 
**ElectionsDepositStake** | Pointer to [**ElectionsDepositStakeAction**](ElectionsDepositStakeAction.md) |  | [optional] 
**ElectionsRecoverStake** | Pointer to [**ElectionsRecoverStakeAction**](ElectionsRecoverStakeAction.md) |  | [optional] 
**JettonSwap** | Pointer to [**JettonSwapAction**](JettonSwapAction.md) |  | [optional] 
**SmartContractExec** | Pointer to [**SmartContractAction**](SmartContractAction.md) |  | [optional] 
**DomainRenew** | Pointer to [**DomainRenewAction**](DomainRenewAction.md) |  | [optional] 
**InscriptionTransfer** | Pointer to [**InscriptionTransferAction**](InscriptionTransferAction.md) |  | [optional] 
**InscriptionMint** | Pointer to [**InscriptionMintAction**](InscriptionMintAction.md) |  | [optional] 
**SimplePreview** | [**ActionSimplePreview**](ActionSimplePreview.md) |  | 
**BaseTransactions** | **[]string** |  | 

## Methods

### NewAction

`func NewAction(type_ string, status string, simplePreview ActionSimplePreview, baseTransactions []string, ) *Action`

NewAction instantiates a new Action object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionWithDefaults

`func NewActionWithDefaults() *Action`

NewActionWithDefaults instantiates a new Action object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Action) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Action) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Action) SetType(v string)`

SetType sets Type field to given value.


### GetStatus

`func (o *Action) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Action) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Action) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTonTransfer

`func (o *Action) GetTonTransfer() TonTransferAction`

GetTonTransfer returns the TonTransfer field if non-nil, zero value otherwise.

### GetTonTransferOk

`func (o *Action) GetTonTransferOk() (*TonTransferAction, bool)`

GetTonTransferOk returns a tuple with the TonTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTonTransfer

`func (o *Action) SetTonTransfer(v TonTransferAction)`

SetTonTransfer sets TonTransfer field to given value.

### HasTonTransfer

`func (o *Action) HasTonTransfer() bool`

HasTonTransfer returns a boolean if a field has been set.

### GetContractDeploy

`func (o *Action) GetContractDeploy() ContractDeployAction`

GetContractDeploy returns the ContractDeploy field if non-nil, zero value otherwise.

### GetContractDeployOk

`func (o *Action) GetContractDeployOk() (*ContractDeployAction, bool)`

GetContractDeployOk returns a tuple with the ContractDeploy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractDeploy

`func (o *Action) SetContractDeploy(v ContractDeployAction)`

SetContractDeploy sets ContractDeploy field to given value.

### HasContractDeploy

`func (o *Action) HasContractDeploy() bool`

HasContractDeploy returns a boolean if a field has been set.

### GetJettonTransfer

`func (o *Action) GetJettonTransfer() JettonTransferAction`

GetJettonTransfer returns the JettonTransfer field if non-nil, zero value otherwise.

### GetJettonTransferOk

`func (o *Action) GetJettonTransferOk() (*JettonTransferAction, bool)`

GetJettonTransferOk returns a tuple with the JettonTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJettonTransfer

`func (o *Action) SetJettonTransfer(v JettonTransferAction)`

SetJettonTransfer sets JettonTransfer field to given value.

### HasJettonTransfer

`func (o *Action) HasJettonTransfer() bool`

HasJettonTransfer returns a boolean if a field has been set.

### GetJettonBurn

`func (o *Action) GetJettonBurn() JettonBurnAction`

GetJettonBurn returns the JettonBurn field if non-nil, zero value otherwise.

### GetJettonBurnOk

`func (o *Action) GetJettonBurnOk() (*JettonBurnAction, bool)`

GetJettonBurnOk returns a tuple with the JettonBurn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJettonBurn

`func (o *Action) SetJettonBurn(v JettonBurnAction)`

SetJettonBurn sets JettonBurn field to given value.

### HasJettonBurn

`func (o *Action) HasJettonBurn() bool`

HasJettonBurn returns a boolean if a field has been set.

### GetJettonMint

`func (o *Action) GetJettonMint() JettonMintAction`

GetJettonMint returns the JettonMint field if non-nil, zero value otherwise.

### GetJettonMintOk

`func (o *Action) GetJettonMintOk() (*JettonMintAction, bool)`

GetJettonMintOk returns a tuple with the JettonMint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJettonMint

`func (o *Action) SetJettonMint(v JettonMintAction)`

SetJettonMint sets JettonMint field to given value.

### HasJettonMint

`func (o *Action) HasJettonMint() bool`

HasJettonMint returns a boolean if a field has been set.

### GetNftItemTransfer

`func (o *Action) GetNftItemTransfer() NftItemTransferAction`

GetNftItemTransfer returns the NftItemTransfer field if non-nil, zero value otherwise.

### GetNftItemTransferOk

`func (o *Action) GetNftItemTransferOk() (*NftItemTransferAction, bool)`

GetNftItemTransferOk returns a tuple with the NftItemTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNftItemTransfer

`func (o *Action) SetNftItemTransfer(v NftItemTransferAction)`

SetNftItemTransfer sets NftItemTransfer field to given value.

### HasNftItemTransfer

`func (o *Action) HasNftItemTransfer() bool`

HasNftItemTransfer returns a boolean if a field has been set.

### GetSubscribe

`func (o *Action) GetSubscribe() SubscriptionAction`

GetSubscribe returns the Subscribe field if non-nil, zero value otherwise.

### GetSubscribeOk

`func (o *Action) GetSubscribeOk() (*SubscriptionAction, bool)`

GetSubscribeOk returns a tuple with the Subscribe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribe

`func (o *Action) SetSubscribe(v SubscriptionAction)`

SetSubscribe sets Subscribe field to given value.

### HasSubscribe

`func (o *Action) HasSubscribe() bool`

HasSubscribe returns a boolean if a field has been set.

### GetUnSubscribe

`func (o *Action) GetUnSubscribe() UnSubscriptionAction`

GetUnSubscribe returns the UnSubscribe field if non-nil, zero value otherwise.

### GetUnSubscribeOk

`func (o *Action) GetUnSubscribeOk() (*UnSubscriptionAction, bool)`

GetUnSubscribeOk returns a tuple with the UnSubscribe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnSubscribe

`func (o *Action) SetUnSubscribe(v UnSubscriptionAction)`

SetUnSubscribe sets UnSubscribe field to given value.

### HasUnSubscribe

`func (o *Action) HasUnSubscribe() bool`

HasUnSubscribe returns a boolean if a field has been set.

### GetAuctionBid

`func (o *Action) GetAuctionBid() AuctionBidAction`

GetAuctionBid returns the AuctionBid field if non-nil, zero value otherwise.

### GetAuctionBidOk

`func (o *Action) GetAuctionBidOk() (*AuctionBidAction, bool)`

GetAuctionBidOk returns a tuple with the AuctionBid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuctionBid

`func (o *Action) SetAuctionBid(v AuctionBidAction)`

SetAuctionBid sets AuctionBid field to given value.

### HasAuctionBid

`func (o *Action) HasAuctionBid() bool`

HasAuctionBid returns a boolean if a field has been set.

### GetNftPurchase

`func (o *Action) GetNftPurchase() NftPurchaseAction`

GetNftPurchase returns the NftPurchase field if non-nil, zero value otherwise.

### GetNftPurchaseOk

`func (o *Action) GetNftPurchaseOk() (*NftPurchaseAction, bool)`

GetNftPurchaseOk returns a tuple with the NftPurchase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNftPurchase

`func (o *Action) SetNftPurchase(v NftPurchaseAction)`

SetNftPurchase sets NftPurchase field to given value.

### HasNftPurchase

`func (o *Action) HasNftPurchase() bool`

HasNftPurchase returns a boolean if a field has been set.

### GetDepositStake

`func (o *Action) GetDepositStake() DepositStakeAction`

GetDepositStake returns the DepositStake field if non-nil, zero value otherwise.

### GetDepositStakeOk

`func (o *Action) GetDepositStakeOk() (*DepositStakeAction, bool)`

GetDepositStakeOk returns a tuple with the DepositStake field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepositStake

`func (o *Action) SetDepositStake(v DepositStakeAction)`

SetDepositStake sets DepositStake field to given value.

### HasDepositStake

`func (o *Action) HasDepositStake() bool`

HasDepositStake returns a boolean if a field has been set.

### GetWithdrawStake

`func (o *Action) GetWithdrawStake() WithdrawStakeAction`

GetWithdrawStake returns the WithdrawStake field if non-nil, zero value otherwise.

### GetWithdrawStakeOk

`func (o *Action) GetWithdrawStakeOk() (*WithdrawStakeAction, bool)`

GetWithdrawStakeOk returns a tuple with the WithdrawStake field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawStake

`func (o *Action) SetWithdrawStake(v WithdrawStakeAction)`

SetWithdrawStake sets WithdrawStake field to given value.

### HasWithdrawStake

`func (o *Action) HasWithdrawStake() bool`

HasWithdrawStake returns a boolean if a field has been set.

### GetWithdrawStakeRequest

`func (o *Action) GetWithdrawStakeRequest() WithdrawStakeRequestAction`

GetWithdrawStakeRequest returns the WithdrawStakeRequest field if non-nil, zero value otherwise.

### GetWithdrawStakeRequestOk

`func (o *Action) GetWithdrawStakeRequestOk() (*WithdrawStakeRequestAction, bool)`

GetWithdrawStakeRequestOk returns a tuple with the WithdrawStakeRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawStakeRequest

`func (o *Action) SetWithdrawStakeRequest(v WithdrawStakeRequestAction)`

SetWithdrawStakeRequest sets WithdrawStakeRequest field to given value.

### HasWithdrawStakeRequest

`func (o *Action) HasWithdrawStakeRequest() bool`

HasWithdrawStakeRequest returns a boolean if a field has been set.

### GetElectionsDepositStake

`func (o *Action) GetElectionsDepositStake() ElectionsDepositStakeAction`

GetElectionsDepositStake returns the ElectionsDepositStake field if non-nil, zero value otherwise.

### GetElectionsDepositStakeOk

`func (o *Action) GetElectionsDepositStakeOk() (*ElectionsDepositStakeAction, bool)`

GetElectionsDepositStakeOk returns a tuple with the ElectionsDepositStake field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElectionsDepositStake

`func (o *Action) SetElectionsDepositStake(v ElectionsDepositStakeAction)`

SetElectionsDepositStake sets ElectionsDepositStake field to given value.

### HasElectionsDepositStake

`func (o *Action) HasElectionsDepositStake() bool`

HasElectionsDepositStake returns a boolean if a field has been set.

### GetElectionsRecoverStake

`func (o *Action) GetElectionsRecoverStake() ElectionsRecoverStakeAction`

GetElectionsRecoverStake returns the ElectionsRecoverStake field if non-nil, zero value otherwise.

### GetElectionsRecoverStakeOk

`func (o *Action) GetElectionsRecoverStakeOk() (*ElectionsRecoverStakeAction, bool)`

GetElectionsRecoverStakeOk returns a tuple with the ElectionsRecoverStake field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElectionsRecoverStake

`func (o *Action) SetElectionsRecoverStake(v ElectionsRecoverStakeAction)`

SetElectionsRecoverStake sets ElectionsRecoverStake field to given value.

### HasElectionsRecoverStake

`func (o *Action) HasElectionsRecoverStake() bool`

HasElectionsRecoverStake returns a boolean if a field has been set.

### GetJettonSwap

`func (o *Action) GetJettonSwap() JettonSwapAction`

GetJettonSwap returns the JettonSwap field if non-nil, zero value otherwise.

### GetJettonSwapOk

`func (o *Action) GetJettonSwapOk() (*JettonSwapAction, bool)`

GetJettonSwapOk returns a tuple with the JettonSwap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJettonSwap

`func (o *Action) SetJettonSwap(v JettonSwapAction)`

SetJettonSwap sets JettonSwap field to given value.

### HasJettonSwap

`func (o *Action) HasJettonSwap() bool`

HasJettonSwap returns a boolean if a field has been set.

### GetSmartContractExec

`func (o *Action) GetSmartContractExec() SmartContractAction`

GetSmartContractExec returns the SmartContractExec field if non-nil, zero value otherwise.

### GetSmartContractExecOk

`func (o *Action) GetSmartContractExecOk() (*SmartContractAction, bool)`

GetSmartContractExecOk returns a tuple with the SmartContractExec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartContractExec

`func (o *Action) SetSmartContractExec(v SmartContractAction)`

SetSmartContractExec sets SmartContractExec field to given value.

### HasSmartContractExec

`func (o *Action) HasSmartContractExec() bool`

HasSmartContractExec returns a boolean if a field has been set.

### GetDomainRenew

`func (o *Action) GetDomainRenew() DomainRenewAction`

GetDomainRenew returns the DomainRenew field if non-nil, zero value otherwise.

### GetDomainRenewOk

`func (o *Action) GetDomainRenewOk() (*DomainRenewAction, bool)`

GetDomainRenewOk returns a tuple with the DomainRenew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainRenew

`func (o *Action) SetDomainRenew(v DomainRenewAction)`

SetDomainRenew sets DomainRenew field to given value.

### HasDomainRenew

`func (o *Action) HasDomainRenew() bool`

HasDomainRenew returns a boolean if a field has been set.

### GetInscriptionTransfer

`func (o *Action) GetInscriptionTransfer() InscriptionTransferAction`

GetInscriptionTransfer returns the InscriptionTransfer field if non-nil, zero value otherwise.

### GetInscriptionTransferOk

`func (o *Action) GetInscriptionTransferOk() (*InscriptionTransferAction, bool)`

GetInscriptionTransferOk returns a tuple with the InscriptionTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInscriptionTransfer

`func (o *Action) SetInscriptionTransfer(v InscriptionTransferAction)`

SetInscriptionTransfer sets InscriptionTransfer field to given value.

### HasInscriptionTransfer

`func (o *Action) HasInscriptionTransfer() bool`

HasInscriptionTransfer returns a boolean if a field has been set.

### GetInscriptionMint

`func (o *Action) GetInscriptionMint() InscriptionMintAction`

GetInscriptionMint returns the InscriptionMint field if non-nil, zero value otherwise.

### GetInscriptionMintOk

`func (o *Action) GetInscriptionMintOk() (*InscriptionMintAction, bool)`

GetInscriptionMintOk returns a tuple with the InscriptionMint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInscriptionMint

`func (o *Action) SetInscriptionMint(v InscriptionMintAction)`

SetInscriptionMint sets InscriptionMint field to given value.

### HasInscriptionMint

`func (o *Action) HasInscriptionMint() bool`

HasInscriptionMint returns a boolean if a field has been set.

### GetSimplePreview

`func (o *Action) GetSimplePreview() ActionSimplePreview`

GetSimplePreview returns the SimplePreview field if non-nil, zero value otherwise.

### GetSimplePreviewOk

`func (o *Action) GetSimplePreviewOk() (*ActionSimplePreview, bool)`

GetSimplePreviewOk returns a tuple with the SimplePreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimplePreview

`func (o *Action) SetSimplePreview(v ActionSimplePreview)`

SetSimplePreview sets SimplePreview field to given value.


### GetBaseTransactions

`func (o *Action) GetBaseTransactions() []string`

GetBaseTransactions returns the BaseTransactions field if non-nil, zero value otherwise.

### GetBaseTransactionsOk

`func (o *Action) GetBaseTransactionsOk() (*[]string, bool)`

GetBaseTransactionsOk returns a tuple with the BaseTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseTransactions

`func (o *Action) SetBaseTransactions(v []string)`

SetBaseTransactions sets BaseTransactions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


