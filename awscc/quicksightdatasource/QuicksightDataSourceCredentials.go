// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightdatasource


type QuicksightDataSourceCredentials struct {
	// <p>The Amazon Resource Name (ARN) of a data source that has the credential pair that you             want to use.
	//
	// When <code>CopySourceArn</code> is not null, the credential pair from the
	//             data source in the ARN is used as the credentials for the
	//             <code>DataSourceCredentials</code> structure.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/quicksight_data_source#copy_source_arn QuicksightDataSource#copy_source_arn}
	CopySourceArn *string `field:"optional" json:"copySourceArn" yaml:"copySourceArn"`
	// <p>The combination of user name and password that are used as credentials.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/quicksight_data_source#credential_pair QuicksightDataSource#credential_pair}
	CredentialPair *QuicksightDataSourceCredentialsCredentialPair `field:"optional" json:"credentialPair" yaml:"credentialPair"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/quicksight_data_source#key_pair_credentials QuicksightDataSource#key_pair_credentials}.
	KeyPairCredentials *QuicksightDataSourceCredentialsKeyPairCredentials `field:"optional" json:"keyPairCredentials" yaml:"keyPairCredentials"`
	// <p>The Amazon Resource Name (ARN) of the secret associated with the data source in Amazon Secrets Manager.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/quicksight_data_source#secret_arn QuicksightDataSource#secret_arn}
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
}

