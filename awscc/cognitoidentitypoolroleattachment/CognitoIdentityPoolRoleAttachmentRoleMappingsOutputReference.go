// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cognitoidentitypoolroleattachment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/cognitoidentitypoolroleattachment/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CognitoIdentityPoolRoleAttachmentRoleMappingsOutputReference interface {
	cdktn.ComplexObject
	AmbiguousRoleResolution() *string
	SetAmbiguousRoleResolution(val *string)
	AmbiguousRoleResolutionInput() *string
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
	IdentityProvider() *string
	SetIdentityProvider(val *string)
	IdentityProviderInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	RulesConfiguration() CognitoIdentityPoolRoleAttachmentRoleMappingsRulesConfigurationOutputReference
	RulesConfigurationInput() interface{}
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
	PutRulesConfiguration(value *CognitoIdentityPoolRoleAttachmentRoleMappingsRulesConfiguration)
	ResetAmbiguousRoleResolution()
	ResetIdentityProvider()
	ResetRulesConfiguration()
	ResetType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CognitoIdentityPoolRoleAttachmentRoleMappingsOutputReference
type jsiiProxy_CognitoIdentityPoolRoleAttachmentRoleMappingsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_CognitoIdentityPoolRoleAttachmentRoleMappingsOutputReference) AmbiguousRoleResolution() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ambiguousRoleResolution",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPoolRoleAttachmentRoleMappingsOutputReference) AmbiguousRoleResolutionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ambiguousRoleResolutionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPoolRoleAttachmentRoleMappingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPoolRoleAttachmentRoleMappingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPoolRoleAttachmentRoleMappingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPoolRoleAttachmentRoleMappingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPoolRoleAttachmentRoleMappingsOutputReference) IdentityProvider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPoolRoleAttachmentRoleMappingsOutputReference) IdentityProviderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityProviderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPoolRoleAttachmentRoleMappingsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPoolRoleAttachmentRoleMappingsOutputReference) RulesConfiguration() CognitoIdentityPoolRoleAttachmentRoleMappingsRulesConfigurationOutputReference {
	var returns CognitoIdentityPoolRoleAttachmentRoleMappingsRulesConfigurationOutputReference
	_jsii_.Get(
		j,
		"rulesConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPoolRoleAttachmentRoleMappingsOutputReference) RulesConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rulesConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPoolRoleAttachmentRoleMappingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPoolRoleAttachmentRoleMappingsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPoolRoleAttachmentRoleMappingsOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPoolRoleAttachmentRoleMappingsOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewCognitoIdentityPoolRoleAttachmentRoleMappingsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectKey *string) CognitoIdentityPoolRoleAttachmentRoleMappingsOutputReference {
	_init_.Initialize()

	if err := validateNewCognitoIdentityPoolRoleAttachmentRoleMappingsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectKey); err != nil {
		panic(err)
	}
	j := jsiiProxy_CognitoIdentityPoolRoleAttachmentRoleMappingsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.cognitoIdentityPoolRoleAttachment.CognitoIdentityPoolRoleAttachmentRoleMappingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectKey},
		&j,
	)

	return &j
}

func NewCognitoIdentityPoolRoleAttachmentRoleMappingsOutputReference_Override(c CognitoIdentityPoolRoleAttachmentRoleMappingsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectKey *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.cognitoIdentityPoolRoleAttachment.CognitoIdentityPoolRoleAttachmentRoleMappingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectKey},
		c,
	)
}

func (j *jsiiProxy_CognitoIdentityPoolRoleAttachmentRoleMappingsOutputReference)SetAmbiguousRoleResolution(val *string) {
	if err := j.validateSetAmbiguousRoleResolutionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ambiguousRoleResolution",
		val,
	)
}

func (j *jsiiProxy_CognitoIdentityPoolRoleAttachmentRoleMappingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CognitoIdentityPoolRoleAttachmentRoleMappingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CognitoIdentityPoolRoleAttachmentRoleMappingsOutputReference)SetIdentityProvider(val *string) {
	if err := j.validateSetIdentityProviderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identityProvider",
		val,
	)
}

func (j *jsiiProxy_CognitoIdentityPoolRoleAttachmentRoleMappingsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CognitoIdentityPoolRoleAttachmentRoleMappingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CognitoIdentityPoolRoleAttachmentRoleMappingsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CognitoIdentityPoolRoleAttachmentRoleMappingsOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (c *jsiiProxy_CognitoIdentityPoolRoleAttachmentRoleMappingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoIdentityPoolRoleAttachmentRoleMappingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoIdentityPoolRoleAttachmentRoleMappingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoIdentityPoolRoleAttachmentRoleMappingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoIdentityPoolRoleAttachmentRoleMappingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoIdentityPoolRoleAttachmentRoleMappingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoIdentityPoolRoleAttachmentRoleMappingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoIdentityPoolRoleAttachmentRoleMappingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoIdentityPoolRoleAttachmentRoleMappingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoIdentityPoolRoleAttachmentRoleMappingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoIdentityPoolRoleAttachmentRoleMappingsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoIdentityPoolRoleAttachmentRoleMappingsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoIdentityPoolRoleAttachmentRoleMappingsOutputReference) PutRulesConfiguration(value *CognitoIdentityPoolRoleAttachmentRoleMappingsRulesConfiguration) {
	if err := c.validatePutRulesConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRulesConfiguration",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CognitoIdentityPoolRoleAttachmentRoleMappingsOutputReference) ResetAmbiguousRoleResolution() {
	_jsii_.InvokeVoid(
		c,
		"resetAmbiguousRoleResolution",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoIdentityPoolRoleAttachmentRoleMappingsOutputReference) ResetIdentityProvider() {
	_jsii_.InvokeVoid(
		c,
		"resetIdentityProvider",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoIdentityPoolRoleAttachmentRoleMappingsOutputReference) ResetRulesConfiguration() {
	_jsii_.InvokeVoid(
		c,
		"resetRulesConfiguration",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoIdentityPoolRoleAttachmentRoleMappingsOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		c,
		"resetType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoIdentityPoolRoleAttachmentRoleMappingsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := c.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoIdentityPoolRoleAttachmentRoleMappingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

