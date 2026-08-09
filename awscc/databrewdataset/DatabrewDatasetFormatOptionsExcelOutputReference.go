// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package databrewdataset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/databrewdataset/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DatabrewDatasetFormatOptionsExcelOutputReference interface {
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
	HeaderRow() interface{}
	SetHeaderRow(val interface{})
	HeaderRowInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	SheetIndexes() *[]*float64
	SetSheetIndexes(val *[]*float64)
	SheetIndexesInput() *[]*float64
	SheetNames() *[]*string
	SetSheetNames(val *[]*string)
	SheetNamesInput() *[]*string
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
	ResetHeaderRow()
	ResetSheetIndexes()
	ResetSheetNames()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DatabrewDatasetFormatOptionsExcelOutputReference
type jsiiProxy_DatabrewDatasetFormatOptionsExcelOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DatabrewDatasetFormatOptionsExcelOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewDatasetFormatOptionsExcelOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewDatasetFormatOptionsExcelOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewDatasetFormatOptionsExcelOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewDatasetFormatOptionsExcelOutputReference) HeaderRow() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"headerRow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewDatasetFormatOptionsExcelOutputReference) HeaderRowInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"headerRowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewDatasetFormatOptionsExcelOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewDatasetFormatOptionsExcelOutputReference) SheetIndexes() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"sheetIndexes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewDatasetFormatOptionsExcelOutputReference) SheetIndexesInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"sheetIndexesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewDatasetFormatOptionsExcelOutputReference) SheetNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sheetNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewDatasetFormatOptionsExcelOutputReference) SheetNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sheetNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewDatasetFormatOptionsExcelOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewDatasetFormatOptionsExcelOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDatabrewDatasetFormatOptionsExcelOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DatabrewDatasetFormatOptionsExcelOutputReference {
	_init_.Initialize()

	if err := validateNewDatabrewDatasetFormatOptionsExcelOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatabrewDatasetFormatOptionsExcelOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.databrewDataset.DatabrewDatasetFormatOptionsExcelOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDatabrewDatasetFormatOptionsExcelOutputReference_Override(d DatabrewDatasetFormatOptionsExcelOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.databrewDataset.DatabrewDatasetFormatOptionsExcelOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DatabrewDatasetFormatOptionsExcelOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DatabrewDatasetFormatOptionsExcelOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DatabrewDatasetFormatOptionsExcelOutputReference)SetHeaderRow(val interface{}) {
	if err := j.validateSetHeaderRowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"headerRow",
		val,
	)
}

func (j *jsiiProxy_DatabrewDatasetFormatOptionsExcelOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatabrewDatasetFormatOptionsExcelOutputReference)SetSheetIndexes(val *[]*float64) {
	if err := j.validateSetSheetIndexesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sheetIndexes",
		val,
	)
}

func (j *jsiiProxy_DatabrewDatasetFormatOptionsExcelOutputReference)SetSheetNames(val *[]*string) {
	if err := j.validateSetSheetNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sheetNames",
		val,
	)
}

func (j *jsiiProxy_DatabrewDatasetFormatOptionsExcelOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatabrewDatasetFormatOptionsExcelOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DatabrewDatasetFormatOptionsExcelOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabrewDatasetFormatOptionsExcelOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DatabrewDatasetFormatOptionsExcelOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DatabrewDatasetFormatOptionsExcelOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DatabrewDatasetFormatOptionsExcelOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DatabrewDatasetFormatOptionsExcelOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DatabrewDatasetFormatOptionsExcelOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DatabrewDatasetFormatOptionsExcelOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DatabrewDatasetFormatOptionsExcelOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DatabrewDatasetFormatOptionsExcelOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DatabrewDatasetFormatOptionsExcelOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabrewDatasetFormatOptionsExcelOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DatabrewDatasetFormatOptionsExcelOutputReference) ResetHeaderRow() {
	_jsii_.InvokeVoid(
		d,
		"resetHeaderRow",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabrewDatasetFormatOptionsExcelOutputReference) ResetSheetIndexes() {
	_jsii_.InvokeVoid(
		d,
		"resetSheetIndexes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabrewDatasetFormatOptionsExcelOutputReference) ResetSheetNames() {
	_jsii_.InvokeVoid(
		d,
		"resetSheetNames",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabrewDatasetFormatOptionsExcelOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DatabrewDatasetFormatOptionsExcelOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

