// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakercluster


type SagemakerClusterInstanceGroupsLifeCycleConfig struct {
	// The file name of the entrypoint script of lifecycle scripts under SourceS3Uri.
	//
	// This entrypoint script runs during cluster creation. Mutually exclusive with OnInitComplete.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_cluster#on_create SagemakerCluster#on_create}
	OnCreate *string `field:"optional" json:"onCreate" yaml:"onCreate"`
	// The file name of the extension script under SourceS3Uri.
	//
	// This script runs after HyperPod configures the default software on the instance. Mutually exclusive with OnCreate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_cluster#on_init_complete SagemakerCluster#on_init_complete}
	OnInitComplete *string `field:"optional" json:"onInitComplete" yaml:"onInitComplete"`
	// An Amazon S3 bucket path where your lifecycle scripts are stored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_cluster#source_s3_uri SagemakerCluster#source_s3_uri}
	SourceS3Uri *string `field:"optional" json:"sourceS3Uri" yaml:"sourceS3Uri"`
}

