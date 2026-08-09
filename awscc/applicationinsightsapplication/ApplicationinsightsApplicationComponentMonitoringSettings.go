// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package applicationinsightsapplication


type ApplicationinsightsApplicationComponentMonitoringSettings struct {
	// The ARN of the compnonent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/applicationinsights_application#component_arn ApplicationinsightsApplication#component_arn}
	ComponentArn *string `field:"optional" json:"componentArn" yaml:"componentArn"`
	// The component monitoring configuration mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/applicationinsights_application#component_configuration_mode ApplicationinsightsApplication#component_configuration_mode}
	ComponentConfigurationMode *string `field:"optional" json:"componentConfigurationMode" yaml:"componentConfigurationMode"`
	// The name of the component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/applicationinsights_application#component_name ApplicationinsightsApplication#component_name}
	ComponentName *string `field:"optional" json:"componentName" yaml:"componentName"`
	// The monitoring configuration of the component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/applicationinsights_application#custom_component_configuration ApplicationinsightsApplication#custom_component_configuration}
	CustomComponentConfiguration *ApplicationinsightsApplicationComponentMonitoringSettingsCustomComponentConfiguration `field:"optional" json:"customComponentConfiguration" yaml:"customComponentConfiguration"`
	// The overwritten settings on default component monitoring configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/applicationinsights_application#default_overwrite_component_configuration ApplicationinsightsApplication#default_overwrite_component_configuration}
	DefaultOverwriteComponentConfiguration *ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfiguration `field:"optional" json:"defaultOverwriteComponentConfiguration" yaml:"defaultOverwriteComponentConfiguration"`
	// The tier of the application component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/applicationinsights_application#tier ApplicationinsightsApplication#tier}
	Tier *string `field:"optional" json:"tier" yaml:"tier"`
}

