// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datazoneconnection


type DatazoneConnectionProps struct {
	// Amazon Q properties of the connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/datazone_connection#amazon_q_properties DatazoneConnection#amazon_q_properties}
	AmazonQProperties *DatazoneConnectionPropsAmazonQProperties `field:"optional" json:"amazonQProperties" yaml:"amazonQProperties"`
	// Athena Properties Input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/datazone_connection#athena_properties DatazoneConnection#athena_properties}
	AthenaProperties *DatazoneConnectionPropsAthenaProperties `field:"optional" json:"athenaProperties" yaml:"athenaProperties"`
	// Glue Properties Input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/datazone_connection#glue_properties DatazoneConnection#glue_properties}
	GlueProperties *DatazoneConnectionPropsGlueProperties `field:"optional" json:"glueProperties" yaml:"glueProperties"`
	// HyperPod Properties Input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/datazone_connection#hyper_pod_properties DatazoneConnection#hyper_pod_properties}
	HyperPodProperties *DatazoneConnectionPropsHyperPodProperties `field:"optional" json:"hyperPodProperties" yaml:"hyperPodProperties"`
	// IAM Properties Input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/datazone_connection#iam_properties DatazoneConnection#iam_properties}
	IamProperties *DatazoneConnectionPropsIamProperties `field:"optional" json:"iamProperties" yaml:"iamProperties"`
	// Lakehouse Properties Input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/datazone_connection#lakehouse_properties DatazoneConnection#lakehouse_properties}
	LakehouseProperties *DatazoneConnectionPropsLakehouseProperties `field:"optional" json:"lakehouseProperties" yaml:"lakehouseProperties"`
	// MLflow Properties Input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/datazone_connection#mlflow_properties DatazoneConnection#mlflow_properties}
	MlflowProperties *DatazoneConnectionPropsMlflowProperties `field:"optional" json:"mlflowProperties" yaml:"mlflowProperties"`
	// Redshift Properties Input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/datazone_connection#redshift_properties DatazoneConnection#redshift_properties}
	RedshiftProperties *DatazoneConnectionPropsRedshiftProperties `field:"optional" json:"redshiftProperties" yaml:"redshiftProperties"`
	// S3 Properties Input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/datazone_connection#s3_properties DatazoneConnection#s3_properties}
	S3Properties *DatazoneConnectionPropsS3Properties `field:"optional" json:"s3Properties" yaml:"s3Properties"`
	// Spark EMR Properties Input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/datazone_connection#spark_emr_properties DatazoneConnection#spark_emr_properties}
	SparkEmrProperties *DatazoneConnectionPropsSparkEmrProperties `field:"optional" json:"sparkEmrProperties" yaml:"sparkEmrProperties"`
	// Spark Glue Properties Input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/datazone_connection#spark_glue_properties DatazoneConnection#spark_glue_properties}
	SparkGlueProperties *DatazoneConnectionPropsSparkGlueProperties `field:"optional" json:"sparkGlueProperties" yaml:"sparkGlueProperties"`
	// Workflows MWAA Properties Input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/datazone_connection#workflows_mwaa_properties DatazoneConnection#workflows_mwaa_properties}
	WorkflowsMwaaProperties *DatazoneConnectionPropsWorkflowsMwaaProperties `field:"optional" json:"workflowsMwaaProperties" yaml:"workflowsMwaaProperties"`
	// Workflows Serverless Properties Input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/datazone_connection#workflows_serverless_properties DatazoneConnection#workflows_serverless_properties}
	WorkflowsServerlessProperties *string `field:"optional" json:"workflowsServerlessProperties" yaml:"workflowsServerlessProperties"`
}

