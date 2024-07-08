# BlockchainConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Raw** | **string** | config boc in base64 format | 
**Var0** | **string** | config address | 
**Var1** | **string** | elector address | 
**Var2** | **string** | minter address | 
**Var3** | Pointer to **string** | The address of the transaction fee collector. | [optional] 
**Var4** | **string** | dns root address | 
**Var5** | Pointer to [**BlockchainConfig5**](BlockchainConfig5.md) |  | [optional] 
**Var6** | Pointer to [**BlockchainConfig6**](BlockchainConfig6.md) |  | [optional] 
**Var7** | Pointer to [**BlockchainConfig7**](BlockchainConfig7.md) |  | [optional] 
**Var8** | Pointer to [**BlockchainConfig8**](BlockchainConfig8.md) |  | [optional] 
**Var9** | Pointer to [**BlockchainConfig9**](BlockchainConfig9.md) |  | [optional] 
**Var10** | Pointer to [**BlockchainConfig10**](BlockchainConfig10.md) |  | [optional] 
**Var11** | Pointer to [**BlockchainConfig11**](BlockchainConfig11.md) |  | [optional] 
**Var12** | Pointer to [**BlockchainConfig12**](BlockchainConfig12.md) |  | [optional] 
**Var13** | Pointer to [**BlockchainConfig13**](BlockchainConfig13.md) |  | [optional] 
**Var14** | Pointer to [**BlockchainConfig14**](BlockchainConfig14.md) |  | [optional] 
**Var15** | Pointer to [**BlockchainConfig15**](BlockchainConfig15.md) |  | [optional] 
**Var16** | Pointer to [**BlockchainConfig16**](BlockchainConfig16.md) |  | [optional] 
**Var17** | Pointer to [**BlockchainConfig17**](BlockchainConfig17.md) |  | [optional] 
**Var18** | Pointer to [**BlockchainConfig18**](BlockchainConfig18.md) |  | [optional] 
**Var20** | Pointer to [**BlockchainConfig20**](BlockchainConfig20.md) |  | [optional] 
**Var21** | Pointer to [**BlockchainConfig21**](BlockchainConfig21.md) |  | [optional] 
**Var22** | Pointer to [**BlockchainConfig22**](BlockchainConfig22.md) |  | [optional] 
**Var23** | Pointer to [**BlockchainConfig23**](BlockchainConfig23.md) |  | [optional] 
**Var24** | Pointer to [**BlockchainConfig24**](BlockchainConfig24.md) |  | [optional] 
**Var25** | Pointer to [**BlockchainConfig25**](BlockchainConfig25.md) |  | [optional] 
**Var28** | Pointer to [**BlockchainConfig28**](BlockchainConfig28.md) |  | [optional] 
**Var29** | Pointer to [**BlockchainConfig29**](BlockchainConfig29.md) |  | [optional] 
**Var31** | Pointer to [**BlockchainConfig31**](BlockchainConfig31.md) |  | [optional] 
**Var32** | Pointer to [**ValidatorsSet**](ValidatorsSet.md) |  | [optional] 
**Var33** | Pointer to [**ValidatorsSet**](ValidatorsSet.md) |  | [optional] 
**Var34** | Pointer to [**ValidatorsSet**](ValidatorsSet.md) |  | [optional] 
**Var35** | Pointer to [**ValidatorsSet**](ValidatorsSet.md) |  | [optional] 
**Var36** | Pointer to [**ValidatorsSet**](ValidatorsSet.md) |  | [optional] 
**Var37** | Pointer to [**ValidatorsSet**](ValidatorsSet.md) |  | [optional] 
**Var40** | Pointer to [**BlockchainConfig40**](BlockchainConfig40.md) |  | [optional] 
**Var43** | Pointer to [**BlockchainConfig43**](BlockchainConfig43.md) |  | [optional] 
**Var44** | [**BlockchainConfig44**](BlockchainConfig44.md) |  | 
**Var71** | Pointer to [**BlockchainConfig71**](BlockchainConfig71.md) |  | [optional] 
**Var72** | Pointer to [**BlockchainConfig71**](BlockchainConfig71.md) |  | [optional] 
**Var73** | Pointer to [**BlockchainConfig71**](BlockchainConfig71.md) |  | [optional] 
**Var79** | Pointer to [**BlockchainConfig79**](BlockchainConfig79.md) |  | [optional] 
**Var81** | Pointer to [**BlockchainConfig79**](BlockchainConfig79.md) |  | [optional] 
**Var82** | Pointer to [**BlockchainConfig79**](BlockchainConfig79.md) |  | [optional] 

