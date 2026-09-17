// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lightsailcontainer


type LightsailContainerContainerServiceDeploymentPublicEndpointHealthCheckConfig struct {
	// The number of consecutive health checks successes required before moving the container to the Healthy state.
	//
	// The default value is 2.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lightsail_container#healthy_threshold LightsailContainer#healthy_threshold}
	HealthyThreshold *float64 `field:"optional" json:"healthyThreshold" yaml:"healthyThreshold"`
	// The approximate interval, in seconds, between health checks of an individual container.
	//
	// You can specify between 5 and 300 seconds. The default value is 5.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lightsail_container#interval_seconds LightsailContainer#interval_seconds}
	IntervalSeconds *float64 `field:"optional" json:"intervalSeconds" yaml:"intervalSeconds"`
	// The path on the container on which to perform the health check. The default value is /.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lightsail_container#path LightsailContainer#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// The HTTP codes to use when checking for a successful response from a container.
	//
	// You can specify values between 200 and 499. You can specify multiple values (for example, 200,202) or a range of values (for example, 200-299).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lightsail_container#success_codes LightsailContainer#success_codes}
	SuccessCodes *string `field:"optional" json:"successCodes" yaml:"successCodes"`
	// The amount of time, in seconds, during which no response means a failed health check.
	//
	// You can specify between 2 and 60 seconds. The default value is 2.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lightsail_container#timeout_seconds LightsailContainer#timeout_seconds}
	TimeoutSeconds *float64 `field:"optional" json:"timeoutSeconds" yaml:"timeoutSeconds"`
	// The number of consecutive health check failures required before moving the container to the Unhealthy state.
	//
	// The default value is 2.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lightsail_container#unhealthy_threshold LightsailContainer#unhealthy_threshold}
	UnhealthyThreshold *float64 `field:"optional" json:"unhealthyThreshold" yaml:"unhealthyThreshold"`
}

