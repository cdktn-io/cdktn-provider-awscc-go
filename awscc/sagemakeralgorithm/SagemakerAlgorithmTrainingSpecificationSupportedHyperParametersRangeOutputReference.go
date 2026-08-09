// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakeralgorithm

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/sagemakeralgorithm/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerAlgorithmTrainingSpecificationSupportedHyperParametersRangeOutputReference interface {
	cdktn.ComplexObject
	CategoricalParameterRangeSpecification() SagemakerAlgorithmTrainingSpecificationSupportedHyperParametersRangeCategoricalParameterRangeSpecificationOutputReference
	CategoricalParameterRangeSpecificationInput() interface{}
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
	ContinuousParameterRangeSpecification() SagemakerAlgorithmTrainingSpecificationSupportedHyperParametersRangeContinuousParameterRangeSpecificationOutputReference
	ContinuousParameterRangeSpecificationInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	IntegerParameterRangeSpecification() SagemakerAlgorithmTrainingSpecificationSupportedHyperParametersRangeIntegerParameterRangeSpecificationOutputReference
	IntegerParameterRangeSpecificationInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
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
	PutCategoricalParameterRangeSpecification(value *SagemakerAlgorithmTrainingSpecificationSupportedHyperParametersRangeCategoricalParameterRangeSpecification)
	PutContinuousParameterRangeSpecification(value *SagemakerAlgorithmTrainingSpecificationSupportedHyperParametersRangeContinuousParameterRangeSpecification)
	PutIntegerParameterRangeSpecification(value *SagemakerAlgorithmTrainingSpecificationSupportedHyperParametersRangeIntegerParameterRangeSpecification)
	ResetCategoricalParameterRangeSpecification()
	ResetContinuousParameterRangeSpecification()
	ResetIntegerParameterRangeSpecification()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerAlgorithmTrainingSpecificationSupportedHyperParametersRangeOutputReference
type jsiiProxy_SagemakerAlgorithmTrainingSpecificationSupportedHyperParametersRangeOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_SagemakerAlgorithmTrainingSpecificationSupportedHyperParametersRangeOutputReference) CategoricalParameterRangeSpecification() SagemakerAlgorithmTrainingSpecificationSupportedHyperParametersRangeCategoricalParameterRangeSpecificationOutputReference {
	var returns SagemakerAlgorithmTrainingSpecificationSupportedHyperParametersRangeCategoricalParameterRangeSpecificationOutputReference
	_jsii_.Get(
		j,
		"categoricalParameterRangeSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmTrainingSpecificationSupportedHyperParametersRangeOutputReference) CategoricalParameterRangeSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"categoricalParameterRangeSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmTrainingSpecificationSupportedHyperParametersRangeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmTrainingSpecificationSupportedHyperParametersRangeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmTrainingSpecificationSupportedHyperParametersRangeOutputReference) ContinuousParameterRangeSpecification() SagemakerAlgorithmTrainingSpecificationSupportedHyperParametersRangeContinuousParameterRangeSpecificationOutputReference {
	var returns SagemakerAlgorithmTrainingSpecificationSupportedHyperParametersRangeContinuousParameterRangeSpecificationOutputReference
	_jsii_.Get(
		j,
		"continuousParameterRangeSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmTrainingSpecificationSupportedHyperParametersRangeOutputReference) ContinuousParameterRangeSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"continuousParameterRangeSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmTrainingSpecificationSupportedHyperParametersRangeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmTrainingSpecificationSupportedHyperParametersRangeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmTrainingSpecificationSupportedHyperParametersRangeOutputReference) IntegerParameterRangeSpecification() SagemakerAlgorithmTrainingSpecificationSupportedHyperParametersRangeIntegerParameterRangeSpecificationOutputReference {
	var returns SagemakerAlgorithmTrainingSpecificationSupportedHyperParametersRangeIntegerParameterRangeSpecificationOutputReference
	_jsii_.Get(
		j,
		"integerParameterRangeSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmTrainingSpecificationSupportedHyperParametersRangeOutputReference) IntegerParameterRangeSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"integerParameterRangeSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmTrainingSpecificationSupportedHyperParametersRangeOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmTrainingSpecificationSupportedHyperParametersRangeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmTrainingSpecificationSupportedHyperParametersRangeOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSagemakerAlgorithmTrainingSpecificationSupportedHyperParametersRangeOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) SagemakerAlgorithmTrainingSpecificationSupportedHyperParametersRangeOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerAlgorithmTrainingSpecificationSupportedHyperParametersRangeOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerAlgorithmTrainingSpecificationSupportedHyperParametersRangeOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.sagemakerAlgorithm.SagemakerAlgorithmTrainingSpecificationSupportedHyperParametersRangeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSagemakerAlgorithmTrainingSpecificationSupportedHyperParametersRangeOutputReference_Override(s SagemakerAlgorithmTrainingSpecificationSupportedHyperParametersRangeOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.sagemakerAlgorithm.SagemakerAlgorithmTrainingSpecificationSupportedHyperParametersRangeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SagemakerAlgorithmTrainingSpecificationSupportedHyperParametersRangeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerAlgorithmTrainingSpecificationSupportedHyperParametersRangeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerAlgorithmTrainingSpecificationSupportedHyperParametersRangeOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerAlgorithmTrainingSpecificationSupportedHyperParametersRangeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerAlgorithmTrainingSpecificationSupportedHyperParametersRangeOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SagemakerAlgorithmTrainingSpecificationSupportedHyperParametersRangeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerAlgorithmTrainingSpecificationSupportedHyperParametersRangeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SagemakerAlgorithmTrainingSpecificationSupportedHyperParametersRangeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerAlgorithmTrainingSpecificationSupportedHyperParametersRangeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SagemakerAlgorithmTrainingSpecificationSupportedHyperParametersRangeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SagemakerAlgorithmTrainingSpecificationSupportedHyperParametersRangeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SagemakerAlgorithmTrainingSpecificationSupportedHyperParametersRangeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SagemakerAlgorithmTrainingSpecificationSupportedHyperParametersRangeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SagemakerAlgorithmTrainingSpecificationSupportedHyperParametersRangeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SagemakerAlgorithmTrainingSpecificationSupportedHyperParametersRangeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SagemakerAlgorithmTrainingSpecificationSupportedHyperParametersRangeOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerAlgorithmTrainingSpecificationSupportedHyperParametersRangeOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerAlgorithmTrainingSpecificationSupportedHyperParametersRangeOutputReference) PutCategoricalParameterRangeSpecification(value *SagemakerAlgorithmTrainingSpecificationSupportedHyperParametersRangeCategoricalParameterRangeSpecification) {
	if err := s.validatePutCategoricalParameterRangeSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putCategoricalParameterRangeSpecification",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerAlgorithmTrainingSpecificationSupportedHyperParametersRangeOutputReference) PutContinuousParameterRangeSpecification(value *SagemakerAlgorithmTrainingSpecificationSupportedHyperParametersRangeContinuousParameterRangeSpecification) {
	if err := s.validatePutContinuousParameterRangeSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putContinuousParameterRangeSpecification",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerAlgorithmTrainingSpecificationSupportedHyperParametersRangeOutputReference) PutIntegerParameterRangeSpecification(value *SagemakerAlgorithmTrainingSpecificationSupportedHyperParametersRangeIntegerParameterRangeSpecification) {
	if err := s.validatePutIntegerParameterRangeSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putIntegerParameterRangeSpecification",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerAlgorithmTrainingSpecificationSupportedHyperParametersRangeOutputReference) ResetCategoricalParameterRangeSpecification() {
	_jsii_.InvokeVoid(
		s,
		"resetCategoricalParameterRangeSpecification",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerAlgorithmTrainingSpecificationSupportedHyperParametersRangeOutputReference) ResetContinuousParameterRangeSpecification() {
	_jsii_.InvokeVoid(
		s,
		"resetContinuousParameterRangeSpecification",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerAlgorithmTrainingSpecificationSupportedHyperParametersRangeOutputReference) ResetIntegerParameterRangeSpecification() {
	_jsii_.InvokeVoid(
		s,
		"resetIntegerParameterRangeSpecification",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerAlgorithmTrainingSpecificationSupportedHyperParametersRangeOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_SagemakerAlgorithmTrainingSpecificationSupportedHyperParametersRangeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

