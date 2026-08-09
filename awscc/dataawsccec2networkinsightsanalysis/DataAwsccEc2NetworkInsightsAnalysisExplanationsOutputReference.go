// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccec2networkinsightsanalysis

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawsccec2networkinsightsanalysis/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference interface {
	cdktn.ComplexObject
	Acl() DataAwsccEc2NetworkInsightsAnalysisExplanationsAclOutputReference
	AclRule() DataAwsccEc2NetworkInsightsAnalysisExplanationsAclRuleOutputReference
	Address() *string
	Addresses() *[]*string
	AttachedTo() DataAwsccEc2NetworkInsightsAnalysisExplanationsAttachedToOutputReference
	AvailabilityZones() *[]*string
	Cidrs() *[]*string
	ClassicLoadBalancerListener() DataAwsccEc2NetworkInsightsAnalysisExplanationsClassicLoadBalancerListenerOutputReference
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
	Component() DataAwsccEc2NetworkInsightsAnalysisExplanationsComponentOutputReference
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CustomerGateway() DataAwsccEc2NetworkInsightsAnalysisExplanationsCustomerGatewayOutputReference
	Destination() DataAwsccEc2NetworkInsightsAnalysisExplanationsDestinationOutputReference
	DestinationVpc() DataAwsccEc2NetworkInsightsAnalysisExplanationsDestinationVpcOutputReference
	Direction() *string
	ElasticLoadBalancerListener() DataAwsccEc2NetworkInsightsAnalysisExplanationsElasticLoadBalancerListenerOutputReference
	ExplanationCode() *string
	// Experimental.
	Fqn() *string
	IngressRouteTable() DataAwsccEc2NetworkInsightsAnalysisExplanationsIngressRouteTableOutputReference
	InternalValue() *DataAwsccEc2NetworkInsightsAnalysisExplanations
	SetInternalValue(val *DataAwsccEc2NetworkInsightsAnalysisExplanations)
	InternetGateway() DataAwsccEc2NetworkInsightsAnalysisExplanationsInternetGatewayOutputReference
	LoadBalancerArn() *string
	LoadBalancerListenerPort() *float64
	LoadBalancerTarget() DataAwsccEc2NetworkInsightsAnalysisExplanationsLoadBalancerTargetOutputReference
	LoadBalancerTargetGroup() DataAwsccEc2NetworkInsightsAnalysisExplanationsLoadBalancerTargetGroupOutputReference
	LoadBalancerTargetGroups() DataAwsccEc2NetworkInsightsAnalysisExplanationsLoadBalancerTargetGroupsList
	LoadBalancerTargetPort() *float64
	MissingComponent() *string
	NatGateway() DataAwsccEc2NetworkInsightsAnalysisExplanationsNatGatewayOutputReference
	NetworkInterface() DataAwsccEc2NetworkInsightsAnalysisExplanationsNetworkInterfaceOutputReference
	PacketField() *string
	Port() *float64
	PortRanges() DataAwsccEc2NetworkInsightsAnalysisExplanationsPortRangesList
	PrefixList() DataAwsccEc2NetworkInsightsAnalysisExplanationsPrefixListStructOutputReference
	Protocols() *[]*string
	RouteTable() DataAwsccEc2NetworkInsightsAnalysisExplanationsRouteTableOutputReference
	RouteTableRoute() DataAwsccEc2NetworkInsightsAnalysisExplanationsRouteTableRouteOutputReference
	SecurityGroup() DataAwsccEc2NetworkInsightsAnalysisExplanationsSecurityGroupOutputReference
	SecurityGroupRule() DataAwsccEc2NetworkInsightsAnalysisExplanationsSecurityGroupRuleOutputReference
	SecurityGroups() DataAwsccEc2NetworkInsightsAnalysisExplanationsSecurityGroupsList
	SourceVpc() DataAwsccEc2NetworkInsightsAnalysisExplanationsSourceVpcOutputReference
	State() *string
	Subnet() DataAwsccEc2NetworkInsightsAnalysisExplanationsSubnetOutputReference
	SubnetRouteTable() DataAwsccEc2NetworkInsightsAnalysisExplanationsSubnetRouteTableOutputReference
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Vpc() DataAwsccEc2NetworkInsightsAnalysisExplanationsVpcOutputReference
	VpcEndpoint() DataAwsccEc2NetworkInsightsAnalysisExplanationsVpcEndpointOutputReference
	VpcPeeringConnection() DataAwsccEc2NetworkInsightsAnalysisExplanationsVpcPeeringConnectionOutputReference
	VpnConnection() DataAwsccEc2NetworkInsightsAnalysisExplanationsVpnConnectionOutputReference
	VpnGateway() DataAwsccEc2NetworkInsightsAnalysisExplanationsVpnGatewayOutputReference
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

// The jsii proxy struct for DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference
type jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference) Acl() DataAwsccEc2NetworkInsightsAnalysisExplanationsAclOutputReference {
	var returns DataAwsccEc2NetworkInsightsAnalysisExplanationsAclOutputReference
	_jsii_.Get(
		j,
		"acl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference) AclRule() DataAwsccEc2NetworkInsightsAnalysisExplanationsAclRuleOutputReference {
	var returns DataAwsccEc2NetworkInsightsAnalysisExplanationsAclRuleOutputReference
	_jsii_.Get(
		j,
		"aclRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference) Address() *string {
	var returns *string
	_jsii_.Get(
		j,
		"address",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference) Addresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"addresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference) AttachedTo() DataAwsccEc2NetworkInsightsAnalysisExplanationsAttachedToOutputReference {
	var returns DataAwsccEc2NetworkInsightsAnalysisExplanationsAttachedToOutputReference
	_jsii_.Get(
		j,
		"attachedTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference) AvailabilityZones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availabilityZones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference) Cidrs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cidrs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference) ClassicLoadBalancerListener() DataAwsccEc2NetworkInsightsAnalysisExplanationsClassicLoadBalancerListenerOutputReference {
	var returns DataAwsccEc2NetworkInsightsAnalysisExplanationsClassicLoadBalancerListenerOutputReference
	_jsii_.Get(
		j,
		"classicLoadBalancerListener",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference) Component() DataAwsccEc2NetworkInsightsAnalysisExplanationsComponentOutputReference {
	var returns DataAwsccEc2NetworkInsightsAnalysisExplanationsComponentOutputReference
	_jsii_.Get(
		j,
		"component",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference) CustomerGateway() DataAwsccEc2NetworkInsightsAnalysisExplanationsCustomerGatewayOutputReference {
	var returns DataAwsccEc2NetworkInsightsAnalysisExplanationsCustomerGatewayOutputReference
	_jsii_.Get(
		j,
		"customerGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference) Destination() DataAwsccEc2NetworkInsightsAnalysisExplanationsDestinationOutputReference {
	var returns DataAwsccEc2NetworkInsightsAnalysisExplanationsDestinationOutputReference
	_jsii_.Get(
		j,
		"destination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference) DestinationVpc() DataAwsccEc2NetworkInsightsAnalysisExplanationsDestinationVpcOutputReference {
	var returns DataAwsccEc2NetworkInsightsAnalysisExplanationsDestinationVpcOutputReference
	_jsii_.Get(
		j,
		"destinationVpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference) Direction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"direction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference) ElasticLoadBalancerListener() DataAwsccEc2NetworkInsightsAnalysisExplanationsElasticLoadBalancerListenerOutputReference {
	var returns DataAwsccEc2NetworkInsightsAnalysisExplanationsElasticLoadBalancerListenerOutputReference
	_jsii_.Get(
		j,
		"elasticLoadBalancerListener",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference) ExplanationCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"explanationCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference) IngressRouteTable() DataAwsccEc2NetworkInsightsAnalysisExplanationsIngressRouteTableOutputReference {
	var returns DataAwsccEc2NetworkInsightsAnalysisExplanationsIngressRouteTableOutputReference
	_jsii_.Get(
		j,
		"ingressRouteTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference) InternalValue() *DataAwsccEc2NetworkInsightsAnalysisExplanations {
	var returns *DataAwsccEc2NetworkInsightsAnalysisExplanations
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference) InternetGateway() DataAwsccEc2NetworkInsightsAnalysisExplanationsInternetGatewayOutputReference {
	var returns DataAwsccEc2NetworkInsightsAnalysisExplanationsInternetGatewayOutputReference
	_jsii_.Get(
		j,
		"internetGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference) LoadBalancerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference) LoadBalancerListenerPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"loadBalancerListenerPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference) LoadBalancerTarget() DataAwsccEc2NetworkInsightsAnalysisExplanationsLoadBalancerTargetOutputReference {
	var returns DataAwsccEc2NetworkInsightsAnalysisExplanationsLoadBalancerTargetOutputReference
	_jsii_.Get(
		j,
		"loadBalancerTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference) LoadBalancerTargetGroup() DataAwsccEc2NetworkInsightsAnalysisExplanationsLoadBalancerTargetGroupOutputReference {
	var returns DataAwsccEc2NetworkInsightsAnalysisExplanationsLoadBalancerTargetGroupOutputReference
	_jsii_.Get(
		j,
		"loadBalancerTargetGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference) LoadBalancerTargetGroups() DataAwsccEc2NetworkInsightsAnalysisExplanationsLoadBalancerTargetGroupsList {
	var returns DataAwsccEc2NetworkInsightsAnalysisExplanationsLoadBalancerTargetGroupsList
	_jsii_.Get(
		j,
		"loadBalancerTargetGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference) LoadBalancerTargetPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"loadBalancerTargetPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference) MissingComponent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"missingComponent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference) NatGateway() DataAwsccEc2NetworkInsightsAnalysisExplanationsNatGatewayOutputReference {
	var returns DataAwsccEc2NetworkInsightsAnalysisExplanationsNatGatewayOutputReference
	_jsii_.Get(
		j,
		"natGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference) NetworkInterface() DataAwsccEc2NetworkInsightsAnalysisExplanationsNetworkInterfaceOutputReference {
	var returns DataAwsccEc2NetworkInsightsAnalysisExplanationsNetworkInterfaceOutputReference
	_jsii_.Get(
		j,
		"networkInterface",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference) PacketField() *string {
	var returns *string
	_jsii_.Get(
		j,
		"packetField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference) PortRanges() DataAwsccEc2NetworkInsightsAnalysisExplanationsPortRangesList {
	var returns DataAwsccEc2NetworkInsightsAnalysisExplanationsPortRangesList
	_jsii_.Get(
		j,
		"portRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference) PrefixList() DataAwsccEc2NetworkInsightsAnalysisExplanationsPrefixListStructOutputReference {
	var returns DataAwsccEc2NetworkInsightsAnalysisExplanationsPrefixListStructOutputReference
	_jsii_.Get(
		j,
		"prefixList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference) Protocols() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"protocols",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference) RouteTable() DataAwsccEc2NetworkInsightsAnalysisExplanationsRouteTableOutputReference {
	var returns DataAwsccEc2NetworkInsightsAnalysisExplanationsRouteTableOutputReference
	_jsii_.Get(
		j,
		"routeTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference) RouteTableRoute() DataAwsccEc2NetworkInsightsAnalysisExplanationsRouteTableRouteOutputReference {
	var returns DataAwsccEc2NetworkInsightsAnalysisExplanationsRouteTableRouteOutputReference
	_jsii_.Get(
		j,
		"routeTableRoute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference) SecurityGroup() DataAwsccEc2NetworkInsightsAnalysisExplanationsSecurityGroupOutputReference {
	var returns DataAwsccEc2NetworkInsightsAnalysisExplanationsSecurityGroupOutputReference
	_jsii_.Get(
		j,
		"securityGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference) SecurityGroupRule() DataAwsccEc2NetworkInsightsAnalysisExplanationsSecurityGroupRuleOutputReference {
	var returns DataAwsccEc2NetworkInsightsAnalysisExplanationsSecurityGroupRuleOutputReference
	_jsii_.Get(
		j,
		"securityGroupRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference) SecurityGroups() DataAwsccEc2NetworkInsightsAnalysisExplanationsSecurityGroupsList {
	var returns DataAwsccEc2NetworkInsightsAnalysisExplanationsSecurityGroupsList
	_jsii_.Get(
		j,
		"securityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference) SourceVpc() DataAwsccEc2NetworkInsightsAnalysisExplanationsSourceVpcOutputReference {
	var returns DataAwsccEc2NetworkInsightsAnalysisExplanationsSourceVpcOutputReference
	_jsii_.Get(
		j,
		"sourceVpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference) Subnet() DataAwsccEc2NetworkInsightsAnalysisExplanationsSubnetOutputReference {
	var returns DataAwsccEc2NetworkInsightsAnalysisExplanationsSubnetOutputReference
	_jsii_.Get(
		j,
		"subnet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference) SubnetRouteTable() DataAwsccEc2NetworkInsightsAnalysisExplanationsSubnetRouteTableOutputReference {
	var returns DataAwsccEc2NetworkInsightsAnalysisExplanationsSubnetRouteTableOutputReference
	_jsii_.Get(
		j,
		"subnetRouteTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference) Vpc() DataAwsccEc2NetworkInsightsAnalysisExplanationsVpcOutputReference {
	var returns DataAwsccEc2NetworkInsightsAnalysisExplanationsVpcOutputReference
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference) VpcEndpoint() DataAwsccEc2NetworkInsightsAnalysisExplanationsVpcEndpointOutputReference {
	var returns DataAwsccEc2NetworkInsightsAnalysisExplanationsVpcEndpointOutputReference
	_jsii_.Get(
		j,
		"vpcEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference) VpcPeeringConnection() DataAwsccEc2NetworkInsightsAnalysisExplanationsVpcPeeringConnectionOutputReference {
	var returns DataAwsccEc2NetworkInsightsAnalysisExplanationsVpcPeeringConnectionOutputReference
	_jsii_.Get(
		j,
		"vpcPeeringConnection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference) VpnConnection() DataAwsccEc2NetworkInsightsAnalysisExplanationsVpnConnectionOutputReference {
	var returns DataAwsccEc2NetworkInsightsAnalysisExplanationsVpnConnectionOutputReference
	_jsii_.Get(
		j,
		"vpnConnection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference) VpnGateway() DataAwsccEc2NetworkInsightsAnalysisExplanationsVpnGatewayOutputReference {
	var returns DataAwsccEc2NetworkInsightsAnalysisExplanationsVpnGatewayOutputReference
	_jsii_.Get(
		j,
		"vpnGateway",
		&returns,
	)
	return returns
}


func NewDataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccEc2NetworkInsightsAnalysis.DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference_Override(d DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccEc2NetworkInsightsAnalysis.DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference)SetInternalValue(val *DataAwsccEc2NetworkInsightsAnalysisExplanations) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisExplanationsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

