// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package supportauthzsupportpermit

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SupportauthzSupportPermitConfig struct {
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
	// The name of the support permit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/supportauthz_support_permit#name SupportauthzSupportPermit#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The grant definition: which actions on which resources, optionally constrained by time conditions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/supportauthz_support_permit#permit SupportauthzSupportPermit#permit}
	Permit *SupportauthzSupportPermitPermit `field:"required" json:"permit" yaml:"permit"`
	// The signing key used by the permit. Exactly one key type must be provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/supportauthz_support_permit#signing_key_info SupportauthzSupportPermit#signing_key_info}
	SigningKeyInfo *SupportauthzSupportPermitSigningKeyInfo `field:"required" json:"signingKeyInfo" yaml:"signingKeyInfo"`
	// An optional description of the support permit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/supportauthz_support_permit#description SupportauthzSupportPermit#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The support case display identifier associated with the permit.
	//
	// When provided, the permit is linked to the specified AWS Support case.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/supportauthz_support_permit#support_case_display_id SupportauthzSupportPermit#support_case_display_id}
	SupportCaseDisplayId *string `field:"optional" json:"supportCaseDisplayId" yaml:"supportCaseDisplayId"`
	// A list of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/supportauthz_support_permit#tags SupportauthzSupportPermit#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

