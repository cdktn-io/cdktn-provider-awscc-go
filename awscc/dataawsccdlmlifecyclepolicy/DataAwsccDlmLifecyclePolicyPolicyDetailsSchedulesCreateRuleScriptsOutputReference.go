// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccdlmlifecyclepolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawsccdlmlifecyclepolicy/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference interface {
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
	ExecuteOperationOnScriptFailure() cdktn.IResolvable
	ExecutionHandler() *string
	ExecutionHandlerService() *string
	ExecutionTimeout() *float64
	// Experimental.
	Fqn() *string
	InternalValue() *DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScripts
	SetInternalValue(val *DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScripts)
	MaximumRetryCount() *float64
	Stages() *[]*string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference
type jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference) ExecuteOperationOnScriptFailure() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"executeOperationOnScriptFailure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference) ExecutionHandler() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionHandler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference) ExecutionHandlerService() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionHandlerService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference) ExecutionTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"executionTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference) InternalValue() *DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScripts {
	var returns *DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScripts
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference) MaximumRetryCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumRetryCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference) Stages() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"stages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccDlmLifecyclePolicy.DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference_Override(d DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccDlmLifecyclePolicy.DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference)SetInternalValue(val *DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScripts) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScriptsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

