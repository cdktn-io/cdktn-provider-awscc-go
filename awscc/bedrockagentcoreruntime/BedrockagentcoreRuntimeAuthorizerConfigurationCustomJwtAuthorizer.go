// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreruntime


type BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizer struct {
	// List of allowed audiences.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_runtime#allowed_audience BedrockagentcoreRuntime#allowed_audience}
	AllowedAudience *[]*string `field:"optional" json:"allowedAudience" yaml:"allowedAudience"`
	// List of allowed clients.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_runtime#allowed_clients BedrockagentcoreRuntime#allowed_clients}
	AllowedClients *[]*string `field:"optional" json:"allowedClients" yaml:"allowedClients"`
	// List of allowed scopes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_runtime#allowed_scopes BedrockagentcoreRuntime#allowed_scopes}
	AllowedScopes *[]*string `field:"optional" json:"allowedScopes" yaml:"allowedScopes"`
	// Allow-list of upstream workloads permitted to reach this resource via the workload identity chain.
	//
	// When set, the data plane enforces that the introspected workload chain's caller matches one of the configured hosting environments or workload identities; absent means no chain enforcement.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_runtime#allowed_workload_configuration BedrockagentcoreRuntime#allowed_workload_configuration}
	AllowedWorkloadConfiguration *BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerAllowedWorkloadConfiguration `field:"optional" json:"allowedWorkloadConfiguration" yaml:"allowedWorkloadConfiguration"`
	// List of required custom claims.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_runtime#custom_claims BedrockagentcoreRuntime#custom_claims}
	CustomClaims interface{} `field:"optional" json:"customClaims" yaml:"customClaims"`
	// OpenID Connect discovery URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_runtime#discovery_url BedrockagentcoreRuntime#discovery_url}
	DiscoveryUrl *string `field:"optional" json:"discoveryUrl" yaml:"discoveryUrl"`
	// Private endpoint configuration. Exactly one of SelfManagedLatticeResource or ManagedVpcResource must be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_runtime#private_endpoint BedrockagentcoreRuntime#private_endpoint}
	PrivateEndpoint *BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerPrivateEndpoint `field:"optional" json:"privateEndpoint" yaml:"privateEndpoint"`
	// List of private endpoint overrides.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_runtime#private_endpoint_overrides BedrockagentcoreRuntime#private_endpoint_overrides}
	PrivateEndpointOverrides interface{} `field:"optional" json:"privateEndpointOverrides" yaml:"privateEndpointOverrides"`
}

