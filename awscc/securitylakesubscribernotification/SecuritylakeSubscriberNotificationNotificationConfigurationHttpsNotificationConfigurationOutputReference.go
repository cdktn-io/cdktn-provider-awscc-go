// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securitylakesubscribernotification

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/securitylakesubscribernotification/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SecuritylakeSubscriberNotificationNotificationConfigurationHttpsNotificationConfigurationOutputReference interface {
	cdktn.ComplexObject
	AuthorizationApiKeyName() *string
	SetAuthorizationApiKeyName(val *string)
	AuthorizationApiKeyNameInput() *string
	AuthorizationApiKeyValue() *string
	SetAuthorizationApiKeyValue(val *string)
	AuthorizationApiKeyValueInput() *string
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
	Endpoint() *string
	SetEndpoint(val *string)
	EndpointInput() *string
	// Experimental.
	Fqn() *string
	HttpMethod() *string
	SetHttpMethod(val *string)
	HttpMethodInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	TargetRoleArn() *string
	SetTargetRoleArn(val *string)
	TargetRoleArnInput() *string
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
	ResetAuthorizationApiKeyName()
	ResetAuthorizationApiKeyValue()
	ResetEndpoint()
	ResetHttpMethod()
	ResetTargetRoleArn()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SecuritylakeSubscriberNotificationNotificationConfigurationHttpsNotificationConfigurationOutputReference
type jsiiProxy_SecuritylakeSubscriberNotificationNotificationConfigurationHttpsNotificationConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_SecuritylakeSubscriberNotificationNotificationConfigurationHttpsNotificationConfigurationOutputReference) AuthorizationApiKeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationApiKeyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecuritylakeSubscriberNotificationNotificationConfigurationHttpsNotificationConfigurationOutputReference) AuthorizationApiKeyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationApiKeyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecuritylakeSubscriberNotificationNotificationConfigurationHttpsNotificationConfigurationOutputReference) AuthorizationApiKeyValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationApiKeyValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecuritylakeSubscriberNotificationNotificationConfigurationHttpsNotificationConfigurationOutputReference) AuthorizationApiKeyValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationApiKeyValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecuritylakeSubscriberNotificationNotificationConfigurationHttpsNotificationConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecuritylakeSubscriberNotificationNotificationConfigurationHttpsNotificationConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecuritylakeSubscriberNotificationNotificationConfigurationHttpsNotificationConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecuritylakeSubscriberNotificationNotificationConfigurationHttpsNotificationConfigurationOutputReference) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecuritylakeSubscriberNotificationNotificationConfigurationHttpsNotificationConfigurationOutputReference) EndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecuritylakeSubscriberNotificationNotificationConfigurationHttpsNotificationConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecuritylakeSubscriberNotificationNotificationConfigurationHttpsNotificationConfigurationOutputReference) HttpMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecuritylakeSubscriberNotificationNotificationConfigurationHttpsNotificationConfigurationOutputReference) HttpMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecuritylakeSubscriberNotificationNotificationConfigurationHttpsNotificationConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecuritylakeSubscriberNotificationNotificationConfigurationHttpsNotificationConfigurationOutputReference) TargetRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecuritylakeSubscriberNotificationNotificationConfigurationHttpsNotificationConfigurationOutputReference) TargetRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecuritylakeSubscriberNotificationNotificationConfigurationHttpsNotificationConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecuritylakeSubscriberNotificationNotificationConfigurationHttpsNotificationConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSecuritylakeSubscriberNotificationNotificationConfigurationHttpsNotificationConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) SecuritylakeSubscriberNotificationNotificationConfigurationHttpsNotificationConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewSecuritylakeSubscriberNotificationNotificationConfigurationHttpsNotificationConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SecuritylakeSubscriberNotificationNotificationConfigurationHttpsNotificationConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.securitylakeSubscriberNotification.SecuritylakeSubscriberNotificationNotificationConfigurationHttpsNotificationConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSecuritylakeSubscriberNotificationNotificationConfigurationHttpsNotificationConfigurationOutputReference_Override(s SecuritylakeSubscriberNotificationNotificationConfigurationHttpsNotificationConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.securitylakeSubscriberNotification.SecuritylakeSubscriberNotificationNotificationConfigurationHttpsNotificationConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SecuritylakeSubscriberNotificationNotificationConfigurationHttpsNotificationConfigurationOutputReference)SetAuthorizationApiKeyName(val *string) {
	if err := j.validateSetAuthorizationApiKeyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authorizationApiKeyName",
		val,
	)
}

