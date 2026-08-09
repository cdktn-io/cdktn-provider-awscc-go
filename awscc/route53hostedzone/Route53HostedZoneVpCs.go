// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package route53hostedzone


type Route53HostedZoneVpCs struct {
	// *Private hosted zones only:* The ID of an Amazon VPC.
	//
	// For public hosted zones, omit ``VPCs``, ``VPCId``, and ``VPCRegion``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/route53_hosted_zone#vpc_id Route53HostedZone#vpc_id}
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
	// *Private hosted zones only:* The region that an Amazon VPC was created in.
	//
	// For public hosted zones, omit ``VPCs``, ``VPCId``, and ``VPCRegion``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/route53_hosted_zone#vpc_region Route53HostedZone#vpc_region}
	VpcRegion *string `field:"optional" json:"vpcRegion" yaml:"vpcRegion"`
}

