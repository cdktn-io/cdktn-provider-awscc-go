// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package acmpcacertificate


type AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNames struct {
	// Contains information about the certificate subject.
	//
	// The certificate can be one issued by your private certificate authority (CA) or it can be your private CA certificate. The Subject field in the certificate identifies the entity that owns or controls the public key in the certificate. The entity can be a user, computer, device, or service. The Subject must contain an X.500 distinguished name (DN). A DN is a sequence of relative distinguished names (RDNs). The RDNs are separated by commas in the certificate. The DN must be unique for each entity, but your private CA can issue more than one certificate with the same DN to the same entity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/acmpca_certificate#directory_name AcmpcaCertificate#directory_name}
	DirectoryName *AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesDirectoryName `field:"optional" json:"directoryName" yaml:"directoryName"`
	// Represents ``GeneralName`` as a DNS name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/acmpca_certificate#dns_name AcmpcaCertificate#dns_name}
	DnsName *string `field:"optional" json:"dnsName" yaml:"dnsName"`
	// Represents ``GeneralName`` as an ``EdiPartyName`` object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/acmpca_certificate#edi_party_name AcmpcaCertificate#edi_party_name}
	EdiPartyName *AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesEdiPartyName `field:"optional" json:"ediPartyName" yaml:"ediPartyName"`
	// Represents ``GeneralName`` as an IPv4 or IPv6 address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/acmpca_certificate#ip_address AcmpcaCertificate#ip_address}
	IpAddress *string `field:"optional" json:"ipAddress" yaml:"ipAddress"`
	// Represents ``GeneralName`` using an ``OtherName`` object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/acmpca_certificate#other_name AcmpcaCertificate#other_name}
	OtherName *AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesOtherName `field:"optional" json:"otherName" yaml:"otherName"`
	// Represents ``GeneralName`` as an object identifier (OID).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/acmpca_certificate#registered_id AcmpcaCertificate#registered_id}
	RegisteredId *string `field:"optional" json:"registeredId" yaml:"registeredId"`
	// Represents ``GeneralName`` as an [RFC 822](https://docs.aws.amazon.com/https://datatracker.ietf.org/doc/html/rfc822) email address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/acmpca_certificate#rfc_822_name AcmpcaCertificate#rfc_822_name}
	Rfc822Name *string `field:"optional" json:"rfc822Name" yaml:"rfc822Name"`
	// Represents ``GeneralName`` as a URI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/acmpca_certificate#uniform_resource_identifier AcmpcaCertificate#uniform_resource_identifier}
	UniformResourceIdentifier *string `field:"optional" json:"uniformResourceIdentifier" yaml:"uniformResourceIdentifier"`
}

