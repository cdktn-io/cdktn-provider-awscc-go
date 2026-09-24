// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakernotebookinstance


type SagemakerNotebookInstanceInstanceMetadataServiceConfiguration struct {
	// Indicates the minimum IMDS version that the notebook instance supports.
	//
	// When passed as part of CreateNotebookInstance, if no value is selected, then it defaults to IMDSv1. This means that both IMDSv1 and IMDSv2 are supported. If passed as part of UpdateNotebookInstance, there is no default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_notebook_instance#minimum_instance_metadata_service_version SagemakerNotebookInstance#minimum_instance_metadata_service_version}
	MinimumInstanceMetadataServiceVersion *string `field:"optional" json:"minimumInstanceMetadataServiceVersion" yaml:"minimumInstanceMetadataServiceVersion"`
}

