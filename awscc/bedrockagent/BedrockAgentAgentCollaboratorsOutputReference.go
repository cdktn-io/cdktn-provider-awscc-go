// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/bedrockagent/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BedrockAgentAgentCollaboratorsOutputReference interface {
	cdktn.ComplexObject
	AgentDescriptor() BedrockAgentAgentCollaboratorsAgentDescriptorOutputReference
	AgentDescriptorInput() interface{}
	CollaborationInstruction() *string
	SetCollaborationInstruction(val *string)
	CollaborationInstructionInput() *string
	CollaboratorName() *string
	SetCollaboratorName(val *string)
	CollaboratorNameInput() *string
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
	RelayConversationHistory() *string
	SetRelayConversationHistory(val *string)
	RelayConversationHistoryInput() *string
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
	PutAgentDescriptor(value *BedrockAgentAgentCollaboratorsAgentDescriptor)
	ResetAgentDescriptor()
	ResetCollaborationInstruction()
	ResetCollaboratorName()
	ResetRelayConversationHistory()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BedrockAgentAgentCollaboratorsOutputReference
type jsiiProxy_BedrockAgentAgentCollaboratorsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_BedrockAgentAgentCollaboratorsOutputReference) AgentDescriptor() BedrockAgentAgentCollaboratorsAgentDescriptorOutputReference {
	var returns BedrockAgentAgentCollaboratorsAgentDescriptorOutputReference
	_jsii_.Get(
		j,
		"agentDescriptor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgentAgentCollaboratorsOutputReference) AgentDescriptorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"agentDescriptorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgentAgentCollaboratorsOutputReference) CollaborationInstruction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collaborationInstruction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgentAgentCollaboratorsOutputReference) CollaborationInstructionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collaborationInstructionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgentAgentCollaboratorsOutputReference) CollaboratorName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collaboratorName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgentAgentCollaboratorsOutputReference) CollaboratorNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collaboratorNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgentAgentCollaboratorsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgentAgentCollaboratorsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgentAgentCollaboratorsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgentAgentCollaboratorsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgentAgentCollaboratorsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgentAgentCollaboratorsOutputReference) RelayConversationHistory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"relayConversationHistory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgentAgentCollaboratorsOutputReference) RelayConversationHistoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"relayConversationHistoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgentAgentCollaboratorsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgentAgentCollaboratorsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBedrockAgentAgentCollaboratorsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) BedrockAgentAgentCollaboratorsOutputReference {
	_init_.Initialize()

	if err := validateNewBedrockAgentAgentCollaboratorsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_BedrockAgentAgentCollaboratorsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.bedrockAgent.BedrockAgentAgentCollaboratorsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewBedrockAgentAgentCollaboratorsOutputReference_Override(b BedrockAgentAgentCollaboratorsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.bedrockAgent.BedrockAgentAgentCollaboratorsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		b,
	)
}

func (j *jsiiProxy_BedrockAgentAgentCollaboratorsOutputReference)SetCollaborationInstruction(val *string) {
	if err := j.validateSetCollaborationInstructionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"collaborationInstruction",
		val,
	)
}

func (j *jsiiProxy_BedrockAgentAgentCollaboratorsOutputReference)SetCollaboratorName(val *string) {
	if err := j.validateSetCollaboratorNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"collaboratorName",
		val,
	)
}

func (j *jsiiProxy_BedrockAgentAgentCollaboratorsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BedrockAgentAgentCollaboratorsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BedrockAgentAgentCollaboratorsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BedrockAgentAgentCollaboratorsOutputReference)SetRelayConversationHistory(val *string) {
	if err := j.validateSetRelayConversationHistoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"relayConversationHistory",
		val,
	)
}

func (j *jsiiProxy_BedrockAgentAgentCollaboratorsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BedrockAgentAgentCollaboratorsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BedrockAgentAgentCollaboratorsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockAgentAgentCollaboratorsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockAgentAgentCollaboratorsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockAgentAgentCollaboratorsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockAgentAgentCollaboratorsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockAgentAgentCollaboratorsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockAgentAgentCollaboratorsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockAgentAgentCollaboratorsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockAgentAgentCollaboratorsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockAgentAgentCollaboratorsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockAgentAgentCollaboratorsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockAgentAgentCollaboratorsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockAgentAgentCollaboratorsOutputReference) PutAgentDescriptor(value *BedrockAgentAgentCollaboratorsAgentDescriptor) {
	if err := b.validatePutAgentDescriptorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putAgentDescriptor",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockAgentAgentCollaboratorsOutputReference) ResetAgentDescriptor() {
	_jsii_.InvokeVoid(
		b,
		"resetAgentDescriptor",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockAgentAgentCollaboratorsOutputReference) ResetCollaborationInstruction() {
	_jsii_.InvokeVoid(
		b,
		"resetCollaborationInstruction",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockAgentAgentCollaboratorsOutputReference) ResetCollaboratorName() {
	_jsii_.InvokeVoid(
		b,
		"resetCollaboratorName",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockAgentAgentCollaboratorsOutputReference) ResetRelayConversationHistory() {
	_jsii_.InvokeVoid(
		b,
		"resetRelayConversationHistory",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockAgentAgentCollaboratorsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := b.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockAgentAgentCollaboratorsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

