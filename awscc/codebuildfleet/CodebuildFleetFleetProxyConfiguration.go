// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package codebuildfleet


type CodebuildFleetFleetProxyConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/codebuild_fleet#default_behavior CodebuildFleet#default_behavior}.
	DefaultBehavior *string `field:"optional" json:"defaultBehavior" yaml:"defaultBehavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/codebuild_fleet#ordered_proxy_rules CodebuildFleet#ordered_proxy_rules}.
	OrderedProxyRules interface{} `field:"optional" json:"orderedProxyRules" yaml:"orderedProxyRules"`
}

