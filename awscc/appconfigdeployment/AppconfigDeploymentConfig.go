// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appconfigdeployment

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AppconfigDeploymentConfig struct {
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
	// The application ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appconfig_deployment#application_id AppconfigDeployment#application_id}
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// The configuration profile ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appconfig_deployment#configuration_profile_id AppconfigDeployment#configuration_profile_id}
	ConfigurationProfileId *string `field:"required" json:"configurationProfileId" yaml:"configurationProfileId"`
	// The configuration version to deploy.
	//
	// If deploying an AWS AppConfig hosted configuration version, you can specify either the version number or version label. For all other configurations, you must specify the version number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appconfig_deployment#configuration_version AppconfigDeployment#configuration_version}
	ConfigurationVersion *string `field:"required" json:"configurationVersion" yaml:"configurationVersion"`
	// The deployment strategy ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appconfig_deployment#deployment_strategy_id AppconfigDeployment#deployment_strategy_id}
	DeploymentStrategyId *string `field:"required" json:"deploymentStrategyId" yaml:"deploymentStrategyId"`
	// The environment ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appconfig_deployment#environment_id AppconfigDeployment#environment_id}
	EnvironmentId *string `field:"required" json:"environmentId" yaml:"environmentId"`
	// A description of the deployment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appconfig_deployment#description AppconfigDeployment#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appconfig_deployment#dynamic_extension_parameters AppconfigDeployment#dynamic_extension_parameters}.
	DynamicExtensionParameters interface{} `field:"optional" json:"dynamicExtensionParameters" yaml:"dynamicExtensionParameters"`
	// The AWS Key Management Service key identifier (key ID, key alias, or key ARN) provided when the resource was created or updated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appconfig_deployment#kms_key_identifier AppconfigDeployment#kms_key_identifier}
	KmsKeyIdentifier *string `field:"optional" json:"kmsKeyIdentifier" yaml:"kmsKeyIdentifier"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appconfig_deployment#tags AppconfigDeployment#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

