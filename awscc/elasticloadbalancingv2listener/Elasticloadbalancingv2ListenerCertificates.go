// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package elasticloadbalancingv2listener


type Elasticloadbalancingv2ListenerCertificates struct {
	// The Amazon Resource Name (ARN) of the certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/elasticloadbalancingv2_listener#certificate_arn Elasticloadbalancingv2Listener#certificate_arn}
	CertificateArn *string `field:"optional" json:"certificateArn" yaml:"certificateArn"`
}

