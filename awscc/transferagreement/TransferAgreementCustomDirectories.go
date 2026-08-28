// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package transferagreement


type TransferAgreementCustomDirectories struct {
	// Specifies a location to store the failed files for an AS2 message.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/transfer_agreement#failed_files_directory TransferAgreement#failed_files_directory}
	FailedFilesDirectory *string `field:"optional" json:"failedFilesDirectory" yaml:"failedFilesDirectory"`
	// Specifies a location to store the MDN file for an AS2 message.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/transfer_agreement#mdn_files_directory TransferAgreement#mdn_files_directory}
	MdnFilesDirectory *string `field:"optional" json:"mdnFilesDirectory" yaml:"mdnFilesDirectory"`
	// Specifies a location to store the payload file for an AS2 message.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/transfer_agreement#payload_files_directory TransferAgreement#payload_files_directory}
	PayloadFilesDirectory *string `field:"optional" json:"payloadFilesDirectory" yaml:"payloadFilesDirectory"`
	// Specifies a location to store the status file for an AS2 message.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/transfer_agreement#status_files_directory TransferAgreement#status_files_directory}
	StatusFilesDirectory *string `field:"optional" json:"statusFilesDirectory" yaml:"statusFilesDirectory"`
	// Specifies a location to store the temporary processing file for an AS2 message.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/transfer_agreement#temporary_files_directory TransferAgreement#temporary_files_directory}
	TemporaryFilesDirectory *string `field:"optional" json:"temporaryFilesDirectory" yaml:"temporaryFilesDirectory"`
}

