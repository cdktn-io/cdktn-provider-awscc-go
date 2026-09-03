// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerprocessingjob


type SagemakerProcessingJobProcessingOutputConfigOutputsFeatureStoreOutput struct {
	// The name of the Amazon SageMaker FeatureGroup to use as the destination for processing job output.
	//
	// Note that your processing script is responsible for putting records into your Feature Store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_processing_job#feature_group_name SagemakerProcessingJob#feature_group_name}
	FeatureGroupName *string `field:"optional" json:"featureGroupName" yaml:"featureGroupName"`
}

