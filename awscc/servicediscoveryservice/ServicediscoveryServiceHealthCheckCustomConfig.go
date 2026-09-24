// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package servicediscoveryservice


type ServicediscoveryServiceHealthCheckCustomConfig struct {
	// The number of consecutive health check failures required before the service is considered unhealthy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/servicediscovery_service#failure_threshold ServicediscoveryService#failure_threshold}
	FailureThreshold *float64 `field:"optional" json:"failureThreshold" yaml:"failureThreshold"`
}

