// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package accountaccessapplication


type AccountaccessApplicationIdentitySourceIdentityCenter struct {
	// The ARN of the Identity Center instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/accountaccess_application#instance_arn AccountaccessApplication#instance_arn}
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
}

