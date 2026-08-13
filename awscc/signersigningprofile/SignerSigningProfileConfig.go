// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package signersigningprofile

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SignerSigningProfileConfig struct {
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
	// The ID of the target signing platform.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/signer_signing_profile#platform_id SignerSigningProfile#platform_id}
	PlatformId *string `field:"required" json:"platformId" yaml:"platformId"`
	// Signature validity period of the profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/signer_signing_profile#signature_validity_period SignerSigningProfile#signature_validity_period}
	SignatureValidityPeriod *SignerSigningProfileSignatureValidityPeriod `field:"optional" json:"signatureValidityPeriod" yaml:"signatureValidityPeriod"`
	// A list of tags associated with the signing profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/signer_signing_profile#tags SignerSigningProfile#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

