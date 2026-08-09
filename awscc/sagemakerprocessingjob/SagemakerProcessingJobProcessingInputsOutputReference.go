// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerprocessingjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/sagemakerprocessingjob/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerProcessingJobProcessingInputsOutputReference interface {
	cdktn.ComplexObject
	AppManaged() interface{}
	SetAppManaged(val interface{})
	AppManagedInput() interface{}
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
	DatasetDefinition() SagemakerProcessingJobProcessingInputsDatasetDefinitionOutputReference
	DatasetDefinitionInput() interface{}
	// Experimental.
	Fqn() *string
	InputName() *string
	SetInputName(val *string)
	InputNameInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	S3Input() SagemakerProcessingJobProcessingInputsS3InputOutputReference
	S3InputInput() interface{}
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
	PutDatasetDefinition(value *SagemakerProcessingJobProcessingInputsDatasetDefinition)
	PutS3Input(value *SagemakerProcessingJobProcessingInputsS3Input)
	ResetAppManaged()
	ResetDatasetDefinition()
	ResetInputName()
	ResetS3Input()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerProcessingJobProcessingInputsOutputReference
type jsiiProxy_SagemakerProcessingJobProcessingInputsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsOutputReference) AppManaged() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"appManaged",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsOutputReference) AppManagedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"appManagedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsOutputReference) DatasetDefinition() SagemakerProcessingJobProcessingInputsDatasetDefinitionOutputReference {
	var returns SagemakerProcessingJobProcessingInputsDatasetDefinitionOutputReference
	_jsii_.Get(
		j,
		"datasetDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsOutputReference) DatasetDefinitionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"datasetDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsOutputReference) InputName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsOutputReference) InputNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsOutputReference) S3Input() SagemakerProcessingJobProcessingInputsS3InputOutputReference {
	var returns SagemakerProcessingJobProcessingInputsS3InputOutputReference
	_jsii_.Get(
		j,
		"s3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsOutputReference) S3InputInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3InputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSagemakerProcessingJobProcessingInputsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) SagemakerProcessingJobProcessingInputsOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerProcessingJobProcessingInputsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerProcessingJobProcessingInputsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.sagemakerProcessingJob.SagemakerProcessingJobProcessingInputsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewSagemakerProcessingJobProcessingInputsOutputReference_Override(s SagemakerProcessingJobProcessingInputsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.sagemakerProcessingJob.SagemakerProcessingJobProcessingInputsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsOutputReference)SetAppManaged(val interface{}) {
	if err := j.validateSetAppManagedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appManaged",
		val,
	)
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsOutputReference)SetInputName(val *string) {
	if err := j.validateSetInputNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputName",
		val,
	)
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SagemakerProcessingJobProcessingInputsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerProcessingJobProcessingInputsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SagemakerProcessingJobProcessingInputsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerProcessingJobProcessingInputsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SagemakerProcessingJobProcessingInputsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SagemakerProcessingJobProcessingInputsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SagemakerProcessingJobProcessingInputsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SagemakerProcessingJobProcessingInputsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SagemakerProcessingJobProcessingInputsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SagemakerProcessingJobProcessingInputsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SagemakerProcessingJobProcessingInputsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerProcessingJobProcessingInputsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerProcessingJobProcessingInputsOutputReference) PutDatasetDefinition(value *SagemakerProcessingJobProcessingInputsDatasetDefinition) {
	if err := s.validatePutDatasetDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putDatasetDefinition",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerProcessingJobProcessingInputsOutputReference) PutS3Input(value *SagemakerProcessingJobProcessingInputsS3Input) {
	if err := s.validatePutS3InputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putS3Input",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerProcessingJobProcessingInputsOutputReference) ResetAppManaged() {
	_jsii_.InvokeVoid(
		s,
		"resetAppManaged",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerProcessingJobProcessingInputsOutputReference) ResetDatasetDefinition() {
	_jsii_.InvokeVoid(
		s,
		"resetDatasetDefinition",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerProcessingJobProcessingInputsOutputReference) ResetInputName() {
	_jsii_.InvokeVoid(
		s,
		"resetInputName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerProcessingJobProcessingInputsOutputReference) ResetS3Input() {
	_jsii_.InvokeVoid(
		s,
		"resetS3Input",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerProcessingJobProcessingInputsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_SagemakerProcessingJobProcessingInputsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

