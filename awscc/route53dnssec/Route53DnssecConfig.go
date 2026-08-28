// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package route53dnssec

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Route53DnssecConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The unique string (ID) used to identify a hosted zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/route53_dnssec#hosted_zone_id Route53Dnssec#hosted_zone_id}
	HostedZoneId *string `field:"required" json:"hostedZoneId" yaml:"hostedZoneId"`
}

