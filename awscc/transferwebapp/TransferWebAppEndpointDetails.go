// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package transferwebapp


type TransferWebAppEndpointDetails struct {
	// You can provide a structure that contains the details for the VPC endpoint to use with your web app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/transfer_web_app#vpc TransferWebApp#vpc}
	Vpc *TransferWebAppEndpointDetailsVpc `field:"optional" json:"vpc" yaml:"vpc"`
}

