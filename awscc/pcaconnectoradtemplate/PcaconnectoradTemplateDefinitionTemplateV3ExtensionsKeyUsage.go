// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pcaconnectoradtemplate


type PcaconnectoradTemplateDefinitionTemplateV3ExtensionsKeyUsage struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/pcaconnectorad_template#critical PcaconnectoradTemplate#critical}.
	Critical interface{} `field:"optional" json:"critical" yaml:"critical"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/pcaconnectorad_template#usage_flags PcaconnectoradTemplate#usage_flags}.
	UsageFlags *PcaconnectoradTemplateDefinitionTemplateV3ExtensionsKeyUsageUsageFlags `field:"optional" json:"usageFlags" yaml:"usageFlags"`
}

