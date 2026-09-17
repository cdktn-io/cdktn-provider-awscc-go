// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakercluster

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerClusterConfig struct {
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
	// Configuration for cluster auto-scaling.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_cluster#auto_scaling SagemakerCluster#auto_scaling}
	AutoScaling *SagemakerClusterAutoScaling `field:"optional" json:"autoScaling" yaml:"autoScaling"`
	// The name of the HyperPod Cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_cluster#cluster_name SagemakerCluster#cluster_name}
	ClusterName *string `field:"optional" json:"clusterName" yaml:"clusterName"`
	// The cluster role for the autoscaler to assume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_cluster#cluster_role SagemakerCluster#cluster_role}
	ClusterRole *string `field:"optional" json:"clusterRole" yaml:"clusterRole"`
	// The instance groups of the SageMaker HyperPod cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_cluster#instance_groups SagemakerCluster#instance_groups}
	InstanceGroups interface{} `field:"optional" json:"instanceGroups" yaml:"instanceGroups"`
	// Determines the scaling strategy for the SageMaker HyperPod cluster.
	//
	// When set to 'Continuous', enables continuous scaling which dynamically manages node provisioning. If the parameter is omitted, uses the standard scaling approach in previous release.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_cluster#node_provisioning_mode SagemakerCluster#node_provisioning_mode}
	NodeProvisioningMode *string `field:"optional" json:"nodeProvisioningMode" yaml:"nodeProvisioningMode"`
	// If node auto-recovery is set to true, faulty nodes will be replaced or rebooted when a failure is detected.
	//
	// If set to false, nodes will be labelled when a fault is detected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_cluster#node_recovery SagemakerCluster#node_recovery}
	NodeRecovery *string `field:"optional" json:"nodeRecovery" yaml:"nodeRecovery"`
	// Specifies parameter(s) specific to the orchestrator, e.g. specify the EKS cluster or Slurm configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_cluster#orchestrator SagemakerCluster#orchestrator}
	Orchestrator *SagemakerClusterOrchestrator `field:"optional" json:"orchestrator" yaml:"orchestrator"`
	// The restricted instance groups of the SageMaker HyperPod cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_cluster#restricted_instance_groups SagemakerCluster#restricted_instance_groups}
	RestrictedInstanceGroups interface{} `field:"optional" json:"restrictedInstanceGroups" yaml:"restrictedInstanceGroups"`
	// The cluster-level configuration for restricted instance groups, including shared environment settings for inter-RIG communication and FSx Lustre sharing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_cluster#restricted_instance_groups_config SagemakerCluster#restricted_instance_groups_config}
	RestrictedInstanceGroupsConfig *SagemakerClusterRestrictedInstanceGroupsConfig `field:"optional" json:"restrictedInstanceGroupsConfig" yaml:"restrictedInstanceGroupsConfig"`
	// Custom tags for managing the SageMaker HyperPod cluster as an AWS resource.
	//
	// You can add tags to your cluster in the same way you add them in other AWS services that support tagging.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_cluster#tags SagemakerCluster#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Configuration for tiered storage in the SageMaker HyperPod cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_cluster#tiered_storage_config SagemakerCluster#tiered_storage_config}
	TieredStorageConfig *SagemakerClusterTieredStorageConfig `field:"optional" json:"tieredStorageConfig" yaml:"tieredStorageConfig"`
	// Specifies an Amazon Virtual Private Cloud (VPC) that your SageMaker jobs, hosted models, and compute resources have access to.
	//
	// You can control access to and from your resources by configuring a VPC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_cluster#vpc_config SagemakerCluster#vpc_config}
	VpcConfig *SagemakerClusterVpcConfig `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}

