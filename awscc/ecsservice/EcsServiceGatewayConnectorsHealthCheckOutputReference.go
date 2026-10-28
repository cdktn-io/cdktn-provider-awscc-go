// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/ecsservice/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EcsServiceGatewayConnectorsHealthCheckOutputReference interface {
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
	HealthCheckIntervalSeconds() *float64
	SetHealthCheckIntervalSeconds(val *float64)
	HealthCheckIntervalSecondsInput() *float64
	HealthCheckPath() *string
	SetHealthCheckPath(val *string)
	HealthCheckPathInput() *string
	HealthCheckPort() *float64
	SetHealthCheckPort(val *float64)
	HealthCheckPortInput() *float64
	HealthCheckProtocol() *string
	SetHealthCheckProtocol(val *string)
	HealthCheckProtocolInput() *string
	HealthCheckTimeoutSeconds() *float64
	SetHealthCheckTimeoutSeconds(val *float64)
	HealthCheckTimeoutSecondsInput() *float64
	HealthyThresholdCount() *float64
	SetHealthyThresholdCount(val *float64)
	HealthyThresholdCountInput() *float64
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Matcher() EcsServiceGatewayConnectorsHealthCheckMatcherOutputReference
	MatcherInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UnhealthyThresholdCount() *float64
	SetUnhealthyThresholdCount(val *float64)
	UnhealthyThresholdCountInput() *float64
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
	PutMatcher(value *EcsServiceGatewayConnectorsHealthCheckMatcher)
	ResetHealthCheckIntervalSeconds()
	ResetHealthCheckPath()
	ResetHealthCheckPort()
	ResetHealthCheckProtocol()
	ResetHealthCheckTimeoutSeconds()
	ResetHealthyThresholdCount()
	ResetMatcher()
	ResetUnhealthyThresholdCount()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EcsServiceGatewayConnectorsHealthCheckOutputReference
type jsiiProxy_EcsServiceGatewayConnectorsHealthCheckOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_EcsServiceGatewayConnectorsHealthCheckOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceGatewayConnectorsHealthCheckOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceGatewayConnectorsHealthCheckOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceGatewayConnectorsHealthCheckOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceGatewayConnectorsHealthCheckOutputReference) HealthCheckIntervalSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckIntervalSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceGatewayConnectorsHealthCheckOutputReference) HealthCheckIntervalSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckIntervalSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceGatewayConnectorsHealthCheckOutputReference) HealthCheckPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceGatewayConnectorsHealthCheckOutputReference) HealthCheckPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceGatewayConnectorsHealthCheckOutputReference) HealthCheckPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceGatewayConnectorsHealthCheckOutputReference) HealthCheckPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceGatewayConnectorsHealthCheckOutputReference) HealthCheckProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceGatewayConnectorsHealthCheckOutputReference) HealthCheckProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceGatewayConnectorsHealthCheckOutputReference) HealthCheckTimeoutSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckTimeoutSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceGatewayConnectorsHealthCheckOutputReference) HealthCheckTimeoutSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckTimeoutSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceGatewayConnectorsHealthCheckOutputReference) HealthyThresholdCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthyThresholdCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceGatewayConnectorsHealthCheckOutputReference) HealthyThresholdCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthyThresholdCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceGatewayConnectorsHealthCheckOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceGatewayConnectorsHealthCheckOutputReference) Matcher() EcsServiceGatewayConnectorsHealthCheckMatcherOutputReference {
	var returns EcsServiceGatewayConnectorsHealthCheckMatcherOutputReference
	_jsii_.Get(
		j,
		"matcher",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceGatewayConnectorsHealthCheckOutputReference) MatcherInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"matcherInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceGatewayConnectorsHealthCheckOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceGatewayConnectorsHealthCheckOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceGatewayConnectorsHealthCheckOutputReference) UnhealthyThresholdCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"unhealthyThresholdCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceGatewayConnectorsHealthCheckOutputReference) UnhealthyThresholdCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"unhealthyThresholdCountInput",
		&returns,
	)
	return returns
}


func NewEcsServiceGatewayConnectorsHealthCheckOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) EcsServiceGatewayConnectorsHealthCheckOutputReference {
	_init_.Initialize()

	if err := validateNewEcsServiceGatewayConnectorsHealthCheckOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_EcsServiceGatewayConnectorsHealthCheckOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.ecsService.EcsServiceGatewayConnectorsHealthCheckOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEcsServiceGatewayConnectorsHealthCheckOutputReference_Override(e EcsServiceGatewayConnectorsHealthCheckOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.ecsService.EcsServiceGatewayConnectorsHealthCheckOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EcsServiceGatewayConnectorsHealthCheckOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EcsServiceGatewayConnectorsHealthCheckOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EcsServiceGatewayConnectorsHealthCheckOutputReference)SetHealthCheckIntervalSeconds(val *float64) {
	if err := j.validateSetHealthCheckIntervalSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheckIntervalSeconds",
		val,
	)
}

