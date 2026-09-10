// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package configconfigrule


type ConfigConfigRuleSourceSourceDetails struct {
	// The source of the event, such as an AWS service, that triggers CC to evaluate your AWS resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/config_config_rule#event_source ConfigConfigRule#event_source}
	EventSource *string `field:"optional" json:"eventSource" yaml:"eventSource"`
	// The frequency at which you want CC to run evaluations for a custom rule with a periodic trigger.
	//
	// If you specify a value for ``MaximumExecutionFrequency``, then ``MessageType`` must use the ``ScheduledNotification`` value.
	//   By default, rules with a periodic trigger are evaluated every 24 hours. To change the frequency, specify a valid value for the ``MaximumExecutionFrequency`` parameter.
	//  Based on the valid value you choose, CC runs evaluations once for each valid value. For example, if you choose ``Three_Hours``, CC runs evaluations once every three hours. In this case, ``Three_Hours`` is the frequency of this rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/config_config_rule#maximum_execution_frequency ConfigConfigRule#maximum_execution_frequency}
	MaximumExecutionFrequency *string `field:"optional" json:"maximumExecutionFrequency" yaml:"maximumExecutionFrequency"`
	// The type of notification that triggers CC to run an evaluation for a rule.
	//
	// You can specify the following notification types:
	//   +  ``ConfigurationItemChangeNotification`` - Triggers an evaluation when CC delivers a configuration item as a result of a resource change.
	//   +  ``OversizedConfigurationItemChangeNotification`` - Triggers an evaluation when CC delivers an oversized configuration item. CC may generate this notification type when a resource changes and the notification exceeds the maximum size allowed by Amazon SNS.
	//   +  ``ScheduledNotification`` - Triggers a periodic evaluation at the frequency specified for ``MaximumExecutionFrequency``.
	//   +  ``ConfigurationSnapshotDeliveryCompleted`` - Triggers a periodic evaluation when CC delivers a configuration snapshot.
	//
	//  If you want your custom rule to be triggered by configuration changes, specify two SourceDetail objects, one for ``ConfigurationItemChangeNotification`` and one for ``OversizedConfigurationItemChangeNotification``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/config_config_rule#message_type ConfigConfigRule#message_type}
	MessageType *string `field:"optional" json:"messageType" yaml:"messageType"`
}

