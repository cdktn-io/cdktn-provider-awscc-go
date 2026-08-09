// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sesmailmanagerruleset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/sesmailmanagerruleset/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SesMailManagerRuleSetRulesUnlessStringExpressionEvaluateOutputReference interface {
	cdktn.ComplexObject
	Analysis() SesMailManagerRuleSetRulesUnlessStringExpressionEvaluateAnalysisOutputReference
	AnalysisInput() interface{}
	Attribute() *string
	SetAttribute(val *string)
	AttributeInput() *string
	ClientCertificateAttribute() *string
	SetClientCertificateAttribute(val *string)
	ClientCertificateAttributeInput() *string
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
	MimeHeaderAttribute() *string
	SetMimeHeaderAttribute(val *string)
	MimeHeaderAttributeInput() *string
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
	PutAnalysis(value *SesMailManagerRuleSetRulesUnlessStringExpressionEvaluateAnalysis)
	ResetAnalysis()
	ResetAttribute()
	ResetClientCertificateAttribute()
	ResetMimeHeaderAttribute()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SesMailManagerRuleSetRulesUnlessStringExpressionEvaluateOutputReference
type jsiiProxy_SesMailManagerRuleSetRulesUnlessStringExpressionEvaluateOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesUnlessStringExpressionEvaluateOutputReference) Analysis() SesMailManagerRuleSetRulesUnlessStringExpressionEvaluateAnalysisOutputReference {
	var returns SesMailManagerRuleSetRulesUnlessStringExpressionEvaluateAnalysisOutputReference
	_jsii_.Get(
		j,
		"analysis",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesUnlessStringExpressionEvaluateOutputReference) AnalysisInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"analysisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesUnlessStringExpressionEvaluateOutputReference) Attribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesUnlessStringExpressionEvaluateOutputReference) AttributeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesUnlessStringExpressionEvaluateOutputReference) ClientCertificateAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificateAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesUnlessStringExpressionEvaluateOutputReference) ClientCertificateAttributeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificateAttributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesUnlessStringExpressionEvaluateOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesUnlessStringExpressionEvaluateOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesUnlessStringExpressionEvaluateOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesUnlessStringExpressionEvaluateOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesUnlessStringExpressionEvaluateOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesUnlessStringExpressionEvaluateOutputReference) MimeHeaderAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mimeHeaderAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesUnlessStringExpressionEvaluateOutputReference) MimeHeaderAttributeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mimeHeaderAttributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesUnlessStringExpressionEvaluateOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesUnlessStringExpressionEvaluateOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSesMailManagerRuleSetRulesUnlessStringExpressionEvaluateOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) SesMailManagerRuleSetRulesUnlessStringExpressionEvaluateOutputReference {
	_init_.Initialize()

	if err := validateNewSesMailManagerRuleSetRulesUnlessStringExpressionEvaluateOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SesMailManagerRuleSetRulesUnlessStringExpressionEvaluateOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.sesMailManagerRuleSet.SesMailManagerRuleSetRulesUnlessStringExpressionEvaluateOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSesMailManagerRuleSetRulesUnlessStringExpressionEvaluateOutputReference_Override(s SesMailManagerRuleSetRulesUnlessStringExpressionEvaluateOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.sesMailManagerRuleSet.SesMailManagerRuleSetRulesUnlessStringExpressionEvaluateOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesUnlessStringExpressionEvaluateOutputReference)SetAttribute(val *string) {
	if err := j.validateSetAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"attribute",
		val,
	)
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesUnlessStringExpressionEvaluateOutputReference)SetClientCertificateAttribute(val *string) {
	if err := j.validateSetClientCertificateAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientCertificateAttribute",
		val,
	)
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesUnlessStringExpressionEvaluateOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesUnlessStringExpressionEvaluateOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesUnlessStringExpressionEvaluateOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesUnlessStringExpressionEvaluateOutputReference)SetMimeHeaderAttribute(val *string) {
	if err := j.validateSetMimeHeaderAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mimeHeaderAttribute",
		val,
	)
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesUnlessStringExpressionEvaluateOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesUnlessStringExpressionEvaluateOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesUnlessStringExpressionEvaluateOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesUnlessStringExpressionEvaluateOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SesMailManagerRuleSetRulesUnlessStringExpressionEvaluateOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SesMailManagerRuleSetRulesUnlessStringExpressionEvaluateOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SesMailManagerRuleSetRulesUnlessStringExpressionEvaluateOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SesMailManagerRuleSetRulesUnlessStringExpressionEvaluateOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SesMailManagerRuleSetRulesUnlessStringExpressionEvaluateOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SesMailManagerRuleSetRulesUnlessStringExpressionEvaluateOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SesMailManagerRuleSetRulesUnlessStringExpressionEvaluateOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SesMailManagerRuleSetRulesUnlessStringExpressionEvaluateOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SesMailManagerRuleSetRulesUnlessStringExpressionEvaluateOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesUnlessStringExpressionEvaluateOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SesMailManagerRuleSetRulesUnlessStringExpressionEvaluateOutputReference) PutAnalysis(value *SesMailManagerRuleSetRulesUnlessStringExpressionEvaluateAnalysis) {
	if err := s.validatePutAnalysisParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAnalysis",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesUnlessStringExpressionEvaluateOutputReference) ResetAnalysis() {
	_jsii_.InvokeVoid(
		s,
		"resetAnalysis",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesUnlessStringExpressionEvaluateOutputReference) ResetAttribute() {
	_jsii_.InvokeVoid(
		s,
		"resetAttribute",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesUnlessStringExpressionEvaluateOutputReference) ResetClientCertificateAttribute() {
	_jsii_.InvokeVoid(
		s,
		"resetClientCertificateAttribute",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesUnlessStringExpressionEvaluateOutputReference) ResetMimeHeaderAttribute() {
	_jsii_.InvokeVoid(
		s,
		"resetMimeHeaderAttribute",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesUnlessStringExpressionEvaluateOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_SesMailManagerRuleSetRulesUnlessStringExpressionEvaluateOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

