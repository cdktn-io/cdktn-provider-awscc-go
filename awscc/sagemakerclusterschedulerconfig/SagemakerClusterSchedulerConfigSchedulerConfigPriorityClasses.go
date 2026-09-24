// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerclusterschedulerconfig


type SagemakerClusterSchedulerConfigSchedulerConfigPriorityClasses struct {
	// Name of the priority class.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_cluster_scheduler_config#name SagemakerClusterSchedulerConfig#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Weight of the priority class. Range 0-100, default 0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_cluster_scheduler_config#weight SagemakerClusterSchedulerConfig#weight}
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

