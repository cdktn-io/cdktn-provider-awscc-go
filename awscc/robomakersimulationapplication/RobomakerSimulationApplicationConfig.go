// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package robomakersimulationapplication

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type RobomakerSimulationApplicationConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The robot software suite used by the simulation application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/robomaker_simulation_application#robot_software_suite RobomakerSimulationApplication#robot_software_suite}
	RobotSoftwareSuite *RobomakerSimulationApplicationRobotSoftwareSuite `field:"required" json:"robotSoftwareSuite" yaml:"robotSoftwareSuite"`
	// The simulation software suite used by the simulation application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/robomaker_simulation_application#simulation_software_suite RobomakerSimulationApplication#simulation_software_suite}
	SimulationSoftwareSuite *RobomakerSimulationApplicationSimulationSoftwareSuite `field:"required" json:"simulationSoftwareSuite" yaml:"simulationSoftwareSuite"`
	// The current revision id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/robomaker_simulation_application#current_revision_id RobomakerSimulationApplication#current_revision_id}
	CurrentRevisionId *string `field:"optional" json:"currentRevisionId" yaml:"currentRevisionId"`
	// The URI of the Docker image for the robot application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/robomaker_simulation_application#environment RobomakerSimulationApplication#environment}
	Environment *string `field:"optional" json:"environment" yaml:"environment"`
	// The name of the simulation application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/robomaker_simulation_application#name RobomakerSimulationApplication#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The rendering engine for the simulation application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/robomaker_simulation_application#rendering_engine RobomakerSimulationApplication#rendering_engine}
	RenderingEngine *RobomakerSimulationApplicationRenderingEngine `field:"optional" json:"renderingEngine" yaml:"renderingEngine"`
	// The sources of the simulation application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/robomaker_simulation_application#sources RobomakerSimulationApplication#sources}
	Sources interface{} `field:"optional" json:"sources" yaml:"sources"`
	// A key-value pair to associate with a resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/robomaker_simulation_application#tags RobomakerSimulationApplication#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

