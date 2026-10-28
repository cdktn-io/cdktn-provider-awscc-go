// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package imagebuilderinfrastructureconfiguration


type ImagebuilderInfrastructureConfigurationParentInstanceAutoScalingGroup struct {
	// The number of seconds that Image Builder waits for an instance to become available in the group before the build fails.
	//
	// If you set this value to 0 or omit it, the build fails as soon as the group has no instance available.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/imagebuilder_infrastructure_configuration#capacity_wait_timeout_seconds ImagebuilderInfrastructureConfiguration#capacity_wait_timeout_seconds}
	CapacityWaitTimeoutSeconds *float64 `field:"optional" json:"capacityWaitTimeoutSeconds" yaml:"capacityWaitTimeoutSeconds"`
	// The name of the Auto Scaling group, in the same account and AWS Region as the infrastructure configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/imagebuilder_infrastructure_configuration#name ImagebuilderInfrastructureConfiguration#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

