// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightdatasource


type QuicksightDataSourceDataSourceParametersRedshiftParametersIamParameters struct {
	// <p>Automatically creates a database user.
	//
	// If your database doesn't have a <code>DatabaseUser</code>, set this parameter to <code>True</code>. If there is no <code>DatabaseUser</code>, Amazon QuickSight can't connect to your cluster. The <code>RoleArn</code> that you use for this operation must grant access to <code>redshift:CreateClusterUser</code> to successfully create the user.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/quicksight_data_source#auto_create_database_user QuicksightDataSource#auto_create_database_user}
	AutoCreateDatabaseUser interface{} `field:"optional" json:"autoCreateDatabaseUser" yaml:"autoCreateDatabaseUser"`
	// <p>A list of groups whose permissions will be granted to Amazon QuickSight to access the cluster.
	//
	// These permissions are combined with the permissions granted to Amazon QuickSight by the <code>DatabaseUser</code>. If you choose to include this parameter, the <code>RoleArn</code> must grant access to <code>redshift:JoinGroup</code>.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/quicksight_data_source#database_groups QuicksightDataSource#database_groups}
	DatabaseGroups *[]*string `field:"optional" json:"databaseGroups" yaml:"databaseGroups"`
	// <p>The user whose permissions and group memberships will be used by Amazon QuickSight to access the cluster.
	//
	// If this user already exists in your database, Amazon QuickSight is granted the same permissions that the user has. If the user doesn't exist, set the value of <code>AutoCreateDatabaseUser</code> to <code>True</code> to create a new user with PUBLIC permissions.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/quicksight_data_source#database_user QuicksightDataSource#database_user}
	DatabaseUser *string `field:"optional" json:"databaseUser" yaml:"databaseUser"`
	// <p>Use the <code>RoleArn</code> structure to allow Amazon QuickSight to call <code>redshift:GetClusterCredentials</code> on your cluster.
	//
	// The calling principal must have <code>iam:PassRole</code> access to pass the role to Amazon QuickSight. The role's trust policy must allow the Amazon QuickSight service principal to assume the role.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/quicksight_data_source#role_arn QuicksightDataSource#role_arn}
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

