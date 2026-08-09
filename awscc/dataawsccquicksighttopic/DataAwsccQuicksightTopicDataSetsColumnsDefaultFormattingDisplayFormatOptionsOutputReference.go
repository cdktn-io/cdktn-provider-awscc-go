// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccquicksighttopic

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawsccquicksighttopic/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccQuicksightTopicDataSetsColumnsDefaultFormattingDisplayFormatOptionsOutputReference interface {
	cdktn.ComplexObject
	BlankCellFormat() *string
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
	CurrencySymbol() *string
	DateFormat() *string
	DecimalSeparator() *string
	// Experimental.
	Fqn() *string
	FractionDigits() *float64
	GroupingSeparator() *string
	InternalValue() *DataAwsccQuicksightTopicDataSetsColumnsDefaultFormattingDisplayFormatOptions
	SetInternalValue(val *DataAwsccQuicksightTopicDataSetsColumnsDefaultFormattingDisplayFormatOptions)
	NegativeFormat() DataAwsccQuicksightTopicDataSetsColumnsDefaultFormattingDisplayFormatOptionsNegativeFormatOutputReference
	Prefix() *string
	Suffix() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UnitScaler() *string
	UseBlankCellFormat() cdktn.IResolvable
	UseGrouping() cdktn.IResolvable
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

// The jsii proxy struct for DataAwsccQuicksightTopicDataSetsColumnsDefaultFormattingDisplayFormatOptionsOutputReference
type jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsDefaultFormattingDisplayFormatOptionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsDefaultFormattingDisplayFormatOptionsOutputReference) BlankCellFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blankCellFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsDefaultFormattingDisplayFormatOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsDefaultFormattingDisplayFormatOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsDefaultFormattingDisplayFormatOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsDefaultFormattingDisplayFormatOptionsOutputReference) CurrencySymbol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"currencySymbol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsDefaultFormattingDisplayFormatOptionsOutputReference) DateFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dateFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsDefaultFormattingDisplayFormatOptionsOutputReference) DecimalSeparator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"decimalSeparator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsDefaultFormattingDisplayFormatOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsDefaultFormattingDisplayFormatOptionsOutputReference) FractionDigits() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fractionDigits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsDefaultFormattingDisplayFormatOptionsOutputReference) GroupingSeparator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupingSeparator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsDefaultFormattingDisplayFormatOptionsOutputReference) InternalValue() *DataAwsccQuicksightTopicDataSetsColumnsDefaultFormattingDisplayFormatOptions {
	var returns *DataAwsccQuicksightTopicDataSetsColumnsDefaultFormattingDisplayFormatOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsDefaultFormattingDisplayFormatOptionsOutputReference) NegativeFormat() DataAwsccQuicksightTopicDataSetsColumnsDefaultFormattingDisplayFormatOptionsNegativeFormatOutputReference {
	var returns DataAwsccQuicksightTopicDataSetsColumnsDefaultFormattingDisplayFormatOptionsNegativeFormatOutputReference
	_jsii_.Get(
		j,
		"negativeFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsDefaultFormattingDisplayFormatOptionsOutputReference) Prefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsDefaultFormattingDisplayFormatOptionsOutputReference) Suffix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"suffix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsDefaultFormattingDisplayFormatOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsDefaultFormattingDisplayFormatOptionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsDefaultFormattingDisplayFormatOptionsOutputReference) UnitScaler() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unitScaler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsDefaultFormattingDisplayFormatOptionsOutputReference) UseBlankCellFormat() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"useBlankCellFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsDefaultFormattingDisplayFormatOptionsOutputReference) UseGrouping() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"useGrouping",
		&returns,
	)
	return returns
}


func NewDataAwsccQuicksightTopicDataSetsColumnsDefaultFormattingDisplayFormatOptionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataAwsccQuicksightTopicDataSetsColumnsDefaultFormattingDisplayFormatOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccQuicksightTopicDataSetsColumnsDefaultFormattingDisplayFormatOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsDefaultFormattingDisplayFormatOptionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccQuicksightTopic.DataAwsccQuicksightTopicDataSetsColumnsDefaultFormattingDisplayFormatOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccQuicksightTopicDataSetsColumnsDefaultFormattingDisplayFormatOptionsOutputReference_Override(d DataAwsccQuicksightTopicDataSetsColumnsDefaultFormattingDisplayFormatOptionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccQuicksightTopic.DataAwsccQuicksightTopicDataSetsColumnsDefaultFormattingDisplayFormatOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsDefaultFormattingDisplayFormatOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsDefaultFormattingDisplayFormatOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsDefaultFormattingDisplayFormatOptionsOutputReference)SetInternalValue(val *DataAwsccQuicksightTopicDataSetsColumnsDefaultFormattingDisplayFormatOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsDefaultFormattingDisplayFormatOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsDefaultFormattingDisplayFormatOptionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsDefaultFormattingDisplayFormatOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsDefaultFormattingDisplayFormatOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsDefaultFormattingDisplayFormatOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsDefaultFormattingDisplayFormatOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsDefaultFormattingDisplayFormatOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsDefaultFormattingDisplayFormatOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsDefaultFormattingDisplayFormatOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsDefaultFormattingDisplayFormatOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsDefaultFormattingDisplayFormatOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsDefaultFormattingDisplayFormatOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsDefaultFormattingDisplayFormatOptionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsDefaultFormattingDisplayFormatOptionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsDefaultFormattingDisplayFormatOptionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsDefaultFormattingDisplayFormatOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

