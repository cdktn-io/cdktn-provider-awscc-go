// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package certificatemanagercertificate


type CertificatemanagerCertificateTags struct {
	// The tag's key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/certificatemanager_certificate#key CertificatemanagerCertificate#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The tag's value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/certificatemanager_certificate#value CertificatemanagerCertificate#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