## Methods

### NewBlockchainConfig

`func NewBlockchainConfig(raw string, var0 string, var1 string, var2 string, var4 string, var44 BlockchainConfig44, ) *BlockchainConfig`

NewBlockchainConfig instantiates a new BlockchainConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockchainConfigWithDefaults

`func NewBlockchainConfigWithDefaults() *BlockchainConfig`

NewBlockchainConfigWithDefaults instantiates a new BlockchainConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRaw

`func (o *BlockchainConfig) GetRaw() string`

GetRaw returns the Raw field if non-nil, zero value otherwise.

### GetRawOk

`func (o *BlockchainConfig) GetRawOk() (*string, bool)`

GetRawOk returns a tuple with the Raw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaw

`func (o *BlockchainConfig) SetRaw(v string)`

SetRaw sets Raw field to given value.


### GetVar0

`func (o *BlockchainConfig) GetVar0() string`

GetVar0 returns the Var0 field if non-nil, zero value otherwise.

### GetVar0Ok

`func (o *BlockchainConfig) GetVar0Ok() (*string, bool)`

GetVar0Ok returns a tuple with the Var0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar0

`func (o *BlockchainConfig) SetVar0(v string)`

SetVar0 sets Var0 field to given value.


### GetVar1

`func (o *BlockchainConfig) GetVar1() string`

GetVar1 returns the Var1 field if non-nil, zero value otherwise.

### GetVar1Ok

`func (o *BlockchainConfig) GetVar1Ok() (*string, bool)`

GetVar1Ok returns a tuple with the Var1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar1

`func (o *BlockchainConfig) SetVar1(v string)`

SetVar1 sets Var1 field to given value.


### GetVar2

`func (o *BlockchainConfig) GetVar2() string`

GetVar2 returns the Var2 field if non-nil, zero value otherwise.

### GetVar2Ok

`func (o *BlockchainConfig) GetVar2Ok() (*string, bool)`

GetVar2Ok returns a tuple with the Var2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar2

`func (o *BlockchainConfig) SetVar2(v string)`

SetVar2 sets Var2 field to given value.


### GetVar3

`func (o *BlockchainConfig) GetVar3() string`

GetVar3 returns the Var3 field if non-nil, zero value otherwise.

### GetVar3Ok

`func (o *BlockchainConfig) GetVar3Ok() (*string, bool)`

GetVar3Ok returns a tuple with the Var3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar3

`func (o *BlockchainConfig) SetVar3(v string)`

SetVar3 sets Var3 field to given value.

### HasVar3

`func (o *BlockchainConfig) HasVar3() bool`

HasVar3 returns a boolean if a field has been set.

### GetVar4

`func (o *BlockchainConfig) GetVar4() string`

GetVar4 returns the Var4 field if non-nil, zero value otherwise.

### GetVar4Ok

`func (o *BlockchainConfig) GetVar4Ok() (*string, bool)`

GetVar4Ok returns a tuple with the Var4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar4

`func (o *BlockchainConfig) SetVar4(v string)`

SetVar4 sets Var4 field to given value.


### GetVar5

`func (o *BlockchainConfig) GetVar5() BlockchainConfig5`

GetVar5 returns the Var5 field if non-nil, zero value otherwise.

### GetVar5Ok

`func (o *BlockchainConfig) GetVar5Ok() (*BlockchainConfig5, bool)`

GetVar5Ok returns a tuple with the Var5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar5

`func (o *BlockchainConfig) SetVar5(v BlockchainConfig5)`

SetVar5 sets Var5 field to given value.

### HasVar5

`func (o *BlockchainConfig) HasVar5() bool`

HasVar5 returns a boolean if a field has been set.

### GetVar6

`func (o *BlockchainConfig) GetVar6() BlockchainConfig6`

GetVar6 returns the Var6 field if non-nil, zero value otherwise.

### GetVar6Ok

`func (o *BlockchainConfig) GetVar6Ok() (*BlockchainConfig6, bool)`

GetVar6Ok returns a tuple with the Var6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar6

`func (o *BlockchainConfig) SetVar6(v BlockchainConfig6)`

SetVar6 sets Var6 field to given value.

### HasVar6

