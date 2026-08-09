// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lexbot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/lexbot/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LexBotBotLocalesIntentsOutputReference interface {
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DialogCodeHook() LexBotBotLocalesIntentsDialogCodeHookOutputReference
	DialogCodeHookInput() interface{}
	// Experimental.
	Fqn() *string
	FulfillmentCodeHook() LexBotBotLocalesIntentsFulfillmentCodeHookOutputReference
	FulfillmentCodeHookInput() interface{}
	InputContexts() LexBotBotLocalesIntentsInputContextsList
	InputContextsInput() interface{}
	IntentClosingSetting() LexBotBotLocalesIntentsIntentClosingSettingOutputReference
	IntentClosingSettingInput() interface{}
	IntentConfirmationSetting() LexBotBotLocalesIntentsIntentConfirmationSettingOutputReference
	IntentConfirmationSettingInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	KendraConfiguration() LexBotBotLocalesIntentsKendraConfigurationOutputReference
	KendraConfigurationInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	OutputContexts() LexBotBotLocalesIntentsOutputContextsList
	OutputContextsInput() interface{}
	ParentIntentSignature() *string
	SetParentIntentSignature(val *string)
	ParentIntentSignatureInput() *string
	SampleUtterances() LexBotBotLocalesIntentsSampleUtterancesList
	SampleUtterancesInput() interface{}
	SlotPriorities() LexBotBotLocalesIntentsSlotPrioritiesList
	SlotPrioritiesInput() interface{}
	Slots() LexBotBotLocalesIntentsSlotsList
	SlotsInput() interface{}
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
	PutDialogCodeHook(value *LexBotBotLocalesIntentsDialogCodeHook)
	PutFulfillmentCodeHook(value *LexBotBotLocalesIntentsFulfillmentCodeHook)
	PutInputContexts(value interface{})
	PutIntentClosingSetting(value *LexBotBotLocalesIntentsIntentClosingSetting)
	PutIntentConfirmationSetting(value *LexBotBotLocalesIntentsIntentConfirmationSetting)
	PutKendraConfiguration(value *LexBotBotLocalesIntentsKendraConfiguration)
	PutOutputContexts(value interface{})
	PutSampleUtterances(value interface{})
	PutSlotPriorities(value interface{})
	PutSlots(value interface{})
	ResetDescription()
	ResetDialogCodeHook()
	ResetFulfillmentCodeHook()
	ResetInputContexts()
	ResetIntentClosingSetting()
	ResetIntentConfirmationSetting()
	ResetKendraConfiguration()
	ResetName()
	ResetOutputContexts()
	ResetParentIntentSignature()
	ResetSampleUtterances()
	ResetSlotPriorities()
	ResetSlots()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LexBotBotLocalesIntentsOutputReference
