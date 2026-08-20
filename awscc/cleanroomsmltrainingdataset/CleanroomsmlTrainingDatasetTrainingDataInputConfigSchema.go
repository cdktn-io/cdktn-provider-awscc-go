// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cleanroomsmltrainingdataset


type CleanroomsmlTrainingDatasetTrainingDataInputConfigSchema struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cleanroomsml_training_dataset#column_name CleanroomsmlTrainingDataset#column_name}.
	ColumnName *string `field:"required" json:"columnName" yaml:"columnName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cleanroomsml_training_dataset#column_types CleanroomsmlTrainingDataset#column_types}.
	ColumnTypes *[]*string `field:"required" json:"columnTypes" yaml:"columnTypes"`
}

