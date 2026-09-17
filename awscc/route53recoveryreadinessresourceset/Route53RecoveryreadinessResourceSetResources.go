// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package route53recoveryreadinessresourceset


type Route53RecoveryreadinessResourceSetResources struct {
	// The component identifier of the resource, generated when DNS target resource is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/route53recoveryreadiness_resource_set#component_id Route53RecoveryreadinessResourceSet#component_id}
	ComponentId *string `field:"optional" json:"componentId" yaml:"componentId"`
	// A component for DNS/routing control readiness checks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/route53recoveryreadiness_resource_set#dns_target_resource Route53RecoveryreadinessResourceSet#dns_target_resource}
	DnsTargetResource *Route53RecoveryreadinessResourceSetResourcesDnsTargetResource `field:"optional" json:"dnsTargetResource" yaml:"dnsTargetResource"`
	// A list of recovery group Amazon Resource Names (ARNs) and cell ARNs that this resource is contained within.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/route53recoveryreadiness_resource_set#readiness_scopes Route53RecoveryreadinessResourceSet#readiness_scopes}
	ReadinessScopes *[]*string `field:"optional" json:"readinessScopes" yaml:"readinessScopes"`
	// The Amazon Resource Name (ARN) of the AWS resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/route53recoveryreadiness_resource_set#resource_arn Route53RecoveryreadinessResourceSet#resource_arn}
	ResourceArn *string `field:"optional" json:"resourceArn" yaml:"resourceArn"`
}

