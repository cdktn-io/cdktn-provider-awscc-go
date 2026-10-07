// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediatailorfunction

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/mediatailorfunction/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MediatailorFunctionConcurrentExecutorConfigurationOutputReference interface {
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
	FunctionList() MediatailorFunctionConcurrentExecutorConfigurationFunctionListStructList
	FunctionListInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MaxConcurrency() *float64
	SetMaxConcurrency(val *float64)
	MaxConcurrencyInput() *float64
	Output() *map[string]*string
	SetOutput(val *map[string]*string)
	OutputInput() *map[string]*string
	Runtime() *string
	SetRuntime(val *string)
	RuntimeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TimeoutMilliseconds() *float64
	SetTimeoutMilliseconds(val *float64)
	TimeoutMillisecondsInput() *float64
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
	PutFunctionList(value interface{})
	ResetFunctionList()
	ResetMaxConcurrency()
	ResetOutput()
	ResetRuntime()
	ResetTimeoutMilliseconds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MediatailorFunctionConcurrentExecutorConfigurationOutputReference
type jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationOutputReference) FunctionList() MediatailorFunctionConcurrentExecutorConfigurationFunctionListStructList {
	var returns MediatailorFunctionConcurrentExecutorConfigurationFunctionListStructList
	_jsii_.Get(
		j,
		"functionList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationOutputReference) FunctionListInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"functionListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationOutputReference) MaxConcurrency() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationOutputReference) MaxConcurrencyInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationOutputReference) Output() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"output",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationOutputReference) OutputInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"outputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationOutputReference) Runtime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationOutputReference) RuntimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationOutputReference) TimeoutMilliseconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutMilliseconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationOutputReference) TimeoutMillisecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutMillisecondsInput",
		&returns,
	)
	return returns
}


func NewMediatailorFunctionConcurrentExecutorConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) MediatailorFunctionConcurrentExecutorConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewMediatailorFunctionConcurrentExecutorConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.mediatailorFunction.MediatailorFunctionConcurrentExecutorConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMediatailorFunctionConcurrentExecutorConfigurationOutputReference_Override(m MediatailorFunctionConcurrentExecutorConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.mediatailorFunction.MediatailorFunctionConcurrentExecutorConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationOutputReference)SetMaxConcurrency(val *float64) {
	if err := j.validateSetMaxConcurrencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxConcurrency",
		val,
	)
}

func (j *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationOutputReference)SetOutput(val *map[string]*string) {
	if err := j.validateSetOutputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"output",
		val,
	)
}

func (j *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationOutputReference)SetRuntime(val *string) {
	if err := j.validateSetRuntimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtime",
		val,
	)
}

func (j *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationOutputReference)SetTimeoutMilliseconds(val *float64) {
	if err := j.validateSetTimeoutMillisecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeoutMilliseconds",
		val,
	)
}

func (m *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationOutputReference) PutFunctionList(value interface{}) {
	if err := m.validatePutFunctionListParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putFunctionList",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationOutputReference) ResetFunctionList() {
	_jsii_.InvokeVoid(
		m,
		"resetFunctionList",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationOutputReference) ResetMaxConcurrency() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxConcurrency",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationOutputReference) ResetOutput() {
	_jsii_.InvokeVoid(
		m,
		"resetOutput",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationOutputReference) ResetRuntime() {
	_jsii_.InvokeVoid(
		m,
		"resetRuntime",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationOutputReference) ResetTimeoutMilliseconds() {
	_jsii_.InvokeVoid(
		m,
		"resetTimeoutMilliseconds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (m *jsiiProxy_MediatailorFunctionConcurrentExecutorConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

