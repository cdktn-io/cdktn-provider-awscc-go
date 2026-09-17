// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package certificatemanagercertificate


type CertificatemanagerCertificateDomainValidationOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/certificatemanager_certificate#domain_name CertificatemanagerCertificate#domain_name}.
	DomainName *string `field:"optional" json:"domainName" yaml:"domainName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/certificatemanager_certificate#hosted_zone_id CertificatemanagerCertificate#hosted_zone_id}.
	HostedZoneId *string `field:"optional" json:"hostedZoneId" yaml:"hostedZoneId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/certificatemanager_certificate#validation_domain CertificatemanagerCertificate#validation_domain}.
	ValidationDomain *string `field:"optional" json:"validationDomain" yaml:"validationDomain"`
}