`func (o *BlockchainConfig) HasVar6() bool`

HasVar6 returns a boolean if a field has been set.

### GetVar7

`func (o *BlockchainConfig) GetVar7() BlockchainConfig7`

GetVar7 returns the Var7 field if non-nil, zero value otherwise.

### GetVar7Ok

`func (o *BlockchainConfig) GetVar7Ok() (*BlockchainConfig7, bool)`

GetVar7Ok returns a tuple with the Var7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7

`func (o *BlockchainConfig) SetVar7(v BlockchainConfig7)`

SetVar7 sets Var7 field to given value.

### HasVar7

`func (o *BlockchainConfig) HasVar7() bool`

HasVar7 returns a boolean if a field has been set.

### GetVar8

`func (o *BlockchainConfig) GetVar8() BlockchainConfig8`

GetVar8 returns the Var8 field if non-nil, zero value otherwise.

### GetVar8Ok

`func (o *BlockchainConfig) GetVar8Ok() (*BlockchainConfig8, bool)`

GetVar8Ok returns a tuple with the Var8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar8

`func (o *BlockchainConfig) SetVar8(v BlockchainConfig8)`

SetVar8 sets Var8 field to given value.

### HasVar8

`func (o *BlockchainConfig) HasVar8() bool`

HasVar8 returns a boolean if a field has been set.

### GetVar9

`func (o *BlockchainConfig) GetVar9() BlockchainConfig9`

GetVar9 returns the Var9 field if non-nil, zero value otherwise.

### GetVar9Ok

`func (o *BlockchainConfig) GetVar9Ok() (*BlockchainConfig9, bool)`

GetVar9Ok returns a tuple with the Var9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar9

`func (o *BlockchainConfig) SetVar9(v BlockchainConfig9)`

SetVar9 sets Var9 field to given value.

### HasVar9

`func (o *BlockchainConfig) HasVar9() bool`

HasVar9 returns a boolean if a field has been set.

### GetVar10

`func (o *BlockchainConfig) GetVar10() BlockchainConfig10`

GetVar10 returns the Var10 field if non-nil, zero value otherwise.

### GetVar10Ok

`func (o *BlockchainConfig) GetVar10Ok() (*BlockchainConfig10, bool)`

GetVar10Ok returns a tuple with the Var10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar10

`func (o *BlockchainConfig) SetVar10(v BlockchainConfig10)`

SetVar10 sets Var10 field to given value.

### HasVar10

`func (o *BlockchainConfig) HasVar10() bool`

HasVar10 returns a boolean if a field has been set.

### GetVar11

`func (o *BlockchainConfig) GetVar11() BlockchainConfig11`

GetVar11 returns the Var11 field if non-nil, zero value otherwise.

### GetVar11Ok

`func (o *BlockchainConfig) GetVar11Ok() (*BlockchainConfig11, bool)`

GetVar11Ok returns a tuple with the Var11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar11

`func (o *BlockchainConfig) SetVar11(v BlockchainConfig11)`

SetVar11 sets Var11 field to given value.

### HasVar11

`func (o *BlockchainConfig) HasVar11() bool`

HasVar11 returns a boolean if a field has been set.

### GetVar12

`func (o *BlockchainConfig) GetVar12() BlockchainConfig12`

GetVar12 returns the Var12 field if non-nil, zero value otherwise.

### GetVar12Ok

`func (o *BlockchainConfig) GetVar12Ok() (*BlockchainConfig12, bool)`

GetVar12Ok returns a tuple with the Var12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar12

`func (o *BlockchainConfig) SetVar12(v BlockchainConfig12)`

SetVar12 sets Var12 field to given value.

### HasVar12

`func (o *BlockchainConfig) HasVar12() bool`

HasVar12 returns a boolean if a field has been set.

### GetVar13

`func (o *BlockchainConfig) GetVar13() BlockchainConfig13`

GetVar13 returns the Var13 field if non-nil, zero value otherwise.

### GetVar13Ok

`func (o *BlockchainConfig) GetVar13Ok() (*BlockchainConfig13, bool)`

GetVar13Ok returns a tuple with the Var13 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar13

`func (o *BlockchainConfig) SetVar13(v BlockchainConfig13)`

SetVar13 sets Var13 field to given value.

### HasVar13

`func (o *BlockchainConfig) HasVar13() bool`

HasVar13 returns a boolean if a field has been set.

### GetVar14

