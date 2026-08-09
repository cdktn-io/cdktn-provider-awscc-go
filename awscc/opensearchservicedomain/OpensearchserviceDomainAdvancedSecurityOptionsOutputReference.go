// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package opensearchservicedomain

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/opensearchservicedomain/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OpensearchserviceDomainAdvancedSecurityOptionsOutputReference interface {
	cdktn.ComplexObject
	AnonymousAuthDisableDate() *string
	AnonymousAuthEnabled() interface{}
	SetAnonymousAuthEnabled(val interface{})
	AnonymousAuthEnabledInput() interface{}
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
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	// Experimental.
	Fqn() *string
	IamFederationOptions() OpensearchserviceDomainAdvancedSecurityOptionsIamFederationOptionsOutputReference
	IamFederationOptionsInput() interface{}
	InternalUserDatabaseEnabled() interface{}
	SetInternalUserDatabaseEnabled(val interface{})
	InternalUserDatabaseEnabledInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	JwtOptions() OpensearchserviceDomainAdvancedSecurityOptionsJwtOptionsOutputReference
	JwtOptionsInput() interface{}
	MasterUserOptions() OpensearchserviceDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference
	MasterUserOptionsInput() interface{}
	SamlOptions() OpensearchserviceDomainAdvancedSecurityOptionsSamlOptionsOutputReference
	SamlOptionsInput() interface{}
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
	PutIamFederationOptions(value *OpensearchserviceDomainAdvancedSecurityOptionsIamFederationOptions)
	PutJwtOptions(value *OpensearchserviceDomainAdvancedSecurityOptionsJwtOptions)
	PutMasterUserOptions(value *OpensearchserviceDomainAdvancedSecurityOptionsMasterUserOptions)
	PutSamlOptions(value *OpensearchserviceDomainAdvancedSecurityOptionsSamlOptions)
	ResetAnonymousAuthEnabled()
	ResetEnabled()
	ResetIamFederationOptions()
	ResetInternalUserDatabaseEnabled()
	ResetJwtOptions()
	ResetMasterUserOptions()
	ResetSamlOptions()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OpensearchserviceDomainAdvancedSecurityOptionsOutputReference
type jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsOutputReference) AnonymousAuthDisableDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"anonymousAuthDisableDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsOutputReference) AnonymousAuthEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"anonymousAuthEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsOutputReference) AnonymousAuthEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"anonymousAuthEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsOutputReference) IamFederationOptions() OpensearchserviceDomainAdvancedSecurityOptionsIamFederationOptionsOutputReference {
	var returns OpensearchserviceDomainAdvancedSecurityOptionsIamFederationOptionsOutputReference
	_jsii_.Get(
		j,
		"iamFederationOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsOutputReference) IamFederationOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"iamFederationOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsOutputReference) InternalUserDatabaseEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalUserDatabaseEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsOutputReference) InternalUserDatabaseEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalUserDatabaseEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsOutputReference) JwtOptions() OpensearchserviceDomainAdvancedSecurityOptionsJwtOptionsOutputReference {
	var returns OpensearchserviceDomainAdvancedSecurityOptionsJwtOptionsOutputReference
	_jsii_.Get(
		j,
		"jwtOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsOutputReference) JwtOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jwtOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsOutputReference) MasterUserOptions() OpensearchserviceDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference {
	var returns OpensearchserviceDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference
	_jsii_.Get(
		j,
		"masterUserOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsOutputReference) MasterUserOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"masterUserOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsOutputReference) SamlOptions() OpensearchserviceDomainAdvancedSecurityOptionsSamlOptionsOutputReference {
	var returns OpensearchserviceDomainAdvancedSecurityOptionsSamlOptionsOutputReference
	_jsii_.Get(
		j,
		"samlOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsOutputReference) SamlOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"samlOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewOpensearchserviceDomainAdvancedSecurityOptionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) OpensearchserviceDomainAdvancedSecurityOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewOpensearchserviceDomainAdvancedSecurityOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.opensearchserviceDomain.OpensearchserviceDomainAdvancedSecurityOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOpensearchserviceDomainAdvancedSecurityOptionsOutputReference_Override(o OpensearchserviceDomainAdvancedSecurityOptionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.opensearchserviceDomain.OpensearchserviceDomainAdvancedSecurityOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsOutputReference)SetAnonymousAuthEnabled(val interface{}) {
	if err := j.validateSetAnonymousAuthEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"anonymousAuthEnabled",
		val,
	)
}

func (j *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsOutputReference)SetInternalUserDatabaseEnabled(val interface{}) {
	if err := j.validateSetInternalUserDatabaseEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalUserDatabaseEnabled",
		val,
	)
}

func (j *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsOutputReference) PutIamFederationOptions(value *OpensearchserviceDomainAdvancedSecurityOptionsIamFederationOptions) {
	if err := o.validatePutIamFederationOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putIamFederationOptions",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsOutputReference) PutJwtOptions(value *OpensearchserviceDomainAdvancedSecurityOptionsJwtOptions) {
	if err := o.validatePutJwtOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putJwtOptions",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsOutputReference) PutMasterUserOptions(value *OpensearchserviceDomainAdvancedSecurityOptionsMasterUserOptions) {
	if err := o.validatePutMasterUserOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putMasterUserOptions",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsOutputReference) PutSamlOptions(value *OpensearchserviceDomainAdvancedSecurityOptionsSamlOptions) {
	if err := o.validatePutSamlOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putSamlOptions",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsOutputReference) ResetAnonymousAuthEnabled() {
	_jsii_.InvokeVoid(
		o,
		"resetAnonymousAuthEnabled",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		o,
		"resetEnabled",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsOutputReference) ResetIamFederationOptions() {
	_jsii_.InvokeVoid(
		o,
		"resetIamFederationOptions",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsOutputReference) ResetInternalUserDatabaseEnabled() {
	_jsii_.InvokeVoid(
		o,
		"resetInternalUserDatabaseEnabled",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsOutputReference) ResetJwtOptions() {
	_jsii_.InvokeVoid(
		o,
		"resetJwtOptions",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsOutputReference) ResetMasterUserOptions() {
	_jsii_.InvokeVoid(
		o,
		"resetMasterUserOptions",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsOutputReference) ResetSamlOptions() {
	_jsii_.InvokeVoid(
		o,
		"resetSamlOptions",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := o.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

