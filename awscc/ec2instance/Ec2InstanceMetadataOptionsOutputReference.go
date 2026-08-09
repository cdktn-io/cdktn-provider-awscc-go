// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2instance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/ec2instance/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Ec2InstanceMetadataOptionsOutputReference interface {
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
	HttpEndpoint() *string
	SetHttpEndpoint(val *string)
	HttpEndpointInput() *string
	HttpProtocolIpv6() *string
	SetHttpProtocolIpv6(val *string)
	HttpProtocolIpv6Input() *string
	HttpPutResponseHopLimit() *float64
	SetHttpPutResponseHopLimit(val *float64)
	HttpPutResponseHopLimitInput() *float64
	HttpTokens() *string
	SetHttpTokens(val *string)
	HttpTokensInput() *string
	InstanceMetadataTags() *string
	SetInstanceMetadataTags(val *string)
	InstanceMetadataTagsInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
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
	ResetHttpEndpoint()
	ResetHttpProtocolIpv6()
	ResetHttpPutResponseHopLimit()
	ResetHttpTokens()
	ResetInstanceMetadataTags()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Ec2InstanceMetadataOptionsOutputReference
type jsiiProxy_Ec2InstanceMetadataOptionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_Ec2InstanceMetadataOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstanceMetadataOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstanceMetadataOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstanceMetadataOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstanceMetadataOptionsOutputReference) HttpEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstanceMetadataOptionsOutputReference) HttpEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstanceMetadataOptionsOutputReference) HttpProtocolIpv6() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpProtocolIpv6",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstanceMetadataOptionsOutputReference) HttpProtocolIpv6Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpProtocolIpv6Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstanceMetadataOptionsOutputReference) HttpPutResponseHopLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpPutResponseHopLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstanceMetadataOptionsOutputReference) HttpPutResponseHopLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpPutResponseHopLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstanceMetadataOptionsOutputReference) HttpTokens() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpTokens",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstanceMetadataOptionsOutputReference) HttpTokensInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpTokensInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstanceMetadataOptionsOutputReference) InstanceMetadataTags() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceMetadataTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstanceMetadataOptionsOutputReference) InstanceMetadataTagsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceMetadataTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstanceMetadataOptionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstanceMetadataOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstanceMetadataOptionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEc2InstanceMetadataOptionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) Ec2InstanceMetadataOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewEc2InstanceMetadataOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Ec2InstanceMetadataOptionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.ec2Instance.Ec2InstanceMetadataOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEc2InstanceMetadataOptionsOutputReference_Override(e Ec2InstanceMetadataOptionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.ec2Instance.Ec2InstanceMetadataOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_Ec2InstanceMetadataOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Ec2InstanceMetadataOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Ec2InstanceMetadataOptionsOutputReference)SetHttpEndpoint(val *string) {
	if err := j.validateSetHttpEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpEndpoint",
		val,
	)
}

func (j *jsiiProxy_Ec2InstanceMetadataOptionsOutputReference)SetHttpProtocolIpv6(val *string) {
	if err := j.validateSetHttpProtocolIpv6Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpProtocolIpv6",
		val,
	)
}

func (j *jsiiProxy_Ec2InstanceMetadataOptionsOutputReference)SetHttpPutResponseHopLimit(val *float64) {
	if err := j.validateSetHttpPutResponseHopLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpPutResponseHopLimit",
		val,
	)
}

func (j *jsiiProxy_Ec2InstanceMetadataOptionsOutputReference)SetHttpTokens(val *string) {
	if err := j.validateSetHttpTokensParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpTokens",
		val,
	)
}

func (j *jsiiProxy_Ec2InstanceMetadataOptionsOutputReference)SetInstanceMetadataTags(val *string) {
	if err := j.validateSetInstanceMetadataTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceMetadataTags",
		val,
	)
}

func (j *jsiiProxy_Ec2InstanceMetadataOptionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Ec2InstanceMetadataOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Ec2InstanceMetadataOptionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_Ec2InstanceMetadataOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2InstanceMetadataOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_Ec2InstanceMetadataOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_Ec2InstanceMetadataOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_Ec2InstanceMetadataOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_Ec2InstanceMetadataOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_Ec2InstanceMetadataOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_Ec2InstanceMetadataOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_Ec2InstanceMetadataOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_Ec2InstanceMetadataOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_Ec2InstanceMetadataOptionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2InstanceMetadataOptionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_Ec2InstanceMetadataOptionsOutputReference) ResetHttpEndpoint() {
	_jsii_.InvokeVoid(
		e,
		"resetHttpEndpoint",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2InstanceMetadataOptionsOutputReference) ResetHttpProtocolIpv6() {
	_jsii_.InvokeVoid(
		e,
		"resetHttpProtocolIpv6",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2InstanceMetadataOptionsOutputReference) ResetHttpPutResponseHopLimit() {
	_jsii_.InvokeVoid(
		e,
		"resetHttpPutResponseHopLimit",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2InstanceMetadataOptionsOutputReference) ResetHttpTokens() {
	_jsii_.InvokeVoid(
		e,
		"resetHttpTokens",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2InstanceMetadataOptionsOutputReference) ResetInstanceMetadataTags() {
	_jsii_.InvokeVoid(
		e,
		"resetInstanceMetadataTags",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2InstanceMetadataOptionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (e *jsiiProxy_Ec2InstanceMetadataOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

