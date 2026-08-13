// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package acmpcacertificateauthority


type AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccess struct {
	// Structure that contains X.509 GeneralName information. Assign one and ONLY one field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/acmpca_certificate_authority#access_location AcmpcaCertificateAuthority#access_location}
	AccessLocation *AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocation `field:"optional" json:"accessLocation" yaml:"accessLocation"`
	// Structure that contains X.509 AccessMethod information. Assign one and ONLY one field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/acmpca_certificate_authority#access_method AcmpcaCertificateAuthority#access_method}
	AccessMethod *AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessMethod `field:"optional" json:"accessMethod" yaml:"accessMethod"`
}

