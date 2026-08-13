// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package acmpcacertificateauthority


type AcmpcaCertificateAuthorityCsrExtensions struct {
	// Structure that contains X.509 KeyUsage information.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/acmpca_certificate_authority#key_usage AcmpcaCertificateAuthority#key_usage}
	KeyUsage *AcmpcaCertificateAuthorityCsrExtensionsKeyUsage `field:"optional" json:"keyUsage" yaml:"keyUsage"`
	// Array of X.509 AccessDescription.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/acmpca_certificate_authority#subject_information_access AcmpcaCertificateAuthority#subject_information_access}
	SubjectInformationAccess interface{} `field:"optional" json:"subjectInformationAccess" yaml:"subjectInformationAccess"`
}

