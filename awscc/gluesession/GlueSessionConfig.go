// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluesession

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GlueSessionConfig struct {
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
	// The SessionCommand that runs the job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_session#command GlueSession#command}
	Command *GlueSessionCommand `field:"required" json:"command" yaml:"command"`
	// The IAM Role ARN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_session#role GlueSession#role}
	Role *string `field:"required" json:"role" yaml:"role"`
	// The ID of the session.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_session#session_id GlueSession#session_id}
	SessionId *string `field:"required" json:"sessionId" yaml:"sessionId"`
	// Specifies the connections used by the session.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_session#connections GlueSession#connections}
	Connections *GlueSessionConnections `field:"optional" json:"connections" yaml:"connections"`
	// A map array of key-value pairs. Max is 75 pairs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_session#default_arguments GlueSession#default_arguments}
	DefaultArguments *map[string]*string `field:"optional" json:"defaultArguments" yaml:"defaultArguments"`
	// The description of the session.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_session#description GlueSession#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The Glue version determines the versions of Apache Spark and Python that Glue supports.
	//
	// The GlueVersion must be greater than 2.0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_session#glue_version GlueSession#glue_version}
	GlueVersion *string `field:"optional" json:"glueVersion" yaml:"glueVersion"`
	// The number of minutes when idle before session times out. Default is the value of Timeout.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_session#idle_timeout GlueSession#idle_timeout}
	IdleTimeout *float64 `field:"optional" json:"idleTimeout" yaml:"idleTimeout"`
	// The number of Glue data processing units (DPUs) that can be allocated when the job runs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_session#max_capacity GlueSession#max_capacity}
	MaxCapacity *float64 `field:"optional" json:"maxCapacity" yaml:"maxCapacity"`
	// The number of workers of a defined WorkerType to use for the session.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_session#number_of_workers GlueSession#number_of_workers}
	NumberOfWorkers *float64 `field:"optional" json:"numberOfWorkers" yaml:"numberOfWorkers"`
	// The origin of the request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_session#request_origin GlueSession#request_origin}
	RequestOrigin *string `field:"optional" json:"requestOrigin" yaml:"requestOrigin"`
	// The name of the SecurityConfiguration structure to be used with the session.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_session#security_configuration GlueSession#security_configuration}
	SecurityConfiguration *string `field:"optional" json:"securityConfiguration" yaml:"securityConfiguration"`
	// The tags belonging to the session.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_session#tags GlueSession#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The number of minutes before session times out.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_session#timeout GlueSession#timeout}
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
	// The type of predefined worker that is allocated when a session runs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_session#worker_type GlueSession#worker_type}
	WorkerType *string `field:"optional" json:"workerType" yaml:"workerType"`
}

