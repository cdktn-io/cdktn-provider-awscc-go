// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package route53recordset


type Route53RecordSetAliasTarget struct {
	// The value that you specify depends on where you want to route queries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/route53_record_set#dns_name Route53RecordSet#dns_name}
	DnsName *string `field:"optional" json:"dnsName" yaml:"dnsName"`
	// When EvaluateTargetHealth is true, an alias resource record set inherits the health of the referenced AWS resource, such as an ELB load balancer or another resource record set in the hosted zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/route53_record_set#evaluate_target_health Route53RecordSet#evaluate_target_health}
	EvaluateTargetHealth interface{} `field:"optional" json:"evaluateTargetHealth" yaml:"evaluateTargetHealth"`
	// The value used depends on where you want to route traffic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/route53_record_set#hosted_zone_id Route53RecordSet#hosted_zone_id}
	HostedZoneId *string `field:"optional" json:"hostedZoneId" yaml:"hostedZoneId"`
}

