// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dmsdataprovider


type DmsDataProviderSettings struct {
	// DocDbSettings property identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dms_data_provider#doc_db_settings DmsDataProvider#doc_db_settings}
	DocDbSettings *DmsDataProviderSettingsDocDbSettings `field:"optional" json:"docDbSettings" yaml:"docDbSettings"`
	// IbmDb2LuwSettings property identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dms_data_provider#ibm_db_2_luw_settings DmsDataProvider#ibm_db_2_luw_settings}
	IbmDb2LuwSettings *DmsDataProviderSettingsIbmDb2LuwSettings `field:"optional" json:"ibmDb2LuwSettings" yaml:"ibmDb2LuwSettings"`
	// IbmDb2zOsSettings property identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dms_data_provider#ibm_db_2_z_os_settings DmsDataProvider#ibm_db_2_z_os_settings}
	IbmDb2ZOsSettings *DmsDataProviderSettingsIbmDb2ZOsSettings `field:"optional" json:"ibmDb2ZOsSettings" yaml:"ibmDb2ZOsSettings"`
	// MariaDbSettings property identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dms_data_provider#maria_db_settings DmsDataProvider#maria_db_settings}
	MariaDbSettings *DmsDataProviderSettingsMariaDbSettings `field:"optional" json:"mariaDbSettings" yaml:"mariaDbSettings"`
	// MicrosoftSqlServerSettings property identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dms_data_provider#microsoft_sql_server_settings DmsDataProvider#microsoft_sql_server_settings}
	MicrosoftSqlServerSettings *DmsDataProviderSettingsMicrosoftSqlServerSettings `field:"optional" json:"microsoftSqlServerSettings" yaml:"microsoftSqlServerSettings"`
	// MongoDbSettings property identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dms_data_provider#mongo_db_settings DmsDataProvider#mongo_db_settings}
	MongoDbSettings *DmsDataProviderSettingsMongoDbSettings `field:"optional" json:"mongoDbSettings" yaml:"mongoDbSettings"`
	// MySqlSettings property identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dms_data_provider#my_sql_settings DmsDataProvider#my_sql_settings}
	MySqlSettings *DmsDataProviderSettingsMySqlSettings `field:"optional" json:"mySqlSettings" yaml:"mySqlSettings"`
	// OracleSettings property identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dms_data_provider#oracle_settings DmsDataProvider#oracle_settings}
	OracleSettings *DmsDataProviderSettingsOracleSettings `field:"optional" json:"oracleSettings" yaml:"oracleSettings"`
	// PostgreSqlSettings property identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dms_data_provider#postgre_sql_settings DmsDataProvider#postgre_sql_settings}
	PostgreSqlSettings *DmsDataProviderSettingsPostgreSqlSettings `field:"optional" json:"postgreSqlSettings" yaml:"postgreSqlSettings"`
	// RedshiftSettings property identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dms_data_provider#redshift_settings DmsDataProvider#redshift_settings}
	RedshiftSettings *DmsDataProviderSettingsRedshiftSettings `field:"optional" json:"redshiftSettings" yaml:"redshiftSettings"`
	// SybaseAseSettings property identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dms_data_provider#sybase_ase_settings DmsDataProvider#sybase_ase_settings}
	SybaseAseSettings *DmsDataProviderSettingsSybaseAseSettings `field:"optional" json:"sybaseAseSettings" yaml:"sybaseAseSettings"`
}

