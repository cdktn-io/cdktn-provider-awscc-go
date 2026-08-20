// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securitylakesubscribernotification


type SecuritylakeSubscriberNotificationNotificationConfigurationHttpsNotificationConfiguration struct {
	// The key name for the notification subscription.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/securitylake_subscriber_notification#authorization_api_key_name SecuritylakeSubscriberNotification#authorization_api_key_name}
	AuthorizationApiKeyName *string `field:"optional" json:"authorizationApiKeyName" yaml:"authorizationApiKeyName"`
	// The key value for the notification subscription.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/securitylake_subscriber_notification#authorization_api_key_value SecuritylakeSubscriberNotification#authorization_api_key_value}
	AuthorizationApiKeyValue *string `field:"optional" json:"authorizationApiKeyValue" yaml:"authorizationApiKeyValue"`
	// The subscription endpoint in Security Lake.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/securitylake_subscriber_notification#endpoint SecuritylakeSubscriberNotification#endpoint}
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
	// The HTTPS method used for the notification subscription.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/securitylake_subscriber_notification#http_method SecuritylakeSubscriberNotification#http_method}
	HttpMethod *string `field:"optional" json:"httpMethod" yaml:"httpMethod"`
	// The Amazon Resource Name (ARN) of the EventBridge API destinations IAM role that you created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/securitylake_subscriber_notification#target_role_arn SecuritylakeSubscriberNotification#target_role_arn}
	TargetRoleArn *string `field:"optional" json:"targetRoleArn" yaml:"targetRoleArn"`
}

