// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dynamodbtable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dynamodbtable/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DynamodbTableVectorIndexesOutputReference interface {
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
	Dimensions() *float64
	SetDimensions(val *float64)
	DimensionsInput() *float64
	DistanceFunction() *string
	SetDistanceFunction(val *string)
	DistanceFunctionInput() *string
	// Experimental.
	Fqn() *string
	IndexName() *string
	SetIndexName(val *string)
	IndexNameInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Projection() DynamodbTableVectorIndexesProjectionOutputReference
	ProjectionInput() interface{}
	SearchSchema() DynamodbTableVectorIndexesSearchSchemaList
	SearchSchemaInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	VectorAttribute() DynamodbTableVectorIndexesVectorAttributeOutputReference
	VectorAttributeInput() interface{}
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
	PutProjection(value *DynamodbTableVectorIndexesProjection)
	PutSearchSchema(value interface{})
	PutVectorAttribute(value *DynamodbTableVectorIndexesVectorAttribute)
	ResetDimensions()
	ResetDistanceFunction()
	ResetIndexName()
	ResetProjection()
	ResetSearchSchema()
	ResetVectorAttribute()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DynamodbTableVectorIndexesOutputReference
type jsiiProxy_DynamodbTableVectorIndexesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DynamodbTableVectorIndexesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableVectorIndexesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableVectorIndexesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableVectorIndexesOutputReference) Dimensions() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dimensions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableVectorIndexesOutputReference) DimensionsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dimensionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableVectorIndexesOutputReference) DistanceFunction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"distanceFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableVectorIndexesOutputReference) DistanceFunctionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"distanceFunctionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableVectorIndexesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableVectorIndexesOutputReference) IndexName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableVectorIndexesOutputReference) IndexNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableVectorIndexesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableVectorIndexesOutputReference) Projection() DynamodbTableVectorIndexesProjectionOutputReference {
	var returns DynamodbTableVectorIndexesProjectionOutputReference
	_jsii_.Get(
		j,
		"projection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableVectorIndexesOutputReference) ProjectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"projectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableVectorIndexesOutputReference) SearchSchema() DynamodbTableVectorIndexesSearchSchemaList {
	var returns DynamodbTableVectorIndexesSearchSchemaList
	_jsii_.Get(
		j,
		"searchSchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableVectorIndexesOutputReference) SearchSchemaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"searchSchemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableVectorIndexesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableVectorIndexesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableVectorIndexesOutputReference) VectorAttribute() DynamodbTableVectorIndexesVectorAttributeOutputReference {
	var returns DynamodbTableVectorIndexesVectorAttributeOutputReference
	_jsii_.Get(
		j,
		"vectorAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableVectorIndexesOutputReference) VectorAttributeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vectorAttributeInput",
		&returns,
	)
	return returns
}


func NewDynamodbTableVectorIndexesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DynamodbTableVectorIndexesOutputReference {
	_init_.Initialize()

	if err := validateNewDynamodbTableVectorIndexesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DynamodbTableVectorIndexesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dynamodbTable.DynamodbTableVectorIndexesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDynamodbTableVectorIndexesOutputReference_Override(d DynamodbTableVectorIndexesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dynamodbTable.DynamodbTableVectorIndexesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DynamodbTableVectorIndexesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DynamodbTableVectorIndexesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DynamodbTableVectorIndexesOutputReference)SetDimensions(val *float64) {
	if err := j.validateSetDimensionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dimensions",
		val,
	)
}

func (j *jsiiProxy_DynamodbTableVectorIndexesOutputReference)SetDistanceFunction(val *string) {
	if err := j.validateSetDistanceFunctionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"distanceFunction",
		val,
	)
}

func (j *jsiiProxy_DynamodbTableVectorIndexesOutputReference)SetIndexName(val *string) {
	if err := j.validateSetIndexNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"indexName",
		val,
	)
}

func (j *jsiiProxy_DynamodbTableVectorIndexesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DynamodbTableVectorIndexesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DynamodbTableVectorIndexesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DynamodbTableVectorIndexesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamodbTableVectorIndexesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DynamodbTableVectorIndexesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DynamodbTableVectorIndexesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DynamodbTableVectorIndexesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DynamodbTableVectorIndexesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DynamodbTableVectorIndexesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DynamodbTableVectorIndexesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DynamodbTableVectorIndexesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DynamodbTableVectorIndexesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DynamodbTableVectorIndexesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamodbTableVectorIndexesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DynamodbTableVectorIndexesOutputReference) PutProjection(value *DynamodbTableVectorIndexesProjection) {
	if err := d.validatePutProjectionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putProjection",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbTableVectorIndexesOutputReference) PutSearchSchema(value interface{}) {
	if err := d.validatePutSearchSchemaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSearchSchema",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbTableVectorIndexesOutputReference) PutVectorAttribute(value *DynamodbTableVectorIndexesVectorAttribute) {
	if err := d.validatePutVectorAttributeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putVectorAttribute",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbTableVectorIndexesOutputReference) ResetDimensions() {
	_jsii_.InvokeVoid(
		d,
		"resetDimensions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTableVectorIndexesOutputReference) ResetDistanceFunction() {
	_jsii_.InvokeVoid(
		d,
		"resetDistanceFunction",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTableVectorIndexesOutputReference) ResetIndexName() {
	_jsii_.InvokeVoid(
		d,
		"resetIndexName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTableVectorIndexesOutputReference) ResetProjection() {
	_jsii_.InvokeVoid(
		d,
		"resetProjection",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTableVectorIndexesOutputReference) ResetSearchSchema() {
	_jsii_.InvokeVoid(
		d,
		"resetSearchSchema",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTableVectorIndexesOutputReference) ResetVectorAttribute() {
	_jsii_.InvokeVoid(
		d,
		"resetVectorAttribute",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTableVectorIndexesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DynamodbTableVectorIndexesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

