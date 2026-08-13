// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package databrewdataset


type DatabrewDatasetPathOptionsParametersDatasetParameterDatetimeOptions struct {
	// Date/time format of a date parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/databrew_dataset#format DatabrewDataset#format}
	Format *string `field:"optional" json:"format" yaml:"format"`
	// Locale code for a date parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/databrew_dataset#locale_code DatabrewDataset#locale_code}
	LocaleCode *string `field:"optional" json:"localeCode" yaml:"localeCode"`
	// Timezone offset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/databrew_dataset#timezone_offset DatabrewDataset#timezone_offset}
	TimezoneOffset *string `field:"optional" json:"timezoneOffset" yaml:"timezoneOffset"`
}

