// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ssmincidentsresponseplan

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/ssmincidentsresponseplan/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SsmincidentsResponsePlanActionsSsmAutomationOutputReference interface {
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
	DocumentName() *string
	SetDocumentName(val *string)
	DocumentNameInput() *string
	DocumentVersion() *string
	SetDocumentVersion(val *string)
	DocumentVersionInput() *string
	DynamicParameters() SsmincidentsResponsePlanActionsSsmAutomationDynamicParametersList
	DynamicParametersInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Parameters() SsmincidentsResponsePlanActionsSsmAutomationParametersList
	ParametersInput() interface{}
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
	TargetAccount() *string
	SetTargetAccount(val *string)
	TargetAccountInput() *string
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
	PutDynamicParameters(value interface{})
	PutParameters(value interface{})
	ResetDocumentName()
	ResetDocumentVersion()
	ResetDynamicParameters()
	ResetParameters()
	ResetRoleArn()
	ResetTargetAccount()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SsmincidentsResponsePlanActionsSsmAutomationOutputReference
type jsiiProxy_SsmincidentsResponsePlanActionsSsmAutomationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_SsmincidentsResponsePlanActionsSsmAutomationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmincidentsResponsePlanActionsSsmAutomationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmincidentsResponsePlanActionsSsmAutomationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmincidentsResponsePlanActionsSsmAutomationOutputReference) DocumentName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmincidentsResponsePlanActionsSsmAutomationOutputReference) DocumentNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmincidentsResponsePlanActionsSsmAutomationOutputReference) DocumentVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmincidentsResponsePlanActionsSsmAutomationOutputReference) DocumentVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmincidentsResponsePlanActionsSsmAutomationOutputReference) DynamicParameters() SsmincidentsResponsePlanActionsSsmAutomationDynamicParametersList {
	var returns SsmincidentsResponsePlanActionsSsmAutomationDynamicParametersList
	_jsii_.Get(
		j,
		"dynamicParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmincidentsResponsePlanActionsSsmAutomationOutputReference) DynamicParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynamicParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmincidentsResponsePlanActionsSsmAutomationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmincidentsResponsePlanActionsSsmAutomationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmincidentsResponsePlanActionsSsmAutomationOutputReference) Parameters() SsmincidentsResponsePlanActionsSsmAutomationParametersList {
	var returns SsmincidentsResponsePlanActionsSsmAutomationParametersList
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmincidentsResponsePlanActionsSsmAutomationOutputReference) ParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmincidentsResponsePlanActionsSsmAutomationOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmincidentsResponsePlanActionsSsmAutomationOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmincidentsResponsePlanActionsSsmAutomationOutputReference) TargetAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmincidentsResponsePlanActionsSsmAutomationOutputReference) TargetAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmincidentsResponsePlanActionsSsmAutomationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmincidentsResponsePlanActionsSsmAutomationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSsmincidentsResponsePlanActionsSsmAutomationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) SsmincidentsResponsePlanActionsSsmAutomationOutputReference {
	_init_.Initialize()

	if err := validateNewSsmincidentsResponsePlanActionsSsmAutomationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SsmincidentsResponsePlanActionsSsmAutomationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.ssmincidentsResponsePlan.SsmincidentsResponsePlanActionsSsmAutomationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSsmincidentsResponsePlanActionsSsmAutomationOutputReference_Override(s SsmincidentsResponsePlanActionsSsmAutomationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.ssmincidentsResponsePlan.SsmincidentsResponsePlanActionsSsmAutomationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SsmincidentsResponsePlanActionsSsmAutomationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SsmincidentsResponsePlanActionsSsmAutomationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SsmincidentsResponsePlanActionsSsmAutomationOutputReference)SetDocumentName(val *string) {
	if err := j.validateSetDocumentNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"documentName",
		val,
	)
}

func (j *jsiiProxy_SsmincidentsResponsePlanActionsSsmAutomationOutputReference)SetDocumentVersion(val *string) {
	if err := j.validateSetDocumentVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"documentVersion",
		val,
	)
}

func (j *jsiiProxy_SsmincidentsResponsePlanActionsSsmAutomationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SsmincidentsResponsePlanActionsSsmAutomationOutputReference)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_SsmincidentsResponsePlanActionsSsmAutomationOutputReference)SetTargetAccount(val *string) {
	if err := j.validateSetTargetAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetAccount",
		val,
	)
}

func (j *jsiiProxy_SsmincidentsResponsePlanActionsSsmAutomationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SsmincidentsResponsePlanActionsSsmAutomationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SsmincidentsResponsePlanActionsSsmAutomationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SsmincidentsResponsePlanActionsSsmAutomationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SsmincidentsResponsePlanActionsSsmAutomationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SsmincidentsResponsePlanActionsSsmAutomationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SsmincidentsResponsePlanActionsSsmAutomationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SsmincidentsResponsePlanActionsSsmAutomationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SsmincidentsResponsePlanActionsSsmAutomationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SsmincidentsResponsePlanActionsSsmAutomationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SsmincidentsResponsePlanActionsSsmAutomationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SsmincidentsResponsePlanActionsSsmAutomationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SsmincidentsResponsePlanActionsSsmAutomationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SsmincidentsResponsePlanActionsSsmAutomationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SsmincidentsResponsePlanActionsSsmAutomationOutputReference) PutDynamicParameters(value interface{}) {
	if err := s.validatePutDynamicParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putDynamicParameters",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SsmincidentsResponsePlanActionsSsmAutomationOutputReference) PutParameters(value interface{}) {
	if err := s.validatePutParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putParameters",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SsmincidentsResponsePlanActionsSsmAutomationOutputReference) ResetDocumentName() {
	_jsii_.InvokeVoid(
		s,
		"resetDocumentName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmincidentsResponsePlanActionsSsmAutomationOutputReference) ResetDocumentVersion() {
	_jsii_.InvokeVoid(
		s,
		"resetDocumentVersion",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmincidentsResponsePlanActionsSsmAutomationOutputReference) ResetDynamicParameters() {
	_jsii_.InvokeVoid(
		s,
		"resetDynamicParameters",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmincidentsResponsePlanActionsSsmAutomationOutputReference) ResetParameters() {
	_jsii_.InvokeVoid(
		s,
		"resetParameters",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmincidentsResponsePlanActionsSsmAutomationOutputReference) ResetRoleArn() {
	_jsii_.InvokeVoid(
		s,
		"resetRoleArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmincidentsResponsePlanActionsSsmAutomationOutputReference) ResetTargetAccount() {
	_jsii_.InvokeVoid(
		s,
		"resetTargetAccount",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmincidentsResponsePlanActionsSsmAutomationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_SsmincidentsResponsePlanActionsSsmAutomationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

