// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wisdomaiagent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/wisdomaiagent/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationOutputReference interface {
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
	ConnectInstanceArn() *string
	SetConnectInstanceArn(val *string)
	ConnectInstanceArnInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Locale() *string
	SetLocale(val *string)
	LocaleInput() *string
	OrchestrationAiGuardrailId() *string
	SetOrchestrationAiGuardrailId(val *string)
	OrchestrationAiGuardrailIdInput() *string
	OrchestrationAiPromptId() *string
	SetOrchestrationAiPromptId(val *string)
	OrchestrationAiPromptIdInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	ToolConfigurations() WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsList
	ToolConfigurationsInput() interface{}
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
	PutToolConfigurations(value interface{})
	ResetConnectInstanceArn()
	ResetLocale()
	ResetOrchestrationAiGuardrailId()
	ResetOrchestrationAiPromptId()
	ResetToolConfigurations()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationOutputReference
type jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationOutputReference) ConnectInstanceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectInstanceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationOutputReference) ConnectInstanceArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectInstanceArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationOutputReference) Locale() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locale",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationOutputReference) LocaleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationOutputReference) OrchestrationAiGuardrailId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orchestrationAiGuardrailId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationOutputReference) OrchestrationAiGuardrailIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orchestrationAiGuardrailIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationOutputReference) OrchestrationAiPromptId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orchestrationAiPromptId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationOutputReference) OrchestrationAiPromptIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orchestrationAiPromptIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationOutputReference) ToolConfigurations() WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsList {
	var returns WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsList
	_jsii_.Get(
		j,
		"toolConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationOutputReference) ToolConfigurationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"toolConfigurationsInput",
		&returns,
	)
	return returns
}


func NewWisdomAiAgentConfigurationOrchestrationAiAgentConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewWisdomAiAgentConfigurationOrchestrationAiAgentConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.wisdomAiAgent.WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWisdomAiAgentConfigurationOrchestrationAiAgentConfigurationOutputReference_Override(w WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.wisdomAiAgent.WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationOutputReference)SetConnectInstanceArn(val *string) {
	if err := j.validateSetConnectInstanceArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectInstanceArn",
		val,
	)
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationOutputReference)SetLocale(val *string) {
	if err := j.validateSetLocaleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"locale",
		val,
	)
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationOutputReference)SetOrchestrationAiGuardrailId(val *string) {
	if err := j.validateSetOrchestrationAiGuardrailIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"orchestrationAiGuardrailId",
		val,
	)
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationOutputReference)SetOrchestrationAiPromptId(val *string) {
	if err := j.validateSetOrchestrationAiPromptIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"orchestrationAiPromptId",
		val,
	)
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (w *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (w *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (w *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (w *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (w *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (w *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (w *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (w *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (w *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (w *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationOutputReference) PutToolConfigurations(value interface{}) {
	if err := w.validatePutToolConfigurationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putToolConfigurations",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationOutputReference) ResetConnectInstanceArn() {
	_jsii_.InvokeVoid(
		w,
		"resetConnectInstanceArn",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationOutputReference) ResetLocale() {
	_jsii_.InvokeVoid(
		w,
		"resetLocale",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationOutputReference) ResetOrchestrationAiGuardrailId() {
	_jsii_.InvokeVoid(
		w,
		"resetOrchestrationAiGuardrailId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationOutputReference) ResetOrchestrationAiPromptId() {
	_jsii_.InvokeVoid(
		w,
		"resetOrchestrationAiPromptId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationOutputReference) ResetToolConfigurations() {
	_jsii_.InvokeVoid(
		w,
		"resetToolConfigurations",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (w *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