func (j *jsiiProxy_EcsServiceGatewayConnectorsHealthCheckOutputReference)SetHealthCheckPath(val *string) {
	if err := j.validateSetHealthCheckPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheckPath",
		val,
	)
}

func (j *jsiiProxy_EcsServiceGatewayConnectorsHealthCheckOutputReference)SetHealthCheckPort(val *float64) {
	if err := j.validateSetHealthCheckPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheckPort",
		val,
	)
}

func (j *jsiiProxy_EcsServiceGatewayConnectorsHealthCheckOutputReference)SetHealthCheckProtocol(val *string) {
	if err := j.validateSetHealthCheckProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheckProtocol",
		val,
	)
}

func (j *jsiiProxy_EcsServiceGatewayConnectorsHealthCheckOutputReference)SetHealthCheckTimeoutSeconds(val *float64) {
	if err := j.validateSetHealthCheckTimeoutSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheckTimeoutSeconds",
		val,
	)
}

func (j *jsiiProxy_EcsServiceGatewayConnectorsHealthCheckOutputReference)SetHealthyThresholdCount(val *float64) {
	if err := j.validateSetHealthyThresholdCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthyThresholdCount",
		val,
	)
}

func (j *jsiiProxy_EcsServiceGatewayConnectorsHealthCheckOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EcsServiceGatewayConnectorsHealthCheckOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EcsServiceGatewayConnectorsHealthCheckOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EcsServiceGatewayConnectorsHealthCheckOutputReference)SetUnhealthyThresholdCount(val *float64) {
	if err := j.validateSetUnhealthyThresholdCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unhealthyThresholdCount",
		val,
	)
}

func (e *jsiiProxy_EcsServiceGatewayConnectorsHealthCheckOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsServiceGatewayConnectorsHealthCheckOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_EcsServiceGatewayConnectorsHealthCheckOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_EcsServiceGatewayConnectorsHealthCheckOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_EcsServiceGatewayConnectorsHealthCheckOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_EcsServiceGatewayConnectorsHealthCheckOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_EcsServiceGatewayConnectorsHealthCheckOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_EcsServiceGatewayConnectorsHealthCheckOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_EcsServiceGatewayConnectorsHealthCheckOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_EcsServiceGatewayConnectorsHealthCheckOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_EcsServiceGatewayConnectorsHealthCheckOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsServiceGatewayConnectorsHealthCheckOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_EcsServiceGatewayConnectorsHealthCheckOutputReference) PutMatcher(value *EcsServiceGatewayConnectorsHealthCheckMatcher) {
	if err := e.validatePutMatcherParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putMatcher",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsServiceGatewayConnectorsHealthCheckOutputReference) ResetHealthCheckIntervalSeconds() {
	_jsii_.InvokeVoid(
		e,
		"resetHealthCheckIntervalSeconds",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsServiceGatewayConnectorsHealthCheckOutputReference) ResetHealthCheckPath() {
	_jsii_.InvokeVoid(
		e,
		"resetHealthCheckPath",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsServiceGatewayConnectorsHealthCheckOutputReference) ResetHealthCheckPort() {
	_jsii_.InvokeVoid(
		e,
		"resetHealthCheckPort",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsServiceGatewayConnectorsHealthCheckOutputReference) ResetHealthCheckProtocol() {
	_jsii_.InvokeVoid(
		e,
		"resetHealthCheckProtocol",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsServiceGatewayConnectorsHealthCheckOutputReference) ResetHealthCheckTimeoutSeconds() {
	_jsii_.InvokeVoid(
		e,
		"resetHealthCheckTimeoutSeconds",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsServiceGatewayConnectorsHealthCheckOutputReference) ResetHealthyThresholdCount() {
	_jsii_.InvokeVoid(
		e,
		"resetHealthyThresholdCount",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsServiceGatewayConnectorsHealthCheckOutputReference) ResetMatcher() {
	_jsii_.InvokeVoid(
		e,
		"resetMatcher",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsServiceGatewayConnectorsHealthCheckOutputReference) ResetUnhealthyThresholdCount() {
	_jsii_.InvokeVoid(
		e,
		"resetUnhealthyThresholdCount",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsServiceGatewayConnectorsHealthCheckOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (e *jsiiProxy_EcsServiceGatewayConnectorsHealthCheckOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

