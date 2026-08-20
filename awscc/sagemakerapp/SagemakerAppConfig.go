// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerapp

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerAppConfig struct {
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
	// The name of the app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_app#app_name SagemakerApp#app_name}
	AppName *string `field:"required" json:"appName" yaml:"appName"`
	// The type of app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_app#app_type SagemakerApp#app_type}
	AppType *string `field:"required" json:"appType" yaml:"appType"`
	// The domain ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_app#domain_id SagemakerApp#domain_id}
	DomainId *string `field:"required" json:"domainId" yaml:"domainId"`
	// The user profile name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_app#user_profile_name SagemakerApp#user_profile_name}
	UserProfileName *string `field:"required" json:"userProfileName" yaml:"userProfileName"`
	// Indicates whether the application is launched in recovery mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_app#recovery_mode SagemakerApp#recovery_mode}
	RecoveryMode interface{} `field:"optional" json:"recoveryMode" yaml:"recoveryMode"`
	// The instance type and the Amazon Resource Name (ARN) of the SageMaker image created on the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_app#resource_spec SagemakerApp#resource_spec}
	ResourceSpec *SagemakerAppResourceSpec `field:"optional" json:"resourceSpec" yaml:"resourceSpec"`
	// A list of tags to apply to the app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_app#tags SagemakerApp#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

