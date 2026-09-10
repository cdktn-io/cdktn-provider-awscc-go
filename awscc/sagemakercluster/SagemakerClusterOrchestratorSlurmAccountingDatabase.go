// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakercluster


type SagemakerClusterOrchestratorSlurmAccountingDatabase struct {
	// Hostname or endpoint of the accounting database, such as an RDS endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_cluster#endpoint SagemakerCluster#endpoint}
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
	// Name of the accounting database schema. Defaults to slurm_acct_db when omitted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_cluster#name SagemakerCluster#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// TCP port of the accounting database. Defaults to 3306 when omitted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_cluster#port SagemakerCluster#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// ARN of the Secrets Manager secret holding the database credentials.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_cluster#secret_arn SagemakerCluster#secret_arn}
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
}