`func (o *BlockchainConfig) GetVar14() BlockchainConfig14`

GetVar14 returns the Var14 field if non-nil, zero value otherwise.

### GetVar14Ok

`func (o *BlockchainConfig) GetVar14Ok() (*BlockchainConfig14, bool)`

GetVar14Ok returns a tuple with the Var14 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar14

`func (o *BlockchainConfig) SetVar14(v BlockchainConfig14)`

SetVar14 sets Var14 field to given value.

### HasVar14

`func (o *BlockchainConfig) HasVar14() bool`

HasVar14 returns a boolean if a field has been set.

### GetVar15

`func (o *BlockchainConfig) GetVar15() BlockchainConfig15`

GetVar15 returns the Var15 field if non-nil, zero value otherwise.

### GetVar15Ok

`func (o *BlockchainConfig) GetVar15Ok() (*BlockchainConfig15, bool)`

GetVar15Ok returns a tuple with the Var15 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar15

`func (o *BlockchainConfig) SetVar15(v BlockchainConfig15)`

SetVar15 sets Var15 field to given value.

### HasVar15

`func (o *BlockchainConfig) HasVar15() bool`

HasVar15 returns a boolean if a field has been set.

### GetVar16

`func (o *BlockchainConfig) GetVar16() BlockchainConfig16`

GetVar16 returns the Var16 field if non-nil, zero value otherwise.

### GetVar16Ok

`func (o *BlockchainConfig) GetVar16Ok() (*BlockchainConfig16, bool)`

GetVar16Ok returns a tuple with the Var16 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar16

`func (o *BlockchainConfig) SetVar16(v BlockchainConfig16)`

SetVar16 sets Var16 field to given value.

### HasVar16

`func (o *BlockchainConfig) HasVar16() bool`

HasVar16 returns a boolean if a field has been set.

### GetVar17

`func (o *BlockchainConfig) GetVar17() BlockchainConfig17`

GetVar17 returns the Var17 field if non-nil, zero value otherwise.

### GetVar17Ok

`func (o *BlockchainConfig) GetVar17Ok() (*BlockchainConfig17, bool)`

GetVar17Ok returns a tuple with the Var17 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar17

`func (o *BlockchainConfig) SetVar17(v BlockchainConfig17)`

SetVar17 sets Var17 field to given value.

### HasVar17

`func (o *BlockchainConfig) HasVar17() bool`

HasVar17 returns a boolean if a field has been set.

### GetVar18

`func (o *BlockchainConfig) GetVar18() BlockchainConfig18`

GetVar18 returns the Var18 field if non-nil, zero value otherwise.

### GetVar18Ok

`func (o *BlockchainConfig) GetVar18Ok() (*BlockchainConfig18, bool)`

GetVar18Ok returns a tuple with the Var18 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar18

`func (o *BlockchainConfig) SetVar18(v BlockchainConfig18)`

SetVar18 sets Var18 field to given value.

### HasVar18

`func (o *BlockchainConfig) HasVar18() bool`

HasVar18 returns a boolean if a field has been set.

### GetVar20

`func (o *BlockchainConfig) GetVar20() BlockchainConfig20`

GetVar20 returns the Var20 field if non-nil, zero value otherwise.

### GetVar20Ok

`func (o *BlockchainConfig) GetVar20Ok() (*BlockchainConfig20, bool)`

GetVar20Ok returns a tuple with the Var20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar20

`func (o *BlockchainConfig) SetVar20(v BlockchainConfig20)`

SetVar20 sets Var20 field to given value.

### HasVar20

`func (o *BlockchainConfig) HasVar20() bool`

HasVar20 returns a boolean if a field has been set.

### GetVar21

`func (o *BlockchainConfig) GetVar21() BlockchainConfig21`

GetVar21 returns the Var21 field if non-nil, zero value otherwise.

### GetVar21Ok

`func (o *BlockchainConfig) GetVar21Ok() (*BlockchainConfig21, bool)`

GetVar21Ok returns a tuple with the Var21 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar21

`func (o *BlockchainConfig) SetVar21(v BlockchainConfig21)`

SetVar21 sets Var21 field to given value.

### HasVar21

`func (o *BlockchainConfig) HasVar21() bool`

HasVar21 returns a boolean if a field has been set.

### GetVar22

`func (o *BlockchainConfig) GetVar22() BlockchainConfig22`

