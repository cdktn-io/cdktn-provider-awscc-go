// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakercluster


type SagemakerClusterInstanceGroupsScheduledUpdateConfigDeploymentConfigRollingUpdatePolicyRollbackMaximumBatchSize struct {
	// Specifies whether SageMaker should process the update by amount or percentage of instances.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_cluster#type SagemakerCluster#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Specifies the amount or percentage of instances SageMaker updates at a time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_cluster#value SagemakerCluster#value}
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}

