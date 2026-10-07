// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package eventsv2subscriber

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/eventsv2subscriber/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Eventsv2SubscriberInvokeConfigurationEventBusV2ParametersSystemMetadataOutputReference interface {
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
	DeduplicationId() *string
	SetDeduplicationId(val *string)
	DeduplicationIdInput() *string
	EventGroupId() *string
	SetEventGroupId(val *string)
	EventGroupIdInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
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
	ResetDeduplicationId()
	ResetEventGroupId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Eventsv2SubscriberInvokeConfigurationEventBusV2ParametersSystemMetadataOutputReference
type jsiiProxy_Eventsv2SubscriberInvokeConfigurationEventBusV2ParametersSystemMetadataOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationEventBusV2ParametersSystemMetadataOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationEventBusV2ParametersSystemMetadataOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationEventBusV2ParametersSystemMetadataOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationEventBusV2ParametersSystemMetadataOutputReference) DeduplicationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deduplicationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationEventBusV2ParametersSystemMetadataOutputReference) DeduplicationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deduplicationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationEventBusV2ParametersSystemMetadataOutputReference) EventGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationEventBusV2ParametersSystemMetadataOutputReference) EventGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationEventBusV2ParametersSystemMetadataOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationEventBusV2ParametersSystemMetadataOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationEventBusV2ParametersSystemMetadataOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationEventBusV2ParametersSystemMetadataOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEventsv2SubscriberInvokeConfigurationEventBusV2ParametersSystemMetadataOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) Eventsv2SubscriberInvokeConfigurationEventBusV2ParametersSystemMetadataOutputReference {
	_init_.Initialize()

	if err := validateNewEventsv2SubscriberInvokeConfigurationEventBusV2ParametersSystemMetadataOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Eventsv2SubscriberInvokeConfigurationEventBusV2ParametersSystemMetadataOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.eventsv2Subscriber.Eventsv2SubscriberInvokeConfigurationEventBusV2ParametersSystemMetadataOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEventsv2SubscriberInvokeConfigurationEventBusV2ParametersSystemMetadataOutputReference_Override(e Eventsv2SubscriberInvokeConfigurationEventBusV2ParametersSystemMetadataOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.eventsv2Subscriber.Eventsv2SubscriberInvokeConfigurationEventBusV2ParametersSystemMetadataOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationEventBusV2ParametersSystemMetadataOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationEventBusV2ParametersSystemMetadataOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationEventBusV2ParametersSystemMetadataOutputReference)SetDeduplicationId(val *string) {
	if err := j.validateSetDeduplicationIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deduplicationId",
		val,
	)
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationEventBusV2ParametersSystemMetadataOutputReference)SetEventGroupId(val *string) {
	if err := j.validateSetEventGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventGroupId",
		val,
	)
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationEventBusV2ParametersSystemMetadataOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationEventBusV2ParametersSystemMetadataOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationEventBusV2ParametersSystemMetadataOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationEventBusV2ParametersSystemMetadataOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationEventBusV2ParametersSystemMetadataOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationEventBusV2ParametersSystemMetadataOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationEventBusV2ParametersSystemMetadataOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationEventBusV2ParametersSystemMetadataOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationEventBusV2ParametersSystemMetadataOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationEventBusV2ParametersSystemMetadataOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationEventBusV2ParametersSystemMetadataOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationEventBusV2ParametersSystemMetadataOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationEventBusV2ParametersSystemMetadataOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationEventBusV2ParametersSystemMetadataOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationEventBusV2ParametersSystemMetadataOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationEventBusV2ParametersSystemMetadataOutputReference) ResetDeduplicationId() {
	_jsii_.InvokeVoid(
		e,
		"resetDeduplicationId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationEventBusV2ParametersSystemMetadataOutputReference) ResetEventGroupId() {
	_jsii_.InvokeVoid(
		e,
		"resetEventGroupId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationEventBusV2ParametersSystemMetadataOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationEventBusV2ParametersSystemMetadataOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

