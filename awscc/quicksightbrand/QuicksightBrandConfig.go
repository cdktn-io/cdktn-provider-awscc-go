// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightbrand

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type QuicksightBrandConfig struct {
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
	// The ID of the QuickSight brand.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/quicksight_brand#brand_id QuicksightBrand#brand_id}
	BrandId *string `field:"required" json:"brandId" yaml:"brandId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/quicksight_brand#brand_definition QuicksightBrand#brand_definition}.
	BrandDefinition *QuicksightBrandBrandDefinition `field:"optional" json:"brandDefinition" yaml:"brandDefinition"`
	// The tags for the brand.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/quicksight_brand#tags QuicksightBrand#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

