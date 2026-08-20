// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lightsaildistribution


type LightsailDistributionDefaultCacheBehavior struct {
	// The cache behavior of the distribution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/lightsail_distribution#behavior LightsailDistribution#behavior}
	Behavior *string `field:"optional" json:"behavior" yaml:"behavior"`
}

