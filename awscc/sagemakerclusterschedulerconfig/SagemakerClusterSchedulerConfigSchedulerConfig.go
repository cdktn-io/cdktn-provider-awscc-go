// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerclusterschedulerconfig


type SagemakerClusterSchedulerConfigSchedulerConfig struct {
	// When enabled, entities borrow idle compute based on assigned FairShareWeight.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_cluster_scheduler_config#fair_share SagemakerClusterSchedulerConfig#fair_share}
	FairShare *string `field:"optional" json:"fairShare" yaml:"fairShare"`
	// Configuration for sharing idle compute resources across entities.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_cluster_scheduler_config#idle_resource_sharing SagemakerClusterSchedulerConfig#idle_resource_sharing}
	IdleResourceSharing *string `field:"optional" json:"idleResourceSharing" yaml:"idleResourceSharing"`
	// List of priority class configurations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_cluster_scheduler_config#priority_classes SagemakerClusterSchedulerConfig#priority_classes}
	PriorityClasses interface{} `field:"optional" json:"priorityClasses" yaml:"priorityClasses"`
}

