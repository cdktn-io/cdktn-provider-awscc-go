// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ekscluster


type EksClusterOutpostConfig struct {
	// The EC2 instance type for the Kubernetes control plane instances of your local Amazon EKS cluster on AWS Outposts.
	//
	// This instance type applies to all control plane instances and cannot be changed after cluster creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/eks_cluster#control_plane_instance_type EksCluster#control_plane_instance_type}
	ControlPlaneInstanceType *string `field:"optional" json:"controlPlaneInstanceType" yaml:"controlPlaneInstanceType"`
	// An object representing the placement configuration for all the control plane instances of your local Amazon EKS cluster on an AWS Outpost.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/eks_cluster#control_plane_placement EksCluster#control_plane_placement}
	ControlPlanePlacement *EksClusterOutpostConfigControlPlanePlacement `field:"optional" json:"controlPlanePlacement" yaml:"controlPlanePlacement"`
	// The EC2 instance type for etcd instances of your local Amazon EKS cluster on AWS Outposts.
	//
	// This instance type applies to all etcd instances and cannot be changed after cluster creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/eks_cluster#etcd_instance_type EksCluster#etcd_instance_type}
	EtcdInstanceType *string `field:"optional" json:"etcdInstanceType" yaml:"etcdInstanceType"`
	// An object representing the placement configuration for the etcd instances of your local Amazon EKS cluster on an AWS Outpost.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/eks_cluster#etcd_placement EksCluster#etcd_placement}
	EtcdPlacement *EksClusterOutpostConfigEtcdPlacement `field:"optional" json:"etcdPlacement" yaml:"etcdPlacement"`
	// The ARN of the Outpost that you want to use for your local Amazon EKS cluster on Outposts.
	//
	// Only a single Outpost ARN is supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/eks_cluster#outpost_arns EksCluster#outpost_arns}
	OutpostArns *[]*string `field:"optional" json:"outpostArns" yaml:"outpostArns"`
}

