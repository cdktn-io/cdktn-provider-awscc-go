// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package redshiftserverlessnamespace

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type RedshiftserverlessNamespaceConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// A unique identifier for the namespace.
	//
	// You use this identifier to refer to the namespace for any subsequent namespace operations such as deleting or modifying. All alphabetical characters must be lower case. Namespace name should be unique for all namespaces within an AWS account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/redshiftserverless_namespace#namespace_name RedshiftserverlessNamespace#namespace_name}
	NamespaceName *string `field:"required" json:"namespaceName" yaml:"namespaceName"`
	// The ID of the AWS Key Management Service (KMS) key used to encrypt and store the namespace's admin credentials secret.
	//
	// You can only use this parameter if manageAdminPassword is true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/redshiftserverless_namespace#admin_password_secret_kms_key_id RedshiftserverlessNamespace#admin_password_secret_kms_key_id}
	AdminPasswordSecretKmsKeyId *string `field:"optional" json:"adminPasswordSecretKmsKeyId" yaml:"adminPasswordSecretKmsKeyId"`
	// The user name associated with the admin user for the namespace that is being created.
	//
	// Only alphanumeric characters and underscores are allowed. It should start with an alphabet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/redshiftserverless_namespace#admin_username RedshiftserverlessNamespace#admin_username}
	AdminUsername *string `field:"optional" json:"adminUsername" yaml:"adminUsername"`
	// The password associated with the admin user for the namespace that is being created.
	//
	// Password must be at least 8 characters in length, should be any printable ASCII character. Must contain at least one lowercase letter, one uppercase letter and one decimal digit. You can't use adminUserPassword if manageAdminPassword is true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/redshiftserverless_namespace#admin_user_password RedshiftserverlessNamespace#admin_user_password}
	AdminUserPassword *string `field:"optional" json:"adminUserPassword" yaml:"adminUserPassword"`
	// The database name associated for the namespace that is being created.
	//
	// Only alphanumeric characters and underscores are allowed. It should start with an alphabet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/redshiftserverless_namespace#db_name RedshiftserverlessNamespace#db_name}
	DbName *string `field:"optional" json:"dbName" yaml:"dbName"`
	// The default IAM role ARN for the namespace that is being created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/redshiftserverless_namespace#default_iam_role_arn RedshiftserverlessNamespace#default_iam_role_arn}
	DefaultIamRoleArn *string `field:"optional" json:"defaultIamRoleArn" yaml:"defaultIamRoleArn"`
	// The name of the namespace the source snapshot was created from.
	//
	// Please specify the name if needed before deleting namespace
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/redshiftserverless_namespace#final_snapshot_name RedshiftserverlessNamespace#final_snapshot_name}
	FinalSnapshotName *string `field:"optional" json:"finalSnapshotName" yaml:"finalSnapshotName"`
	// The number of days to retain automated snapshot in the destination region after they are copied from the source region.
	//
	// If the value is -1, the manual snapshot is retained indefinitely. The value must be either -1 or an integer between 1 and 3,653.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/redshiftserverless_namespace#final_snapshot_retention_period RedshiftserverlessNamespace#final_snapshot_retention_period}
	FinalSnapshotRetentionPeriod *float64 `field:"optional" json:"finalSnapshotRetentionPeriod" yaml:"finalSnapshotRetentionPeriod"`
	// A list of AWS Identity and Access Management (IAM) roles that can be used by the namespace to access other AWS services.
	//
	// You must supply the IAM roles in their Amazon Resource Name (ARN) format. The Default role limit for each request is 10.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/redshiftserverless_namespace#iam_roles RedshiftserverlessNamespace#iam_roles}
	IamRoles *[]*string `field:"optional" json:"iamRoles" yaml:"iamRoles"`
	// The AWS Key Management Service (KMS) key ID of the encryption key that you want to use to encrypt data in the namespace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/redshiftserverless_namespace#kms_key_id RedshiftserverlessNamespace#kms_key_id}
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The collection of log types to be exported provided by the customer.
	//
	// Should only be one of the three supported log types: userlog, useractivitylog and connectionlog
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/redshiftserverless_namespace#log_exports RedshiftserverlessNamespace#log_exports}
	LogExports *[]*string `field:"optional" json:"logExports" yaml:"logExports"`
	// If true, Amazon Redshift uses AWS Secrets Manager to manage the namespace's admin credentials.
	//
	// You can't use adminUserPassword if manageAdminPassword is true. If manageAdminPassword is false or not set, Amazon Redshift uses adminUserPassword for the admin user account's password.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/redshiftserverless_namespace#manage_admin_password RedshiftserverlessNamespace#manage_admin_password}
	ManageAdminPassword interface{} `field:"optional" json:"manageAdminPassword" yaml:"manageAdminPassword"`
	// The resource policy document that will be attached to the namespace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/redshiftserverless_namespace#namespace_resource_policy RedshiftserverlessNamespace#namespace_resource_policy}
	NamespaceResourcePolicy *string `field:"optional" json:"namespaceResourcePolicy" yaml:"namespaceResourcePolicy"`
	// The ARN for the Redshift application that integrates with IAM Identity Center.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/redshiftserverless_namespace#redshift_idc_application_arn RedshiftserverlessNamespace#redshift_idc_application_arn}
	RedshiftIdcApplicationArn *string `field:"optional" json:"redshiftIdcApplicationArn" yaml:"redshiftIdcApplicationArn"`
	// The snapshot copy configurations for the namespace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/redshiftserverless_namespace#snapshot_copy_configurations RedshiftserverlessNamespace#snapshot_copy_configurations}
	SnapshotCopyConfigurations interface{} `field:"optional" json:"snapshotCopyConfigurations" yaml:"snapshotCopyConfigurations"`
	// The list of tags for the namespace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/redshiftserverless_namespace#tags RedshiftserverlessNamespace#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

