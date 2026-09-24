// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pcaconnectoradtemplate


type PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyAttributesKeyUsagePropertyPropertyFlags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/pcaconnectorad_template#decrypt PcaconnectoradTemplate#decrypt}.
	Decrypt interface{} `field:"optional" json:"decrypt" yaml:"decrypt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/pcaconnectorad_template#key_agreement PcaconnectoradTemplate#key_agreement}.
	KeyAgreement interface{} `field:"optional" json:"keyAgreement" yaml:"keyAgreement"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/pcaconnectorad_template#sign PcaconnectoradTemplate#sign}.
	Sign interface{} `field:"optional" json:"sign" yaml:"sign"`
}

