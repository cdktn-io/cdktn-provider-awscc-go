// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2clientvpnendpoint


type Ec2ClientVpnEndpointAuthenticationOptionsMutualAuthentication struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/ec2_client_vpn_endpoint#client_root_certificate_chain_arn Ec2ClientVpnEndpoint#client_root_certificate_chain_arn}.
	ClientRootCertificateChainArn *string `field:"optional" json:"clientRootCertificateChainArn" yaml:"clientRootCertificateChainArn"`
}

