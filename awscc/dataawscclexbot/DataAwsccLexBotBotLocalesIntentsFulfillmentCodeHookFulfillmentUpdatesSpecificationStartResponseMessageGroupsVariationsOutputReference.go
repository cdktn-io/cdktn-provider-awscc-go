// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawscclexbot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawscclexbot/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupsVariationsOutputReference interface {
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
	CustomPayload() DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupsVariationsCustomPayloadOutputReference
	// Experimental.
	Fqn() *string
	ImageResponseCard() DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupsVariationsImageResponseCardOutputReference
	InternalValue() *DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupsVariations
	SetInternalValue(val *DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupsVariations)
	PlainTextMessage() DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupsVariationsPlainTextMessageOutputReference
	SsmlMessage() DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupsVariationsSsmlMessageOutputReference
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupsVariationsOutputReference
type jsiiProxy_DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupsVariationsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupsVariationsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupsVariationsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupsVariationsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupsVariationsOutputReference) CustomPayload() DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupsVariationsCustomPayloadOutputReference {
	var returns DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupsVariationsCustomPayloadOutputReference
	_jsii_.Get(
		j,
		"customPayload",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupsVariationsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupsVariationsOutputReference) ImageResponseCard() DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupsVariationsImageResponseCardOutputReference {
	var returns DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupsVariationsImageResponseCardOutputReference
	_jsii_.Get(
		j,
		"imageResponseCard",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupsVariationsOutputReference) InternalValue() *DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupsVariations {
	var returns *DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupsVariations
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupsVariationsOutputReference) PlainTextMessage() DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupsVariationsPlainTextMessageOutputReference {
	var returns DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupsVariationsPlainTextMessageOutputReference
	_jsii_.Get(
		j,
		"plainTextMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupsVariationsOutputReference) SsmlMessage() DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupsVariationsSsmlMessageOutputReference {
	var returns DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupsVariationsSsmlMessageOutputReference
	_jsii_.Get(
		j,
		"ssmlMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupsVariationsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupsVariationsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupsVariationsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupsVariationsOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupsVariationsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupsVariationsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccLexBot.DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupsVariationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupsVariationsOutputReference_Override(d DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupsVariationsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccLexBot.DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupsVariationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupsVariationsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupsVariationsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupsVariationsOutputReference)SetInternalValue(val *DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupsVariations) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupsVariationsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupsVariationsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupsVariationsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupsVariationsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupsVariationsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupsVariationsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupsVariationsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupsVariationsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupsVariationsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupsVariationsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupsVariationsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupsVariationsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupsVariationsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupsVariationsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupsVariationsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupsVariationsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

