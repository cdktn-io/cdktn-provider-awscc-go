// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iottopicrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/iottopicrule/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IotTopicRuleTopicRulePayloadErrorActionRepublishHeadersOutputReference interface {
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
	ContentType() *string
	SetContentType(val *string)
	ContentTypeInput() *string
	CorrelationData() *string
	SetCorrelationData(val *string)
	CorrelationDataInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MessageExpiry() *string
	SetMessageExpiry(val *string)
	MessageExpiryInput() *string
	PayloadFormatIndicator() *string
	SetPayloadFormatIndicator(val *string)
	PayloadFormatIndicatorInput() *string
	ResponseTopic() *string
	SetResponseTopic(val *string)
	ResponseTopicInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UserProperties() IotTopicRuleTopicRulePayloadErrorActionRepublishHeadersUserPropertiesList
	UserPropertiesInput() interface{}
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
	PutUserProperties(value interface{})
	ResetContentType()
	ResetCorrelationData()
	ResetMessageExpiry()
	ResetPayloadFormatIndicator()
	ResetResponseTopic()
	ResetUserProperties()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IotTopicRuleTopicRulePayloadErrorActionRepublishHeadersOutputReference
type jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionRepublishHeadersOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionRepublishHeadersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionRepublishHeadersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionRepublishHeadersOutputReference) ContentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionRepublishHeadersOutputReference) ContentTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionRepublishHeadersOutputReference) CorrelationData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"correlationData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionRepublishHeadersOutputReference) CorrelationDataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"correlationDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionRepublishHeadersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionRepublishHeadersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionRepublishHeadersOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionRepublishHeadersOutputReference) MessageExpiry() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageExpiry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionRepublishHeadersOutputReference) MessageExpiryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageExpiryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionRepublishHeadersOutputReference) PayloadFormatIndicator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"payloadFormatIndicator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionRepublishHeadersOutputReference) PayloadFormatIndicatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"payloadFormatIndicatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionRepublishHeadersOutputReference) ResponseTopic() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseTopic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionRepublishHeadersOutputReference) ResponseTopicInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseTopicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionRepublishHeadersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionRepublishHeadersOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionRepublishHeadersOutputReference) UserProperties() IotTopicRuleTopicRulePayloadErrorActionRepublishHeadersUserPropertiesList {
	var returns IotTopicRuleTopicRulePayloadErrorActionRepublishHeadersUserPropertiesList
	_jsii_.Get(
		j,
		"userProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionRepublishHeadersOutputReference) UserPropertiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userPropertiesInput",
		&returns,
	)
	return returns
}


func NewIotTopicRuleTopicRulePayloadErrorActionRepublishHeadersOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) IotTopicRuleTopicRulePayloadErrorActionRepublishHeadersOutputReference {
	_init_.Initialize()

	if err := validateNewIotTopicRuleTopicRulePayloadErrorActionRepublishHeadersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionRepublishHeadersOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.iotTopicRule.IotTopicRuleTopicRulePayloadErrorActionRepublishHeadersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIotTopicRuleTopicRulePayloadErrorActionRepublishHeadersOutputReference_Override(i IotTopicRuleTopicRulePayloadErrorActionRepublishHeadersOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.iotTopicRule.IotTopicRuleTopicRulePayloadErrorActionRepublishHeadersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionRepublishHeadersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionRepublishHeadersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionRepublishHeadersOutputReference)SetContentType(val *string) {
	if err := j.validateSetContentTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentType",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionRepublishHeadersOutputReference)SetCorrelationData(val *string) {
	if err := j.validateSetCorrelationDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"correlationData",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionRepublishHeadersOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionRepublishHeadersOutputReference)SetMessageExpiry(val *string) {
	if err := j.validateSetMessageExpiryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"messageExpiry",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionRepublishHeadersOutputReference)SetPayloadFormatIndicator(val *string) {
	if err := j.validateSetPayloadFormatIndicatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"payloadFormatIndicator",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionRepublishHeadersOutputReference)SetResponseTopic(val *string) {
	if err := j.validateSetResponseTopicParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"responseTopic",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionRepublishHeadersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionRepublishHeadersOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionRepublishHeadersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionRepublishHeadersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionRepublishHeadersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionRepublishHeadersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionRepublishHeadersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionRepublishHeadersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionRepublishHeadersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionRepublishHeadersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionRepublishHeadersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionRepublishHeadersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionRepublishHeadersOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionRepublishHeadersOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionRepublishHeadersOutputReference) PutUserProperties(value interface{}) {
	if err := i.validatePutUserPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putUserProperties",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionRepublishHeadersOutputReference) ResetContentType() {
	_jsii_.InvokeVoid(
		i,
		"resetContentType",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionRepublishHeadersOutputReference) ResetCorrelationData() {
	_jsii_.InvokeVoid(
		i,
		"resetCorrelationData",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionRepublishHeadersOutputReference) ResetMessageExpiry() {
	_jsii_.InvokeVoid(
		i,
		"resetMessageExpiry",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionRepublishHeadersOutputReference) ResetPayloadFormatIndicator() {
	_jsii_.InvokeVoid(
		i,
		"resetPayloadFormatIndicator",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionRepublishHeadersOutputReference) ResetResponseTopic() {
	_jsii_.InvokeVoid(
		i,
		"resetResponseTopic",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionRepublishHeadersOutputReference) ResetUserProperties() {
	_jsii_.InvokeVoid(
		i,
		"resetUserProperties",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionRepublishHeadersOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := i.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionRepublishHeadersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

