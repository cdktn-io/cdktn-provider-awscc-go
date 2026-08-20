// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cleanroomsmembership


type CleanroomsMembershipDefaultResultConfigurationOutputConfigurationS3 struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cleanrooms_membership#bucket CleanroomsMembership#bucket}.
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cleanrooms_membership#key_prefix CleanroomsMembership#key_prefix}.
	KeyPrefix *string `field:"optional" json:"keyPrefix" yaml:"keyPrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cleanrooms_membership#result_format CleanroomsMembership#result_format}.
	ResultFormat *string `field:"optional" json:"resultFormat" yaml:"resultFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cleanrooms_membership#single_file_output CleanroomsMembership#single_file_output}.
	SingleFileOutput interface{} `field:"optional" json:"singleFileOutput" yaml:"singleFileOutput"`
}

