// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package athenaworkgroup


type AthenaWorkGroupWorkGroupConfigurationResultConfigurationAclConfiguration struct {
	// The Amazon S3 canned ACL that Athena should specify when storing query results.
	//
	// Currently the only supported canned ACL is BUCKET_OWNER_FULL_CONTROL
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/athena_work_group#s3_acl_option AthenaWorkGroup#s3_acl_option}
	S3AclOption *string `field:"optional" json:"s3AclOption" yaml:"s3AclOption"`
}

