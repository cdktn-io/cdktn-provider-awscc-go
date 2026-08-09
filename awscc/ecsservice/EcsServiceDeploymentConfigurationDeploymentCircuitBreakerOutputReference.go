// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/ecsservice/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EcsServiceDeploymentConfigurationDeploymentCircuitBreakerOutputReference interface {
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
	Enable() interface{}
	SetEnable(val interface{})
	EnableInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ResetOnHealthyTask() interface{}
	SetResetOnHealthyTask(val interface{})
	ResetOnHealthyTaskInput() interface{}
	Rollback() interface{}
	SetRollback(val interface{})
	RollbackInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	ThresholdConfiguration() EcsServiceDeploymentConfigurationDeploymentCircuitBreakerThresholdConfigurationOutputReference
	ThresholdConfigurationInput() interface{}
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
	PutThresholdConfiguration(value *EcsServiceDeploymentConfigurationDeploymentCircuitBreakerThresholdConfiguration)
	ResetEnable()
	ResetResetOnHealthyTask()
	ResetRollback()
	ResetThresholdConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EcsServiceDeploymentConfigurationDeploymentCircuitBreakerOutputReference
type jsiiProxy_EcsServiceDeploymentConfigurationDeploymentCircuitBreakerOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationDeploymentCircuitBreakerOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationDeploymentCircuitBreakerOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationDeploymentCircuitBreakerOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationDeploymentCircuitBreakerOutputReference) Enable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationDeploymentCircuitBreakerOutputReference) EnableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationDeploymentCircuitBreakerOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationDeploymentCircuitBreakerOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationDeploymentCircuitBreakerOutputReference) ResetOnHealthyTask() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resetOnHealthyTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationDeploymentCircuitBreakerOutputReference) ResetOnHealthyTaskInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resetOnHealthyTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationDeploymentCircuitBreakerOutputReference) Rollback() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rollback",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationDeploymentCircuitBreakerOutputReference) RollbackInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rollbackInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationDeploymentCircuitBreakerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationDeploymentCircuitBreakerOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationDeploymentCircuitBreakerOutputReference) ThresholdConfiguration() EcsServiceDeploymentConfigurationDeploymentCircuitBreakerThresholdConfigurationOutputReference {
	var returns EcsServiceDeploymentConfigurationDeploymentCircuitBreakerThresholdConfigurationOutputReference
	_jsii_.Get(
		j,
		"thresholdConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationDeploymentCircuitBreakerOutputReference) ThresholdConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"thresholdConfigurationInput",
		&returns,
	)
	return returns
}


func NewEcsServiceDeploymentConfigurationDeploymentCircuitBreakerOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) EcsServiceDeploymentConfigurationDeploymentCircuitBreakerOutputReference {
	_init_.Initialize()

	if err := validateNewEcsServiceDeploymentConfigurationDeploymentCircuitBreakerOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_EcsServiceDeploymentConfigurationDeploymentCircuitBreakerOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.ecsService.EcsServiceDeploymentConfigurationDeploymentCircuitBreakerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEcsServiceDeploymentConfigurationDeploymentCircuitBreakerOutputReference_Override(e EcsServiceDeploymentConfigurationDeploymentCircuitBreakerOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.ecsService.EcsServiceDeploymentConfigurationDeploymentCircuitBreakerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationDeploymentCircuitBreakerOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationDeploymentCircuitBreakerOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationDeploymentCircuitBreakerOutputReference)SetEnable(val interface{}) {
	if err := j.validateSetEnableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enable",
		val,
	)
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationDeploymentCircuitBreakerOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationDeploymentCircuitBreakerOutputReference)SetResetOnHealthyTask(val interface{}) {
	if err := j.validateSetResetOnHealthyTaskParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resetOnHealthyTask",
		val,
	)
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationDeploymentCircuitBreakerOutputReference)SetRollback(val interface{}) {
	if err := j.validateSetRollbackParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rollback",
		val,
	)
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationDeploymentCircuitBreakerOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationDeploymentCircuitBreakerOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EcsServiceDeploymentConfigurationDeploymentCircuitBreakerOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsServiceDeploymentConfigurationDeploymentCircuitBreakerOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_EcsServiceDeploymentConfigurationDeploymentCircuitBreakerOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_EcsServiceDeploymentConfigurationDeploymentCircuitBreakerOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_EcsServiceDeploymentConfigurationDeploymentCircuitBreakerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_EcsServiceDeploymentConfigurationDeploymentCircuitBreakerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_EcsServiceDeploymentConfigurationDeploymentCircuitBreakerOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_EcsServiceDeploymentConfigurationDeploymentCircuitBreakerOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_EcsServiceDeploymentConfigurationDeploymentCircuitBreakerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_EcsServiceDeploymentConfigurationDeploymentCircuitBreakerOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_EcsServiceDeploymentConfigurationDeploymentCircuitBreakerOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsServiceDeploymentConfigurationDeploymentCircuitBreakerOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_EcsServiceDeploymentConfigurationDeploymentCircuitBreakerOutputReference) PutThresholdConfiguration(value *EcsServiceDeploymentConfigurationDeploymentCircuitBreakerThresholdConfiguration) {
	if err := e.validatePutThresholdConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putThresholdConfiguration",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsServiceDeploymentConfigurationDeploymentCircuitBreakerOutputReference) ResetEnable() {
	_jsii_.InvokeVoid(
		e,
		"resetEnable",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsServiceDeploymentConfigurationDeploymentCircuitBreakerOutputReference) ResetResetOnHealthyTask() {
	_jsii_.InvokeVoid(
		e,
		"resetResetOnHealthyTask",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsServiceDeploymentConfigurationDeploymentCircuitBreakerOutputReference) ResetRollback() {
	_jsii_.InvokeVoid(
		e,
		"resetRollback",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsServiceDeploymentConfigurationDeploymentCircuitBreakerOutputReference) ResetThresholdConfiguration() {
	_jsii_.InvokeVoid(
		e,
		"resetThresholdConfiguration",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsServiceDeploymentConfigurationDeploymentCircuitBreakerOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (e *jsiiProxy_EcsServiceDeploymentConfigurationDeploymentCircuitBreakerOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

