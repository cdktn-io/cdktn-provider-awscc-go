// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/ecsservice/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EcsServiceDeploymentConfigurationLifecycleHooksOutputReference interface {
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
	HookDetails() *string
	SetHookDetails(val *string)
	HookDetailsInput() *string
	HookTargetArn() *string
	SetHookTargetArn(val *string)
	HookTargetArnInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LifecycleStages() *[]*string
	SetLifecycleStages(val *[]*string)
	LifecycleStagesInput() *[]*string
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
	TargetType() *string
	SetTargetType(val *string)
	TargetTypeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TimeoutConfiguration() EcsServiceDeploymentConfigurationLifecycleHooksTimeoutConfigurationOutputReference
	TimeoutConfigurationInput() interface{}
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
	PutTimeoutConfiguration(value *EcsServiceDeploymentConfigurationLifecycleHooksTimeoutConfiguration)
	ResetHookDetails()
	ResetHookTargetArn()
	ResetLifecycleStages()
	ResetRoleArn()
	ResetTargetType()
	ResetTimeoutConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EcsServiceDeploymentConfigurationLifecycleHooksOutputReference
type jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHooksOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHooksOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHooksOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHooksOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHooksOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHooksOutputReference) HookDetails() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hookDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHooksOutputReference) HookDetailsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hookDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHooksOutputReference) HookTargetArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hookTargetArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHooksOutputReference) HookTargetArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hookTargetArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHooksOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHooksOutputReference) LifecycleStages() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"lifecycleStages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHooksOutputReference) LifecycleStagesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"lifecycleStagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHooksOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHooksOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHooksOutputReference) TargetType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHooksOutputReference) TargetTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHooksOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHooksOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHooksOutputReference) TimeoutConfiguration() EcsServiceDeploymentConfigurationLifecycleHooksTimeoutConfigurationOutputReference {
	var returns EcsServiceDeploymentConfigurationLifecycleHooksTimeoutConfigurationOutputReference
	_jsii_.Get(
		j,
		"timeoutConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHooksOutputReference) TimeoutConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutConfigurationInput",
		&returns,
	)
	return returns
}


func NewEcsServiceDeploymentConfigurationLifecycleHooksOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) EcsServiceDeploymentConfigurationLifecycleHooksOutputReference {
	_init_.Initialize()

	if err := validateNewEcsServiceDeploymentConfigurationLifecycleHooksOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHooksOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.ecsService.EcsServiceDeploymentConfigurationLifecycleHooksOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEcsServiceDeploymentConfigurationLifecycleHooksOutputReference_Override(e EcsServiceDeploymentConfigurationLifecycleHooksOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.ecsService.EcsServiceDeploymentConfigurationLifecycleHooksOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHooksOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHooksOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHooksOutputReference)SetHookDetails(val *string) {
	if err := j.validateSetHookDetailsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hookDetails",
		val,
	)
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHooksOutputReference)SetHookTargetArn(val *string) {
	if err := j.validateSetHookTargetArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hookTargetArn",
		val,
	)
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHooksOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHooksOutputReference)SetLifecycleStages(val *[]*string) {
	if err := j.validateSetLifecycleStagesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycleStages",
		val,
	)
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHooksOutputReference)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHooksOutputReference)SetTargetType(val *string) {
	if err := j.validateSetTargetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetType",
		val,
	)
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHooksOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHooksOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHooksOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHooksOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHooksOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHooksOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHooksOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHooksOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHooksOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHooksOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHooksOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHooksOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHooksOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHooksOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHooksOutputReference) PutTimeoutConfiguration(value *EcsServiceDeploymentConfigurationLifecycleHooksTimeoutConfiguration) {
	if err := e.validatePutTimeoutConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putTimeoutConfiguration",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHooksOutputReference) ResetHookDetails() {
	_jsii_.InvokeVoid(
		e,
		"resetHookDetails",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHooksOutputReference) ResetHookTargetArn() {
	_jsii_.InvokeVoid(
		e,
		"resetHookTargetArn",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHooksOutputReference) ResetLifecycleStages() {
	_jsii_.InvokeVoid(
		e,
		"resetLifecycleStages",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHooksOutputReference) ResetRoleArn() {
	_jsii_.InvokeVoid(
		e,
		"resetRoleArn",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHooksOutputReference) ResetTargetType() {
	_jsii_.InvokeVoid(
		e,
		"resetTargetType",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHooksOutputReference) ResetTimeoutConfiguration() {
	_jsii_.InvokeVoid(
		e,
		"resetTimeoutConfiguration",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHooksOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := e.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHooksOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

