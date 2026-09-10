// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityhubfindingaggregator

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SecurityhubFindingAggregatorConfig struct {
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
	// Indicates whether to aggregate findings from all of the available Regions in the current partition.
	//
	// Also determines whether to automatically aggregate findings from new Regions as Security Hub supports them and you opt into them.
	//  The selected option also determines how to use the Regions provided in the Regions list.
	//  In CFN, the options for this property are as follows:
	//   +  ``ALL_REGIONS`` - Indicates to aggregate findings from all of the Regions where Security Hub is enabled. When you choose this option, Security Hub also automatically aggregates findings from new Regions as Security Hub supports them and you opt into them.
	//   +  ``ALL_REGIONS_EXCEPT_SPECIFIED`` - Indicates to aggregate findings from all of the Regions where Security Hub is enabled, except for the Regions listed in the ``Regions`` parameter. When you choose this option, Security Hub also automatically aggregates findings from new Regions as Security Hub supports them and you opt into them.
	//   +  ``SPECIFIED_REGIONS`` - Indicates to aggregate findings only from the Regions listed in the ``Regions`` parameter. Security Hub does not automatically aggregate findings from new Regions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/securityhub_finding_aggregator#region_linking_mode SecurityhubFindingAggregator#region_linking_mode}
	RegionLinkingMode *string `field:"required" json:"regionLinkingMode" yaml:"regionLinkingMode"`
	// If ``RegionLinkingMode`` is ``ALL_REGIONS_EXCEPT_SPECIFIED``, then this is a space-separated list of Regions that do not aggregate findings to the aggregation Region.
	//
	// If ``RegionLinkingMode`` is ``SPECIFIED_REGIONS``, then this is a space-separated list of Regions that do aggregate findings to the aggregation Region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/securityhub_finding_aggregator#regions SecurityhubFindingAggregator#regions}
	Regions *[]*string `field:"optional" json:"regions" yaml:"regions"`
}

