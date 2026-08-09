// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cleanroomsmlconfiguredmodelalgorithm


type CleanroomsmlConfiguredModelAlgorithmTrainingContainerConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/cleanroomsml_configured_model_algorithm#arguments CleanroomsmlConfiguredModelAlgorithm#arguments}.
	Arguments *[]*string `field:"optional" json:"arguments" yaml:"arguments"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/cleanroomsml_configured_model_algorithm#entrypoint CleanroomsmlConfiguredModelAlgorithm#entrypoint}.
	Entrypoint *[]*string `field:"optional" json:"entrypoint" yaml:"entrypoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/cleanroomsml_configured_model_algorithm#image_uri CleanroomsmlConfiguredModelAlgorithm#image_uri}.
	ImageUri *string `field:"optional" json:"imageUri" yaml:"imageUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/cleanroomsml_configured_model_algorithm#metric_definitions CleanroomsmlConfiguredModelAlgorithm#metric_definitions}.
	MetricDefinitions interface{} `field:"optional" json:"metricDefinitions" yaml:"metricDefinitions"`
}

