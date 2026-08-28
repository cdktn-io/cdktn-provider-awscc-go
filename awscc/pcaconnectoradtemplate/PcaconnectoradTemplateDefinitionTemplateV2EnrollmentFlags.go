// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pcaconnectoradtemplate


type PcaconnectoradTemplateDefinitionTemplateV2EnrollmentFlags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/pcaconnectorad_template#enable_key_reuse_on_nt_token_keyset_storage_full PcaconnectoradTemplate#enable_key_reuse_on_nt_token_keyset_storage_full}.
	EnableKeyReuseOnNtTokenKeysetStorageFull interface{} `field:"optional" json:"enableKeyReuseOnNtTokenKeysetStorageFull" yaml:"enableKeyReuseOnNtTokenKeysetStorageFull"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/pcaconnectorad_template#include_symmetric_algorithms PcaconnectoradTemplate#include_symmetric_algorithms}.
	IncludeSymmetricAlgorithms interface{} `field:"optional" json:"includeSymmetricAlgorithms" yaml:"includeSymmetricAlgorithms"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/pcaconnectorad_template#no_security_extension PcaconnectoradTemplate#no_security_extension}.
	NoSecurityExtension interface{} `field:"optional" json:"noSecurityExtension" yaml:"noSecurityExtension"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/pcaconnectorad_template#remove_invalid_certificate_from_personal_store PcaconnectoradTemplate#remove_invalid_certificate_from_personal_store}.
	RemoveInvalidCertificateFromPersonalStore interface{} `field:"optional" json:"removeInvalidCertificateFromPersonalStore" yaml:"removeInvalidCertificateFromPersonalStore"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/pcaconnectorad_template#user_interaction_required PcaconnectoradTemplate#user_interaction_required}.
	UserInteractionRequired interface{} `field:"optional" json:"userInteractionRequired" yaml:"userInteractionRequired"`
}

