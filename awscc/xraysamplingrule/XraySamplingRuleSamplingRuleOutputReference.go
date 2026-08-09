// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package xraysamplingrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/xraysamplingrule/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type XraySamplingRuleSamplingRuleOutputReference interface {
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
	SamplingRateBoost() XraySamplingRuleSamplingRuleSamplingRateBoostOutputReference
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
	Version() *float64
	SetVersion(val *float64)
	VersionInput() *float64
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
	PutSamplingRateBoost(value *XraySamplingRuleSamplingRuleSamplingRateBoost)
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
	ResetVersion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for XraySamplingRuleSamplingRuleOutputReference
type jsiiProxy_XraySamplingRuleSamplingRuleOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference) Attributes() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"attributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference) AttributesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"attributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference) FixedRate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fixedRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference) FixedRateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fixedRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference) Host() *string {
	var returns *string
	_jsii_.Get(
		j,
		"host",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference) HostInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference) HttpMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference) HttpMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference) PriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference) ReservoirSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"reservoirSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference) ReservoirSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"reservoirSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference) ResourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference) ResourceArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference) RuleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference) RuleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference) RuleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference) RuleNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference) SamplingRateBoost() XraySamplingRuleSamplingRuleSamplingRateBoostOutputReference {
	var returns XraySamplingRuleSamplingRuleSamplingRateBoostOutputReference
	_jsii_.Get(
		j,
		"samplingRateBoost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference) SamplingRateBoostInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"samplingRateBoostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference) ServiceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference) ServiceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference) ServiceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference) ServiceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference) UrlPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference) UrlPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference) Version() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference) VersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


func NewXraySamplingRuleSamplingRuleOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) XraySamplingRuleSamplingRuleOutputReference {
	_init_.Initialize()

	if err := validateNewXraySamplingRuleSamplingRuleOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_XraySamplingRuleSamplingRuleOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.xraySamplingRule.XraySamplingRuleSamplingRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewXraySamplingRuleSamplingRuleOutputReference_Override(x XraySamplingRuleSamplingRuleOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.xraySamplingRule.XraySamplingRuleSamplingRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		x,
	)
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference)SetAttributes(val *map[string]*string) {
	if err := j.validateSetAttributesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"attributes",
		val,
	)
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference)SetFixedRate(val *float64) {
	if err := j.validateSetFixedRateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fixedRate",
		val,
	)
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference)SetHost(val *string) {
	if err := j.validateSetHostParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"host",
		val,
	)
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference)SetHttpMethod(val *string) {
	if err := j.validateSetHttpMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpMethod",
		val,
	)
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference)SetPriority(val *float64) {
	if err := j.validateSetPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference)SetReservoirSize(val *float64) {
	if err := j.validateSetReservoirSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reservoirSize",
		val,
	)
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference)SetResourceArn(val *string) {
	if err := j.validateSetResourceArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceArn",
		val,
	)
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference)SetRuleArn(val *string) {
	if err := j.validateSetRuleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ruleArn",
		val,
	)
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference)SetRuleName(val *string) {
	if err := j.validateSetRuleNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ruleName",
		val,
	)
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference)SetServiceName(val *string) {
	if err := j.validateSetServiceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceName",
		val,
	)
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference)SetServiceType(val *string) {
	if err := j.validateSetServiceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceType",
		val,
	)
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference)SetUrlPath(val *string) {
	if err := j.validateSetUrlPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"urlPath",
		val,
	)
}

func (j *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference)SetVersion(val *float64) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (x *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		x,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (x *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (x *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (x *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (x *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (x *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (x *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (x *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (x *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (x *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		x,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (x *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference) PutSamplingRateBoost(value *XraySamplingRuleSamplingRuleSamplingRateBoost) {
	if err := x.validatePutSamplingRateBoostParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		x,
		"putSamplingRateBoost",
		[]interface{}{value},
	)
}

func (x *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference) ResetAttributes() {
	_jsii_.InvokeVoid(
		x,
		"resetAttributes",
		nil, // no parameters
	)
}

func (x *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference) ResetFixedRate() {
	_jsii_.InvokeVoid(
		x,
		"resetFixedRate",
		nil, // no parameters
	)
}

func (x *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference) ResetHost() {
	_jsii_.InvokeVoid(
		x,
		"resetHost",
		nil, // no parameters
	)
}

func (x *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference) ResetHttpMethod() {
	_jsii_.InvokeVoid(
		x,
		"resetHttpMethod",
		nil, // no parameters
	)
}

func (x *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference) ResetPriority() {
	_jsii_.InvokeVoid(
		x,
		"resetPriority",
		nil, // no parameters
	)
}

func (x *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference) ResetReservoirSize() {
	_jsii_.InvokeVoid(
		x,
		"resetReservoirSize",
		nil, // no parameters
	)
}

func (x *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference) ResetResourceArn() {
	_jsii_.InvokeVoid(
		x,
		"resetResourceArn",
		nil, // no parameters
	)
}

func (x *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference) ResetRuleArn() {
	_jsii_.InvokeVoid(
		x,
		"resetRuleArn",
		nil, // no parameters
	)
}

func (x *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference) ResetRuleName() {
	_jsii_.InvokeVoid(
		x,
		"resetRuleName",
		nil, // no parameters
	)
}

func (x *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference) ResetSamplingRateBoost() {
	_jsii_.InvokeVoid(
		x,
		"resetSamplingRateBoost",
		nil, // no parameters
	)
}

func (x *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference) ResetServiceName() {
	_jsii_.InvokeVoid(
		x,
		"resetServiceName",
		nil, // no parameters
	)
}

func (x *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference) ResetServiceType() {
	_jsii_.InvokeVoid(
		x,
		"resetServiceType",
		nil, // no parameters
	)
}

func (x *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference) ResetUrlPath() {
	_jsii_.InvokeVoid(
		x,
		"resetUrlPath",
		nil, // no parameters
	)
}

func (x *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference) ResetVersion() {
	_jsii_.InvokeVoid(
		x,
		"resetVersion",
		nil, // no parameters
	)
}

func (x *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (x *jsiiProxy_XraySamplingRuleSamplingRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		x,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

