// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package eventsv2subscriber

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/eventsv2subscriber/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Eventsv2SubscriberInvokeConfigurationKinesisParametersOutputReference interface {
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
	ExplicitHashKey() *string
	SetExplicitHashKey(val *string)
	ExplicitHashKeyInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PartitionKey() *string
	SetPartitionKey(val *string)
	PartitionKeyInput() *string
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
	ResetExplicitHashKey()
	ResetPartitionKey()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Eventsv2SubscriberInvokeConfigurationKinesisParametersOutputReference
type jsiiProxy_Eventsv2SubscriberInvokeConfigurationKinesisParametersOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationKinesisParametersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationKinesisParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationKinesisParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationKinesisParametersOutputReference) ExplicitHashKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"explicitHashKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationKinesisParametersOutputReference) ExplicitHashKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"explicitHashKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationKinesisParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationKinesisParametersOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationKinesisParametersOutputReference) PartitionKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partitionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationKinesisParametersOutputReference) PartitionKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partitionKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationKinesisParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationKinesisParametersOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEventsv2SubscriberInvokeConfigurationKinesisParametersOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) Eventsv2SubscriberInvokeConfigurationKinesisParametersOutputReference {
	_init_.Initialize()

	if err := validateNewEventsv2SubscriberInvokeConfigurationKinesisParametersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Eventsv2SubscriberInvokeConfigurationKinesisParametersOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.eventsv2Subscriber.Eventsv2SubscriberInvokeConfigurationKinesisParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEventsv2SubscriberInvokeConfigurationKinesisParametersOutputReference_Override(e Eventsv2SubscriberInvokeConfigurationKinesisParametersOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.eventsv2Subscriber.Eventsv2SubscriberInvokeConfigurationKinesisParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationKinesisParametersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationKinesisParametersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationKinesisParametersOutputReference)SetExplicitHashKey(val *string) {
	if err := j.validateSetExplicitHashKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"explicitHashKey",
		val,
	)
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationKinesisParametersOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationKinesisParametersOutputReference)SetPartitionKey(val *string) {
	if err := j.validateSetPartitionKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"partitionKey",
		val,
	)
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationKinesisParametersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationKinesisParametersOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationKinesisParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationKinesisParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationKinesisParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationKinesisParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationKinesisParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationKinesisParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationKinesisParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationKinesisParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationKinesisParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationKinesisParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationKinesisParametersOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationKinesisParametersOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationKinesisParametersOutputReference) ResetExplicitHashKey() {
	_jsii_.InvokeVoid(
		e,
		"resetExplicitHashKey",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationKinesisParametersOutputReference) ResetPartitionKey() {
	_jsii_.InvokeVoid(
		e,
		"resetPartitionKey",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationKinesisParametersOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationKinesisParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

