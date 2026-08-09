// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/ecsservice/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EcsServiceDeploymentConfigurationOutputReference interface {
	cdktn.ComplexObject
	Alarms() EcsServiceDeploymentConfigurationAlarmsOutputReference
	AlarmsInput() interface{}
	BakeTimeInMinutes() *float64
	SetBakeTimeInMinutes(val *float64)
	BakeTimeInMinutesInput() *float64
	CanaryConfiguration() EcsServiceDeploymentConfigurationCanaryConfigurationOutputReference
	CanaryConfigurationInput() interface{}
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
	DeploymentCircuitBreaker() EcsServiceDeploymentConfigurationDeploymentCircuitBreakerOutputReference
	DeploymentCircuitBreakerInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LifecycleHooks() EcsServiceDeploymentConfigurationLifecycleHooksList
	LifecycleHooksInput() interface{}
	LinearConfiguration() EcsServiceDeploymentConfigurationLinearConfigurationOutputReference
	LinearConfigurationInput() interface{}
	MaximumPercent() *float64
	SetMaximumPercent(val *float64)
	MaximumPercentInput() *float64
	MinimumHealthyPercent() *float64
	SetMinimumHealthyPercent(val *float64)
	MinimumHealthyPercentInput() *float64
	Strategy() *string
	SetStrategy(val *string)
	StrategyInput() *string
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
	PutAlarms(value *EcsServiceDeploymentConfigurationAlarms)
	PutCanaryConfiguration(value *EcsServiceDeploymentConfigurationCanaryConfiguration)
	PutDeploymentCircuitBreaker(value *EcsServiceDeploymentConfigurationDeploymentCircuitBreaker)
	PutLifecycleHooks(value interface{})
	PutLinearConfiguration(value *EcsServiceDeploymentConfigurationLinearConfiguration)
	ResetAlarms()
	ResetBakeTimeInMinutes()
	ResetCanaryConfiguration()
	ResetDeploymentCircuitBreaker()
	ResetLifecycleHooks()
	ResetLinearConfiguration()
	ResetMaximumPercent()
	ResetMinimumHealthyPercent()
	ResetStrategy()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EcsServiceDeploymentConfigurationOutputReference
type jsiiProxy_EcsServiceDeploymentConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) Alarms() EcsServiceDeploymentConfigurationAlarmsOutputReference {
	var returns EcsServiceDeploymentConfigurationAlarmsOutputReference
	_jsii_.Get(
		j,
		"alarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) AlarmsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alarmsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) BakeTimeInMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bakeTimeInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) BakeTimeInMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bakeTimeInMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) CanaryConfiguration() EcsServiceDeploymentConfigurationCanaryConfigurationOutputReference {
	var returns EcsServiceDeploymentConfigurationCanaryConfigurationOutputReference
	_jsii_.Get(
		j,
		"canaryConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) CanaryConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"canaryConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) DeploymentCircuitBreaker() EcsServiceDeploymentConfigurationDeploymentCircuitBreakerOutputReference {
	var returns EcsServiceDeploymentConfigurationDeploymentCircuitBreakerOutputReference
	_jsii_.Get(
		j,
		"deploymentCircuitBreaker",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) DeploymentCircuitBreakerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deploymentCircuitBreakerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) LifecycleHooks() EcsServiceDeploymentConfigurationLifecycleHooksList {
	var returns EcsServiceDeploymentConfigurationLifecycleHooksList
	_jsii_.Get(
		j,
		"lifecycleHooks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) LifecycleHooksInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lifecycleHooksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) LinearConfiguration() EcsServiceDeploymentConfigurationLinearConfigurationOutputReference {
	var returns EcsServiceDeploymentConfigurationLinearConfigurationOutputReference
	_jsii_.Get(
		j,
		"linearConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) LinearConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"linearConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) MaximumPercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumPercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) MaximumPercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumPercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) MinimumHealthyPercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumHealthyPercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) MinimumHealthyPercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumHealthyPercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) Strategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) StrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"strategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEcsServiceDeploymentConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) EcsServiceDeploymentConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewEcsServiceDeploymentConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_EcsServiceDeploymentConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.ecsService.EcsServiceDeploymentConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEcsServiceDeploymentConfigurationOutputReference_Override(e EcsServiceDeploymentConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.ecsService.EcsServiceDeploymentConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference)SetBakeTimeInMinutes(val *float64) {
	if err := j.validateSetBakeTimeInMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bakeTimeInMinutes",
		val,
	)
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference)SetMaximumPercent(val *float64) {
	if err := j.validateSetMaximumPercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumPercent",
		val,
	)
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference)SetMinimumHealthyPercent(val *float64) {
	if err := j.validateSetMinimumHealthyPercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimumHealthyPercent",
		val,
	)
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference)SetStrategy(val *string) {
	if err := j.validateSetStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"strategy",
		val,
	)
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) PutAlarms(value *EcsServiceDeploymentConfigurationAlarms) {
	if err := e.validatePutAlarmsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putAlarms",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) PutCanaryConfiguration(value *EcsServiceDeploymentConfigurationCanaryConfiguration) {
	if err := e.validatePutCanaryConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putCanaryConfiguration",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) PutDeploymentCircuitBreaker(value *EcsServiceDeploymentConfigurationDeploymentCircuitBreaker) {
	if err := e.validatePutDeploymentCircuitBreakerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putDeploymentCircuitBreaker",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) PutLifecycleHooks(value interface{}) {
	if err := e.validatePutLifecycleHooksParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putLifecycleHooks",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) PutLinearConfiguration(value *EcsServiceDeploymentConfigurationLinearConfiguration) {
	if err := e.validatePutLinearConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putLinearConfiguration",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) ResetAlarms() {
	_jsii_.InvokeVoid(
		e,
		"resetAlarms",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) ResetBakeTimeInMinutes() {
	_jsii_.InvokeVoid(
		e,
		"resetBakeTimeInMinutes",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) ResetCanaryConfiguration() {
	_jsii_.InvokeVoid(
		e,
		"resetCanaryConfiguration",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) ResetDeploymentCircuitBreaker() {
	_jsii_.InvokeVoid(
		e,
		"resetDeploymentCircuitBreaker",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) ResetLifecycleHooks() {
	_jsii_.InvokeVoid(
		e,
		"resetLifecycleHooks",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) ResetLinearConfiguration() {
	_jsii_.InvokeVoid(
		e,
		"resetLinearConfiguration",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) ResetMaximumPercent() {
	_jsii_.InvokeVoid(
		e,
		"resetMaximumPercent",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) ResetMinimumHealthyPercent() {
	_jsii_.InvokeVoid(
		e,
		"resetMinimumHealthyPercent",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) ResetStrategy() {
	_jsii_.InvokeVoid(
		e,
		"resetStrategy",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (e *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

