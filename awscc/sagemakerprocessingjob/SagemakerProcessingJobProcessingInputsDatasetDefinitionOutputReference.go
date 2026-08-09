// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerprocessingjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/sagemakerprocessingjob/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerProcessingJobProcessingInputsDatasetDefinitionOutputReference interface {
	cdktn.ComplexObject
	AthenaDatasetDefinition() SagemakerProcessingJobProcessingInputsDatasetDefinitionAthenaDatasetDefinitionOutputReference
	AthenaDatasetDefinitionInput() interface{}
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
	DataDistributionType() *string
	SetDataDistributionType(val *string)
	DataDistributionTypeInput() *string
	// Experimental.
	Fqn() *string
	InputMode() *string
	SetInputMode(val *string)
	InputModeInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LocalPath() *string
	SetLocalPath(val *string)
	LocalPathInput() *string
	RedshiftDatasetDefinition() SagemakerProcessingJobProcessingInputsDatasetDefinitionRedshiftDatasetDefinitionOutputReference
	RedshiftDatasetDefinitionInput() interface{}
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
	PutAthenaDatasetDefinition(value *SagemakerProcessingJobProcessingInputsDatasetDefinitionAthenaDatasetDefinition)
	PutRedshiftDatasetDefinition(value *SagemakerProcessingJobProcessingInputsDatasetDefinitionRedshiftDatasetDefinition)
	ResetAthenaDatasetDefinition()
	ResetDataDistributionType()
	ResetInputMode()
	ResetLocalPath()
	ResetRedshiftDatasetDefinition()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerProcessingJobProcessingInputsDatasetDefinitionOutputReference
type jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionOutputReference) AthenaDatasetDefinition() SagemakerProcessingJobProcessingInputsDatasetDefinitionAthenaDatasetDefinitionOutputReference {
	var returns SagemakerProcessingJobProcessingInputsDatasetDefinitionAthenaDatasetDefinitionOutputReference
	_jsii_.Get(
		j,
		"athenaDatasetDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionOutputReference) AthenaDatasetDefinitionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"athenaDatasetDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionOutputReference) DataDistributionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataDistributionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionOutputReference) DataDistributionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataDistributionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionOutputReference) InputMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionOutputReference) InputModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionOutputReference) LocalPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionOutputReference) LocalPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionOutputReference) RedshiftDatasetDefinition() SagemakerProcessingJobProcessingInputsDatasetDefinitionRedshiftDatasetDefinitionOutputReference {
	var returns SagemakerProcessingJobProcessingInputsDatasetDefinitionRedshiftDatasetDefinitionOutputReference
	_jsii_.Get(
		j,
		"redshiftDatasetDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionOutputReference) RedshiftDatasetDefinitionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"redshiftDatasetDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSagemakerProcessingJobProcessingInputsDatasetDefinitionOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) SagemakerProcessingJobProcessingInputsDatasetDefinitionOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerProcessingJobProcessingInputsDatasetDefinitionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.sagemakerProcessingJob.SagemakerProcessingJobProcessingInputsDatasetDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSagemakerProcessingJobProcessingInputsDatasetDefinitionOutputReference_Override(s SagemakerProcessingJobProcessingInputsDatasetDefinitionOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.sagemakerProcessingJob.SagemakerProcessingJobProcessingInputsDatasetDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionOutputReference)SetDataDistributionType(val *string) {
	if err := j.validateSetDataDistributionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataDistributionType",
		val,
	)
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionOutputReference)SetInputMode(val *string) {
	if err := j.validateSetInputModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputMode",
		val,
	)
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionOutputReference)SetLocalPath(val *string) {
	if err := j.validateSetLocalPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localPath",
		val,
	)
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionOutputReference) PutAthenaDatasetDefinition(value *SagemakerProcessingJobProcessingInputsDatasetDefinitionAthenaDatasetDefinition) {
	if err := s.validatePutAthenaDatasetDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAthenaDatasetDefinition",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionOutputReference) PutRedshiftDatasetDefinition(value *SagemakerProcessingJobProcessingInputsDatasetDefinitionRedshiftDatasetDefinition) {
	if err := s.validatePutRedshiftDatasetDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putRedshiftDatasetDefinition",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionOutputReference) ResetAthenaDatasetDefinition() {
	_jsii_.InvokeVoid(
		s,
		"resetAthenaDatasetDefinition",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionOutputReference) ResetDataDistributionType() {
	_jsii_.InvokeVoid(
		s,
		"resetDataDistributionType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionOutputReference) ResetInputMode() {
	_jsii_.InvokeVoid(
		s,
		"resetInputMode",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionOutputReference) ResetLocalPath() {
	_jsii_.InvokeVoid(
		s,
		"resetLocalPath",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionOutputReference) ResetRedshiftDatasetDefinition() {
	_jsii_.InvokeVoid(
		s,
		"resetRedshiftDatasetDefinition",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_SagemakerProcessingJobProcessingInputsDatasetDefinitionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

