// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package agentregistryregistryrecord

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/agentregistryregistryrecord/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlCredentialProviderConfigurationsCredentialProviderOauthCredentialProviderOutputReference interface {
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
	CustomParameters() *map[string]*string
	SetCustomParameters(val *map[string]*string)
	CustomParametersInput() *map[string]*string
	// Experimental.
	Fqn() *string
	GrantType() *string
	SetGrantType(val *string)
	GrantTypeInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ProviderArn() *string
	SetProviderArn(val *string)
	ProviderArnInput() *string
	Scopes() *[]*string
	SetScopes(val *[]*string)
	ScopesInput() *[]*string
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
	ResetCustomParameters()
	ResetGrantType()
	ResetProviderArn()
	ResetScopes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlCredentialProviderConfigurationsCredentialProviderOauthCredentialProviderOutputReference
type jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlCredentialProviderConfigurationsCredentialProviderOauthCredentialProviderOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlCredentialProviderConfigurationsCredentialProviderOauthCredentialProviderOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlCredentialProviderConfigurationsCredentialProviderOauthCredentialProviderOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlCredentialProviderConfigurationsCredentialProviderOauthCredentialProviderOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlCredentialProviderConfigurationsCredentialProviderOauthCredentialProviderOutputReference) CustomParameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlCredentialProviderConfigurationsCredentialProviderOauthCredentialProviderOutputReference) CustomParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlCredentialProviderConfigurationsCredentialProviderOauthCredentialProviderOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlCredentialProviderConfigurationsCredentialProviderOauthCredentialProviderOutputReference) GrantType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"grantType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlCredentialProviderConfigurationsCredentialProviderOauthCredentialProviderOutputReference) GrantTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"grantTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlCredentialProviderConfigurationsCredentialProviderOauthCredentialProviderOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlCredentialProviderConfigurationsCredentialProviderOauthCredentialProviderOutputReference) ProviderArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlCredentialProviderConfigurationsCredentialProviderOauthCredentialProviderOutputReference) ProviderArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlCredentialProviderConfigurationsCredentialProviderOauthCredentialProviderOutputReference) Scopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlCredentialProviderConfigurationsCredentialProviderOauthCredentialProviderOutputReference) ScopesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scopesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlCredentialProviderConfigurationsCredentialProviderOauthCredentialProviderOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlCredentialProviderConfigurationsCredentialProviderOauthCredentialProviderOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlCredentialProviderConfigurationsCredentialProviderOauthCredentialProviderOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlCredentialProviderConfigurationsCredentialProviderOauthCredentialProviderOutputReference {
	_init_.Initialize()

	if err := validateNewAgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlCredentialProviderConfigurationsCredentialProviderOauthCredentialProviderOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlCredentialProviderConfigurationsCredentialProviderOauthCredentialProviderOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.agentregistryRegistryRecord.AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlCredentialProviderConfigurationsCredentialProviderOauthCredentialProviderOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlCredentialProviderConfigurationsCredentialProviderOauthCredentialProviderOutputReference_Override(a AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlCredentialProviderConfigurationsCredentialProviderOauthCredentialProviderOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.agentregistryRegistryRecord.AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlCredentialProviderConfigurationsCredentialProviderOauthCredentialProviderOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlCredentialProviderConfigurationsCredentialProviderOauthCredentialProviderOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlCredentialProviderConfigurationsCredentialProviderOauthCredentialProviderOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlCredentialProviderConfigurationsCredentialProviderOauthCredentialProviderOutputReference)SetCustomParameters(val *map[string]*string) {
	if err := j.validateSetCustomParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customParameters",
		val,
	)
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlCredentialProviderConfigurationsCredentialProviderOauthCredentialProviderOutputReference)SetGrantType(val *string) {
	if err := j.validateSetGrantTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"grantType",
		val,
	)
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlCredentialProviderConfigurationsCredentialProviderOauthCredentialProviderOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlCredentialProviderConfigurationsCredentialProviderOauthCredentialProviderOutputReference)SetProviderArn(val *string) {
	if err := j.validateSetProviderArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"providerArn",
		val,
	)
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlCredentialProviderConfigurationsCredentialProviderOauthCredentialProviderOutputReference)SetScopes(val *[]*string) {
	if err := j.validateSetScopesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scopes",
		val,
	)
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlCredentialProviderConfigurationsCredentialProviderOauthCredentialProviderOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlCredentialProviderConfigurationsCredentialProviderOauthCredentialProviderOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlCredentialProviderConfigurationsCredentialProviderOauthCredentialProviderOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlCredentialProviderConfigurationsCredentialProviderOauthCredentialProviderOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlCredentialProviderConfigurationsCredentialProviderOauthCredentialProviderOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlCredentialProviderConfigurationsCredentialProviderOauthCredentialProviderOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlCredentialProviderConfigurationsCredentialProviderOauthCredentialProviderOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlCredentialProviderConfigurationsCredentialProviderOauthCredentialProviderOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlCredentialProviderConfigurationsCredentialProviderOauthCredentialProviderOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlCredentialProviderConfigurationsCredentialProviderOauthCredentialProviderOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlCredentialProviderConfigurationsCredentialProviderOauthCredentialProviderOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlCredentialProviderConfigurationsCredentialProviderOauthCredentialProviderOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlCredentialProviderConfigurationsCredentialProviderOauthCredentialProviderOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlCredentialProviderConfigurationsCredentialProviderOauthCredentialProviderOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlCredentialProviderConfigurationsCredentialProviderOauthCredentialProviderOutputReference) ResetCustomParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlCredentialProviderConfigurationsCredentialProviderOauthCredentialProviderOutputReference) ResetGrantType() {
	_jsii_.InvokeVoid(
		a,
		"resetGrantType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlCredentialProviderConfigurationsCredentialProviderOauthCredentialProviderOutputReference) ResetProviderArn() {
	_jsii_.InvokeVoid(
		a,
		"resetProviderArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlCredentialProviderConfigurationsCredentialProviderOauthCredentialProviderOutputReference) ResetScopes() {
	_jsii_.InvokeVoid(
		a,
		"resetScopes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlCredentialProviderConfigurationsCredentialProviderOauthCredentialProviderOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlCredentialProviderConfigurationsCredentialProviderOauthCredentialProviderOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

