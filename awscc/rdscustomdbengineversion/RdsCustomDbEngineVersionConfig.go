// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package rdscustomdbengineversion

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type RdsCustomDbEngineVersionConfig struct {
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
	// The database engine to use for your custom engine version (CEV).
	//
	// Valid values:
	//   +   ``custom-oracle-ee``
	//   +   ``custom-oracle-ee-cdb``
	//   +   ``sqlserver-dev-ee``
	//   +   ``sqlserver-ee``
	//   +   ``sqlserver-se``
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/rds_custom_db_engine_version#engine RdsCustomDbEngineVersion#engine}
	Engine *string `field:"required" json:"engine" yaml:"engine"`
	// The name of your CEV.
	//
	// The name format is ``major version.customized_string``. For example, a valid CEV name is ``19.my_cev1``. This setting is required for RDS Custom for Oracle, but optional for Amazon RDS. The combination of ``Engine`` and ``EngineVersion`` is unique per customer per Region.
	//  *Constraints:* Minimum length is 1. Maximum length is 60.
	//  *Pattern:*``^[a-z0-9_.-]{1,60$``}
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/rds_custom_db_engine_version#engine_version RdsCustomDbEngineVersion#engine_version}
	EngineVersion *string `field:"required" json:"engineVersion" yaml:"engineVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/rds_custom_db_engine_version#database_installation_files RdsCustomDbEngineVersion#database_installation_files}.
	DatabaseInstallationFiles *[]*string `field:"optional" json:"databaseInstallationFiles" yaml:"databaseInstallationFiles"`
	// The name of an Amazon S3 bucket that contains database installation files for your CEV.
	//
	// For example, a valid bucket name is ``my-custom-installation-files``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/rds_custom_db_engine_version#database_installation_files_s3_bucket_name RdsCustomDbEngineVersion#database_installation_files_s3_bucket_name}
	DatabaseInstallationFilesS3BucketName *string `field:"optional" json:"databaseInstallationFilesS3BucketName" yaml:"databaseInstallationFilesS3BucketName"`
	// The Amazon S3 directory that contains the database installation files for your CEV.
	//
	// For example, a valid bucket name is ``123456789012/cev1``. If this setting isn't specified, no prefix is assumed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/rds_custom_db_engine_version#database_installation_files_s3_prefix RdsCustomDbEngineVersion#database_installation_files_s3_prefix}
	DatabaseInstallationFilesS3Prefix *string `field:"optional" json:"databaseInstallationFilesS3Prefix" yaml:"databaseInstallationFilesS3Prefix"`
	// An optional description of your CEV.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/rds_custom_db_engine_version#description RdsCustomDbEngineVersion#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A value that indicates the ID of the AMI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/rds_custom_db_engine_version#image_id RdsCustomDbEngineVersion#image_id}
	ImageId *string `field:"optional" json:"imageId" yaml:"imageId"`
	// The AWS KMS key identifier for an encrypted CEV.
	//
	// A symmetric encryption KMS key is required for RDS Custom, but optional for Amazon RDS.
	//  If you have an existing symmetric encryption KMS key in your account, you can use it with RDS Custom. No further action is necessary. If you don't already have a symmetric encryption KMS key in your account, follow the instructions in [Creating a symmetric encryption KMS key](https://docs.aws.amazon.com/kms/latest/developerguide/create-keys.html#create-symmetric-cmk) in the *Key Management Service Developer Guide*.
	//  You can choose the same symmetric encryption key when you create a CEV and a DB instance, or choose different keys.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/rds_custom_db_engine_version#kms_key_id RdsCustomDbEngineVersion#kms_key_id}
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The CEV manifest, which is a JSON document that describes the installation .zip files stored in Amazon S3. Specify the name/value pairs in a file or a quoted string. RDS Custom applies the patches in the order in which they are listed.  The following JSON fields are valid:   + MediaImportTemplateVersion Version of the CEV manifest. The date is in the format YYYY-MM-DD. + databaseInstallationFileNames Ordered list of installation files for the CEV. + opatchFileNames Ordered list of OPatch installers used for the Oracle DB engine. + psuRuPatchFileNames The PSU and RU patches for this CEV. + OtherPatchFileNames The patches that are not in the list of PSU and RU patches. Amazon RDS applies these patches after applying the PSU and RU patches.   For more information, see [Creating the CEV manifest](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-cev.html#custom-cev.preparing.manifest) in the *Amazon RDS User Guide*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/rds_custom_db_engine_version#manifest RdsCustomDbEngineVersion#manifest}
	Manifest *string `field:"optional" json:"manifest" yaml:"manifest"`
	// The ARN of a CEV to use as a source for creating a new CEV.
	//
	// You can specify a different Amazon Machine Imagine (AMI) by using either ``Source`` or ``UseAwsProvidedLatestImage``. You can't specify a different JSON manifest when you specify ``SourceCustomDbEngineVersionIdentifier``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/rds_custom_db_engine_version#source_custom_db_engine_version_identifier RdsCustomDbEngineVersion#source_custom_db_engine_version_identifier}
	SourceCustomDbEngineVersionIdentifier *string `field:"optional" json:"sourceCustomDbEngineVersionIdentifier" yaml:"sourceCustomDbEngineVersionIdentifier"`
	// A value that indicates the status of a custom engine version (CEV).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/rds_custom_db_engine_version#status RdsCustomDbEngineVersion#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
	// A list of tags. For more information, see [Tagging Amazon RDS Resources](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.html) in the *Amazon RDS User Guide.*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/rds_custom_db_engine_version#tags RdsCustomDbEngineVersion#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Specifies whether to use the latest service-provided Amazon Machine Image (AMI) for the CEV.
	//
	// If you specify ``UseAwsProvidedLatestImage``, you can't also specify ``ImageId``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/rds_custom_db_engine_version#use_aws_provided_latest_image RdsCustomDbEngineVersion#use_aws_provided_latest_image}
	UseAwsProvidedLatestImage interface{} `field:"optional" json:"useAwsProvidedLatestImage" yaml:"useAwsProvidedLatestImage"`
}

