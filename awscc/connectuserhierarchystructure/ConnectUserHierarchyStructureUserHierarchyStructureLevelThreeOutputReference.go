// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectuserhierarchystructure

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/connectuserhierarchystructure/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ConnectUserHierarchyStructureUserHierarchyStructureLevelThreeOutputReference interface {
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
	HierarchyLevelArn() *string
	SetHierarchyLevelArn(val *string)
	HierarchyLevelArnInput() *string
	HierarchyLevelId() *string
	SetHierarchyLevelId(val *string)
	HierarchyLevelIdInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
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
	ResetHierarchyLevelArn()
	ResetHierarchyLevelId()
	ResetName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ConnectUserHierarchyStructureUserHierarchyStructureLevelThreeOutputReference
type jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureLevelThreeOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureLevelThreeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureLevelThreeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureLevelThreeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureLevelThreeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureLevelThreeOutputReference) HierarchyLevelArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hierarchyLevelArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureLevelThreeOutputReference) HierarchyLevelArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hierarchyLevelArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureLevelThreeOutputReference) HierarchyLevelId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hierarchyLevelId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureLevelThreeOutputReference) HierarchyLevelIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hierarchyLevelIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureLevelThreeOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureLevelThreeOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureLevelThreeOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureLevelThreeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureLevelThreeOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewConnectUserHierarchyStructureUserHierarchyStructureLevelThreeOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ConnectUserHierarchyStructureUserHierarchyStructureLevelThreeOutputReference {
	_init_.Initialize()

	if err := validateNewConnectUserHierarchyStructureUserHierarchyStructureLevelThreeOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureLevelThreeOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.connectUserHierarchyStructure.ConnectUserHierarchyStructureUserHierarchyStructureLevelThreeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewConnectUserHierarchyStructureUserHierarchyStructureLevelThreeOutputReference_Override(c ConnectUserHierarchyStructureUserHierarchyStructureLevelThreeOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.connectUserHierarchyStructure.ConnectUserHierarchyStructureUserHierarchyStructureLevelThreeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureLevelThreeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureLevelThreeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureLevelThreeOutputReference)SetHierarchyLevelArn(val *string) {
	if err := j.validateSetHierarchyLevelArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hierarchyLevelArn",
		val,
	)
}

func (j *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureLevelThreeOutputReference)SetHierarchyLevelId(val *string) {
	if err := j.validateSetHierarchyLevelIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hierarchyLevelId",
		val,
	)
}

func (j *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureLevelThreeOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureLevelThreeOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureLevelThreeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureLevelThreeOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureLevelThreeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureLevelThreeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureLevelThreeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureLevelThreeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureLevelThreeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureLevelThreeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureLevelThreeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureLevelThreeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureLevelThreeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureLevelThreeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureLevelThreeOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureLevelThreeOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureLevelThreeOutputReference) ResetHierarchyLevelArn() {
	_jsii_.InvokeVoid(
		c,
		"resetHierarchyLevelArn",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureLevelThreeOutputReference) ResetHierarchyLevelId() {
	_jsii_.InvokeVoid(
		c,
		"resetHierarchyLevelId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureLevelThreeOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		c,
		"resetName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureLevelThreeOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_ConnectUserHierarchyStructureUserHierarchyStructureLevelThreeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

