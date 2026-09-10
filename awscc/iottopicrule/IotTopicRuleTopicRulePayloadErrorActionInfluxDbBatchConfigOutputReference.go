// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iottopicrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/iottopicrule/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IotTopicRuleTopicRulePayloadErrorActionInfluxDbBatchConfigOutputReference interface {
	cdktn.ComplexObject
	BatchAcrossTopics() interface{}
	SetBatchAcrossTopics(val interface{})
	BatchAcrossTopicsInput() interface{}
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
	MaxBatchOpenMs() *float64
	SetMaxBatchOpenMs(val *float64)
	MaxBatchOpenMsInput() *float64
	MaxBatchSize() *float64
	SetMaxBatchSize(val *float64)
	MaxBatchSizeBytes() *float64
	SetMaxBatchSizeBytes(val *float64)
	MaxBatchSizeBytesInput() *float64
	MaxBatchSizeInput() *float64
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
	ResetBatchAcrossTopics()
	ResetMaxBatchOpenMs()
	ResetMaxBatchSize()
	ResetMaxBatchSizeBytes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IotTopicRuleTopicRulePayloadErrorActionInfluxDbBatchConfigOutputReference
type jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionInfluxDbBatchConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionInfluxDbBatchConfigOutputReference) BatchAcrossTopics() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"batchAcrossTopics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionInfluxDbBatchConfigOutputReference) BatchAcrossTopicsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"batchAcrossTopicsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionInfluxDbBatchConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionInfluxDbBatchConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionInfluxDbBatchConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionInfluxDbBatchConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionInfluxDbBatchConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionInfluxDbBatchConfigOutputReference) MaxBatchOpenMs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBatchOpenMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionInfluxDbBatchConfigOutputReference) MaxBatchOpenMsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBatchOpenMsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionInfluxDbBatchConfigOutputReference) MaxBatchSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBatchSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionInfluxDbBatchConfigOutputReference) MaxBatchSizeBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBatchSizeBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionInfluxDbBatchConfigOutputReference) MaxBatchSizeBytesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBatchSizeBytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionInfluxDbBatchConfigOutputReference) MaxBatchSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBatchSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionInfluxDbBatchConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionInfluxDbBatchConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewIotTopicRuleTopicRulePayloadErrorActionInfluxDbBatchConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) IotTopicRuleTopicRulePayloadErrorActionInfluxDbBatchConfigOutputReference {
	_init_.Initialize()

	if err := validateNewIotTopicRuleTopicRulePayloadErrorActionInfluxDbBatchConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionInfluxDbBatchConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.iotTopicRule.IotTopicRuleTopicRulePayloadErrorActionInfluxDbBatchConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIotTopicRuleTopicRulePayloadErrorActionInfluxDbBatchConfigOutputReference_Override(i IotTopicRuleTopicRulePayloadErrorActionInfluxDbBatchConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.iotTopicRule.IotTopicRuleTopicRulePayloadErrorActionInfluxDbBatchConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionInfluxDbBatchConfigOutputReference)SetBatchAcrossTopics(val interface{}) {
	if err := j.validateSetBatchAcrossTopicsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"batchAcrossTopics",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionInfluxDbBatchConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionInfluxDbBatchConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionInfluxDbBatchConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionInfluxDbBatchConfigOutputReference)SetMaxBatchOpenMs(val *float64) {
	if err := j.validateSetMaxBatchOpenMsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxBatchOpenMs",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionInfluxDbBatchConfigOutputReference)SetMaxBatchSize(val *float64) {
	if err := j.validateSetMaxBatchSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxBatchSize",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionInfluxDbBatchConfigOutputReference)SetMaxBatchSizeBytes(val *float64) {
	if err := j.validateSetMaxBatchSizeBytesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxBatchSizeBytes",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionInfluxDbBatchConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionInfluxDbBatchConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionInfluxDbBatchConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionInfluxDbBatchConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionInfluxDbBatchConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionInfluxDbBatchConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionInfluxDbBatchConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionInfluxDbBatchConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionInfluxDbBatchConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionInfluxDbBatchConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionInfluxDbBatchConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionInfluxDbBatchConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionInfluxDbBatchConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionInfluxDbBatchConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionInfluxDbBatchConfigOutputReference) ResetBatchAcrossTopics() {
	_jsii_.InvokeVoid(
		i,
		"resetBatchAcrossTopics",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionInfluxDbBatchConfigOutputReference) ResetMaxBatchOpenMs() {
	_jsii_.InvokeVoid(
		i,
		"resetMaxBatchOpenMs",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionInfluxDbBatchConfigOutputReference) ResetMaxBatchSize() {
	_jsii_.InvokeVoid(
		i,
		"resetMaxBatchSize",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionInfluxDbBatchConfigOutputReference) ResetMaxBatchSizeBytes() {
	_jsii_.InvokeVoid(
		i,
		"resetMaxBatchSizeBytes",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionInfluxDbBatchConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionInfluxDbBatchConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

