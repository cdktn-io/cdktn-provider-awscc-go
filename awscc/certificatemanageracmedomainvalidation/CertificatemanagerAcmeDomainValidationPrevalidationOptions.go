// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package certificatemanageracmedomainvalidation


type CertificatemanagerAcmeDomainValidationPrevalidationOptions struct {
	// DNS-based prevalidation options for the domain validation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/certificatemanager_acme_domain_validation#dns_prevalidation CertificatemanagerAcmeDomainValidation#dns_prevalidation}
	DnsPrevalidation *CertificatemanagerAcmeDomainValidationPrevalidationOptionsDnsPrevalidation `field:"required" json:"dnsPrevalidation" yaml:"dnsPrevalidation"`
}

