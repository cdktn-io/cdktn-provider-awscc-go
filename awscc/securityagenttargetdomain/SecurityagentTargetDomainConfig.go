// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityagenttargetdomain

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SecurityagentTargetDomainConfig struct {
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
	// Domain name of the target domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/securityagent_target_domain#target_domain_name SecurityagentTargetDomain#target_domain_name}
	TargetDomainName *string `field:"required" json:"targetDomainName" yaml:"targetDomainName"`
	// Verification method for the target domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/securityagent_target_domain#verification_method SecurityagentTargetDomain#verification_method}
	VerificationMethod *string `field:"required" json:"verificationMethod" yaml:"verificationMethod"`
	// Tags for the target domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/securityagent_target_domain#tags SecurityagentTargetDomain#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

