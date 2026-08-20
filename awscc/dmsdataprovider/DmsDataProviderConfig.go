// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dmsdataprovider

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DmsDataProviderConfig struct {
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
	// The property describes a data engine for the data provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/dms_data_provider#engine DmsDataProvider#engine}
	Engine *string `field:"required" json:"engine" yaml:"engine"`
	// The property describes an identifier for the data provider. It is used for describing/deleting/modifying can be name/arn.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/dms_data_provider#data_provider_identifier DmsDataProvider#data_provider_identifier}
	DataProviderIdentifier *string `field:"optional" json:"dataProviderIdentifier" yaml:"dataProviderIdentifier"`
	// The property describes a name to identify the data provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/dms_data_provider#data_provider_name DmsDataProvider#data_provider_name}
	DataProviderName *string `field:"optional" json:"dataProviderName" yaml:"dataProviderName"`
	// The optional description of the data provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/dms_data_provider#description DmsDataProvider#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The property describes the exact settings which can be modified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/dms_data_provider#exact_settings DmsDataProvider#exact_settings}
	ExactSettings interface{} `field:"optional" json:"exactSettings" yaml:"exactSettings"`
	// The property identifies the exact type of settings for the data provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/dms_data_provider#settings DmsDataProvider#settings}
	Settings *DmsDataProviderSettings `field:"optional" json:"settings" yaml:"settings"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/dms_data_provider#tags DmsDataProvider#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

