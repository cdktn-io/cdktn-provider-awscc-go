// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kendradatasource


type KendraDataSourceDataSourceConfigurationServiceNowConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/kendra_data_source#authentication_type KendraDataSource#authentication_type}.
	AuthenticationType *string `field:"optional" json:"authenticationType" yaml:"authenticationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/kendra_data_source#host_url KendraDataSource#host_url}.
	HostUrl *string `field:"optional" json:"hostUrl" yaml:"hostUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/kendra_data_source#knowledge_article_configuration KendraDataSource#knowledge_article_configuration}.
	KnowledgeArticleConfiguration *KendraDataSourceDataSourceConfigurationServiceNowConfigurationKnowledgeArticleConfiguration `field:"optional" json:"knowledgeArticleConfiguration" yaml:"knowledgeArticleConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/kendra_data_source#secret_arn KendraDataSource#secret_arn}.
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/kendra_data_source#service_catalog_configuration KendraDataSource#service_catalog_configuration}.
	ServiceCatalogConfiguration *KendraDataSourceDataSourceConfigurationServiceNowConfigurationServiceCatalogConfiguration `field:"optional" json:"serviceCatalogConfiguration" yaml:"serviceCatalogConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/kendra_data_source#service_now_build_version KendraDataSource#service_now_build_version}.
	ServiceNowBuildVersion *string `field:"optional" json:"serviceNowBuildVersion" yaml:"serviceNowBuildVersion"`
}

