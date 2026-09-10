// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package robomakersimulationapplication


type RobomakerSimulationApplicationRenderingEngine struct {
	// The name of the rendering engine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/robomaker_simulation_application#name RobomakerSimulationApplication#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The version of the rendering engine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/robomaker_simulation_application#version RobomakerSimulationApplication#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

