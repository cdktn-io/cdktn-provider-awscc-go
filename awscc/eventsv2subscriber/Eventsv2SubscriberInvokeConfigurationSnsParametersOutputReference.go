// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package eventsv2subscriber

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/eventsv2subscriber/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Eventsv2SubscriberInvokeConfigurationSnsParametersOutputReference interface {
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
	MessageAttributes() Eventsv2SubscriberInvokeConfigurationSnsParametersMessageAttributesMap
	MessageAttributesInput() interface{}
	MessageDeduplicationId() *string
	SetMessageDeduplicationId(val *string)
	MessageDeduplicationIdInput() *string
	MessageGroupId() *string
	SetMessageGroupId(val *string)
	MessageGroupIdInput() *string
	MessageStructure() *string
	SetMessageStructure(val *string)
	MessageStructureInput() *string
	Subject() *string
	SetSubject(val *string)
	SubjectInput() *string
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
	PutMessageAttributes(value interface{})
	ResetMessageAttributes()
	ResetMessageDeduplicationId()
	ResetMessageGroupId()
	ResetMessageStructure()
	ResetSubject()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Eventsv2SubscriberInvokeConfigurationSnsParametersOutputReference
type jsiiProxy_Eventsv2SubscriberInvokeConfigurationSnsParametersOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationSnsParametersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationSnsParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationSnsParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationSnsParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationSnsParametersOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationSnsParametersOutputReference) MessageAttributes() Eventsv2SubscriberInvokeConfigurationSnsParametersMessageAttributesMap {
	var returns Eventsv2SubscriberInvokeConfigurationSnsParametersMessageAttributesMap
	_jsii_.Get(
		j,
		"messageAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationSnsParametersOutputReference) MessageAttributesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"messageAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationSnsParametersOutputReference) MessageDeduplicationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageDeduplicationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationSnsParametersOutputReference) MessageDeduplicationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageDeduplicationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationSnsParametersOutputReference) MessageGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationSnsParametersOutputReference) MessageGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationSnsParametersOutputReference) MessageStructure() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageStructure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationSnsParametersOutputReference) MessageStructureInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageStructureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationSnsParametersOutputReference) Subject() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationSnsParametersOutputReference) SubjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationSnsParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationSnsParametersOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEventsv2SubscriberInvokeConfigurationSnsParametersOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) Eventsv2SubscriberInvokeConfigurationSnsParametersOutputReference {
	_init_.Initialize()

	if err := validateNewEventsv2SubscriberInvokeConfigurationSnsParametersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Eventsv2SubscriberInvokeConfigurationSnsParametersOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.eventsv2Subscriber.Eventsv2SubscriberInvokeConfigurationSnsParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEventsv2SubscriberInvokeConfigurationSnsParametersOutputReference_Override(e Eventsv2SubscriberInvokeConfigurationSnsParametersOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.eventsv2Subscriber.Eventsv2SubscriberInvokeConfigurationSnsParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationSnsParametersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationSnsParametersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationSnsParametersOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationSnsParametersOutputReference)SetMessageDeduplicationId(val *string) {
	if err := j.validateSetMessageDeduplicationIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"messageDeduplicationId",
		val,
	)
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationSnsParametersOutputReference)SetMessageGroupId(val *string) {
	if err := j.validateSetMessageGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"messageGroupId",
		val,
	)
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationSnsParametersOutputReference)SetMessageStructure(val *string) {
	if err := j.validateSetMessageStructureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"messageStructure",
		val,
	)
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationSnsParametersOutputReference)SetSubject(val *string) {
	if err := j.validateSetSubjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subject",
		val,
	)
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationSnsParametersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationSnsParametersOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationSnsParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationSnsParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationSnsParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationSnsParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationSnsParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationSnsParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationSnsParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationSnsParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationSnsParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationSnsParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationSnsParametersOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationSnsParametersOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationSnsParametersOutputReference) PutMessageAttributes(value interface{}) {
	if err := e.validatePutMessageAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putMessageAttributes",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationSnsParametersOutputReference) ResetMessageAttributes() {
	_jsii_.InvokeVoid(
		e,
		"resetMessageAttributes",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationSnsParametersOutputReference) ResetMessageDeduplicationId() {
	_jsii_.InvokeVoid(
		e,
		"resetMessageDeduplicationId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationSnsParametersOutputReference) ResetMessageGroupId() {
	_jsii_.InvokeVoid(
		e,
		"resetMessageGroupId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationSnsParametersOutputReference) ResetMessageStructure() {
	_jsii_.InvokeVoid(
		e,
		"resetMessageStructure",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationSnsParametersOutputReference) ResetSubject() {
	_jsii_.InvokeVoid(
		e,
		"resetSubject",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationSnsParametersOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationSnsParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

