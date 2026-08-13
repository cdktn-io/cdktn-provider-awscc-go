// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package medialivemultiplex

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MedialiveMultiplexConfig struct {
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
	// A list of availability zones for the multiplex.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/medialive_multiplex#availability_zones MedialiveMultiplex#availability_zones}
	AvailabilityZones *[]*string `field:"required" json:"availabilityZones" yaml:"availabilityZones"`
	// Configuration for a multiplex event.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/medialive_multiplex#multiplex_settings MedialiveMultiplex#multiplex_settings}
	MultiplexSettings *MedialiveMultiplexMultiplexSettings `field:"required" json:"multiplexSettings" yaml:"multiplexSettings"`
	// Name of multiplex.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/medialive_multiplex#name MedialiveMultiplex#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A list of the multiplex output destinations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/medialive_multiplex#destinations MedialiveMultiplex#destinations}
	Destinations interface{} `field:"optional" json:"destinations" yaml:"destinations"`
	// A collection of key-value pairs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/medialive_multiplex#tags MedialiveMultiplex#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

