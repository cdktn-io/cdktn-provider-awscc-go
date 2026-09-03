// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerendpointconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/sagemakerendpointconfig/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference interface {
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
	ContentTemplate() *string
	SetContentTemplate(val *string)
	ContentTemplateInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	FeatureHeaders() *[]*string
	SetFeatureHeaders(val *[]*string)
	FeatureHeadersInput() *[]*string
	FeaturesAttribute() *string
	SetFeaturesAttribute(val *string)
	FeaturesAttributeInput() *string
	FeatureTypes() *[]*string
	SetFeatureTypes(val *[]*string)
	FeatureTypesInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LabelAttribute() *string
	SetLabelAttribute(val *string)
	LabelAttributeInput() *string
	LabelHeaders() *[]*string
	SetLabelHeaders(val *[]*string)
	LabelHeadersInput() *[]*string
	LabelIndex() *float64
	SetLabelIndex(val *float64)
	LabelIndexInput() *float64
	MaxPayloadInMb() *float64
	SetMaxPayloadInMb(val *float64)
	MaxPayloadInMbInput() *float64
	MaxRecordCount() *float64
	SetMaxRecordCount(val *float64)
	MaxRecordCountInput() *float64
	ProbabilityAttribute() *string
	SetProbabilityAttribute(val *string)
	ProbabilityAttributeInput() *string
	ProbabilityIndex() *float64
	SetProbabilityIndex(val *float64)
	ProbabilityIndexInput() *float64
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
	ResetContentTemplate()
	ResetFeatureHeaders()
	ResetFeaturesAttribute()
	ResetFeatureTypes()
	ResetLabelAttribute()
	ResetLabelHeaders()
	ResetLabelIndex()
	ResetMaxPayloadInMb()
	ResetMaxRecordCount()
	ResetProbabilityAttribute()
	ResetProbabilityIndex()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference
type jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference) ContentTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference) ContentTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference) FeatureHeaders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"featureHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference) FeatureHeadersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"featureHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference) FeaturesAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"featuresAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference) FeaturesAttributeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"featuresAttributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference) FeatureTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"featureTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference) FeatureTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"featureTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference) LabelAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference) LabelAttributeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelAttributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference) LabelHeaders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"labelHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference) LabelHeadersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"labelHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference) LabelIndex() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"labelIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference) LabelIndexInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"labelIndexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference) MaxPayloadInMb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxPayloadInMb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference) MaxPayloadInMbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxPayloadInMbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference) MaxRecordCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRecordCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference) MaxRecordCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRecordCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference) ProbabilityAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"probabilityAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference) ProbabilityAttributeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"probabilityAttributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference) ProbabilityIndex() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"probabilityIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference) ProbabilityIndexInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"probabilityIndexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.sagemakerEndpointConfig.SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference_Override(s SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.sagemakerEndpointConfig.SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference)SetContentTemplate(val *string) {
	if err := j.validateSetContentTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentTemplate",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference)SetFeatureHeaders(val *[]*string) {
	if err := j.validateSetFeatureHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"featureHeaders",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference)SetFeaturesAttribute(val *string) {
	if err := j.validateSetFeaturesAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"featuresAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference)SetFeatureTypes(val *[]*string) {
	if err := j.validateSetFeatureTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"featureTypes",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference)SetLabelAttribute(val *string) {
	if err := j.validateSetLabelAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labelAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference)SetLabelHeaders(val *[]*string) {
	if err := j.validateSetLabelHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labelHeaders",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference)SetLabelIndex(val *float64) {
	if err := j.validateSetLabelIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labelIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference)SetMaxPayloadInMb(val *float64) {
	if err := j.validateSetMaxPayloadInMbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxPayloadInMb",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference)SetMaxRecordCount(val *float64) {
	if err := j.validateSetMaxRecordCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxRecordCount",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference)SetProbabilityAttribute(val *string) {
	if err := j.validateSetProbabilityAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"probabilityAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference)SetProbabilityIndex(val *float64) {
	if err := j.validateSetProbabilityIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"probabilityIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference) ResetContentTemplate() {
	_jsii_.InvokeVoid(
		s,
		"resetContentTemplate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference) ResetFeatureHeaders() {
	_jsii_.InvokeVoid(
		s,
		"resetFeatureHeaders",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference) ResetFeaturesAttribute() {
	_jsii_.InvokeVoid(
		s,
		"resetFeaturesAttribute",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference) ResetFeatureTypes() {
	_jsii_.InvokeVoid(
		s,
		"resetFeatureTypes",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference) ResetLabelAttribute() {
	_jsii_.InvokeVoid(
		s,
		"resetLabelAttribute",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference) ResetLabelHeaders() {
	_jsii_.InvokeVoid(
		s,
		"resetLabelHeaders",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference) ResetLabelIndex() {
	_jsii_.InvokeVoid(
		s,
		"resetLabelIndex",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference) ResetMaxPayloadInMb() {
	_jsii_.InvokeVoid(
		s,
		"resetMaxPayloadInMb",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference) ResetMaxRecordCount() {
	_jsii_.InvokeVoid(
		s,
		"resetMaxRecordCount",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference) ResetProbabilityAttribute() {
	_jsii_.InvokeVoid(
		s,
		"resetProbabilityAttribute",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference) ResetProbabilityIndex() {
	_jsii_.InvokeVoid(
		s,
		"resetProbabilityIndex",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

