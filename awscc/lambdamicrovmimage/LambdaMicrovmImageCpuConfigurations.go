// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lambdamicrovmimage


type LambdaMicrovmImageCpuConfigurations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lambda_microvm_image#architecture LambdaMicrovmImage#architecture}.
	Architecture *string `field:"required" json:"architecture" yaml:"architecture"`
}

