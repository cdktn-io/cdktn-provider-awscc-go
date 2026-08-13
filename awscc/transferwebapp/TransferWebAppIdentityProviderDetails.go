// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package transferwebapp


type TransferWebAppIdentityProviderDetails struct {
	// The Amazon Resource Name (ARN) for the IAM Identity Center used for the web app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/transfer_web_app#instance_arn TransferWebApp#instance_arn}
	InstanceArn *string `field:"optional" json:"instanceArn" yaml:"instanceArn"`
	// The IAM role in IAM Identity Center used for the web app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/transfer_web_app#role TransferWebApp#role}
	Role *string `field:"optional" json:"role" yaml:"role"`
}

