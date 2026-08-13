// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lambdanetworkconnector


type LambdaNetworkConnectorConfigurationVpcEgressConfiguration struct {
	// The types of Lambda compute resources that can use this connector. Currently, only MicroVm is supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/lambda_network_connector#associated_compute_resource_types LambdaNetworkConnector#associated_compute_resource_types}
	AssociatedComputeResourceTypes *[]*string `field:"required" json:"associatedComputeResourceTypes" yaml:"associatedComputeResourceTypes"`
	// The IDs of the VPC subnets where Lambda provisions elastic network interfaces (ENIs).
	//
	// Specify 1 to 16 subnets. All subnets must be in the same VPC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/lambda_network_connector#subnet_ids LambdaNetworkConnector#subnet_ids}
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// The network protocol for the connector. Specify IPv4 for IPv4-only networking, or DualStack for both IPv4 and IPv6.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/lambda_network_connector#network_protocol LambdaNetworkConnector#network_protocol}
	NetworkProtocol *string `field:"optional" json:"networkProtocol" yaml:"networkProtocol"`
	// The IDs of the VPC security groups to attach to the ENIs.
	//
	// Specify 0 to 5 security groups. All security groups must be in the same VPC as the subnets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/lambda_network_connector#security_group_ids LambdaNetworkConnector#security_group_ids}
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
}

