// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appstreamappblock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/appstreamappblock/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AppstreamAppBlockPostSetupScriptDetailsOutputReference interface {
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
	ExecutableParameters() *string
	SetExecutableParameters(val *string)
	ExecutableParametersInput() *string
	ExecutablePath() *string
	SetExecutablePath(val *string)
	ExecutablePathInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ScriptS3Location() AppstreamAppBlockPostSetupScriptDetailsScriptS3LocationOutputReference
	ScriptS3LocationInput() interface{}
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
	PutScriptS3Location(value *AppstreamAppBlockPostSetupScriptDetailsScriptS3Location)
	ResetExecutableParameters()
	ResetExecutablePath()
	ResetScriptS3Location()
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

// The jsii proxy struct for AppstreamAppBlockPostSetupScriptDetailsOutputReference
type jsiiProxy_AppstreamAppBlockPostSetupScriptDetailsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AppstreamAppBlockPostSetupScriptDetailsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamAppBlockPostSetupScriptDetailsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamAppBlockPostSetupScriptDetailsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamAppBlockPostSetupScriptDetailsOutputReference) ExecutableParameters() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executableParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamAppBlockPostSetupScriptDetailsOutputReference) ExecutableParametersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executableParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamAppBlockPostSetupScriptDetailsOutputReference) ExecutablePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executablePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamAppBlockPostSetupScriptDetailsOutputReference) ExecutablePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executablePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamAppBlockPostSetupScriptDetailsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamAppBlockPostSetupScriptDetailsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamAppBlockPostSetupScriptDetailsOutputReference) ScriptS3Location() AppstreamAppBlockPostSetupScriptDetailsScriptS3LocationOutputReference {
	var returns AppstreamAppBlockPostSetupScriptDetailsScriptS3LocationOutputReference
	_jsii_.Get(
		j,
		"scriptS3Location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamAppBlockPostSetupScriptDetailsOutputReference) ScriptS3LocationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scriptS3LocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamAppBlockPostSetupScriptDetailsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamAppBlockPostSetupScriptDetailsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamAppBlockPostSetupScriptDetailsOutputReference) TimeoutInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamAppBlockPostSetupScriptDetailsOutputReference) TimeoutInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInSecondsInput",
		&returns,
	)
	return returns
}


func NewAppstreamAppBlockPostSetupScriptDetailsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AppstreamAppBlockPostSetupScriptDetailsOutputReference {
	_init_.Initialize()

	if err := validateNewAppstreamAppBlockPostSetupScriptDetailsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppstreamAppBlockPostSetupScriptDetailsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.appstreamAppBlock.AppstreamAppBlockPostSetupScriptDetailsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppstreamAppBlockPostSetupScriptDetailsOutputReference_Override(a AppstreamAppBlockPostSetupScriptDetailsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.appstreamAppBlock.AppstreamAppBlockPostSetupScriptDetailsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppstreamAppBlockPostSetupScriptDetailsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppstreamAppBlockPostSetupScriptDetailsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppstreamAppBlockPostSetupScriptDetailsOutputReference)SetExecutableParameters(val *string) {
	if err := j.validateSetExecutableParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executableParameters",
		val,
	)
}

func (j *jsiiProxy_AppstreamAppBlockPostSetupScriptDetailsOutputReference)SetExecutablePath(val *string) {
	if err := j.validateSetExecutablePathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executablePath",
		val,
	)
}

func (j *jsiiProxy_AppstreamAppBlockPostSetupScriptDetailsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppstreamAppBlockPostSetupScriptDetailsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppstreamAppBlockPostSetupScriptDetailsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AppstreamAppBlockPostSetupScriptDetailsOutputReference)SetTimeoutInSeconds(val *float64) {
	if err := j.validateSetTimeoutInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeoutInSeconds",
		val,
	)
}

func (a *jsiiProxy_AppstreamAppBlockPostSetupScriptDetailsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppstreamAppBlockPostSetupScriptDetailsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppstreamAppBlockPostSetupScriptDetailsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppstreamAppBlockPostSetupScriptDetailsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppstreamAppBlockPostSetupScriptDetailsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppstreamAppBlockPostSetupScriptDetailsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppstreamAppBlockPostSetupScriptDetailsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppstreamAppBlockPostSetupScriptDetailsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppstreamAppBlockPostSetupScriptDetailsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppstreamAppBlockPostSetupScriptDetailsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppstreamAppBlockPostSetupScriptDetailsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppstreamAppBlockPostSetupScriptDetailsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppstreamAppBlockPostSetupScriptDetailsOutputReference) PutScriptS3Location(value *AppstreamAppBlockPostSetupScriptDetailsScriptS3Location) {
	if err := a.validatePutScriptS3LocationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putScriptS3Location",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppstreamAppBlockPostSetupScriptDetailsOutputReference) ResetExecutableParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetExecutableParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamAppBlockPostSetupScriptDetailsOutputReference) ResetExecutablePath() {
	_jsii_.InvokeVoid(
		a,
		"resetExecutablePath",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamAppBlockPostSetupScriptDetailsOutputReference) ResetScriptS3Location() {
	_jsii_.InvokeVoid(
		a,
		"resetScriptS3Location",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamAppBlockPostSetupScriptDetailsOutputReference) ResetTimeoutInSeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeoutInSeconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamAppBlockPostSetupScriptDetailsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppstreamAppBlockPostSetupScriptDetailsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

