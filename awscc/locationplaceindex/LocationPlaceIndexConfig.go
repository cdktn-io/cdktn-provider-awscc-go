// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package locationplaceindex

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LocationPlaceIndexConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/location_place_index#data_source LocationPlaceIndex#data_source}.
	DataSource *string `field:"required" json:"dataSource" yaml:"dataSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/location_place_index#index_name LocationPlaceIndex#index_name}.
	IndexName *string `field:"required" json:"indexName" yaml:"indexName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/location_place_index#data_source_configuration LocationPlaceIndex#data_source_configuration}.
	DataSourceConfiguration *LocationPlaceIndexDataSourceConfiguration `field:"optional" json:"dataSourceConfiguration" yaml:"dataSourceConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/location_place_index#description LocationPlaceIndex#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/location_place_index#pricing_plan LocationPlaceIndex#pricing_plan}.
	PricingPlan *string `field:"optional" json:"pricingPlan" yaml:"pricingPlan"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/location_place_index#tags LocationPlaceIndex#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

