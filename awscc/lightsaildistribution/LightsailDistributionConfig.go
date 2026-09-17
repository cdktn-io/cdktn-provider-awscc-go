// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lightsaildistribution

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LightsailDistributionConfig struct {
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
	// The bundle ID to use for the distribution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lightsail_distribution#bundle_id LightsailDistribution#bundle_id}
	BundleId *string `field:"required" json:"bundleId" yaml:"bundleId"`
	// An object that describes the default cache behavior for the distribution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lightsail_distribution#default_cache_behavior LightsailDistribution#default_cache_behavior}
	DefaultCacheBehavior *LightsailDistributionDefaultCacheBehavior `field:"required" json:"defaultCacheBehavior" yaml:"defaultCacheBehavior"`
	// The name for the distribution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lightsail_distribution#distribution_name LightsailDistribution#distribution_name}
	DistributionName *string `field:"required" json:"distributionName" yaml:"distributionName"`
	// An object that describes the origin resource for the distribution, such as a Lightsail instance or load balancer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lightsail_distribution#origin LightsailDistribution#origin}
	Origin *LightsailDistributionOrigin `field:"required" json:"origin" yaml:"origin"`
	// An array of objects that describe the per-path cache behavior for the distribution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lightsail_distribution#cache_behaviors LightsailDistribution#cache_behaviors}
	CacheBehaviors interface{} `field:"optional" json:"cacheBehaviors" yaml:"cacheBehaviors"`
	// An object that describes the cache behavior settings for the distribution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lightsail_distribution#cache_behavior_settings LightsailDistribution#cache_behavior_settings}
	CacheBehaviorSettings *LightsailDistributionCacheBehaviorSettings `field:"optional" json:"cacheBehaviorSettings" yaml:"cacheBehaviorSettings"`
	// The certificate attached to the Distribution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lightsail_distribution#certificate_name LightsailDistribution#certificate_name}
	CertificateName *string `field:"optional" json:"certificateName" yaml:"certificateName"`
	// The IP address type for the distribution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lightsail_distribution#ip_address_type LightsailDistribution#ip_address_type}
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// Indicates whether the distribution is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lightsail_distribution#is_enabled LightsailDistribution#is_enabled}
	IsEnabled interface{} `field:"optional" json:"isEnabled" yaml:"isEnabled"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lightsail_distribution#tags LightsailDistribution#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

