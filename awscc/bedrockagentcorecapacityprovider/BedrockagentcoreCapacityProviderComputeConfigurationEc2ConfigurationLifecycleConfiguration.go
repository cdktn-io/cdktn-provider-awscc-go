// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcorecapacityprovider


type BedrockagentcoreCapacityProviderComputeConfigurationEc2ConfigurationLifecycleConfiguration struct {
	// The number of seconds an instance can remain idle before it is stopped.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/bedrockagentcore_capacity_provider#idle_instance_timeout BedrockagentcoreCapacityProvider#idle_instance_timeout}
	IdleInstanceTimeout *float64 `field:"optional" json:"idleInstanceTimeout" yaml:"idleInstanceTimeout"`
	// Maximum lifetime for the instance in seconds.
	//
	// Once reached, instances will be automatically terminated regardless of activity. Default: 28800 seconds (8 hours). Maximum: 1209600 seconds (14 days).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/bedrockagentcore_capacity_provider#max_lifetime BedrockagentcoreCapacityProvider#max_lifetime}
	MaxLifetime *float64 `field:"optional" json:"maxLifetime" yaml:"maxLifetime"`
}

