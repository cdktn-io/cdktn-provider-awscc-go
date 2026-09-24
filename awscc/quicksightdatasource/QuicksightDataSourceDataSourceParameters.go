// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightdatasource


type QuicksightDataSourceDataSourceParameters struct {
	// <p>The parameters for OpenSearch.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_data_source#amazon_elasticsearch_parameters QuicksightDataSource#amazon_elasticsearch_parameters}
	AmazonElasticsearchParameters *QuicksightDataSourceDataSourceParametersAmazonElasticsearchParameters `field:"optional" json:"amazonElasticsearchParameters" yaml:"amazonElasticsearchParameters"`
	// <p>The parameters for OpenSearch.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_data_source#amazon_open_search_parameters QuicksightDataSource#amazon_open_search_parameters}
	AmazonOpenSearchParameters *QuicksightDataSourceDataSourceParametersAmazonOpenSearchParameters `field:"optional" json:"amazonOpenSearchParameters" yaml:"amazonOpenSearchParameters"`
	// <p>Parameters for Amazon Athena.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_data_source#athena_parameters QuicksightDataSource#athena_parameters}
	AthenaParameters *QuicksightDataSourceDataSourceParametersAthenaParameters `field:"optional" json:"athenaParameters" yaml:"athenaParameters"`
	// <p>Parameters for Amazon Aurora.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_data_source#aurora_parameters QuicksightDataSource#aurora_parameters}
	AuroraParameters *QuicksightDataSourceDataSourceParametersAuroraParameters `field:"optional" json:"auroraParameters" yaml:"auroraParameters"`
	// <p>Parameters for Amazon Aurora PostgreSQL-Compatible Edition.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_data_source#aurora_postgre_sql_parameters QuicksightDataSource#aurora_postgre_sql_parameters}
	AuroraPostgreSqlParameters *QuicksightDataSourceDataSourceParametersAuroraPostgreSqlParameters `field:"optional" json:"auroraPostgreSqlParameters" yaml:"auroraPostgreSqlParameters"`
	// <p>The parameters that are required to connect to a Databricks data source.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_data_source#databricks_parameters QuicksightDataSource#databricks_parameters}
	DatabricksParameters *QuicksightDataSourceDataSourceParametersDatabricksParameters `field:"optional" json:"databricksParameters" yaml:"databricksParameters"`
	// <p>The parameters for MariaDB.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_data_source#maria_db_parameters QuicksightDataSource#maria_db_parameters}
	MariaDbParameters *QuicksightDataSourceDataSourceParametersMariaDbParameters `field:"optional" json:"mariaDbParameters" yaml:"mariaDbParameters"`
	// <p>The parameters for MySQL.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_data_source#my_sql_parameters QuicksightDataSource#my_sql_parameters}
	MySqlParameters *QuicksightDataSourceDataSourceParametersMySqlParameters `field:"optional" json:"mySqlParameters" yaml:"mySqlParameters"`
	// <p>The parameters for Oracle.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_data_source#oracle_parameters QuicksightDataSource#oracle_parameters}
	OracleParameters *QuicksightDataSourceDataSourceParametersOracleParameters `field:"optional" json:"oracleParameters" yaml:"oracleParameters"`
	// <p>The parameters for PostgreSQL.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_data_source#postgre_sql_parameters QuicksightDataSource#postgre_sql_parameters}
	PostgreSqlParameters *QuicksightDataSourceDataSourceParametersPostgreSqlParameters `field:"optional" json:"postgreSqlParameters" yaml:"postgreSqlParameters"`
	// <p>The parameters for Presto.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_data_source#presto_parameters QuicksightDataSource#presto_parameters}
	PrestoParameters *QuicksightDataSourceDataSourceParametersPrestoParameters `field:"optional" json:"prestoParameters" yaml:"prestoParameters"`
	// <p>The parameters for Amazon RDS.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_data_source#rds_parameters QuicksightDataSource#rds_parameters}
	RdsParameters *QuicksightDataSourceDataSourceParametersRdsParameters `field:"optional" json:"rdsParameters" yaml:"rdsParameters"`
	// <p>The parameters for Amazon Redshift.
	//
	// The <code>ClusterId</code> field can be blank if
	//             <code>Host</code> and <code>Port</code> are both set. The <code>Host</code> and <code>Port</code> fields can be blank if the <code>ClusterId</code> field is set.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_data_source#redshift_parameters QuicksightDataSource#redshift_parameters}
	RedshiftParameters *QuicksightDataSourceDataSourceParametersRedshiftParameters `field:"optional" json:"redshiftParameters" yaml:"redshiftParameters"`
	// <p>The parameters for S3.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_data_source#s3_parameters QuicksightDataSource#s3_parameters}
	S3Parameters *QuicksightDataSourceDataSourceParametersS3Parameters `field:"optional" json:"s3Parameters" yaml:"s3Parameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_data_source#s3_tables_parameters QuicksightDataSource#s3_tables_parameters}.
	S3TablesParameters *QuicksightDataSourceDataSourceParametersS3TablesParameters `field:"optional" json:"s3TablesParameters" yaml:"s3TablesParameters"`
	// <p>The parameters for Snowflake.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_data_source#snowflake_parameters QuicksightDataSource#snowflake_parameters}
	SnowflakeParameters *QuicksightDataSourceDataSourceParametersSnowflakeParameters `field:"optional" json:"snowflakeParameters" yaml:"snowflakeParameters"`
	// <p>The parameters for Spark.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_data_source#spark_parameters QuicksightDataSource#spark_parameters}
	SparkParameters *QuicksightDataSourceDataSourceParametersSparkParameters `field:"optional" json:"sparkParameters" yaml:"sparkParameters"`
	// <p>The parameters for SQL Server.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_data_source#sql_server_parameters QuicksightDataSource#sql_server_parameters}
	SqlServerParameters *QuicksightDataSourceDataSourceParametersSqlServerParameters `field:"optional" json:"sqlServerParameters" yaml:"sqlServerParameters"`
	// <p>The parameters that are required to connect to a Starburst data source.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_data_source#starburst_parameters QuicksightDataSource#starburst_parameters}
	StarburstParameters *QuicksightDataSourceDataSourceParametersStarburstParameters `field:"optional" json:"starburstParameters" yaml:"starburstParameters"`
	// <p>The parameters for Teradata.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_data_source#teradata_parameters QuicksightDataSource#teradata_parameters}
	TeradataParameters *QuicksightDataSourceDataSourceParametersTeradataParameters `field:"optional" json:"teradataParameters" yaml:"teradataParameters"`
	// <p>The parameters that are required to connect to a Trino data source.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_data_source#trino_parameters QuicksightDataSource#trino_parameters}
	TrinoParameters *QuicksightDataSourceDataSourceParametersTrinoParameters `field:"optional" json:"trinoParameters" yaml:"trinoParameters"`
}

