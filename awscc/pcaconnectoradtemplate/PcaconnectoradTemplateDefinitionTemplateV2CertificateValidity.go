// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pcaconnectoradtemplate


type PcaconnectoradTemplateDefinitionTemplateV2CertificateValidity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/pcaconnectorad_template#renewal_period PcaconnectoradTemplate#renewal_period}.
	RenewalPeriod *PcaconnectoradTemplateDefinitionTemplateV2CertificateValidityRenewalPeriod `field:"optional" json:"renewalPeriod" yaml:"renewalPeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/pcaconnectorad_template#validity_period PcaconnectoradTemplate#validity_period}.
	ValidityPeriod *PcaconnectoradTemplateDefinitionTemplateV2CertificateValidityValidityPeriod `field:"optional" json:"validityPeriod" yaml:"validityPeriod"`
}