GetVar22 returns the Var22 field if non-nil, zero value otherwise.

### GetVar22Ok

`func (o *BlockchainConfig) GetVar22Ok() (*BlockchainConfig22, bool)`

GetVar22Ok returns a tuple with the Var22 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar22

`func (o *BlockchainConfig) SetVar22(v BlockchainConfig22)`

SetVar22 sets Var22 field to given value.

### HasVar22

`func (o *BlockchainConfig) HasVar22() bool`

HasVar22 returns a boolean if a field has been set.

### GetVar23

`func (o *BlockchainConfig) GetVar23() BlockchainConfig23`

GetVar23 returns the Var23 field if non-nil, zero value otherwise.

### GetVar23Ok

`func (o *BlockchainConfig) GetVar23Ok() (*BlockchainConfig23, bool)`

GetVar23Ok returns a tuple with the Var23 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar23

`func (o *BlockchainConfig) SetVar23(v BlockchainConfig23)`

SetVar23 sets Var23 field to given value.

### HasVar23

`func (o *BlockchainConfig) HasVar23() bool`

HasVar23 returns a boolean if a field has been set.

### GetVar24

`func (o *BlockchainConfig) GetVar24() BlockchainConfig24`

GetVar24 returns the Var24 field if non-nil, zero value otherwise.

### GetVar24Ok

`func (o *BlockchainConfig) GetVar24Ok() (*BlockchainConfig24, bool)`

GetVar24Ok returns a tuple with the Var24 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar24

`func (o *BlockchainConfig) SetVar24(v BlockchainConfig24)`

SetVar24 sets Var24 field to given value.

### HasVar24

`func (o *BlockchainConfig) HasVar24() bool`

HasVar24 returns a boolean if a field has been set.

### GetVar25

`func (o *BlockchainConfig) GetVar25() BlockchainConfig25`

GetVar25 returns the Var25 field if non-nil, zero value otherwise.

### GetVar25Ok

`func (o *BlockchainConfig) GetVar25Ok() (*BlockchainConfig25, bool)`

GetVar25Ok returns a tuple with the Var25 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar25

`func (o *BlockchainConfig) SetVar25(v BlockchainConfig25)`

SetVar25 sets Var25 field to given value.

### HasVar25

`func (o *BlockchainConfig) HasVar25() bool`

HasVar25 returns a boolean if a field has been set.

### GetVar28

`func (o *BlockchainConfig) GetVar28() BlockchainConfig28`

GetVar28 returns the Var28 field if non-nil, zero value otherwise.

### GetVar28Ok

`func (o *BlockchainConfig) GetVar28Ok() (*BlockchainConfig28, bool)`

GetVar28Ok returns a tuple with the Var28 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar28

`func (o *BlockchainConfig) SetVar28(v BlockchainConfig28)`

SetVar28 sets Var28 field to given value.

### HasVar28

`func (o *BlockchainConfig) HasVar28() bool`

HasVar28 returns a boolean if a field has been set.

### GetVar29

`func (o *BlockchainConfig) GetVar29() BlockchainConfig29`

GetVar29 returns the Var29 field if non-nil, zero value otherwise.

### GetVar29Ok

`func (o *BlockchainConfig) GetVar29Ok() (*BlockchainConfig29, bool)`

GetVar29Ok returns a tuple with the Var29 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar29

`func (o *BlockchainConfig) SetVar29(v BlockchainConfig29)`

SetVar29 sets Var29 field to given value.

### HasVar29

`func (o *BlockchainConfig) HasVar29() bool`

HasVar29 returns a boolean if a field has been set.

### GetVar31

`func (o *BlockchainConfig) GetVar31() BlockchainConfig31`

GetVar31 returns the Var31 field if non-nil, zero value otherwise.

### GetVar31Ok

`func (o *BlockchainConfig) GetVar31Ok() (*BlockchainConfig31, bool)`

GetVar31Ok returns a tuple with the Var31 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar31

`func (o *BlockchainConfig) SetVar31(v BlockchainConfig31)`

SetVar31 sets Var31 field to given value.

### HasVar31

`func (o *BlockchainConfig) HasVar31() bool`

HasVar31 returns a boolean if a field has been set.

### GetVar32

`func (o *BlockchainConfig) GetVar32() ValidatorsSet`

GetVar32 returns the Var32 field if non-nil, zero value otherwise.

### GetVar32Ok

