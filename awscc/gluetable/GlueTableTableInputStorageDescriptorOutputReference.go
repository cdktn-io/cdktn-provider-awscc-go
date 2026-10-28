// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluetable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/gluetable/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GlueTableTableInputStorageDescriptorOutputReference interface {
	cdktn.ComplexObject
	BucketColumns() *[]*string
	SetBucketColumns(val *[]*string)
	BucketColumnsInput() *[]*string
	Columns() GlueTableTableInputStorageDescriptorColumnsList
	ColumnsInput() interface{}
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
	Compressed() interface{}
	SetCompressed(val interface{})
	CompressedInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InputFormat() *string
	SetInputFormat(val *string)
	InputFormatInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	NumberOfBuckets() *float64
	SetNumberOfBuckets(val *float64)
	NumberOfBucketsInput() *float64
	OutputFormat() *string
	SetOutputFormat(val *string)
	OutputFormatInput() *string
	Parameters() *string
	SetParameters(val *string)
	ParametersInput() *string
	SchemaReference() GlueTableTableInputStorageDescriptorSchemaReferenceOutputReference
	SchemaReferenceInput() interface{}
	SerdeInfo() GlueTableTableInputStorageDescriptorSerdeInfoOutputReference
	SerdeInfoInput() interface{}
	SkewedInfo() GlueTableTableInputStorageDescriptorSkewedInfoOutputReference
	SkewedInfoInput() interface{}
	SortColumns() GlueTableTableInputStorageDescriptorSortColumnsList
	SortColumnsInput() interface{}
	StoredAsSubDirectories() interface{}
	SetStoredAsSubDirectories(val interface{})
	StoredAsSubDirectoriesInput() interface{}
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
	PutColumns(value interface{})
	PutSchemaReference(value *GlueTableTableInputStorageDescriptorSchemaReference)
	PutSerdeInfo(value *GlueTableTableInputStorageDescriptorSerdeInfo)
	PutSkewedInfo(value *GlueTableTableInputStorageDescriptorSkewedInfo)
	PutSortColumns(value interface{})
	ResetBucketColumns()
	ResetColumns()
	ResetCompressed()
	ResetInputFormat()
	ResetLocation()
	ResetNumberOfBuckets()
	ResetOutputFormat()
	ResetParameters()
	ResetSchemaReference()
	ResetSerdeInfo()
	ResetSkewedInfo()
	ResetSortColumns()
	ResetStoredAsSubDirectories()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GlueTableTableInputStorageDescriptorOutputReference
type jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference) BucketColumns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"bucketColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference) BucketColumnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"bucketColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference) Columns() GlueTableTableInputStorageDescriptorColumnsList {
	var returns GlueTableTableInputStorageDescriptorColumnsList
	_jsii_.Get(
		j,
		"columns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference) ColumnsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"columnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference) Compressed() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"compressed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference) CompressedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"compressedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference) InputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference) InputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference) NumberOfBuckets() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfBuckets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference) NumberOfBucketsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfBucketsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference) OutputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference) OutputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference) Parameters() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference) ParametersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference) SchemaReference() GlueTableTableInputStorageDescriptorSchemaReferenceOutputReference {
	var returns GlueTableTableInputStorageDescriptorSchemaReferenceOutputReference
	_jsii_.Get(
		j,
		"schemaReference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference) SchemaReferenceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"schemaReferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference) SerdeInfo() GlueTableTableInputStorageDescriptorSerdeInfoOutputReference {
	var returns GlueTableTableInputStorageDescriptorSerdeInfoOutputReference
	_jsii_.Get(
		j,
		"serdeInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference) SerdeInfoInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serdeInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference) SkewedInfo() GlueTableTableInputStorageDescriptorSkewedInfoOutputReference {
	var returns GlueTableTableInputStorageDescriptorSkewedInfoOutputReference
	_jsii_.Get(
		j,
		"skewedInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference) SkewedInfoInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skewedInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference) SortColumns() GlueTableTableInputStorageDescriptorSortColumnsList {
	var returns GlueTableTableInputStorageDescriptorSortColumnsList
	_jsii_.Get(
		j,
		"sortColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference) SortColumnsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sortColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference) StoredAsSubDirectories() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storedAsSubDirectories",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference) StoredAsSubDirectoriesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storedAsSubDirectoriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGlueTableTableInputStorageDescriptorOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GlueTableTableInputStorageDescriptorOutputReference {
	_init_.Initialize()

	if err := validateNewGlueTableTableInputStorageDescriptorOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.glueTable.GlueTableTableInputStorageDescriptorOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGlueTableTableInputStorageDescriptorOutputReference_Override(g GlueTableTableInputStorageDescriptorOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.glueTable.GlueTableTableInputStorageDescriptorOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference)SetBucketColumns(val *[]*string) {
	if err := j.validateSetBucketColumnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucketColumns",
		val,
	)
}

func (j *jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference)SetCompressed(val interface{}) {
	if err := j.validateSetCompressedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compressed",
		val,
	)
}

func (j *jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference)SetInputFormat(val *string) {
	if err := j.validateSetInputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputFormat",
		val,
	)
}

func (j *jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference)SetNumberOfBuckets(val *float64) {
	if err := j.validateSetNumberOfBucketsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numberOfBuckets",
		val,
	)
}

func (j *jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference)SetOutputFormat(val *string) {
	if err := j.validateSetOutputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputFormat",
		val,
	)
}

