// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package acmpcacertificateauthority


type AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocation struct {
	// Structure that contains X.500 distinguished name information for your CA.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/acmpca_certificate_authority#directory_name AcmpcaCertificateAuthority#directory_name}
	DirectoryName *AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationDirectoryName `field:"optional" json:"directoryName" yaml:"directoryName"`
	// String that contains X.509 DnsName information.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/acmpca_certificate_authority#dns_name AcmpcaCertificateAuthority#dns_name}
	DnsName *string `field:"optional" json:"dnsName" yaml:"dnsName"`
	// Structure that contains X.509 EdiPartyName information.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/acmpca_certificate_authority#edi_party_name AcmpcaCertificateAuthority#edi_party_name}
	EdiPartyName *AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationEdiPartyName `field:"optional" json:"ediPartyName" yaml:"ediPartyName"`
	// String that contains X.509 IpAddress information.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/acmpca_certificate_authority#ip_address AcmpcaCertificateAuthority#ip_address}
	IpAddress *string `field:"optional" json:"ipAddress" yaml:"ipAddress"`
	// Structure that contains X.509 OtherName information.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/acmpca_certificate_authority#other_name AcmpcaCertificateAuthority#other_name}
	OtherName *AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOtherName `field:"optional" json:"otherName" yaml:"otherName"`
	// String that contains X.509 ObjectIdentifier information.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/acmpca_certificate_authority#registered_id AcmpcaCertificateAuthority#registered_id}
	RegisteredId *string `field:"optional" json:"registeredId" yaml:"registeredId"`
	// String that contains X.509 Rfc822Name information.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/acmpca_certificate_authority#rfc_822_name AcmpcaCertificateAuthority#rfc_822_name}
	Rfc822Name *string `field:"optional" json:"rfc822Name" yaml:"rfc822Name"`
	// String that contains X.509 UniformResourceIdentifier information.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/acmpca_certificate_authority#uniform_resource_identifier AcmpcaCertificateAuthority#uniform_resource_identifier}
	UniformResourceIdentifier *string `field:"optional" json:"uniformResourceIdentifier" yaml:"uniformResourceIdentifier"`
}

