// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightdatasource


type QuicksightDataSourceCredentialsKeyPairCredentials struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/quicksight_data_source#key_pair_username QuicksightDataSource#key_pair_username}.
	KeyPairUsername *string `field:"optional" json:"keyPairUsername" yaml:"keyPairUsername"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/quicksight_data_source#private_key QuicksightDataSource#private_key}.
	PrivateKey *string `field:"optional" json:"privateKey" yaml:"privateKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/quicksight_data_source#private_key_passphrase QuicksightDataSource#private_key_passphrase}.
	PrivateKeyPassphrase *string `field:"optional" json:"privateKeyPassphrase" yaml:"privateKeyPassphrase"`
}

