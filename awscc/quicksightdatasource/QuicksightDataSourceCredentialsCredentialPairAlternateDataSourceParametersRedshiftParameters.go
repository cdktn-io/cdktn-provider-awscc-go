// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightdatasource


type QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParameters struct {
	// <p>Cluster ID.
	//
	// This field can be blank if the <code>Host</code> and <code>Port</code> are
	//             provided.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/quicksight_data_source#cluster_id QuicksightDataSource#cluster_id}
	ClusterId *string `field:"optional" json:"clusterId" yaml:"clusterId"`
	// <p>Database.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/quicksight_data_source#database QuicksightDataSource#database}
	Database *string `field:"optional" json:"database" yaml:"database"`
	// <p>Host. This field can be blank if <code>ClusterId</code> is provided.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/quicksight_data_source#host QuicksightDataSource#host}
	Host *string `field:"optional" json:"host" yaml:"host"`
	// <p>A structure that grants Amazon QuickSight access to your cluster and make a call to the <code>redshift:GetClusterCredentials</code> API.
	//
	// For more information on the <code>redshift:GetClusterCredentials</code> API, see <a href="https://docs.aws.amazon.com/redshift/latest/APIReference/API_GetClusterCredentials.html">
	//                <code>GetClusterCredentials</code>
	//             </a>.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/quicksight_data_source#iam_parameters QuicksightDataSource#iam_parameters}
	IamParameters *QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersIamParameters `field:"optional" json:"iamParameters" yaml:"iamParameters"`
	// <p>The parameters for an IAM Identity Center configuration.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/quicksight_data_source#identity_center_configuration QuicksightDataSource#identity_center_configuration}
	IdentityCenterConfiguration *QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersIdentityCenterConfiguration `field:"optional" json:"identityCenterConfiguration" yaml:"identityCenterConfiguration"`
	// <p>Port. This field can be blank if the <code>ClusterId</code> is provided.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/quicksight_data_source#port QuicksightDataSource#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

