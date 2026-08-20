// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package agentregistryregistryrecord

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/agentregistryregistryrecord/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlOutputReference interface {
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
	CredentialProviderConfigurations() AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlCredentialProviderConfigurationsList
	CredentialProviderConfigurationsInput() interface{}
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
	Url() *string
	SetUrl(val *string)
	UrlInput() *string
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
	PutCredentialProviderConfigurations(value interface{})
	ResetCredentialProviderConfigurations()
	ResetUrl()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlOutputReference
type jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlOutputReference) CredentialProviderConfigurations() AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlCredentialProviderConfigurationsList {
	var returns AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlCredentialProviderConfigurationsList
	_jsii_.Get(
		j,
		"credentialProviderConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlOutputReference) CredentialProviderConfigurationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"credentialProviderConfigurationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlOutputReference) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlOutputReference) UrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlInput",
		&returns,
	)
	return returns
}


func NewAgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlOutputReference {
	_init_.Initialize()

	if err := validateNewAgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.agentregistryRegistryRecord.AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlOutputReference_Override(a AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.agentregistryRegistryRecord.AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlOutputReference)SetUrl(val *string) {
	if err := j.validateSetUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

func (a *jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlOutputReference) PutCredentialProviderConfigurations(value interface{}) {
	if err := a.validatePutCredentialProviderConfigurationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCredentialProviderConfigurations",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlOutputReference) ResetCredentialProviderConfigurations() {
	_jsii_.InvokeVoid(
		a,
		"resetCredentialProviderConfigurations",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlOutputReference) ResetUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

