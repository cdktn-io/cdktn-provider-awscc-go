// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityagentagentspace

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/securityagentagentspace/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SecurityagentAgentSpaceAwsResourcesOutputReference interface {
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
	IamRoles() *[]*string
	SetIamRoles(val *[]*string)
	IamRolesInput() *[]*string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LambdaFunctionArns() *[]*string
	SetLambdaFunctionArns(val *[]*string)
	LambdaFunctionArnsInput() *[]*string
	LogGroups() *[]*string
	SetLogGroups(val *[]*string)
	LogGroupsInput() *[]*string
	S3Buckets() *[]*string
	SetS3Buckets(val *[]*string)
	S3BucketsInput() *[]*string
	SecretArns() *[]*string
	SetSecretArns(val *[]*string)
	SecretArnsInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Vpcs() SecurityagentAgentSpaceAwsResourcesVpcsList
	VpcsInput() interface{}
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
	PutVpcs(value interface{})
	ResetIamRoles()
	ResetLambdaFunctionArns()
	ResetLogGroups()
	ResetS3Buckets()
	ResetSecretArns()
	ResetVpcs()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SecurityagentAgentSpaceAwsResourcesOutputReference
type jsiiProxy_SecurityagentAgentSpaceAwsResourcesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_SecurityagentAgentSpaceAwsResourcesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityagentAgentSpaceAwsResourcesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityagentAgentSpaceAwsResourcesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityagentAgentSpaceAwsResourcesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityagentAgentSpaceAwsResourcesOutputReference) IamRoles() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"iamRoles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityagentAgentSpaceAwsResourcesOutputReference) IamRolesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"iamRolesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityagentAgentSpaceAwsResourcesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityagentAgentSpaceAwsResourcesOutputReference) LambdaFunctionArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"lambdaFunctionArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityagentAgentSpaceAwsResourcesOutputReference) LambdaFunctionArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"lambdaFunctionArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityagentAgentSpaceAwsResourcesOutputReference) LogGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"logGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityagentAgentSpaceAwsResourcesOutputReference) LogGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"logGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityagentAgentSpaceAwsResourcesOutputReference) S3Buckets() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"s3Buckets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityagentAgentSpaceAwsResourcesOutputReference) S3BucketsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"s3BucketsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityagentAgentSpaceAwsResourcesOutputReference) SecretArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"secretArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityagentAgentSpaceAwsResourcesOutputReference) SecretArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"secretArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityagentAgentSpaceAwsResourcesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityagentAgentSpaceAwsResourcesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityagentAgentSpaceAwsResourcesOutputReference) Vpcs() SecurityagentAgentSpaceAwsResourcesVpcsList {
	var returns SecurityagentAgentSpaceAwsResourcesVpcsList
	_jsii_.Get(
		j,
		"vpcs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityagentAgentSpaceAwsResourcesOutputReference) VpcsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vpcsInput",
		&returns,
	)
	return returns
}


func NewSecurityagentAgentSpaceAwsResourcesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) SecurityagentAgentSpaceAwsResourcesOutputReference {
	_init_.Initialize()

	if err := validateNewSecurityagentAgentSpaceAwsResourcesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SecurityagentAgentSpaceAwsResourcesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.securityagentAgentSpace.SecurityagentAgentSpaceAwsResourcesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSecurityagentAgentSpaceAwsResourcesOutputReference_Override(s SecurityagentAgentSpaceAwsResourcesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.securityagentAgentSpace.SecurityagentAgentSpaceAwsResourcesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SecurityagentAgentSpaceAwsResourcesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SecurityagentAgentSpaceAwsResourcesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SecurityagentAgentSpaceAwsResourcesOutputReference)SetIamRoles(val *[]*string) {
	if err := j.validateSetIamRolesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iamRoles",
		val,
	)
}

func (j *jsiiProxy_SecurityagentAgentSpaceAwsResourcesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SecurityagentAgentSpaceAwsResourcesOutputReference)SetLambdaFunctionArns(val *[]*string) {
	if err := j.validateSetLambdaFunctionArnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lambdaFunctionArns",
		val,
	)
}

func (j *jsiiProxy_SecurityagentAgentSpaceAwsResourcesOutputReference)SetLogGroups(val *[]*string) {
	if err := j.validateSetLogGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logGroups",
		val,
	)
}

func (j *jsiiProxy_SecurityagentAgentSpaceAwsResourcesOutputReference)SetS3Buckets(val *[]*string) {
	if err := j.validateSetS3BucketsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3Buckets",
		val,
	)
}

func (j *jsiiProxy_SecurityagentAgentSpaceAwsResourcesOutputReference)SetSecretArns(val *[]*string) {
	if err := j.validateSetSecretArnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretArns",
		val,
	)
}

func (j *jsiiProxy_SecurityagentAgentSpaceAwsResourcesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SecurityagentAgentSpaceAwsResourcesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SecurityagentAgentSpaceAwsResourcesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityagentAgentSpaceAwsResourcesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SecurityagentAgentSpaceAwsResourcesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SecurityagentAgentSpaceAwsResourcesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SecurityagentAgentSpaceAwsResourcesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SecurityagentAgentSpaceAwsResourcesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SecurityagentAgentSpaceAwsResourcesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SecurityagentAgentSpaceAwsResourcesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SecurityagentAgentSpaceAwsResourcesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SecurityagentAgentSpaceAwsResourcesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SecurityagentAgentSpaceAwsResourcesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityagentAgentSpaceAwsResourcesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SecurityagentAgentSpaceAwsResourcesOutputReference) PutVpcs(value interface{}) {
	if err := s.validatePutVpcsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putVpcs",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityagentAgentSpaceAwsResourcesOutputReference) ResetIamRoles() {
	_jsii_.InvokeVoid(
		s,
		"resetIamRoles",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityagentAgentSpaceAwsResourcesOutputReference) ResetLambdaFunctionArns() {
	_jsii_.InvokeVoid(
		s,
		"resetLambdaFunctionArns",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityagentAgentSpaceAwsResourcesOutputReference) ResetLogGroups() {
	_jsii_.InvokeVoid(
		s,
		"resetLogGroups",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityagentAgentSpaceAwsResourcesOutputReference) ResetS3Buckets() {
	_jsii_.InvokeVoid(
		s,
		"resetS3Buckets",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityagentAgentSpaceAwsResourcesOutputReference) ResetSecretArns() {
	_jsii_.InvokeVoid(
		s,
		"resetSecretArns",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityagentAgentSpaceAwsResourcesOutputReference) ResetVpcs() {
	_jsii_.InvokeVoid(
		s,
		"resetVpcs",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityagentAgentSpaceAwsResourcesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_SecurityagentAgentSpaceAwsResourcesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

