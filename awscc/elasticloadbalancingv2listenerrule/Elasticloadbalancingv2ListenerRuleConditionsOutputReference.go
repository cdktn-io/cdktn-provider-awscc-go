// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package elasticloadbalancingv2listenerrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/elasticloadbalancingv2listenerrule/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Elasticloadbalancingv2ListenerRuleConditionsOutputReference interface {
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
	Field() *string
	SetField(val *string)
	FieldInput() *string
	// Experimental.
	Fqn() *string
	HostHeaderConfig() Elasticloadbalancingv2ListenerRuleConditionsHostHeaderConfigOutputReference
	HostHeaderConfigInput() interface{}
	HttpHeaderConfig() Elasticloadbalancingv2ListenerRuleConditionsHttpHeaderConfigOutputReference
	HttpHeaderConfigInput() interface{}
	HttpRequestMethodConfig() Elasticloadbalancingv2ListenerRuleConditionsHttpRequestMethodConfigOutputReference
	HttpRequestMethodConfigInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PathPatternConfig() Elasticloadbalancingv2ListenerRuleConditionsPathPatternConfigOutputReference
	PathPatternConfigInput() interface{}
	QueryStringConfig() Elasticloadbalancingv2ListenerRuleConditionsQueryStringConfigOutputReference
	QueryStringConfigInput() interface{}
	RegexValues() *[]*string
	SetRegexValues(val *[]*string)
	RegexValuesInput() *[]*string
	SourceIpConfig() Elasticloadbalancingv2ListenerRuleConditionsSourceIpConfigOutputReference
	SourceIpConfigInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Values() *[]*string
	SetValues(val *[]*string)
	ValuesInput() *[]*string
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
	PutHostHeaderConfig(value *Elasticloadbalancingv2ListenerRuleConditionsHostHeaderConfig)
	PutHttpHeaderConfig(value *Elasticloadbalancingv2ListenerRuleConditionsHttpHeaderConfig)
	PutHttpRequestMethodConfig(value *Elasticloadbalancingv2ListenerRuleConditionsHttpRequestMethodConfig)
	PutPathPatternConfig(value *Elasticloadbalancingv2ListenerRuleConditionsPathPatternConfig)
	PutQueryStringConfig(value *Elasticloadbalancingv2ListenerRuleConditionsQueryStringConfig)
	PutSourceIpConfig(value *Elasticloadbalancingv2ListenerRuleConditionsSourceIpConfig)
	ResetField()
	ResetHostHeaderConfig()
	ResetHttpHeaderConfig()
	ResetHttpRequestMethodConfig()
	ResetPathPatternConfig()
	ResetQueryStringConfig()
	ResetRegexValues()
	ResetSourceIpConfig()
	ResetValues()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Elasticloadbalancingv2ListenerRuleConditionsOutputReference
