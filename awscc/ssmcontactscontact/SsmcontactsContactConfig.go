// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ssmcontactscontact

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SsmcontactsContactConfig struct {
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
	// Alias of the contact. String value with 20 to 256 characters. Only alphabetical, numeric characters, dash, or underscore allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/ssmcontacts_contact#alias SsmcontactsContact#alias}
	Alias *string `field:"required" json:"alias" yaml:"alias"`
	// Name of the contact.
	//
	// String value with 3 to 256 characters. Only alphabetical, space, numeric characters, dash, or underscore allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/ssmcontacts_contact#display_name SsmcontactsContact#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Contact type, which specify type of contact. Currently supported values: ?PERSONAL?, ?SHARED?, ?OTHER?.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/ssmcontacts_contact#type SsmcontactsContact#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The stages that an escalation plan or engagement plan engages contacts and contact methods in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/ssmcontacts_contact#plan SsmcontactsContact#plan}
	Plan interface{} `field:"optional" json:"plan" yaml:"plan"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/ssmcontacts_contact#tags SsmcontactsContact#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

