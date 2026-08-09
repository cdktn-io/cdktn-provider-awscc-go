// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lambdaeventsourcemapping

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/lambdaeventsourcemapping/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LambdaEventSourceMappingProvisionedPollerConfigOutputReference interface {
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
	MaximumPollers() *float64
	SetMaximumPollers(val *float64)
	MaximumPollersInput() *float64
	MinimumPollers() *float64
	SetMinimumPollers(val *float64)
	MinimumPollersInput() *float64
	PollerGroupName() *string
	SetPollerGroupName(val *string)
	PollerGroupNameInput() *string
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
	ResetMaximumPollers()
	ResetMinimumPollers()
	ResetPollerGroupName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LambdaEventSourceMappingProvisionedPollerConfigOutputReference
type jsiiProxy_LambdaEventSourceMappingProvisionedPollerConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_LambdaEventSourceMappingProvisionedPollerConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMappingProvisionedPollerConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMappingProvisionedPollerConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMappingProvisionedPollerConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMappingProvisionedPollerConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMappingProvisionedPollerConfigOutputReference) MaximumPollers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumPollers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMappingProvisionedPollerConfigOutputReference) MaximumPollersInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumPollersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMappingProvisionedPollerConfigOutputReference) MinimumPollers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumPollers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMappingProvisionedPollerConfigOutputReference) MinimumPollersInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumPollersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMappingProvisionedPollerConfigOutputReference) PollerGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pollerGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMappingProvisionedPollerConfigOutputReference) PollerGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pollerGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMappingProvisionedPollerConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMappingProvisionedPollerConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewLambdaEventSourceMappingProvisionedPollerConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) LambdaEventSourceMappingProvisionedPollerConfigOutputReference {
	_init_.Initialize()

	if err := validateNewLambdaEventSourceMappingProvisionedPollerConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LambdaEventSourceMappingProvisionedPollerConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.lambdaEventSourceMapping.LambdaEventSourceMappingProvisionedPollerConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLambdaEventSourceMappingProvisionedPollerConfigOutputReference_Override(l LambdaEventSourceMappingProvisionedPollerConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.lambdaEventSourceMapping.LambdaEventSourceMappingProvisionedPollerConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LambdaEventSourceMappingProvisionedPollerConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LambdaEventSourceMappingProvisionedPollerConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LambdaEventSourceMappingProvisionedPollerConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LambdaEventSourceMappingProvisionedPollerConfigOutputReference)SetMaximumPollers(val *float64) {
	if err := j.validateSetMaximumPollersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumPollers",
		val,
	)
}

func (j *jsiiProxy_LambdaEventSourceMappingProvisionedPollerConfigOutputReference)SetMinimumPollers(val *float64) {
	if err := j.validateSetMinimumPollersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimumPollers",
		val,
	)
}

func (j *jsiiProxy_LambdaEventSourceMappingProvisionedPollerConfigOutputReference)SetPollerGroupName(val *string) {
	if err := j.validateSetPollerGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pollerGroupName",
		val,
	)
}

func (j *jsiiProxy_LambdaEventSourceMappingProvisionedPollerConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LambdaEventSourceMappingProvisionedPollerConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LambdaEventSourceMappingProvisionedPollerConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaEventSourceMappingProvisionedPollerConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_LambdaEventSourceMappingProvisionedPollerConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (l *jsiiProxy_LambdaEventSourceMappingProvisionedPollerConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_LambdaEventSourceMappingProvisionedPollerConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_LambdaEventSourceMappingProvisionedPollerConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_LambdaEventSourceMappingProvisionedPollerConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_LambdaEventSourceMappingProvisionedPollerConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_LambdaEventSourceMappingProvisionedPollerConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_LambdaEventSourceMappingProvisionedPollerConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_LambdaEventSourceMappingProvisionedPollerConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaEventSourceMappingProvisionedPollerConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (l *jsiiProxy_LambdaEventSourceMappingProvisionedPollerConfigOutputReference) ResetMaximumPollers() {
	_jsii_.InvokeVoid(
		l,
		"resetMaximumPollers",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaEventSourceMappingProvisionedPollerConfigOutputReference) ResetMinimumPollers() {
	_jsii_.InvokeVoid(
		l,
		"resetMinimumPollers",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaEventSourceMappingProvisionedPollerConfigOutputReference) ResetPollerGroupName() {
	_jsii_.InvokeVoid(
		l,
		"resetPollerGroupName",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaEventSourceMappingProvisionedPollerConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (l *jsiiProxy_LambdaEventSourceMappingProvisionedPollerConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

