// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakercomputequota

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerComputeQuotaConfig struct {
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
	// The ARN of the EKS-orchestrated SageMaker HyperPod cluster the compute allocation definition is bound to.
	//
	// Slurm-orchestrated clusters are rejected by the service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/sagemaker_compute_quota#cluster_arn SagemakerComputeQuota#cluster_arn}
	ClusterArn *string `field:"required" json:"clusterArn" yaml:"clusterArn"`
	// Compute allocation rules: per-resource quotas, idle-borrowing strategy, and team-task preemption policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/sagemaker_compute_quota#compute_quota_config SagemakerComputeQuota#compute_quota_config}
	ComputeQuotaConfig *SagemakerComputeQuotaComputeQuotaConfig `field:"required" json:"computeQuotaConfig" yaml:"computeQuotaConfig"`
	// The team this compute allocation rule applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/sagemaker_compute_quota#compute_quota_target SagemakerComputeQuota#compute_quota_target}
	ComputeQuotaTarget *SagemakerComputeQuotaComputeQuotaTarget `field:"required" json:"computeQuotaTarget" yaml:"computeQuotaTarget"`
	// The name of the compute allocation definition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/sagemaker_compute_quota#name SagemakerComputeQuota#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Whether the compute allocation rule is active. Defaults service-side to 'Enabled' when omitted on Create.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/sagemaker_compute_quota#activation_state SagemakerComputeQuota#activation_state}
	ActivationState *string `field:"optional" json:"activationState" yaml:"activationState"`
	// A free-form description of the compute allocation definition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/sagemaker_compute_quota#description SagemakerComputeQuota#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Tags to associate with the compute allocation definition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/sagemaker_compute_quota#tags SagemakerComputeQuota#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

