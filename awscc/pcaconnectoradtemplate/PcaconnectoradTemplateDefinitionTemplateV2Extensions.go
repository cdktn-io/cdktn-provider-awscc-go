// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pcaconnectoradtemplate


type PcaconnectoradTemplateDefinitionTemplateV2Extensions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/pcaconnectorad_template#application_policies PcaconnectoradTemplate#application_policies}.
	ApplicationPolicies *PcaconnectoradTemplateDefinitionTemplateV2ExtensionsApplicationPolicies `field:"optional" json:"applicationPolicies" yaml:"applicationPolicies"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/pcaconnectorad_template#key_usage PcaconnectoradTemplate#key_usage}.
	KeyUsage *PcaconnectoradTemplateDefinitionTemplateV2ExtensionsKeyUsage `field:"optional" json:"keyUsage" yaml:"keyUsage"`
}