func (j *jsiiProxy_SecuritylakeSubscriberNotificationNotificationConfigurationHttpsNotificationConfigurationOutputReference)SetAuthorizationApiKeyValue(val *string) {
	if err := j.validateSetAuthorizationApiKeyValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authorizationApiKeyValue",
		val,
	)
}

func (j *jsiiProxy_SecuritylakeSubscriberNotificationNotificationConfigurationHttpsNotificationConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SecuritylakeSubscriberNotificationNotificationConfigurationHttpsNotificationConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SecuritylakeSubscriberNotificationNotificationConfigurationHttpsNotificationConfigurationOutputReference)SetEndpoint(val *string) {
	if err := j.validateSetEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpoint",
		val,
	)
}

func (j *jsiiProxy_SecuritylakeSubscriberNotificationNotificationConfigurationHttpsNotificationConfigurationOutputReference)SetHttpMethod(val *string) {
	if err := j.validateSetHttpMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpMethod",
		val,
	)
}

func (j *jsiiProxy_SecuritylakeSubscriberNotificationNotificationConfigurationHttpsNotificationConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SecuritylakeSubscriberNotificationNotificationConfigurationHttpsNotificationConfigurationOutputReference)SetTargetRoleArn(val *string) {
	if err := j.validateSetTargetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetRoleArn",
		val,
	)
}

func (j *jsiiProxy_SecuritylakeSubscriberNotificationNotificationConfigurationHttpsNotificationConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SecuritylakeSubscriberNotificationNotificationConfigurationHttpsNotificationConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SecuritylakeSubscriberNotificationNotificationConfigurationHttpsNotificationConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecuritylakeSubscriberNotificationNotificationConfigurationHttpsNotificationConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SecuritylakeSubscriberNotificationNotificationConfigurationHttpsNotificationConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SecuritylakeSubscriberNotificationNotificationConfigurationHttpsNotificationConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SecuritylakeSubscriberNotificationNotificationConfigurationHttpsNotificationConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SecuritylakeSubscriberNotificationNotificationConfigurationHttpsNotificationConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SecuritylakeSubscriberNotificationNotificationConfigurationHttpsNotificationConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SecuritylakeSubscriberNotificationNotificationConfigurationHttpsNotificationConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SecuritylakeSubscriberNotificationNotificationConfigurationHttpsNotificationConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SecuritylakeSubscriberNotificationNotificationConfigurationHttpsNotificationConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SecuritylakeSubscriberNotificationNotificationConfigurationHttpsNotificationConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecuritylakeSubscriberNotificationNotificationConfigurationHttpsNotificationConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SecuritylakeSubscriberNotificationNotificationConfigurationHttpsNotificationConfigurationOutputReference) ResetAuthorizationApiKeyName() {
	_jsii_.InvokeVoid(
		s,
		"resetAuthorizationApiKeyName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecuritylakeSubscriberNotificationNotificationConfigurationHttpsNotificationConfigurationOutputReference) ResetAuthorizationApiKeyValue() {
	_jsii_.InvokeVoid(
		s,
		"resetAuthorizationApiKeyValue",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecuritylakeSubscriberNotificationNotificationConfigurationHttpsNotificationConfigurationOutputReference) ResetEndpoint() {
	_jsii_.InvokeVoid(
		s,
		"resetEndpoint",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecuritylakeSubscriberNotificationNotificationConfigurationHttpsNotificationConfigurationOutputReference) ResetHttpMethod() {
	_jsii_.InvokeVoid(
		s,
		"resetHttpMethod",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecuritylakeSubscriberNotificationNotificationConfigurationHttpsNotificationConfigurationOutputReference) ResetTargetRoleArn() {
	_jsii_.InvokeVoid(
		s,
		"resetTargetRoleArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecuritylakeSubscriberNotificationNotificationConfigurationHttpsNotificationConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_SecuritylakeSubscriberNotificationNotificationConfigurationHttpsNotificationConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

