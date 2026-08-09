// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerhub

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerHubConfig struct {
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
	// A description of the hub.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/sagemaker_hub#hub_description SagemakerHub#hub_description}
	HubDescription *string `field:"required" json:"hubDescription" yaml:"hubDescription"`
	// The name of the hub.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/sagemaker_hub#hub_name SagemakerHub#hub_name}
	HubName *string `field:"required" json:"hubName" yaml:"hubName"`
	// The display name of the hub.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/sagemaker_hub#hub_display_name SagemakerHub#hub_display_name}
	HubDisplayName *string `field:"optional" json:"hubDisplayName" yaml:"hubDisplayName"`
	// The searchable keywords for the hub.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/sagemaker_hub#hub_search_keywords SagemakerHub#hub_search_keywords}
	HubSearchKeywords *[]*string `field:"optional" json:"hubSearchKeywords" yaml:"hubSearchKeywords"`
	// The Amazon S3 storage configuration for the hub.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/sagemaker_hub#s3_storage_config SagemakerHub#s3_storage_config}
	S3StorageConfig *SagemakerHubS3StorageConfig `field:"optional" json:"s3StorageConfig" yaml:"s3StorageConfig"`
	// Tags to associate with the hub.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/sagemaker_hub#tags SagemakerHub#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

