// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package transferserver


type TransferServerProtocolDetailsProxyConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/transfer_server#sftp_mode TransferServer#sftp_mode}.
	SftpMode *string `field:"optional" json:"sftpMode" yaml:"sftpMode"`
}

