// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package qbusinessapplication


type QbusinessApplicationAutoSubscriptionConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/qbusiness_application#auto_subscribe QbusinessApplication#auto_subscribe}.
	AutoSubscribe *string `field:"optional" json:"autoSubscribe" yaml:"autoSubscribe"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/qbusiness_application#default_subscription_type QbusinessApplication#default_subscription_type}.
	DefaultSubscriptionType *string `field:"optional" json:"defaultSubscriptionType" yaml:"defaultSubscriptionType"`
}

