// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package certificatemanageracmedomainvalidation


type CertificatemanagerAcmeDomainValidationPrevalidationOptionsDnsPrevalidation struct {
	// Controls which certificate types are authorized to be issued for the domain via the ACME endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/certificatemanager_acme_domain_validation#domain_scope CertificatemanagerAcmeDomainValidation#domain_scope}
	DomainScope *CertificatemanagerAcmeDomainValidationPrevalidationOptionsDnsPrevalidationDomainScope `field:"optional" json:"domainScope" yaml:"domainScope"`
	// The Route 53 hosted zone ID for automatic DNS record management.
	//
	// When provided, the service creates the validation DNS record on the customer's behalf.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/certificatemanager_acme_domain_validation#hosted_zone_id CertificatemanagerAcmeDomainValidation#hosted_zone_id}
	HostedZoneId *string `field:"optional" json:"hostedZoneId" yaml:"hostedZoneId"`
}

