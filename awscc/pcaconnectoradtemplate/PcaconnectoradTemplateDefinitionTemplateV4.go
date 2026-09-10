// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pcaconnectoradtemplate


type PcaconnectoradTemplateDefinitionTemplateV4 struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/pcaconnectorad_template#certificate_validity PcaconnectoradTemplate#certificate_validity}.
	CertificateValidity *PcaconnectoradTemplateDefinitionTemplateV4CertificateValidity `field:"optional" json:"certificateValidity" yaml:"certificateValidity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/pcaconnectorad_template#enrollment_flags PcaconnectoradTemplate#enrollment_flags}.
	EnrollmentFlags *PcaconnectoradTemplateDefinitionTemplateV4EnrollmentFlags `field:"optional" json:"enrollmentFlags" yaml:"enrollmentFlags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/pcaconnectorad_template#extensions PcaconnectoradTemplate#extensions}.
	Extensions *PcaconnectoradTemplateDefinitionTemplateV4Extensions `field:"optional" json:"extensions" yaml:"extensions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/pcaconnectorad_template#general_flags PcaconnectoradTemplate#general_flags}.
	GeneralFlags *PcaconnectoradTemplateDefinitionTemplateV4GeneralFlags `field:"optional" json:"generalFlags" yaml:"generalFlags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/pcaconnectorad_template#hash_algorithm PcaconnectoradTemplate#hash_algorithm}.
	HashAlgorithm *string `field:"optional" json:"hashAlgorithm" yaml:"hashAlgorithm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/pcaconnectorad_template#private_key_attributes PcaconnectoradTemplate#private_key_attributes}.
	PrivateKeyAttributes *PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyAttributes `field:"optional" json:"privateKeyAttributes" yaml:"privateKeyAttributes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/pcaconnectorad_template#private_key_flags PcaconnectoradTemplate#private_key_flags}.
	PrivateKeyFlags *PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyFlags `field:"optional" json:"privateKeyFlags" yaml:"privateKeyFlags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/pcaconnectorad_template#subject_name_flags PcaconnectoradTemplate#subject_name_flags}.
	SubjectNameFlags *PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlags `field:"optional" json:"subjectNameFlags" yaml:"subjectNameFlags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/pcaconnectorad_template#superseded_templates PcaconnectoradTemplate#superseded_templates}.
	SupersededTemplates *[]*string `field:"optional" json:"supersededTemplates" yaml:"supersededTemplates"`
}

