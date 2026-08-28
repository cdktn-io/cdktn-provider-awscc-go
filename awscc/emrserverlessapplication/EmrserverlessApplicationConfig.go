// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package emrserverlessapplication

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EmrserverlessApplicationConfig struct {
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
	// EMR release label.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/emrserverless_application#release_label EmrserverlessApplication#release_label}
	ReleaseLabel *string `field:"required" json:"releaseLabel" yaml:"releaseLabel"`
	// The type of the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/emrserverless_application#type EmrserverlessApplication#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The cpu architecture of an application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/emrserverless_application#architecture EmrserverlessApplication#architecture}
	Architecture *string `field:"optional" json:"architecture" yaml:"architecture"`
	// Configuration for Auto Start of Application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/emrserverless_application#auto_start_configuration EmrserverlessApplication#auto_start_configuration}
	AutoStartConfiguration *EmrserverlessApplicationAutoStartConfiguration `field:"optional" json:"autoStartConfiguration" yaml:"autoStartConfiguration"`
	// Configuration for Auto Stop of Application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/emrserverless_application#auto_stop_configuration EmrserverlessApplication#auto_stop_configuration}
	AutoStopConfiguration *EmrserverlessApplicationAutoStopConfiguration `field:"optional" json:"autoStopConfiguration" yaml:"autoStopConfiguration"`
	// The image configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/emrserverless_application#image_configuration EmrserverlessApplication#image_configuration}
	ImageConfiguration *EmrserverlessApplicationImageConfiguration `field:"optional" json:"imageConfiguration" yaml:"imageConfiguration"`
	// Initial capacity initialized when an Application is started.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/emrserverless_application#initial_capacity EmrserverlessApplication#initial_capacity}
	InitialCapacity interface{} `field:"optional" json:"initialCapacity" yaml:"initialCapacity"`
	// Maximum allowed cumulative resources for an Application. No new resources will be created once the limit is hit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/emrserverless_application#maximum_capacity EmrserverlessApplication#maximum_capacity}
	MaximumCapacity *EmrserverlessApplicationMaximumCapacity `field:"optional" json:"maximumCapacity" yaml:"maximumCapacity"`
	// User friendly Application name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/emrserverless_application#name EmrserverlessApplication#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Network Configuration for customer VPC connectivity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/emrserverless_application#network_configuration EmrserverlessApplication#network_configuration}
	NetworkConfiguration *EmrserverlessApplicationNetworkConfiguration `field:"optional" json:"networkConfiguration" yaml:"networkConfiguration"`
	// Tag map with key and value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/emrserverless_application#tags EmrserverlessApplication#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The key-value pairs that specify worker type to WorkerTypeSpecificationInput.
	//
	// This parameter must contain all valid worker types for a Spark or Hive application. Valid worker types include Driver and Executor for Spark applications and HiveDriver and TezTask for Hive applications. You can either set image details in this parameter for each worker type, or in imageConfiguration for all worker types.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/emrserverless_application#worker_type_specifications EmrserverlessApplication#worker_type_specifications}
	WorkerTypeSpecifications interface{} `field:"optional" json:"workerTypeSpecifications" yaml:"workerTypeSpecifications"`
}

