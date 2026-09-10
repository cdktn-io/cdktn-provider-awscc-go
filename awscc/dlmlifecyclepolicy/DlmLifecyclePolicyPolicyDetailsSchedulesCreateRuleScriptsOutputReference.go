// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dlmlifecyclepolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dlmlifecyclepolicy/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference interface {
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
	ExecuteOperationOnScriptFailure() interface{}
	SetExecuteOperationOnScriptFailure(val interface{})
	ExecuteOperationOnScriptFailureInput() interface{}
	ExecutionHandler() *string
	SetExecutionHandler(val *string)
	ExecutionHandlerInput() *string
	ExecutionHandlerService() *string
	SetExecutionHandlerService(val *string)
	ExecutionHandlerServiceInput() *string
	ExecutionTimeout() *float64
	SetExecutionTimeout(val *float64)
	ExecutionTimeoutInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MaximumRetryCount() *float64
	SetMaximumRetryCount(val *float64)
	MaximumRetryCountInput() *float64
	Stages() *[]*string
	SetStages(val *[]*string)
	StagesInput() *[]*string
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
	ResetExecuteOperationOnScriptFailure()
	ResetExecutionHandler()
	ResetExecutionHandlerService()
	ResetExecutionTimeout()
	ResetMaximumRetryCount()
	ResetStages()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference
type jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference) ExecuteOperationOnScriptFailure() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"executeOperationOnScriptFailure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference) ExecuteOperationOnScriptFailureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"executeOperationOnScriptFailureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference) ExecutionHandler() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionHandler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference) ExecutionHandlerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionHandlerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference) ExecutionHandlerService() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionHandlerService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference) ExecutionHandlerServiceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionHandlerServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference) ExecutionTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"executionTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference) ExecutionTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"executionTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference) MaximumRetryCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumRetryCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference) MaximumRetryCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumRetryCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference) Stages() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"stages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference) StagesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"stagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference {
	_init_.Initialize()

	if err := validateNewDlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dlmLifecyclePolicy.DlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference_Override(d DlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dlmLifecyclePolicy.DlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference)SetExecuteOperationOnScriptFailure(val interface{}) {
	if err := j.validateSetExecuteOperationOnScriptFailureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executeOperationOnScriptFailure",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference)SetExecutionHandler(val *string) {
	if err := j.validateSetExecutionHandlerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionHandler",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference)SetExecutionHandlerService(val *string) {
	if err := j.validateSetExecutionHandlerServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionHandlerService",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference)SetExecutionTimeout(val *float64) {
	if err := j.validateSetExecutionTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionTimeout",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference)SetMaximumRetryCount(val *float64) {
	if err := j.validateSetMaximumRetryCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumRetryCount",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference)SetStages(val *[]*string) {
	if err := j.validateSetStagesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stages",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference) ResetExecuteOperationOnScriptFailure() {
	_jsii_.InvokeVoid(
		d,
		"resetExecuteOperationOnScriptFailure",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference) ResetExecutionHandler() {
	_jsii_.InvokeVoid(
		d,
		"resetExecutionHandler",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference) ResetExecutionHandlerService() {
	_jsii_.InvokeVoid(
		d,
		"resetExecutionHandlerService",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference) ResetExecutionTimeout() {
	_jsii_.InvokeVoid(
		d,
		"resetExecutionTimeout",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference) ResetMaximumRetryCount() {
	_jsii_.InvokeVoid(
		d,
		"resetMaximumRetryCount",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference) ResetStages() {
	_jsii_.InvokeVoid(
		d,
		"resetStages",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

