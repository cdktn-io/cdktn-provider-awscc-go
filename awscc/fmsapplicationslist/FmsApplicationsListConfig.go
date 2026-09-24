// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package fmsapplicationslist

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type FmsApplicationsListConfig struct {
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
	// An array of applications in the Firewall Manager applications list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/fms_applications_list#apps_list FmsApplicationsList#apps_list}
	AppsList interface{} `field:"required" json:"appsList" yaml:"appsList"`
	// The name of the Firewall Manager applications list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/fms_applications_list#list_name FmsApplicationsList#list_name}
	ListName *string `field:"required" json:"listName" yaml:"listName"`
	// An array of key-value pairs to apply to the applications list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/fms_applications_list#tags FmsApplicationsList#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

