// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoregatewayrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/bedrockagentcoregatewayrule/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackOutputReference interface {
	cdktn.ComplexObject
	CircuitBreaker() BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackCircuitBreakerOutputReference
	CircuitBreakerInput() interface{}
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
	FirstByteTimeoutMilliseconds() *float64
	SetFirstByteTimeoutMilliseconds(val *float64)
	FirstByteTimeoutMillisecondsInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	TargetFallback() BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackTargetFallbackList
	TargetFallbackInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TriggerOn() BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackTriggerOnOutputReference
	TriggerOnInput() interface{}
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
	PutCircuitBreaker(value *BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackCircuitBreaker)
	PutTargetFallback(value interface{})
	PutTriggerOn(value *BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackTriggerOn)
	ResetCircuitBreaker()
	ResetFirstByteTimeoutMilliseconds()
	ResetTargetFallback()
	ResetTriggerOn()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackOutputReference
type jsiiProxy_BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackOutputReference) CircuitBreaker() BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackCircuitBreakerOutputReference {
	var returns BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackCircuitBreakerOutputReference
	_jsii_.Get(
		j,
		"circuitBreaker",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackOutputReference) CircuitBreakerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"circuitBreakerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackOutputReference) FirstByteTimeoutMilliseconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"firstByteTimeoutMilliseconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackOutputReference) FirstByteTimeoutMillisecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"firstByteTimeoutMillisecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackOutputReference) TargetFallback() BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackTargetFallbackList {
	var returns BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackTargetFallbackList
	_jsii_.Get(
		j,
		"targetFallback",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackOutputReference) TargetFallbackInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targetFallbackInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackOutputReference) TriggerOn() BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackTriggerOnOutputReference {
	var returns BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackTriggerOnOutputReference
	_jsii_.Get(
		j,
		"triggerOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackOutputReference) TriggerOnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"triggerOnInput",
		&returns,
	)
	return returns
}


func NewBedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackOutputReference {
	_init_.Initialize()

	if err := validateNewBedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.bedrockagentcoreGatewayRule.BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackOutputReference_Override(b BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.bedrockagentcoreGatewayRule.BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackOutputReference)SetFirstByteTimeoutMilliseconds(val *float64) {
	if err := j.validateSetFirstByteTimeoutMillisecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"firstByteTimeoutMilliseconds",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackOutputReference) PutCircuitBreaker(value *BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackCircuitBreaker) {
	if err := b.validatePutCircuitBreakerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putCircuitBreaker",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackOutputReference) PutTargetFallback(value interface{}) {
	if err := b.validatePutTargetFallbackParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putTargetFallback",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackOutputReference) PutTriggerOn(value *BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackTriggerOn) {
	if err := b.validatePutTriggerOnParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putTriggerOn",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackOutputReference) ResetCircuitBreaker() {
	_jsii_.InvokeVoid(
		b,
		"resetCircuitBreaker",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackOutputReference) ResetFirstByteTimeoutMilliseconds() {
	_jsii_.InvokeVoid(
		b,
		"resetFirstByteTimeoutMilliseconds",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackOutputReference) ResetTargetFallback() {
	_jsii_.InvokeVoid(
		b,
		"resetTargetFallback",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackOutputReference) ResetTriggerOn() {
	_jsii_.InvokeVoid(
		b,
		"resetTriggerOn",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := b.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

