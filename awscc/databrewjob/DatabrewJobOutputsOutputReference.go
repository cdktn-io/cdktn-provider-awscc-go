// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package databrewjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/databrewjob/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DatabrewJobOutputsOutputReference interface {
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
	CompressionFormat() *string
	SetCompressionFormat(val *string)
	CompressionFormatInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Format() *string
	SetFormat(val *string)
	FormatInput() *string
	FormatOptions() DatabrewJobOutputsFormatOptionsOutputReference
	FormatOptionsInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Location() DatabrewJobOutputsLocationOutputReference
	LocationInput() interface{}
	MaxOutputFiles() *float64
	SetMaxOutputFiles(val *float64)
	MaxOutputFilesInput() *float64
	Overwrite() interface{}
	SetOverwrite(val interface{})
	OverwriteInput() interface{}
	PartitionColumns() *[]*string
	SetPartitionColumns(val *[]*string)
	PartitionColumnsInput() *[]*string
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
	PutFormatOptions(value *DatabrewJobOutputsFormatOptions)
	PutLocation(value *DatabrewJobOutputsLocation)
	ResetCompressionFormat()
	ResetFormat()
	ResetFormatOptions()
	ResetLocation()
	ResetMaxOutputFiles()
	ResetOverwrite()
	ResetPartitionColumns()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DatabrewJobOutputsOutputReference
type jsiiProxy_DatabrewJobOutputsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DatabrewJobOutputsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJobOutputsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJobOutputsOutputReference) CompressionFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJobOutputsOutputReference) CompressionFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJobOutputsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJobOutputsOutputReference) Format() *string {
	var returns *string
	_jsii_.Get(
		j,
		"format",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJobOutputsOutputReference) FormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"formatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJobOutputsOutputReference) FormatOptions() DatabrewJobOutputsFormatOptionsOutputReference {
	var returns DatabrewJobOutputsFormatOptionsOutputReference
	_jsii_.Get(
		j,
		"formatOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJobOutputsOutputReference) FormatOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"formatOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJobOutputsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJobOutputsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJobOutputsOutputReference) Location() DatabrewJobOutputsLocationOutputReference {
	var returns DatabrewJobOutputsLocationOutputReference
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJobOutputsOutputReference) LocationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJobOutputsOutputReference) MaxOutputFiles() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxOutputFiles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJobOutputsOutputReference) MaxOutputFilesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxOutputFilesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJobOutputsOutputReference) Overwrite() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"overwrite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJobOutputsOutputReference) OverwriteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"overwriteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJobOutputsOutputReference) PartitionColumns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"partitionColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJobOutputsOutputReference) PartitionColumnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"partitionColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJobOutputsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJobOutputsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDatabrewJobOutputsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DatabrewJobOutputsOutputReference {
	_init_.Initialize()

	if err := validateNewDatabrewJobOutputsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatabrewJobOutputsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.databrewJob.DatabrewJobOutputsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDatabrewJobOutputsOutputReference_Override(d DatabrewJobOutputsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.databrewJob.DatabrewJobOutputsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DatabrewJobOutputsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DatabrewJobOutputsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DatabrewJobOutputsOutputReference)SetCompressionFormat(val *string) {
	if err := j.validateSetCompressionFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compressionFormat",
		val,
	)
}

func (j *jsiiProxy_DatabrewJobOutputsOutputReference)SetFormat(val *string) {
	if err := j.validateSetFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"format",
		val,
	)
}

func (j *jsiiProxy_DatabrewJobOutputsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatabrewJobOutputsOutputReference)SetMaxOutputFiles(val *float64) {
	if err := j.validateSetMaxOutputFilesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxOutputFiles",
		val,
	)
}

func (j *jsiiProxy_DatabrewJobOutputsOutputReference)SetOverwrite(val interface{}) {
	if err := j.validateSetOverwriteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"overwrite",
		val,
	)
}

func (j *jsiiProxy_DatabrewJobOutputsOutputReference)SetPartitionColumns(val *[]*string) {
	if err := j.validateSetPartitionColumnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"partitionColumns",
		val,
	)
}

func (j *jsiiProxy_DatabrewJobOutputsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatabrewJobOutputsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DatabrewJobOutputsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabrewJobOutputsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DatabrewJobOutputsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DatabrewJobOutputsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DatabrewJobOutputsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DatabrewJobOutputsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DatabrewJobOutputsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DatabrewJobOutputsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DatabrewJobOutputsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DatabrewJobOutputsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DatabrewJobOutputsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabrewJobOutputsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DatabrewJobOutputsOutputReference) PutFormatOptions(value *DatabrewJobOutputsFormatOptions) {
	if err := d.validatePutFormatOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFormatOptions",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabrewJobOutputsOutputReference) PutLocation(value *DatabrewJobOutputsLocation) {
	if err := d.validatePutLocationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putLocation",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabrewJobOutputsOutputReference) ResetCompressionFormat() {
	_jsii_.InvokeVoid(
		d,
		"resetCompressionFormat",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabrewJobOutputsOutputReference) ResetFormat() {
	_jsii_.InvokeVoid(
		d,
		"resetFormat",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabrewJobOutputsOutputReference) ResetFormatOptions() {
	_jsii_.InvokeVoid(
		d,
		"resetFormatOptions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabrewJobOutputsOutputReference) ResetLocation() {
	_jsii_.InvokeVoid(
		d,
		"resetLocation",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabrewJobOutputsOutputReference) ResetMaxOutputFiles() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxOutputFiles",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabrewJobOutputsOutputReference) ResetOverwrite() {
	_jsii_.InvokeVoid(
		d,
		"resetOverwrite",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabrewJobOutputsOutputReference) ResetPartitionColumns() {
	_jsii_.InvokeVoid(
		d,
		"resetPartitionColumns",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabrewJobOutputsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DatabrewJobOutputsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

