// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2spotfleet

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Ec2SpotFleetConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_spot_fleet#spot_fleet_request_config_data Ec2SpotFleet#spot_fleet_request_config_data}.
	SpotFleetRequestConfigData *Ec2SpotFleetSpotFleetRequestConfigData `field:"required" json:"spotFleetRequestConfigData" yaml:"spotFleetRequestConfigData"`
	// The tags to specify in SpotFleetRequestConfigData.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_spot_fleet#tags Ec2SpotFleet#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

