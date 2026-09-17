// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package casesdomain

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CasesDomainConfig struct {
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
	// The name for your Cases domain. It must be unique for your AWS account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cases_domain#name CasesDomain#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The tags that you attach to this domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cases_domain#tags CasesDomain#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

