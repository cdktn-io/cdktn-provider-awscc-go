// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityagentprivateconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/securityagentprivateconnection/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SecurityagentPrivateConnectionModeServiceManagedOutputReference interface {
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
	DnsResolution() *string
	SetDnsResolution(val *string)
	DnsResolutionInput() *string
	// Experimental.
	Fqn() *string
	HostAddress() *string
	SetHostAddress(val *string)
	HostAddressInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IpAddressType() *string
	SetIpAddressType(val *string)
	IpAddressTypeInput() *string
	Ipv4AddressesPerEni() *float64
	SetIpv4AddressesPerEni(val *float64)
	Ipv4AddressesPerEniInput() *float64
	PortRanges() *[]*string
	SetPortRanges(val *[]*string)
	PortRangesInput() *[]*string
	SecurityGroupIds() *[]*string
	SetSecurityGroupIds(val *[]*string)
	SecurityGroupIdsInput() *[]*string
	SubnetIds() *[]*string
	SetSubnetIds(val *[]*string)
	SubnetIdsInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	VpcId() *string
	SetVpcId(val *string)
	VpcIdInput() *string
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
	ResetDnsResolution()
	ResetHostAddress()
	ResetIpAddressType()
	ResetIpv4AddressesPerEni()
	ResetPortRanges()
	ResetSecurityGroupIds()
	ResetSubnetIds()
	ResetVpcId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SecurityagentPrivateConnectionModeServiceManagedOutputReference
type jsiiProxy_SecurityagentPrivateConnectionModeServiceManagedOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_SecurityagentPrivateConnectionModeServiceManagedOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityagentPrivateConnectionModeServiceManagedOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityagentPrivateConnectionModeServiceManagedOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityagentPrivateConnectionModeServiceManagedOutputReference) DnsResolution() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsResolution",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityagentPrivateConnectionModeServiceManagedOutputReference) DnsResolutionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsResolutionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityagentPrivateConnectionModeServiceManagedOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityagentPrivateConnectionModeServiceManagedOutputReference) HostAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityagentPrivateConnectionModeServiceManagedOutputReference) HostAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityagentPrivateConnectionModeServiceManagedOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityagentPrivateConnectionModeServiceManagedOutputReference) IpAddressType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityagentPrivateConnectionModeServiceManagedOutputReference) IpAddressTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityagentPrivateConnectionModeServiceManagedOutputReference) Ipv4AddressesPerEni() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv4AddressesPerEni",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityagentPrivateConnectionModeServiceManagedOutputReference) Ipv4AddressesPerEniInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv4AddressesPerEniInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityagentPrivateConnectionModeServiceManagedOutputReference) PortRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"portRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityagentPrivateConnectionModeServiceManagedOutputReference) PortRangesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"portRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityagentPrivateConnectionModeServiceManagedOutputReference) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityagentPrivateConnectionModeServiceManagedOutputReference) SecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityagentPrivateConnectionModeServiceManagedOutputReference) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityagentPrivateConnectionModeServiceManagedOutputReference) SubnetIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityagentPrivateConnectionModeServiceManagedOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityagentPrivateConnectionModeServiceManagedOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityagentPrivateConnectionModeServiceManagedOutputReference) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityagentPrivateConnectionModeServiceManagedOutputReference) VpcIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcIdInput",
		&returns,
	)
	return returns
}


func NewSecurityagentPrivateConnectionModeServiceManagedOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) SecurityagentPrivateConnectionModeServiceManagedOutputReference {
	_init_.Initialize()

	if err := validateNewSecurityagentPrivateConnectionModeServiceManagedOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SecurityagentPrivateConnectionModeServiceManagedOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.securityagentPrivateConnection.SecurityagentPrivateConnectionModeServiceManagedOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSecurityagentPrivateConnectionModeServiceManagedOutputReference_Override(s SecurityagentPrivateConnectionModeServiceManagedOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.securityagentPrivateConnection.SecurityagentPrivateConnectionModeServiceManagedOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SecurityagentPrivateConnectionModeServiceManagedOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SecurityagentPrivateConnectionModeServiceManagedOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SecurityagentPrivateConnectionModeServiceManagedOutputReference)SetDnsResolution(val *string) {
	if err := j.validateSetDnsResolutionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsResolution",
		val,
	)
}

