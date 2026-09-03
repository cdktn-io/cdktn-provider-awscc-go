// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lightsaildisk

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LightsailDiskConfig struct {
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
	// The names to use for your new Lightsail disk.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/lightsail_disk#disk_name LightsailDisk#disk_name}
	DiskName *string `field:"required" json:"diskName" yaml:"diskName"`
	// Size of the Lightsail disk.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/lightsail_disk#size_in_gb LightsailDisk#size_in_gb}
	SizeInGb *float64 `field:"required" json:"sizeInGb" yaml:"sizeInGb"`
	// An array of objects representing the add-ons to enable for the new instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/lightsail_disk#add_ons LightsailDisk#add_ons}
	AddOns interface{} `field:"optional" json:"addOns" yaml:"addOns"`
	// The Availability Zone in which to create your instance.
	//
	// Use the following format: us-east-2a (case sensitive). Be sure to add the include Availability Zones parameter to your request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/lightsail_disk#availability_zone LightsailDisk#availability_zone}
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// Location of a resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/lightsail_disk#location LightsailDisk#location}
	Location *LightsailDiskLocation `field:"optional" json:"location" yaml:"location"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/lightsail_disk#tags LightsailDisk#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

