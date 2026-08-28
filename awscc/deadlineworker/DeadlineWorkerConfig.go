// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package deadlineworker

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DeadlineWorkerConfig struct {
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
	// The farm ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/deadline_worker#farm_id DeadlineWorker#farm_id}
	FarmId *string `field:"required" json:"farmId" yaml:"farmId"`
	// The fleet ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/deadline_worker#fleet_id DeadlineWorker#fleet_id}
	FleetId *string `field:"required" json:"fleetId" yaml:"fleetId"`
	// The IP address and host name of the worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/deadline_worker#host_properties DeadlineWorker#host_properties}
	HostProperties *DeadlineWorkerHostProperties `field:"optional" json:"hostProperties" yaml:"hostProperties"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/deadline_worker#tags DeadlineWorker#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

