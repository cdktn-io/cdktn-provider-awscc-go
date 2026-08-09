// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wisdomaiagent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/wisdomaiagent/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference interface {
	cdktn.ComplexObject
	Annotations() *string
	SetAnnotations(val *string)
	AnnotationsInput() *string
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	// Experimental.
	Fqn() *string
	InputSchema() *string
	SetInputSchema(val *string)
	InputSchemaInput() *string
	Instruction() WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsInstructionOutputReference
	InstructionInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	OutputFilters() WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputFiltersList
	OutputFiltersInput() interface{}
	OutputSchema() *string
	SetOutputSchema(val *string)
	OutputSchemaInput() *string
	OverrideInputValues() WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOverrideInputValuesList
	OverrideInputValuesInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Title() *string
	SetTitle(val *string)
	TitleInput() *string
	ToolId() *string
	SetToolId(val *string)
	ToolIdInput() *string
	ToolName() *string
	SetToolName(val *string)
	ToolNameInput() *string
	ToolType() *string
	SetToolType(val *string)
	ToolTypeInput() *string
	UserInteractionConfiguration() WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsUserInteractionConfigurationOutputReference
	UserInteractionConfigurationInput() interface{}
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
	PutInstruction(value *WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsInstruction)
	PutOutputFilters(value interface{})
	PutOverrideInputValues(value interface{})
	PutUserInteractionConfiguration(value *WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsUserInteractionConfiguration)
	ResetAnnotations()
	ResetDescription()
	ResetInputSchema()
	ResetInstruction()
	ResetOutputFilters()
	ResetOutputSchema()
	ResetOverrideInputValues()
	ResetTitle()
	ResetToolId()
	ResetToolName()
	ResetToolType()
	ResetUserInteractionConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference
type jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference) Annotations() *string {
	var returns *string
	_jsii_.Get(
		j,
		"annotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference) AnnotationsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"annotationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference) InputSchema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputSchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference) InputSchemaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputSchemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference) Instruction() WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsInstructionOutputReference {
	var returns WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsInstructionOutputReference
	_jsii_.Get(
		j,
		"instruction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference) InstructionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"instructionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference) OutputFilters() WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputFiltersList {
	var returns WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputFiltersList
	_jsii_.Get(
		j,
		"outputFilters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference) OutputFiltersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"outputFiltersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference) OutputSchema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputSchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference) OutputSchemaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputSchemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference) OverrideInputValues() WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOverrideInputValuesList {
	var returns WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOverrideInputValuesList
	_jsii_.Get(
		j,
		"overrideInputValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference) OverrideInputValuesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"overrideInputValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference) TitleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference) ToolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"toolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference) ToolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"toolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference) ToolName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"toolName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference) ToolNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"toolNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference) ToolType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"toolType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference) ToolTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"toolTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference) UserInteractionConfiguration() WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsUserInteractionConfigurationOutputReference {
	var returns WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsUserInteractionConfigurationOutputReference
	_jsii_.Get(
		j,
		"userInteractionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference) UserInteractionConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userInteractionConfigurationInput",
		&returns,
	)
	return returns
}


func NewWisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference {
	_init_.Initialize()

	if err := validateNewWisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.wisdomAiAgent.WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewWisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference_Override(w WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.wisdomAiAgent.WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		w,
	)
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference)SetAnnotations(val *string) {
	if err := j.validateSetAnnotationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"annotations",
		val,
	)
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference)SetInputSchema(val *string) {
	if err := j.validateSetInputSchemaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputSchema",
		val,
	)
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference)SetOutputSchema(val *string) {
	if err := j.validateSetOutputSchemaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputSchema",
		val,
	)
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference)SetTitle(val *string) {
	if err := j.validateSetTitleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"title",
		val,
	)
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference)SetToolId(val *string) {
	if err := j.validateSetToolIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"toolId",
		val,
	)
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference)SetToolName(val *string) {
	if err := j.validateSetToolNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"toolName",
		val,
	)
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference)SetToolType(val *string) {
	if err := j.validateSetToolTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"toolType",
		val,
	)
}

func (w *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := w.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := w.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := w.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := w.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := w.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := w.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := w.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := w.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := w.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference) PutInstruction(value *WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsInstruction) {
	if err := w.validatePutInstructionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putInstruction",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference) PutOutputFilters(value interface{}) {
	if err := w.validatePutOutputFiltersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putOutputFilters",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference) PutOverrideInputValues(value interface{}) {
	if err := w.validatePutOverrideInputValuesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putOverrideInputValues",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference) PutUserInteractionConfiguration(value *WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsUserInteractionConfiguration) {
	if err := w.validatePutUserInteractionConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putUserInteractionConfiguration",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference) ResetAnnotations() {
	_jsii_.InvokeVoid(
		w,
		"resetAnnotations",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		w,
		"resetDescription",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference) ResetInputSchema() {
	_jsii_.InvokeVoid(
		w,
		"resetInputSchema",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference) ResetInstruction() {
	_jsii_.InvokeVoid(
		w,
		"resetInstruction",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference) ResetOutputFilters() {
	_jsii_.InvokeVoid(
		w,
		"resetOutputFilters",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference) ResetOutputSchema() {
	_jsii_.InvokeVoid(
		w,
		"resetOutputSchema",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference) ResetOverrideInputValues() {
	_jsii_.InvokeVoid(
		w,
		"resetOverrideInputValues",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference) ResetTitle() {
	_jsii_.InvokeVoid(
		w,
		"resetTitle",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference) ResetToolId() {
	_jsii_.InvokeVoid(
		w,
		"resetToolId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference) ResetToolName() {
	_jsii_.InvokeVoid(
		w,
		"resetToolName",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference) ResetToolType() {
	_jsii_.InvokeVoid(
		w,
		"resetToolType",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference) ResetUserInteractionConfiguration() {
	_jsii_.InvokeVoid(
		w,
		"resetUserInteractionConfiguration",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := w.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

