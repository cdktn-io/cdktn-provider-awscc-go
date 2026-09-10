// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lambdafunction


type LambdaFunctionTenancyConfig struct {
	// Tenant isolation mode allows for invocation to be sent to a corresponding execution environment dedicated to a specific tenant ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/lambda_function#tenant_isolation_mode LambdaFunction#tenant_isolation_mode}
	TenantIsolationMode *string `field:"optional" json:"tenantIsolationMode" yaml:"tenantIsolationMode"`
}