`func (o *BlockchainConfig) GetVar32Ok() (*ValidatorsSet, bool)`

GetVar32Ok returns a tuple with the Var32 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar32

`func (o *BlockchainConfig) SetVar32(v ValidatorsSet)`

SetVar32 sets Var32 field to given value.

### HasVar32

`func (o *BlockchainConfig) HasVar32() bool`

HasVar32 returns a boolean if a field has been set.

### GetVar33

`func (o *BlockchainConfig) GetVar33() ValidatorsSet`

GetVar33 returns the Var33 field if non-nil, zero value otherwise.

### GetVar33Ok

`func (o *BlockchainConfig) GetVar33Ok() (*ValidatorsSet, bool)`

GetVar33Ok returns a tuple with the Var33 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar33

`func (o *BlockchainConfig) SetVar33(v ValidatorsSet)`

SetVar33 sets Var33 field to given value.

### HasVar33

`func (o *BlockchainConfig) HasVar33() bool`

HasVar33 returns a boolean if a field has been set.

### GetVar34

`func (o *BlockchainConfig) GetVar34() ValidatorsSet`

GetVar34 returns the Var34 field if non-nil, zero value otherwise.

### GetVar34Ok

`func (o *BlockchainConfig) GetVar34Ok() (*ValidatorsSet, bool)`

GetVar34Ok returns a tuple with the Var34 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar34

`func (o *BlockchainConfig) SetVar34(v ValidatorsSet)`

SetVar34 sets Var34 field to given value.

### HasVar34

`func (o *BlockchainConfig) HasVar34() bool`

HasVar34 returns a boolean if a field has been set.

### GetVar35

`func (o *BlockchainConfig) GetVar35() ValidatorsSet`

GetVar35 returns the Var35 field if non-nil, zero value otherwise.

### GetVar35Ok

`func (o *BlockchainConfig) GetVar35Ok() (*ValidatorsSet, bool)`

GetVar35Ok returns a tuple with the Var35 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar35

`func (o *BlockchainConfig) SetVar35(v ValidatorsSet)`

SetVar35 sets Var35 field to given value.

### HasVar35

`func (o *BlockchainConfig) HasVar35() bool`

HasVar35 returns a boolean if a field has been set.

### GetVar36

`func (o *BlockchainConfig) GetVar36() ValidatorsSet`

GetVar36 returns the Var36 field if non-nil, zero value otherwise.

### GetVar36Ok

`func (o *BlockchainConfig) GetVar36Ok() (*ValidatorsSet, bool)`

GetVar36Ok returns a tuple with the Var36 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar36

`func (o *BlockchainConfig) SetVar36(v ValidatorsSet)`

SetVar36 sets Var36 field to given value.

### HasVar36

`func (o *BlockchainConfig) HasVar36() bool`

HasVar36 returns a boolean if a field has been set.

### GetVar37

`func (o *BlockchainConfig) GetVar37() ValidatorsSet`

GetVar37 returns the Var37 field if non-nil, zero value otherwise.

### GetVar37Ok

`func (o *BlockchainConfig) GetVar37Ok() (*ValidatorsSet, bool)`

GetVar37Ok returns a tuple with the Var37 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar37

`func (o *BlockchainConfig) SetVar37(v ValidatorsSet)`

SetVar37 sets Var37 field to given value.

### HasVar37

`func (o *BlockchainConfig) HasVar37() bool`

HasVar37 returns a boolean if a field has been set.

### GetVar40

`func (o *BlockchainConfig) GetVar40() BlockchainConfig40`

GetVar40 returns the Var40 field if non-nil, zero value otherwise.

### GetVar40Ok

`func (o *BlockchainConfig) GetVar40Ok() (*BlockchainConfig40, bool)`

GetVar40Ok returns a tuple with the Var40 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar40

`func (o *BlockchainConfig) SetVar40(v BlockchainConfig40)`

SetVar40 sets Var40 field to given value.

### HasVar40

`func (o *BlockchainConfig) HasVar40() bool`

HasVar40 returns a boolean if a field has been set.

### GetVar43

`func (o *BlockchainConfig) GetVar43() BlockchainConfig43`

GetVar43 returns the Var43 field if non-nil, zero value otherwise.

### GetVar43Ok

`func (o *BlockchainConfig) GetVar43Ok() (*BlockchainConfig43, bool)`

GetVar43Ok returns a tuple with the Var43 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar43

