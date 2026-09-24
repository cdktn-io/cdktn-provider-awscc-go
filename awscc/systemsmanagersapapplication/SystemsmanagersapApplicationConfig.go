// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package systemsmanagersapapplication

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SystemsmanagersapApplicationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/systemsmanagersap_application#application_id SystemsmanagersapApplication#application_id}.
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/systemsmanagersap_application#application_type SystemsmanagersapApplication#application_type}.
	ApplicationType *string `field:"required" json:"applicationType" yaml:"applicationType"`
	// This is an optional parameter for component details to which the SAP ABAP application is attached, such as Web Dispatcher.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/systemsmanagersap_application#components_info SystemsmanagersapApplication#components_info}
	ComponentsInfo interface{} `field:"optional" json:"componentsInfo" yaml:"componentsInfo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/systemsmanagersap_application#credentials SystemsmanagersapApplication#credentials}.
	Credentials interface{} `field:"optional" json:"credentials" yaml:"credentials"`
	// The ARN of the SAP HANA database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/systemsmanagersap_application#database_arn SystemsmanagersapApplication#database_arn}
	DatabaseArn *string `field:"optional" json:"databaseArn" yaml:"databaseArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/systemsmanagersap_application#instances SystemsmanagersapApplication#instances}.
	Instances *[]*string `field:"optional" json:"instances" yaml:"instances"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/systemsmanagersap_application#sap_instance_number SystemsmanagersapApplication#sap_instance_number}.
	SapInstanceNumber *string `field:"optional" json:"sapInstanceNumber" yaml:"sapInstanceNumber"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/systemsmanagersap_application#sid SystemsmanagersapApplication#sid}.
	Sid *string `field:"optional" json:"sid" yaml:"sid"`
	// The tags of a SystemsManagerSAP application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/systemsmanagersap_application#tags SystemsmanagersapApplication#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

