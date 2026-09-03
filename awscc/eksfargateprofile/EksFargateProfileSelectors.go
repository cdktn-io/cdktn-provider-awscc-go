// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package eksfargateprofile


type EksFargateProfileSelectors struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/eks_fargate_profile#namespace EksFargateProfile#namespace}.
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/eks_fargate_profile#labels EksFargateProfile#labels}.
	Labels interface{} `field:"optional" json:"labels" yaml:"labels"`
}

