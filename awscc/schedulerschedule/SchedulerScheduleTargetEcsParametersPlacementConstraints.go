// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package schedulerschedule


type SchedulerScheduleTargetEcsParametersPlacementConstraints struct {
	// A cluster query language expression to apply to the constraint.
	//
	// You cannot specify an expression if the constraint type is distinctInstance. To learn more, see Cluster Query Language in the Amazon Elastic Container Service Developer Guide.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/scheduler_schedule#expression SchedulerSchedule#expression}
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	// The type of constraint.
	//
	// Use distinctInstance to ensure that each task in a particular group is running on a different container instance. Use memberOf to restrict the selection to a group of valid candidates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/scheduler_schedule#type SchedulerSchedule#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

