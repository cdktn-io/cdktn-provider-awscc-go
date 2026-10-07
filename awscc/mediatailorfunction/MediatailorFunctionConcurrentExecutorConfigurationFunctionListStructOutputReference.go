// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediatailorfunction

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/mediatailorfunction/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MediatailorFunctionConcurrentExecutorConfigurationFunctionListStructOutputReference interface {
	cdktn.ComplexObject
	Alias() *string
	SetAlias(val *string)
	AliasInput() *string
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
	FunctionId() *string
	SetFunctionId(val *string)
	FunctionIdInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	RunCondition() *string
	SetRunCondition(val *string)
	RunConditionInput() *string
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
	ResetAlias()
	ResetFunctionId()
	ResetRunCondition()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MediatailorFunctionConcurrentExecutorConfigurationFunctionListStructOutputReference
type jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationFunctionListStructOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationFunctionListStructOutputReference) Alias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationFunctionListStructOutputReference) AliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationFunctionListStructOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationFunctionListStructOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationFunctionListStructOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationFunctionListStructOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationFunctionListStructOutputReference) FunctionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationFunctionListStructOutputReference) FunctionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationFunctionListStructOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationFunctionListStructOutputReference) RunCondition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationFunctionListStructOutputReference) RunConditionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationFunctionListStructOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationFunctionListStructOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMediatailorFunctionConcurrentExecutorConfigurationFunctionListStructOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) MediatailorFunctionConcurrentExecutorConfigurationFunctionListStructOutputReference {
	_init_.Initialize()

	if err := validateNewMediatailorFunctionConcurrentExecutorConfigurationFunctionListStructOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationFunctionListStructOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.mediatailorFunction.MediatailorFunctionConcurrentExecutorConfigurationFunctionListStructOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewMediatailorFunctionConcurrentExecutorConfigurationFunctionListStructOutputReference_Override(m MediatailorFunctionConcurrentExecutorConfigurationFunctionListStructOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.mediatailorFunction.MediatailorFunctionConcurrentExecutorConfigurationFunctionListStructOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationFunctionListStructOutputReference)SetAlias(val *string) {
	if err := j.validateSetAliasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alias",
		val,
	)
}

func (j *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationFunctionListStructOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationFunctionListStructOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationFunctionListStructOutputReference)SetFunctionId(val *string) {
	if err := j.validateSetFunctionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"functionId",
		val,
	)
}

func (j *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationFunctionListStructOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationFunctionListStructOutputReference)SetRunCondition(val *string) {
	if err := j.validateSetRunConditionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runCondition",
		val,
	)
}

func (j *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationFunctionListStructOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationFunctionListStructOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationFunctionListStructOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationFunctionListStructOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationFunctionListStructOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationFunctionListStructOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationFunctionListStructOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationFunctionListStructOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationFunctionListStructOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationFunctionListStructOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationFunctionListStructOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationFunctionListStructOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationFunctionListStructOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationFunctionListStructOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationFunctionListStructOutputReference) ResetAlias() {
	_jsii_.InvokeVoid(
		m,
		"resetAlias",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationFunctionListStructOutputReference) ResetFunctionId() {
	_jsii_.InvokeVoid(
		m,
		"resetFunctionId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationFunctionListStructOutputReference) ResetRunCondition() {
	_jsii_.InvokeVoid(
		m,
		"resetRunCondition",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationFunctionListStructOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := m.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationFunctionListStructOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

