// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockdatasource


type BedrockDataSourceDataSourceConfigurationSalesforceConfigurationSourceConfiguration struct {
	// The supported authentication type to authenticate and connect to your Salesforce instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_data_source#auth_type BedrockDataSource#auth_type}
	AuthType *string `field:"optional" json:"authType" yaml:"authType"`
	// The Amazon Resource Name of an AWS Secrets Manager secret that stores your authentication credentials for your Salesforce instance URL.
	//
	// For more information on the key-value pairs that must be included in your secret, depending on your authentication type, see Salesforce connection configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_data_source#credentials_secret_arn BedrockDataSource#credentials_secret_arn}
	CredentialsSecretArn *string `field:"optional" json:"credentialsSecretArn" yaml:"credentialsSecretArn"`
	// The Salesforce host URL or instance URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_data_source#host_url BedrockDataSource#host_url}
	HostUrl *string `field:"optional" json:"hostUrl" yaml:"hostUrl"`
}

