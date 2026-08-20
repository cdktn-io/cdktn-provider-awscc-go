// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package databrewdataset


type DatabrewDatasetPathOptionsParametersDatasetParameterFilter struct {
	// Filtering expression for a parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/databrew_dataset#expression DatabrewDataset#expression}
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/databrew_dataset#values_map DatabrewDataset#values_map}.
	ValuesMap interface{} `field:"optional" json:"valuesMap" yaml:"valuesMap"`
}

