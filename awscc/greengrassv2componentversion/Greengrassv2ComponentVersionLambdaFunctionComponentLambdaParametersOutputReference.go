// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package greengrassv2componentversion

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/greengrassv2componentversion/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference interface {
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
	EnvironmentVariables() *map[string]*string
	SetEnvironmentVariables(val *map[string]*string)
	EnvironmentVariablesInput() *map[string]*string
	EventSources() Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersEventSourcesList
	EventSourcesInput() interface{}
	ExecArgs() *[]*string
	SetExecArgs(val *[]*string)
	ExecArgsInput() *[]*string
	// Experimental.
	Fqn() *string
	InputPayloadEncodingType() *string
	SetInputPayloadEncodingType(val *string)
	InputPayloadEncodingTypeInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LinuxProcessParams() Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersLinuxProcessParamsOutputReference
	LinuxProcessParamsInput() interface{}
	MaxIdleTimeInSeconds() *float64
	SetMaxIdleTimeInSeconds(val *float64)
	MaxIdleTimeInSecondsInput() *float64
	MaxInstancesCount() *float64
	SetMaxInstancesCount(val *float64)
	MaxInstancesCountInput() *float64
	MaxQueueSize() *float64
	SetMaxQueueSize(val *float64)
	MaxQueueSizeInput() *float64
	Pinned() interface{}
	SetPinned(val interface{})
	PinnedInput() interface{}
	StatusTimeoutInSeconds() *float64
	SetStatusTimeoutInSeconds(val *float64)
	StatusTimeoutInSecondsInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TimeoutInSeconds() *float64
	SetTimeoutInSeconds(val *float64)
	TimeoutInSecondsInput() *float64
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
	PutEventSources(value interface{})
	PutLinuxProcessParams(value *Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersLinuxProcessParams)
	ResetEnvironmentVariables()
	ResetEventSources()
	ResetExecArgs()
	ResetInputPayloadEncodingType()
	ResetLinuxProcessParams()
	ResetMaxIdleTimeInSeconds()
	ResetMaxInstancesCount()
	ResetMaxQueueSize()
	ResetPinned()
	ResetStatusTimeoutInSeconds()
	ResetTimeoutInSeconds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference
type jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference) EnvironmentVariables() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environmentVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference) EnvironmentVariablesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environmentVariablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference) EventSources() Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersEventSourcesList {
	var returns Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersEventSourcesList
	_jsii_.Get(
		j,
		"eventSources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference) EventSourcesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"eventSourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference) ExecArgs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"execArgs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference) ExecArgsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"execArgsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference) InputPayloadEncodingType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputPayloadEncodingType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference) InputPayloadEncodingTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputPayloadEncodingTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference) LinuxProcessParams() Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersLinuxProcessParamsOutputReference {
	var returns Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersLinuxProcessParamsOutputReference
	_jsii_.Get(
		j,
		"linuxProcessParams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference) LinuxProcessParamsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"linuxProcessParamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference) MaxIdleTimeInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxIdleTimeInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference) MaxIdleTimeInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxIdleTimeInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference) MaxInstancesCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxInstancesCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference) MaxInstancesCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxInstancesCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference) MaxQueueSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxQueueSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference) MaxQueueSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxQueueSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference) Pinned() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pinned",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference) PinnedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pinnedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference) StatusTimeoutInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"statusTimeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference) StatusTimeoutInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"statusTimeoutInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference) TimeoutInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference) TimeoutInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInSecondsInput",
		&returns,
	)
	return returns
}


func NewGreengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference {
	_init_.Initialize()

	if err := validateNewGreengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.greengrassv2ComponentVersion.Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGreengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference_Override(g Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.greengrassv2ComponentVersion.Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference)SetEnvironmentVariables(val *map[string]*string) {
	if err := j.validateSetEnvironmentVariablesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environmentVariables",
		val,
	)
}

func (j *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference)SetExecArgs(val *[]*string) {
	if err := j.validateSetExecArgsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"execArgs",
		val,
	)
}

func (j *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference)SetInputPayloadEncodingType(val *string) {
	if err := j.validateSetInputPayloadEncodingTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputPayloadEncodingType",
		val,
	)
}

func (j *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference)SetMaxIdleTimeInSeconds(val *float64) {
	if err := j.validateSetMaxIdleTimeInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxIdleTimeInSeconds",
		val,
	)
}

func (j *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference)SetMaxInstancesCount(val *float64) {
	if err := j.validateSetMaxInstancesCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxInstancesCount",
		val,
	)
}

func (j *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference)SetMaxQueueSize(val *float64) {
	if err := j.validateSetMaxQueueSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxQueueSize",
		val,
	)
}

func (j *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference)SetPinned(val interface{}) {
	if err := j.validateSetPinnedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pinned",
		val,
	)
}

func (j *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference)SetStatusTimeoutInSeconds(val *float64) {
	if err := j.validateSetStatusTimeoutInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"statusTimeoutInSeconds",
		val,
	)
}

func (j *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference)SetTimeoutInSeconds(val *float64) {
	if err := j.validateSetTimeoutInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeoutInSeconds",
		val,
	)
}

func (g *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference) PutEventSources(value interface{}) {
	if err := g.validatePutEventSourcesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEventSources",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference) PutLinuxProcessParams(value *Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersLinuxProcessParams) {
	if err := g.validatePutLinuxProcessParamsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLinuxProcessParams",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference) ResetEnvironmentVariables() {
	_jsii_.InvokeVoid(
		g,
		"resetEnvironmentVariables",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference) ResetEventSources() {
	_jsii_.InvokeVoid(
		g,
		"resetEventSources",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference) ResetExecArgs() {
	_jsii_.InvokeVoid(
		g,
		"resetExecArgs",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference) ResetInputPayloadEncodingType() {
	_jsii_.InvokeVoid(
		g,
		"resetInputPayloadEncodingType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference) ResetLinuxProcessParams() {
	_jsii_.InvokeVoid(
		g,
		"resetLinuxProcessParams",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference) ResetMaxIdleTimeInSeconds() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxIdleTimeInSeconds",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference) ResetMaxInstancesCount() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxInstancesCount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference) ResetMaxQueueSize() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxQueueSize",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference) ResetPinned() {
	_jsii_.InvokeVoid(
		g,
		"resetPinned",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference) ResetStatusTimeoutInSeconds() {
	_jsii_.InvokeVoid(
		g,
		"resetStatusTimeoutInSeconds",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference) ResetTimeoutInSeconds() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeoutInSeconds",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := g.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

