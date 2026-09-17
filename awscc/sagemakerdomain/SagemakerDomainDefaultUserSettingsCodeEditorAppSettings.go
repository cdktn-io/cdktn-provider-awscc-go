// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerdomain


type SagemakerDomainDefaultUserSettingsCodeEditorAppSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_domain#app_lifecycle_management SagemakerDomain#app_lifecycle_management}.
	AppLifecycleManagement *SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagement `field:"optional" json:"appLifecycleManagement" yaml:"appLifecycleManagement"`
	// The lifecycle configuration that runs before the default lifecycle configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_domain#built_in_lifecycle_config_arn SagemakerDomain#built_in_lifecycle_config_arn}
	BuiltInLifecycleConfigArn *string `field:"optional" json:"builtInLifecycleConfigArn" yaml:"builtInLifecycleConfigArn"`
	// A list of custom images for use for CodeEditor apps.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_domain#custom_images SagemakerDomain#custom_images}
	CustomImages interface{} `field:"optional" json:"customImages" yaml:"customImages"`
	// The default instance type and the Amazon Resource Name (ARN) of the default SageMaker image used by the CodeEditor app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_domain#default_resource_spec SagemakerDomain#default_resource_spec}
	DefaultResourceSpec *SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsDefaultResourceSpec `field:"optional" json:"defaultResourceSpec" yaml:"defaultResourceSpec"`
	// A list of LifecycleConfigArns available for use with CodeEditor apps.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_domain#lifecycle_config_arns SagemakerDomain#lifecycle_config_arns}
	LifecycleConfigArns *[]*string `field:"optional" json:"lifecycleConfigArns" yaml:"lifecycleConfigArns"`
}

