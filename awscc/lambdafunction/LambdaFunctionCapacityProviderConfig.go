// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lambdafunction


type LambdaFunctionCapacityProviderConfig struct {
	// Configuration for Lambda-managed instances used by the capacity provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/lambda_function#lambda_managed_instances_capacity_provider_config LambdaFunction#lambda_managed_instances_capacity_provider_config}
	LambdaManagedInstancesCapacityProviderConfig *LambdaFunctionCapacityProviderConfigLambdaManagedInstancesCapacityProviderConfig `field:"optional" json:"lambdaManagedInstancesCapacityProviderConfig" yaml:"lambdaManagedInstancesCapacityProviderConfig"`
}

