// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lambdaalias


type LambdaAliasRoutingConfigAdditionalVersionWeights struct {
	// The qualifier of the second version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lambda_alias#function_version LambdaAlias#function_version}
	FunctionVersion *string `field:"optional" json:"functionVersion" yaml:"functionVersion"`
	// The percentage of traffic that the alias routes to the second version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lambda_alias#function_weight LambdaAlias#function_weight}
	FunctionWeight *float64 `field:"optional" json:"functionWeight" yaml:"functionWeight"`
}

