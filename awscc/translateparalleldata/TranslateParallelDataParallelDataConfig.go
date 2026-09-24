// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package translateparalleldata


type TranslateParallelDataParallelDataConfig struct {
	// The format of the parallel data input file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/translate_parallel_data#format TranslateParallelData#format}
	Format *string `field:"required" json:"format" yaml:"format"`
	// The URI of the Amazon S3 folder that contains the parallel data input file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/translate_parallel_data#s3_uri TranslateParallelData#s3_uri}
	S3Uri *string `field:"required" json:"s3Uri" yaml:"s3Uri"`
}

