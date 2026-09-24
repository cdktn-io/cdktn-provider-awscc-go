// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package redshiftusagelimit


type RedshiftUsageLimitTags struct {
	// The key, or name, for the resource tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/redshift_usage_limit#key RedshiftUsageLimit#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value for the resource tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/redshift_usage_limit#value RedshiftUsageLimit#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

