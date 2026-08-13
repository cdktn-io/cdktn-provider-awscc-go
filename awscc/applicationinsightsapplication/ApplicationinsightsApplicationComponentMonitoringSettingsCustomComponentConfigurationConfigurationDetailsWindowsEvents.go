// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package applicationinsightsapplication


type ApplicationinsightsApplicationComponentMonitoringSettingsCustomComponentConfigurationConfigurationDetailsWindowsEvents struct {
	// The levels of event to log.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/applicationinsights_application#event_levels ApplicationinsightsApplication#event_levels}
	EventLevels *[]*string `field:"optional" json:"eventLevels" yaml:"eventLevels"`
	// The type of Windows Events to log.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/applicationinsights_application#event_name ApplicationinsightsApplication#event_name}
	EventName *string `field:"optional" json:"eventName" yaml:"eventName"`
	// The CloudWatch log group name to be associated to the monitored log.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/applicationinsights_application#log_group_name ApplicationinsightsApplication#log_group_name}
	LogGroupName *string `field:"optional" json:"logGroupName" yaml:"logGroupName"`
	// The name of the log pattern set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/applicationinsights_application#pattern_set ApplicationinsightsApplication#pattern_set}
	PatternSet *string `field:"optional" json:"patternSet" yaml:"patternSet"`
}

