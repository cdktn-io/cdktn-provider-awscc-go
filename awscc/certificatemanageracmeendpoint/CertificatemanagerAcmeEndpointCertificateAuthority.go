// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package certificatemanageracmeendpoint


type CertificatemanagerAcmeEndpointCertificateAuthority struct {
	// Configuration for the public certificate authority.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/certificatemanager_acme_endpoint#public_certificate_authority CertificatemanagerAcmeEndpoint#public_certificate_authority}
	PublicCertificateAuthority *CertificatemanagerAcmeEndpointCertificateAuthorityPublicCertificateAuthority `field:"required" json:"publicCertificateAuthority" yaml:"publicCertificateAuthority"`
}

