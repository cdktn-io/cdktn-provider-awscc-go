// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package m2environment

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type M2EnvironmentConfig struct {
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
	// The target platform for the environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/m2_environment#engine_type M2Environment#engine_type}
	EngineType *string `field:"required" json:"engineType" yaml:"engineType"`
	// The type of instance underlying the environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/m2_environment#instance_type M2Environment#instance_type}
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
	// The name of the environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/m2_environment#name M2Environment#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The description of the environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/m2_environment#description M2Environment#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The version of the runtime engine for the environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/m2_environment#engine_version M2Environment#engine_version}
	EngineVersion *string `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	// Defines the details of a high availability configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/m2_environment#high_availability_config M2Environment#high_availability_config}
	HighAvailabilityConfig *M2EnvironmentHighAvailabilityConfig `field:"optional" json:"highAvailabilityConfig" yaml:"highAvailabilityConfig"`
	// The ID or the Amazon Resource Name (ARN) of the customer managed KMS Key used for encrypting environment-related resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/m2_environment#kms_key_id M2Environment#kms_key_id}
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/m2_environment#network_type M2Environment#network_type}.
	NetworkType *string `field:"optional" json:"networkType" yaml:"networkType"`
	// Configures a desired maintenance window for the environment.
	//
	// If you do not provide a value, a random system-generated value will be assigned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/m2_environment#preferred_maintenance_window M2Environment#preferred_maintenance_window}
	PreferredMaintenanceWindow *string `field:"optional" json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// Specifies whether the environment is publicly accessible.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/m2_environment#publicly_accessible M2Environment#publicly_accessible}
	PubliclyAccessible interface{} `field:"optional" json:"publiclyAccessible" yaml:"publiclyAccessible"`
	// The list of security groups for the VPC associated with this environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/m2_environment#security_group_ids M2Environment#security_group_ids}
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// The storage configurations defined for the runtime environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/m2_environment#storage_configurations M2Environment#storage_configurations}
	StorageConfigurations interface{} `field:"optional" json:"storageConfigurations" yaml:"storageConfigurations"`
	// The unique identifiers of the subnets assigned to this runtime environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/m2_environment#subnet_ids M2Environment#subnet_ids}
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
	// Tags associated to this environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/m2_environment#tags M2Environment#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

