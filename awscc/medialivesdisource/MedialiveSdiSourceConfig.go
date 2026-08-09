// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package medialivesdisource

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MedialiveSdiSourceConfig struct {
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
	// The name of the SdiSource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/medialive_sdi_source#name MedialiveSdiSource#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The interface mode of the SdiSource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/medialive_sdi_source#type MedialiveSdiSource#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The current state of the SdiSource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/medialive_sdi_source#mode MedialiveSdiSource#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// A collection of key-value pairs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/medialive_sdi_source#tags MedialiveSdiSource#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