`func (o *BlockchainConfig) SetVar43(v BlockchainConfig43)`

SetVar43 sets Var43 field to given value.

### HasVar43

`func (o *BlockchainConfig) HasVar43() bool`

HasVar43 returns a boolean if a field has been set.

### GetVar44

`func (o *BlockchainConfig) GetVar44() BlockchainConfig44`

GetVar44 returns the Var44 field if non-nil, zero value otherwise.

### GetVar44Ok

`func (o *BlockchainConfig) GetVar44Ok() (*BlockchainConfig44, bool)`

GetVar44Ok returns a tuple with the Var44 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar44

`func (o *BlockchainConfig) SetVar44(v BlockchainConfig44)`

SetVar44 sets Var44 field to given value.


### GetVar71

`func (o *BlockchainConfig) GetVar71() BlockchainConfig71`

GetVar71 returns the Var71 field if non-nil, zero value otherwise.

### GetVar71Ok

`func (o *BlockchainConfig) GetVar71Ok() (*BlockchainConfig71, bool)`

GetVar71Ok returns a tuple with the Var71 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar71

`func (o *BlockchainConfig) SetVar71(v BlockchainConfig71)`

SetVar71 sets Var71 field to given value.

### HasVar71

`func (o *BlockchainConfig) HasVar71() bool`

HasVar71 returns a boolean if a field has been set.

### GetVar72

`func (o *BlockchainConfig) GetVar72() BlockchainConfig71`

GetVar72 returns the Var72 field if non-nil, zero value otherwise.

### GetVar72Ok

`func (o *BlockchainConfig) GetVar72Ok() (*BlockchainConfig71, bool)`

GetVar72Ok returns a tuple with the Var72 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar72

`func (o *BlockchainConfig) SetVar72(v BlockchainConfig71)`

SetVar72 sets Var72 field to given value.

### HasVar72

`func (o *BlockchainConfig) HasVar72() bool`

HasVar72 returns a boolean if a field has been set.

### GetVar73

`func (o *BlockchainConfig) GetVar73() BlockchainConfig71`

GetVar73 returns the Var73 field if non-nil, zero value otherwise.

### GetVar73Ok

`func (o *BlockchainConfig) GetVar73Ok() (*BlockchainConfig71, bool)`

GetVar73Ok returns a tuple with the Var73 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar73

`func (o *BlockchainConfig) SetVar73(v BlockchainConfig71)`

SetVar73 sets Var73 field to given value.

### HasVar73

`func (o *BlockchainConfig) HasVar73() bool`

HasVar73 returns a boolean if a field has been set.

### GetVar79

`func (o *BlockchainConfig) GetVar79() BlockchainConfig79`

GetVar79 returns the Var79 field if non-nil, zero value otherwise.

### GetVar79Ok

`func (o *BlockchainConfig) GetVar79Ok() (*BlockchainConfig79, bool)`

GetVar79Ok returns a tuple with the Var79 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar79

`func (o *BlockchainConfig) SetVar79(v BlockchainConfig79)`

SetVar79 sets Var79 field to given value.

### HasVar79

`func (o *BlockchainConfig) HasVar79() bool`

HasVar79 returns a boolean if a field has been set.

### GetVar81

`func (o *BlockchainConfig) GetVar81() BlockchainConfig79`

GetVar81 returns the Var81 field if non-nil, zero value otherwise.

### GetVar81Ok

`func (o *BlockchainConfig) GetVar81Ok() (*BlockchainConfig79, bool)`

GetVar81Ok returns a tuple with the Var81 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar81

`func (o *BlockchainConfig) SetVar81(v BlockchainConfig79)`

SetVar81 sets Var81 field to given value.

### HasVar81

`func (o *BlockchainConfig) HasVar81() bool`

HasVar81 returns a boolean if a field has been set.

### GetVar82

`func (o *BlockchainConfig) GetVar82() BlockchainConfig79`

GetVar82 returns the Var82 field if non-nil, zero value otherwise.

### GetVar82Ok

`func (o *BlockchainConfig) GetVar82Ok() (*BlockchainConfig79, bool)`

GetVar82Ok returns a tuple with the Var82 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar82

`func (o *BlockchainConfig) SetVar82(v BlockchainConfig79)`

SetVar82 sets Var82 field to given value.

### HasVar82

`func (o *BlockchainConfig) HasVar82() bool`

HasVar82 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


