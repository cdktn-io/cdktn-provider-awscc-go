// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datazonedatasource


type DatazoneDataSourceRecommendation struct {
	// Specifies whether automatic business name generation is to be enabled or not as part of the recommendation configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/datazone_data_source#enable_business_name_generation DatazoneDataSource#enable_business_name_generation}
	EnableBusinessNameGeneration interface{} `field:"optional" json:"enableBusinessNameGeneration" yaml:"enableBusinessNameGeneration"`
}

