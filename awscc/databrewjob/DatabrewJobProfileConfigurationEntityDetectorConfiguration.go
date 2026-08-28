// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package databrewjob


type DatabrewJobProfileConfigurationEntityDetectorConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/databrew_job#allowed_statistics DatabrewJob#allowed_statistics}.
	AllowedStatistics *DatabrewJobProfileConfigurationEntityDetectorConfigurationAllowedStatistics `field:"optional" json:"allowedStatistics" yaml:"allowedStatistics"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/databrew_job#entity_types DatabrewJob#entity_types}.
	EntityTypes *[]*string `field:"optional" json:"entityTypes" yaml:"entityTypes"`
}

