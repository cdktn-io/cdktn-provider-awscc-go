// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datazonedatasource

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DatazoneDataSourceConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The ID of the Amazon DataZone domain where the data source is created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/datazone_data_source#domain_identifier DatazoneDataSource#domain_identifier}
	DomainIdentifier *string `field:"required" json:"domainIdentifier" yaml:"domainIdentifier"`
	// The name of the data source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/datazone_data_source#name DatazoneDataSource#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The identifier of the Amazon DataZone project in which you want to add the data source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/datazone_data_source#project_identifier DatazoneDataSource#project_identifier}
	ProjectIdentifier *string `field:"required" json:"projectIdentifier" yaml:"projectIdentifier"`
	// The type of the data source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/datazone_data_source#type DatazoneDataSource#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The metadata forms that are to be attached to the assets that this data source works with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/datazone_data_source#asset_forms_input DatazoneDataSource#asset_forms_input}
	AssetFormsInput interface{} `field:"optional" json:"assetFormsInput" yaml:"assetFormsInput"`
	// Configuration of the data source. It can be set to either glueRunConfiguration or redshiftRunConfiguration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/datazone_data_source#configuration DatazoneDataSource#configuration}
	Configuration *DatazoneDataSourceConfiguration `field:"optional" json:"configuration" yaml:"configuration"`
	// The unique identifier of a connection used to fetch relevant parameters from connection during Datasource run.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/datazone_data_source#connection_identifier DatazoneDataSource#connection_identifier}
	ConnectionIdentifier *string `field:"optional" json:"connectionIdentifier" yaml:"connectionIdentifier"`
	// The description of the data source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/datazone_data_source#description DatazoneDataSource#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies whether the data source is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/datazone_data_source#enable_setting DatazoneDataSource#enable_setting}
	EnableSetting *string `field:"optional" json:"enableSetting" yaml:"enableSetting"`
	// The unique identifier of the Amazon DataZone environment to which the data source publishes assets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/datazone_data_source#environment_identifier DatazoneDataSource#environment_identifier}
	EnvironmentIdentifier *string `field:"optional" json:"environmentIdentifier" yaml:"environmentIdentifier"`
	// Specifies whether the assets that this data source creates in the inventory are to be also automatically published to the catalog.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/datazone_data_source#publish_on_import DatazoneDataSource#publish_on_import}
	PublishOnImport interface{} `field:"optional" json:"publishOnImport" yaml:"publishOnImport"`
	// Specifies whether the business name generation is to be enabled for this data source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/datazone_data_source#recommendation DatazoneDataSource#recommendation}
	Recommendation *DatazoneDataSourceRecommendation `field:"optional" json:"recommendation" yaml:"recommendation"`
	// The schedule of the data source runs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/datazone_data_source#schedule DatazoneDataSource#schedule}
	Schedule *DatazoneDataSourceSchedule `field:"optional" json:"schedule" yaml:"schedule"`
}

