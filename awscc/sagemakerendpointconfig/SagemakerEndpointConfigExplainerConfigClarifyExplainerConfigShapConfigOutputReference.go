// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerendpointconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/sagemakerendpointconfig/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigShapConfigOutputReference interface {
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
	NumberOfSamples() *float64
	SetNumberOfSamples(val *float64)
	NumberOfSamplesInput() *float64
	Seed() *float64
	SetSeed(val *float64)
	SeedInput() *float64
	ShapBaselineConfig() SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigShapConfigShapBaselineConfigOutputReference
	ShapBaselineConfigInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TextConfig() SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigShapConfigTextConfigOutputReference
	TextConfigInput() interface{}
	UseLogit() interface{}
	SetUseLogit(val interface{})
	UseLogitInput() interface{}
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
	PutShapBaselineConfig(value *SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigShapConfigShapBaselineConfig)
	PutTextConfig(value *SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigShapConfigTextConfig)
	ResetNumberOfSamples()
	ResetSeed()
	ResetShapBaselineConfig()
	ResetTextConfig()
	ResetUseLogit()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigShapConfigOutputReference
type jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigShapConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigShapConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigShapConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigShapConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigShapConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigShapConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigShapConfigOutputReference) NumberOfSamples() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfSamples",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigShapConfigOutputReference) NumberOfSamplesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfSamplesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigShapConfigOutputReference) Seed() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"seed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigShapConfigOutputReference) SeedInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"seedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigShapConfigOutputReference) ShapBaselineConfig() SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigShapConfigShapBaselineConfigOutputReference {
	var returns SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigShapConfigShapBaselineConfigOutputReference
	_jsii_.Get(
		j,
		"shapBaselineConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigShapConfigOutputReference) ShapBaselineConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shapBaselineConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigShapConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigShapConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigShapConfigOutputReference) TextConfig() SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigShapConfigTextConfigOutputReference {
	var returns SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigShapConfigTextConfigOutputReference
	_jsii_.Get(
		j,
		"textConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigShapConfigOutputReference) TextConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"textConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigShapConfigOutputReference) UseLogit() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useLogit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigShapConfigOutputReference) UseLogitInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useLogitInput",
		&returns,
	)
	return returns
}


func NewSagemakerEndpointConfigExplainerConfigClarifyExplainerConfigShapConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigShapConfigOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerEndpointConfigExplainerConfigClarifyExplainerConfigShapConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigShapConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.sagemakerEndpointConfig.SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigShapConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSagemakerEndpointConfigExplainerConfigClarifyExplainerConfigShapConfigOutputReference_Override(s SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigShapConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.sagemakerEndpointConfig.SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigShapConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigShapConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigShapConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigShapConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigShapConfigOutputReference)SetNumberOfSamples(val *float64) {
	if err := j.validateSetNumberOfSamplesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numberOfSamples",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigShapConfigOutputReference)SetSeed(val *float64) {
	if err := j.validateSetSeedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"seed",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigShapConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigShapConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigShapConfigOutputReference)SetUseLogit(val interface{}) {
	if err := j.validateSetUseLogitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useLogit",
		val,
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigShapConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigShapConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigShapConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigShapConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigShapConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigShapConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigShapConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigShapConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigShapConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigShapConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigShapConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigShapConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigShapConfigOutputReference) PutShapBaselineConfig(value *SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigShapConfigShapBaselineConfig) {
	if err := s.validatePutShapBaselineConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putShapBaselineConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigShapConfigOutputReference) PutTextConfig(value *SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigShapConfigTextConfig) {
	if err := s.validatePutTextConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTextConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigShapConfigOutputReference) ResetNumberOfSamples() {
	_jsii_.InvokeVoid(
		s,
		"resetNumberOfSamples",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigShapConfigOutputReference) ResetSeed() {
	_jsii_.InvokeVoid(
		s,
		"resetSeed",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigShapConfigOutputReference) ResetShapBaselineConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetShapBaselineConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigShapConfigOutputReference) ResetTextConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetTextConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigShapConfigOutputReference) ResetUseLogit() {
	_jsii_.InvokeVoid(
		s,
		"resetUseLogit",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigShapConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigShapConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

