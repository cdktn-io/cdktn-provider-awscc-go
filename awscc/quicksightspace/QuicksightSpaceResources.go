// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightspace


type QuicksightSpaceResources struct {
	// The ARN of the QuickSight resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/quicksight_space#resource_arn QuicksightSpace#resource_arn}
	ResourceArn *string `field:"optional" json:"resourceArn" yaml:"resourceArn"`
	// The type of QuickSight resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/quicksight_space#resource_type QuicksightSpace#resource_type}
	ResourceType *string `field:"optional" json:"resourceType" yaml:"resourceType"`
}

