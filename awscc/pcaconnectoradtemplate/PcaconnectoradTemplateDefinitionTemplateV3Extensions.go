// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pcaconnectoradtemplate


type PcaconnectoradTemplateDefinitionTemplateV3Extensions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/pcaconnectorad_template#application_policies PcaconnectoradTemplate#application_policies}.
	ApplicationPolicies *PcaconnectoradTemplateDefinitionTemplateV3ExtensionsApplicationPolicies `field:"optional" json:"applicationPolicies" yaml:"applicationPolicies"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/pcaconnectorad_template#key_usage PcaconnectoradTemplate#key_usage}.
	KeyUsage *PcaconnectoradTemplateDefinitionTemplateV3ExtensionsKeyUsage `field:"optional" json:"keyUsage" yaml:"keyUsage"`
}

