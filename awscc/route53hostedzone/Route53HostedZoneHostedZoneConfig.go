// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package route53hostedzone


type Route53HostedZoneHostedZoneConfig struct {
	// Any comments that you want to include about the hosted zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/route53_hosted_zone#comment Route53HostedZone#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
}

