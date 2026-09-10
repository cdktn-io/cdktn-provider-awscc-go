// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kendradatasource


type KendraDataSourceDataSourceConfigurationSharePointConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/kendra_data_source#crawl_attachments KendraDataSource#crawl_attachments}.
	CrawlAttachments interface{} `field:"optional" json:"crawlAttachments" yaml:"crawlAttachments"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/kendra_data_source#disable_local_groups KendraDataSource#disable_local_groups}.
	DisableLocalGroups interface{} `field:"optional" json:"disableLocalGroups" yaml:"disableLocalGroups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/kendra_data_source#document_title_field_name KendraDataSource#document_title_field_name}.
	DocumentTitleFieldName *string `field:"optional" json:"documentTitleFieldName" yaml:"documentTitleFieldName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/kendra_data_source#exclusion_patterns KendraDataSource#exclusion_patterns}.
	ExclusionPatterns *[]*string `field:"optional" json:"exclusionPatterns" yaml:"exclusionPatterns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/kendra_data_source#field_mappings KendraDataSource#field_mappings}.
	FieldMappings interface{} `field:"optional" json:"fieldMappings" yaml:"fieldMappings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/kendra_data_source#inclusion_patterns KendraDataSource#inclusion_patterns}.
	InclusionPatterns *[]*string `field:"optional" json:"inclusionPatterns" yaml:"inclusionPatterns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/kendra_data_source#secret_arn KendraDataSource#secret_arn}.
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/kendra_data_source#share_point_version KendraDataSource#share_point_version}.
	SharePointVersion *string `field:"optional" json:"sharePointVersion" yaml:"sharePointVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/kendra_data_source#ssl_certificate_s3_path KendraDataSource#ssl_certificate_s3_path}.
	SslCertificateS3Path *KendraDataSourceDataSourceConfigurationSharePointConfigurationSslCertificateS3Path `field:"optional" json:"sslCertificateS3Path" yaml:"sslCertificateS3Path"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/kendra_data_source#urls KendraDataSource#urls}.
	Urls *[]*string `field:"optional" json:"urls" yaml:"urls"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/kendra_data_source#use_change_log KendraDataSource#use_change_log}.
	UseChangeLog interface{} `field:"optional" json:"useChangeLog" yaml:"useChangeLog"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/kendra_data_source#vpc_configuration KendraDataSource#vpc_configuration}.
	VpcConfiguration *KendraDataSourceDataSourceConfigurationSharePointConfigurationVpcConfiguration `field:"optional" json:"vpcConfiguration" yaml:"vpcConfiguration"`
}

