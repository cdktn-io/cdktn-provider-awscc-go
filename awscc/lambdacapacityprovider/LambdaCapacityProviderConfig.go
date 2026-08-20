// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lambdacapacityprovider

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LambdaCapacityProviderConfig struct {
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
	// The permissions configuration for the capacity provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/lambda_capacity_provider#permissions_config LambdaCapacityProvider#permissions_config}
	PermissionsConfig *LambdaCapacityProviderPermissionsConfig `field:"required" json:"permissionsConfig" yaml:"permissionsConfig"`
	// The VPC configuration for the capacity provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/lambda_capacity_provider#vpc_config LambdaCapacityProvider#vpc_config}
	VpcConfig *LambdaCapacityProviderVpcConfig `field:"required" json:"vpcConfig" yaml:"vpcConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/lambda_capacity_provider#capacity_provider_name LambdaCapacityProvider#capacity_provider_name}.
	CapacityProviderName *string `field:"optional" json:"capacityProviderName" yaml:"capacityProviderName"`
	// The scaling configuration for the capacity provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/lambda_capacity_provider#capacity_provider_scaling_config LambdaCapacityProvider#capacity_provider_scaling_config}
	CapacityProviderScalingConfig *LambdaCapacityProviderCapacityProviderScalingConfig `field:"optional" json:"capacityProviderScalingConfig" yaml:"capacityProviderScalingConfig"`
	// The instance requirements for compute resources managed by the capacity provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/lambda_capacity_provider#instance_requirements LambdaCapacityProvider#instance_requirements}
	InstanceRequirements *LambdaCapacityProviderInstanceRequirements `field:"optional" json:"instanceRequirements" yaml:"instanceRequirements"`
	// The ARN of the KMS key used to encrypt the capacity provider's resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/lambda_capacity_provider#kms_key_arn LambdaCapacityProvider#kms_key_arn}
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// Configuration for tag propagation to managed resources launched by the capacity provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/lambda_capacity_provider#propagate_tags LambdaCapacityProvider#propagate_tags}
	PropagateTags *LambdaCapacityProviderPropagateTags `field:"optional" json:"propagateTags" yaml:"propagateTags"`
	// A key-value pair that provides metadata for the capacity provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/lambda_capacity_provider#tags LambdaCapacityProvider#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The telemetry configuration for the capacity provider, including logging settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/lambda_capacity_provider#telemetry_config LambdaCapacityProvider#telemetry_config}
	TelemetryConfig *LambdaCapacityProviderTelemetryConfig `field:"optional" json:"telemetryConfig" yaml:"telemetryConfig"`
}

