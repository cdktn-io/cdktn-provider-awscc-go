// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ekscluster


type EksClusterResourcesVpcConfig struct {
	// Specify subnets for your Amazon EKS nodes.
	//
	// Amazon EKS creates cross-account elastic network interfaces in these subnets to allow communication between your nodes and the Kubernetes control plane.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/eks_cluster#subnet_ids EksCluster#subnet_ids}
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// Specify the egress mode for the cluster control plane.
	//
	// If you set this to CUSTOMER_ROUTED, the control plane routes traffic through your VPC subnets instead of using AWS managed networking.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/eks_cluster#control_plane_egress_mode EksCluster#control_plane_egress_mode}
	ControlPlaneEgressMode *string `field:"optional" json:"controlPlaneEgressMode" yaml:"controlPlaneEgressMode"`
	// Set this value to true to enable private access for your cluster's Kubernetes API server endpoint.
	//
	// If you enable private access, Kubernetes API requests from within your cluster's VPC use the private VPC endpoint. The default value for this parameter is false, which disables private access for your Kubernetes API server. If you disable private access and you have nodes or AWS Fargate pods in the cluster, then ensure that publicAccessCidrs includes the necessary CIDR blocks for communication with the nodes or Fargate pods.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/eks_cluster#endpoint_private_access EksCluster#endpoint_private_access}
	EndpointPrivateAccess interface{} `field:"optional" json:"endpointPrivateAccess" yaml:"endpointPrivateAccess"`
	// Set this value to false to disable public access to your cluster's Kubernetes API server endpoint.
	//
	// If you disable public access, your cluster's Kubernetes API server can only receive requests from within the cluster VPC. The default value for this parameter is true, which enables public access for your Kubernetes API server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/eks_cluster#endpoint_public_access EksCluster#endpoint_public_access}
	EndpointPublicAccess interface{} `field:"optional" json:"endpointPublicAccess" yaml:"endpointPublicAccess"`
	// The CIDR blocks that are allowed access to your cluster's public Kubernetes API server endpoint.
	//
	// Communication to the endpoint from addresses outside of the CIDR blocks that you specify is denied. The default value is 0.0.0.0/0. If you've disabled private endpoint access and you have nodes or AWS Fargate pods in the cluster, then ensure that you specify the necessary CIDR blocks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/eks_cluster#public_access_cidrs EksCluster#public_access_cidrs}
	PublicAccessCidrs *[]*string `field:"optional" json:"publicAccessCidrs" yaml:"publicAccessCidrs"`
	// Specify one or more security groups for the cross-account elastic network interfaces that Amazon EKS creates to use to allow communication between your worker nodes and the Kubernetes control plane.
	//
	// If you don't specify a security group, the default security group for your VPC is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/eks_cluster#security_group_ids EksCluster#security_group_ids}
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
}

