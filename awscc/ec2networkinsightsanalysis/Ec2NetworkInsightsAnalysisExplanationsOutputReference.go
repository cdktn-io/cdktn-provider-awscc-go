// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2networkinsightsanalysis

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/ec2networkinsightsanalysis/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Ec2NetworkInsightsAnalysisExplanationsOutputReference interface {
	cdktn.ComplexObject
	Acl() Ec2NetworkInsightsAnalysisExplanationsAclOutputReference
	AclRule() Ec2NetworkInsightsAnalysisExplanationsAclRuleOutputReference
	Address() *string
	Addresses() *[]*string
	AttachedTo() Ec2NetworkInsightsAnalysisExplanationsAttachedToOutputReference
	AvailabilityZones() *[]*string
	Cidrs() *[]*string
	ClassicLoadBalancerListener() Ec2NetworkInsightsAnalysisExplanationsClassicLoadBalancerListenerOutputReference
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
	Component() Ec2NetworkInsightsAnalysisExplanationsComponentOutputReference
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CustomerGateway() Ec2NetworkInsightsAnalysisExplanationsCustomerGatewayOutputReference
	Destination() Ec2NetworkInsightsAnalysisExplanationsDestinationOutputReference
	DestinationVpc() Ec2NetworkInsightsAnalysisExplanationsDestinationVpcOutputReference
	Direction() *string
	ElasticLoadBalancerListener() Ec2NetworkInsightsAnalysisExplanationsElasticLoadBalancerListenerOutputReference
	ExplanationCode() *string
	// Experimental.
	Fqn() *string
	IngressRouteTable() Ec2NetworkInsightsAnalysisExplanationsIngressRouteTableOutputReference
	InternalValue() *Ec2NetworkInsightsAnalysisExplanations
	SetInternalValue(val *Ec2NetworkInsightsAnalysisExplanations)
	InternetGateway() Ec2NetworkInsightsAnalysisExplanationsInternetGatewayOutputReference
	LoadBalancerArn() *string
	LoadBalancerListenerPort() *float64
	LoadBalancerTarget() Ec2NetworkInsightsAnalysisExplanationsLoadBalancerTargetOutputReference
	LoadBalancerTargetGroup() Ec2NetworkInsightsAnalysisExplanationsLoadBalancerTargetGroupOutputReference
	LoadBalancerTargetGroups() Ec2NetworkInsightsAnalysisExplanationsLoadBalancerTargetGroupsList
	LoadBalancerTargetPort() *float64
	MissingComponent() *string
	NatGateway() Ec2NetworkInsightsAnalysisExplanationsNatGatewayOutputReference
	NetworkInterface() Ec2NetworkInsightsAnalysisExplanationsNetworkInterfaceOutputReference
	PacketField() *string
	Port() *float64
	PortRanges() Ec2NetworkInsightsAnalysisExplanationsPortRangesList
	PrefixList() Ec2NetworkInsightsAnalysisExplanationsPrefixListStructOutputReference
	Protocols() *[]*string
	RouteTable() Ec2NetworkInsightsAnalysisExplanationsRouteTableOutputReference
	RouteTableRoute() Ec2NetworkInsightsAnalysisExplanationsRouteTableRouteOutputReference
	SecurityGroup() Ec2NetworkInsightsAnalysisExplanationsSecurityGroupOutputReference
	SecurityGroupRule() Ec2NetworkInsightsAnalysisExplanationsSecurityGroupRuleOutputReference
	SecurityGroups() Ec2NetworkInsightsAnalysisExplanationsSecurityGroupsList
	SourceVpc() Ec2NetworkInsightsAnalysisExplanationsSourceVpcOutputReference
	State() *string
	Subnet() Ec2NetworkInsightsAnalysisExplanationsSubnetOutputReference
	SubnetRouteTable() Ec2NetworkInsightsAnalysisExplanationsSubnetRouteTableOutputReference
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Vpc() Ec2NetworkInsightsAnalysisExplanationsVpcOutputReference
	VpcEndpoint() Ec2NetworkInsightsAnalysisExplanationsVpcEndpointOutputReference
	VpcPeeringConnection() Ec2NetworkInsightsAnalysisExplanationsVpcPeeringConnectionOutputReference
	VpnConnection() Ec2NetworkInsightsAnalysisExplanationsVpnConnectionOutputReference
	VpnGateway() Ec2NetworkInsightsAnalysisExplanationsVpnGatewayOutputReference
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Ec2NetworkInsightsAnalysisExplanationsOutputReference
type jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) Acl() Ec2NetworkInsightsAnalysisExplanationsAclOutputReference {
	var returns Ec2NetworkInsightsAnalysisExplanationsAclOutputReference
	_jsii_.Get(
		j,
		"acl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) AclRule() Ec2NetworkInsightsAnalysisExplanationsAclRuleOutputReference {
	var returns Ec2NetworkInsightsAnalysisExplanationsAclRuleOutputReference
	_jsii_.Get(
		j,
		"aclRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) Address() *string {
	var returns *string
	_jsii_.Get(
		j,
		"address",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) Addresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"addresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) AttachedTo() Ec2NetworkInsightsAnalysisExplanationsAttachedToOutputReference {
	var returns Ec2NetworkInsightsAnalysisExplanationsAttachedToOutputReference
	_jsii_.Get(
		j,
		"attachedTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) AvailabilityZones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availabilityZones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) Cidrs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cidrs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) ClassicLoadBalancerListener() Ec2NetworkInsightsAnalysisExplanationsClassicLoadBalancerListenerOutputReference {
	var returns Ec2NetworkInsightsAnalysisExplanationsClassicLoadBalancerListenerOutputReference
	_jsii_.Get(
		j,
		"classicLoadBalancerListener",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) Component() Ec2NetworkInsightsAnalysisExplanationsComponentOutputReference {
	var returns Ec2NetworkInsightsAnalysisExplanationsComponentOutputReference
	_jsii_.Get(
		j,
		"component",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) CustomerGateway() Ec2NetworkInsightsAnalysisExplanationsCustomerGatewayOutputReference {
	var returns Ec2NetworkInsightsAnalysisExplanationsCustomerGatewayOutputReference
	_jsii_.Get(
		j,
		"customerGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) Destination() Ec2NetworkInsightsAnalysisExplanationsDestinationOutputReference {
	var returns Ec2NetworkInsightsAnalysisExplanationsDestinationOutputReference
	_jsii_.Get(
		j,
		"destination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) DestinationVpc() Ec2NetworkInsightsAnalysisExplanationsDestinationVpcOutputReference {
	var returns Ec2NetworkInsightsAnalysisExplanationsDestinationVpcOutputReference
	_jsii_.Get(
		j,
		"destinationVpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) Direction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"direction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) ElasticLoadBalancerListener() Ec2NetworkInsightsAnalysisExplanationsElasticLoadBalancerListenerOutputReference {
	var returns Ec2NetworkInsightsAnalysisExplanationsElasticLoadBalancerListenerOutputReference
	_jsii_.Get(
		j,
		"elasticLoadBalancerListener",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) ExplanationCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"explanationCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) IngressRouteTable() Ec2NetworkInsightsAnalysisExplanationsIngressRouteTableOutputReference {
	var returns Ec2NetworkInsightsAnalysisExplanationsIngressRouteTableOutputReference
	_jsii_.Get(
		j,
		"ingressRouteTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) InternalValue() *Ec2NetworkInsightsAnalysisExplanations {
	var returns *Ec2NetworkInsightsAnalysisExplanations
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) InternetGateway() Ec2NetworkInsightsAnalysisExplanationsInternetGatewayOutputReference {
	var returns Ec2NetworkInsightsAnalysisExplanationsInternetGatewayOutputReference
	_jsii_.Get(
		j,
		"internetGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) LoadBalancerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) LoadBalancerListenerPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"loadBalancerListenerPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) LoadBalancerTarget() Ec2NetworkInsightsAnalysisExplanationsLoadBalancerTargetOutputReference {
	var returns Ec2NetworkInsightsAnalysisExplanationsLoadBalancerTargetOutputReference
	_jsii_.Get(
		j,
		"loadBalancerTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) LoadBalancerTargetGroup() Ec2NetworkInsightsAnalysisExplanationsLoadBalancerTargetGroupOutputReference {
	var returns Ec2NetworkInsightsAnalysisExplanationsLoadBalancerTargetGroupOutputReference
	_jsii_.Get(
		j,
		"loadBalancerTargetGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) LoadBalancerTargetGroups() Ec2NetworkInsightsAnalysisExplanationsLoadBalancerTargetGroupsList {
	var returns Ec2NetworkInsightsAnalysisExplanationsLoadBalancerTargetGroupsList
	_jsii_.Get(
		j,
		"loadBalancerTargetGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) LoadBalancerTargetPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"loadBalancerTargetPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) MissingComponent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"missingComponent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) NatGateway() Ec2NetworkInsightsAnalysisExplanationsNatGatewayOutputReference {
	var returns Ec2NetworkInsightsAnalysisExplanationsNatGatewayOutputReference
	_jsii_.Get(
		j,
		"natGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) NetworkInterface() Ec2NetworkInsightsAnalysisExplanationsNetworkInterfaceOutputReference {
	var returns Ec2NetworkInsightsAnalysisExplanationsNetworkInterfaceOutputReference
	_jsii_.Get(
		j,
		"networkInterface",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) PacketField() *string {
	var returns *string
	_jsii_.Get(
		j,
		"packetField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) PortRanges() Ec2NetworkInsightsAnalysisExplanationsPortRangesList {
	var returns Ec2NetworkInsightsAnalysisExplanationsPortRangesList
	_jsii_.Get(
		j,
		"portRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) PrefixList() Ec2NetworkInsightsAnalysisExplanationsPrefixListStructOutputReference {
	var returns Ec2NetworkInsightsAnalysisExplanationsPrefixListStructOutputReference
	_jsii_.Get(
		j,
		"prefixList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) Protocols() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"protocols",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) RouteTable() Ec2NetworkInsightsAnalysisExplanationsRouteTableOutputReference {
	var returns Ec2NetworkInsightsAnalysisExplanationsRouteTableOutputReference
	_jsii_.Get(
		j,
		"routeTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) RouteTableRoute() Ec2NetworkInsightsAnalysisExplanationsRouteTableRouteOutputReference {
	var returns Ec2NetworkInsightsAnalysisExplanationsRouteTableRouteOutputReference
	_jsii_.Get(
		j,
		"routeTableRoute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) SecurityGroup() Ec2NetworkInsightsAnalysisExplanationsSecurityGroupOutputReference {
	var returns Ec2NetworkInsightsAnalysisExplanationsSecurityGroupOutputReference
	_jsii_.Get(
		j,
		"securityGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) SecurityGroupRule() Ec2NetworkInsightsAnalysisExplanationsSecurityGroupRuleOutputReference {
	var returns Ec2NetworkInsightsAnalysisExplanationsSecurityGroupRuleOutputReference
	_jsii_.Get(
		j,
		"securityGroupRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) SecurityGroups() Ec2NetworkInsightsAnalysisExplanationsSecurityGroupsList {
	var returns Ec2NetworkInsightsAnalysisExplanationsSecurityGroupsList
	_jsii_.Get(
		j,
		"securityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) SourceVpc() Ec2NetworkInsightsAnalysisExplanationsSourceVpcOutputReference {
	var returns Ec2NetworkInsightsAnalysisExplanationsSourceVpcOutputReference
	_jsii_.Get(
		j,
		"sourceVpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) Subnet() Ec2NetworkInsightsAnalysisExplanationsSubnetOutputReference {
	var returns Ec2NetworkInsightsAnalysisExplanationsSubnetOutputReference
	_jsii_.Get(
		j,
		"subnet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) SubnetRouteTable() Ec2NetworkInsightsAnalysisExplanationsSubnetRouteTableOutputReference {
	var returns Ec2NetworkInsightsAnalysisExplanationsSubnetRouteTableOutputReference
	_jsii_.Get(
		j,
		"subnetRouteTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) Vpc() Ec2NetworkInsightsAnalysisExplanationsVpcOutputReference {
	var returns Ec2NetworkInsightsAnalysisExplanationsVpcOutputReference
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) VpcEndpoint() Ec2NetworkInsightsAnalysisExplanationsVpcEndpointOutputReference {
	var returns Ec2NetworkInsightsAnalysisExplanationsVpcEndpointOutputReference
	_jsii_.Get(
		j,
		"vpcEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) VpcPeeringConnection() Ec2NetworkInsightsAnalysisExplanationsVpcPeeringConnectionOutputReference {
	var returns Ec2NetworkInsightsAnalysisExplanationsVpcPeeringConnectionOutputReference
	_jsii_.Get(
		j,
		"vpcPeeringConnection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) VpnConnection() Ec2NetworkInsightsAnalysisExplanationsVpnConnectionOutputReference {
	var returns Ec2NetworkInsightsAnalysisExplanationsVpnConnectionOutputReference
	_jsii_.Get(
		j,
		"vpnConnection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) VpnGateway() Ec2NetworkInsightsAnalysisExplanationsVpnGatewayOutputReference {
	var returns Ec2NetworkInsightsAnalysisExplanationsVpnGatewayOutputReference
	_jsii_.Get(
		j,
		"vpnGateway",
		&returns,
	)
	return returns
}


func NewEc2NetworkInsightsAnalysisExplanationsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) Ec2NetworkInsightsAnalysisExplanationsOutputReference {
	_init_.Initialize()

	if err := validateNewEc2NetworkInsightsAnalysisExplanationsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.ec2NetworkInsightsAnalysis.Ec2NetworkInsightsAnalysisExplanationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEc2NetworkInsightsAnalysisExplanationsOutputReference_Override(e Ec2NetworkInsightsAnalysisExplanationsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.ec2NetworkInsightsAnalysis.Ec2NetworkInsightsAnalysisExplanationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference)SetInternalValue(val *Ec2NetworkInsightsAnalysisExplanations) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

