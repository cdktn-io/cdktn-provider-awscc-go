// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksighttopic

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/quicksighttopic/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference interface {
	cdktn.ComplexObject
	BlankCellFormat() *string
	SetBlankCellFormat(val *string)
	BlankCellFormatInput() *string
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
	SetCurrencySymbol(val *string)
	CurrencySymbolInput() *string
	DateFormat() *string
	SetDateFormat(val *string)
	DateFormatInput() *string
	DecimalSeparator() *string
	SetDecimalSeparator(val *string)
	DecimalSeparatorInput() *string
	// Experimental.
	Fqn() *string
	FractionDigits() *float64
	SetFractionDigits(val *float64)
	FractionDigitsInput() *float64
	GroupingSeparator() *string
	SetGroupingSeparator(val *string)
	GroupingSeparatorInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	NegativeFormat() QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsNegativeFormatOutputReference
	NegativeFormatInput() interface{}
	Prefix() *string
	SetPrefix(val *string)
	PrefixInput() *string
	Suffix() *string
	SetSuffix(val *string)
	SuffixInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UnitScaler() *string
	SetUnitScaler(val *string)
	UnitScalerInput() *string
	UseBlankCellFormat() interface{}
	SetUseBlankCellFormat(val interface{})
	UseBlankCellFormatInput() interface{}
	UseGrouping() interface{}
	SetUseGrouping(val interface{})
	UseGroupingInput() interface{}
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
	PutNegativeFormat(value *QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsNegativeFormat)
	ResetBlankCellFormat()
	ResetCurrencySymbol()
	ResetDateFormat()
	ResetDecimalSeparator()
	ResetFractionDigits()
	ResetGroupingSeparator()
	ResetNegativeFormat()
	ResetPrefix()
	ResetSuffix()
	ResetUnitScaler()
	ResetUseBlankCellFormat()
	ResetUseGrouping()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference
type jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference) BlankCellFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blankCellFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference) BlankCellFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blankCellFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference) CurrencySymbol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"currencySymbol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference) CurrencySymbolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"currencySymbolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference) DateFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dateFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference) DateFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dateFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference) DecimalSeparator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"decimalSeparator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference) DecimalSeparatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"decimalSeparatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference) FractionDigits() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fractionDigits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference) FractionDigitsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fractionDigitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference) GroupingSeparator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupingSeparator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference) GroupingSeparatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupingSeparatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference) NegativeFormat() QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsNegativeFormatOutputReference {
	var returns QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsNegativeFormatOutputReference
	_jsii_.Get(
		j,
		"negativeFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference) NegativeFormatInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negativeFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference) Prefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference) PrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference) Suffix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"suffix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference) SuffixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"suffixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference) UnitScaler() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unitScaler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference) UnitScalerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unitScalerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference) UseBlankCellFormat() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useBlankCellFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference) UseBlankCellFormatInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useBlankCellFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference) UseGrouping() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useGrouping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference) UseGroupingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useGroupingInput",
		&returns,
	)
	return returns
}


func NewQuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewQuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.quicksightTopic.QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewQuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference_Override(q QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.quicksightTopic.QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		q,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference)SetBlankCellFormat(val *string) {
	if err := j.validateSetBlankCellFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blankCellFormat",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference)SetCurrencySymbol(val *string) {
	if err := j.validateSetCurrencySymbolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"currencySymbol",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference)SetDateFormat(val *string) {
	if err := j.validateSetDateFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dateFormat",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference)SetDecimalSeparator(val *string) {
	if err := j.validateSetDecimalSeparatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"decimalSeparator",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference)SetFractionDigits(val *float64) {
	if err := j.validateSetFractionDigitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fractionDigits",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference)SetGroupingSeparator(val *string) {
	if err := j.validateSetGroupingSeparatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupingSeparator",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference)SetPrefix(val *string) {
	if err := j.validateSetPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prefix",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference)SetSuffix(val *string) {
	if err := j.validateSetSuffixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"suffix",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference)SetUnitScaler(val *string) {
	if err := j.validateSetUnitScalerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unitScaler",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference)SetUseBlankCellFormat(val interface{}) {
	if err := j.validateSetUseBlankCellFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useBlankCellFormat",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference)SetUseGrouping(val interface{}) {
	if err := j.validateSetUseGroupingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useGrouping",
		val,
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := q.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		q,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := q.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		q,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := q.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		q,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := q.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		q,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := q.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		q,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := q.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		q,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := q.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		q,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := q.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		q,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := q.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		q,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := q.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference) PutNegativeFormat(value *QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsNegativeFormat) {
	if err := q.validatePutNegativeFormatParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putNegativeFormat",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference) ResetBlankCellFormat() {
	_jsii_.InvokeVoid(
		q,
		"resetBlankCellFormat",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference) ResetCurrencySymbol() {
	_jsii_.InvokeVoid(
		q,
		"resetCurrencySymbol",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference) ResetDateFormat() {
	_jsii_.InvokeVoid(
		q,
		"resetDateFormat",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference) ResetDecimalSeparator() {
	_jsii_.InvokeVoid(
		q,
		"resetDecimalSeparator",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference) ResetFractionDigits() {
	_jsii_.InvokeVoid(
		q,
		"resetFractionDigits",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference) ResetGroupingSeparator() {
	_jsii_.InvokeVoid(
		q,
		"resetGroupingSeparator",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference) ResetNegativeFormat() {
	_jsii_.InvokeVoid(
		q,
		"resetNegativeFormat",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference) ResetPrefix() {
	_jsii_.InvokeVoid(
		q,
		"resetPrefix",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference) ResetSuffix() {
	_jsii_.InvokeVoid(
		q,
		"resetSuffix",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference) ResetUnitScaler() {
	_jsii_.InvokeVoid(
		q,
		"resetUnitScaler",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference) ResetUseBlankCellFormat() {
	_jsii_.InvokeVoid(
		q,
		"resetUseBlankCellFormat",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference) ResetUseGrouping() {
	_jsii_.InvokeVoid(
		q,
		"resetUseGrouping",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := q.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		q,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingDisplayFormatOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

