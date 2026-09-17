// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appconfigdeploymentstrategy

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AppconfigDeploymentStrategyConfig struct {
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
	// Total amount of time for a deployment to last.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appconfig_deployment_strategy#deployment_duration_in_minutes AppconfigDeploymentStrategy#deployment_duration_in_minutes}
	DeploymentDurationInMinutes *float64 `field:"required" json:"deploymentDurationInMinutes" yaml:"deploymentDurationInMinutes"`
	// The percentage of targets to receive a deployed configuration during each interval.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appconfig_deployment_strategy#growth_factor AppconfigDeploymentStrategy#growth_factor}
	GrowthFactor *float64 `field:"required" json:"growthFactor" yaml:"growthFactor"`
	// A name for the deployment strategy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appconfig_deployment_strategy#name AppconfigDeploymentStrategy#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Save the deployment strategy to a Systems Manager (SSM) document.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appconfig_deployment_strategy#replicate_to AppconfigDeploymentStrategy#replicate_to}
	ReplicateTo *string `field:"required" json:"replicateTo" yaml:"replicateTo"`
	// A description of the deployment strategy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appconfig_deployment_strategy#description AppconfigDeploymentStrategy#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies the amount of time AWS AppConfig monitors for Amazon CloudWatch alarms after the configuration has been deployed to 100% of its targets, before considering the deployment to be complete.
	//
	// If an alarm is triggered during this time, AWS AppConfig rolls back the deployment. You must configure permissions for AWS AppConfig to roll back based on CloudWatch alarms. For more information, see Configuring permissions for rollback based on Amazon CloudWatch alarms in the AWS AppConfig User Guide.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appconfig_deployment_strategy#final_bake_time_in_minutes AppconfigDeploymentStrategy#final_bake_time_in_minutes}
	FinalBakeTimeInMinutes *float64 `field:"optional" json:"finalBakeTimeInMinutes" yaml:"finalBakeTimeInMinutes"`
	// The algorithm used to define how percentage grows over time. AWS AppConfig supports the following growth types:.
	//
	// Linear: For this type, AWS AppConfig processes the deployment by dividing the total number of targets by the value specified for Step percentage. For example, a linear deployment that uses a Step percentage of 10 deploys the configuration to 10 percent of the hosts. After those deployments are complete, the system deploys the configuration to the next 10 percent. This continues until 100% of the targets have successfully received the configuration.
	//
	// Exponential: For this type, AWS AppConfig processes the deployment exponentially using the following formula: G*(2^N). In this formula, G is the growth factor specified by the user and N is the number of steps until the configuration is deployed to all targets. For example, if you specify a growth factor of 2, then the system rolls out the configuration as follows:
	//
	// 2*(2^0)
	//
	// 2*(2^1)
	//
	// 2*(2^2)
	//
	// Expressed numerically, the deployment rolls out as follows: 2% of the targets, 4% of the targets, 8% of the targets, and continues until the configuration has been deployed to all targets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appconfig_deployment_strategy#growth_type AppconfigDeploymentStrategy#growth_type}
	GrowthType *string `field:"optional" json:"growthType" yaml:"growthType"`
	// Assigns metadata to an AWS AppConfig resource.
	//
	// Tags help organize and categorize your AWS AppConfig resources. Each tag consists of a key and an optional value, both of which you define. You can specify a maximum of 50 tags for a resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appconfig_deployment_strategy#tags AppconfigDeploymentStrategy#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

