// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lambdamicrovmimage


type LambdaMicrovmImageResources struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lambda_microvm_image#minimum_memory_in_mi_b LambdaMicrovmImage#minimum_memory_in_mi_b}.
	MinimumMemoryInMiB *float64 `field:"required" json:"minimumMemoryInMiB" yaml:"minimumMemoryInMiB"`
}

