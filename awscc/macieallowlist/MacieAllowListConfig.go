// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package macieallowlist

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MacieAllowListConfig struct {
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
	// AllowList criteria.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/macie_allow_list#criteria MacieAllowList#criteria}
	Criteria *MacieAllowListCriteria `field:"required" json:"criteria" yaml:"criteria"`
	// Name of AllowList.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/macie_allow_list#name MacieAllowList#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Description of AllowList.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/macie_allow_list#description MacieAllowList#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A collection of tags associated with a resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/macie_allow_list#tags MacieAllowList#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

