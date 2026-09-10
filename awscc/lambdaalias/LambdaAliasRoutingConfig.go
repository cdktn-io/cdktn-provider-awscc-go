// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lambdaalias


type LambdaAliasRoutingConfig struct {
	// The second version, and the percentage of traffic that's routed to it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/lambda_alias#additional_version_weights LambdaAlias#additional_version_weights}
	AdditionalVersionWeights interface{} `field:"optional" json:"additionalVersionWeights" yaml:"additionalVersionWeights"`
}

