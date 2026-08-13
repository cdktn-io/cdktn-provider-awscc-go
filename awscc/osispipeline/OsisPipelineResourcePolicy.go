// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package osispipeline


type OsisPipelineResourcePolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/osis_pipeline#policy OsisPipeline#policy}.
	Policy *string `field:"optional" json:"policy" yaml:"policy"`
}

