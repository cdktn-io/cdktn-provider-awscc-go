// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lambdacapacityprovider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/lambdacapacityprovider/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LambdaCapacityProviderCapacityProviderScalingConfigOutputReference interface {
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
	MaxVCpuCount() *float64
	SetMaxVCpuCount(val *float64)
	MaxVCpuCountInput() *float64
	ScalingMode() *string
	SetScalingMode(val *string)
	ScalingModeInput() *string
	ScalingPolicies() LambdaCapacityProviderCapacityProviderScalingConfigScalingPoliciesList
	ScalingPoliciesInput() interface{}
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
	PutScalingPolicies(value interface{})
	ResetMaxVCpuCount()
	ResetScalingMode()
	ResetScalingPolicies()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LambdaCapacityProviderCapacityProviderScalingConfigOutputReference
type jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference) MaxVCpuCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxVCpuCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference) MaxVCpuCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxVCpuCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference) ScalingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scalingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference) ScalingModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scalingModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference) ScalingPolicies() LambdaCapacityProviderCapacityProviderScalingConfigScalingPoliciesList {
	var returns LambdaCapacityProviderCapacityProviderScalingConfigScalingPoliciesList
	_jsii_.Get(
		j,
		"scalingPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference) ScalingPoliciesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scalingPoliciesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewLambdaCapacityProviderCapacityProviderScalingConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) LambdaCapacityProviderCapacityProviderScalingConfigOutputReference {
	_init_.Initialize()

	if err := validateNewLambdaCapacityProviderCapacityProviderScalingConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.lambdaCapacityProvider.LambdaCapacityProviderCapacityProviderScalingConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLambdaCapacityProviderCapacityProviderScalingConfigOutputReference_Override(l LambdaCapacityProviderCapacityProviderScalingConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.lambdaCapacityProvider.LambdaCapacityProviderCapacityProviderScalingConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference)SetMaxVCpuCount(val *float64) {
	if err := j.validateSetMaxVCpuCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxVCpuCount",
		val,
	)
}

func (j *jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference)SetScalingMode(val *string) {
	if err := j.validateSetScalingModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scalingMode",
		val,
	)
}

func (j *jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference) PutScalingPolicies(value interface{}) {
	if err := l.validatePutScalingPoliciesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putScalingPolicies",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference) ResetMaxVCpuCount() {
	_jsii_.InvokeVoid(
		l,
		"resetMaxVCpuCount",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference) ResetScalingMode() {
	_jsii_.InvokeVoid(
		l,
		"resetScalingMode",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference) ResetScalingPolicies() {
	_jsii_.InvokeVoid(
		l,
		"resetScalingPolicies",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := l.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

