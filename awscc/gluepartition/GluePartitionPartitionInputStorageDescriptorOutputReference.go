// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluepartition

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/gluepartition/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GluePartitionPartitionInputStorageDescriptorOutputReference interface {
	cdktn.ComplexObject
	BucketColumns() *[]*string
	SetBucketColumns(val *[]*string)
	BucketColumnsInput() *[]*string
	Columns() GluePartitionPartitionInputStorageDescriptorColumnsList
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
	SchemaReference() GluePartitionPartitionInputStorageDescriptorSchemaReferenceOutputReference
	SchemaReferenceInput() interface{}
	SerdeInfo() GluePartitionPartitionInputStorageDescriptorSerdeInfoOutputReference
	SerdeInfoInput() interface{}
	SkewedInfo() GluePartitionPartitionInputStorageDescriptorSkewedInfoOutputReference
	SkewedInfoInput() interface{}
	SortColumns() GluePartitionPartitionInputStorageDescriptorSortColumnsList
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
	PutSchemaReference(value *GluePartitionPartitionInputStorageDescriptorSchemaReference)
	PutSerdeInfo(value *GluePartitionPartitionInputStorageDescriptorSerdeInfo)
	PutSkewedInfo(value *GluePartitionPartitionInputStorageDescriptorSkewedInfo)
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

// The jsii proxy struct for GluePartitionPartitionInputStorageDescriptorOutputReference
type jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference) BucketColumns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"bucketColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference) BucketColumnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"bucketColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference) Columns() GluePartitionPartitionInputStorageDescriptorColumnsList {
	var returns GluePartitionPartitionInputStorageDescriptorColumnsList
	_jsii_.Get(
		j,
		"columns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference) ColumnsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"columnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference) Compressed() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"compressed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference) CompressedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"compressedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference) InputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference) InputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference) NumberOfBuckets() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfBuckets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference) NumberOfBucketsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfBucketsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference) OutputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference) OutputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference) Parameters() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference) ParametersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference) SchemaReference() GluePartitionPartitionInputStorageDescriptorSchemaReferenceOutputReference {
	var returns GluePartitionPartitionInputStorageDescriptorSchemaReferenceOutputReference
	_jsii_.Get(
		j,
		"schemaReference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference) SchemaReferenceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"schemaReferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference) SerdeInfo() GluePartitionPartitionInputStorageDescriptorSerdeInfoOutputReference {
	var returns GluePartitionPartitionInputStorageDescriptorSerdeInfoOutputReference
	_jsii_.Get(
		j,
		"serdeInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference) SerdeInfoInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serdeInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference) SkewedInfo() GluePartitionPartitionInputStorageDescriptorSkewedInfoOutputReference {
	var returns GluePartitionPartitionInputStorageDescriptorSkewedInfoOutputReference
	_jsii_.Get(
		j,
		"skewedInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference) SkewedInfoInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skewedInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference) SortColumns() GluePartitionPartitionInputStorageDescriptorSortColumnsList {
	var returns GluePartitionPartitionInputStorageDescriptorSortColumnsList
	_jsii_.Get(
		j,
		"sortColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference) SortColumnsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sortColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference) StoredAsSubDirectories() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storedAsSubDirectories",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference) StoredAsSubDirectoriesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storedAsSubDirectoriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGluePartitionPartitionInputStorageDescriptorOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GluePartitionPartitionInputStorageDescriptorOutputReference {
	_init_.Initialize()

	if err := validateNewGluePartitionPartitionInputStorageDescriptorOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.gluePartition.GluePartitionPartitionInputStorageDescriptorOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGluePartitionPartitionInputStorageDescriptorOutputReference_Override(g GluePartitionPartitionInputStorageDescriptorOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.gluePartition.GluePartitionPartitionInputStorageDescriptorOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference)SetBucketColumns(val *[]*string) {
	if err := j.validateSetBucketColumnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucketColumns",
		val,
	)
}

func (j *jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference)SetCompressed(val interface{}) {
	if err := j.validateSetCompressedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compressed",
		val,
	)
}

func (j *jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference)SetInputFormat(val *string) {
	if err := j.validateSetInputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputFormat",
		val,
	)
}

func (j *jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference)SetNumberOfBuckets(val *float64) {
	if err := j.validateSetNumberOfBucketsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numberOfBuckets",
		val,
	)
}

func (j *jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference)SetOutputFormat(val *string) {
	if err := j.validateSetOutputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputFormat",
		val,
	)
}

func (j *jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference)SetParameters(val *string) {
	if err := j.validateSetParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference)SetStoredAsSubDirectories(val interface{}) {
	if err := j.validateSetStoredAsSubDirectoriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storedAsSubDirectories",
		val,
	)
}

func (j *jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference) PutColumns(value interface{}) {
	if err := g.validatePutColumnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putColumns",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference) PutSchemaReference(value *GluePartitionPartitionInputStorageDescriptorSchemaReference) {
	if err := g.validatePutSchemaReferenceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSchemaReference",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference) PutSerdeInfo(value *GluePartitionPartitionInputStorageDescriptorSerdeInfo) {
	if err := g.validatePutSerdeInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSerdeInfo",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference) PutSkewedInfo(value *GluePartitionPartitionInputStorageDescriptorSkewedInfo) {
	if err := g.validatePutSkewedInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSkewedInfo",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference) PutSortColumns(value interface{}) {
	if err := g.validatePutSortColumnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSortColumns",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference) ResetBucketColumns() {
	_jsii_.InvokeVoid(
		g,
		"resetBucketColumns",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference) ResetColumns() {
	_jsii_.InvokeVoid(
		g,
		"resetColumns",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference) ResetCompressed() {
	_jsii_.InvokeVoid(
		g,
		"resetCompressed",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference) ResetInputFormat() {
	_jsii_.InvokeVoid(
		g,
		"resetInputFormat",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference) ResetLocation() {
	_jsii_.InvokeVoid(
		g,
		"resetLocation",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference) ResetNumberOfBuckets() {
	_jsii_.InvokeVoid(
		g,
		"resetNumberOfBuckets",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference) ResetOutputFormat() {
	_jsii_.InvokeVoid(
		g,
		"resetOutputFormat",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference) ResetParameters() {
	_jsii_.InvokeVoid(
		g,
		"resetParameters",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference) ResetSchemaReference() {
	_jsii_.InvokeVoid(
		g,
		"resetSchemaReference",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference) ResetSerdeInfo() {
	_jsii_.InvokeVoid(
		g,
		"resetSerdeInfo",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference) ResetSkewedInfo() {
	_jsii_.InvokeVoid(
		g,
		"resetSkewedInfo",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference) ResetSortColumns() {
	_jsii_.InvokeVoid(
		g,
		"resetSortColumns",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference) ResetStoredAsSubDirectories() {
	_jsii_.InvokeVoid(
		g,
		"resetStoredAsSubDirectories",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GluePartitionPartitionInputStorageDescriptorOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

