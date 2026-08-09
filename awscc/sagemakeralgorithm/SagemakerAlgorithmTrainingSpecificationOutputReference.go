// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakeralgorithm

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/sagemakeralgorithm/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerAlgorithmTrainingSpecificationOutputReference interface {
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MetricDefinitions() SagemakerAlgorithmTrainingSpecificationMetricDefinitionsList
	MetricDefinitionsInput() interface{}
	SupportedHyperParameters() SagemakerAlgorithmTrainingSpecificationSupportedHyperParametersList
	SupportedHyperParametersInput() interface{}
	SupportedTrainingInstanceTypes() *[]*string
	SetSupportedTrainingInstanceTypes(val *[]*string)
	SupportedTrainingInstanceTypesInput() *[]*string
	SupportedTuningJobObjectiveMetrics() SagemakerAlgorithmTrainingSpecificationSupportedTuningJobObjectiveMetricsList
	SupportedTuningJobObjectiveMetricsInput() interface{}
	SupportsDistributedTraining() interface{}
	SetSupportsDistributedTraining(val interface{})
	SupportsDistributedTrainingInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TrainingChannels() SagemakerAlgorithmTrainingSpecificationTrainingChannelsList
	TrainingChannelsInput() interface{}
	TrainingImage() *string
	SetTrainingImage(val *string)
	TrainingImageDigest() *string
	SetTrainingImageDigest(val *string)
	TrainingImageDigestInput() *string
	TrainingImageInput() *string
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
	PutMetricDefinitions(value interface{})
	PutSupportedHyperParameters(value interface{})
	PutSupportedTuningJobObjectiveMetrics(value interface{})
	PutTrainingChannels(value interface{})
	ResetMetricDefinitions()
	ResetSupportedHyperParameters()
	ResetSupportedTuningJobObjectiveMetrics()
	ResetSupportsDistributedTraining()
	ResetTrainingImageDigest()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerAlgorithmTrainingSpecificationOutputReference
type jsiiProxy_SagemakerAlgorithmTrainingSpecificationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_SagemakerAlgorithmTrainingSpecificationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmTrainingSpecificationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmTrainingSpecificationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmTrainingSpecificationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmTrainingSpecificationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmTrainingSpecificationOutputReference) MetricDefinitions() SagemakerAlgorithmTrainingSpecificationMetricDefinitionsList {
	var returns SagemakerAlgorithmTrainingSpecificationMetricDefinitionsList
	_jsii_.Get(
		j,
		"metricDefinitions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmTrainingSpecificationOutputReference) MetricDefinitionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metricDefinitionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmTrainingSpecificationOutputReference) SupportedHyperParameters() SagemakerAlgorithmTrainingSpecificationSupportedHyperParametersList {
	var returns SagemakerAlgorithmTrainingSpecificationSupportedHyperParametersList
	_jsii_.Get(
		j,
		"supportedHyperParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmTrainingSpecificationOutputReference) SupportedHyperParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"supportedHyperParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmTrainingSpecificationOutputReference) SupportedTrainingInstanceTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedTrainingInstanceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmTrainingSpecificationOutputReference) SupportedTrainingInstanceTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedTrainingInstanceTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmTrainingSpecificationOutputReference) SupportedTuningJobObjectiveMetrics() SagemakerAlgorithmTrainingSpecificationSupportedTuningJobObjectiveMetricsList {
	var returns SagemakerAlgorithmTrainingSpecificationSupportedTuningJobObjectiveMetricsList
	_jsii_.Get(
		j,
		"supportedTuningJobObjectiveMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmTrainingSpecificationOutputReference) SupportedTuningJobObjectiveMetricsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"supportedTuningJobObjectiveMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmTrainingSpecificationOutputReference) SupportsDistributedTraining() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"supportsDistributedTraining",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmTrainingSpecificationOutputReference) SupportsDistributedTrainingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"supportsDistributedTrainingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmTrainingSpecificationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmTrainingSpecificationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmTrainingSpecificationOutputReference) TrainingChannels() SagemakerAlgorithmTrainingSpecificationTrainingChannelsList {
	var returns SagemakerAlgorithmTrainingSpecificationTrainingChannelsList
	_jsii_.Get(
		j,
		"trainingChannels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmTrainingSpecificationOutputReference) TrainingChannelsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trainingChannelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmTrainingSpecificationOutputReference) TrainingImage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trainingImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmTrainingSpecificationOutputReference) TrainingImageDigest() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trainingImageDigest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmTrainingSpecificationOutputReference) TrainingImageDigestInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trainingImageDigestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmTrainingSpecificationOutputReference) TrainingImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trainingImageInput",
		&returns,
	)
	return returns
}


