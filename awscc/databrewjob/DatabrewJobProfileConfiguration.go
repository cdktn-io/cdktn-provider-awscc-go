// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package databrewjob


type DatabrewJobProfileConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/databrew_job#column_statistics_configurations DatabrewJob#column_statistics_configurations}.
	ColumnStatisticsConfigurations interface{} `field:"optional" json:"columnStatisticsConfigurations" yaml:"columnStatisticsConfigurations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/databrew_job#dataset_statistics_configuration DatabrewJob#dataset_statistics_configuration}.
	DatasetStatisticsConfiguration *DatabrewJobProfileConfigurationDatasetStatisticsConfiguration `field:"optional" json:"datasetStatisticsConfiguration" yaml:"datasetStatisticsConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/databrew_job#entity_detector_configuration DatabrewJob#entity_detector_configuration}.
	EntityDetectorConfiguration *DatabrewJobProfileConfigurationEntityDetectorConfiguration `field:"optional" json:"entityDetectorConfiguration" yaml:"entityDetectorConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/databrew_job#profile_columns DatabrewJob#profile_columns}.
	ProfileColumns interface{} `field:"optional" json:"profileColumns" yaml:"profileColumns"`
}

