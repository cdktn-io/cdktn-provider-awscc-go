// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerprocessingjob


type SagemakerProcessingJobProcessingResources struct {
	// Configuration for the cluster used to run a processing job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_processing_job#cluster_config SagemakerProcessingJob#cluster_config}
	ClusterConfig *SagemakerProcessingJobProcessingResourcesClusterConfig `field:"required" json:"clusterConfig" yaml:"clusterConfig"`
}

