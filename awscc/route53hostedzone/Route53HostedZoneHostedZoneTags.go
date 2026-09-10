// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package route53hostedzone


type Route53HostedZoneHostedZoneTags struct {
	// The value of ``Key`` depends on the operation that you want to perform:   +  *Add a tag to a health check or hosted zone*: ``Key`` is the name that you want to give the new tag.
	//
	// +  *Edit a tag*: ``Key`` is the name of the tag that you want to change the ``Value`` for.
	//   +  *Delete a key*: ``Key`` is the name of the tag you want to remove.
	//   +  *Give a name to a health check*: Edit the default ``Name`` tag. In the Amazon Route 53 console, the list of your health checks includes a *Name* column that lets you see the name that you've given to each health check.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/route53_hosted_zone#key Route53HostedZone#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value of ``Value`` depends on the operation that you want to perform:   +  *Add a tag to a health check or hosted zone*: ``Value`` is the value that you want to give the new tag.
	//
	// +  *Edit a tag*: ``Value`` is the new value that you want to assign the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/route53_hosted_zone#value Route53HostedZone#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

