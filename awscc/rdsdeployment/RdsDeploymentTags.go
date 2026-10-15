// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package rdsdeployment


type RdsDeploymentTags struct {
	// The key name of the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/rds_deployment#key RdsDeployment#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value of the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/rds_deployment#value RdsDeployment#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

