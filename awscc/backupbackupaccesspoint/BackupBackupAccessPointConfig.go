// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package backupbackupaccesspoint

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BackupBackupAccessPointConfig struct {
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
	// The Amazon Resource Name (ARN) of the recovery point for which to create the backup access point.
	//
	// The recovery point must be an Amazon S3 recovery point in the AVAILABLE, STOPPED, or COMPLETED state.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/backup_backup_access_point#recovery_point_arn BackupBackupAccessPoint#recovery_point_arn}
	RecoveryPointArn *string `field:"required" json:"recoveryPointArn" yaml:"recoveryPointArn"`
	// For continuous (point-in-time) recovery points, the timestamp in ISO 8601 format (for example 2021-11-27T03:30:27Z) that determines the point-in-time view of the backup data.
	//
	// The timestamp must be within the continuous backup's retention period at the time the access point is created. Do not specify this for snapshot recovery points. The service reports it back under AccessPointMetadata.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/backup_backup_access_point#access_point_in_time BackupBackupAccessPoint#access_point_in_time}
	AccessPointInTime *string `field:"optional" json:"accessPointInTime" yaml:"accessPointInTime"`
	// An optional resource-based policy document, in JSON, to apply to the underlying Amazon S3 access point.
	//
	// The policy controls how backup data can be accessed through the access point. If you do not specify a policy, access is governed by the caller's IAM permissions. The serialized policy may not exceed 200000 characters. This value is not returned by the service and cannot be updated in place.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/backup_backup_access_point#access_point_policy BackupBackupAccessPoint#access_point_policy}
	AccessPointPolicy *string `field:"optional" json:"accessPointPolicy" yaml:"accessPointPolicy"`
	// The name of the backup access point.
	//
	// This name is shared with the Amazon S3 access point namespace. It must be unique within your account and Region and cannot conflict with an existing Amazon S3 access point. If omitted, a name is generated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/backup_backup_access_point#name BackupBackupAccessPoint#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The tags to assign to the backup access point.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/backup_backup_access_point#tags BackupBackupAccessPoint#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

