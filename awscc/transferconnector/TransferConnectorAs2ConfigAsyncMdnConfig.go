// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package transferconnector


type TransferConnectorAs2ConfigAsyncMdnConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/transfer_connector#server_ids TransferConnector#server_ids}.
	ServerIds *[]*string `field:"optional" json:"serverIds" yaml:"serverIds"`
	// URL of the server to receive the MDN response on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/transfer_connector#url TransferConnector#url}
	Url *string `field:"optional" json:"url" yaml:"url"`
}

