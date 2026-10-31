// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iototaupdate


type IotOtaUpdateFilesCodeSigningStartSigningJobParameter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.108.0/docs/resources/iot_ota_update#destination IotOtaUpdate#destination}.
	Destination *IotOtaUpdateFilesCodeSigningStartSigningJobParameterDestination `field:"optional" json:"destination" yaml:"destination"`
	// The code-signing profile name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.108.0/docs/resources/iot_ota_update#signing_profile_name IotOtaUpdate#signing_profile_name}
	SigningProfileName *string `field:"optional" json:"signingProfileName" yaml:"signingProfileName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.108.0/docs/resources/iot_ota_update#signing_profile_parameter IotOtaUpdate#signing_profile_parameter}.
	SigningProfileParameter *IotOtaUpdateFilesCodeSigningStartSigningJobParameterSigningProfileParameter `field:"optional" json:"signingProfileParameter" yaml:"signingProfileParameter"`
}

