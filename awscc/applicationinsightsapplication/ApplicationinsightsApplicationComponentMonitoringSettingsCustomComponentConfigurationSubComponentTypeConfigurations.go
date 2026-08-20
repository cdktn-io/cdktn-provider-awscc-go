// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package applicationinsightsapplication


type ApplicationinsightsApplicationComponentMonitoringSettingsCustomComponentConfigurationSubComponentTypeConfigurations struct {
	// The configuration settings of sub components.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/applicationinsights_application#sub_component_configuration_details ApplicationinsightsApplication#sub_component_configuration_details}
	SubComponentConfigurationDetails *ApplicationinsightsApplicationComponentMonitoringSettingsCustomComponentConfigurationSubComponentTypeConfigurationsSubComponentConfigurationDetails `field:"optional" json:"subComponentConfigurationDetails" yaml:"subComponentConfigurationDetails"`
	// The sub component type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/applicationinsights_application#sub_component_type ApplicationinsightsApplication#sub_component_type}
	SubComponentType *string `field:"optional" json:"subComponentType" yaml:"subComponentType"`
}

