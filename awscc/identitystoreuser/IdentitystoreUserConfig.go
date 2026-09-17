// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package identitystoreuser

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IdentitystoreUserConfig struct {
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
	// The globally unique identifier for the identity store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/identitystore_user#identity_store_id IdentitystoreUser#identity_store_id}
	IdentityStoreId *string `field:"required" json:"identityStoreId" yaml:"identityStoreId"`
	// A list of addresses associated with the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/identitystore_user#addresses IdentitystoreUser#addresses}
	Addresses interface{} `field:"optional" json:"addresses" yaml:"addresses"`
	// The user's birthdate in YYYY-MM-DD format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/identitystore_user#birthdate IdentitystoreUser#birthdate}
	Birthdate *string `field:"optional" json:"birthdate" yaml:"birthdate"`
	// A string containing the name of the user for display.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/identitystore_user#display_name IdentitystoreUser#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// A list of email addresses associated with the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/identitystore_user#emails IdentitystoreUser#emails}
	Emails interface{} `field:"optional" json:"emails" yaml:"emails"`
	// The geographical region or location of the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/identitystore_user#locale IdentitystoreUser#locale}
	Locale *string `field:"optional" json:"locale" yaml:"locale"`
	// The name of the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/identitystore_user#name IdentitystoreUser#name}
	Name *IdentitystoreUserName `field:"optional" json:"name" yaml:"name"`
	// An alternate name for the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/identitystore_user#nick_name IdentitystoreUser#nick_name}
	NickName *string `field:"optional" json:"nickName" yaml:"nickName"`
	// A list of phone numbers associated with the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/identitystore_user#phone_numbers IdentitystoreUser#phone_numbers}
	PhoneNumbers interface{} `field:"optional" json:"phoneNumbers" yaml:"phoneNumbers"`
	// A list of photos associated with the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/identitystore_user#photos IdentitystoreUser#photos}
	Photos interface{} `field:"optional" json:"photos" yaml:"photos"`
	// The preferred language of the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/identitystore_user#preferred_language IdentitystoreUser#preferred_language}
	PreferredLanguage *string `field:"optional" json:"preferredLanguage" yaml:"preferredLanguage"`
	// A URL associated with the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/identitystore_user#profile_url IdentitystoreUser#profile_url}
	ProfileUrl *string `field:"optional" json:"profileUrl" yaml:"profileUrl"`
	// A list of roles associated with the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/identitystore_user#roles IdentitystoreUser#roles}
	Roles interface{} `field:"optional" json:"roles" yaml:"roles"`
	// The time zone for the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/identitystore_user#timezone IdentitystoreUser#timezone}
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
	// The title of the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/identitystore_user#title IdentitystoreUser#title}
	Title *string `field:"optional" json:"title" yaml:"title"`
	// A unique string used to identify the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/identitystore_user#user_name IdentitystoreUser#user_name}
	UserName *string `field:"optional" json:"userName" yaml:"userName"`
	// A string indicating the type of user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/identitystore_user#user_type IdentitystoreUser#user_type}
	UserType *string `field:"optional" json:"userType" yaml:"userType"`
	// The user's personal website or blog URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/identitystore_user#website IdentitystoreUser#website}
	Website *string `field:"optional" json:"website" yaml:"website"`
}

