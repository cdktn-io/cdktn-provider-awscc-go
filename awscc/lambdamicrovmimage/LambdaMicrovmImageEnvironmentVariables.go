// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lambdamicrovmimage


type LambdaMicrovmImageEnvironmentVariables struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/lambda_microvm_image#key LambdaMicrovmImage#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/lambda_microvm_image#value LambdaMicrovmImage#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

