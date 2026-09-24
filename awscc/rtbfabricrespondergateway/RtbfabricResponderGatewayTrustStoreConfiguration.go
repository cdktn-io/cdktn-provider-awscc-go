// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package rtbfabricrespondergateway


type RtbfabricResponderGatewayTrustStoreConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/rtbfabric_responder_gateway#certificate_authority_certificates RtbfabricResponderGateway#certificate_authority_certificates}.
	CertificateAuthorityCertificates *[]*string `field:"optional" json:"certificateAuthorityCertificates" yaml:"certificateAuthorityCertificates"`
}

