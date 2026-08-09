// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3tablestable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/s3tablestable/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type S3TablesTableIcebergMetadataIcebergSchemaV2OutputReference interface {
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
	IdentifierFieldIds() *[]*float64
	SetIdentifierFieldIds(val *[]*float64)
	IdentifierFieldIdsInput() *[]*float64
	InternalValue() interface{}
	SetInternalValue(val interface{})
	SchemaId() *float64
	SetSchemaId(val *float64)
	SchemaIdInput() *float64
	SchemaV2FieldList() S3TablesTableIcebergMetadataIcebergSchemaV2SchemaV2FieldListStructList
	SchemaV2FieldListInput() interface{}
	SchemaV2FieldType() *string
	SetSchemaV2FieldType(val *string)
	SchemaV2FieldTypeInput() *string
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
	PutSchemaV2FieldList(value interface{})
	ResetIdentifierFieldIds()
	ResetSchemaId()
	ResetSchemaV2FieldList()
	ResetSchemaV2FieldType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for S3TablesTableIcebergMetadataIcebergSchemaV2OutputReference
type jsiiProxy_S3TablesTableIcebergMetadataIcebergSchemaV2OutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_S3TablesTableIcebergMetadataIcebergSchemaV2OutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3TablesTableIcebergMetadataIcebergSchemaV2OutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3TablesTableIcebergMetadataIcebergSchemaV2OutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3TablesTableIcebergMetadataIcebergSchemaV2OutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3TablesTableIcebergMetadataIcebergSchemaV2OutputReference) IdentifierFieldIds() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"identifierFieldIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3TablesTableIcebergMetadataIcebergSchemaV2OutputReference) IdentifierFieldIdsInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"identifierFieldIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3TablesTableIcebergMetadataIcebergSchemaV2OutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3TablesTableIcebergMetadataIcebergSchemaV2OutputReference) SchemaId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"schemaId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3TablesTableIcebergMetadataIcebergSchemaV2OutputReference) SchemaIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"schemaIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3TablesTableIcebergMetadataIcebergSchemaV2OutputReference) SchemaV2FieldList() S3TablesTableIcebergMetadataIcebergSchemaV2SchemaV2FieldListStructList {
	var returns S3TablesTableIcebergMetadataIcebergSchemaV2SchemaV2FieldListStructList
	_jsii_.Get(
		j,
		"schemaV2FieldList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3TablesTableIcebergMetadataIcebergSchemaV2OutputReference) SchemaV2FieldListInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"schemaV2FieldListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3TablesTableIcebergMetadataIcebergSchemaV2OutputReference) SchemaV2FieldType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaV2FieldType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3TablesTableIcebergMetadataIcebergSchemaV2OutputReference) SchemaV2FieldTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaV2FieldTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3TablesTableIcebergMetadataIcebergSchemaV2OutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3TablesTableIcebergMetadataIcebergSchemaV2OutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewS3TablesTableIcebergMetadataIcebergSchemaV2OutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) S3TablesTableIcebergMetadataIcebergSchemaV2OutputReference {
	_init_.Initialize()

	if err := validateNewS3TablesTableIcebergMetadataIcebergSchemaV2OutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3TablesTableIcebergMetadataIcebergSchemaV2OutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.s3TablesTable.S3TablesTableIcebergMetadataIcebergSchemaV2OutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewS3TablesTableIcebergMetadataIcebergSchemaV2OutputReference_Override(s S3TablesTableIcebergMetadataIcebergSchemaV2OutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.s3TablesTable.S3TablesTableIcebergMetadataIcebergSchemaV2OutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_S3TablesTableIcebergMetadataIcebergSchemaV2OutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_S3TablesTableIcebergMetadataIcebergSchemaV2OutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_S3TablesTableIcebergMetadataIcebergSchemaV2OutputReference)SetIdentifierFieldIds(val *[]*float64) {
	if err := j.validateSetIdentifierFieldIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identifierFieldIds",
		val,
	)
}

func (j *jsiiProxy_S3TablesTableIcebergMetadataIcebergSchemaV2OutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_S3TablesTableIcebergMetadataIcebergSchemaV2OutputReference)SetSchemaId(val *float64) {
	if err := j.validateSetSchemaIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schemaId",
		val,
	)
}

func (j *jsiiProxy_S3TablesTableIcebergMetadataIcebergSchemaV2OutputReference)SetSchemaV2FieldType(val *string) {
	if err := j.validateSetSchemaV2FieldTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schemaV2FieldType",
		val,
	)
}

func (j *jsiiProxy_S3TablesTableIcebergMetadataIcebergSchemaV2OutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_S3TablesTableIcebergMetadataIcebergSchemaV2OutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_S3TablesTableIcebergMetadataIcebergSchemaV2OutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3TablesTableIcebergMetadataIcebergSchemaV2OutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3TablesTableIcebergMetadataIcebergSchemaV2OutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3TablesTableIcebergMetadataIcebergSchemaV2OutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3TablesTableIcebergMetadataIcebergSchemaV2OutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3TablesTableIcebergMetadataIcebergSchemaV2OutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3TablesTableIcebergMetadataIcebergSchemaV2OutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3TablesTableIcebergMetadataIcebergSchemaV2OutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3TablesTableIcebergMetadataIcebergSchemaV2OutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3TablesTableIcebergMetadataIcebergSchemaV2OutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3TablesTableIcebergMetadataIcebergSchemaV2OutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3TablesTableIcebergMetadataIcebergSchemaV2OutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3TablesTableIcebergMetadataIcebergSchemaV2OutputReference) PutSchemaV2FieldList(value interface{}) {
	if err := s.validatePutSchemaV2FieldListParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSchemaV2FieldList",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3TablesTableIcebergMetadataIcebergSchemaV2OutputReference) ResetIdentifierFieldIds() {
	_jsii_.InvokeVoid(
		s,
		"resetIdentifierFieldIds",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3TablesTableIcebergMetadataIcebergSchemaV2OutputReference) ResetSchemaId() {
	_jsii_.InvokeVoid(
		s,
		"resetSchemaId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3TablesTableIcebergMetadataIcebergSchemaV2OutputReference) ResetSchemaV2FieldList() {
	_jsii_.InvokeVoid(
		s,
		"resetSchemaV2FieldList",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3TablesTableIcebergMetadataIcebergSchemaV2OutputReference) ResetSchemaV2FieldType() {
	_jsii_.InvokeVoid(
		s,
		"resetSchemaV2FieldType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3TablesTableIcebergMetadataIcebergSchemaV2OutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := s.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3TablesTableIcebergMetadataIcebergSchemaV2OutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

