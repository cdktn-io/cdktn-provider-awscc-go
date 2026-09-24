// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lambdafunction


type LambdaFunctionEphemeralStorage struct {
	// The size of the function's ``/tmp`` directory.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/lambda_function#size LambdaFunction#size}
	Size *float64 `field:"optional" json:"size" yaml:"size"`
}

