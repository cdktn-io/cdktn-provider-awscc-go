// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package eventsv2subscriber


type Eventsv2SubscriberInvokeConfigurationLambdaParameters struct {
	// A unique name for a durable function execution. Accepts a literal value or a JSONata expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/eventsv2_subscriber#durable_execution_name Eventsv2Subscriber#durable_execution_name}
	DurableExecutionName *string `field:"optional" json:"durableExecutionName" yaml:"durableExecutionName"`
	// The timeout in seconds for each invocation of the target, written as a string.
	//
	// Accepts a literal value or a JSONata expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/eventsv2_subscriber#invocation_timeout_seconds Eventsv2Subscriber#invocation_timeout_seconds}
	InvocationTimeoutSeconds *string `field:"optional" json:"invocationTimeoutSeconds" yaml:"invocationTimeoutSeconds"`
	// How the function is invoked: EVENT (asynchronous) or REQUEST_RESPONSE (synchronous).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/eventsv2_subscriber#invocation_type Eventsv2Subscriber#invocation_type}
	InvocationType *string `field:"optional" json:"invocationType" yaml:"invocationType"`
	// The version or alias of the Lambda function to invoke. Accepts a literal value or a JSONata expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/eventsv2_subscriber#qualifier Eventsv2Subscriber#qualifier}
	Qualifier *string `field:"optional" json:"qualifier" yaml:"qualifier"`
	// The tenant identifier for multi-tenant Lambda functions. Accepts a literal value or a JSONata expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/eventsv2_subscriber#tenant_id Eventsv2Subscriber#tenant_id}
	TenantId *string `field:"optional" json:"tenantId" yaml:"tenantId"`
}

