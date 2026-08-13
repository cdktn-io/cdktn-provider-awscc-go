// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ivsplaybackrestrictionpolicy

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IvsPlaybackRestrictionPolicyConfig struct {
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
	// A list of country codes that control geoblocking restriction.
	//
	// Allowed values are the officially assigned ISO 3166-1 alpha-2 codes. Default: All countries (an empty array).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ivs_playback_restriction_policy#allowed_countries IvsPlaybackRestrictionPolicy#allowed_countries}
	AllowedCountries *[]*string `field:"optional" json:"allowedCountries" yaml:"allowedCountries"`
	// A list of origin sites that control CORS restriction.
	//
	// Allowed values are the same as valid values of the Origin header defined at https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Origin
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ivs_playback_restriction_policy#allowed_origins IvsPlaybackRestrictionPolicy#allowed_origins}
	AllowedOrigins *[]*string `field:"optional" json:"allowedOrigins" yaml:"allowedOrigins"`
	// Whether channel playback is constrained by origin site.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ivs_playback_restriction_policy#enable_strict_origin_enforcement IvsPlaybackRestrictionPolicy#enable_strict_origin_enforcement}
	EnableStrictOriginEnforcement interface{} `field:"optional" json:"enableStrictOriginEnforcement" yaml:"enableStrictOriginEnforcement"`
	// Playback-restriction-policy name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ivs_playback_restriction_policy#name IvsPlaybackRestrictionPolicy#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ivs_playback_restriction_policy#tags IvsPlaybackRestrictionPolicy#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

