// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package syntheticscanary


type SyntheticsCanaryRunConfig struct {
	// Enable active tracing if set to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/synthetics_canary#active_tracing SyntheticsCanary#active_tracing}
	ActiveTracing interface{} `field:"optional" json:"activeTracing" yaml:"activeTracing"`
	// Environment variable key-value pairs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/synthetics_canary#environment_variables SyntheticsCanary#environment_variables}
	EnvironmentVariables *map[string]*string `field:"optional" json:"environmentVariables" yaml:"environmentVariables"`
	// Provide ephemeralStorage available for canary in MB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/synthetics_canary#ephemeral_storage SyntheticsCanary#ephemeral_storage}
	EphemeralStorage *float64 `field:"optional" json:"ephemeralStorage" yaml:"ephemeralStorage"`
	// Provide maximum memory available for canary in MB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/synthetics_canary#memory_in_mb SyntheticsCanary#memory_in_mb}
	MemoryInMb *float64 `field:"optional" json:"memoryInMb" yaml:"memoryInMb"`
	// Provide maximum canary timeout per run in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/synthetics_canary#timeout_in_seconds SyntheticsCanary#timeout_in_seconds}
	TimeoutInSeconds *float64 `field:"optional" json:"timeoutInSeconds" yaml:"timeoutInSeconds"`
}

