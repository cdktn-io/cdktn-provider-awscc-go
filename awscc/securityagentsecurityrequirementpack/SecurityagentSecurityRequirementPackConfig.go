// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityagentsecurityrequirementpack

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SecurityagentSecurityRequirementPackConfig struct {
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
	// Name of the security requirement pack.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/securityagent_security_requirement_pack#name SecurityagentSecurityRequirementPack#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Description of the pack.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/securityagent_security_requirement_pack#description SecurityagentSecurityRequirementPack#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// KMS key for client-side encryption of pack contents.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/securityagent_security_requirement_pack#kms_key_id SecurityagentSecurityRequirementPack#kms_key_id}
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Security requirements within this pack.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/securityagent_security_requirement_pack#security_requirements SecurityagentSecurityRequirementPack#security_requirements}
	SecurityRequirements interface{} `field:"optional" json:"securityRequirements" yaml:"securityRequirements"`
	// Whether the pack is enabled or disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/securityagent_security_requirement_pack#status SecurityagentSecurityRequirementPack#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Tags for the security requirement pack.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/securityagent_security_requirement_pack#tags SecurityagentSecurityRequirementPack#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

