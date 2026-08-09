// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package evidentlyexperiment


type EvidentlyExperimentMetricGoals struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/evidently_experiment#desired_change EvidentlyExperiment#desired_change}.
	DesiredChange *string `field:"required" json:"desiredChange" yaml:"desiredChange"`
	// The JSON path to reference the entity id in the event.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/evidently_experiment#entity_id_key EvidentlyExperiment#entity_id_key}
	EntityIdKey *string `field:"required" json:"entityIdKey" yaml:"entityIdKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/evidently_experiment#metric_name EvidentlyExperiment#metric_name}.
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// The JSON path to reference the numerical metric value in the event.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/evidently_experiment#value_key EvidentlyExperiment#value_key}
	ValueKey *string `field:"required" json:"valueKey" yaml:"valueKey"`
	// Event patterns have the same structure as the events they match.
	//
	// Rules use event patterns to select events. An event pattern either matches an event or it doesn't.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/evidently_experiment#event_pattern EvidentlyExperiment#event_pattern}
	EventPattern *string `field:"optional" json:"eventPattern" yaml:"eventPattern"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/evidently_experiment#unit_label EvidentlyExperiment#unit_label}.
	UnitLabel *string `field:"optional" json:"unitLabel" yaml:"unitLabel"`
}

