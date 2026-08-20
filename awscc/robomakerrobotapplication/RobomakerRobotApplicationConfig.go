// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package robomakerrobotapplication

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type RobomakerRobotApplicationConfig struct {
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
	// The robot software suite used by the robot application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/robomaker_robot_application#robot_software_suite RobomakerRobotApplication#robot_software_suite}
	RobotSoftwareSuite *RobomakerRobotApplicationRobotSoftwareSuite `field:"required" json:"robotSoftwareSuite" yaml:"robotSoftwareSuite"`
	// The revision ID of robot application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/robomaker_robot_application#current_revision_id RobomakerRobotApplication#current_revision_id}
	CurrentRevisionId *string `field:"optional" json:"currentRevisionId" yaml:"currentRevisionId"`
	// The URI of the Docker image for the robot application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/robomaker_robot_application#environment RobomakerRobotApplication#environment}
	Environment *string `field:"optional" json:"environment" yaml:"environment"`
	// The name of the robot application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/robomaker_robot_application#name RobomakerRobotApplication#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The sources of the robot application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/robomaker_robot_application#sources RobomakerRobotApplication#sources}
	Sources interface{} `field:"optional" json:"sources" yaml:"sources"`
	// A key-value pair to associate with a resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/robomaker_robot_application#tags RobomakerRobotApplication#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

