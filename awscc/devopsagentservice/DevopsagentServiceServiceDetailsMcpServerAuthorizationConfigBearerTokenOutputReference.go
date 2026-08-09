// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package devopsagentservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/devopsagentservice/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DevopsagentServiceServiceDetailsMcpServerAuthorizationConfigBearerTokenOutputReference interface {
	cdktn.ComplexObject
	AuthorizationHeader() *string
	SetAuthorizationHeader(val *string)
	AuthorizationHeaderInput() *string
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
	TokenName() *string
	SetTokenName(val *string)
	TokenNameInput() *string
	TokenValue() *string
	SetTokenValue(val *string)
	TokenValueInput() *string
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
	ResetAuthorizationHeader()
	ResetTokenName()
	ResetTokenValue()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DevopsagentServiceServiceDetailsMcpServerAuthorizationConfigBearerTokenOutputReference
type jsiiProxy_DevopsagentServiceServiceDetailsMcpServerAuthorizationConfigBearerTokenOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DevopsagentServiceServiceDetailsMcpServerAuthorizationConfigBearerTokenOutputReference) AuthorizationHeader() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentServiceServiceDetailsMcpServerAuthorizationConfigBearerTokenOutputReference) AuthorizationHeaderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentServiceServiceDetailsMcpServerAuthorizationConfigBearerTokenOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentServiceServiceDetailsMcpServerAuthorizationConfigBearerTokenOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentServiceServiceDetailsMcpServerAuthorizationConfigBearerTokenOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentServiceServiceDetailsMcpServerAuthorizationConfigBearerTokenOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentServiceServiceDetailsMcpServerAuthorizationConfigBearerTokenOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentServiceServiceDetailsMcpServerAuthorizationConfigBearerTokenOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentServiceServiceDetailsMcpServerAuthorizationConfigBearerTokenOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentServiceServiceDetailsMcpServerAuthorizationConfigBearerTokenOutputReference) TokenName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentServiceServiceDetailsMcpServerAuthorizationConfigBearerTokenOutputReference) TokenNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentServiceServiceDetailsMcpServerAuthorizationConfigBearerTokenOutputReference) TokenValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentServiceServiceDetailsMcpServerAuthorizationConfigBearerTokenOutputReference) TokenValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenValueInput",
		&returns,
	)
	return returns
}


func NewDevopsagentServiceServiceDetailsMcpServerAuthorizationConfigBearerTokenOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DevopsagentServiceServiceDetailsMcpServerAuthorizationConfigBearerTokenOutputReference {
	_init_.Initialize()

	if err := validateNewDevopsagentServiceServiceDetailsMcpServerAuthorizationConfigBearerTokenOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DevopsagentServiceServiceDetailsMcpServerAuthorizationConfigBearerTokenOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.devopsagentService.DevopsagentServiceServiceDetailsMcpServerAuthorizationConfigBearerTokenOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDevopsagentServiceServiceDetailsMcpServerAuthorizationConfigBearerTokenOutputReference_Override(d DevopsagentServiceServiceDetailsMcpServerAuthorizationConfigBearerTokenOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.devopsagentService.DevopsagentServiceServiceDetailsMcpServerAuthorizationConfigBearerTokenOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DevopsagentServiceServiceDetailsMcpServerAuthorizationConfigBearerTokenOutputReference)SetAuthorizationHeader(val *string) {
	if err := j.validateSetAuthorizationHeaderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authorizationHeader",
		val,
	)
}

func (j *jsiiProxy_DevopsagentServiceServiceDetailsMcpServerAuthorizationConfigBearerTokenOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DevopsagentServiceServiceDetailsMcpServerAuthorizationConfigBearerTokenOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DevopsagentServiceServiceDetailsMcpServerAuthorizationConfigBearerTokenOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DevopsagentServiceServiceDetailsMcpServerAuthorizationConfigBearerTokenOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DevopsagentServiceServiceDetailsMcpServerAuthorizationConfigBearerTokenOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DevopsagentServiceServiceDetailsMcpServerAuthorizationConfigBearerTokenOutputReference)SetTokenName(val *string) {
	if err := j.validateSetTokenNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenName",
		val,
	)
}

func (j *jsiiProxy_DevopsagentServiceServiceDetailsMcpServerAuthorizationConfigBearerTokenOutputReference)SetTokenValue(val *string) {
	if err := j.validateSetTokenValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenValue",
		val,
	)
}

func (d *jsiiProxy_DevopsagentServiceServiceDetailsMcpServerAuthorizationConfigBearerTokenOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevopsagentServiceServiceDetailsMcpServerAuthorizationConfigBearerTokenOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevopsagentServiceServiceDetailsMcpServerAuthorizationConfigBearerTokenOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevopsagentServiceServiceDetailsMcpServerAuthorizationConfigBearerTokenOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevopsagentServiceServiceDetailsMcpServerAuthorizationConfigBearerTokenOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevopsagentServiceServiceDetailsMcpServerAuthorizationConfigBearerTokenOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevopsagentServiceServiceDetailsMcpServerAuthorizationConfigBearerTokenOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevopsagentServiceServiceDetailsMcpServerAuthorizationConfigBearerTokenOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevopsagentServiceServiceDetailsMcpServerAuthorizationConfigBearerTokenOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevopsagentServiceServiceDetailsMcpServerAuthorizationConfigBearerTokenOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevopsagentServiceServiceDetailsMcpServerAuthorizationConfigBearerTokenOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevopsagentServiceServiceDetailsMcpServerAuthorizationConfigBearerTokenOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevopsagentServiceServiceDetailsMcpServerAuthorizationConfigBearerTokenOutputReference) ResetAuthorizationHeader() {
	_jsii_.InvokeVoid(
		d,
		"resetAuthorizationHeader",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevopsagentServiceServiceDetailsMcpServerAuthorizationConfigBearerTokenOutputReference) ResetTokenName() {
	_jsii_.InvokeVoid(
		d,
		"resetTokenName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevopsagentServiceServiceDetailsMcpServerAuthorizationConfigBearerTokenOutputReference) ResetTokenValue() {
	_jsii_.InvokeVoid(
		d,
		"resetTokenValue",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevopsagentServiceServiceDetailsMcpServerAuthorizationConfigBearerTokenOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevopsagentServiceServiceDetailsMcpServerAuthorizationConfigBearerTokenOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

