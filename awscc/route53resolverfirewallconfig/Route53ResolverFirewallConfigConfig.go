// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package route53resolverfirewallconfig

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Route53ResolverFirewallConfigConfig struct {
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
	// Determines how DNS Firewall operates during failures.
	//
	// By default, fail open is disabled, which means the failure mode is closed. If you enable this option, the failure mode is open. This behavior is only enforced for VPCs that have at least one DNS Firewall rule group association.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/route53resolver_firewall_config#firewall_fail_open Route53ResolverFirewallConfig#firewall_fail_open}
	FirewallFailOpen *string `field:"required" json:"firewallFailOpen" yaml:"firewallFailOpen"`
	// The ID of the VPC that this firewall configuration applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/route53resolver_firewall_config#resource_id Route53ResolverFirewallConfig#resource_id}
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
}

