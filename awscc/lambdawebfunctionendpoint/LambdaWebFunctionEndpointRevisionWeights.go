// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lambdawebfunctionendpoint


type LambdaWebFunctionEndpointRevisionWeights struct {
	// The revision identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/lambda_web_function_endpoint#revision_id LambdaWebFunctionEndpoint#revision_id}
	RevisionId *string `field:"optional" json:"revisionId" yaml:"revisionId"`
	// The traffic weight for this revision.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/lambda_web_function_endpoint#weight LambdaWebFunctionEndpoint#weight}
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

