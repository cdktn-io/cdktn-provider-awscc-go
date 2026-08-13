// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package certificatemanageracmedomainvalidation


type CertificatemanagerAcmeDomainValidationPrevalidationOptionsDnsPrevalidationDomainScope struct {
	// Whether certificates may be issued for the exact domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/certificatemanager_acme_domain_validation#exact_domain CertificatemanagerAcmeDomainValidation#exact_domain}
	ExactDomain *string `field:"optional" json:"exactDomain" yaml:"exactDomain"`
	// Whether certificates may be issued for subdomains of the domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/certificatemanager_acme_domain_validation#subdomains CertificatemanagerAcmeDomainValidation#subdomains}
	Subdomains *string `field:"optional" json:"subdomains" yaml:"subdomains"`
	// Whether wildcard certificates may be issued for the domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/certificatemanager_acme_domain_validation#wildcards CertificatemanagerAcmeDomainValidation#wildcards}
	Wildcards *string `field:"optional" json:"wildcards" yaml:"wildcards"`
}

