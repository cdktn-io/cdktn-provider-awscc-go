// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iototaupdate


type IotOtaUpdateFiles struct {
	// A list of name-attribute pairs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.108.0/docs/resources/iot_ota_update#attributes IotOtaUpdate#attributes}
	Attributes *map[string]*string `field:"optional" json:"attributes" yaml:"attributes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.108.0/docs/resources/iot_ota_update#code_signing IotOtaUpdate#code_signing}.
	CodeSigning *IotOtaUpdateFilesCodeSigning `field:"optional" json:"codeSigning" yaml:"codeSigning"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.108.0/docs/resources/iot_ota_update#file_location IotOtaUpdate#file_location}.
	FileLocation *IotOtaUpdateFilesFileLocation `field:"optional" json:"fileLocation" yaml:"fileLocation"`
	// The name of the file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.108.0/docs/resources/iot_ota_update#file_name IotOtaUpdate#file_name}
	FileName *string `field:"optional" json:"fileName" yaml:"fileName"`
	// An integer value you can include in the job document to allow your devices to identify the type of file received from the cloud.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.108.0/docs/resources/iot_ota_update#file_type IotOtaUpdate#file_type}
	FileType *float64 `field:"optional" json:"fileType" yaml:"fileType"`
	// The file version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.108.0/docs/resources/iot_ota_update#file_version IotOtaUpdate#file_version}
	FileVersion *string `field:"optional" json:"fileVersion" yaml:"fileVersion"`
}

