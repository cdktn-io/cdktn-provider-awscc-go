// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerendpointconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/sagemakerendpointconfig/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigOutputReference interface {
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
	EnableExplanations() *string
	SetEnableExplanations(val *string)
	EnableExplanationsInput() *string
	// Experimental.
	Fqn() *string
	InferenceConfig() SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference
	InferenceConfigInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ShapConfig() SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigShapConfigOutputReference
	ShapConfigInput() interface{}
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
	PutInferenceConfig(value *SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfig)
	PutShapConfig(value *SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigShapConfig)
	ResetEnableExplanations()
	ResetInferenceConfig()
	ResetShapConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigOutputReference
type jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigOutputReference) EnableExplanations() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enableExplanations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigOutputReference) EnableExplanationsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enableExplanationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigOutputReference) InferenceConfig() SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference {
	var returns SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference
	_jsii_.Get(
		j,
		"inferenceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigOutputReference) InferenceConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inferenceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigOutputReference) ShapConfig() SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigShapConfigOutputReference {
	var returns SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigShapConfigOutputReference
	_jsii_.Get(
		j,
		"shapConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigOutputReference) ShapConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shapConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSagemakerEndpointConfigExplainerConfigClarifyExplainerConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerEndpointConfigExplainerConfigClarifyExplainerConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.sagemakerEndpointConfig.SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSagemakerEndpointConfigExplainerConfigClarifyExplainerConfigOutputReference_Override(s SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.sagemakerEndpointConfig.SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigOutputReference)SetEnableExplanations(val *string) {
	if err := j.validateSetEnableExplanationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableExplanations",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigOutputReference) PutInferenceConfig(value *SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfig) {
	if err := s.validatePutInferenceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putInferenceConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigOutputReference) PutShapConfig(value *SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigShapConfig) {
	if err := s.validatePutShapConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putShapConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigOutputReference) ResetEnableExplanations() {
	_jsii_.InvokeVoid(
		s,
		"resetEnableExplanations",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigOutputReference) ResetInferenceConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetInferenceConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigOutputReference) ResetShapConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetShapConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

