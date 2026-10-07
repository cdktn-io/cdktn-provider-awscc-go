// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package redshiftredshiftidcapplication


type RedshiftRedshiftIdcApplicationServiceIntegrationsS3AccessGrants struct {
	// The S3 Access Grants scope.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/redshift_redshift_idc_application#read_write_access RedshiftRedshiftIdcApplication#read_write_access}
	ReadWriteAccess *RedshiftRedshiftIdcApplicationServiceIntegrationsS3AccessGrantsReadWriteAccess `field:"optional" json:"readWriteAccess" yaml:"readWriteAccess"`
}