type jsiiProxy_Elasticloadbalancingv2ListenerRuleConditionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleConditionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleConditionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleConditionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleConditionsOutputReference) Field() *string {
	var returns *string
	_jsii_.Get(
		j,
		"field",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleConditionsOutputReference) FieldInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleConditionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleConditionsOutputReference) HostHeaderConfig() Elasticloadbalancingv2ListenerRuleConditionsHostHeaderConfigOutputReference {
	var returns Elasticloadbalancingv2ListenerRuleConditionsHostHeaderConfigOutputReference
	_jsii_.Get(
		j,
		"hostHeaderConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleConditionsOutputReference) HostHeaderConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hostHeaderConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleConditionsOutputReference) HttpHeaderConfig() Elasticloadbalancingv2ListenerRuleConditionsHttpHeaderConfigOutputReference {
	var returns Elasticloadbalancingv2ListenerRuleConditionsHttpHeaderConfigOutputReference
	_jsii_.Get(
		j,
		"httpHeaderConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleConditionsOutputReference) HttpHeaderConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpHeaderConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleConditionsOutputReference) HttpRequestMethodConfig() Elasticloadbalancingv2ListenerRuleConditionsHttpRequestMethodConfigOutputReference {
	var returns Elasticloadbalancingv2ListenerRuleConditionsHttpRequestMethodConfigOutputReference
	_jsii_.Get(
		j,
		"httpRequestMethodConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleConditionsOutputReference) HttpRequestMethodConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpRequestMethodConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleConditionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleConditionsOutputReference) PathPatternConfig() Elasticloadbalancingv2ListenerRuleConditionsPathPatternConfigOutputReference {
	var returns Elasticloadbalancingv2ListenerRuleConditionsPathPatternConfigOutputReference
	_jsii_.Get(
		j,
		"pathPatternConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleConditionsOutputReference) PathPatternConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pathPatternConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleConditionsOutputReference) QueryStringConfig() Elasticloadbalancingv2ListenerRuleConditionsQueryStringConfigOutputReference {
	var returns Elasticloadbalancingv2ListenerRuleConditionsQueryStringConfigOutputReference
	_jsii_.Get(
		j,
		"queryStringConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleConditionsOutputReference) QueryStringConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queryStringConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleConditionsOutputReference) RegexValues() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"regexValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleConditionsOutputReference) RegexValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"regexValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleConditionsOutputReference) SourceIpConfig() Elasticloadbalancingv2ListenerRuleConditionsSourceIpConfigOutputReference {
	var returns Elasticloadbalancingv2ListenerRuleConditionsSourceIpConfigOutputReference
	_jsii_.Get(
		j,
		"sourceIpConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleConditionsOutputReference) SourceIpConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceIpConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleConditionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleConditionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleConditionsOutputReference) Values() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"values",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleConditionsOutputReference) ValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"valuesInput",
		&returns,
	)
	return returns
}


func NewElasticloadbalancingv2ListenerRuleConditionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) Elasticloadbalancingv2ListenerRuleConditionsOutputReference {
	_init_.Initialize()

	if err := validateNewElasticloadbalancingv2ListenerRuleConditionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_Elasticloadbalancingv2ListenerRuleConditionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.elasticloadbalancingv2ListenerRule.Elasticloadbalancingv2ListenerRuleConditionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewElasticloadbalancingv2ListenerRuleConditionsOutputReference_Override(e Elasticloadbalancingv2ListenerRuleConditionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.elasticloadbalancingv2ListenerRule.Elasticloadbalancingv2ListenerRuleConditionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleConditionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleConditionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleConditionsOutputReference)SetField(val *string) {
	if err := j.validateSetFieldParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"field",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleConditionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleConditionsOutputReference)SetRegexValues(val *[]*string) {
	if err := j.validateSetRegexValuesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"regexValues",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleConditionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleConditionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleConditionsOutputReference)SetValues(val *[]*string) {
	if err := j.validateSetValuesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"values",
		val,
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerRuleConditionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerRuleConditionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_Elasticloadbalancingv2ListenerRuleConditionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_Elasticloadbalancingv2ListenerRuleConditionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_Elasticloadbalancingv2ListenerRuleConditionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_Elasticloadbalancingv2ListenerRuleConditionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_Elasticloadbalancingv2ListenerRuleConditionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_Elasticloadbalancingv2ListenerRuleConditionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_Elasticloadbalancingv2ListenerRuleConditionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_Elasticloadbalancingv2ListenerRuleConditionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_Elasticloadbalancingv2ListenerRuleConditionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerRuleConditionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_Elasticloadbalancingv2ListenerRuleConditionsOutputReference) PutHostHeaderConfig(value *Elasticloadbalancingv2ListenerRuleConditionsHostHeaderConfig) {
	if err := e.validatePutHostHeaderConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putHostHeaderConfig",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerRuleConditionsOutputReference) PutHttpHeaderConfig(value *Elasticloadbalancingv2ListenerRuleConditionsHttpHeaderConfig) {
	if err := e.validatePutHttpHeaderConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putHttpHeaderConfig",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerRuleConditionsOutputReference) PutHttpRequestMethodConfig(value *Elasticloadbalancingv2ListenerRuleConditionsHttpRequestMethodConfig) {
	if err := e.validatePutHttpRequestMethodConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putHttpRequestMethodConfig",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerRuleConditionsOutputReference) PutPathPatternConfig(value *Elasticloadbalancingv2ListenerRuleConditionsPathPatternConfig) {
	if err := e.validatePutPathPatternConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putPathPatternConfig",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerRuleConditionsOutputReference) PutQueryStringConfig(value *Elasticloadbalancingv2ListenerRuleConditionsQueryStringConfig) {
	if err := e.validatePutQueryStringConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putQueryStringConfig",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerRuleConditionsOutputReference) PutSourceIpConfig(value *Elasticloadbalancingv2ListenerRuleConditionsSourceIpConfig) {
	if err := e.validatePutSourceIpConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putSourceIpConfig",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerRuleConditionsOutputReference) ResetField() {
	_jsii_.InvokeVoid(
		e,
		"resetField",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerRuleConditionsOutputReference) ResetHostHeaderConfig() {
	_jsii_.InvokeVoid(
		e,
		"resetHostHeaderConfig",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerRuleConditionsOutputReference) ResetHttpHeaderConfig() {
	_jsii_.InvokeVoid(
		e,
		"resetHttpHeaderConfig",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerRuleConditionsOutputReference) ResetHttpRequestMethodConfig() {
	_jsii_.InvokeVoid(
		e,
		"resetHttpRequestMethodConfig",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerRuleConditionsOutputReference) ResetPathPatternConfig() {
	_jsii_.InvokeVoid(
		e,
		"resetPathPatternConfig",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerRuleConditionsOutputReference) ResetQueryStringConfig() {
	_jsii_.InvokeVoid(
		e,
		"resetQueryStringConfig",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerRuleConditionsOutputReference) ResetRegexValues() {
	_jsii_.InvokeVoid(
		e,
		"resetRegexValues",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerRuleConditionsOutputReference) ResetSourceIpConfig() {
	_jsii_.InvokeVoid(
		e,
		"resetSourceIpConfig",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerRuleConditionsOutputReference) ResetValues() {
	_jsii_.InvokeVoid(
		e,
		"resetValues",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerRuleConditionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (e *jsiiProxy_Elasticloadbalancingv2ListenerRuleConditionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

