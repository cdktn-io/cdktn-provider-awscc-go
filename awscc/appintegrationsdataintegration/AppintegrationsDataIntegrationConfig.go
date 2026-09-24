// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appintegrationsdataintegration

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AppintegrationsDataIntegrationConfig struct {
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
	// The KMS key of the data integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/appintegrations_data_integration#kms_key AppintegrationsDataIntegration#kms_key}
	KmsKey *string `field:"required" json:"kmsKey" yaml:"kmsKey"`
	// The name of the data integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/appintegrations_data_integration#name AppintegrationsDataIntegration#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The URI of the data source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/appintegrations_data_integration#source_uri AppintegrationsDataIntegration#source_uri}
	SourceUri *string `field:"required" json:"sourceUri" yaml:"sourceUri"`
	// The data integration description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/appintegrations_data_integration#description AppintegrationsDataIntegration#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The configuration for what files should be pulled from the source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/appintegrations_data_integration#file_configuration AppintegrationsDataIntegration#file_configuration}
	FileConfiguration *AppintegrationsDataIntegrationFileConfiguration `field:"optional" json:"fileConfiguration" yaml:"fileConfiguration"`
	// The configuration for what data should be pulled from the source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/appintegrations_data_integration#object_configuration AppintegrationsDataIntegration#object_configuration}
	ObjectConfiguration interface{} `field:"optional" json:"objectConfiguration" yaml:"objectConfiguration"`
	// The name of the data and how often it should be pulled from the source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/appintegrations_data_integration#schedule_config AppintegrationsDataIntegration#schedule_config}
	ScheduleConfig *AppintegrationsDataIntegrationScheduleConfig `field:"optional" json:"scheduleConfig" yaml:"scheduleConfig"`
	// The tags (keys and values) associated with the data integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/appintegrations_data_integration#tags AppintegrationsDataIntegration#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

