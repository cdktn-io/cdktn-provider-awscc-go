// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerprocessingjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/sagemakerprocessingjob/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerProcessingJobProcessingInputsDatasetDefinitionAthenaDatasetDefinitionOutputReference interface {
	cdktn.ComplexObject
	Catalog() *string
	SetCatalog(val *string)
	CatalogInput() *string
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
	Database() *string
	SetDatabase(val *string)
	DatabaseInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	KmsKeyId() *string
	SetKmsKeyId(val *string)
	KmsKeyIdInput() *string
	OutputCompression() *string
	SetOutputCompression(val *string)
	OutputCompressionInput() *string
	OutputFormat() *string
	SetOutputFormat(val *string)
	OutputFormatInput() *string
	OutputS3Uri() *string
	SetOutputS3Uri(val *string)
	OutputS3UriInput() *string
	QueryString() *string
	SetQueryString(val *string)
	QueryStringInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	WorkGroup() *string
	SetWorkGroup(val *string)
	WorkGroupInput() *string
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
	ResetCatalog()
	ResetDatabase()
	ResetKmsKeyId()
	ResetOutputCompression()
	ResetOutputFormat()
	ResetOutputS3Uri()
	ResetQueryString()
	ResetWorkGroup()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerProcessingJobProcessingInputsDatasetDefinitionAthenaDatasetDefinitionOutputReference
type jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionAthenaDatasetDefinitionOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionAthenaDatasetDefinitionOutputReference) Catalog() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionAthenaDatasetDefinitionOutputReference) CatalogInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionAthenaDatasetDefinitionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionAthenaDatasetDefinitionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionAthenaDatasetDefinitionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionAthenaDatasetDefinitionOutputReference) Database() *string {
	var returns *string
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionAthenaDatasetDefinitionOutputReference) DatabaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionAthenaDatasetDefinitionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionAthenaDatasetDefinitionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionAthenaDatasetDefinitionOutputReference) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionAthenaDatasetDefinitionOutputReference) KmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionAthenaDatasetDefinitionOutputReference) OutputCompression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputCompression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionAthenaDatasetDefinitionOutputReference) OutputCompressionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputCompressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionAthenaDatasetDefinitionOutputReference) OutputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionAthenaDatasetDefinitionOutputReference) OutputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionAthenaDatasetDefinitionOutputReference) OutputS3Uri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputS3Uri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionAthenaDatasetDefinitionOutputReference) OutputS3UriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputS3UriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionAthenaDatasetDefinitionOutputReference) QueryString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionAthenaDatasetDefinitionOutputReference) QueryStringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionAthenaDatasetDefinitionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionAthenaDatasetDefinitionOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionAthenaDatasetDefinitionOutputReference) WorkGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionAthenaDatasetDefinitionOutputReference) WorkGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workGroupInput",
		&returns,
	)
	return returns
}


func NewSagemakerProcessingJobProcessingInputsDatasetDefinitionAthenaDatasetDefinitionOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) SagemakerProcessingJobProcessingInputsDatasetDefinitionAthenaDatasetDefinitionOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerProcessingJobProcessingInputsDatasetDefinitionAthenaDatasetDefinitionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionAthenaDatasetDefinitionOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.sagemakerProcessingJob.SagemakerProcessingJobProcessingInputsDatasetDefinitionAthenaDatasetDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSagemakerProcessingJobProcessingInputsDatasetDefinitionAthenaDatasetDefinitionOutputReference_Override(s SagemakerProcessingJobProcessingInputsDatasetDefinitionAthenaDatasetDefinitionOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.sagemakerProcessingJob.SagemakerProcessingJobProcessingInputsDatasetDefinitionAthenaDatasetDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionAthenaDatasetDefinitionOutputReference)SetCatalog(val *string) {
	if err := j.validateSetCatalogParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"catalog",
		val,
	)
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionAthenaDatasetDefinitionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionAthenaDatasetDefinitionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionAthenaDatasetDefinitionOutputReference)SetDatabase(val *string) {
	if err := j.validateSetDatabaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"database",
		val,
	)
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionAthenaDatasetDefinitionOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionAthenaDatasetDefinitionOutputReference)SetKmsKeyId(val *string) {
	if err := j.validateSetKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionAthenaDatasetDefinitionOutputReference)SetOutputCompression(val *string) {
	if err := j.validateSetOutputCompressionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputCompression",
		val,
	)
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionAthenaDatasetDefinitionOutputReference)SetOutputFormat(val *string) {
	if err := j.validateSetOutputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputFormat",
		val,
	)
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionAthenaDatasetDefinitionOutputReference)SetOutputS3Uri(val *string) {
	if err := j.validateSetOutputS3UriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputS3Uri",
		val,
	)
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionAthenaDatasetDefinitionOutputReference)SetQueryString(val *string) {
	if err := j.validateSetQueryStringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryString",
		val,
	)
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionAthenaDatasetDefinitionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionAthenaDatasetDefinitionOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionAthenaDatasetDefinitionOutputReference)SetWorkGroup(val *string) {
	if err := j.validateSetWorkGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workGroup",
		val,
	)
}

func (s *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionAthenaDatasetDefinitionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionAthenaDatasetDefinitionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionAthenaDatasetDefinitionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionAthenaDatasetDefinitionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionAthenaDatasetDefinitionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionAthenaDatasetDefinitionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionAthenaDatasetDefinitionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionAthenaDatasetDefinitionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionAthenaDatasetDefinitionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionAthenaDatasetDefinitionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionAthenaDatasetDefinitionOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionAthenaDatasetDefinitionOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionAthenaDatasetDefinitionOutputReference) ResetCatalog() {
	_jsii_.InvokeVoid(
		s,
		"resetCatalog",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionAthenaDatasetDefinitionOutputReference) ResetDatabase() {
	_jsii_.InvokeVoid(
		s,
		"resetDatabase",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionAthenaDatasetDefinitionOutputReference) ResetKmsKeyId() {
	_jsii_.InvokeVoid(
		s,
		"resetKmsKeyId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionAthenaDatasetDefinitionOutputReference) ResetOutputCompression() {
	_jsii_.InvokeVoid(
		s,
		"resetOutputCompression",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionAthenaDatasetDefinitionOutputReference) ResetOutputFormat() {
	_jsii_.InvokeVoid(
		s,
		"resetOutputFormat",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionAthenaDatasetDefinitionOutputReference) ResetOutputS3Uri() {
	_jsii_.InvokeVoid(
		s,
		"resetOutputS3Uri",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionAthenaDatasetDefinitionOutputReference) ResetQueryString() {
	_jsii_.InvokeVoid(
		s,
		"resetQueryString",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionAthenaDatasetDefinitionOutputReference) ResetWorkGroup() {
	_jsii_.InvokeVoid(
		s,
		"resetWorkGroup",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionAthenaDatasetDefinitionOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionAthenaDatasetDefinitionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

