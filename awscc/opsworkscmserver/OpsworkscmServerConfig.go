// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package opsworkscmserver

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OpsworkscmServerConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/opsworkscm_server#instance_profile_arn OpsworkscmServer#instance_profile_arn}.
	InstanceProfileArn *string `field:"required" json:"instanceProfileArn" yaml:"instanceProfileArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/opsworkscm_server#instance_type OpsworkscmServer#instance_type}.
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/opsworkscm_server#service_role_arn OpsworkscmServer#service_role_arn}.
	ServiceRoleArn *string `field:"required" json:"serviceRoleArn" yaml:"serviceRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/opsworkscm_server#associate_public_ip_address OpsworkscmServer#associate_public_ip_address}.
	AssociatePublicIpAddress interface{} `field:"optional" json:"associatePublicIpAddress" yaml:"associatePublicIpAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/opsworkscm_server#backup_id OpsworkscmServer#backup_id}.
	BackupId *string `field:"optional" json:"backupId" yaml:"backupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/opsworkscm_server#backup_retention_count OpsworkscmServer#backup_retention_count}.
	BackupRetentionCount *float64 `field:"optional" json:"backupRetentionCount" yaml:"backupRetentionCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/opsworkscm_server#custom_certificate OpsworkscmServer#custom_certificate}.
	CustomCertificate *string `field:"optional" json:"customCertificate" yaml:"customCertificate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/opsworkscm_server#custom_domain OpsworkscmServer#custom_domain}.
	CustomDomain *string `field:"optional" json:"customDomain" yaml:"customDomain"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/opsworkscm_server#custom_private_key OpsworkscmServer#custom_private_key}.
	CustomPrivateKey *string `field:"optional" json:"customPrivateKey" yaml:"customPrivateKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/opsworkscm_server#disable_automated_backup OpsworkscmServer#disable_automated_backup}.
	DisableAutomatedBackup interface{} `field:"optional" json:"disableAutomatedBackup" yaml:"disableAutomatedBackup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/opsworkscm_server#engine OpsworkscmServer#engine}.
	Engine *string `field:"optional" json:"engine" yaml:"engine"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/opsworkscm_server#engine_attributes OpsworkscmServer#engine_attributes}.
	EngineAttributes interface{} `field:"optional" json:"engineAttributes" yaml:"engineAttributes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/opsworkscm_server#engine_model OpsworkscmServer#engine_model}.
	EngineModel *string `field:"optional" json:"engineModel" yaml:"engineModel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/opsworkscm_server#engine_version OpsworkscmServer#engine_version}.
	EngineVersion *string `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/opsworkscm_server#key_pair OpsworkscmServer#key_pair}.
	KeyPair *string `field:"optional" json:"keyPair" yaml:"keyPair"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/opsworkscm_server#preferred_backup_window OpsworkscmServer#preferred_backup_window}.
	PreferredBackupWindow *string `field:"optional" json:"preferredBackupWindow" yaml:"preferredBackupWindow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/opsworkscm_server#preferred_maintenance_window OpsworkscmServer#preferred_maintenance_window}.
	PreferredMaintenanceWindow *string `field:"optional" json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/opsworkscm_server#security_group_ids OpsworkscmServer#security_group_ids}.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/opsworkscm_server#server_name OpsworkscmServer#server_name}.
	ServerName *string `field:"optional" json:"serverName" yaml:"serverName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/opsworkscm_server#subnet_ids OpsworkscmServer#subnet_ids}.
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/opsworkscm_server#tags OpsworkscmServer#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

