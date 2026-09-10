// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package greengrassv2componentversion


type Greengrassv2ComponentVersionLambdaFunction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/greengrassv2_component_version#component_dependencies Greengrassv2ComponentVersion#component_dependencies}.
	ComponentDependencies interface{} `field:"optional" json:"componentDependencies" yaml:"componentDependencies"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/greengrassv2_component_version#component_lambda_parameters Greengrassv2ComponentVersion#component_lambda_parameters}.
	ComponentLambdaParameters *Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParameters `field:"optional" json:"componentLambdaParameters" yaml:"componentLambdaParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/greengrassv2_component_version#component_name Greengrassv2ComponentVersion#component_name}.
	ComponentName *string `field:"optional" json:"componentName" yaml:"componentName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/greengrassv2_component_version#component_platforms Greengrassv2ComponentVersion#component_platforms}.
	ComponentPlatforms interface{} `field:"optional" json:"componentPlatforms" yaml:"componentPlatforms"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/greengrassv2_component_version#component_version Greengrassv2ComponentVersion#component_version}.
	ComponentVersion *string `field:"optional" json:"componentVersion" yaml:"componentVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/greengrassv2_component_version#lambda_arn Greengrassv2ComponentVersion#lambda_arn}.
	LambdaArn *string `field:"optional" json:"lambdaArn" yaml:"lambdaArn"`
}

