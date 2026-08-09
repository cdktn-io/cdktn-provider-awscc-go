// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package resiliencehubv2service


type Resiliencehubv2ServiceInputSourcesResourceConfigurationResourceTags struct {
	// Tag key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/resiliencehubv2_service#key Resiliencehubv2Service#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Tag values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/resiliencehubv2_service#values Resiliencehubv2Service#values}
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

