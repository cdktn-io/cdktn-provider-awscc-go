// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package transferserver


type TransferServerProtocolDetailsSftpPorts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/transfer_server#communication_mode TransferServer#communication_mode}.
	CommunicationMode *string `field:"optional" json:"communicationMode" yaml:"communicationMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/transfer_server#sftp_port TransferServer#sftp_port}.
	SftpPort *float64 `field:"optional" json:"sftpPort" yaml:"sftpPort"`
}