func NewSagemakerAlgorithmTrainingSpecificationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) SagemakerAlgorithmTrainingSpecificationOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerAlgorithmTrainingSpecificationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerAlgorithmTrainingSpecificationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.sagemakerAlgorithm.SagemakerAlgorithmTrainingSpecificationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSagemakerAlgorithmTrainingSpecificationOutputReference_Override(s SagemakerAlgorithmTrainingSpecificationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.sagemakerAlgorithm.SagemakerAlgorithmTrainingSpecificationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SagemakerAlgorithmTrainingSpecificationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerAlgorithmTrainingSpecificationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerAlgorithmTrainingSpecificationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerAlgorithmTrainingSpecificationOutputReference)SetSupportedTrainingInstanceTypes(val *[]*string) {
	if err := j.validateSetSupportedTrainingInstanceTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supportedTrainingInstanceTypes",
		val,
	)
}

func (j *jsiiProxy_SagemakerAlgorithmTrainingSpecificationOutputReference)SetSupportsDistributedTraining(val interface{}) {
	if err := j.validateSetSupportsDistributedTrainingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supportsDistributedTraining",
		val,
	)
}

func (j *jsiiProxy_SagemakerAlgorithmTrainingSpecificationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerAlgorithmTrainingSpecificationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SagemakerAlgorithmTrainingSpecificationOutputReference)SetTrainingImage(val *string) {
	if err := j.validateSetTrainingImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trainingImage",
		val,
	)
}

func (j *jsiiProxy_SagemakerAlgorithmTrainingSpecificationOutputReference)SetTrainingImageDigest(val *string) {
	if err := j.validateSetTrainingImageDigestParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trainingImageDigest",
		val,
	)
}

func (s *jsiiProxy_SagemakerAlgorithmTrainingSpecificationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerAlgorithmTrainingSpecificationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SagemakerAlgorithmTrainingSpecificationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerAlgorithmTrainingSpecificationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SagemakerAlgorithmTrainingSpecificationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SagemakerAlgorithmTrainingSpecificationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SagemakerAlgorithmTrainingSpecificationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SagemakerAlgorithmTrainingSpecificationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SagemakerAlgorithmTrainingSpecificationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SagemakerAlgorithmTrainingSpecificationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SagemakerAlgorithmTrainingSpecificationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerAlgorithmTrainingSpecificationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerAlgorithmTrainingSpecificationOutputReference) PutMetricDefinitions(value interface{}) {
	if err := s.validatePutMetricDefinitionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putMetricDefinitions",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerAlgorithmTrainingSpecificationOutputReference) PutSupportedHyperParameters(value interface{}) {
	if err := s.validatePutSupportedHyperParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSupportedHyperParameters",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerAlgorithmTrainingSpecificationOutputReference) PutSupportedTuningJobObjectiveMetrics(value interface{}) {
	if err := s.validatePutSupportedTuningJobObjectiveMetricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSupportedTuningJobObjectiveMetrics",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerAlgorithmTrainingSpecificationOutputReference) PutTrainingChannels(value interface{}) {
	if err := s.validatePutTrainingChannelsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTrainingChannels",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerAlgorithmTrainingSpecificationOutputReference) ResetMetricDefinitions() {
	_jsii_.InvokeVoid(
		s,
		"resetMetricDefinitions",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerAlgorithmTrainingSpecificationOutputReference) ResetSupportedHyperParameters() {
	_jsii_.InvokeVoid(
		s,
		"resetSupportedHyperParameters",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerAlgorithmTrainingSpecificationOutputReference) ResetSupportedTuningJobObjectiveMetrics() {
	_jsii_.InvokeVoid(
		s,
		"resetSupportedTuningJobObjectiveMetrics",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerAlgorithmTrainingSpecificationOutputReference) ResetSupportsDistributedTraining() {
	_jsii_.InvokeVoid(
		s,
		"resetSupportsDistributedTraining",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerAlgorithmTrainingSpecificationOutputReference) ResetTrainingImageDigest() {
	_jsii_.InvokeVoid(
		s,
		"resetTrainingImageDigest",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerAlgorithmTrainingSpecificationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_SagemakerAlgorithmTrainingSpecificationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

