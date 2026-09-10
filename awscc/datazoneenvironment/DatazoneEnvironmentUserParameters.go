// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datazoneenvironment


type DatazoneEnvironmentUserParameters struct {
	// The name of an environment parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/datazone_environment#name DatazoneEnvironment#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The value of an environment parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/datazone_environment#value DatazoneEnvironment#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

