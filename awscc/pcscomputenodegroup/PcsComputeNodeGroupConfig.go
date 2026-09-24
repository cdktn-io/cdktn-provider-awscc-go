// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pcscomputenodegroup

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PcsComputeNodeGroupConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The ID of the cluster of the compute node group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/pcs_compute_node_group#cluster_id PcsComputeNodeGroup#cluster_id}
	ClusterId *string `field:"required" json:"clusterId" yaml:"clusterId"`
	// An Amazon EC2 launch template AWS PCS uses to launch compute nodes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/pcs_compute_node_group#custom_launch_template PcsComputeNodeGroup#custom_launch_template}
	CustomLaunchTemplate *PcsComputeNodeGroupCustomLaunchTemplate `field:"required" json:"customLaunchTemplate" yaml:"customLaunchTemplate"`
	// The Amazon Resource Name (ARN) of the IAM instance profile used to pass an IAM role when launching EC2 instances.
	//
	// The role contained in your instance profile must have pcs:RegisterComputeNodeGroupInstance permissions attached to provision instances correctly.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/pcs_compute_node_group#iam_instance_profile_arn PcsComputeNodeGroup#iam_instance_profile_arn}
	IamInstanceProfileArn *string `field:"required" json:"iamInstanceProfileArn" yaml:"iamInstanceProfileArn"`
	// A list of EC2 instance configurations that AWS PCS can provision in the compute node group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/pcs_compute_node_group#instance_configs PcsComputeNodeGroup#instance_configs}
	InstanceConfigs interface{} `field:"required" json:"instanceConfigs" yaml:"instanceConfigs"`
	// Specifies the boundaries of the compute node group auto scaling.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/pcs_compute_node_group#scaling_configuration PcsComputeNodeGroup#scaling_configuration}
	ScalingConfiguration *PcsComputeNodeGroupScalingConfiguration `field:"required" json:"scalingConfiguration" yaml:"scalingConfiguration"`
	// The list of subnet IDs where instances are provisioned by the compute node group.
	//
	// The subnets must be in the same VPC as the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/pcs_compute_node_group#subnet_ids PcsComputeNodeGroup#subnet_ids}
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// The ID of the Amazon Machine Image (AMI) that AWS PCS uses to launch instances.
	//
	// If not provided, AWS PCS uses the AMI ID specified in the custom launch template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/pcs_compute_node_group#ami_id PcsComputeNodeGroup#ami_id}
	AmiId *string `field:"optional" json:"amiId" yaml:"amiId"`
	// The name that identifies the compute node group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/pcs_compute_node_group#name PcsComputeNodeGroup#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Custom scripts that run at defined points in a compute node's lifecycle.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/pcs_compute_node_group#node_lifecycle_actions PcsComputeNodeGroup#node_lifecycle_actions}
	NodeLifecycleActions *PcsComputeNodeGroupNodeLifecycleActions `field:"optional" json:"nodeLifecycleActions" yaml:"nodeLifecycleActions"`
	// Specifies how EC2 instances are purchased on your behalf.
	//
	// AWS PCS supports On-Demand, Spot, Capacity Block, and Interruptible Capacity Reservation instances. For more information, see Instance purchasing options in the Amazon Elastic Compute Cloud User Guide. If you don't provide this option, it defaults to On-Demand.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/pcs_compute_node_group#purchase_option PcsComputeNodeGroup#purchase_option}
	PurchaseOption *string `field:"optional" json:"purchaseOption" yaml:"purchaseOption"`
	// Additional options related to the Slurm scheduler.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/pcs_compute_node_group#slurm_configuration PcsComputeNodeGroup#slurm_configuration}
	SlurmConfiguration *PcsComputeNodeGroupSlurmConfiguration `field:"optional" json:"slurmConfiguration" yaml:"slurmConfiguration"`
	// Additional configuration when you specify SPOT as the purchase option.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/pcs_compute_node_group#spot_options PcsComputeNodeGroup#spot_options}
	SpotOptions *PcsComputeNodeGroupSpotOptions `field:"optional" json:"spotOptions" yaml:"spotOptions"`
	// 1 or more tags added to the resource.
	//
	// Each tag consists of a tag key and tag value. The tag value is optional and can be an empty string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/pcs_compute_node_group#tags PcsComputeNodeGroup#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

