// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apprunnerservice


type ApprunnerServiceSourceConfigurationCodeRepositorySourceCodeVersion struct {
	// Source Code Version Type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/apprunner_service#type ApprunnerService#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Source Code Version Value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/apprunner_service#value ApprunnerService#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

