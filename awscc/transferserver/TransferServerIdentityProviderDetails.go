// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package transferserver


type TransferServerIdentityProviderDetails struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/transfer_server#directory_id TransferServer#directory_id}.
	DirectoryId *string `field:"optional" json:"directoryId" yaml:"directoryId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/transfer_server#function TransferServer#function}.
	Function *string `field:"optional" json:"function" yaml:"function"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/transfer_server#invocation_role TransferServer#invocation_role}.
	InvocationRole *string `field:"optional" json:"invocationRole" yaml:"invocationRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/transfer_server#sftp_authentication_methods TransferServer#sftp_authentication_methods}.
	SftpAuthenticationMethods *string `field:"optional" json:"sftpAuthenticationMethods" yaml:"sftpAuthenticationMethods"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/transfer_server#url TransferServer#url}.
	Url *string `field:"optional" json:"url" yaml:"url"`
}

