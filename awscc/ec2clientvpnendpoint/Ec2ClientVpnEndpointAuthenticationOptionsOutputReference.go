// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2clientvpnendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/ec2clientvpnendpoint/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Ec2ClientVpnEndpointAuthenticationOptionsOutputReference interface {
	cdktn.ComplexObject
	ActiveDirectory() Ec2ClientVpnEndpointAuthenticationOptionsActiveDirectoryOutputReference
	ActiveDirectoryInput() interface{}
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
	FederatedAuthentication() Ec2ClientVpnEndpointAuthenticationOptionsFederatedAuthenticationOutputReference
	FederatedAuthenticationInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MutualAuthentication() Ec2ClientVpnEndpointAuthenticationOptionsMutualAuthenticationOutputReference
	MutualAuthenticationInput() interface{}
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
	PutActiveDirectory(value *Ec2ClientVpnEndpointAuthenticationOptionsActiveDirectory)
	PutFederatedAuthentication(value *Ec2ClientVpnEndpointAuthenticationOptionsFederatedAuthentication)
	PutMutualAuthentication(value *Ec2ClientVpnEndpointAuthenticationOptionsMutualAuthentication)
	ResetActiveDirectory()
	ResetFederatedAuthentication()
	ResetMutualAuthentication()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Ec2ClientVpnEndpointAuthenticationOptionsOutputReference
type jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference) ActiveDirectory() Ec2ClientVpnEndpointAuthenticationOptionsActiveDirectoryOutputReference {
	var returns Ec2ClientVpnEndpointAuthenticationOptionsActiveDirectoryOutputReference
	_jsii_.Get(
		j,
		"activeDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference) ActiveDirectoryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"activeDirectoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference) FederatedAuthentication() Ec2ClientVpnEndpointAuthenticationOptionsFederatedAuthenticationOutputReference {
	var returns Ec2ClientVpnEndpointAuthenticationOptionsFederatedAuthenticationOutputReference
	_jsii_.Get(
		j,
		"federatedAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference) FederatedAuthenticationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"federatedAuthenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference) MutualAuthentication() Ec2ClientVpnEndpointAuthenticationOptionsMutualAuthenticationOutputReference {
	var returns Ec2ClientVpnEndpointAuthenticationOptionsMutualAuthenticationOutputReference
	_jsii_.Get(
		j,
		"mutualAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference) MutualAuthenticationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mutualAuthenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewEc2ClientVpnEndpointAuthenticationOptionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) Ec2ClientVpnEndpointAuthenticationOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewEc2ClientVpnEndpointAuthenticationOptionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.ec2ClientVpnEndpoint.Ec2ClientVpnEndpointAuthenticationOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEc2ClientVpnEndpointAuthenticationOptionsOutputReference_Override(e Ec2ClientVpnEndpointAuthenticationOptionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.ec2ClientVpnEndpoint.Ec2ClientVpnEndpointAuthenticationOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (e *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference) PutActiveDirectory(value *Ec2ClientVpnEndpointAuthenticationOptionsActiveDirectory) {
	if err := e.validatePutActiveDirectoryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putActiveDirectory",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference) PutFederatedAuthentication(value *Ec2ClientVpnEndpointAuthenticationOptionsFederatedAuthentication) {
	if err := e.validatePutFederatedAuthenticationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putFederatedAuthentication",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference) PutMutualAuthentication(value *Ec2ClientVpnEndpointAuthenticationOptionsMutualAuthentication) {
	if err := e.validatePutMutualAuthenticationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putMutualAuthentication",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference) ResetActiveDirectory() {
	_jsii_.InvokeVoid(
		e,
		"resetActiveDirectory",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference) ResetFederatedAuthentication() {
	_jsii_.InvokeVoid(
		e,
		"resetFederatedAuthentication",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference) ResetMutualAuthentication() {
	_jsii_.InvokeVoid(
		e,
		"resetMutualAuthentication",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (e *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

