// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package eventsv2eventsource

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/eventsv2eventsource/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Eventsv2EventSourceConfigurationAwsServiceEventsConfigurationOutputReference interface {
	cdktn.ComplexObject
	AwsService() *string
	SetAwsService(val *string)
	AwsServiceInput() *string
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
	OnFailureConfiguration() Eventsv2EventSourceConfigurationAwsServiceEventsConfigurationOnFailureConfigurationOutputReference
	OnFailureConfigurationInput() interface{}
	Pattern() *string
	SetPattern(val *string)
	PatternInput() *string
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
	PutOnFailureConfiguration(value *Eventsv2EventSourceConfigurationAwsServiceEventsConfigurationOnFailureConfiguration)
	ResetAwsService()
	ResetOnFailureConfiguration()
	ResetPattern()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Eventsv2EventSourceConfigurationAwsServiceEventsConfigurationOutputReference
type jsiiProxy_Eventsv2EventSourceConfigurationAwsServiceEventsConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_Eventsv2EventSourceConfigurationAwsServiceEventsConfigurationOutputReference) AwsService() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2EventSourceConfigurationAwsServiceEventsConfigurationOutputReference) AwsServiceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2EventSourceConfigurationAwsServiceEventsConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2EventSourceConfigurationAwsServiceEventsConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2EventSourceConfigurationAwsServiceEventsConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2EventSourceConfigurationAwsServiceEventsConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2EventSourceConfigurationAwsServiceEventsConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2EventSourceConfigurationAwsServiceEventsConfigurationOutputReference) OnFailureConfiguration() Eventsv2EventSourceConfigurationAwsServiceEventsConfigurationOnFailureConfigurationOutputReference {
	var returns Eventsv2EventSourceConfigurationAwsServiceEventsConfigurationOnFailureConfigurationOutputReference
	_jsii_.Get(
		j,
		"onFailureConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2EventSourceConfigurationAwsServiceEventsConfigurationOutputReference) OnFailureConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onFailureConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2EventSourceConfigurationAwsServiceEventsConfigurationOutputReference) Pattern() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2EventSourceConfigurationAwsServiceEventsConfigurationOutputReference) PatternInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"patternInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2EventSourceConfigurationAwsServiceEventsConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2EventSourceConfigurationAwsServiceEventsConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEventsv2EventSourceConfigurationAwsServiceEventsConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) Eventsv2EventSourceConfigurationAwsServiceEventsConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewEventsv2EventSourceConfigurationAwsServiceEventsConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Eventsv2EventSourceConfigurationAwsServiceEventsConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.eventsv2EventSource.Eventsv2EventSourceConfigurationAwsServiceEventsConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEventsv2EventSourceConfigurationAwsServiceEventsConfigurationOutputReference_Override(e Eventsv2EventSourceConfigurationAwsServiceEventsConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.eventsv2EventSource.Eventsv2EventSourceConfigurationAwsServiceEventsConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_Eventsv2EventSourceConfigurationAwsServiceEventsConfigurationOutputReference)SetAwsService(val *string) {
	if err := j.validateSetAwsServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"awsService",
		val,
	)
}

func (j *jsiiProxy_Eventsv2EventSourceConfigurationAwsServiceEventsConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Eventsv2EventSourceConfigurationAwsServiceEventsConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Eventsv2EventSourceConfigurationAwsServiceEventsConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Eventsv2EventSourceConfigurationAwsServiceEventsConfigurationOutputReference)SetPattern(val *string) {
	if err := j.validateSetPatternParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pattern",
		val,
	)
}

func (j *jsiiProxy_Eventsv2EventSourceConfigurationAwsServiceEventsConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Eventsv2EventSourceConfigurationAwsServiceEventsConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_Eventsv2EventSourceConfigurationAwsServiceEventsConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Eventsv2EventSourceConfigurationAwsServiceEventsConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Eventsv2EventSourceConfigurationAwsServiceEventsConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Eventsv2EventSourceConfigurationAwsServiceEventsConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Eventsv2EventSourceConfigurationAwsServiceEventsConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Eventsv2EventSourceConfigurationAwsServiceEventsConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Eventsv2EventSourceConfigurationAwsServiceEventsConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Eventsv2EventSourceConfigurationAwsServiceEventsConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Eventsv2EventSourceConfigurationAwsServiceEventsConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Eventsv2EventSourceConfigurationAwsServiceEventsConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Eventsv2EventSourceConfigurationAwsServiceEventsConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Eventsv2EventSourceConfigurationAwsServiceEventsConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Eventsv2EventSourceConfigurationAwsServiceEventsConfigurationOutputReference) PutOnFailureConfiguration(value *Eventsv2EventSourceConfigurationAwsServiceEventsConfigurationOnFailureConfiguration) {
	if err := e.validatePutOnFailureConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putOnFailureConfiguration",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Eventsv2EventSourceConfigurationAwsServiceEventsConfigurationOutputReference) ResetAwsService() {
	_jsii_.InvokeVoid(
		e,
		"resetAwsService",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Eventsv2EventSourceConfigurationAwsServiceEventsConfigurationOutputReference) ResetOnFailureConfiguration() {
	_jsii_.InvokeVoid(
		e,
		"resetOnFailureConfiguration",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Eventsv2EventSourceConfigurationAwsServiceEventsConfigurationOutputReference) ResetPattern() {
	_jsii_.InvokeVoid(
		e,
		"resetPattern",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Eventsv2EventSourceConfigurationAwsServiceEventsConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := e.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Eventsv2EventSourceConfigurationAwsServiceEventsConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

