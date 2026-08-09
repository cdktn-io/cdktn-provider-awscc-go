// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3tablestable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/s3tablestable/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type S3TablesTableIcebergMetadataOutputReference interface {
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
	IcebergPartitionSpec() S3TablesTableIcebergMetadataIcebergPartitionSpecOutputReference
	IcebergPartitionSpecInput() interface{}
	IcebergSchema() S3TablesTableIcebergMetadataIcebergSchemaOutputReference
	IcebergSchemaInput() interface{}
	IcebergSchemaV2() S3TablesTableIcebergMetadataIcebergSchemaV2OutputReference
	IcebergSchemaV2Input() interface{}
	IcebergSortOrder() S3TablesTableIcebergMetadataIcebergSortOrderOutputReference
	IcebergSortOrderInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	TableProperties() *map[string]*string
	SetTableProperties(val *map[string]*string)
	TablePropertiesInput() *map[string]*string
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
	PutIcebergPartitionSpec(value *S3TablesTableIcebergMetadataIcebergPartitionSpec)
	PutIcebergSchema(value *S3TablesTableIcebergMetadataIcebergSchema)
	PutIcebergSchemaV2(value *S3TablesTableIcebergMetadataIcebergSchemaV2)
	PutIcebergSortOrder(value *S3TablesTableIcebergMetadataIcebergSortOrder)
	ResetIcebergPartitionSpec()
	ResetIcebergSchema()
	ResetIcebergSchemaV2()
	ResetIcebergSortOrder()
	ResetTableProperties()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for S3TablesTableIcebergMetadataOutputReference
type jsiiProxy_S3TablesTableIcebergMetadataOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_S3TablesTableIcebergMetadataOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3TablesTableIcebergMetadataOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3TablesTableIcebergMetadataOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3TablesTableIcebergMetadataOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3TablesTableIcebergMetadataOutputReference) IcebergPartitionSpec() S3TablesTableIcebergMetadataIcebergPartitionSpecOutputReference {
	var returns S3TablesTableIcebergMetadataIcebergPartitionSpecOutputReference
	_jsii_.Get(
		j,
		"icebergPartitionSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3TablesTableIcebergMetadataOutputReference) IcebergPartitionSpecInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"icebergPartitionSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3TablesTableIcebergMetadataOutputReference) IcebergSchema() S3TablesTableIcebergMetadataIcebergSchemaOutputReference {
	var returns S3TablesTableIcebergMetadataIcebergSchemaOutputReference
	_jsii_.Get(
		j,
		"icebergSchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3TablesTableIcebergMetadataOutputReference) IcebergSchemaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"icebergSchemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3TablesTableIcebergMetadataOutputReference) IcebergSchemaV2() S3TablesTableIcebergMetadataIcebergSchemaV2OutputReference {
	var returns S3TablesTableIcebergMetadataIcebergSchemaV2OutputReference
	_jsii_.Get(
		j,
		"icebergSchemaV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3TablesTableIcebergMetadataOutputReference) IcebergSchemaV2Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"icebergSchemaV2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3TablesTableIcebergMetadataOutputReference) IcebergSortOrder() S3TablesTableIcebergMetadataIcebergSortOrderOutputReference {
	var returns S3TablesTableIcebergMetadataIcebergSortOrderOutputReference
	_jsii_.Get(
		j,
		"icebergSortOrder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3TablesTableIcebergMetadataOutputReference) IcebergSortOrderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"icebergSortOrderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3TablesTableIcebergMetadataOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3TablesTableIcebergMetadataOutputReference) TableProperties() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tableProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3TablesTableIcebergMetadataOutputReference) TablePropertiesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tablePropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3TablesTableIcebergMetadataOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3TablesTableIcebergMetadataOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewS3TablesTableIcebergMetadataOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) S3TablesTableIcebergMetadataOutputReference {
	_init_.Initialize()

	if err := validateNewS3TablesTableIcebergMetadataOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3TablesTableIcebergMetadataOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.s3TablesTable.S3TablesTableIcebergMetadataOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewS3TablesTableIcebergMetadataOutputReference_Override(s S3TablesTableIcebergMetadataOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.s3TablesTable.S3TablesTableIcebergMetadataOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_S3TablesTableIcebergMetadataOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_S3TablesTableIcebergMetadataOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_S3TablesTableIcebergMetadataOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_S3TablesTableIcebergMetadataOutputReference)SetTableProperties(val *map[string]*string) {
	if err := j.validateSetTablePropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tableProperties",
		val,
	)
}

func (j *jsiiProxy_S3TablesTableIcebergMetadataOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_S3TablesTableIcebergMetadataOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_S3TablesTableIcebergMetadataOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3TablesTableIcebergMetadataOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_S3TablesTableIcebergMetadataOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_S3TablesTableIcebergMetadataOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_S3TablesTableIcebergMetadataOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_S3TablesTableIcebergMetadataOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_S3TablesTableIcebergMetadataOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_S3TablesTableIcebergMetadataOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_S3TablesTableIcebergMetadataOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_S3TablesTableIcebergMetadataOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_S3TablesTableIcebergMetadataOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3TablesTableIcebergMetadataOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_S3TablesTableIcebergMetadataOutputReference) PutIcebergPartitionSpec(value *S3TablesTableIcebergMetadataIcebergPartitionSpec) {
	if err := s.validatePutIcebergPartitionSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putIcebergPartitionSpec",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3TablesTableIcebergMetadataOutputReference) PutIcebergSchema(value *S3TablesTableIcebergMetadataIcebergSchema) {
	if err := s.validatePutIcebergSchemaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putIcebergSchema",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3TablesTableIcebergMetadataOutputReference) PutIcebergSchemaV2(value *S3TablesTableIcebergMetadataIcebergSchemaV2) {
	if err := s.validatePutIcebergSchemaV2Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putIcebergSchemaV2",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3TablesTableIcebergMetadataOutputReference) PutIcebergSortOrder(value *S3TablesTableIcebergMetadataIcebergSortOrder) {
	if err := s.validatePutIcebergSortOrderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putIcebergSortOrder",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3TablesTableIcebergMetadataOutputReference) ResetIcebergPartitionSpec() {
	_jsii_.InvokeVoid(
		s,
		"resetIcebergPartitionSpec",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3TablesTableIcebergMetadataOutputReference) ResetIcebergSchema() {
	_jsii_.InvokeVoid(
		s,
		"resetIcebergSchema",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3TablesTableIcebergMetadataOutputReference) ResetIcebergSchemaV2() {
	_jsii_.InvokeVoid(
		s,
		"resetIcebergSchemaV2",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3TablesTableIcebergMetadataOutputReference) ResetIcebergSortOrder() {
	_jsii_.InvokeVoid(
		s,
		"resetIcebergSortOrder",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3TablesTableIcebergMetadataOutputReference) ResetTableProperties() {
	_jsii_.InvokeVoid(
		s,
		"resetTableProperties",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3TablesTableIcebergMetadataOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_S3TablesTableIcebergMetadataOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

