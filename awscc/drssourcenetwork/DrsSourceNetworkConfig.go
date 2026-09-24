// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package drssourcenetwork

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DrsSourceNetworkConfig struct {
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
	// The account ID containing the VPC to protect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/drs_source_network#origin_account_id DrsSourceNetwork#origin_account_id}
	OriginAccountId *string `field:"required" json:"originAccountId" yaml:"originAccountId"`
	// The region containing the VPC to protect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/drs_source_network#origin_region DrsSourceNetwork#origin_region}
	OriginRegion *string `field:"required" json:"originRegion" yaml:"originRegion"`
	// The VPC ID to protect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/drs_source_network#vpc_id DrsSourceNetwork#vpc_id}
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// A set of tags associated with the Source Network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/drs_source_network#tags DrsSourceNetwork#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

