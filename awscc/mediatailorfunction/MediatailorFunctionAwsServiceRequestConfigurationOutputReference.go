// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediatailorfunction

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/mediatailorfunction/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MediatailorFunctionAwsServiceRequestConfigurationOutputReference interface {
	cdktn.ComplexObject
	Body() *string
	SetBody(val *string)
	BodyInput() *string
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
	Headers() *map[string]*string
	SetHeaders(val *map[string]*string)
	HeadersInput() *map[string]*string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MethodType() *string
	SetMethodType(val *string)
	MethodTypeInput() *string
	Output() *map[string]*string
	SetOutput(val *map[string]*string)
	OutputInput() *map[string]*string
	RequestTimeoutMilliseconds() *float64
	SetRequestTimeoutMilliseconds(val *float64)
	RequestTimeoutMillisecondsInput() *float64
	Runtime() *string
	SetRuntime(val *string)
	RuntimeInput() *string
	TargetRegion() *string
	SetTargetRegion(val *string)
	TargetRegionInput() *string
	TargetService() *string
	SetTargetService(val *string)
	TargetServiceInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Url() *string
	SetUrl(val *string)
	UrlInput() *string
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
	ResetBody()
	ResetHeaders()
	ResetMethodType()
	ResetOutput()
	ResetRequestTimeoutMilliseconds()
	ResetRuntime()
	ResetTargetRegion()
	ResetTargetService()
	ResetUrl()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MediatailorFunctionAwsServiceRequestConfigurationOutputReference
type jsiiProxy_MediatailorFunctionAwsServiceRequestConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_MediatailorFunctionAwsServiceRequestConfigurationOutputReference) Body() *string {
	var returns *string
	_jsii_.Get(
		j,
		"body",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorFunctionAwsServiceRequestConfigurationOutputReference) BodyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorFunctionAwsServiceRequestConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorFunctionAwsServiceRequestConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorFunctionAwsServiceRequestConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorFunctionAwsServiceRequestConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorFunctionAwsServiceRequestConfigurationOutputReference) Headers() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"headers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorFunctionAwsServiceRequestConfigurationOutputReference) HeadersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"headersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorFunctionAwsServiceRequestConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorFunctionAwsServiceRequestConfigurationOutputReference) MethodType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"methodType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorFunctionAwsServiceRequestConfigurationOutputReference) MethodTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"methodTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorFunctionAwsServiceRequestConfigurationOutputReference) Output() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"output",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorFunctionAwsServiceRequestConfigurationOutputReference) OutputInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"outputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorFunctionAwsServiceRequestConfigurationOutputReference) RequestTimeoutMilliseconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requestTimeoutMilliseconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorFunctionAwsServiceRequestConfigurationOutputReference) RequestTimeoutMillisecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requestTimeoutMillisecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorFunctionAwsServiceRequestConfigurationOutputReference) Runtime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorFunctionAwsServiceRequestConfigurationOutputReference) RuntimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorFunctionAwsServiceRequestConfigurationOutputReference) TargetRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorFunctionAwsServiceRequestConfigurationOutputReference) TargetRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorFunctionAwsServiceRequestConfigurationOutputReference) TargetService() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorFunctionAwsServiceRequestConfigurationOutputReference) TargetServiceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorFunctionAwsServiceRequestConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorFunctionAwsServiceRequestConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorFunctionAwsServiceRequestConfigurationOutputReference) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorFunctionAwsServiceRequestConfigurationOutputReference) UrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlInput",
		&returns,
	)
	return returns
}


func NewMediatailorFunctionAwsServiceRequestConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) MediatailorFunctionAwsServiceRequestConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewMediatailorFunctionAwsServiceRequestConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MediatailorFunctionAwsServiceRequestConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.mediatailorFunction.MediatailorFunctionAwsServiceRequestConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMediatailorFunctionAwsServiceRequestConfigurationOutputReference_Override(m MediatailorFunctionAwsServiceRequestConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.mediatailorFunction.MediatailorFunctionAwsServiceRequestConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MediatailorFunctionAwsServiceRequestConfigurationOutputReference)SetBody(val *string) {
	if err := j.validateSetBodyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"body",
		val,
	)
}

