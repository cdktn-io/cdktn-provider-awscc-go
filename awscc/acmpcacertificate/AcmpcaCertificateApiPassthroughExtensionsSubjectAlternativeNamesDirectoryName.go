// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package acmpcacertificate


type AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesDirectoryName struct {
	// For CA and end-entity certificates in a private PKI, the common name (CN) can be any string within the length limit.
	//
	// Note: In publicly trusted certificates, the common name must be a fully qualified domain name (FQDN) associated with the certificate subject.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/acmpca_certificate#common_name AcmpcaCertificate#common_name}
	CommonName *string `field:"optional" json:"commonName" yaml:"commonName"`
	// Two-digit code that specifies the country in which the certificate subject located.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/acmpca_certificate#country AcmpcaCertificate#country}
	Country *string `field:"optional" json:"country" yaml:"country"`
	// Contains a sequence of one or more X.500 relative distinguished names (RDNs), each of which consists of an object identifier (OID) and a value. For more information, see NIST?s definition of [Object Identifier (OID)](https://docs.aws.amazon.com/https://csrc.nist.gov/glossary/term/Object_Identifier).   Custom attributes cannot be used in combination with standard attributes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/acmpca_certificate#custom_attributes AcmpcaCertificate#custom_attributes}
	CustomAttributes interface{} `field:"optional" json:"customAttributes" yaml:"customAttributes"`
	// Disambiguating information for the certificate subject.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/acmpca_certificate#distinguished_name_qualifier AcmpcaCertificate#distinguished_name_qualifier}
	DistinguishedNameQualifier *string `field:"optional" json:"distinguishedNameQualifier" yaml:"distinguishedNameQualifier"`
	// Typically a qualifier appended to the name of an individual.
	//
	// Examples include Jr. for junior, Sr. for senior, and III for third.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/acmpca_certificate#generation_qualifier AcmpcaCertificate#generation_qualifier}
	GenerationQualifier *string `field:"optional" json:"generationQualifier" yaml:"generationQualifier"`
	// First name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/acmpca_certificate#given_name AcmpcaCertificate#given_name}
	GivenName *string `field:"optional" json:"givenName" yaml:"givenName"`
	// Concatenation that typically contains the first letter of the *GivenName*, the first letter of the middle name if one exists, and the first letter of the *Surname*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/acmpca_certificate#initials AcmpcaCertificate#initials}
	Initials *string `field:"optional" json:"initials" yaml:"initials"`
	// The locality (such as a city or town) in which the certificate subject is located.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/acmpca_certificate#locality AcmpcaCertificate#locality}
	Locality *string `field:"optional" json:"locality" yaml:"locality"`
	// Legal name of the organization with which the certificate subject is affiliated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/acmpca_certificate#organization AcmpcaCertificate#organization}
	Organization *string `field:"optional" json:"organization" yaml:"organization"`
	// A subdivision or unit of the organization (such as sales or finance) with which the certificate subject is affiliated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/acmpca_certificate#organizational_unit AcmpcaCertificate#organizational_unit}
	OrganizationalUnit *string `field:"optional" json:"organizationalUnit" yaml:"organizationalUnit"`
	// Typically a shortened version of a longer *GivenName*.
	//
	// For example, Jonathan is often shortened to John. Elizabeth is often shortened to Beth, Liz, or Eliza.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/acmpca_certificate#pseudonym AcmpcaCertificate#pseudonym}
	Pseudonym *string `field:"optional" json:"pseudonym" yaml:"pseudonym"`
	// The certificate serial number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/acmpca_certificate#serial_number AcmpcaCertificate#serial_number}
	SerialNumber *string `field:"optional" json:"serialNumber" yaml:"serialNumber"`
	// State in which the subject of the certificate is located.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/acmpca_certificate#state AcmpcaCertificate#state}
	State *string `field:"optional" json:"state" yaml:"state"`
	// Family name.
	//
	// In the US and the UK, for example, the surname of an individual is ordered last. In Asian cultures the surname is typically ordered first.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/acmpca_certificate#surname AcmpcaCertificate#surname}
	Surname *string `field:"optional" json:"surname" yaml:"surname"`
	// A title such as Mr.
	//
	// or Ms., which is pre-pended to the name to refer formally to the certificate subject.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/acmpca_certificate#title AcmpcaCertificate#title}
	Title *string `field:"optional" json:"title" yaml:"title"`
}

