// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerstudiolifecycleconfig

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerStudioLifecycleConfigConfig struct {
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
	// The App type that the Lifecycle Configuration is attached to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_studio_lifecycle_config#studio_lifecycle_config_app_type SagemakerStudioLifecycleConfig#studio_lifecycle_config_app_type}
	StudioLifecycleConfigAppType *string `field:"required" json:"studioLifecycleConfigAppType" yaml:"studioLifecycleConfigAppType"`
	// The content of your Amazon SageMaker Studio Lifecycle Configuration script. This content must be base64 encoded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_studio_lifecycle_config#studio_lifecycle_config_content SagemakerStudioLifecycleConfig#studio_lifecycle_config_content}
	StudioLifecycleConfigContent *string `field:"required" json:"studioLifecycleConfigContent" yaml:"studioLifecycleConfigContent"`
	// The name of the Amazon SageMaker Studio Lifecycle Configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_studio_lifecycle_config#studio_lifecycle_config_name SagemakerStudioLifecycleConfig#studio_lifecycle_config_name}
	StudioLifecycleConfigName *string `field:"required" json:"studioLifecycleConfigName" yaml:"studioLifecycleConfigName"`
	// Tags to be associated with the Lifecycle Configuration.
	//
	// Each tag consists of a key and an optional value. Tag keys must be unique per resource. Tags are searchable using the Search API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_studio_lifecycle_config#tags SagemakerStudioLifecycleConfig#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

