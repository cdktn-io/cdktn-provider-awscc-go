// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package oamlink

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OamLinkConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/oam_link#resource_types OamLink#resource_types}.
	ResourceTypes *[]*string `field:"required" json:"resourceTypes" yaml:"resourceTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/oam_link#sink_identifier OamLink#sink_identifier}.
	SinkIdentifier *string `field:"required" json:"sinkIdentifier" yaml:"sinkIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/oam_link#label_template OamLink#label_template}.
	LabelTemplate *string `field:"optional" json:"labelTemplate" yaml:"labelTemplate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/oam_link#link_configuration OamLink#link_configuration}.
	LinkConfiguration *OamLinkLinkConfiguration `field:"optional" json:"linkConfiguration" yaml:"linkConfiguration"`
	// Tags to apply to the link.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/oam_link#tags OamLink#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