func (j *jsiiProxy_MediatailorFunctionAwsServiceRequestConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediatailorFunctionAwsServiceRequestConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediatailorFunctionAwsServiceRequestConfigurationOutputReference)SetHeaders(val *map[string]*string) {
	if err := j.validateSetHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"headers",
		val,
	)
}

func (j *jsiiProxy_MediatailorFunctionAwsServiceRequestConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediatailorFunctionAwsServiceRequestConfigurationOutputReference)SetMethodType(val *string) {
	if err := j.validateSetMethodTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"methodType",
		val,
	)
}

func (j *jsiiProxy_MediatailorFunctionAwsServiceRequestConfigurationOutputReference)SetOutput(val *map[string]*string) {
	if err := j.validateSetOutputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"output",
		val,
	)
}

func (j *jsiiProxy_MediatailorFunctionAwsServiceRequestConfigurationOutputReference)SetRequestTimeoutMilliseconds(val *float64) {
	if err := j.validateSetRequestTimeoutMillisecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestTimeoutMilliseconds",
		val,
	)
}

func (j *jsiiProxy_MediatailorFunctionAwsServiceRequestConfigurationOutputReference)SetRuntime(val *string) {
	if err := j.validateSetRuntimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtime",
		val,
	)
}

func (j *jsiiProxy_MediatailorFunctionAwsServiceRequestConfigurationOutputReference)SetTargetRegion(val *string) {
	if err := j.validateSetTargetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetRegion",
		val,
	)
}

func (j *jsiiProxy_MediatailorFunctionAwsServiceRequestConfigurationOutputReference)SetTargetService(val *string) {
	if err := j.validateSetTargetServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetService",
		val,
	)
}

func (j *jsiiProxy_MediatailorFunctionAwsServiceRequestConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediatailorFunctionAwsServiceRequestConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MediatailorFunctionAwsServiceRequestConfigurationOutputReference)SetUrl(val *string) {
	if err := j.validateSetUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

func (m *jsiiProxy_MediatailorFunctionAwsServiceRequestConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediatailorFunctionAwsServiceRequestConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MediatailorFunctionAwsServiceRequestConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MediatailorFunctionAwsServiceRequestConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MediatailorFunctionAwsServiceRequestConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MediatailorFunctionAwsServiceRequestConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MediatailorFunctionAwsServiceRequestConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MediatailorFunctionAwsServiceRequestConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MediatailorFunctionAwsServiceRequestConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MediatailorFunctionAwsServiceRequestConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MediatailorFunctionAwsServiceRequestConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediatailorFunctionAwsServiceRequestConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MediatailorFunctionAwsServiceRequestConfigurationOutputReference) ResetBody() {
	_jsii_.InvokeVoid(
		m,
		"resetBody",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediatailorFunctionAwsServiceRequestConfigurationOutputReference) ResetHeaders() {
	_jsii_.InvokeVoid(
		m,
		"resetHeaders",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediatailorFunctionAwsServiceRequestConfigurationOutputReference) ResetMethodType() {
	_jsii_.InvokeVoid(
		m,
		"resetMethodType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediatailorFunctionAwsServiceRequestConfigurationOutputReference) ResetOutput() {
	_jsii_.InvokeVoid(
		m,
		"resetOutput",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediatailorFunctionAwsServiceRequestConfigurationOutputReference) ResetRequestTimeoutMilliseconds() {
	_jsii_.InvokeVoid(
		m,
		"resetRequestTimeoutMilliseconds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediatailorFunctionAwsServiceRequestConfigurationOutputReference) ResetRuntime() {
	_jsii_.InvokeVoid(
		m,
		"resetRuntime",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediatailorFunctionAwsServiceRequestConfigurationOutputReference) ResetTargetRegion() {
	_jsii_.InvokeVoid(
		m,
		"resetTargetRegion",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediatailorFunctionAwsServiceRequestConfigurationOutputReference) ResetTargetService() {
	_jsii_.InvokeVoid(
		m,
		"resetTargetService",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediatailorFunctionAwsServiceRequestConfigurationOutputReference) ResetUrl() {
	_jsii_.InvokeVoid(
		m,
		"resetUrl",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediatailorFunctionAwsServiceRequestConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (m *jsiiProxy_MediatailorFunctionAwsServiceRequestConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

