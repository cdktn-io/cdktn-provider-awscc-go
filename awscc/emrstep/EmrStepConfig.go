// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package emrstep

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EmrStepConfig struct {
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
	// This specifies what action to take when the cluster step fails. Possible values are CANCEL_AND_WAIT and CONTINUE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/emr_step#action_on_failure EmrStep#action_on_failure}
	ActionOnFailure *string `field:"required" json:"actionOnFailure" yaml:"actionOnFailure"`
	// The HadoopJarStepConfig property type specifies a job flow step consisting of a JAR file whose main function will be executed.
	//
	// The main function submits a job for the cluster to execute as a step on the master node, and then waits for the job to finish or fail before executing subsequent steps.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/emr_step#hadoop_jar_step EmrStep#hadoop_jar_step}
	HadoopJarStep *EmrStepHadoopJarStep `field:"required" json:"hadoopJarStep" yaml:"hadoopJarStep"`
	// A string that uniquely identifies the cluster (job flow).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/emr_step#job_flow_id EmrStep#job_flow_id}
	JobFlowId *string `field:"required" json:"jobFlowId" yaml:"jobFlowId"`
	// The name of the cluster step.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/emr_step#name EmrStep#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The KMS key ARN to encrypt the logs published to the given Amazon S3 destination.
	//
	// When omitted, EMR falls back to cluster-level logging behavior.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/emr_step#encryption_key_arn EmrStep#encryption_key_arn}
	EncryptionKeyArn *string `field:"optional" json:"encryptionKeyArn" yaml:"encryptionKeyArn"`
	// The Amazon S3 destination URI for log publishing. When omitted, EMR falls back to cluster-level logging behavior.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/emr_step#log_uri EmrStep#log_uri}
	LogUri *string `field:"optional" json:"logUri" yaml:"logUri"`
}

