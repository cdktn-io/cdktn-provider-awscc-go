// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerimage

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerImageConfig struct {
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
	// The name of the image.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_image#image_name SagemakerImage#image_name}
	ImageName *string `field:"required" json:"imageName" yaml:"imageName"`
	// The Amazon Resource Name (ARN) of an IAM role that enables Amazon SageMaker to perform tasks on behalf of the customer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_image#image_role_arn SagemakerImage#image_role_arn}
	ImageRoleArn *string `field:"required" json:"imageRoleArn" yaml:"imageRoleArn"`
	// A description of the image.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_image#image_description SagemakerImage#image_description}
	ImageDescription *string `field:"optional" json:"imageDescription" yaml:"imageDescription"`
	// The display name of the image.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_image#image_display_name SagemakerImage#image_display_name}
	ImageDisplayName *string `field:"optional" json:"imageDisplayName" yaml:"imageDisplayName"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_image#tags SagemakerImage#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

