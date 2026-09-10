// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appintegrationsdataintegration


type AppintegrationsDataIntegrationScheduleConfig struct {
	// The start date for objects to import in the first flow run. Epoch or ISO timestamp format is supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/appintegrations_data_integration#first_execution_from AppintegrationsDataIntegration#first_execution_from}
	FirstExecutionFrom *string `field:"optional" json:"firstExecutionFrom" yaml:"firstExecutionFrom"`
	// The name of the object to pull from the data source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/appintegrations_data_integration#object AppintegrationsDataIntegration#object}
	Object *string `field:"optional" json:"object" yaml:"object"`
	// How often the data should be pulled from data source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/appintegrations_data_integration#schedule_expression AppintegrationsDataIntegration#schedule_expression}
	ScheduleExpression *string `field:"optional" json:"scheduleExpression" yaml:"scheduleExpression"`
}

