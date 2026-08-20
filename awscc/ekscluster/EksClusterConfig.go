// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ekscluster

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EksClusterConfig struct {
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
	// An object representing the VPC configuration to use for an Amazon EKS cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/eks_cluster#resources_vpc_config EksCluster#resources_vpc_config}
	ResourcesVpcConfig *EksClusterResourcesVpcConfig `field:"required" json:"resourcesVpcConfig" yaml:"resourcesVpcConfig"`
	// The Amazon Resource Name (ARN) of the IAM role that provides permissions for the Kubernetes control plane to make calls to AWS API operations on your behalf.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/eks_cluster#role_arn EksCluster#role_arn}
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// An object representing the Access Config to use for the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/eks_cluster#access_config EksCluster#access_config}
	AccessConfig *EksClusterAccessConfig `field:"optional" json:"accessConfig" yaml:"accessConfig"`
	// Set this value to false to avoid creating the default networking add-ons when the cluster is created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/eks_cluster#bootstrap_self_managed_addons EksCluster#bootstrap_self_managed_addons}
	BootstrapSelfManagedAddons interface{} `field:"optional" json:"bootstrapSelfManagedAddons" yaml:"bootstrapSelfManagedAddons"`
	// Todo: add description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/eks_cluster#compute_config EksCluster#compute_config}
	ComputeConfig *EksClusterComputeConfig `field:"optional" json:"computeConfig" yaml:"computeConfig"`
	// Configuration for provisioned control plane scaling.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/eks_cluster#control_plane_scaling_config EksCluster#control_plane_scaling_config}
	ControlPlaneScalingConfig *EksClusterControlPlaneScalingConfig `field:"optional" json:"controlPlaneScalingConfig" yaml:"controlPlaneScalingConfig"`
	// Set this value to true to enable deletion protection for the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/eks_cluster#deletion_protection EksCluster#deletion_protection}
	DeletionProtection interface{} `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/eks_cluster#encryption_config EksCluster#encryption_config}.
	EncryptionConfig interface{} `field:"optional" json:"encryptionConfig" yaml:"encryptionConfig"`
	// Force cluster version update.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/eks_cluster#force EksCluster#force}
	Force interface{} `field:"optional" json:"force" yaml:"force"`
	// The configuration for the Kubernetes API server on an Amazon EKS cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/eks_cluster#kube_api_server_config EksCluster#kube_api_server_config}
	KubeApiServerConfig *EksClusterKubeApiServerConfig `field:"optional" json:"kubeApiServerConfig" yaml:"kubeApiServerConfig"`
	// The configuration for the Kubernetes controller manager on an Amazon EKS cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/eks_cluster#kube_controller_manager_config EksCluster#kube_controller_manager_config}
	KubeControllerManagerConfig *EksClusterKubeControllerManagerConfig `field:"optional" json:"kubeControllerManagerConfig" yaml:"kubeControllerManagerConfig"`
	// The Kubernetes network configuration for the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/eks_cluster#kubernetes_network_config EksCluster#kubernetes_network_config}
	KubernetesNetworkConfig *EksClusterKubernetesNetworkConfig `field:"optional" json:"kubernetesNetworkConfig" yaml:"kubernetesNetworkConfig"`
	// The configuration for the Kubernetes scheduler on an Amazon EKS cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/eks_cluster#kube_scheduler_config EksCluster#kube_scheduler_config}
	KubeSchedulerConfig *EksClusterKubeSchedulerConfig `field:"optional" json:"kubeSchedulerConfig" yaml:"kubeSchedulerConfig"`
	// Enable exporting the Kubernetes control plane logs for your cluster to CloudWatch Logs based on log types.
	//
	// By default, cluster control plane logs aren't exported to CloudWatch Logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/eks_cluster#logging EksCluster#logging}
	Logging *EksClusterLogging `field:"optional" json:"logging" yaml:"logging"`
	// The unique name to give to your cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/eks_cluster#name EksCluster#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// An object representing the Outpost configuration to use for AWS EKS outpost cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/eks_cluster#outpost_config EksCluster#outpost_config}
	OutpostConfig *EksClusterOutpostConfig `field:"optional" json:"outpostConfig" yaml:"outpostConfig"`
	// Configuration fields for specifying on-premises node and pod CIDRs that are external to the VPC passed during cluster creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/eks_cluster#remote_network_config EksCluster#remote_network_config}
	RemoteNetworkConfig *EksClusterRemoteNetworkConfig `field:"optional" json:"remoteNetworkConfig" yaml:"remoteNetworkConfig"`
	// The rollback configuration to use for the cluster version rollback.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/eks_cluster#rollback_config EksCluster#rollback_config}
	RollbackConfig *EksClusterRollbackConfig `field:"optional" json:"rollbackConfig" yaml:"rollbackConfig"`
	// Todo: add description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/eks_cluster#storage_config EksCluster#storage_config}
	StorageConfig *EksClusterStorageConfig `field:"optional" json:"storageConfig" yaml:"storageConfig"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/eks_cluster#tags EksCluster#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// An object representing the Upgrade Policy to use for the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/eks_cluster#upgrade_policy EksCluster#upgrade_policy}
	UpgradePolicy *EksClusterUpgradePolicy `field:"optional" json:"upgradePolicy" yaml:"upgradePolicy"`
	// The desired Kubernetes version for your cluster.
	//
	// If you don't specify a value here, the latest version available in Amazon EKS is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/eks_cluster#version EksCluster#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
	// The current zonal shift configuration to use for the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/eks_cluster#zonal_shift_config EksCluster#zonal_shift_config}
	ZonalShiftConfig *EksClusterZonalShiftConfig `field:"optional" json:"zonalShiftConfig" yaml:"zonalShiftConfig"`
}

