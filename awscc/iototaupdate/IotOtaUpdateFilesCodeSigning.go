// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iototaupdate


type IotOtaUpdateFilesCodeSigning struct {
	// The ID of the AWSSignerJob which was created to sign the file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.108.0/docs/resources/iot_ota_update#aws_signer_job_id IotOtaUpdate#aws_signer_job_id}
	AwsSignerJobId *string `field:"optional" json:"awsSignerJobId" yaml:"awsSignerJobId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.108.0/docs/resources/iot_ota_update#custom_code_signing IotOtaUpdate#custom_code_signing}.
	CustomCodeSigning *IotOtaUpdateFilesCodeSigningCustomCodeSigning `field:"optional" json:"customCodeSigning" yaml:"customCodeSigning"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.108.0/docs/resources/iot_ota_update#start_signing_job_parameter IotOtaUpdate#start_signing_job_parameter}.
	StartSigningJobParameter *IotOtaUpdateFilesCodeSigningStartSigningJobParameter `field:"optional" json:"startSigningJobParameter" yaml:"startSigningJobParameter"`
}

