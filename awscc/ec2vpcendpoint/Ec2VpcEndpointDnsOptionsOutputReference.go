// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2vpcendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/ec2vpcendpoint/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Ec2VpcEndpointDnsOptionsOutputReference interface {
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
	DnsRecordIpType() *string
	SetDnsRecordIpType(val *string)
	DnsRecordIpTypeInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PrivateDnsOnlyForInboundResolverEndpoint() *string
	SetPrivateDnsOnlyForInboundResolverEndpoint(val *string)
	PrivateDnsOnlyForInboundResolverEndpointInput() *string
	PrivateDnsPreference() *string
	SetPrivateDnsPreference(val *string)
	PrivateDnsPreferenceInput() *string
	PrivateDnsSpecifiedDomains() *[]*string
	SetPrivateDnsSpecifiedDomains(val *[]*string)
	PrivateDnsSpecifiedDomainsInput() *[]*string
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
	ResetDnsRecordIpType()
	ResetPrivateDnsOnlyForInboundResolverEndpoint()
	ResetPrivateDnsPreference()
	ResetPrivateDnsSpecifiedDomains()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Ec2VpcEndpointDnsOptionsOutputReference
type jsiiProxy_Ec2VpcEndpointDnsOptionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_Ec2VpcEndpointDnsOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcEndpointDnsOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcEndpointDnsOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcEndpointDnsOptionsOutputReference) DnsRecordIpType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsRecordIpType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcEndpointDnsOptionsOutputReference) DnsRecordIpTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsRecordIpTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcEndpointDnsOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcEndpointDnsOptionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcEndpointDnsOptionsOutputReference) PrivateDnsOnlyForInboundResolverEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateDnsOnlyForInboundResolverEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcEndpointDnsOptionsOutputReference) PrivateDnsOnlyForInboundResolverEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateDnsOnlyForInboundResolverEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcEndpointDnsOptionsOutputReference) PrivateDnsPreference() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateDnsPreference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcEndpointDnsOptionsOutputReference) PrivateDnsPreferenceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateDnsPreferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcEndpointDnsOptionsOutputReference) PrivateDnsSpecifiedDomains() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"privateDnsSpecifiedDomains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcEndpointDnsOptionsOutputReference) PrivateDnsSpecifiedDomainsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"privateDnsSpecifiedDomainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcEndpointDnsOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcEndpointDnsOptionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEc2VpcEndpointDnsOptionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) Ec2VpcEndpointDnsOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewEc2VpcEndpointDnsOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Ec2VpcEndpointDnsOptionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.ec2VpcEndpoint.Ec2VpcEndpointDnsOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEc2VpcEndpointDnsOptionsOutputReference_Override(e Ec2VpcEndpointDnsOptionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.ec2VpcEndpoint.Ec2VpcEndpointDnsOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_Ec2VpcEndpointDnsOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Ec2VpcEndpointDnsOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Ec2VpcEndpointDnsOptionsOutputReference)SetDnsRecordIpType(val *string) {
	if err := j.validateSetDnsRecordIpTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsRecordIpType",
		val,
	)
}

func (j *jsiiProxy_Ec2VpcEndpointDnsOptionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Ec2VpcEndpointDnsOptionsOutputReference)SetPrivateDnsOnlyForInboundResolverEndpoint(val *string) {
	if err := j.validateSetPrivateDnsOnlyForInboundResolverEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateDnsOnlyForInboundResolverEndpoint",
		val,
	)
}

func (j *jsiiProxy_Ec2VpcEndpointDnsOptionsOutputReference)SetPrivateDnsPreference(val *string) {
	if err := j.validateSetPrivateDnsPreferenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateDnsPreference",
		val,
	)
}

func (j *jsiiProxy_Ec2VpcEndpointDnsOptionsOutputReference)SetPrivateDnsSpecifiedDomains(val *[]*string) {
	if err := j.validateSetPrivateDnsSpecifiedDomainsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateDnsSpecifiedDomains",
		val,
	)
}

func (j *jsiiProxy_Ec2VpcEndpointDnsOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Ec2VpcEndpointDnsOptionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_Ec2VpcEndpointDnsOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2VpcEndpointDnsOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_Ec2VpcEndpointDnsOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_Ec2VpcEndpointDnsOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_Ec2VpcEndpointDnsOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_Ec2VpcEndpointDnsOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_Ec2VpcEndpointDnsOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_Ec2VpcEndpointDnsOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_Ec2VpcEndpointDnsOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_Ec2VpcEndpointDnsOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_Ec2VpcEndpointDnsOptionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2VpcEndpointDnsOptionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_Ec2VpcEndpointDnsOptionsOutputReference) ResetDnsRecordIpType() {
	_jsii_.InvokeVoid(
		e,
		"resetDnsRecordIpType",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2VpcEndpointDnsOptionsOutputReference) ResetPrivateDnsOnlyForInboundResolverEndpoint() {
	_jsii_.InvokeVoid(
		e,
		"resetPrivateDnsOnlyForInboundResolverEndpoint",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2VpcEndpointDnsOptionsOutputReference) ResetPrivateDnsPreference() {
	_jsii_.InvokeVoid(
		e,
		"resetPrivateDnsPreference",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2VpcEndpointDnsOptionsOutputReference) ResetPrivateDnsSpecifiedDomains() {
	_jsii_.InvokeVoid(
		e,
		"resetPrivateDnsSpecifiedDomains",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2VpcEndpointDnsOptionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (e *jsiiProxy_Ec2VpcEndpointDnsOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

