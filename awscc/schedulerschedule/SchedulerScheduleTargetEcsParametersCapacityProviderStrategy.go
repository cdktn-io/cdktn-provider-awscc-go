// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package schedulerschedule


type SchedulerScheduleTargetEcsParametersCapacityProviderStrategy struct {
	// The base value designates how many tasks, at a minimum, to run on the specified capacity provider.
	//
	// Only one capacity provider in a capacity provider strategy can have a base defined. If no value is specified, the default value of 0 is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/scheduler_schedule#base SchedulerSchedule#base}
	Base *float64 `field:"optional" json:"base" yaml:"base"`
	// The short name of the capacity provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/scheduler_schedule#capacity_provider SchedulerSchedule#capacity_provider}
	CapacityProvider *string `field:"optional" json:"capacityProvider" yaml:"capacityProvider"`
	// The weight value designates the relative percentage of the total number of tasks launched that should use the specified capacity provider.
	//
	// The weight value is taken into consideration after the base value, if defined, is satisfied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/scheduler_schedule#weight SchedulerSchedule#weight}
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

