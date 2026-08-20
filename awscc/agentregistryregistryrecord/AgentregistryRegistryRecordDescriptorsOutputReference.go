// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package agentregistryregistryrecord

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/agentregistryregistryrecord/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AgentregistryRegistryRecordDescriptorsOutputReference interface {
	cdktn.ComplexObject
	A2AAgentCard() AgentregistryRegistryRecordDescriptorsA2AAgentCardOutputReference
	A2AAgentCardInput() interface{}
	AgentSkillsDefinition() AgentregistryRegistryRecordDescriptorsAgentSkillsDefinitionOutputReference
	AgentSkillsDefinitionInput() interface{}
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
	Custom() AgentregistryRegistryRecordDescriptorsCustomOutputReference
	CustomInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	McpServer() AgentregistryRegistryRecordDescriptorsMcpServerOutputReference
	McpServerInput() interface{}
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
	PutA2AAgentCard(value *AgentregistryRegistryRecordDescriptorsA2AAgentCard)
	PutAgentSkillsDefinition(value *AgentregistryRegistryRecordDescriptorsAgentSkillsDefinition)
	PutCustom(value *AgentregistryRegistryRecordDescriptorsCustom)
	PutMcpServer(value *AgentregistryRegistryRecordDescriptorsMcpServer)
	ResetA2AAgentCard()
	ResetAgentSkillsDefinition()
	ResetCustom()
	ResetMcpServer()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AgentregistryRegistryRecordDescriptorsOutputReference
type jsiiProxy_AgentregistryRegistryRecordDescriptorsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsOutputReference) A2AAgentCard() AgentregistryRegistryRecordDescriptorsA2AAgentCardOutputReference {
	var returns AgentregistryRegistryRecordDescriptorsA2AAgentCardOutputReference
	_jsii_.Get(
		j,
		"a2AAgentCard",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsOutputReference) A2AAgentCardInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"a2AAgentCardInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsOutputReference) AgentSkillsDefinition() AgentregistryRegistryRecordDescriptorsAgentSkillsDefinitionOutputReference {
	var returns AgentregistryRegistryRecordDescriptorsAgentSkillsDefinitionOutputReference
	_jsii_.Get(
		j,
		"agentSkillsDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsOutputReference) AgentSkillsDefinitionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"agentSkillsDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsOutputReference) Custom() AgentregistryRegistryRecordDescriptorsCustomOutputReference {
	var returns AgentregistryRegistryRecordDescriptorsCustomOutputReference
	_jsii_.Get(
		j,
		"custom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsOutputReference) CustomInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsOutputReference) McpServer() AgentregistryRegistryRecordDescriptorsMcpServerOutputReference {
	var returns AgentregistryRegistryRecordDescriptorsMcpServerOutputReference
	_jsii_.Get(
		j,
		"mcpServer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsOutputReference) McpServerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mcpServerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAgentregistryRegistryRecordDescriptorsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AgentregistryRegistryRecordDescriptorsOutputReference {
	_init_.Initialize()

	if err := validateNewAgentregistryRegistryRecordDescriptorsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AgentregistryRegistryRecordDescriptorsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.agentregistryRegistryRecord.AgentregistryRegistryRecordDescriptorsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAgentregistryRegistryRecordDescriptorsOutputReference_Override(a AgentregistryRegistryRecordDescriptorsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.agentregistryRegistryRecord.AgentregistryRegistryRecordDescriptorsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AgentregistryRegistryRecordDescriptorsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgentregistryRegistryRecordDescriptorsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AgentregistryRegistryRecordDescriptorsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AgentregistryRegistryRecordDescriptorsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AgentregistryRegistryRecordDescriptorsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AgentregistryRegistryRecordDescriptorsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AgentregistryRegistryRecordDescriptorsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AgentregistryRegistryRecordDescriptorsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AgentregistryRegistryRecordDescriptorsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AgentregistryRegistryRecordDescriptorsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AgentregistryRegistryRecordDescriptorsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgentregistryRegistryRecordDescriptorsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AgentregistryRegistryRecordDescriptorsOutputReference) PutA2AAgentCard(value *AgentregistryRegistryRecordDescriptorsA2AAgentCard) {
	if err := a.validatePutA2AAgentCardParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putA2AAgentCard",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AgentregistryRegistryRecordDescriptorsOutputReference) PutAgentSkillsDefinition(value *AgentregistryRegistryRecordDescriptorsAgentSkillsDefinition) {
	if err := a.validatePutAgentSkillsDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAgentSkillsDefinition",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AgentregistryRegistryRecordDescriptorsOutputReference) PutCustom(value *AgentregistryRegistryRecordDescriptorsCustom) {
	if err := a.validatePutCustomParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustom",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AgentregistryRegistryRecordDescriptorsOutputReference) PutMcpServer(value *AgentregistryRegistryRecordDescriptorsMcpServer) {
	if err := a.validatePutMcpServerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMcpServer",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AgentregistryRegistryRecordDescriptorsOutputReference) ResetA2AAgentCard() {
	_jsii_.InvokeVoid(
		a,
		"resetA2AAgentCard",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgentregistryRegistryRecordDescriptorsOutputReference) ResetAgentSkillsDefinition() {
	_jsii_.InvokeVoid(
		a,
		"resetAgentSkillsDefinition",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgentregistryRegistryRecordDescriptorsOutputReference) ResetCustom() {
	_jsii_.InvokeVoid(
		a,
		"resetCustom",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgentregistryRegistryRecordDescriptorsOutputReference) ResetMcpServer() {
	_jsii_.InvokeVoid(
		a,
		"resetMcpServer",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgentregistryRegistryRecordDescriptorsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AgentregistryRegistryRecordDescriptorsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

