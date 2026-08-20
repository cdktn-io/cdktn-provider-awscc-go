// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pcaconnectoradtemplate


type PcaconnectoradTemplateDefinitionTemplateV4ExtensionsKeyUsageUsageFlags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/pcaconnectorad_template#data_encipherment PcaconnectoradTemplate#data_encipherment}.
	DataEncipherment interface{} `field:"optional" json:"dataEncipherment" yaml:"dataEncipherment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/pcaconnectorad_template#digital_signature PcaconnectoradTemplate#digital_signature}.
	DigitalSignature interface{} `field:"optional" json:"digitalSignature" yaml:"digitalSignature"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/pcaconnectorad_template#key_agreement PcaconnectoradTemplate#key_agreement}.
	KeyAgreement interface{} `field:"optional" json:"keyAgreement" yaml:"keyAgreement"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/pcaconnectorad_template#key_encipherment PcaconnectoradTemplate#key_encipherment}.
	KeyEncipherment interface{} `field:"optional" json:"keyEncipherment" yaml:"keyEncipherment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/pcaconnectorad_template#non_repudiation PcaconnectoradTemplate#non_repudiation}.
	NonRepudiation interface{} `field:"optional" json:"nonRepudiation" yaml:"nonRepudiation"`
}

