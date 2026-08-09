// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package transferserver


type TransferServerWorkflowDetails struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/transfer_server#on_partial_upload TransferServer#on_partial_upload}.
	OnPartialUpload interface{} `field:"optional" json:"onPartialUpload" yaml:"onPartialUpload"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/transfer_server#on_upload TransferServer#on_upload}.
	OnUpload interface{} `field:"optional" json:"onUpload" yaml:"onUpload"`
}

