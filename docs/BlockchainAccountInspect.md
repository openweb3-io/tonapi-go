# BlockchainAccountInspect

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** |  | 
**CodeHash** | **string** |  | 
**Methods** | [**[]BlockchainAccountInspectMethodsInner**](BlockchainAccountInspectMethodsInner.md) |  | 
**Compiler** | Pointer to **string** |  | [optional] 

## Methods

### NewBlockchainAccountInspect

`func NewBlockchainAccountInspect(code string, codeHash string, methods []BlockchainAccountInspectMethodsInner, ) *BlockchainAccountInspect`

NewBlockchainAccountInspect instantiates a new BlockchainAccountInspect object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockchainAccountInspectWithDefaults

`func NewBlockchainAccountInspectWithDefaults() *BlockchainAccountInspect`

NewBlockchainAccountInspectWithDefaults instantiates a new BlockchainAccountInspect object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *BlockchainAccountInspect) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *BlockchainAccountInspect) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *BlockchainAccountInspect) SetCode(v string)`

SetCode sets Code field to given value.


### GetCodeHash

`func (o *BlockchainAccountInspect) GetCodeHash() string`

GetCodeHash returns the CodeHash field if non-nil, zero value otherwise.

### GetCodeHashOk

`func (o *BlockchainAccountInspect) GetCodeHashOk() (*string, bool)`

GetCodeHashOk returns a tuple with the CodeHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeHash

`func (o *BlockchainAccountInspect) SetCodeHash(v string)`

SetCodeHash sets CodeHash field to given value.


### GetMethods

`func (o *BlockchainAccountInspect) GetMethods() []BlockchainAccountInspectMethodsInner`

GetMethods returns the Methods field if non-nil, zero value otherwise.

### GetMethodsOk

`func (o *BlockchainAccountInspect) GetMethodsOk() (*[]BlockchainAccountInspectMethodsInner, bool)`

GetMethodsOk returns a tuple with the Methods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethods

`func (o *BlockchainAccountInspect) SetMethods(v []BlockchainAccountInspectMethodsInner)`

SetMethods sets Methods field to given value.


### GetCompiler

`func (o *BlockchainAccountInspect) GetCompiler() string`

GetCompiler returns the Compiler field if non-nil, zero value otherwise.

### GetCompilerOk

`func (o *BlockchainAccountInspect) GetCompilerOk() (*string, bool)`

GetCompilerOk returns a tuple with the Compiler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompiler

`func (o *BlockchainAccountInspect) SetCompiler(v string)`

SetCompiler sets Compiler field to given value.

### HasCompiler

`func (o *BlockchainAccountInspect) HasCompiler() bool`

HasCompiler returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


