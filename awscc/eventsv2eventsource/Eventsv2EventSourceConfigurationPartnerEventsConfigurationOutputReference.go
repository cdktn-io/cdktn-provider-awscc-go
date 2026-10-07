// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package eventsv2eventsource

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/eventsv2eventsource/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Eventsv2EventSourceConfigurationPartnerEventsConfigurationOutputReference interface {
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
	OnFailureConfiguration() Eventsv2EventSourceConfigurationPartnerEventsConfigurationOnFailureConfigurationOutputReference
	OnFailureConfigurationInput() interface{}
	PartnerBusKmsKeyIdentifier() *string
	SetPartnerBusKmsKeyIdentifier(val *string)
	PartnerBusKmsKeyIdentifierInput() *string
	PartnerEventSourceArn() *string
	SetPartnerEventSourceArn(val *string)
	PartnerEventSourceArnInput() *string
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
	PutOnFailureConfiguration(value *Eventsv2EventSourceConfigurationPartnerEventsConfigurationOnFailureConfiguration)
	ResetOnFailureConfiguration()
	ResetPartnerBusKmsKeyIdentifier()
	ResetPartnerEventSourceArn()
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

// The jsii proxy struct for Eventsv2EventSourceConfigurationPartnerEventsConfigurationOutputReference
type jsiiProxy_Eventsv2EventSourceConfigurationPartnerEventsConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_Eventsv2EventSourceConfigurationPartnerEventsConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2EventSourceConfigurationPartnerEventsConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2EventSourceConfigurationPartnerEventsConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2EventSourceConfigurationPartnerEventsConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2EventSourceConfigurationPartnerEventsConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2EventSourceConfigurationPartnerEventsConfigurationOutputReference) OnFailureConfiguration() Eventsv2EventSourceConfigurationPartnerEventsConfigurationOnFailureConfigurationOutputReference {
	var returns Eventsv2EventSourceConfigurationPartnerEventsConfigurationOnFailureConfigurationOutputReference
	_jsii_.Get(
		j,
		"onFailureConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2EventSourceConfigurationPartnerEventsConfigurationOutputReference) OnFailureConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onFailureConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2EventSourceConfigurationPartnerEventsConfigurationOutputReference) PartnerBusKmsKeyIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partnerBusKmsKeyIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2EventSourceConfigurationPartnerEventsConfigurationOutputReference) PartnerBusKmsKeyIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partnerBusKmsKeyIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2EventSourceConfigurationPartnerEventsConfigurationOutputReference) PartnerEventSourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partnerEventSourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2EventSourceConfigurationPartnerEventsConfigurationOutputReference) PartnerEventSourceArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partnerEventSourceArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2EventSourceConfigurationPartnerEventsConfigurationOutputReference) Pattern() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2EventSourceConfigurationPartnerEventsConfigurationOutputReference) PatternInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"patternInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2EventSourceConfigurationPartnerEventsConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2EventSourceConfigurationPartnerEventsConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEventsv2EventSourceConfigurationPartnerEventsConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) Eventsv2EventSourceConfigurationPartnerEventsConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewEventsv2EventSourceConfigurationPartnerEventsConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Eventsv2EventSourceConfigurationPartnerEventsConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.eventsv2EventSource.Eventsv2EventSourceConfigurationPartnerEventsConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEventsv2EventSourceConfigurationPartnerEventsConfigurationOutputReference_Override(e Eventsv2EventSourceConfigurationPartnerEventsConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.eventsv2EventSource.Eventsv2EventSourceConfigurationPartnerEventsConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_Eventsv2EventSourceConfigurationPartnerEventsConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Eventsv2EventSourceConfigurationPartnerEventsConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Eventsv2EventSourceConfigurationPartnerEventsConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Eventsv2EventSourceConfigurationPartnerEventsConfigurationOutputReference)SetPartnerBusKmsKeyIdentifier(val *string) {
	if err := j.validateSetPartnerBusKmsKeyIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"partnerBusKmsKeyIdentifier",
		val,
	)
}

func (j *jsiiProxy_Eventsv2EventSourceConfigurationPartnerEventsConfigurationOutputReference)SetPartnerEventSourceArn(val *string) {
	if err := j.validateSetPartnerEventSourceArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"partnerEventSourceArn",
		val,
	)
}

func (j *jsiiProxy_Eventsv2EventSourceConfigurationPartnerEventsConfigurationOutputReference)SetPattern(val *string) {
	if err := j.validateSetPatternParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pattern",
		val,
	)
}

func (j *jsiiProxy_Eventsv2EventSourceConfigurationPartnerEventsConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Eventsv2EventSourceConfigurationPartnerEventsConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_Eventsv2EventSourceConfigurationPartnerEventsConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Eventsv2EventSourceConfigurationPartnerEventsConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_Eventsv2EventSourceConfigurationPartnerEventsConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_Eventsv2EventSourceConfigurationPartnerEventsConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_Eventsv2EventSourceConfigurationPartnerEventsConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_Eventsv2EventSourceConfigurationPartnerEventsConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_Eventsv2EventSourceConfigurationPartnerEventsConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_Eventsv2EventSourceConfigurationPartnerEventsConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_Eventsv2EventSourceConfigurationPartnerEventsConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_Eventsv2EventSourceConfigurationPartnerEventsConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_Eventsv2EventSourceConfigurationPartnerEventsConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Eventsv2EventSourceConfigurationPartnerEventsConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_Eventsv2EventSourceConfigurationPartnerEventsConfigurationOutputReference) PutOnFailureConfiguration(value *Eventsv2EventSourceConfigurationPartnerEventsConfigurationOnFailureConfiguration) {
	if err := e.validatePutOnFailureConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putOnFailureConfiguration",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Eventsv2EventSourceConfigurationPartnerEventsConfigurationOutputReference) ResetOnFailureConfiguration() {
	_jsii_.InvokeVoid(
		e,
		"resetOnFailureConfiguration",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Eventsv2EventSourceConfigurationPartnerEventsConfigurationOutputReference) ResetPartnerBusKmsKeyIdentifier() {
	_jsii_.InvokeVoid(
		e,
		"resetPartnerBusKmsKeyIdentifier",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Eventsv2EventSourceConfigurationPartnerEventsConfigurationOutputReference) ResetPartnerEventSourceArn() {
	_jsii_.InvokeVoid(
		e,
		"resetPartnerEventSourceArn",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Eventsv2EventSourceConfigurationPartnerEventsConfigurationOutputReference) ResetPattern() {
	_jsii_.InvokeVoid(
		e,
		"resetPattern",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Eventsv2EventSourceConfigurationPartnerEventsConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (e *jsiiProxy_Eventsv2EventSourceConfigurationPartnerEventsConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

