// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectuserhierarchystructure

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/connectuserhierarchystructure/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ConnectUserHierarchyStructureUserHierarchyStructureOutputReference interface {
	cdktn.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LevelFive() ConnectUserHierarchyStructureUserHierarchyStructureLevelFiveOutputReference
	LevelFiveInput() interface{}
	LevelFour() ConnectUserHierarchyStructureUserHierarchyStructureLevelFourOutputReference
	LevelFourInput() interface{}
	LevelOne() ConnectUserHierarchyStructureUserHierarchyStructureLevelOneOutputReference
	LevelOneInput() interface{}
	LevelThree() ConnectUserHierarchyStructureUserHierarchyStructureLevelThreeOutputReference
	LevelThreeInput() interface{}
	LevelTwo() ConnectUserHierarchyStructureUserHierarchyStructureLevelTwoOutputReference
	LevelTwoInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	PutLevelFive(value *ConnectUserHierarchyStructureUserHierarchyStructureLevelFive)
	PutLevelFour(value *ConnectUserHierarchyStructureUserHierarchyStructureLevelFour)
	PutLevelOne(value *ConnectUserHierarchyStructureUserHierarchyStructureLevelOne)
	PutLevelThree(value *ConnectUserHierarchyStructureUserHierarchyStructureLevelThree)
	PutLevelTwo(value *ConnectUserHierarchyStructureUserHierarchyStructureLevelTwo)
	ResetLevelFive()
	ResetLevelFour()
	ResetLevelOne()
	ResetLevelThree()
	ResetLevelTwo()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ConnectUserHierarchyStructureUserHierarchyStructureOutputReference
type jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureOutputReference) LevelFive() ConnectUserHierarchyStructureUserHierarchyStructureLevelFiveOutputReference {
	var returns ConnectUserHierarchyStructureUserHierarchyStructureLevelFiveOutputReference
	_jsii_.Get(
		j,
		"levelFive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureOutputReference) LevelFiveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"levelFiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureOutputReference) LevelFour() ConnectUserHierarchyStructureUserHierarchyStructureLevelFourOutputReference {
	var returns ConnectUserHierarchyStructureUserHierarchyStructureLevelFourOutputReference
	_jsii_.Get(
		j,
		"levelFour",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureOutputReference) LevelFourInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"levelFourInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureOutputReference) LevelOne() ConnectUserHierarchyStructureUserHierarchyStructureLevelOneOutputReference {
	var returns ConnectUserHierarchyStructureUserHierarchyStructureLevelOneOutputReference
	_jsii_.Get(
		j,
		"levelOne",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureOutputReference) LevelOneInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"levelOneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureOutputReference) LevelThree() ConnectUserHierarchyStructureUserHierarchyStructureLevelThreeOutputReference {
	var returns ConnectUserHierarchyStructureUserHierarchyStructureLevelThreeOutputReference
	_jsii_.Get(
		j,
		"levelThree",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureOutputReference) LevelThreeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"levelThreeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureOutputReference) LevelTwo() ConnectUserHierarchyStructureUserHierarchyStructureLevelTwoOutputReference {
	var returns ConnectUserHierarchyStructureUserHierarchyStructureLevelTwoOutputReference
	_jsii_.Get(
		j,
		"levelTwo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureOutputReference) LevelTwoInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"levelTwoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewConnectUserHierarchyStructureUserHierarchyStructureOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ConnectUserHierarchyStructureUserHierarchyStructureOutputReference {
	_init_.Initialize()

	if err := validateNewConnectUserHierarchyStructureUserHierarchyStructureOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.connectUserHierarchyStructure.ConnectUserHierarchyStructureUserHierarchyStructureOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewConnectUserHierarchyStructureUserHierarchyStructureOutputReference_Override(c ConnectUserHierarchyStructureUserHierarchyStructureOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.connectUserHierarchyStructure.ConnectUserHierarchyStructureUserHierarchyStructureOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureOutputReference) PutLevelFive(value *ConnectUserHierarchyStructureUserHierarchyStructureLevelFive) {
	if err := c.validatePutLevelFiveParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putLevelFive",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureOutputReference) PutLevelFour(value *ConnectUserHierarchyStructureUserHierarchyStructureLevelFour) {
	if err := c.validatePutLevelFourParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putLevelFour",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureOutputReference) PutLevelOne(value *ConnectUserHierarchyStructureUserHierarchyStructureLevelOne) {
	if err := c.validatePutLevelOneParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putLevelOne",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureOutputReference) PutLevelThree(value *ConnectUserHierarchyStructureUserHierarchyStructureLevelThree) {
	if err := c.validatePutLevelThreeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putLevelThree",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureOutputReference) PutLevelTwo(value *ConnectUserHierarchyStructureUserHierarchyStructureLevelTwo) {
	if err := c.validatePutLevelTwoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putLevelTwo",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureOutputReference) ResetLevelFive() {
	_jsii_.InvokeVoid(
		c,
		"resetLevelFive",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureOutputReference) ResetLevelFour() {
	_jsii_.InvokeVoid(
		c,
		"resetLevelFour",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureOutputReference) ResetLevelOne() {
	_jsii_.InvokeVoid(
		c,
		"resetLevelOne",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureOutputReference) ResetLevelThree() {
	_jsii_.InvokeVoid(
		c,
		"resetLevelThree",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureOutputReference) ResetLevelTwo() {
	_jsii_.InvokeVoid(
		c,
		"resetLevelTwo",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := c.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

