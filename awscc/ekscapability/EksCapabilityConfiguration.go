// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ekscapability


type EksCapabilityConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/eks_capability#ack EksCapability#ack}.
	Ack *string `field:"optional" json:"ack" yaml:"ack"`
	// Configuration settings for an Argo CD capability.
	//
	// This includes the Kubernetes namespace, IAM Identity Center integration, RBAC role mappings, and network access configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/eks_capability#argo_cd EksCapability#argo_cd}
	ArgoCd *EksCapabilityConfigurationArgoCd `field:"optional" json:"argoCd" yaml:"argoCd"`
}

