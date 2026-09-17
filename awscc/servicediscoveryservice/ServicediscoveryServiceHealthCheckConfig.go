// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package servicediscoveryservice


type ServicediscoveryServiceHealthCheckConfig struct {
	// The number of consecutive health check failures that must occur before declaring the service unhealthy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/servicediscovery_service#failure_threshold ServicediscoveryService#failure_threshold}
	FailureThreshold *float64 `field:"optional" json:"failureThreshold" yaml:"failureThreshold"`
	// The path to ping on the service for health checks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/servicediscovery_service#resource_path ServicediscoveryService#resource_path}
	ResourcePath *string `field:"optional" json:"resourcePath" yaml:"resourcePath"`
	// The type of health check (e.g., HTTP, HTTPS, TCP).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/servicediscovery_service#type ServicediscoveryService#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

