// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityhubaggregatorv2

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SecurityhubAggregatorV2Config struct {
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
	// The list of included Regions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/securityhub_aggregator_v2#linked_regions SecurityhubAggregatorV2#linked_regions}
	LinkedRegions *[]*string `field:"required" json:"linkedRegions" yaml:"linkedRegions"`
	// Indicates to link a list of included Regions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/securityhub_aggregator_v2#region_linking_mode SecurityhubAggregatorV2#region_linking_mode}
	RegionLinkingMode *string `field:"required" json:"regionLinkingMode" yaml:"regionLinkingMode"`
	// A key-value pair to associate with the Security Hub V2 resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/securityhub_aggregator_v2#tags SecurityhubAggregatorV2#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

