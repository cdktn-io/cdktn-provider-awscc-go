// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package rdsdbcluster


type RdsDbClusterAssociatedRoles struct {
	// The name of the feature associated with the AWS Identity and Access Management (IAM) role.
	//
	// IAM roles that are associated with a DB cluster grant permission for the DB cluster to access other AWS services on your behalf. For the list of supported feature names, see the ``SupportedFeatureNames`` description in [DBEngineVersion](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_DBEngineVersion.html) in the *Amazon RDS API Reference*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/rds_db_cluster#feature_name RdsDbCluster#feature_name}
	FeatureName *string `field:"optional" json:"featureName" yaml:"featureName"`
	// The Amazon Resource Name (ARN) of the IAM role that is associated with the DB cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/rds_db_cluster#role_arn RdsDbCluster#role_arn}
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

