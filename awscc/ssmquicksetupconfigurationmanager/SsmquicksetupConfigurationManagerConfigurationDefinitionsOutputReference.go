// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ssmquicksetupconfigurationmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/ssmquicksetupconfigurationmanager/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SsmquicksetupConfigurationManagerConfigurationDefinitionsOutputReference interface {
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
	Id() *string
	SetId(val *string)
	IdInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LocalDeploymentAdministrationRoleArn() *string
	SetLocalDeploymentAdministrationRoleArn(val *string)
	LocalDeploymentAdministrationRoleArnInput() *string
	LocalDeploymentExecutionRoleName() *string
	SetLocalDeploymentExecutionRoleName(val *string)
	LocalDeploymentExecutionRoleNameInput() *string
	Parameters() *map[string]*string
	SetParameters(val *map[string]*string)
	ParametersInput() *map[string]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
	TypeVersion() *string
	SetTypeVersion(val *string)
	TypeVersionInput() *string
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
	ResetId()
	ResetLocalDeploymentAdministrationRoleArn()
	ResetLocalDeploymentExecutionRoleName()
	ResetTypeVersion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SsmquicksetupConfigurationManagerConfigurationDefinitionsOutputReference
type jsiiProxy_SsmquicksetupConfigurationManagerConfigurationDefinitionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_SsmquicksetupConfigurationManagerConfigurationDefinitionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmquicksetupConfigurationManagerConfigurationDefinitionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmquicksetupConfigurationManagerConfigurationDefinitionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmquicksetupConfigurationManagerConfigurationDefinitionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmquicksetupConfigurationManagerConfigurationDefinitionsOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmquicksetupConfigurationManagerConfigurationDefinitionsOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmquicksetupConfigurationManagerConfigurationDefinitionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmquicksetupConfigurationManagerConfigurationDefinitionsOutputReference) LocalDeploymentAdministrationRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localDeploymentAdministrationRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmquicksetupConfigurationManagerConfigurationDefinitionsOutputReference) LocalDeploymentAdministrationRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localDeploymentAdministrationRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmquicksetupConfigurationManagerConfigurationDefinitionsOutputReference) LocalDeploymentExecutionRoleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localDeploymentExecutionRoleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmquicksetupConfigurationManagerConfigurationDefinitionsOutputReference) LocalDeploymentExecutionRoleNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localDeploymentExecutionRoleNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmquicksetupConfigurationManagerConfigurationDefinitionsOutputReference) Parameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmquicksetupConfigurationManagerConfigurationDefinitionsOutputReference) ParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmquicksetupConfigurationManagerConfigurationDefinitionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmquicksetupConfigurationManagerConfigurationDefinitionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmquicksetupConfigurationManagerConfigurationDefinitionsOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmquicksetupConfigurationManagerConfigurationDefinitionsOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmquicksetupConfigurationManagerConfigurationDefinitionsOutputReference) TypeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmquicksetupConfigurationManagerConfigurationDefinitionsOutputReference) TypeVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeVersionInput",
		&returns,
	)
	return returns
}


func NewSsmquicksetupConfigurationManagerConfigurationDefinitionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) SsmquicksetupConfigurationManagerConfigurationDefinitionsOutputReference {
	_init_.Initialize()

	if err := validateNewSsmquicksetupConfigurationManagerConfigurationDefinitionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_SsmquicksetupConfigurationManagerConfigurationDefinitionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.ssmquicksetupConfigurationManager.SsmquicksetupConfigurationManagerConfigurationDefinitionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewSsmquicksetupConfigurationManagerConfigurationDefinitionsOutputReference_Override(s SsmquicksetupConfigurationManagerConfigurationDefinitionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.ssmquicksetupConfigurationManager.SsmquicksetupConfigurationManagerConfigurationDefinitionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_SsmquicksetupConfigurationManagerConfigurationDefinitionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SsmquicksetupConfigurationManagerConfigurationDefinitionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SsmquicksetupConfigurationManagerConfigurationDefinitionsOutputReference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SsmquicksetupConfigurationManagerConfigurationDefinitionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SsmquicksetupConfigurationManagerConfigurationDefinitionsOutputReference)SetLocalDeploymentAdministrationRoleArn(val *string) {
	if err := j.validateSetLocalDeploymentAdministrationRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localDeploymentAdministrationRoleArn",
		val,
	)
}

func (j *jsiiProxy_SsmquicksetupConfigurationManagerConfigurationDefinitionsOutputReference)SetLocalDeploymentExecutionRoleName(val *string) {
	if err := j.validateSetLocalDeploymentExecutionRoleNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localDeploymentExecutionRoleName",
		val,
	)
}

func (j *jsiiProxy_SsmquicksetupConfigurationManagerConfigurationDefinitionsOutputReference)SetParameters(val *map[string]*string) {
	if err := j.validateSetParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_SsmquicksetupConfigurationManagerConfigurationDefinitionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SsmquicksetupConfigurationManagerConfigurationDefinitionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SsmquicksetupConfigurationManagerConfigurationDefinitionsOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_SsmquicksetupConfigurationManagerConfigurationDefinitionsOutputReference)SetTypeVersion(val *string) {
	if err := j.validateSetTypeVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"typeVersion",
		val,
	)
}

func (s *jsiiProxy_SsmquicksetupConfigurationManagerConfigurationDefinitionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SsmquicksetupConfigurationManagerConfigurationDefinitionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SsmquicksetupConfigurationManagerConfigurationDefinitionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SsmquicksetupConfigurationManagerConfigurationDefinitionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SsmquicksetupConfigurationManagerConfigurationDefinitionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SsmquicksetupConfigurationManagerConfigurationDefinitionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SsmquicksetupConfigurationManagerConfigurationDefinitionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SsmquicksetupConfigurationManagerConfigurationDefinitionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SsmquicksetupConfigurationManagerConfigurationDefinitionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SsmquicksetupConfigurationManagerConfigurationDefinitionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SsmquicksetupConfigurationManagerConfigurationDefinitionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SsmquicksetupConfigurationManagerConfigurationDefinitionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SsmquicksetupConfigurationManagerConfigurationDefinitionsOutputReference) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmquicksetupConfigurationManagerConfigurationDefinitionsOutputReference) ResetLocalDeploymentAdministrationRoleArn() {
	_jsii_.InvokeVoid(
		s,
		"resetLocalDeploymentAdministrationRoleArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmquicksetupConfigurationManagerConfigurationDefinitionsOutputReference) ResetLocalDeploymentExecutionRoleName() {
	_jsii_.InvokeVoid(
		s,
		"resetLocalDeploymentExecutionRoleName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmquicksetupConfigurationManagerConfigurationDefinitionsOutputReference) ResetTypeVersion() {
	_jsii_.InvokeVoid(
		s,
		"resetTypeVersion",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmquicksetupConfigurationManagerConfigurationDefinitionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_SsmquicksetupConfigurationManagerConfigurationDefinitionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

