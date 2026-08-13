// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package oamsink

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OamSinkConfig struct {
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
	// The name of the ObservabilityAccessManager Sink.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/oam_sink#name OamSink#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The policy of this ObservabilityAccessManager Sink.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/oam_sink#policy OamSink#policy}
	Policy *string `field:"optional" json:"policy" yaml:"policy"`
	// Tags to apply to the sink.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/oam_sink#tags OamSink#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

