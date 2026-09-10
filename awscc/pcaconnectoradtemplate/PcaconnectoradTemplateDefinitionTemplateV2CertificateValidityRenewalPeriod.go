// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pcaconnectoradtemplate


type PcaconnectoradTemplateDefinitionTemplateV2CertificateValidityRenewalPeriod struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/pcaconnectorad_template#period PcaconnectoradTemplate#period}.
	Period *float64 `field:"optional" json:"period" yaml:"period"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/pcaconnectorad_template#period_type PcaconnectoradTemplate#period_type}.
	PeriodType *string `field:"optional" json:"periodType" yaml:"periodType"`
}

