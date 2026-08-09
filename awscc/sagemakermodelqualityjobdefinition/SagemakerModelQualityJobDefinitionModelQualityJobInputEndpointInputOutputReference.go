// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakermodelqualityjobdefinition

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/sagemakermodelqualityjobdefinition/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerModelQualityJobDefinitionModelQualityJobInputEndpointInputOutputReference interface {
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
	EndpointName() *string
	SetEndpointName(val *string)
	EndpointNameInput() *string
	EndTimeOffset() *string
	SetEndTimeOffset(val *string)
	EndTimeOffsetInput() *string
	// Experimental.
	Fqn() *string
	InferenceAttribute() *string
	SetInferenceAttribute(val *string)
	InferenceAttributeInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LocalPath() *string
	SetLocalPath(val *string)
	LocalPathInput() *string
	ProbabilityAttribute() *string
	SetProbabilityAttribute(val *string)
	ProbabilityAttributeInput() *string
	ProbabilityThresholdAttribute() *float64
	SetProbabilityThresholdAttribute(val *float64)
	ProbabilityThresholdAttributeInput() *float64
	S3DataDistributionType() *string
	SetS3DataDistributionType(val *string)
	S3DataDistributionTypeInput() *string
	S3InputMode() *string
	SetS3InputMode(val *string)
	S3InputModeInput() *string
	StartTimeOffset() *string
	SetStartTimeOffset(val *string)
	StartTimeOffsetInput() *string
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
	ResetEndpointName()
	ResetEndTimeOffset()
	ResetInferenceAttribute()
	ResetLocalPath()
	ResetProbabilityAttribute()
	ResetProbabilityThresholdAttribute()
	ResetS3DataDistributionType()
	ResetS3InputMode()
	ResetStartTimeOffset()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerModelQualityJobDefinitionModelQualityJobInputEndpointInputOutputReference
type jsiiProxy_SagemakerModelQualityJobDefinitionModelQualityJobInputEndpointInputOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinitionModelQualityJobInputEndpointInputOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinitionModelQualityJobInputEndpointInputOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinitionModelQualityJobInputEndpointInputOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinitionModelQualityJobInputEndpointInputOutputReference) EndpointName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinitionModelQualityJobInputEndpointInputOutputReference) EndpointNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinitionModelQualityJobInputEndpointInputOutputReference) EndTimeOffset() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endTimeOffset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinitionModelQualityJobInputEndpointInputOutputReference) EndTimeOffsetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endTimeOffsetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinitionModelQualityJobInputEndpointInputOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinitionModelQualityJobInputEndpointInputOutputReference) InferenceAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inferenceAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinitionModelQualityJobInputEndpointInputOutputReference) InferenceAttributeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inferenceAttributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinitionModelQualityJobInputEndpointInputOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinitionModelQualityJobInputEndpointInputOutputReference) LocalPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinitionModelQualityJobInputEndpointInputOutputReference) LocalPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinitionModelQualityJobInputEndpointInputOutputReference) ProbabilityAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"probabilityAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinitionModelQualityJobInputEndpointInputOutputReference) ProbabilityAttributeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"probabilityAttributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinitionModelQualityJobInputEndpointInputOutputReference) ProbabilityThresholdAttribute() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"probabilityThresholdAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinitionModelQualityJobInputEndpointInputOutputReference) ProbabilityThresholdAttributeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"probabilityThresholdAttributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinitionModelQualityJobInputEndpointInputOutputReference) S3DataDistributionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3DataDistributionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinitionModelQualityJobInputEndpointInputOutputReference) S3DataDistributionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3DataDistributionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinitionModelQualityJobInputEndpointInputOutputReference) S3InputMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3InputMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinitionModelQualityJobInputEndpointInputOutputReference) S3InputModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3InputModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinitionModelQualityJobInputEndpointInputOutputReference) StartTimeOffset() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTimeOffset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinitionModelQualityJobInputEndpointInputOutputReference) StartTimeOffsetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTimeOffsetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinitionModelQualityJobInputEndpointInputOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinitionModelQualityJobInputEndpointInputOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSagemakerModelQualityJobDefinitionModelQualityJobInputEndpointInputOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) SagemakerModelQualityJobDefinitionModelQualityJobInputEndpointInputOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerModelQualityJobDefinitionModelQualityJobInputEndpointInputOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerModelQualityJobDefinitionModelQualityJobInputEndpointInputOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.sagemakerModelQualityJobDefinition.SagemakerModelQualityJobDefinitionModelQualityJobInputEndpointInputOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSagemakerModelQualityJobDefinitionModelQualityJobInputEndpointInputOutputReference_Override(s SagemakerModelQualityJobDefinitionModelQualityJobInputEndpointInputOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.sagemakerModelQualityJobDefinition.SagemakerModelQualityJobDefinitionModelQualityJobInputEndpointInputOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinitionModelQualityJobInputEndpointInputOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinitionModelQualityJobInputEndpointInputOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinitionModelQualityJobInputEndpointInputOutputReference)SetEndpointName(val *string) {
	if err := j.validateSetEndpointNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointName",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinitionModelQualityJobInputEndpointInputOutputReference)SetEndTimeOffset(val *string) {
	if err := j.validateSetEndTimeOffsetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endTimeOffset",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinitionModelQualityJobInputEndpointInputOutputReference)SetInferenceAttribute(val *string) {
	if err := j.validateSetInferenceAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inferenceAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinitionModelQualityJobInputEndpointInputOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinitionModelQualityJobInputEndpointInputOutputReference)SetLocalPath(val *string) {
	if err := j.validateSetLocalPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localPath",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinitionModelQualityJobInputEndpointInputOutputReference)SetProbabilityAttribute(val *string) {
	if err := j.validateSetProbabilityAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"probabilityAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinitionModelQualityJobInputEndpointInputOutputReference)SetProbabilityThresholdAttribute(val *float64) {
	if err := j.validateSetProbabilityThresholdAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"probabilityThresholdAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinitionModelQualityJobInputEndpointInputOutputReference)SetS3DataDistributionType(val *string) {
	if err := j.validateSetS3DataDistributionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3DataDistributionType",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinitionModelQualityJobInputEndpointInputOutputReference)SetS3InputMode(val *string) {
	if err := j.validateSetS3InputModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3InputMode",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinitionModelQualityJobInputEndpointInputOutputReference)SetStartTimeOffset(val *string) {
	if err := j.validateSetStartTimeOffsetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startTimeOffset",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinitionModelQualityJobInputEndpointInputOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinitionModelQualityJobInputEndpointInputOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SagemakerModelQualityJobDefinitionModelQualityJobInputEndpointInputOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelQualityJobDefinitionModelQualityJobInputEndpointInputOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SagemakerModelQualityJobDefinitionModelQualityJobInputEndpointInputOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerModelQualityJobDefinitionModelQualityJobInputEndpointInputOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SagemakerModelQualityJobDefinitionModelQualityJobInputEndpointInputOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SagemakerModelQualityJobDefinitionModelQualityJobInputEndpointInputOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SagemakerModelQualityJobDefinitionModelQualityJobInputEndpointInputOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SagemakerModelQualityJobDefinitionModelQualityJobInputEndpointInputOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SagemakerModelQualityJobDefinitionModelQualityJobInputEndpointInputOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SagemakerModelQualityJobDefinitionModelQualityJobInputEndpointInputOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SagemakerModelQualityJobDefinitionModelQualityJobInputEndpointInputOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelQualityJobDefinitionModelQualityJobInputEndpointInputOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerModelQualityJobDefinitionModelQualityJobInputEndpointInputOutputReference) ResetEndpointName() {
	_jsii_.InvokeVoid(
		s,
		"resetEndpointName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelQualityJobDefinitionModelQualityJobInputEndpointInputOutputReference) ResetEndTimeOffset() {
	_jsii_.InvokeVoid(
		s,
		"resetEndTimeOffset",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelQualityJobDefinitionModelQualityJobInputEndpointInputOutputReference) ResetInferenceAttribute() {
	_jsii_.InvokeVoid(
		s,
		"resetInferenceAttribute",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelQualityJobDefinitionModelQualityJobInputEndpointInputOutputReference) ResetLocalPath() {
	_jsii_.InvokeVoid(
		s,
		"resetLocalPath",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelQualityJobDefinitionModelQualityJobInputEndpointInputOutputReference) ResetProbabilityAttribute() {
	_jsii_.InvokeVoid(
		s,
		"resetProbabilityAttribute",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelQualityJobDefinitionModelQualityJobInputEndpointInputOutputReference) ResetProbabilityThresholdAttribute() {
	_jsii_.InvokeVoid(
		s,
		"resetProbabilityThresholdAttribute",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelQualityJobDefinitionModelQualityJobInputEndpointInputOutputReference) ResetS3DataDistributionType() {
	_jsii_.InvokeVoid(
		s,
		"resetS3DataDistributionType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelQualityJobDefinitionModelQualityJobInputEndpointInputOutputReference) ResetS3InputMode() {
	_jsii_.InvokeVoid(
		s,
		"resetS3InputMode",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelQualityJobDefinitionModelQualityJobInputEndpointInputOutputReference) ResetStartTimeOffset() {
	_jsii_.InvokeVoid(
		s,
		"resetStartTimeOffset",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelQualityJobDefinitionModelQualityJobInputEndpointInputOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_SagemakerModelQualityJobDefinitionModelQualityJobInputEndpointInputOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

