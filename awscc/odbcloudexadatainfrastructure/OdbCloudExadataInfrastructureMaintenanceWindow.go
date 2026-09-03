// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package odbcloudexadatainfrastructure


type OdbCloudExadataInfrastructureMaintenanceWindow struct {
	// The timeout duration for custom actions in minutes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/odb_cloud_exadata_infrastructure#custom_action_timeout_in_mins OdbCloudExadataInfrastructure#custom_action_timeout_in_mins}
	CustomActionTimeoutInMins *float64 `field:"optional" json:"customActionTimeoutInMins" yaml:"customActionTimeoutInMins"`
	// The days of the week when maintenance can be performed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/odb_cloud_exadata_infrastructure#days_of_week OdbCloudExadataInfrastructure#days_of_week}
	DaysOfWeek *[]*string `field:"optional" json:"daysOfWeek" yaml:"daysOfWeek"`
	// The hours of the day when maintenance can be performed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/odb_cloud_exadata_infrastructure#hours_of_day OdbCloudExadataInfrastructure#hours_of_day}
	HoursOfDay *[]*float64 `field:"optional" json:"hoursOfDay" yaml:"hoursOfDay"`
	// Indicates whether custom action timeout is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/odb_cloud_exadata_infrastructure#is_custom_action_timeout_enabled OdbCloudExadataInfrastructure#is_custom_action_timeout_enabled}
	IsCustomActionTimeoutEnabled interface{} `field:"optional" json:"isCustomActionTimeoutEnabled" yaml:"isCustomActionTimeoutEnabled"`
	// The lead time in weeks before the maintenance window.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/odb_cloud_exadata_infrastructure#lead_time_in_weeks OdbCloudExadataInfrastructure#lead_time_in_weeks}
	LeadTimeInWeeks *float64 `field:"optional" json:"leadTimeInWeeks" yaml:"leadTimeInWeeks"`
	// The months when maintenance can be performed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/odb_cloud_exadata_infrastructure#months OdbCloudExadataInfrastructure#months}
	Months *[]*string `field:"optional" json:"months" yaml:"months"`
	// The patching mode for the maintenance window.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/odb_cloud_exadata_infrastructure#patching_mode OdbCloudExadataInfrastructure#patching_mode}
	PatchingMode *string `field:"optional" json:"patchingMode" yaml:"patchingMode"`
	// The preference for the maintenance window scheduling.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/odb_cloud_exadata_infrastructure#preference OdbCloudExadataInfrastructure#preference}
	Preference *string `field:"optional" json:"preference" yaml:"preference"`
	// The weeks of the month when maintenance can be performed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/odb_cloud_exadata_infrastructure#weeks_of_month OdbCloudExadataInfrastructure#weeks_of_month}
	WeeksOfMonth *[]*float64 `field:"optional" json:"weeksOfMonth" yaml:"weeksOfMonth"`
}

