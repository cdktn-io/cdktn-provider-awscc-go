// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dmsendpoint

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DmsEndpointConfig struct {
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
	// The type of endpoint. Valid values are source and target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/dms_endpoint#endpoint_type DmsEndpoint#endpoint_type}
	EndpointType *string `field:"required" json:"endpointType" yaml:"endpointType"`
	// The type of engine for the endpoint, depending on the EndpointType value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/dms_endpoint#engine_name DmsEndpoint#engine_name}
	EngineName *string `field:"required" json:"engineName" yaml:"engineName"`
	// The Amazon Resource Name (ARN) for the certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/dms_endpoint#certificate_arn DmsEndpoint#certificate_arn}
	CertificateArn *string `field:"optional" json:"certificateArn" yaml:"certificateArn"`
	// The name of the endpoint database.
	//
	// For a MySQL source or target endpoint, don't specify DatabaseName. To migrate to a specific database, use this setting and targetDbType.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/dms_endpoint#database_name DmsEndpoint#database_name}
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// Settings in JSON format for the source and target DocumentDB endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/dms_endpoint#doc_db_settings DmsEndpoint#doc_db_settings}
	DocDbSettings *DmsEndpointDocDbSettings `field:"optional" json:"docDbSettings" yaml:"docDbSettings"`
	// Settings in JSON format for the target Amazon DynamoDB endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/dms_endpoint#dynamo_db_settings DmsEndpoint#dynamo_db_settings}
	DynamoDbSettings *DmsEndpointDynamoDbSettings `field:"optional" json:"dynamoDbSettings" yaml:"dynamoDbSettings"`
	// Settings in JSON format for the target OpenSearch endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/dms_endpoint#elasticsearch_settings DmsEndpoint#elasticsearch_settings}
	ElasticsearchSettings *DmsEndpointElasticsearchSettings `field:"optional" json:"elasticsearchSettings" yaml:"elasticsearchSettings"`
	// The database endpoint identifier.
	//
	// Identifiers must begin with a letter and must contain only ASCII letters, digits, and hyphens. They can't end with a hyphen, or contain two consecutive hyphens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/dms_endpoint#endpoint_identifier DmsEndpoint#endpoint_identifier}
	EndpointIdentifier *string `field:"optional" json:"endpointIdentifier" yaml:"endpointIdentifier"`
	// Additional attributes associated with the connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/dms_endpoint#extra_connection_attributes DmsEndpoint#extra_connection_attributes}
	ExtraConnectionAttributes *string `field:"optional" json:"extraConnectionAttributes" yaml:"extraConnectionAttributes"`
	// Settings in JSON format for the source GCP MySQL endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/dms_endpoint#gcp_my_sql_settings DmsEndpoint#gcp_my_sql_settings}
	GcpMySqlSettings *DmsEndpointGcpMySqlSettings `field:"optional" json:"gcpMySqlSettings" yaml:"gcpMySqlSettings"`
	// Settings in JSON format for the source IBM Db2 LUW endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/dms_endpoint#ibm_db_2_settings DmsEndpoint#ibm_db_2_settings}
	IbmDb2Settings *DmsEndpointIbmDb2Settings `field:"optional" json:"ibmDb2Settings" yaml:"ibmDb2Settings"`
	// Settings in JSON format for the target Apache Kafka endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/dms_endpoint#kafka_settings DmsEndpoint#kafka_settings}
	KafkaSettings *DmsEndpointKafkaSettings `field:"optional" json:"kafkaSettings" yaml:"kafkaSettings"`
	// Settings in JSON format for the target endpoint for Amazon Kinesis Data Streams.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/dms_endpoint#kinesis_settings DmsEndpoint#kinesis_settings}
	KinesisSettings *DmsEndpointKinesisSettings `field:"optional" json:"kinesisSettings" yaml:"kinesisSettings"`
	// An AWS KMS key identifier that is used to encrypt the connection parameters for the endpoint.If you don't specify a value for the KmsKeyId parameter, AWS DMS uses your default encryption key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/dms_endpoint#kms_key_id DmsEndpoint#kms_key_id}
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Settings in JSON format for the source and target Microsoft SQL Server endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/dms_endpoint#microsoft_sql_server_settings DmsEndpoint#microsoft_sql_server_settings}
	MicrosoftSqlServerSettings *DmsEndpointMicrosoftSqlServerSettings `field:"optional" json:"microsoftSqlServerSettings" yaml:"microsoftSqlServerSettings"`
	// Settings in JSON format for the source MongoDB endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/dms_endpoint#mongo_db_settings DmsEndpoint#mongo_db_settings}
	MongoDbSettings *DmsEndpointMongoDbSettings `field:"optional" json:"mongoDbSettings" yaml:"mongoDbSettings"`
	// Settings in JSON format for the source and target MySQL endpoin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/dms_endpoint#my_sql_settings DmsEndpoint#my_sql_settings}
	MySqlSettings *DmsEndpointMySqlSettings `field:"optional" json:"mySqlSettings" yaml:"mySqlSettings"`
	// Settings in JSON format for the target Amazon Neptune endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/dms_endpoint#neptune_settings DmsEndpoint#neptune_settings}
	NeptuneSettings *DmsEndpointNeptuneSettings `field:"optional" json:"neptuneSettings" yaml:"neptuneSettings"`
	// Settings in JSON format for the source and target Oracle endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/dms_endpoint#oracle_settings DmsEndpoint#oracle_settings}
	OracleSettings *DmsEndpointOracleSettings `field:"optional" json:"oracleSettings" yaml:"oracleSettings"`
	// The password to be used to log in to the endpoint database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/dms_endpoint#password DmsEndpoint#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// The port used by the endpoint database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/dms_endpoint#port DmsEndpoint#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Settings in JSON format for the source and target PostgreSQL endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/dms_endpoint#postgre_sql_settings DmsEndpoint#postgre_sql_settings}
	PostgreSqlSettings *DmsEndpointPostgreSqlSettings `field:"optional" json:"postgreSqlSettings" yaml:"postgreSqlSettings"`
	// Settings in JSON format for the target Redis endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/dms_endpoint#redis_settings DmsEndpoint#redis_settings}
	RedisSettings *DmsEndpointRedisSettings `field:"optional" json:"redisSettings" yaml:"redisSettings"`
	// Settings in JSON format for the Amazon Redshift endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/dms_endpoint#redshift_settings DmsEndpoint#redshift_settings}
	RedshiftSettings *DmsEndpointRedshiftSettings `field:"optional" json:"redshiftSettings" yaml:"redshiftSettings"`
	// A display name for the resource identifier at the end of the EndpointArn response parameter that is returned in the created Endpoint object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/dms_endpoint#resource_identifier DmsEndpoint#resource_identifier}
	ResourceIdentifier *string `field:"optional" json:"resourceIdentifier" yaml:"resourceIdentifier"`
	// Settings in JSON format for the source and target Amazon S3 endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/dms_endpoint#s3_settings DmsEndpoint#s3_settings}
	S3Settings *DmsEndpointS3Settings `field:"optional" json:"s3Settings" yaml:"s3Settings"`
	// The name of the server where the endpoint database resides.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/dms_endpoint#server_name DmsEndpoint#server_name}
	ServerName *string `field:"optional" json:"serverName" yaml:"serverName"`
	// The Secure Sockets Layer (SSL) mode to use for the SSL connection. The default is none.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/dms_endpoint#ssl_mode DmsEndpoint#ssl_mode}
	SslMode *string `field:"optional" json:"sslMode" yaml:"sslMode"`
	// Settings in JSON format for the source and target SAP ASE endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/dms_endpoint#sybase_settings DmsEndpoint#sybase_settings}
	SybaseSettings *DmsEndpointSybaseSettings `field:"optional" json:"sybaseSettings" yaml:"sybaseSettings"`
	// One or more tags to be assigned to the endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/dms_endpoint#tags DmsEndpoint#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The user name to be used to log in to the endpoint database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/dms_endpoint#username DmsEndpoint#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
}

