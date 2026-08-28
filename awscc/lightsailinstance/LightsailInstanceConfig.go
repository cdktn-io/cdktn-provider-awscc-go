// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lightsailinstance

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LightsailInstanceConfig struct {
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
	// The ID for a virtual private server image (e.g., app_wordpress_4_4 or app_lamp_7_0 ). Use the get blueprints operation to return a list of available images (or blueprints ).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lightsail_instance#blueprint_id LightsailInstance#blueprint_id}
	BlueprintId *string `field:"required" json:"blueprintId" yaml:"blueprintId"`
	// The bundle of specification information for your virtual private server (or instance ), including the pricing plan (e.g., micro_1_0 ).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lightsail_instance#bundle_id LightsailInstance#bundle_id}
	BundleId *string `field:"required" json:"bundleId" yaml:"bundleId"`
	// The names to use for your new Lightsail instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lightsail_instance#instance_name LightsailInstance#instance_name}
	InstanceName *string `field:"required" json:"instanceName" yaml:"instanceName"`
	// An array of objects representing the add-ons to enable for the new instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lightsail_instance#add_ons LightsailInstance#add_ons}
	AddOns interface{} `field:"optional" json:"addOns" yaml:"addOns"`
	// The Availability Zone in which to create your instance.
	//
	// Use the following format: us-east-2a (case sensitive). Be sure to add the include Availability Zones parameter to your request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lightsail_instance#availability_zone LightsailInstance#availability_zone}
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// Hardware of the Instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lightsail_instance#hardware LightsailInstance#hardware}
	Hardware *LightsailInstanceHardware `field:"optional" json:"hardware" yaml:"hardware"`
	// The name of your key pair.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lightsail_instance#key_pair_name LightsailInstance#key_pair_name}
	KeyPairName *string `field:"optional" json:"keyPairName" yaml:"keyPairName"`
	// Location of a resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lightsail_instance#location LightsailInstance#location}
	Location *LightsailInstanceLocation `field:"optional" json:"location" yaml:"location"`
	// Networking of the Instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lightsail_instance#networking LightsailInstance#networking}
	Networking *LightsailInstanceNetworking `field:"optional" json:"networking" yaml:"networking"`
	// Current State of the Instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lightsail_instance#state LightsailInstance#state}
	State *LightsailInstanceState `field:"optional" json:"state" yaml:"state"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lightsail_instance#tags LightsailInstance#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// A launch script you can create that configures a server with additional user data.
	//
	// For example, you might want to run apt-get -y update.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lightsail_instance#user_data LightsailInstance#user_data}
	UserData *string `field:"optional" json:"userData" yaml:"userData"`
}

