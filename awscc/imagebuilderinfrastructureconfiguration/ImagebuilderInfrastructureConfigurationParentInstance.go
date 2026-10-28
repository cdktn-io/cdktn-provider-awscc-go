// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package imagebuilderinfrastructureconfiguration


type ImagebuilderInfrastructureConfigurationParentInstance struct {
	// The Amazon EC2 Auto Scaling group that Image Builder takes the build instance from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/imagebuilder_infrastructure_configuration#auto_scaling_group ImagebuilderInfrastructureConfiguration#auto_scaling_group}
	AutoScalingGroup *ImagebuilderInfrastructureConfigurationParentInstanceAutoScalingGroup `field:"optional" json:"autoScalingGroup" yaml:"autoScalingGroup"`
}

