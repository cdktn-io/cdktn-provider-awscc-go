// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pcscomputenodegroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/pcscomputenodegroup/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PcsComputeNodeGroupNodeLifecycleActionsStagesNodeBootstrappedOutputReference interface {
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
	ScriptSource() PcsComputeNodeGroupNodeLifecycleActionsStagesNodeBootstrappedScriptSourceOutputReference
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
	PutScriptSource(value *PcsComputeNodeGroupNodeLifecycleActionsStagesNodeBootstrappedScriptSource)
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

// The jsii proxy struct for PcsComputeNodeGroupNodeLifecycleActionsStagesNodeBootstrappedOutputReference
type jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeBootstrappedOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeBootstrappedOutputReference) Arguments() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"arguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeBootstrappedOutputReference) ArgumentsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"argumentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeBootstrappedOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeBootstrappedOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeBootstrappedOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeBootstrappedOutputReference) ExecutionPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeBootstrappedOutputReference) ExecutionPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeBootstrappedOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeBootstrappedOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeBootstrappedOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeBootstrappedOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeBootstrappedOutputReference) OnError() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onError",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeBootstrappedOutputReference) OnErrorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onErrorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeBootstrappedOutputReference) ScriptSource() PcsComputeNodeGroupNodeLifecycleActionsStagesNodeBootstrappedScriptSourceOutputReference {
	var returns PcsComputeNodeGroupNodeLifecycleActionsStagesNodeBootstrappedScriptSourceOutputReference
	_jsii_.Get(
		j,
		"scriptSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeBootstrappedOutputReference) ScriptSourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scriptSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeBootstrappedOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeBootstrappedOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPcsComputeNodeGroupNodeLifecycleActionsStagesNodeBootstrappedOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) PcsComputeNodeGroupNodeLifecycleActionsStagesNodeBootstrappedOutputReference {
	_init_.Initialize()

	if err := validateNewPcsComputeNodeGroupNodeLifecycleActionsStagesNodeBootstrappedOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeBootstrappedOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.pcsComputeNodeGroup.PcsComputeNodeGroupNodeLifecycleActionsStagesNodeBootstrappedOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewPcsComputeNodeGroupNodeLifecycleActionsStagesNodeBootstrappedOutputReference_Override(p PcsComputeNodeGroupNodeLifecycleActionsStagesNodeBootstrappedOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.pcsComputeNodeGroup.PcsComputeNodeGroupNodeLifecycleActionsStagesNodeBootstrappedOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		p,
	)
}

func (j *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeBootstrappedOutputReference)SetArguments(val *[]*string) {
	if err := j.validateSetArgumentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"arguments",
		val,
	)
}

func (j *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeBootstrappedOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeBootstrappedOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeBootstrappedOutputReference)SetExecutionPolicy(val *string) {
	if err := j.validateSetExecutionPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionPolicy",
		val,
	)
}

func (j *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeBootstrappedOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeBootstrappedOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeBootstrappedOutputReference)SetOnError(val *string) {
	if err := j.validateSetOnErrorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onError",
		val,
	)
}

func (j *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeBootstrappedOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeBootstrappedOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeBootstrappedOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeBootstrappedOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeBootstrappedOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeBootstrappedOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeBootstrappedOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeBootstrappedOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeBootstrappedOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeBootstrappedOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeBootstrappedOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeBootstrappedOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeBootstrappedOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeBootstrappedOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeBootstrappedOutputReference) PutScriptSource(value *PcsComputeNodeGroupNodeLifecycleActionsStagesNodeBootstrappedScriptSource) {
	if err := p.validatePutScriptSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putScriptSource",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeBootstrappedOutputReference) ResetArguments() {
	_jsii_.InvokeVoid(
		p,
		"resetArguments",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeBootstrappedOutputReference) ResetExecutionPolicy() {
	_jsii_.InvokeVoid(
		p,
		"resetExecutionPolicy",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeBootstrappedOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		p,
		"resetName",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeBootstrappedOutputReference) ResetOnError() {
	_jsii_.InvokeVoid(
		p,
		"resetOnError",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeBootstrappedOutputReference) ResetScriptSource() {
	_jsii_.InvokeVoid(
		p,
		"resetScriptSource",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeBootstrappedOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (p *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeBootstrappedOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

