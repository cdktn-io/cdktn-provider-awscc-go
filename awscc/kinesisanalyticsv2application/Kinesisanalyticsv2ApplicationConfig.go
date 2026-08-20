// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kinesisanalyticsv2application

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Kinesisanalyticsv2ApplicationConfig struct {
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
	// The runtime environment for the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/kinesisanalyticsv2_application#runtime_environment Kinesisanalyticsv2Application#runtime_environment}
	RuntimeEnvironment *string `field:"required" json:"runtimeEnvironment" yaml:"runtimeEnvironment"`
	// Specifies the IAM role that the application uses to access external resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/kinesisanalyticsv2_application#service_execution_role Kinesisanalyticsv2Application#service_execution_role}
	ServiceExecutionRole *string `field:"required" json:"serviceExecutionRole" yaml:"serviceExecutionRole"`
	// Use this parameter to configure the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/kinesisanalyticsv2_application#application_configuration Kinesisanalyticsv2Application#application_configuration}
	ApplicationConfiguration *Kinesisanalyticsv2ApplicationApplicationConfiguration `field:"optional" json:"applicationConfiguration" yaml:"applicationConfiguration"`
	// The description of the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/kinesisanalyticsv2_application#application_description Kinesisanalyticsv2Application#application_description}
	ApplicationDescription *string `field:"optional" json:"applicationDescription" yaml:"applicationDescription"`
	// Used to configure start of maintenance window.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/kinesisanalyticsv2_application#application_maintenance_configuration Kinesisanalyticsv2Application#application_maintenance_configuration}
	ApplicationMaintenanceConfiguration *Kinesisanalyticsv2ApplicationApplicationMaintenanceConfiguration `field:"optional" json:"applicationMaintenanceConfiguration" yaml:"applicationMaintenanceConfiguration"`
	// To create a Kinesis Data Analytics Studio notebook, you must set the mode to `INTERACTIVE`.
	//
	// However, for a Kinesis Data Analytics for Apache Flink application, the mode is optional.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/kinesisanalyticsv2_application#application_mode Kinesisanalyticsv2Application#application_mode}
	ApplicationMode *string `field:"optional" json:"applicationMode" yaml:"applicationMode"`
	// The name of the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/kinesisanalyticsv2_application#application_name Kinesisanalyticsv2Application#application_name}
	ApplicationName *string `field:"optional" json:"applicationName" yaml:"applicationName"`
	// Specifies run configuration (start parameters) of a Kinesis Data Analytics application. Evaluated on update for RUNNING applications an only.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/kinesisanalyticsv2_application#run_configuration Kinesisanalyticsv2Application#run_configuration}
	RunConfiguration *Kinesisanalyticsv2ApplicationRunConfiguration `field:"optional" json:"runConfiguration" yaml:"runConfiguration"`
	// A list of one or more tags to assign to the application.
	//
	// A tag is a key-value pair that identifies an application. Note that the maximum number of application tags includes system tags. The maximum number of user-defined application tags is 50.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/kinesisanalyticsv2_application#tags Kinesisanalyticsv2Application#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

