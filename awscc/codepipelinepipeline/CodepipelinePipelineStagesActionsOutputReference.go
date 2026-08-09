// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package codepipelinepipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/codepipelinepipeline/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CodepipelinePipelineStagesActionsOutputReference interface {
	cdktn.ComplexObject
	ActionTypeId() CodepipelinePipelineStagesActionsActionTypeIdOutputReference
	ActionTypeIdInput() interface{}
	Commands() *[]*string
	SetCommands(val *[]*string)
	CommandsInput() *[]*string
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
	Configuration() *string
	SetConfiguration(val *string)
	ConfigurationInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EnvironmentVariables() CodepipelinePipelineStagesActionsEnvironmentVariablesList
	EnvironmentVariablesInput() interface{}
	// Experimental.
	Fqn() *string
	InputArtifacts() CodepipelinePipelineStagesActionsInputArtifactsList
	InputArtifactsInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	Namespace() *string
	SetNamespace(val *string)
	NamespaceInput() *string
	OutputArtifacts() CodepipelinePipelineStagesActionsOutputArtifactsList
	OutputArtifactsInput() interface{}
	OutputVariables() *[]*string
	SetOutputVariables(val *[]*string)
	OutputVariablesInput() *[]*string
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
	RunOrder() *float64
	SetRunOrder(val *float64)
	RunOrderInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TimeoutInMinutes() *float64
	SetTimeoutInMinutes(val *float64)
	TimeoutInMinutesInput() *float64
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
	PutActionTypeId(value *CodepipelinePipelineStagesActionsActionTypeId)
	PutEnvironmentVariables(value interface{})
	PutInputArtifacts(value interface{})
	PutOutputArtifacts(value interface{})
	ResetCommands()
	ResetConfiguration()
	ResetEnvironmentVariables()
	ResetInputArtifacts()
	ResetNamespace()
	ResetOutputArtifacts()
	ResetOutputVariables()
	ResetRegion()
	ResetRoleArn()
	ResetRunOrder()
	ResetTimeoutInMinutes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CodepipelinePipelineStagesActionsOutputReference
type jsiiProxy_CodepipelinePipelineStagesActionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_CodepipelinePipelineStagesActionsOutputReference) ActionTypeId() CodepipelinePipelineStagesActionsActionTypeIdOutputReference {
	var returns CodepipelinePipelineStagesActionsActionTypeIdOutputReference
	_jsii_.Get(
		j,
		"actionTypeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelinePipelineStagesActionsOutputReference) ActionTypeIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"actionTypeIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelinePipelineStagesActionsOutputReference) Commands() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"commands",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelinePipelineStagesActionsOutputReference) CommandsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"commandsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelinePipelineStagesActionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelinePipelineStagesActionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelinePipelineStagesActionsOutputReference) Configuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelinePipelineStagesActionsOutputReference) ConfigurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelinePipelineStagesActionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelinePipelineStagesActionsOutputReference) EnvironmentVariables() CodepipelinePipelineStagesActionsEnvironmentVariablesList {
	var returns CodepipelinePipelineStagesActionsEnvironmentVariablesList
	_jsii_.Get(
		j,
		"environmentVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelinePipelineStagesActionsOutputReference) EnvironmentVariablesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"environmentVariablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelinePipelineStagesActionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelinePipelineStagesActionsOutputReference) InputArtifacts() CodepipelinePipelineStagesActionsInputArtifactsList {
	var returns CodepipelinePipelineStagesActionsInputArtifactsList
	_jsii_.Get(
		j,
		"inputArtifacts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelinePipelineStagesActionsOutputReference) InputArtifactsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inputArtifactsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelinePipelineStagesActionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelinePipelineStagesActionsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelinePipelineStagesActionsOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelinePipelineStagesActionsOutputReference) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelinePipelineStagesActionsOutputReference) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelinePipelineStagesActionsOutputReference) OutputArtifacts() CodepipelinePipelineStagesActionsOutputArtifactsList {
	var returns CodepipelinePipelineStagesActionsOutputArtifactsList
	_jsii_.Get(
		j,
		"outputArtifacts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelinePipelineStagesActionsOutputReference) OutputArtifactsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"outputArtifactsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelinePipelineStagesActionsOutputReference) OutputVariables() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"outputVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelinePipelineStagesActionsOutputReference) OutputVariablesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"outputVariablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelinePipelineStagesActionsOutputReference) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelinePipelineStagesActionsOutputReference) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelinePipelineStagesActionsOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelinePipelineStagesActionsOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelinePipelineStagesActionsOutputReference) RunOrder() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"runOrder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelinePipelineStagesActionsOutputReference) RunOrderInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"runOrderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelinePipelineStagesActionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelinePipelineStagesActionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelinePipelineStagesActionsOutputReference) TimeoutInMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelinePipelineStagesActionsOutputReference) TimeoutInMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInMinutesInput",
		&returns,
	)
	return returns
}


func NewCodepipelinePipelineStagesActionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CodepipelinePipelineStagesActionsOutputReference {
	_init_.Initialize()

	if err := validateNewCodepipelinePipelineStagesActionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_CodepipelinePipelineStagesActionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.codepipelinePipeline.CodepipelinePipelineStagesActionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCodepipelinePipelineStagesActionsOutputReference_Override(c CodepipelinePipelineStagesActionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.codepipelinePipeline.CodepipelinePipelineStagesActionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CodepipelinePipelineStagesActionsOutputReference)SetCommands(val *[]*string) {
	if err := j.validateSetCommandsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"commands",
		val,
	)
}

func (j *jsiiProxy_CodepipelinePipelineStagesActionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CodepipelinePipelineStagesActionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CodepipelinePipelineStagesActionsOutputReference)SetConfiguration(val *string) {
	if err := j.validateSetConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configuration",
		val,
	)
}

func (j *jsiiProxy_CodepipelinePipelineStagesActionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CodepipelinePipelineStagesActionsOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CodepipelinePipelineStagesActionsOutputReference)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_CodepipelinePipelineStagesActionsOutputReference)SetOutputVariables(val *[]*string) {
	if err := j.validateSetOutputVariablesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputVariables",
		val,
	)
}

func (j *jsiiProxy_CodepipelinePipelineStagesActionsOutputReference)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_CodepipelinePipelineStagesActionsOutputReference)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_CodepipelinePipelineStagesActionsOutputReference)SetRunOrder(val *float64) {
	if err := j.validateSetRunOrderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runOrder",
		val,
	)
}

func (j *jsiiProxy_CodepipelinePipelineStagesActionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CodepipelinePipelineStagesActionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CodepipelinePipelineStagesActionsOutputReference)SetTimeoutInMinutes(val *float64) {
	if err := j.validateSetTimeoutInMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeoutInMinutes",
		val,
	)
}

func (c *jsiiProxy_CodepipelinePipelineStagesActionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodepipelinePipelineStagesActionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CodepipelinePipelineStagesActionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CodepipelinePipelineStagesActionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CodepipelinePipelineStagesActionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CodepipelinePipelineStagesActionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CodepipelinePipelineStagesActionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CodepipelinePipelineStagesActionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CodepipelinePipelineStagesActionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CodepipelinePipelineStagesActionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CodepipelinePipelineStagesActionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodepipelinePipelineStagesActionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CodepipelinePipelineStagesActionsOutputReference) PutActionTypeId(value *CodepipelinePipelineStagesActionsActionTypeId) {
	if err := c.validatePutActionTypeIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putActionTypeId",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodepipelinePipelineStagesActionsOutputReference) PutEnvironmentVariables(value interface{}) {
	if err := c.validatePutEnvironmentVariablesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putEnvironmentVariables",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodepipelinePipelineStagesActionsOutputReference) PutInputArtifacts(value interface{}) {
	if err := c.validatePutInputArtifactsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putInputArtifacts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodepipelinePipelineStagesActionsOutputReference) PutOutputArtifacts(value interface{}) {
	if err := c.validatePutOutputArtifactsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putOutputArtifacts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodepipelinePipelineStagesActionsOutputReference) ResetCommands() {
	_jsii_.InvokeVoid(
		c,
		"resetCommands",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodepipelinePipelineStagesActionsOutputReference) ResetConfiguration() {
	_jsii_.InvokeVoid(
		c,
		"resetConfiguration",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodepipelinePipelineStagesActionsOutputReference) ResetEnvironmentVariables() {
	_jsii_.InvokeVoid(
		c,
		"resetEnvironmentVariables",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodepipelinePipelineStagesActionsOutputReference) ResetInputArtifacts() {
	_jsii_.InvokeVoid(
		c,
		"resetInputArtifacts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodepipelinePipelineStagesActionsOutputReference) ResetNamespace() {
	_jsii_.InvokeVoid(
		c,
		"resetNamespace",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodepipelinePipelineStagesActionsOutputReference) ResetOutputArtifacts() {
	_jsii_.InvokeVoid(
		c,
		"resetOutputArtifacts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodepipelinePipelineStagesActionsOutputReference) ResetOutputVariables() {
	_jsii_.InvokeVoid(
		c,
		"resetOutputVariables",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodepipelinePipelineStagesActionsOutputReference) ResetRegion() {
	_jsii_.InvokeVoid(
		c,
		"resetRegion",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodepipelinePipelineStagesActionsOutputReference) ResetRoleArn() {
	_jsii_.InvokeVoid(
		c,
		"resetRoleArn",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodepipelinePipelineStagesActionsOutputReference) ResetRunOrder() {
	_jsii_.InvokeVoid(
		c,
		"resetRunOrder",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodepipelinePipelineStagesActionsOutputReference) ResetTimeoutInMinutes() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeoutInMinutes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodepipelinePipelineStagesActionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_CodepipelinePipelineStagesActionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

