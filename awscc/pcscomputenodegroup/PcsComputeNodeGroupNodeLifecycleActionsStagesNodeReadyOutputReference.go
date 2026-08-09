// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pcscomputenodegroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/pcscomputenodegroup/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyOutputReference interface {
	cdktn.ComplexObject
	Arguments() *[]*string
	SetArguments(val *[]*string)
	ArgumentsInput() *[]*string
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
	ExecutionPolicy() *string
	SetExecutionPolicy(val *string)
	ExecutionPolicyInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	OnError() *string
	SetOnError(val *string)
	OnErrorInput() *string
	ScriptSource() PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyScriptSourceOutputReference
	ScriptSourceInput() interface{}
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
	PutScriptSource(value *PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyScriptSource)
	ResetArguments()
	ResetExecutionPolicy()
	ResetName()
	ResetOnError()
	ResetScriptSource()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyOutputReference
type jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyOutputReference) Arguments() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"arguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyOutputReference) ArgumentsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"argumentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyOutputReference) ExecutionPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyOutputReference) ExecutionPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyOutputReference) OnError() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onError",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyOutputReference) OnErrorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onErrorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyOutputReference) ScriptSource() PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyScriptSourceOutputReference {
	var returns PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyScriptSourceOutputReference
	_jsii_.Get(
		j,
		"scriptSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyOutputReference) ScriptSourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scriptSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyOutputReference {
	_init_.Initialize()

	if err := validateNewPcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.pcsComputeNodeGroup.PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewPcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyOutputReference_Override(p PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.pcsComputeNodeGroup.PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		p,
	)
}

func (j *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyOutputReference)SetArguments(val *[]*string) {
	if err := j.validateSetArgumentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"arguments",
		val,
	)
}

func (j *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyOutputReference)SetExecutionPolicy(val *string) {
	if err := j.validateSetExecutionPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionPolicy",
		val,
	)
}

func (j *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyOutputReference)SetOnError(val *string) {
	if err := j.validateSetOnErrorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onError",
		val,
	)
}

func (j *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyOutputReference) PutScriptSource(value *PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyScriptSource) {
	if err := p.validatePutScriptSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putScriptSource",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyOutputReference) ResetArguments() {
	_jsii_.InvokeVoid(
		p,
		"resetArguments",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyOutputReference) ResetExecutionPolicy() {
	_jsii_.InvokeVoid(
		p,
		"resetExecutionPolicy",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		p,
		"resetName",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyOutputReference) ResetOnError() {
	_jsii_.InvokeVoid(
		p,
		"resetOnError",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyOutputReference) ResetScriptSource() {
	_jsii_.InvokeVoid(
		p,
		"resetScriptSource",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := p.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

