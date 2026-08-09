// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wisdomaiagent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/wisdomaiagent/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputFiltersOutputConfigurationOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	OutputVariableNameOverride() *string
	SetOutputVariableNameOverride(val *string)
	OutputVariableNameOverrideInput() *string
	SessionDataNamespace() *string
	SetSessionDataNamespace(val *string)
	SessionDataNamespaceInput() *string
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
	ResetOutputVariableNameOverride()
	ResetSessionDataNamespace()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputFiltersOutputConfigurationOutputReference
type jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputFiltersOutputConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputFiltersOutputConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputFiltersOutputConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputFiltersOutputConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputFiltersOutputConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputFiltersOutputConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputFiltersOutputConfigurationOutputReference) OutputVariableNameOverride() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputVariableNameOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputFiltersOutputConfigurationOutputReference) OutputVariableNameOverrideInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputVariableNameOverrideInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputFiltersOutputConfigurationOutputReference) SessionDataNamespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionDataNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputFiltersOutputConfigurationOutputReference) SessionDataNamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionDataNamespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputFiltersOutputConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputFiltersOutputConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputFiltersOutputConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputFiltersOutputConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewWisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputFiltersOutputConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputFiltersOutputConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.wisdomAiAgent.WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputFiltersOutputConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputFiltersOutputConfigurationOutputReference_Override(w WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputFiltersOutputConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.wisdomAiAgent.WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputFiltersOutputConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputFiltersOutputConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputFiltersOutputConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputFiltersOutputConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputFiltersOutputConfigurationOutputReference)SetOutputVariableNameOverride(val *string) {
	if err := j.validateSetOutputVariableNameOverrideParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputVariableNameOverride",
		val,
	)
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputFiltersOutputConfigurationOutputReference)SetSessionDataNamespace(val *string) {
	if err := j.validateSetSessionDataNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionDataNamespace",
		val,
	)
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputFiltersOutputConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputFiltersOutputConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputFiltersOutputConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputFiltersOutputConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (w *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputFiltersOutputConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (w *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputFiltersOutputConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (w *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputFiltersOutputConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (w *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputFiltersOutputConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (w *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputFiltersOutputConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (w *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputFiltersOutputConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (w *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputFiltersOutputConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (w *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputFiltersOutputConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (w *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputFiltersOutputConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputFiltersOutputConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (w *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputFiltersOutputConfigurationOutputReference) ResetOutputVariableNameOverride() {
	_jsii_.InvokeVoid(
		w,
		"resetOutputVariableNameOverride",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputFiltersOutputConfigurationOutputReference) ResetSessionDataNamespace() {
	_jsii_.InvokeVoid(
		w,
		"resetSessionDataNamespace",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputFiltersOutputConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (w *jsiiProxy_WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputFiltersOutputConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

