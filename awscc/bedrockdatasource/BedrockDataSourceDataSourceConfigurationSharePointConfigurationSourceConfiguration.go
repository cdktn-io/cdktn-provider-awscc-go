// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockdatasource


type BedrockDataSourceDataSourceConfigurationSharePointConfigurationSourceConfiguration struct {
	// The supported authentication type to authenticate and connect to your SharePoint site/sites.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/bedrock_data_source#auth_type BedrockDataSource#auth_type}
	AuthType *string `field:"optional" json:"authType" yaml:"authType"`
	// The Amazon Resource Name of an AWS Secrets Manager secret that stores your authentication credentials for your SharePoint site/sites.
	//
	// For more information on the key-value pairs that must be included in your secret, depending on your authentication type, see SharePoint connection configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/bedrock_data_source#credentials_secret_arn BedrockDataSource#credentials_secret_arn}
	CredentialsSecretArn *string `field:"optional" json:"credentialsSecretArn" yaml:"credentialsSecretArn"`
	// The domain of your SharePoint instance or site URL/URLs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/bedrock_data_source#domain BedrockDataSource#domain}
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// The supported host type, whether online/cloud or server/on-premises.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/bedrock_data_source#host_type BedrockDataSource#host_type}
	HostType *string `field:"optional" json:"hostType" yaml:"hostType"`
	// A list of one or more SharePoint site URLs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/bedrock_data_source#site_urls BedrockDataSource#site_urls}
	SiteUrls *[]*string `field:"optional" json:"siteUrls" yaml:"siteUrls"`
	// The identifier of your Microsoft 365 tenant.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/bedrock_data_source#tenant_id BedrockDataSource#tenant_id}
	TenantId *string `field:"optional" json:"tenantId" yaml:"tenantId"`
}

