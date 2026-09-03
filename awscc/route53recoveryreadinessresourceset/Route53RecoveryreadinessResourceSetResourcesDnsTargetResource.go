// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package route53recoveryreadinessresourceset


type Route53RecoveryreadinessResourceSetResourcesDnsTargetResource struct {
	// The domain name that acts as an ingress point to a portion of the customer application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/route53recoveryreadiness_resource_set#domain_name Route53RecoveryreadinessResourceSet#domain_name}
	DomainName *string `field:"optional" json:"domainName" yaml:"domainName"`
	// The hosted zone Amazon Resource Name (ARN) that contains the DNS record with the provided name of the target resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/route53recoveryreadiness_resource_set#hosted_zone_arn Route53RecoveryreadinessResourceSet#hosted_zone_arn}
	HostedZoneArn *string `field:"optional" json:"hostedZoneArn" yaml:"hostedZoneArn"`
	// The Route 53 record set ID that will uniquely identify a DNS record, given a name and a type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/route53recoveryreadiness_resource_set#record_set_id Route53RecoveryreadinessResourceSet#record_set_id}
	RecordSetId *string `field:"optional" json:"recordSetId" yaml:"recordSetId"`
	// The type of DNS record of the target resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/route53recoveryreadiness_resource_set#record_type Route53RecoveryreadinessResourceSet#record_type}
	RecordType *string `field:"optional" json:"recordType" yaml:"recordType"`
	// The target resource that the Route 53 record points to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/route53recoveryreadiness_resource_set#target_resource Route53RecoveryreadinessResourceSet#target_resource}
	TargetResource *Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResource `field:"optional" json:"targetResource" yaml:"targetResource"`
}

