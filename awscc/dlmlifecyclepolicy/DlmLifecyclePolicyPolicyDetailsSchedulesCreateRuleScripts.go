// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dlmlifecyclepolicy


type DlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleScripts struct {
	// Indicates whether Amazon Data Lifecycle Manager should default to crash-consistent snapshots if the pre script fails.
	//
	// - To default to crash consistent snapshot if the pre script fails, specify `true`.
	// - To skip the instance for snapshot creation if the pre script fails, specify `false`.
	//
	// This parameter is supported only if you run a pre script. If you run a post script only, omit this parameter.
	//
	// Default: `true`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dlm_lifecycle_policy#execute_operation_on_script_failure DlmLifecyclePolicy#execute_operation_on_script_failure}
	ExecuteOperationOnScriptFailure interface{} `field:"optional" json:"executeOperationOnScriptFailure" yaml:"executeOperationOnScriptFailure"`
	// The SSM document that includes the pre and/or post scripts to run.
	//
	// If you are automating VSS backups, specify `AWS_VSS_BACKUP`. In this case, Amazon Data Lifecycle Manager automatically uses the `AWSEC2-CreateVssSnapshot` SSM document.
	//
	// If you are using a custom SSM document that you own, specify either the name or ARN of the SSM document. If you are using a custom SSM document that is shared with you, specify the ARN of the SSM document.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dlm_lifecycle_policy#execution_handler DlmLifecyclePolicy#execution_handler}
	ExecutionHandler *string `field:"optional" json:"executionHandler" yaml:"executionHandler"`
	// Indicates the service used to execute the pre and/or post scripts.
	//
	// Default: `AWS_SYSTEMS_MANAGER`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dlm_lifecycle_policy#execution_handler_service DlmLifecyclePolicy#execution_handler_service}
	ExecutionHandlerService *string `field:"optional" json:"executionHandlerService" yaml:"executionHandlerService"`
	// Specifies a timeout period, in seconds, after which Amazon Data Lifecycle Manager fails the script run attempt if it has not completed.
	//
	// If a script does not complete within its timeout period, Amazon Data Lifecycle Manager fails the attempt. The timeout period applies to the pre and post scripts individually.
	//
	// Default: 10
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dlm_lifecycle_policy#execution_timeout DlmLifecyclePolicy#execution_timeout}
	ExecutionTimeout *float64 `field:"optional" json:"executionTimeout" yaml:"executionTimeout"`
	// Specifies the number of times Amazon Data Lifecycle Manager should retry scripts that fail.
	//
	// If the pre script fails, Amazon Data Lifecycle Manager retries the entire snapshot creation process, including running the pre and post scripts.
	//
	// If the post script fails, Amazon Data Lifecycle Manager retries the post script only; in this case, the pre script will have completed and the snapshot might have been created.
	//
	// If you do not want Amazon Data Lifecycle Manager to retry failed scripts, specify `0`.
	//
	// Default: 0
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dlm_lifecycle_policy#maximum_retry_count DlmLifecyclePolicy#maximum_retry_count}
	MaximumRetryCount *float64 `field:"optional" json:"maximumRetryCount" yaml:"maximumRetryCount"`
	// Indicate which scripts Amazon Data Lifecycle Manager should run on target instances.
	//
	// Pre scripts run before Amazon Data Lifecycle Manager initiates snapshot creation. Post scripts run after Amazon Data Lifecycle Manager initiates snapshot creation.
	//
	// - To run a pre script only, specify `PRE`.
	// - To run a post script only, specify `POST`.
	// - To run both pre and post scripts, specify both `PRE` and `POST`.
	//
	// Default: PRE and POST
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dlm_lifecycle_policy#stages DlmLifecyclePolicy#stages}
	Stages *[]*string `field:"optional" json:"stages" yaml:"stages"`
}

