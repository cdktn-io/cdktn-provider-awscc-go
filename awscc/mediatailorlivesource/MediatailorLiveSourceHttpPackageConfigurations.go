// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediatailorlivesource


type MediatailorLiveSourceHttpPackageConfigurations struct {
	// <p>The relative path to the URL for this VOD source.
	//
	// This is combined with <code>SourceLocation::HttpConfiguration::BaseUrl</code> to form a valid URL.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mediatailor_live_source#path MediatailorLiveSource#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// <p>The name of the source group. This has to match one of the <code>Channel::Outputs::SourceGroup</code>.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mediatailor_live_source#source_group MediatailorLiveSource#source_group}
	SourceGroup *string `field:"required" json:"sourceGroup" yaml:"sourceGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mediatailor_live_source#type MediatailorLiveSource#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

