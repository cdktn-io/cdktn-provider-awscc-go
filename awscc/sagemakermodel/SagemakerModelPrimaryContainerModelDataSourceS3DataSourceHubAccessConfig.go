// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakermodel


type SagemakerModelPrimaryContainerModelDataSourceS3DataSourceHubAccessConfig struct {
	// The ARN of the hub content for which deployment access is allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/sagemaker_model#hub_content_arn SagemakerModel#hub_content_arn}
	HubContentArn *string `field:"optional" json:"hubContentArn" yaml:"hubContentArn"`
}

