// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lambdamicrovmimage

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LambdaMicrovmImageConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lambda_microvm_image#additional_os_capabilities LambdaMicrovmImage#additional_os_capabilities}.
	AdditionalOsCapabilities *[]*string `field:"required" json:"additionalOsCapabilities" yaml:"additionalOsCapabilities"`
	// ARN of the base MicroVM image.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lambda_microvm_image#base_image_arn LambdaMicrovmImage#base_image_arn}
	BaseImageArn *string `field:"required" json:"baseImageArn" yaml:"baseImageArn"`
	// Specific version of the base MicroVM image to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lambda_microvm_image#base_image_version LambdaMicrovmImage#base_image_version}
	BaseImageVersion *string `field:"required" json:"baseImageVersion" yaml:"baseImageVersion"`
	// ARN of the IAM build role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lambda_microvm_image#build_role_arn LambdaMicrovmImage#build_role_arn}
	BuildRoleArn *string `field:"required" json:"buildRoleArn" yaml:"buildRoleArn"`
	// Code artifact for the active MicroVM image.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lambda_microvm_image#code_artifact LambdaMicrovmImage#code_artifact}
	CodeArtifact *LambdaMicrovmImageCodeArtifact `field:"required" json:"codeArtifact" yaml:"codeArtifact"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lambda_microvm_image#cpu_configurations LambdaMicrovmImage#cpu_configurations}.
	CpuConfigurations interface{} `field:"required" json:"cpuConfigurations" yaml:"cpuConfigurations"`
	// Human-readable description of the MicroVM image and its purpose.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lambda_microvm_image#description LambdaMicrovmImage#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lambda_microvm_image#egress_network_connectors LambdaMicrovmImage#egress_network_connectors}.
	EgressNetworkConnectors *[]*string `field:"required" json:"egressNetworkConnectors" yaml:"egressNetworkConnectors"`
	// Environment variables to set in the container during the snapshot build.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lambda_microvm_image#environment_variables LambdaMicrovmImage#environment_variables}
	EnvironmentVariables interface{} `field:"required" json:"environmentVariables" yaml:"environmentVariables"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lambda_microvm_image#hooks LambdaMicrovmImage#hooks}.
	Hooks *LambdaMicrovmImageHooks `field:"required" json:"hooks" yaml:"hooks"`
	// Configuration for MicroVM image logging.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lambda_microvm_image#logging LambdaMicrovmImage#logging}
	Logging *LambdaMicrovmImageLogging `field:"required" json:"logging" yaml:"logging"`
	// Unique name for the MicroVM image within the account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lambda_microvm_image#name LambdaMicrovmImage#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lambda_microvm_image#resources LambdaMicrovmImage#resources}.
	Resources interface{} `field:"required" json:"resources" yaml:"resources"`
	// Key-value pairs to associate with the MicroVM image for organization and management.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lambda_microvm_image#tags LambdaMicrovmImage#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

