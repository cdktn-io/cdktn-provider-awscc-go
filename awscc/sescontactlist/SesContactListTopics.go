// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sescontactlist


type SesContactListTopics struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ses_contact_list#default_subscription_status SesContactList#default_subscription_status}.
	DefaultSubscriptionStatus *string `field:"optional" json:"defaultSubscriptionStatus" yaml:"defaultSubscriptionStatus"`
	// The description of the topic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ses_contact_list#description SesContactList#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The display name of the topic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ses_contact_list#display_name SesContactList#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// The name of the topic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ses_contact_list#topic_name SesContactList#topic_name}
	TopicName *string `field:"optional" json:"topicName" yaml:"topicName"`
}

