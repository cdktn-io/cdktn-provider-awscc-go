// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package xraysamplingrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/xraysamplingrule/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type XraySamplingRuleSamplingRuleUpdateOutputReference interface {
	cdktn.ComplexObject
	Attributes() *map[string]*string
	SetAttributes(val *map[string]*string)
	AttributesInput() *map[string]*string
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
	FixedRate() *float64
	SetFixedRate(val *float64)
	FixedRateInput() *float64
	// Experimental.
	Fqn() *string
	Host() *string
	SetHost(val *string)
	HostInput() *string
	HttpMethod() *string
	SetHttpMethod(val *string)
	HttpMethodInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Priority() *float64
	SetPriority(val *float64)
	PriorityInput() *float64
	ReservoirSize() *float64
	SetReservoirSize(val *float64)
	ReservoirSizeInput() *float64
	ResourceArn() *string
	SetResourceArn(val *string)
	ResourceArnInput() *string
	RuleArn() *string
	SetRuleArn(val *string)
	RuleArnInput() *string
	RuleName() *string
	SetRuleName(val *string)
	RuleNameInput() *string
	SamplingRateBoost() XraySamplingRuleSamplingRuleUpdateSamplingRateBoostOutputReference
	SamplingRateBoostInput() interface{}
	ServiceName() *string
	SetServiceName(val *string)
	ServiceNameInput() *string
	ServiceType() *string
	SetServiceType(val *string)
	ServiceTypeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UrlPath() *string
	SetUrlPath(val *string)
	UrlPathInput() *string
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
	PutSamplingRateBoost(value *XraySamplingRuleSamplingRuleUpdateSamplingRateBoost)
	ResetAttributes()
	ResetFixedRate()
	ResetHost()
	ResetHttpMethod()
	ResetPriority()
	ResetReservoirSize()
	ResetResourceArn()
	ResetRuleArn()
	ResetRuleName()
	ResetSamplingRateBoost()
	ResetServiceName()
	ResetServiceType()
	ResetUrlPath()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for XraySamplingRuleSamplingRuleUpdateOutputReference
type jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference) Attributes() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"attributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference) AttributesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"attributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference) FixedRate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fixedRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference) FixedRateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fixedRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference) Host() *string {
	var returns *string
	_jsii_.Get(
		j,
		"host",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference) HostInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference) HttpMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference) HttpMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference) PriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference) ReservoirSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"reservoirSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference) ReservoirSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"reservoirSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference) ResourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference) ResourceArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference) RuleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference) RuleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference) RuleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference) RuleNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference) SamplingRateBoost() XraySamplingRuleSamplingRuleUpdateSamplingRateBoostOutputReference {
	var returns XraySamplingRuleSamplingRuleUpdateSamplingRateBoostOutputReference
	_jsii_.Get(
		j,
		"samplingRateBoost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference) SamplingRateBoostInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"samplingRateBoostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference) ServiceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference) ServiceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference) ServiceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference) ServiceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference) UrlPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference) UrlPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlPathInput",
		&returns,
	)
	return returns
}


func NewXraySamplingRuleSamplingRuleUpdateOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) XraySamplingRuleSamplingRuleUpdateOutputReference {
	_init_.Initialize()

	if err := validateNewXraySamplingRuleSamplingRuleUpdateOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.xraySamplingRule.XraySamplingRuleSamplingRuleUpdateOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewXraySamplingRuleSamplingRuleUpdateOutputReference_Override(x XraySamplingRuleSamplingRuleUpdateOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.xraySamplingRule.XraySamplingRuleSamplingRuleUpdateOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		x,
	)
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference)SetAttributes(val *map[string]*string) {
	if err := j.validateSetAttributesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"attributes",
		val,
	)
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference)SetFixedRate(val *float64) {
	if err := j.validateSetFixedRateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fixedRate",
		val,
	)
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference)SetHost(val *string) {
	if err := j.validateSetHostParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"host",
		val,
	)
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference)SetHttpMethod(val *string) {
	if err := j.validateSetHttpMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpMethod",
		val,
	)
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference)SetPriority(val *float64) {
	if err := j.validateSetPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference)SetReservoirSize(val *float64) {
	if err := j.validateSetReservoirSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reservoirSize",
		val,
	)
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference)SetResourceArn(val *string) {
	if err := j.validateSetResourceArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceArn",
		val,
	)
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference)SetRuleArn(val *string) {
	if err := j.validateSetRuleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ruleArn",
		val,
	)
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference)SetRuleName(val *string) {
	if err := j.validateSetRuleNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ruleName",
		val,
	)
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference)SetServiceName(val *string) {
	if err := j.validateSetServiceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceName",
		val,
	)
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference)SetServiceType(val *string) {
	if err := j.validateSetServiceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceType",
		val,
	)
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference)SetUrlPath(val *string) {
	if err := j.validateSetUrlPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"urlPath",
		val,
	)
}

func (x *jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		x,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := x.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		x,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := x.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		x,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := x.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		x,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := x.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		x,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := x.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		x,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := x.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		x,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := x.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		x,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := x.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		x,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := x.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		x,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		x,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := x.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		x,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference) PutSamplingRateBoost(value *XraySamplingRuleSamplingRuleUpdateSamplingRateBoost) {
	if err := x.validatePutSamplingRateBoostParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		x,
		"putSamplingRateBoost",
		[]interface{}{value},
	)
}

func (x *jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference) ResetAttributes() {
	_jsii_.InvokeVoid(
		x,
		"resetAttributes",
		nil, // no parameters
	)
}

func (x *jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference) ResetFixedRate() {
	_jsii_.InvokeVoid(
		x,
		"resetFixedRate",
		nil, // no parameters
	)
}

func (x *jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference) ResetHost() {
	_jsii_.InvokeVoid(
		x,
		"resetHost",
		nil, // no parameters
	)
}

func (x *jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference) ResetHttpMethod() {
	_jsii_.InvokeVoid(
		x,
		"resetHttpMethod",
		nil, // no parameters
	)
}

func (x *jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference) ResetPriority() {
	_jsii_.InvokeVoid(
		x,
		"resetPriority",
		nil, // no parameters
	)
}

func (x *jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference) ResetReservoirSize() {
	_jsii_.InvokeVoid(
		x,
		"resetReservoirSize",
		nil, // no parameters
	)
}

func (x *jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference) ResetResourceArn() {
	_jsii_.InvokeVoid(
		x,
		"resetResourceArn",
		nil, // no parameters
	)
}

func (x *jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference) ResetRuleArn() {
	_jsii_.InvokeVoid(
		x,
		"resetRuleArn",
		nil, // no parameters
	)
}

func (x *jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference) ResetRuleName() {
	_jsii_.InvokeVoid(
		x,
		"resetRuleName",
		nil, // no parameters
	)
}

func (x *jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference) ResetSamplingRateBoost() {
	_jsii_.InvokeVoid(
		x,
		"resetSamplingRateBoost",
		nil, // no parameters
	)
}

func (x *jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference) ResetServiceName() {
	_jsii_.InvokeVoid(
		x,
		"resetServiceName",
		nil, // no parameters
	)
}

func (x *jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference) ResetServiceType() {
	_jsii_.InvokeVoid(
		x,
		"resetServiceType",
		nil, // no parameters
	)
}

func (x *jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference) ResetUrlPath() {
	_jsii_.InvokeVoid(
		x,
		"resetUrlPath",
		nil, // no parameters
	)
}

func (x *jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := x.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		x,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XraySamplingRuleSamplingRuleUpdateOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		x,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