func (j *jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference)SetParameters(val *string) {
	if err := j.validateSetParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference)SetStoredAsSubDirectories(val interface{}) {
	if err := j.validateSetStoredAsSubDirectoriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storedAsSubDirectories",
		val,
	)
}

func (j *jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference) PutColumns(value interface{}) {
	if err := g.validatePutColumnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putColumns",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference) PutSchemaReference(value *GlueTableTableInputStorageDescriptorSchemaReference) {
	if err := g.validatePutSchemaReferenceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSchemaReference",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference) PutSerdeInfo(value *GlueTableTableInputStorageDescriptorSerdeInfo) {
	if err := g.validatePutSerdeInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSerdeInfo",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference) PutSkewedInfo(value *GlueTableTableInputStorageDescriptorSkewedInfo) {
	if err := g.validatePutSkewedInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSkewedInfo",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference) PutSortColumns(value interface{}) {
	if err := g.validatePutSortColumnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSortColumns",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference) ResetBucketColumns() {
	_jsii_.InvokeVoid(
		g,
		"resetBucketColumns",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference) ResetColumns() {
	_jsii_.InvokeVoid(
		g,
		"resetColumns",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference) ResetCompressed() {
	_jsii_.InvokeVoid(
		g,
		"resetCompressed",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference) ResetInputFormat() {
	_jsii_.InvokeVoid(
		g,
		"resetInputFormat",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference) ResetLocation() {
	_jsii_.InvokeVoid(
		g,
		"resetLocation",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference) ResetNumberOfBuckets() {
	_jsii_.InvokeVoid(
		g,
		"resetNumberOfBuckets",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference) ResetOutputFormat() {
	_jsii_.InvokeVoid(
		g,
		"resetOutputFormat",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference) ResetParameters() {
	_jsii_.InvokeVoid(
		g,
		"resetParameters",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference) ResetSchemaReference() {
	_jsii_.InvokeVoid(
		g,
		"resetSchemaReference",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference) ResetSerdeInfo() {
	_jsii_.InvokeVoid(
		g,
		"resetSerdeInfo",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference) ResetSkewedInfo() {
	_jsii_.InvokeVoid(
		g,
		"resetSkewedInfo",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference) ResetSortColumns() {
	_jsii_.InvokeVoid(
		g,
		"resetSortColumns",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference) ResetStoredAsSubDirectories() {
	_jsii_.InvokeVoid(
		g,
		"resetStoredAsSubDirectories",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := g.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueTableTableInputStorageDescriptorOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

