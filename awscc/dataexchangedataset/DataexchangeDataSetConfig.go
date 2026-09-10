// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataexchangedataset

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataexchangeDataSetConfig struct {
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
	// The type of asset that is added to a data set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dataexchange_data_set#asset_type DataexchangeDataSet#asset_type}
	AssetType *string `field:"required" json:"assetType" yaml:"assetType"`
	// A description for the data set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dataexchange_data_set#description DataexchangeDataSet#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// The name of the data set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dataexchange_data_set#name DataexchangeDataSet#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Tags for the data set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dataexchange_data_set#tags DataexchangeDataSet#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

