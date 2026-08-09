// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package secretsmanagerrotationschedule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/secretsmanagerrotationschedule/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference interface {
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
	ExcludeCharacters() *string
	SetExcludeCharacters(val *string)
	ExcludeCharactersInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	KmsKeyArn() *string
	SetKmsKeyArn(val *string)
	KmsKeyArnInput() *string
	MasterSecretArn() *string
	SetMasterSecretArn(val *string)
	MasterSecretArnInput() *string
	MasterSecretKmsKeyArn() *string
	SetMasterSecretKmsKeyArn(val *string)
	MasterSecretKmsKeyArnInput() *string
	RotationLambdaName() *string
	SetRotationLambdaName(val *string)
	RotationLambdaNameInput() *string
	RotationType() *string
	SetRotationType(val *string)
	RotationTypeInput() *string
	Runtime() *string
	SetRuntime(val *string)
	RuntimeInput() *string
	SuperuserSecretArn() *string
	SetSuperuserSecretArn(val *string)
	SuperuserSecretArnInput() *string
	SuperuserSecretKmsKeyArn() *string
	SetSuperuserSecretKmsKeyArn(val *string)
	SuperuserSecretKmsKeyArnInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	VpcSecurityGroupIds() *string
	SetVpcSecurityGroupIds(val *string)
	VpcSecurityGroupIdsInput() *string
	VpcSubnetIds() *string
	SetVpcSubnetIds(val *string)
	VpcSubnetIdsInput() *string
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
	ResetExcludeCharacters()
	ResetKmsKeyArn()
	ResetMasterSecretArn()
	ResetMasterSecretKmsKeyArn()
	ResetRotationLambdaName()
	ResetRotationType()
	ResetRuntime()
	ResetSuperuserSecretArn()
	ResetSuperuserSecretKmsKeyArn()
	ResetVpcSecurityGroupIds()
	ResetVpcSubnetIds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference
type jsiiProxy_SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference) ExcludeCharacters() *string {
	var returns *string
	_jsii_.Get(
		j,
		"excludeCharacters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference) ExcludeCharactersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"excludeCharactersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference) KmsKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference) KmsKeyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference) MasterSecretArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterSecretArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference) MasterSecretArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterSecretArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference) MasterSecretKmsKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterSecretKmsKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference) MasterSecretKmsKeyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterSecretKmsKeyArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference) RotationLambdaName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rotationLambdaName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference) RotationLambdaNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rotationLambdaNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference) RotationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rotationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference) RotationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rotationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference) Runtime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference) RuntimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference) SuperuserSecretArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"superuserSecretArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference) SuperuserSecretArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"superuserSecretArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference) SuperuserSecretKmsKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"superuserSecretKmsKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference) SuperuserSecretKmsKeyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"superuserSecretKmsKeyArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference) VpcSecurityGroupIds() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcSecurityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference) VpcSecurityGroupIdsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcSecurityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference) VpcSubnetIds() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcSubnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference) VpcSubnetIdsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcSubnetIdsInput",
		&returns,
	)
	return returns
}


func NewSecretsmanagerRotationScheduleHostedRotationLambdaOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference {
	_init_.Initialize()

	if err := validateNewSecretsmanagerRotationScheduleHostedRotationLambdaOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.secretsmanagerRotationSchedule.SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSecretsmanagerRotationScheduleHostedRotationLambdaOutputReference_Override(s SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.secretsmanagerRotationSchedule.SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference)SetExcludeCharacters(val *string) {
	if err := j.validateSetExcludeCharactersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeCharacters",
		val,
	)
}

func (j *jsiiProxy_SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference)SetKmsKeyArn(val *string) {
	if err := j.validateSetKmsKeyArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyArn",
		val,
	)
}

func (j *jsiiProxy_SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference)SetMasterSecretArn(val *string) {
	if err := j.validateSetMasterSecretArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"masterSecretArn",
		val,
	)
}

func (j *jsiiProxy_SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference)SetMasterSecretKmsKeyArn(val *string) {
	if err := j.validateSetMasterSecretKmsKeyArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"masterSecretKmsKeyArn",
		val,
	)
}

func (j *jsiiProxy_SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference)SetRotationLambdaName(val *string) {
	if err := j.validateSetRotationLambdaNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rotationLambdaName",
		val,
	)
}

func (j *jsiiProxy_SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference)SetRotationType(val *string) {
	if err := j.validateSetRotationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rotationType",
		val,
	)
}

func (j *jsiiProxy_SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference)SetRuntime(val *string) {
	if err := j.validateSetRuntimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtime",
		val,
	)
}

func (j *jsiiProxy_SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference)SetSuperuserSecretArn(val *string) {
	if err := j.validateSetSuperuserSecretArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"superuserSecretArn",
		val,
	)
}

func (j *jsiiProxy_SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference)SetSuperuserSecretKmsKeyArn(val *string) {
	if err := j.validateSetSuperuserSecretKmsKeyArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"superuserSecretKmsKeyArn",
		val,
	)
}

func (j *jsiiProxy_SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference)SetVpcSecurityGroupIds(val *string) {
	if err := j.validateSetVpcSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcSecurityGroupIds",
		val,
	)
}

func (j *jsiiProxy_SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference)SetVpcSubnetIds(val *string) {
	if err := j.validateSetVpcSubnetIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcSubnetIds",
		val,
	)
}

func (s *jsiiProxy_SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference) ResetExcludeCharacters() {
	_jsii_.InvokeVoid(
		s,
		"resetExcludeCharacters",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference) ResetKmsKeyArn() {
	_jsii_.InvokeVoid(
		s,
		"resetKmsKeyArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference) ResetMasterSecretArn() {
	_jsii_.InvokeVoid(
		s,
		"resetMasterSecretArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference) ResetMasterSecretKmsKeyArn() {
	_jsii_.InvokeVoid(
		s,
		"resetMasterSecretKmsKeyArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference) ResetRotationLambdaName() {
	_jsii_.InvokeVoid(
		s,
		"resetRotationLambdaName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference) ResetRotationType() {
	_jsii_.InvokeVoid(
		s,
		"resetRotationType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference) ResetRuntime() {
	_jsii_.InvokeVoid(
		s,
		"resetRuntime",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference) ResetSuperuserSecretArn() {
	_jsii_.InvokeVoid(
		s,
		"resetSuperuserSecretArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference) ResetSuperuserSecretKmsKeyArn() {
	_jsii_.InvokeVoid(
		s,
		"resetSuperuserSecretKmsKeyArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference) ResetVpcSecurityGroupIds() {
	_jsii_.InvokeVoid(
		s,
		"resetVpcSecurityGroupIds",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference) ResetVpcSubnetIds() {
	_jsii_.InvokeVoid(
		s,
		"resetVpcSubnetIds",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

