// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package route53recoveryreadinessresourceset


type Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResource struct {
	// The Network Load Balancer resource that a DNS target resource points to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/route53recoveryreadiness_resource_set#nlb_resource Route53RecoveryreadinessResourceSet#nlb_resource}
	NlbResource *Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResourceNlbResource `field:"optional" json:"nlbResource" yaml:"nlbResource"`
	// The Route 53 resource that a DNS target resource record points to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/route53recoveryreadiness_resource_set#r53_resource Route53RecoveryreadinessResourceSet#r53_resource}
	R53Resource *Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResourceR53Resource `field:"optional" json:"r53Resource" yaml:"r53Resource"`
}

