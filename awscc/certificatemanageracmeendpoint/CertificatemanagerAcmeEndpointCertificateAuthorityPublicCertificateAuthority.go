// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package certificatemanageracmeendpoint


type CertificatemanagerAcmeEndpointCertificateAuthorityPublicCertificateAuthority struct {
	// The allowed key algorithms for certificates issued via this endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/certificatemanager_acme_endpoint#allowed_key_algorithms CertificatemanagerAcmeEndpoint#allowed_key_algorithms}
	AllowedKeyAlgorithms *[]*string `field:"optional" json:"allowedKeyAlgorithms" yaml:"allowedKeyAlgorithms"`
}

