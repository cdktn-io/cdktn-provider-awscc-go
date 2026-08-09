// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package elasticloadbalancingv2listener

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/elasticloadbalancingv2listener/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Elasticloadbalancingv2ListenerDefaultActionsOutputReference interface {
	cdktn.ComplexObject
	AuthenticateCognitoConfig() Elasticloadbalancingv2ListenerDefaultActionsAuthenticateCognitoConfigOutputReference
	AuthenticateCognitoConfigInput() interface{}
	AuthenticateOidcConfig() Elasticloadbalancingv2ListenerDefaultActionsAuthenticateOidcConfigOutputReference
	AuthenticateOidcConfigInput() interface{}
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
	FixedResponseConfig() Elasticloadbalancingv2ListenerDefaultActionsFixedResponseConfigOutputReference
	FixedResponseConfigInput() interface{}
	ForwardConfig() Elasticloadbalancingv2ListenerDefaultActionsForwardConfigOutputReference
	ForwardConfigInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	JwtValidationConfig() Elasticloadbalancingv2ListenerDefaultActionsJwtValidationConfigOutputReference
	JwtValidationConfigInput() interface{}
	Order() *float64
	SetOrder(val *float64)
	OrderInput() *float64
	RedirectConfig() Elasticloadbalancingv2ListenerDefaultActionsRedirectConfigOutputReference
	RedirectConfigInput() interface{}
	TargetGroupArn() *string
	SetTargetGroupArn(val *string)
	TargetGroupArnInput() *string
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
	PutAuthenticateCognitoConfig(value *Elasticloadbalancingv2ListenerDefaultActionsAuthenticateCognitoConfig)
	PutAuthenticateOidcConfig(value *Elasticloadbalancingv2ListenerDefaultActionsAuthenticateOidcConfig)
	PutFixedResponseConfig(value *Elasticloadbalancingv2ListenerDefaultActionsFixedResponseConfig)
	PutForwardConfig(value *Elasticloadbalancingv2ListenerDefaultActionsForwardConfig)
	PutJwtValidationConfig(value *Elasticloadbalancingv2ListenerDefaultActionsJwtValidationConfig)
	PutRedirectConfig(value *Elasticloadbalancingv2ListenerDefaultActionsRedirectConfig)
	ResetAuthenticateCognitoConfig()
	ResetAuthenticateOidcConfig()
	ResetFixedResponseConfig()
	ResetForwardConfig()
	ResetJwtValidationConfig()
	ResetOrder()
	ResetRedirectConfig()
	ResetTargetGroupArn()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Elasticloadbalancingv2ListenerDefaultActionsOutputReference
type jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsOutputReference) AuthenticateCognitoConfig() Elasticloadbalancingv2ListenerDefaultActionsAuthenticateCognitoConfigOutputReference {
	var returns Elasticloadbalancingv2ListenerDefaultActionsAuthenticateCognitoConfigOutputReference
	_jsii_.Get(
		j,
		"authenticateCognitoConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsOutputReference) AuthenticateCognitoConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authenticateCognitoConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsOutputReference) AuthenticateOidcConfig() Elasticloadbalancingv2ListenerDefaultActionsAuthenticateOidcConfigOutputReference {
	var returns Elasticloadbalancingv2ListenerDefaultActionsAuthenticateOidcConfigOutputReference
	_jsii_.Get(
		j,
		"authenticateOidcConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsOutputReference) AuthenticateOidcConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authenticateOidcConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsOutputReference) FixedResponseConfig() Elasticloadbalancingv2ListenerDefaultActionsFixedResponseConfigOutputReference {
	var returns Elasticloadbalancingv2ListenerDefaultActionsFixedResponseConfigOutputReference
	_jsii_.Get(
		j,
		"fixedResponseConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsOutputReference) FixedResponseConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fixedResponseConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsOutputReference) ForwardConfig() Elasticloadbalancingv2ListenerDefaultActionsForwardConfigOutputReference {
	var returns Elasticloadbalancingv2ListenerDefaultActionsForwardConfigOutputReference
	_jsii_.Get(
		j,
		"forwardConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsOutputReference) ForwardConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forwardConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsOutputReference) JwtValidationConfig() Elasticloadbalancingv2ListenerDefaultActionsJwtValidationConfigOutputReference {
	var returns Elasticloadbalancingv2ListenerDefaultActionsJwtValidationConfigOutputReference
	_jsii_.Get(
		j,
		"jwtValidationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsOutputReference) JwtValidationConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jwtValidationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsOutputReference) Order() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"order",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsOutputReference) OrderInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"orderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsOutputReference) RedirectConfig() Elasticloadbalancingv2ListenerDefaultActionsRedirectConfigOutputReference {
	var returns Elasticloadbalancingv2ListenerDefaultActionsRedirectConfigOutputReference
	_jsii_.Get(
		j,
		"redirectConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsOutputReference) RedirectConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"redirectConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsOutputReference) TargetGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsOutputReference) TargetGroupArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetGroupArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewElasticloadbalancingv2ListenerDefaultActionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) Elasticloadbalancingv2ListenerDefaultActionsOutputReference {
	_init_.Initialize()

	if err := validateNewElasticloadbalancingv2ListenerDefaultActionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.elasticloadbalancingv2Listener.Elasticloadbalancingv2ListenerDefaultActionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewElasticloadbalancingv2ListenerDefaultActionsOutputReference_Override(e Elasticloadbalancingv2ListenerDefaultActionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.elasticloadbalancingv2Listener.Elasticloadbalancingv2ListenerDefaultActionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsOutputReference)SetOrder(val *float64) {
	if err := j.validateSetOrderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"order",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsOutputReference)SetTargetGroupArn(val *string) {
	if err := j.validateSetTargetGroupArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetGroupArn",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsOutputReference) PutAuthenticateCognitoConfig(value *Elasticloadbalancingv2ListenerDefaultActionsAuthenticateCognitoConfig) {
	if err := e.validatePutAuthenticateCognitoConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putAuthenticateCognitoConfig",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsOutputReference) PutAuthenticateOidcConfig(value *Elasticloadbalancingv2ListenerDefaultActionsAuthenticateOidcConfig) {
	if err := e.validatePutAuthenticateOidcConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putAuthenticateOidcConfig",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsOutputReference) PutFixedResponseConfig(value *Elasticloadbalancingv2ListenerDefaultActionsFixedResponseConfig) {
	if err := e.validatePutFixedResponseConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putFixedResponseConfig",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsOutputReference) PutForwardConfig(value *Elasticloadbalancingv2ListenerDefaultActionsForwardConfig) {
	if err := e.validatePutForwardConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putForwardConfig",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsOutputReference) PutJwtValidationConfig(value *Elasticloadbalancingv2ListenerDefaultActionsJwtValidationConfig) {
	if err := e.validatePutJwtValidationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putJwtValidationConfig",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsOutputReference) PutRedirectConfig(value *Elasticloadbalancingv2ListenerDefaultActionsRedirectConfig) {
	if err := e.validatePutRedirectConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putRedirectConfig",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsOutputReference) ResetAuthenticateCognitoConfig() {
	_jsii_.InvokeVoid(
		e,
		"resetAuthenticateCognitoConfig",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsOutputReference) ResetAuthenticateOidcConfig() {
	_jsii_.InvokeVoid(
		e,
		"resetAuthenticateOidcConfig",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsOutputReference) ResetFixedResponseConfig() {
	_jsii_.InvokeVoid(
		e,
		"resetFixedResponseConfig",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsOutputReference) ResetForwardConfig() {
	_jsii_.InvokeVoid(
		e,
		"resetForwardConfig",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsOutputReference) ResetJwtValidationConfig() {
	_jsii_.InvokeVoid(
		e,
		"resetJwtValidationConfig",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsOutputReference) ResetOrder() {
	_jsii_.InvokeVoid(
		e,
		"resetOrder",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsOutputReference) ResetRedirectConfig() {
	_jsii_.InvokeVoid(
		e,
		"resetRedirectConfig",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsOutputReference) ResetTargetGroupArn() {
	_jsii_.InvokeVoid(
		e,
		"resetTargetGroupArn",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := e.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

