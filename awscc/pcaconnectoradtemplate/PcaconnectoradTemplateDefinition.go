// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pcaconnectoradtemplate


type PcaconnectoradTemplateDefinition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/pcaconnectorad_template#template_v2 PcaconnectoradTemplate#template_v2}.
	TemplateV2 *PcaconnectoradTemplateDefinitionTemplateV2 `field:"optional" json:"templateV2" yaml:"templateV2"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/pcaconnectorad_template#template_v3 PcaconnectoradTemplate#template_v3}.
	TemplateV3 *PcaconnectoradTemplateDefinitionTemplateV3 `field:"optional" json:"templateV3" yaml:"templateV3"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/pcaconnectorad_template#template_v4 PcaconnectoradTemplate#template_v4}.
	TemplateV4 *PcaconnectoradTemplateDefinitionTemplateV4 `field:"optional" json:"templateV4" yaml:"templateV4"`
}

