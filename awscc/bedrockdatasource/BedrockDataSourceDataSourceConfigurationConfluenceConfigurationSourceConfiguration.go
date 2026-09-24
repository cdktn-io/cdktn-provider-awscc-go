// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockdatasource


type BedrockDataSourceDataSourceConfigurationConfluenceConfigurationSourceConfiguration struct {
	// The supported authentication type to authenticate and connect to your Confluence instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrock_data_source#auth_type BedrockDataSource#auth_type}
	AuthType *string `field:"optional" json:"authType" yaml:"authType"`
	// The Amazon Resource Name of an AWS Secrets Manager secret that stores your authentication credentials for your Confluence instance URL.
	//
	// For more information on the key-value pairs that must be included in your secret, depending on your authentication type, see Confluence connection configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrock_data_source#credentials_secret_arn BedrockDataSource#credentials_secret_arn}
	CredentialsSecretArn *string `field:"optional" json:"credentialsSecretArn" yaml:"credentialsSecretArn"`
	// The supported host type, whether online/cloud or server/on-premises.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrock_data_source#host_type BedrockDataSource#host_type}
	HostType *string `field:"optional" json:"hostType" yaml:"hostType"`
	// The Confluence host URL or instance URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrock_data_source#host_url BedrockDataSource#host_url}
	HostUrl *string `field:"optional" json:"hostUrl" yaml:"hostUrl"`
}

