// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package databrewjob


type DatabrewJobProfileConfigurationDatasetStatisticsConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/databrew_job#included_statistics DatabrewJob#included_statistics}.
	IncludedStatistics *[]*string `field:"optional" json:"includedStatistics" yaml:"includedStatistics"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/databrew_job#overrides DatabrewJob#overrides}.
	Overrides interface{} `field:"optional" json:"overrides" yaml:"overrides"`
}