type jsiiProxy_LexBotBotLocalesIntentsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_LexBotBotLocalesIntentsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsOutputReference) DialogCodeHook() LexBotBotLocalesIntentsDialogCodeHookOutputReference {
	var returns LexBotBotLocalesIntentsDialogCodeHookOutputReference
	_jsii_.Get(
		j,
		"dialogCodeHook",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsOutputReference) DialogCodeHookInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dialogCodeHookInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsOutputReference) FulfillmentCodeHook() LexBotBotLocalesIntentsFulfillmentCodeHookOutputReference {
	var returns LexBotBotLocalesIntentsFulfillmentCodeHookOutputReference
	_jsii_.Get(
		j,
		"fulfillmentCodeHook",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsOutputReference) FulfillmentCodeHookInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fulfillmentCodeHookInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsOutputReference) InputContexts() LexBotBotLocalesIntentsInputContextsList {
	var returns LexBotBotLocalesIntentsInputContextsList
	_jsii_.Get(
		j,
		"inputContexts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsOutputReference) InputContextsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inputContextsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsOutputReference) IntentClosingSetting() LexBotBotLocalesIntentsIntentClosingSettingOutputReference {
	var returns LexBotBotLocalesIntentsIntentClosingSettingOutputReference
	_jsii_.Get(
		j,
		"intentClosingSetting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsOutputReference) IntentClosingSettingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"intentClosingSettingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsOutputReference) IntentConfirmationSetting() LexBotBotLocalesIntentsIntentConfirmationSettingOutputReference {
	var returns LexBotBotLocalesIntentsIntentConfirmationSettingOutputReference
	_jsii_.Get(
		j,
		"intentConfirmationSetting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsOutputReference) IntentConfirmationSettingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"intentConfirmationSettingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsOutputReference) KendraConfiguration() LexBotBotLocalesIntentsKendraConfigurationOutputReference {
	var returns LexBotBotLocalesIntentsKendraConfigurationOutputReference
	_jsii_.Get(
		j,
		"kendraConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsOutputReference) KendraConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kendraConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsOutputReference) OutputContexts() LexBotBotLocalesIntentsOutputContextsList {
	var returns LexBotBotLocalesIntentsOutputContextsList
	_jsii_.Get(
		j,
		"outputContexts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsOutputReference) OutputContextsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"outputContextsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsOutputReference) ParentIntentSignature() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentIntentSignature",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsOutputReference) ParentIntentSignatureInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentIntentSignatureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsOutputReference) SampleUtterances() LexBotBotLocalesIntentsSampleUtterancesList {
	var returns LexBotBotLocalesIntentsSampleUtterancesList
	_jsii_.Get(
		j,
		"sampleUtterances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsOutputReference) SampleUtterancesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sampleUtterancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsOutputReference) SlotPriorities() LexBotBotLocalesIntentsSlotPrioritiesList {
	var returns LexBotBotLocalesIntentsSlotPrioritiesList
	_jsii_.Get(
		j,
		"slotPriorities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsOutputReference) SlotPrioritiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"slotPrioritiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsOutputReference) Slots() LexBotBotLocalesIntentsSlotsList {
	var returns LexBotBotLocalesIntentsSlotsList
	_jsii_.Get(
		j,
		"slots",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsOutputReference) SlotsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"slotsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewLexBotBotLocalesIntentsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) LexBotBotLocalesIntentsOutputReference {
	_init_.Initialize()

	if err := validateNewLexBotBotLocalesIntentsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_LexBotBotLocalesIntentsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.lexBot.LexBotBotLocalesIntentsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewLexBotBotLocalesIntentsOutputReference_Override(l LexBotBotLocalesIntentsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.lexBot.LexBotBotLocalesIntentsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		l,
	)
}

func (j *jsiiProxy_LexBotBotLocalesIntentsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LexBotBotLocalesIntentsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LexBotBotLocalesIntentsOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_LexBotBotLocalesIntentsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LexBotBotLocalesIntentsOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_LexBotBotLocalesIntentsOutputReference)SetParentIntentSignature(val *string) {
	if err := j.validateSetParentIntentSignatureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parentIntentSignature",
		val,
	)
}

func (j *jsiiProxy_LexBotBotLocalesIntentsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LexBotBotLocalesIntentsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LexBotBotLocalesIntentsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexBotBotLocalesIntentsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexBotBotLocalesIntentsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexBotBotLocalesIntentsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexBotBotLocalesIntentsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexBotBotLocalesIntentsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexBotBotLocalesIntentsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexBotBotLocalesIntentsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexBotBotLocalesIntentsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexBotBotLocalesIntentsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexBotBotLocalesIntentsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexBotBotLocalesIntentsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexBotBotLocalesIntentsOutputReference) PutDialogCodeHook(value *LexBotBotLocalesIntentsDialogCodeHook) {
	if err := l.validatePutDialogCodeHookParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putDialogCodeHook",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LexBotBotLocalesIntentsOutputReference) PutFulfillmentCodeHook(value *LexBotBotLocalesIntentsFulfillmentCodeHook) {
	if err := l.validatePutFulfillmentCodeHookParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putFulfillmentCodeHook",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LexBotBotLocalesIntentsOutputReference) PutInputContexts(value interface{}) {
	if err := l.validatePutInputContextsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putInputContexts",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LexBotBotLocalesIntentsOutputReference) PutIntentClosingSetting(value *LexBotBotLocalesIntentsIntentClosingSetting) {
	if err := l.validatePutIntentClosingSettingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putIntentClosingSetting",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LexBotBotLocalesIntentsOutputReference) PutIntentConfirmationSetting(value *LexBotBotLocalesIntentsIntentConfirmationSetting) {
	if err := l.validatePutIntentConfirmationSettingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putIntentConfirmationSetting",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LexBotBotLocalesIntentsOutputReference) PutKendraConfiguration(value *LexBotBotLocalesIntentsKendraConfiguration) {
	if err := l.validatePutKendraConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putKendraConfiguration",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LexBotBotLocalesIntentsOutputReference) PutOutputContexts(value interface{}) {
	if err := l.validatePutOutputContextsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putOutputContexts",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LexBotBotLocalesIntentsOutputReference) PutSampleUtterances(value interface{}) {
	if err := l.validatePutSampleUtterancesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putSampleUtterances",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LexBotBotLocalesIntentsOutputReference) PutSlotPriorities(value interface{}) {
	if err := l.validatePutSlotPrioritiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putSlotPriorities",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LexBotBotLocalesIntentsOutputReference) PutSlots(value interface{}) {
	if err := l.validatePutSlotsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putSlots",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LexBotBotLocalesIntentsOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		l,
		"resetDescription",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexBotBotLocalesIntentsOutputReference) ResetDialogCodeHook() {
	_jsii_.InvokeVoid(
		l,
		"resetDialogCodeHook",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexBotBotLocalesIntentsOutputReference) ResetFulfillmentCodeHook() {
	_jsii_.InvokeVoid(
		l,
		"resetFulfillmentCodeHook",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexBotBotLocalesIntentsOutputReference) ResetInputContexts() {
	_jsii_.InvokeVoid(
		l,
		"resetInputContexts",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexBotBotLocalesIntentsOutputReference) ResetIntentClosingSetting() {
	_jsii_.InvokeVoid(
		l,
		"resetIntentClosingSetting",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexBotBotLocalesIntentsOutputReference) ResetIntentConfirmationSetting() {
	_jsii_.InvokeVoid(
		l,
		"resetIntentConfirmationSetting",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexBotBotLocalesIntentsOutputReference) ResetKendraConfiguration() {
	_jsii_.InvokeVoid(
		l,
		"resetKendraConfiguration",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexBotBotLocalesIntentsOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		l,
		"resetName",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexBotBotLocalesIntentsOutputReference) ResetOutputContexts() {
	_jsii_.InvokeVoid(
		l,
		"resetOutputContexts",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexBotBotLocalesIntentsOutputReference) ResetParentIntentSignature() {
	_jsii_.InvokeVoid(
		l,
		"resetParentIntentSignature",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexBotBotLocalesIntentsOutputReference) ResetSampleUtterances() {
	_jsii_.InvokeVoid(
		l,
		"resetSampleUtterances",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexBotBotLocalesIntentsOutputReference) ResetSlotPriorities() {
	_jsii_.InvokeVoid(
		l,
		"resetSlotPriorities",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexBotBotLocalesIntentsOutputReference) ResetSlots() {
	_jsii_.InvokeVoid(
		l,
		"resetSlots",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexBotBotLocalesIntentsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := l.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexBotBotLocalesIntentsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

