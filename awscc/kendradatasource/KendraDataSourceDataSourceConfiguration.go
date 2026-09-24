// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kendradatasource


type KendraDataSourceDataSourceConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/kendra_data_source#confluence_configuration KendraDataSource#confluence_configuration}.
	ConfluenceConfiguration *KendraDataSourceDataSourceConfigurationConfluenceConfiguration `field:"optional" json:"confluenceConfiguration" yaml:"confluenceConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/kendra_data_source#database_configuration KendraDataSource#database_configuration}.
	DatabaseConfiguration *KendraDataSourceDataSourceConfigurationDatabaseConfiguration `field:"optional" json:"databaseConfiguration" yaml:"databaseConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/kendra_data_source#google_drive_configuration KendraDataSource#google_drive_configuration}.
	GoogleDriveConfiguration *KendraDataSourceDataSourceConfigurationGoogleDriveConfiguration `field:"optional" json:"googleDriveConfiguration" yaml:"googleDriveConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/kendra_data_source#one_drive_configuration KendraDataSource#one_drive_configuration}.
	OneDriveConfiguration *KendraDataSourceDataSourceConfigurationOneDriveConfiguration `field:"optional" json:"oneDriveConfiguration" yaml:"oneDriveConfiguration"`
	// S3 data source configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/kendra_data_source#s3_configuration KendraDataSource#s3_configuration}
	S3Configuration *KendraDataSourceDataSourceConfigurationS3Configuration `field:"optional" json:"s3Configuration" yaml:"s3Configuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/kendra_data_source#salesforce_configuration KendraDataSource#salesforce_configuration}.
	SalesforceConfiguration *KendraDataSourceDataSourceConfigurationSalesforceConfiguration `field:"optional" json:"salesforceConfiguration" yaml:"salesforceConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/kendra_data_source#service_now_configuration KendraDataSource#service_now_configuration}.
	ServiceNowConfiguration *KendraDataSourceDataSourceConfigurationServiceNowConfiguration `field:"optional" json:"serviceNowConfiguration" yaml:"serviceNowConfiguration"`
	// SharePoint configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/kendra_data_source#share_point_configuration KendraDataSource#share_point_configuration}
	SharePointConfiguration *KendraDataSourceDataSourceConfigurationSharePointConfiguration `field:"optional" json:"sharePointConfiguration" yaml:"sharePointConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/kendra_data_source#template_configuration KendraDataSource#template_configuration}.
	TemplateConfiguration *KendraDataSourceDataSourceConfigurationTemplateConfiguration `field:"optional" json:"templateConfiguration" yaml:"templateConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/kendra_data_source#web_crawler_configuration KendraDataSource#web_crawler_configuration}.
	WebCrawlerConfiguration *KendraDataSourceDataSourceConfigurationWebCrawlerConfiguration `field:"optional" json:"webCrawlerConfiguration" yaml:"webCrawlerConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/kendra_data_source#work_docs_configuration KendraDataSource#work_docs_configuration}.
	WorkDocsConfiguration *KendraDataSourceDataSourceConfigurationWorkDocsConfiguration `field:"optional" json:"workDocsConfiguration" yaml:"workDocsConfiguration"`
}

