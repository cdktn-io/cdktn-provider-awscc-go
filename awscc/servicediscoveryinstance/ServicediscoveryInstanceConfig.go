// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package servicediscoveryinstance

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ServicediscoveryInstanceConfig struct {
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
	// A string map that contains information for the service that is specified in ServiceId.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/servicediscovery_instance#instance_attributes ServicediscoveryInstance#instance_attributes}
	InstanceAttributes *string `field:"required" json:"instanceAttributes" yaml:"instanceAttributes"`
	// The ID or Amazon Resource Name (ARN) of the service that you want to use for settings for the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/servicediscovery_instance#service_id ServicediscoveryInstance#service_id}
	ServiceId *string `field:"required" json:"serviceId" yaml:"serviceId"`
	// An identifier that you want to associate with the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/servicediscovery_instance#instance_id ServicediscoveryInstance#instance_id}
	InstanceId *string `field:"optional" json:"instanceId" yaml:"instanceId"`
}

