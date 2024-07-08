# BlockchainConfig29

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Flags** | Pointer to **int32** |  | [optional] 
**NewCatchainIds** | Pointer to **bool** |  | [optional] 
**RoundCandidates** | **int64** |  | 
**NextCandidateDelayMs** | **int64** |  | 
**ConsensusTimeoutMs** | **int64** |  | 
**FastAttempts** | **int64** |  | 
**AttemptDuration** | **int64** |  | 
**CatchainMaxDeps** | **int64** |  | 
**MaxBlockBytes** | **int64** |  | 
**MaxCollatedBytes** | **int64** |  | 
**ProtoVersion** | Pointer to **int64** |  | [optional] 
**CatchainMaxBlocksCoeff** | Pointer to **int64** |  | [optional] 

## Methods

### NewBlockchainConfig29

`func NewBlockchainConfig29(roundCandidates int64, nextCandidateDelayMs int64, consensusTimeoutMs int64, fastAttempts int64, attemptDuration int64, catchainMaxDeps int64, maxBlockBytes int64, maxCollatedBytes int64, ) *BlockchainConfig29`

NewBlockchainConfig29 instantiates a new BlockchainConfig29 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockchainConfig29WithDefaults

`func NewBlockchainConfig29WithDefaults() *BlockchainConfig29`

NewBlockchainConfig29WithDefaults instantiates a new BlockchainConfig29 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlags

`func (o *BlockchainConfig29) GetFlags() int32`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *BlockchainConfig29) GetFlagsOk() (*int32, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *BlockchainConfig29) SetFlags(v int32)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *BlockchainConfig29) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetNewCatchainIds

`func (o *BlockchainConfig29) GetNewCatchainIds() bool`

GetNewCatchainIds returns the NewCatchainIds field if non-nil, zero value otherwise.

### GetNewCatchainIdsOk

`func (o *BlockchainConfig29) GetNewCatchainIdsOk() (*bool, bool)`

GetNewCatchainIdsOk returns a tuple with the NewCatchainIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewCatchainIds

`func (o *BlockchainConfig29) SetNewCatchainIds(v bool)`

SetNewCatchainIds sets NewCatchainIds field to given value.

### HasNewCatchainIds

`func (o *BlockchainConfig29) HasNewCatchainIds() bool`

HasNewCatchainIds returns a boolean if a field has been set.

### GetRoundCandidates

`func (o *BlockchainConfig29) GetRoundCandidates() int64`

GetRoundCandidates returns the RoundCandidates field if non-nil, zero value otherwise.

### GetRoundCandidatesOk

`func (o *BlockchainConfig29) GetRoundCandidatesOk() (*int64, bool)`

GetRoundCandidatesOk returns a tuple with the RoundCandidates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoundCandidates

`func (o *BlockchainConfig29) SetRoundCandidates(v int64)`

SetRoundCandidates sets RoundCandidates field to given value.


### GetNextCandidateDelayMs

`func (o *BlockchainConfig29) GetNextCandidateDelayMs() int64`

GetNextCandidateDelayMs returns the NextCandidateDelayMs field if non-nil, zero value otherwise.

### GetNextCandidateDelayMsOk

`func (o *BlockchainConfig29) GetNextCandidateDelayMsOk() (*int64, bool)`

GetNextCandidateDelayMsOk returns a tuple with the NextCandidateDelayMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCandidateDelayMs

`func (o *BlockchainConfig29) SetNextCandidateDelayMs(v int64)`

SetNextCandidateDelayMs sets NextCandidateDelayMs field to given value.


### GetConsensusTimeoutMs

`func (o *BlockchainConfig29) GetConsensusTimeoutMs() int64`

GetConsensusTimeoutMs returns the ConsensusTimeoutMs field if non-nil, zero value otherwise.

### GetConsensusTimeoutMsOk

`func (o *BlockchainConfig29) GetConsensusTimeoutMsOk() (*int64, bool)`

GetConsensusTimeoutMsOk returns a tuple with the ConsensusTimeoutMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsensusTimeoutMs

`func (o *BlockchainConfig29) SetConsensusTimeoutMs(v int64)`

SetConsensusTimeoutMs sets ConsensusTimeoutMs field to given value.


### GetFastAttempts

`func (o *BlockchainConfig29) GetFastAttempts() int64`

GetFastAttempts returns the FastAttempts field if non-nil, zero value otherwise.

