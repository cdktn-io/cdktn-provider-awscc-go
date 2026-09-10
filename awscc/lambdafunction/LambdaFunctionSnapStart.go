// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lambdafunction


type LambdaFunctionSnapStart struct {
	// Set ``ApplyOn`` to ``PublishedVersions`` to create a snapshot of the initialized execution environment when you publish a function version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/lambda_function#apply_on LambdaFunction#apply_on}
	ApplyOn *string `field:"optional" json:"applyOn" yaml:"applyOn"`
}

