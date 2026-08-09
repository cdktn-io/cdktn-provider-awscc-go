// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightdatasource

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/quicksightdatasource/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference interface {
	cdktn.ComplexObject
	AmazonElasticsearchParameters() QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersAmazonElasticsearchParametersOutputReference
	AmazonElasticsearchParametersInput() interface{}
	AmazonOpenSearchParameters() QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersAmazonOpenSearchParametersOutputReference
	AmazonOpenSearchParametersInput() interface{}
	AthenaParameters() QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersAthenaParametersOutputReference
	AthenaParametersInput() interface{}
	AuroraParameters() QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersAuroraParametersOutputReference
	AuroraParametersInput() interface{}
	AuroraPostgreSqlParameters() QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersAuroraPostgreSqlParametersOutputReference
	AuroraPostgreSqlParametersInput() interface{}
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DatabricksParameters() QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersDatabricksParametersOutputReference
	DatabricksParametersInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MariaDbParameters() QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersMariaDbParametersOutputReference
	MariaDbParametersInput() interface{}
	MySqlParameters() QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersMySqlParametersOutputReference
	MySqlParametersInput() interface{}
	OracleParameters() QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOracleParametersOutputReference
	OracleParametersInput() interface{}
	PostgreSqlParameters() QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersPostgreSqlParametersOutputReference
	PostgreSqlParametersInput() interface{}
	PrestoParameters() QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersPrestoParametersOutputReference
	PrestoParametersInput() interface{}
	RdsParameters() QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRdsParametersOutputReference
	RdsParametersInput() interface{}
	RedshiftParameters() QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersOutputReference
	RedshiftParametersInput() interface{}
	S3Parameters() QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersS3ParametersOutputReference
	S3ParametersInput() interface{}
	S3TablesParameters() QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersS3TablesParametersOutputReference
	S3TablesParametersInput() interface{}
	SnowflakeParameters() QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersSnowflakeParametersOutputReference
	SnowflakeParametersInput() interface{}
	SparkParameters() QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersSparkParametersOutputReference
	SparkParametersInput() interface{}
	SqlServerParameters() QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersSqlServerParametersOutputReference
	SqlServerParametersInput() interface{}
	StarburstParameters() QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOutputReference
	StarburstParametersInput() interface{}
	TeradataParameters() QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersTeradataParametersOutputReference
	TeradataParametersInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TrinoParameters() QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersTrinoParametersOutputReference
	TrinoParametersInput() interface{}
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	PutAmazonElasticsearchParameters(value *QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersAmazonElasticsearchParameters)
	PutAmazonOpenSearchParameters(value *QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersAmazonOpenSearchParameters)
	PutAthenaParameters(value *QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersAthenaParameters)
	PutAuroraParameters(value *QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersAuroraParameters)
	PutAuroraPostgreSqlParameters(value *QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersAuroraPostgreSqlParameters)
	PutDatabricksParameters(value *QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersDatabricksParameters)
	PutMariaDbParameters(value *QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersMariaDbParameters)
	PutMySqlParameters(value *QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersMySqlParameters)
	PutOracleParameters(value *QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOracleParameters)
	PutPostgreSqlParameters(value *QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersPostgreSqlParameters)
	PutPrestoParameters(value *QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersPrestoParameters)
	PutRdsParameters(value *QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRdsParameters)
	PutRedshiftParameters(value *QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParameters)
	PutS3Parameters(value *QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersS3Parameters)
	PutS3TablesParameters(value *QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersS3TablesParameters)
	PutSnowflakeParameters(value *QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersSnowflakeParameters)
	PutSparkParameters(value *QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersSparkParameters)
	PutSqlServerParameters(value *QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersSqlServerParameters)
	PutStarburstParameters(value *QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParameters)
	PutTeradataParameters(value *QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersTeradataParameters)
	PutTrinoParameters(value *QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersTrinoParameters)
	ResetAmazonElasticsearchParameters()
	ResetAmazonOpenSearchParameters()
	ResetAthenaParameters()
	ResetAuroraParameters()
	ResetAuroraPostgreSqlParameters()
	ResetDatabricksParameters()
	ResetMariaDbParameters()
	ResetMySqlParameters()
	ResetOracleParameters()
	ResetPostgreSqlParameters()
	ResetPrestoParameters()
	ResetRdsParameters()
	ResetRedshiftParameters()
	ResetS3Parameters()
	ResetS3TablesParameters()
	ResetSnowflakeParameters()
	ResetSparkParameters()
	ResetSqlServerParameters()
	ResetStarburstParameters()
	ResetTeradataParameters()
	ResetTrinoParameters()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference
type jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) AmazonElasticsearchParameters() QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersAmazonElasticsearchParametersOutputReference {
	var returns QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersAmazonElasticsearchParametersOutputReference
	_jsii_.Get(
		j,
		"amazonElasticsearchParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) AmazonElasticsearchParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"amazonElasticsearchParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) AmazonOpenSearchParameters() QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersAmazonOpenSearchParametersOutputReference {
	var returns QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersAmazonOpenSearchParametersOutputReference
	_jsii_.Get(
		j,
		"amazonOpenSearchParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) AmazonOpenSearchParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"amazonOpenSearchParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) AthenaParameters() QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersAthenaParametersOutputReference {
	var returns QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersAthenaParametersOutputReference
	_jsii_.Get(
		j,
		"athenaParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) AthenaParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"athenaParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) AuroraParameters() QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersAuroraParametersOutputReference {
	var returns QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersAuroraParametersOutputReference
	_jsii_.Get(
		j,
		"auroraParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) AuroraParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"auroraParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) AuroraPostgreSqlParameters() QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersAuroraPostgreSqlParametersOutputReference {
	var returns QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersAuroraPostgreSqlParametersOutputReference
	_jsii_.Get(
		j,
		"auroraPostgreSqlParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) AuroraPostgreSqlParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"auroraPostgreSqlParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) DatabricksParameters() QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersDatabricksParametersOutputReference {
	var returns QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersDatabricksParametersOutputReference
	_jsii_.Get(
		j,
		"databricksParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) DatabricksParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"databricksParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) MariaDbParameters() QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersMariaDbParametersOutputReference {
	var returns QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersMariaDbParametersOutputReference
	_jsii_.Get(
		j,
		"mariaDbParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) MariaDbParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mariaDbParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) MySqlParameters() QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersMySqlParametersOutputReference {
	var returns QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersMySqlParametersOutputReference
	_jsii_.Get(
		j,
		"mySqlParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) MySqlParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mySqlParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) OracleParameters() QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOracleParametersOutputReference {
	var returns QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOracleParametersOutputReference
	_jsii_.Get(
		j,
		"oracleParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) OracleParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"oracleParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) PostgreSqlParameters() QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersPostgreSqlParametersOutputReference {
	var returns QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersPostgreSqlParametersOutputReference
	_jsii_.Get(
		j,
		"postgreSqlParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) PostgreSqlParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"postgreSqlParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) PrestoParameters() QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersPrestoParametersOutputReference {
	var returns QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersPrestoParametersOutputReference
	_jsii_.Get(
		j,
		"prestoParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) PrestoParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"prestoParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) RdsParameters() QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRdsParametersOutputReference {
	var returns QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRdsParametersOutputReference
	_jsii_.Get(
		j,
		"rdsParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) RdsParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rdsParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) RedshiftParameters() QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersOutputReference {
	var returns QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersOutputReference
	_jsii_.Get(
		j,
		"redshiftParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) RedshiftParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"redshiftParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) S3Parameters() QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersS3ParametersOutputReference {
	var returns QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersS3ParametersOutputReference
	_jsii_.Get(
		j,
		"s3Parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) S3ParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3ParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) S3TablesParameters() QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersS3TablesParametersOutputReference {
	var returns QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersS3TablesParametersOutputReference
	_jsii_.Get(
		j,
		"s3TablesParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) S3TablesParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3TablesParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) SnowflakeParameters() QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersSnowflakeParametersOutputReference {
	var returns QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersSnowflakeParametersOutputReference
	_jsii_.Get(
		j,
		"snowflakeParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) SnowflakeParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"snowflakeParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) SparkParameters() QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersSparkParametersOutputReference {
	var returns QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersSparkParametersOutputReference
	_jsii_.Get(
		j,
		"sparkParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) SparkParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sparkParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) SqlServerParameters() QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersSqlServerParametersOutputReference {
	var returns QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersSqlServerParametersOutputReference
	_jsii_.Get(
		j,
		"sqlServerParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) SqlServerParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sqlServerParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) StarburstParameters() QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOutputReference {
	var returns QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOutputReference
	_jsii_.Get(
		j,
		"starburstParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) StarburstParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"starburstParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) TeradataParameters() QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersTeradataParametersOutputReference {
	var returns QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersTeradataParametersOutputReference
	_jsii_.Get(
		j,
		"teradataParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) TeradataParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"teradataParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) TrinoParameters() QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersTrinoParametersOutputReference {
	var returns QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersTrinoParametersOutputReference
	_jsii_.Get(
		j,
		"trinoParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) TrinoParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trinoParametersInput",
		&returns,
	)
	return returns
}


func NewQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference {
	_init_.Initialize()

	if err := validateNewQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.quicksightDataSource.QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference_Override(q QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.quicksightDataSource.QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		q,
	)
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := q.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		q,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := q.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		q,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := q.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		q,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := q.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		q,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := q.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		q,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := q.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		q,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := q.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		q,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := q.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		q,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := q.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		q,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := q.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) PutAmazonElasticsearchParameters(value *QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersAmazonElasticsearchParameters) {
	if err := q.validatePutAmazonElasticsearchParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putAmazonElasticsearchParameters",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) PutAmazonOpenSearchParameters(value *QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersAmazonOpenSearchParameters) {
	if err := q.validatePutAmazonOpenSearchParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putAmazonOpenSearchParameters",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) PutAthenaParameters(value *QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersAthenaParameters) {
	if err := q.validatePutAthenaParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putAthenaParameters",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) PutAuroraParameters(value *QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersAuroraParameters) {
	if err := q.validatePutAuroraParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putAuroraParameters",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) PutAuroraPostgreSqlParameters(value *QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersAuroraPostgreSqlParameters) {
	if err := q.validatePutAuroraPostgreSqlParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putAuroraPostgreSqlParameters",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) PutDatabricksParameters(value *QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersDatabricksParameters) {
	if err := q.validatePutDatabricksParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putDatabricksParameters",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) PutMariaDbParameters(value *QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersMariaDbParameters) {
	if err := q.validatePutMariaDbParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putMariaDbParameters",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) PutMySqlParameters(value *QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersMySqlParameters) {
	if err := q.validatePutMySqlParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putMySqlParameters",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) PutOracleParameters(value *QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOracleParameters) {
	if err := q.validatePutOracleParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putOracleParameters",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) PutPostgreSqlParameters(value *QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersPostgreSqlParameters) {
	if err := q.validatePutPostgreSqlParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putPostgreSqlParameters",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) PutPrestoParameters(value *QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersPrestoParameters) {
	if err := q.validatePutPrestoParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putPrestoParameters",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) PutRdsParameters(value *QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRdsParameters) {
	if err := q.validatePutRdsParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putRdsParameters",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) PutRedshiftParameters(value *QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParameters) {
	if err := q.validatePutRedshiftParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putRedshiftParameters",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) PutS3Parameters(value *QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersS3Parameters) {
	if err := q.validatePutS3ParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putS3Parameters",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) PutS3TablesParameters(value *QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersS3TablesParameters) {
	if err := q.validatePutS3TablesParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putS3TablesParameters",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) PutSnowflakeParameters(value *QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersSnowflakeParameters) {
	if err := q.validatePutSnowflakeParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putSnowflakeParameters",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) PutSparkParameters(value *QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersSparkParameters) {
	if err := q.validatePutSparkParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putSparkParameters",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) PutSqlServerParameters(value *QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersSqlServerParameters) {
	if err := q.validatePutSqlServerParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putSqlServerParameters",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) PutStarburstParameters(value *QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParameters) {
	if err := q.validatePutStarburstParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putStarburstParameters",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) PutTeradataParameters(value *QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersTeradataParameters) {
	if err := q.validatePutTeradataParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putTeradataParameters",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) PutTrinoParameters(value *QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersTrinoParameters) {
	if err := q.validatePutTrinoParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putTrinoParameters",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) ResetAmazonElasticsearchParameters() {
	_jsii_.InvokeVoid(
		q,
		"resetAmazonElasticsearchParameters",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) ResetAmazonOpenSearchParameters() {
	_jsii_.InvokeVoid(
		q,
		"resetAmazonOpenSearchParameters",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) ResetAthenaParameters() {
	_jsii_.InvokeVoid(
		q,
		"resetAthenaParameters",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) ResetAuroraParameters() {
	_jsii_.InvokeVoid(
		q,
		"resetAuroraParameters",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) ResetAuroraPostgreSqlParameters() {
	_jsii_.InvokeVoid(
		q,
		"resetAuroraPostgreSqlParameters",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) ResetDatabricksParameters() {
	_jsii_.InvokeVoid(
		q,
		"resetDatabricksParameters",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) ResetMariaDbParameters() {
	_jsii_.InvokeVoid(
		q,
		"resetMariaDbParameters",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) ResetMySqlParameters() {
	_jsii_.InvokeVoid(
		q,
		"resetMySqlParameters",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) ResetOracleParameters() {
	_jsii_.InvokeVoid(
		q,
		"resetOracleParameters",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) ResetPostgreSqlParameters() {
	_jsii_.InvokeVoid(
		q,
		"resetPostgreSqlParameters",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) ResetPrestoParameters() {
	_jsii_.InvokeVoid(
		q,
		"resetPrestoParameters",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) ResetRdsParameters() {
	_jsii_.InvokeVoid(
		q,
		"resetRdsParameters",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) ResetRedshiftParameters() {
	_jsii_.InvokeVoid(
		q,
		"resetRedshiftParameters",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) ResetS3Parameters() {
	_jsii_.InvokeVoid(
		q,
		"resetS3Parameters",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) ResetS3TablesParameters() {
	_jsii_.InvokeVoid(
		q,
		"resetS3TablesParameters",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) ResetSnowflakeParameters() {
	_jsii_.InvokeVoid(
		q,
		"resetSnowflakeParameters",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) ResetSparkParameters() {
	_jsii_.InvokeVoid(
		q,
		"resetSparkParameters",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) ResetSqlServerParameters() {
	_jsii_.InvokeVoid(
		q,
		"resetSqlServerParameters",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) ResetStarburstParameters() {
	_jsii_.InvokeVoid(
		q,
		"resetStarburstParameters",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) ResetTeradataParameters() {
	_jsii_.InvokeVoid(
		q,
		"resetTeradataParameters",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) ResetTrinoParameters() {
	_jsii_.InvokeVoid(
		q,
		"resetTrinoParameters",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := q.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		q,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