### GetFastAttemptsOk

`func (o *BlockchainConfig29) GetFastAttemptsOk() (*int64, bool)`

GetFastAttemptsOk returns a tuple with the FastAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFastAttempts

`func (o *BlockchainConfig29) SetFastAttempts(v int64)`

SetFastAttempts sets FastAttempts field to given value.


### GetAttemptDuration

`func (o *BlockchainConfig29) GetAttemptDuration() int64`

GetAttemptDuration returns the AttemptDuration field if non-nil, zero value otherwise.

### GetAttemptDurationOk

`func (o *BlockchainConfig29) GetAttemptDurationOk() (*int64, bool)`

GetAttemptDurationOk returns a tuple with the AttemptDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptDuration

`func (o *BlockchainConfig29) SetAttemptDuration(v int64)`

SetAttemptDuration sets AttemptDuration field to given value.


### GetCatchainMaxDeps

`func (o *BlockchainConfig29) GetCatchainMaxDeps() int64`

GetCatchainMaxDeps returns the CatchainMaxDeps field if non-nil, zero value otherwise.

### GetCatchainMaxDepsOk

`func (o *BlockchainConfig29) GetCatchainMaxDepsOk() (*int64, bool)`

GetCatchainMaxDepsOk returns a tuple with the CatchainMaxDeps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatchainMaxDeps

`func (o *BlockchainConfig29) SetCatchainMaxDeps(v int64)`

SetCatchainMaxDeps sets CatchainMaxDeps field to given value.


### GetMaxBlockBytes

`func (o *BlockchainConfig29) GetMaxBlockBytes() int64`

GetMaxBlockBytes returns the MaxBlockBytes field if non-nil, zero value otherwise.

### GetMaxBlockBytesOk

`func (o *BlockchainConfig29) GetMaxBlockBytesOk() (*int64, bool)`

GetMaxBlockBytesOk returns a tuple with the MaxBlockBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBlockBytes

`func (o *BlockchainConfig29) SetMaxBlockBytes(v int64)`

SetMaxBlockBytes sets MaxBlockBytes field to given value.


### GetMaxCollatedBytes

`func (o *BlockchainConfig29) GetMaxCollatedBytes() int64`

GetMaxCollatedBytes returns the MaxCollatedBytes field if non-nil, zero value otherwise.

### GetMaxCollatedBytesOk

`func (o *BlockchainConfig29) GetMaxCollatedBytesOk() (*int64, bool)`

GetMaxCollatedBytesOk returns a tuple with the MaxCollatedBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCollatedBytes

`func (o *BlockchainConfig29) SetMaxCollatedBytes(v int64)`

SetMaxCollatedBytes sets MaxCollatedBytes field to given value.


### GetProtoVersion

`func (o *BlockchainConfig29) GetProtoVersion() int64`

GetProtoVersion returns the ProtoVersion field if non-nil, zero value otherwise.

### GetProtoVersionOk

`func (o *BlockchainConfig29) GetProtoVersionOk() (*int64, bool)`

GetProtoVersionOk returns a tuple with the ProtoVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtoVersion

`func (o *BlockchainConfig29) SetProtoVersion(v int64)`

SetProtoVersion sets ProtoVersion field to given value.

### HasProtoVersion

`func (o *BlockchainConfig29) HasProtoVersion() bool`

HasProtoVersion returns a boolean if a field has been set.

### GetCatchainMaxBlocksCoeff

`func (o *BlockchainConfig29) GetCatchainMaxBlocksCoeff() int64`

GetCatchainMaxBlocksCoeff returns the CatchainMaxBlocksCoeff field if non-nil, zero value otherwise.

### GetCatchainMaxBlocksCoeffOk

`func (o *BlockchainConfig29) GetCatchainMaxBlocksCoeffOk() (*int64, bool)`

GetCatchainMaxBlocksCoeffOk returns a tuple with the CatchainMaxBlocksCoeff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatchainMaxBlocksCoeff

`func (o *BlockchainConfig29) SetCatchainMaxBlocksCoeff(v int64)`

SetCatchainMaxBlocksCoeff sets CatchainMaxBlocksCoeff field to given value.

### HasCatchainMaxBlocksCoeff

`func (o *BlockchainConfig29) HasCatchainMaxBlocksCoeff() bool`

HasCatchainMaxBlocksCoeff returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


