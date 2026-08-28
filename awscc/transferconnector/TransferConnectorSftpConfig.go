// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package transferconnector


type TransferConnectorSftpConfig struct {
	// Specifies the number of active connections that your connector can establish with the remote server at the same time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/transfer_connector#max_concurrent_connections TransferConnector#max_concurrent_connections}
	MaxConcurrentConnections *float64 `field:"optional" json:"maxConcurrentConnections" yaml:"maxConcurrentConnections"`
	// List of public host keys, for the external server to which you are connecting.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/transfer_connector#trusted_host_keys TransferConnector#trusted_host_keys}
	TrustedHostKeys *[]*string `field:"optional" json:"trustedHostKeys" yaml:"trustedHostKeys"`
	// ARN or name of the secret in AWS Secrets Manager which contains the SFTP user's private keys or passwords.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/transfer_connector#user_secret_id TransferConnector#user_secret_id}
	UserSecretId *string `field:"optional" json:"userSecretId" yaml:"userSecretId"`
}

