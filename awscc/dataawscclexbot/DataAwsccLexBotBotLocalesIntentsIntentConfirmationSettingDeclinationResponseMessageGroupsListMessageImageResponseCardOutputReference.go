// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawscclexbot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawscclexbot/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccLexBotBotLocalesIntentsIntentConfirmationSettingDeclinationResponseMessageGroupsListMessageImageResponseCardOutputReference interface {
	cdktn.ComplexObject
	Buttons() DataAwsccLexBotBotLocalesIntentsIntentConfirmationSettingDeclinationResponseMessageGroupsListMessageImageResponseCardButtonsList
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
	ImageUrl() *string
	InternalValue() *DataAwsccLexBotBotLocalesIntentsIntentConfirmationSettingDeclinationResponseMessageGroupsListMessageImageResponseCard
	SetInternalValue(val *DataAwsccLexBotBotLocalesIntentsIntentConfirmationSettingDeclinationResponseMessageGroupsListMessageImageResponseCard)
	Subtitle() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Title() *string
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

// The jsii proxy struct for DataAwsccLexBotBotLocalesIntentsIntentConfirmationSettingDeclinationResponseMessageGroupsListMessageImageResponseCardOutputReference
type jsiiProxy_DataAwsccLexBotBotLocalesIntentsIntentConfirmationSettingDeclinationResponseMessageGroupsListMessageImageResponseCardOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccLexBotBotLocalesIntentsIntentConfirmationSettingDeclinationResponseMessageGroupsListMessageImageResponseCardOutputReference) Buttons() DataAwsccLexBotBotLocalesIntentsIntentConfirmationSettingDeclinationResponseMessageGroupsListMessageImageResponseCardButtonsList {
	var returns DataAwsccLexBotBotLocalesIntentsIntentConfirmationSettingDeclinationResponseMessageGroupsListMessageImageResponseCardButtonsList
	_jsii_.Get(
		j,
		"buttons",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccLexBotBotLocalesIntentsIntentConfirmationSettingDeclinationResponseMessageGroupsListMessageImageResponseCardOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccLexBotBotLocalesIntentsIntentConfirmationSettingDeclinationResponseMessageGroupsListMessageImageResponseCardOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccLexBotBotLocalesIntentsIntentConfirmationSettingDeclinationResponseMessageGroupsListMessageImageResponseCardOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccLexBotBotLocalesIntentsIntentConfirmationSettingDeclinationResponseMessageGroupsListMessageImageResponseCardOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccLexBotBotLocalesIntentsIntentConfirmationSettingDeclinationResponseMessageGroupsListMessageImageResponseCardOutputReference) ImageUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccLexBotBotLocalesIntentsIntentConfirmationSettingDeclinationResponseMessageGroupsListMessageImageResponseCardOutputReference) InternalValue() *DataAwsccLexBotBotLocalesIntentsIntentConfirmationSettingDeclinationResponseMessageGroupsListMessageImageResponseCard {
	var returns *DataAwsccLexBotBotLocalesIntentsIntentConfirmationSettingDeclinationResponseMessageGroupsListMessageImageResponseCard
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccLexBotBotLocalesIntentsIntentConfirmationSettingDeclinationResponseMessageGroupsListMessageImageResponseCardOutputReference) Subtitle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subtitle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccLexBotBotLocalesIntentsIntentConfirmationSettingDeclinationResponseMessageGroupsListMessageImageResponseCardOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccLexBotBotLocalesIntentsIntentConfirmationSettingDeclinationResponseMessageGroupsListMessageImageResponseCardOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccLexBotBotLocalesIntentsIntentConfirmationSettingDeclinationResponseMessageGroupsListMessageImageResponseCardOutputReference) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}


func NewDataAwsccLexBotBotLocalesIntentsIntentConfirmationSettingDeclinationResponseMessageGroupsListMessageImageResponseCardOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataAwsccLexBotBotLocalesIntentsIntentConfirmationSettingDeclinationResponseMessageGroupsListMessageImageResponseCardOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccLexBotBotLocalesIntentsIntentConfirmationSettingDeclinationResponseMessageGroupsListMessageImageResponseCardOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccLexBotBotLocalesIntentsIntentConfirmationSettingDeclinationResponseMessageGroupsListMessageImageResponseCardOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccLexBot.DataAwsccLexBotBotLocalesIntentsIntentConfirmationSettingDeclinationResponseMessageGroupsListMessageImageResponseCardOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccLexBotBotLocalesIntentsIntentConfirmationSettingDeclinationResponseMessageGroupsListMessageImageResponseCardOutputReference_Override(d DataAwsccLexBotBotLocalesIntentsIntentConfirmationSettingDeclinationResponseMessageGroupsListMessageImageResponseCardOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccLexBot.DataAwsccLexBotBotLocalesIntentsIntentConfirmationSettingDeclinationResponseMessageGroupsListMessageImageResponseCardOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccLexBotBotLocalesIntentsIntentConfirmationSettingDeclinationResponseMessageGroupsListMessageImageResponseCardOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccLexBotBotLocalesIntentsIntentConfirmationSettingDeclinationResponseMessageGroupsListMessageImageResponseCardOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccLexBotBotLocalesIntentsIntentConfirmationSettingDeclinationResponseMessageGroupsListMessageImageResponseCardOutputReference)SetInternalValue(val *DataAwsccLexBotBotLocalesIntentsIntentConfirmationSettingDeclinationResponseMessageGroupsListMessageImageResponseCard) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccLexBotBotLocalesIntentsIntentConfirmationSettingDeclinationResponseMessageGroupsListMessageImageResponseCardOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccLexBotBotLocalesIntentsIntentConfirmationSettingDeclinationResponseMessageGroupsListMessageImageResponseCardOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccLexBotBotLocalesIntentsIntentConfirmationSettingDeclinationResponseMessageGroupsListMessageImageResponseCardOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccLexBotBotLocalesIntentsIntentConfirmationSettingDeclinationResponseMessageGroupsListMessageImageResponseCardOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccLexBotBotLocalesIntentsIntentConfirmationSettingDeclinationResponseMessageGroupsListMessageImageResponseCardOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccLexBotBotLocalesIntentsIntentConfirmationSettingDeclinationResponseMessageGroupsListMessageImageResponseCardOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccLexBotBotLocalesIntentsIntentConfirmationSettingDeclinationResponseMessageGroupsListMessageImageResponseCardOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccLexBotBotLocalesIntentsIntentConfirmationSettingDeclinationResponseMessageGroupsListMessageImageResponseCardOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccLexBotBotLocalesIntentsIntentConfirmationSettingDeclinationResponseMessageGroupsListMessageImageResponseCardOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccLexBotBotLocalesIntentsIntentConfirmationSettingDeclinationResponseMessageGroupsListMessageImageResponseCardOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccLexBotBotLocalesIntentsIntentConfirmationSettingDeclinationResponseMessageGroupsListMessageImageResponseCardOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccLexBotBotLocalesIntentsIntentConfirmationSettingDeclinationResponseMessageGroupsListMessageImageResponseCardOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccLexBotBotLocalesIntentsIntentConfirmationSettingDeclinationResponseMessageGroupsListMessageImageResponseCardOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccLexBotBotLocalesIntentsIntentConfirmationSettingDeclinationResponseMessageGroupsListMessageImageResponseCardOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccLexBotBotLocalesIntentsIntentConfirmationSettingDeclinationResponseMessageGroupsListMessageImageResponseCardOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccLexBotBotLocalesIntentsIntentConfirmationSettingDeclinationResponseMessageGroupsListMessageImageResponseCardOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

