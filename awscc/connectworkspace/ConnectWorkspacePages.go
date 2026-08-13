// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectworkspace


type ConnectWorkspacePages struct {
	// The input data for the page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/connect_workspace#input_data ConnectWorkspace#input_data}
	InputData *string `field:"optional" json:"inputData" yaml:"inputData"`
	// The page identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/connect_workspace#page ConnectWorkspace#page}
	Page *string `field:"optional" json:"page" yaml:"page"`
	// The Amazon Resource Name (ARN) of the resource associated with the page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/connect_workspace#resource_arn ConnectWorkspace#resource_arn}
	ResourceArn *string `field:"optional" json:"resourceArn" yaml:"resourceArn"`
	// The slug for the page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/connect_workspace#slug ConnectWorkspace#slug}
	Slug *string `field:"optional" json:"slug" yaml:"slug"`
}

