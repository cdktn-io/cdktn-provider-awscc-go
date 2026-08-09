// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package rtbfabriclinkroutingrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/rtbfabriclinkroutingrule/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type RtbfabricLinkRoutingRuleConditionsOutputReference interface {
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
	HostHeader() *string
	SetHostHeader(val *string)
	HostHeaderInput() *string
	HostHeaderWildcard() *string
	SetHostHeaderWildcard(val *string)
	HostHeaderWildcardInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PathExact() *string
	SetPathExact(val *string)
	PathExactInput() *string
	PathPrefix() *string
	SetPathPrefix(val *string)
	PathPrefixInput() *string
	QueryStringEquals() RtbfabricLinkRoutingRuleConditionsQueryStringEqualsOutputReference
	QueryStringEqualsInput() interface{}
	QueryStringExists() *string
	SetQueryStringExists(val *string)
	QueryStringExistsInput() *string
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
	PutQueryStringEquals(value *RtbfabricLinkRoutingRuleConditionsQueryStringEquals)
	ResetHostHeader()
	ResetHostHeaderWildcard()
	ResetPathExact()
	ResetPathPrefix()
	ResetQueryStringEquals()
	ResetQueryStringExists()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for RtbfabricLinkRoutingRuleConditionsOutputReference
type jsiiProxy_RtbfabricLinkRoutingRuleConditionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_RtbfabricLinkRoutingRuleConditionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricLinkRoutingRuleConditionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricLinkRoutingRuleConditionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricLinkRoutingRuleConditionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricLinkRoutingRuleConditionsOutputReference) HostHeader() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricLinkRoutingRuleConditionsOutputReference) HostHeaderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricLinkRoutingRuleConditionsOutputReference) HostHeaderWildcard() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostHeaderWildcard",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricLinkRoutingRuleConditionsOutputReference) HostHeaderWildcardInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostHeaderWildcardInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricLinkRoutingRuleConditionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricLinkRoutingRuleConditionsOutputReference) PathExact() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathExact",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricLinkRoutingRuleConditionsOutputReference) PathExactInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathExactInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricLinkRoutingRuleConditionsOutputReference) PathPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricLinkRoutingRuleConditionsOutputReference) PathPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricLinkRoutingRuleConditionsOutputReference) QueryStringEquals() RtbfabricLinkRoutingRuleConditionsQueryStringEqualsOutputReference {
	var returns RtbfabricLinkRoutingRuleConditionsQueryStringEqualsOutputReference
	_jsii_.Get(
		j,
		"queryStringEquals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricLinkRoutingRuleConditionsOutputReference) QueryStringEqualsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queryStringEqualsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricLinkRoutingRuleConditionsOutputReference) QueryStringExists() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryStringExists",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricLinkRoutingRuleConditionsOutputReference) QueryStringExistsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryStringExistsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricLinkRoutingRuleConditionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricLinkRoutingRuleConditionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewRtbfabricLinkRoutingRuleConditionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) RtbfabricLinkRoutingRuleConditionsOutputReference {
	_init_.Initialize()

	if err := validateNewRtbfabricLinkRoutingRuleConditionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_RtbfabricLinkRoutingRuleConditionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.rtbfabricLinkRoutingRule.RtbfabricLinkRoutingRuleConditionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewRtbfabricLinkRoutingRuleConditionsOutputReference_Override(r RtbfabricLinkRoutingRuleConditionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.rtbfabricLinkRoutingRule.RtbfabricLinkRoutingRuleConditionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		r,
	)
}

func (j *jsiiProxy_RtbfabricLinkRoutingRuleConditionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_RtbfabricLinkRoutingRuleConditionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_RtbfabricLinkRoutingRuleConditionsOutputReference)SetHostHeader(val *string) {
	if err := j.validateSetHostHeaderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostHeader",
		val,
	)
}

func (j *jsiiProxy_RtbfabricLinkRoutingRuleConditionsOutputReference)SetHostHeaderWildcard(val *string) {
	if err := j.validateSetHostHeaderWildcardParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostHeaderWildcard",
		val,
	)
}

func (j *jsiiProxy_RtbfabricLinkRoutingRuleConditionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_RtbfabricLinkRoutingRuleConditionsOutputReference)SetPathExact(val *string) {
	if err := j.validateSetPathExactParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pathExact",
		val,
	)
}

func (j *jsiiProxy_RtbfabricLinkRoutingRuleConditionsOutputReference)SetPathPrefix(val *string) {
	if err := j.validateSetPathPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pathPrefix",
		val,
	)
}

func (j *jsiiProxy_RtbfabricLinkRoutingRuleConditionsOutputReference)SetQueryStringExists(val *string) {
	if err := j.validateSetQueryStringExistsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryStringExists",
		val,
	)
}

func (j *jsiiProxy_RtbfabricLinkRoutingRuleConditionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_RtbfabricLinkRoutingRuleConditionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (r *jsiiProxy_RtbfabricLinkRoutingRuleConditionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RtbfabricLinkRoutingRuleConditionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RtbfabricLinkRoutingRuleConditionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RtbfabricLinkRoutingRuleConditionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RtbfabricLinkRoutingRuleConditionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RtbfabricLinkRoutingRuleConditionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RtbfabricLinkRoutingRuleConditionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RtbfabricLinkRoutingRuleConditionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RtbfabricLinkRoutingRuleConditionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RtbfabricLinkRoutingRuleConditionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RtbfabricLinkRoutingRuleConditionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RtbfabricLinkRoutingRuleConditionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RtbfabricLinkRoutingRuleConditionsOutputReference) PutQueryStringEquals(value *RtbfabricLinkRoutingRuleConditionsQueryStringEquals) {
	if err := r.validatePutQueryStringEqualsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putQueryStringEquals",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RtbfabricLinkRoutingRuleConditionsOutputReference) ResetHostHeader() {
	_jsii_.InvokeVoid(
		r,
		"resetHostHeader",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RtbfabricLinkRoutingRuleConditionsOutputReference) ResetHostHeaderWildcard() {
	_jsii_.InvokeVoid(
		r,
		"resetHostHeaderWildcard",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RtbfabricLinkRoutingRuleConditionsOutputReference) ResetPathExact() {
	_jsii_.InvokeVoid(
		r,
		"resetPathExact",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RtbfabricLinkRoutingRuleConditionsOutputReference) ResetPathPrefix() {
	_jsii_.InvokeVoid(
		r,
		"resetPathPrefix",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RtbfabricLinkRoutingRuleConditionsOutputReference) ResetQueryStringEquals() {
	_jsii_.InvokeVoid(
		r,
		"resetQueryStringEquals",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RtbfabricLinkRoutingRuleConditionsOutputReference) ResetQueryStringExists() {
	_jsii_.InvokeVoid(
		r,
		"resetQueryStringExists",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RtbfabricLinkRoutingRuleConditionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := r.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		r,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RtbfabricLinkRoutingRuleConditionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

