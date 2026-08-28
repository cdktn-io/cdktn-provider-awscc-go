// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lambdaalias


type LambdaAliasProvisionedConcurrencyConfig struct {
	// The amount of provisioned concurrency to allocate for the alias.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lambda_alias#provisioned_concurrent_executions LambdaAlias#provisioned_concurrent_executions}
	ProvisionedConcurrentExecutions *float64 `field:"optional" json:"provisionedConcurrentExecutions" yaml:"provisionedConcurrentExecutions"`
}

