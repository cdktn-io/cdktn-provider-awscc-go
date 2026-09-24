// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cleanroomsmltrainingdataset


type CleanroomsmlTrainingDatasetTrainingData struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/cleanroomsml_training_dataset#input_config CleanroomsmlTrainingDataset#input_config}.
	InputConfig *CleanroomsmlTrainingDatasetTrainingDataInputConfig `field:"required" json:"inputConfig" yaml:"inputConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/cleanroomsml_training_dataset#type CleanroomsmlTrainingDataset#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