func (j *jsiiProxy_SecurityagentPrivateConnectionModeServiceManagedOutputReference)SetHostAddress(val *string) {
	if err := j.validateSetHostAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostAddress",
		val,
	)
}

func (j *jsiiProxy_SecurityagentPrivateConnectionModeServiceManagedOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SecurityagentPrivateConnectionModeServiceManagedOutputReference)SetIpAddressType(val *string) {
	if err := j.validateSetIpAddressTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipAddressType",
		val,
	)
}

func (j *jsiiProxy_SecurityagentPrivateConnectionModeServiceManagedOutputReference)SetIpv4AddressesPerEni(val *float64) {
	if err := j.validateSetIpv4AddressesPerEniParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv4AddressesPerEni",
		val,
	)
}

func (j *jsiiProxy_SecurityagentPrivateConnectionModeServiceManagedOutputReference)SetPortRanges(val *[]*string) {
	if err := j.validateSetPortRangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"portRanges",
		val,
	)
}

func (j *jsiiProxy_SecurityagentPrivateConnectionModeServiceManagedOutputReference)SetSecurityGroupIds(val *[]*string) {
	if err := j.validateSetSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_SecurityagentPrivateConnectionModeServiceManagedOutputReference)SetSubnetIds(val *[]*string) {
	if err := j.validateSetSubnetIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetIds",
		val,
	)
}

func (j *jsiiProxy_SecurityagentPrivateConnectionModeServiceManagedOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SecurityagentPrivateConnectionModeServiceManagedOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SecurityagentPrivateConnectionModeServiceManagedOutputReference)SetVpcId(val *string) {
	if err := j.validateSetVpcIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcId",
		val,
	)
}

func (s *jsiiProxy_SecurityagentPrivateConnectionModeServiceManagedOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityagentPrivateConnectionModeServiceManagedOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityagentPrivateConnectionModeServiceManagedOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityagentPrivateConnectionModeServiceManagedOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityagentPrivateConnectionModeServiceManagedOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityagentPrivateConnectionModeServiceManagedOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityagentPrivateConnectionModeServiceManagedOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityagentPrivateConnectionModeServiceManagedOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityagentPrivateConnectionModeServiceManagedOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityagentPrivateConnectionModeServiceManagedOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityagentPrivateConnectionModeServiceManagedOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityagentPrivateConnectionModeServiceManagedOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityagentPrivateConnectionModeServiceManagedOutputReference) ResetDnsResolution() {
	_jsii_.InvokeVoid(
		s,
		"resetDnsResolution",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityagentPrivateConnectionModeServiceManagedOutputReference) ResetHostAddress() {
	_jsii_.InvokeVoid(
		s,
		"resetHostAddress",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityagentPrivateConnectionModeServiceManagedOutputReference) ResetIpAddressType() {
	_jsii_.InvokeVoid(
		s,
		"resetIpAddressType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityagentPrivateConnectionModeServiceManagedOutputReference) ResetIpv4AddressesPerEni() {
	_jsii_.InvokeVoid(
		s,
		"resetIpv4AddressesPerEni",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityagentPrivateConnectionModeServiceManagedOutputReference) ResetPortRanges() {
	_jsii_.InvokeVoid(
		s,
		"resetPortRanges",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityagentPrivateConnectionModeServiceManagedOutputReference) ResetSecurityGroupIds() {
	_jsii_.InvokeVoid(
		s,
		"resetSecurityGroupIds",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityagentPrivateConnectionModeServiceManagedOutputReference) ResetSubnetIds() {
	_jsii_.InvokeVoid(
		s,
		"resetSubnetIds",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityagentPrivateConnectionModeServiceManagedOutputReference) ResetVpcId() {
	_jsii_.InvokeVoid(
		s,
		"resetVpcId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityagentPrivateConnectionModeServiceManagedOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := s.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityagentPrivateConnectionModeServiceManagedOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

