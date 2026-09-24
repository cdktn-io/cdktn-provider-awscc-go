// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sescontactlist

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SesContactListConfig struct {
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
	// The name of the contact list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ses_contact_list#contact_list_name SesContactList#contact_list_name}
	ContactListName *string `field:"optional" json:"contactListName" yaml:"contactListName"`
	// The description of the contact list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ses_contact_list#description SesContactList#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The tags (keys and values) associated with the contact list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ses_contact_list#tags SesContactList#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The topics associated with the contact list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ses_contact_list#topics SesContactList#topics}
	Topics interface{} `field:"optional" json:"topics" yaml:"topics"`
}

