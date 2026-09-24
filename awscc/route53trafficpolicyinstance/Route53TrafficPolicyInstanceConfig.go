// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package route53trafficpolicyinstance

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Route53TrafficPolicyInstanceConfig struct {
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
	// The ID of the hosted zone that Amazon Route 53 creates the resource record sets in.
	//
	// The bare ID as Route 53 returns it, without a /hostedzone/ prefix.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/route53_traffic_policy_instance#hosted_zone_id Route53TrafficPolicyInstance#hosted_zone_id}
	HostedZoneId *string `field:"required" json:"hostedZoneId" yaml:"hostedZoneId"`
	// The domain name, or subdomain name, for which Amazon Route 53 answers DNS queries by using the resource record sets it creates for this traffic policy instance.
	//
	// Must be lower-case and end with a trailing dot, which is the form Route 53 returns: Route 53 normalizes DNS names, so admitting another form would neither round-trip through Read nor guarantee that a change to this property changes the resource's identity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/route53_traffic_policy_instance#name Route53TrafficPolicyInstance#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ID of the traffic policy that Amazon Route 53 uses to create resource record sets in the specified hosted zone.
	//
	// Lower-case, as Route 53 returns it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/route53_traffic_policy_instance#traffic_policy_id Route53TrafficPolicyInstance#traffic_policy_id}
	TrafficPolicyId *string `field:"required" json:"trafficPolicyId" yaml:"trafficPolicyId"`
	// The version of the traffic policy that Amazon Route 53 uses to create resource record sets in the specified hosted zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/route53_traffic_policy_instance#traffic_policy_version Route53TrafficPolicyInstance#traffic_policy_version}
	TrafficPolicyVersion *float64 `field:"required" json:"trafficPolicyVersion" yaml:"trafficPolicyVersion"`
	// The TTL that Amazon Route 53 assigns to all of the resource record sets that it creates in the specified hosted zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/route53_traffic_policy_instance#ttl Route53TrafficPolicyInstance#ttl}
	Ttl *float64 `field:"required" json:"ttl" yaml:"ttl"`
}

