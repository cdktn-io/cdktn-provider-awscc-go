// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerworkteam

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/sagemakerworkteam/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerWorkteamMemberDefinitionsCognitoMemberDefinitionOutputReference interface {
	cdktn.ComplexObject
	CognitoClientId() *string
	SetCognitoClientId(val *string)
	CognitoClientIdInput() *string
	CognitoUserGroup() *string
	SetCognitoUserGroup(val *string)
	CognitoUserGroupInput() *string
	CognitoUserPool() *string
	SetCognitoUserPool(val *string)
	CognitoUserPoolInput() *string
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
	ResetCognitoClientId()
	ResetCognitoUserGroup()
	ResetCognitoUserPool()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerWorkteamMemberDefinitionsCognitoMemberDefinitionOutputReference
type jsiiProxy_SagemakerWorkteamMemberDefinitionsCognitoMemberDefinitionOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_SagemakerWorkteamMemberDefinitionsCognitoMemberDefinitionOutputReference) CognitoClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cognitoClientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerWorkteamMemberDefinitionsCognitoMemberDefinitionOutputReference) CognitoClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cognitoClientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerWorkteamMemberDefinitionsCognitoMemberDefinitionOutputReference) CognitoUserGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cognitoUserGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerWorkteamMemberDefinitionsCognitoMemberDefinitionOutputReference) CognitoUserGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cognitoUserGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerWorkteamMemberDefinitionsCognitoMemberDefinitionOutputReference) CognitoUserPool() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cognitoUserPool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerWorkteamMemberDefinitionsCognitoMemberDefinitionOutputReference) CognitoUserPoolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cognitoUserPoolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerWorkteamMemberDefinitionsCognitoMemberDefinitionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerWorkteamMemberDefinitionsCognitoMemberDefinitionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerWorkteamMemberDefinitionsCognitoMemberDefinitionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerWorkteamMemberDefinitionsCognitoMemberDefinitionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerWorkteamMemberDefinitionsCognitoMemberDefinitionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerWorkteamMemberDefinitionsCognitoMemberDefinitionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerWorkteamMemberDefinitionsCognitoMemberDefinitionOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSagemakerWorkteamMemberDefinitionsCognitoMemberDefinitionOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) SagemakerWorkteamMemberDefinitionsCognitoMemberDefinitionOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerWorkteamMemberDefinitionsCognitoMemberDefinitionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerWorkteamMemberDefinitionsCognitoMemberDefinitionOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.sagemakerWorkteam.SagemakerWorkteamMemberDefinitionsCognitoMemberDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSagemakerWorkteamMemberDefinitionsCognitoMemberDefinitionOutputReference_Override(s SagemakerWorkteamMemberDefinitionsCognitoMemberDefinitionOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.sagemakerWorkteam.SagemakerWorkteamMemberDefinitionsCognitoMemberDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SagemakerWorkteamMemberDefinitionsCognitoMemberDefinitionOutputReference)SetCognitoClientId(val *string) {
	if err := j.validateSetCognitoClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cognitoClientId",
		val,
	)
}

func (j *jsiiProxy_SagemakerWorkteamMemberDefinitionsCognitoMemberDefinitionOutputReference)SetCognitoUserGroup(val *string) {
	if err := j.validateSetCognitoUserGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cognitoUserGroup",
		val,
	)
}

func (j *jsiiProxy_SagemakerWorkteamMemberDefinitionsCognitoMemberDefinitionOutputReference)SetCognitoUserPool(val *string) {
	if err := j.validateSetCognitoUserPoolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cognitoUserPool",
		val,
	)
}

func (j *jsiiProxy_SagemakerWorkteamMemberDefinitionsCognitoMemberDefinitionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerWorkteamMemberDefinitionsCognitoMemberDefinitionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerWorkteamMemberDefinitionsCognitoMemberDefinitionOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerWorkteamMemberDefinitionsCognitoMemberDefinitionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerWorkteamMemberDefinitionsCognitoMemberDefinitionOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SagemakerWorkteamMemberDefinitionsCognitoMemberDefinitionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerWorkteamMemberDefinitionsCognitoMemberDefinitionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SagemakerWorkteamMemberDefinitionsCognitoMemberDefinitionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerWorkteamMemberDefinitionsCognitoMemberDefinitionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SagemakerWorkteamMemberDefinitionsCognitoMemberDefinitionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SagemakerWorkteamMemberDefinitionsCognitoMemberDefinitionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SagemakerWorkteamMemberDefinitionsCognitoMemberDefinitionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SagemakerWorkteamMemberDefinitionsCognitoMemberDefinitionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SagemakerWorkteamMemberDefinitionsCognitoMemberDefinitionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SagemakerWorkteamMemberDefinitionsCognitoMemberDefinitionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SagemakerWorkteamMemberDefinitionsCognitoMemberDefinitionOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerWorkteamMemberDefinitionsCognitoMemberDefinitionOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerWorkteamMemberDefinitionsCognitoMemberDefinitionOutputReference) ResetCognitoClientId() {
	_jsii_.InvokeVoid(
		s,
		"resetCognitoClientId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerWorkteamMemberDefinitionsCognitoMemberDefinitionOutputReference) ResetCognitoUserGroup() {
	_jsii_.InvokeVoid(
		s,
		"resetCognitoUserGroup",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerWorkteamMemberDefinitionsCognitoMemberDefinitionOutputReference) ResetCognitoUserPool() {
	_jsii_.InvokeVoid(
		s,
		"resetCognitoUserPool",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerWorkteamMemberDefinitionsCognitoMemberDefinitionOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_SagemakerWorkteamMemberDefinitionsCognitoMemberDefinitionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

