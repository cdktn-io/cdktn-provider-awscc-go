// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3storagelensgroup


type S3StorageLensGroupFilter struct {
	// The Storage Lens group will include objects that match all of the specified filter values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/s3_storage_lens_group#and S3StorageLensGroup#and}
	And *S3StorageLensGroupFilterAnd `field:"optional" json:"and" yaml:"and"`
	// Filter to match any of the specified prefixes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/s3_storage_lens_group#match_any_prefix S3StorageLensGroup#match_any_prefix}
	MatchAnyPrefix *[]*string `field:"optional" json:"matchAnyPrefix" yaml:"matchAnyPrefix"`
	// Filter to match any of the specified suffixes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/s3_storage_lens_group#match_any_suffix S3StorageLensGroup#match_any_suffix}
	MatchAnySuffix *[]*string `field:"optional" json:"matchAnySuffix" yaml:"matchAnySuffix"`
	// Filter to match any of the specified object tags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/s3_storage_lens_group#match_any_tag S3StorageLensGroup#match_any_tag}
	MatchAnyTag interface{} `field:"optional" json:"matchAnyTag" yaml:"matchAnyTag"`
	// Filter to match all of the specified values for the minimum and maximum object age.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/s3_storage_lens_group#match_object_age S3StorageLensGroup#match_object_age}
	MatchObjectAge *S3StorageLensGroupFilterMatchObjectAge `field:"optional" json:"matchObjectAge" yaml:"matchObjectAge"`
	// Filter to match all of the specified values for the minimum and maximum object size.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/s3_storage_lens_group#match_object_size S3StorageLensGroup#match_object_size}
	MatchObjectSize *S3StorageLensGroupFilterMatchObjectSize `field:"optional" json:"matchObjectSize" yaml:"matchObjectSize"`
	// The Storage Lens group will include objects that match any of the specified filter values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/s3_storage_lens_group#or S3StorageLensGroup#or}
	Or *S3StorageLensGroupFilterOr `field:"optional" json:"or" yaml:"or"`
}

