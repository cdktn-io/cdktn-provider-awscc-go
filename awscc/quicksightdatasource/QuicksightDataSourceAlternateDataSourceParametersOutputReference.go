// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightdatasource

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/quicksightdatasource/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type QuicksightDataSourceAlternateDataSourceParametersOutputReference interface {
	cdktn.ComplexObject
	AmazonElasticsearchParameters() QuicksightDataSourceAlternateDataSourceParametersAmazonElasticsearchParametersOutputReference
	AmazonElasticsearchParametersInput() interface{}
	AmazonOpenSearchParameters() QuicksightDataSourceAlternateDataSourceParametersAmazonOpenSearchParametersOutputReference
	AmazonOpenSearchParametersInput() interface{}
	AthenaParameters() QuicksightDataSourceAlternateDataSourceParametersAthenaParametersOutputReference
	AthenaParametersInput() interface{}
	AuroraParameters() QuicksightDataSourceAlternateDataSourceParametersAuroraParametersOutputReference
	AuroraParametersInput() interface{}
	AuroraPostgreSqlParameters() QuicksightDataSourceAlternateDataSourceParametersAuroraPostgreSqlParametersOutputReference
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
	DatabricksParameters() QuicksightDataSourceAlternateDataSourceParametersDatabricksParametersOutputReference
	DatabricksParametersInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MariaDbParameters() QuicksightDataSourceAlternateDataSourceParametersMariaDbParametersOutputReference
	MariaDbParametersInput() interface{}
	MySqlParameters() QuicksightDataSourceAlternateDataSourceParametersMySqlParametersOutputReference
	MySqlParametersInput() interface{}
	OracleParameters() QuicksightDataSourceAlternateDataSourceParametersOracleParametersOutputReference
	OracleParametersInput() interface{}
	PostgreSqlParameters() QuicksightDataSourceAlternateDataSourceParametersPostgreSqlParametersOutputReference
	PostgreSqlParametersInput() interface{}
	PrestoParameters() QuicksightDataSourceAlternateDataSourceParametersPrestoParametersOutputReference
	PrestoParametersInput() interface{}
	RdsParameters() QuicksightDataSourceAlternateDataSourceParametersRdsParametersOutputReference
	RdsParametersInput() interface{}
	RedshiftParameters() QuicksightDataSourceAlternateDataSourceParametersRedshiftParametersOutputReference
	RedshiftParametersInput() interface{}
	S3Parameters() QuicksightDataSourceAlternateDataSourceParametersS3ParametersOutputReference
	S3ParametersInput() interface{}
	S3TablesParameters() QuicksightDataSourceAlternateDataSourceParametersS3TablesParametersOutputReference
	S3TablesParametersInput() interface{}
	SnowflakeParameters() QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOutputReference
	SnowflakeParametersInput() interface{}
	SparkParameters() QuicksightDataSourceAlternateDataSourceParametersSparkParametersOutputReference
	SparkParametersInput() interface{}
	SqlServerParameters() QuicksightDataSourceAlternateDataSourceParametersSqlServerParametersOutputReference
	SqlServerParametersInput() interface{}
	StarburstParameters() QuicksightDataSourceAlternateDataSourceParametersStarburstParametersOutputReference
	StarburstParametersInput() interface{}
	TeradataParameters() QuicksightDataSourceAlternateDataSourceParametersTeradataParametersOutputReference
	TeradataParametersInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TrinoParameters() QuicksightDataSourceAlternateDataSourceParametersTrinoParametersOutputReference
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
	PutAmazonElasticsearchParameters(value *QuicksightDataSourceAlternateDataSourceParametersAmazonElasticsearchParameters)
	PutAmazonOpenSearchParameters(value *QuicksightDataSourceAlternateDataSourceParametersAmazonOpenSearchParameters)
	PutAthenaParameters(value *QuicksightDataSourceAlternateDataSourceParametersAthenaParameters)
	PutAuroraParameters(value *QuicksightDataSourceAlternateDataSourceParametersAuroraParameters)
	PutAuroraPostgreSqlParameters(value *QuicksightDataSourceAlternateDataSourceParametersAuroraPostgreSqlParameters)
	PutDatabricksParameters(value *QuicksightDataSourceAlternateDataSourceParametersDatabricksParameters)
	PutMariaDbParameters(value *QuicksightDataSourceAlternateDataSourceParametersMariaDbParameters)
	PutMySqlParameters(value *QuicksightDataSourceAlternateDataSourceParametersMySqlParameters)
	PutOracleParameters(value *QuicksightDataSourceAlternateDataSourceParametersOracleParameters)
	PutPostgreSqlParameters(value *QuicksightDataSourceAlternateDataSourceParametersPostgreSqlParameters)
	PutPrestoParameters(value *QuicksightDataSourceAlternateDataSourceParametersPrestoParameters)
	PutRdsParameters(value *QuicksightDataSourceAlternateDataSourceParametersRdsParameters)
	PutRedshiftParameters(value *QuicksightDataSourceAlternateDataSourceParametersRedshiftParameters)
	PutS3Parameters(value *QuicksightDataSourceAlternateDataSourceParametersS3Parameters)
	PutS3TablesParameters(value *QuicksightDataSourceAlternateDataSourceParametersS3TablesParameters)
	PutSnowflakeParameters(value *QuicksightDataSourceAlternateDataSourceParametersSnowflakeParameters)
	PutSparkParameters(value *QuicksightDataSourceAlternateDataSourceParametersSparkParameters)
	PutSqlServerParameters(value *QuicksightDataSourceAlternateDataSourceParametersSqlServerParameters)
	PutStarburstParameters(value *QuicksightDataSourceAlternateDataSourceParametersStarburstParameters)
	PutTeradataParameters(value *QuicksightDataSourceAlternateDataSourceParametersTeradataParameters)
	PutTrinoParameters(value *QuicksightDataSourceAlternateDataSourceParametersTrinoParameters)
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

// The jsii proxy struct for QuicksightDataSourceAlternateDataSourceParametersOutputReference
type jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) AmazonElasticsearchParameters() QuicksightDataSourceAlternateDataSourceParametersAmazonElasticsearchParametersOutputReference {
	var returns QuicksightDataSourceAlternateDataSourceParametersAmazonElasticsearchParametersOutputReference
	_jsii_.Get(
		j,
		"amazonElasticsearchParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) AmazonElasticsearchParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"amazonElasticsearchParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) AmazonOpenSearchParameters() QuicksightDataSourceAlternateDataSourceParametersAmazonOpenSearchParametersOutputReference {
	var returns QuicksightDataSourceAlternateDataSourceParametersAmazonOpenSearchParametersOutputReference
	_jsii_.Get(
		j,
		"amazonOpenSearchParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) AmazonOpenSearchParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"amazonOpenSearchParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) AthenaParameters() QuicksightDataSourceAlternateDataSourceParametersAthenaParametersOutputReference {
	var returns QuicksightDataSourceAlternateDataSourceParametersAthenaParametersOutputReference
	_jsii_.Get(
		j,
		"athenaParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) AthenaParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"athenaParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) AuroraParameters() QuicksightDataSourceAlternateDataSourceParametersAuroraParametersOutputReference {
	var returns QuicksightDataSourceAlternateDataSourceParametersAuroraParametersOutputReference
	_jsii_.Get(
		j,
		"auroraParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) AuroraParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"auroraParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) AuroraPostgreSqlParameters() QuicksightDataSourceAlternateDataSourceParametersAuroraPostgreSqlParametersOutputReference {
	var returns QuicksightDataSourceAlternateDataSourceParametersAuroraPostgreSqlParametersOutputReference
	_jsii_.Get(
		j,
		"auroraPostgreSqlParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) AuroraPostgreSqlParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"auroraPostgreSqlParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) DatabricksParameters() QuicksightDataSourceAlternateDataSourceParametersDatabricksParametersOutputReference {
	var returns QuicksightDataSourceAlternateDataSourceParametersDatabricksParametersOutputReference
	_jsii_.Get(
		j,
		"databricksParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) DatabricksParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"databricksParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) MariaDbParameters() QuicksightDataSourceAlternateDataSourceParametersMariaDbParametersOutputReference {
	var returns QuicksightDataSourceAlternateDataSourceParametersMariaDbParametersOutputReference
	_jsii_.Get(
		j,
		"mariaDbParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) MariaDbParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mariaDbParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) MySqlParameters() QuicksightDataSourceAlternateDataSourceParametersMySqlParametersOutputReference {
	var returns QuicksightDataSourceAlternateDataSourceParametersMySqlParametersOutputReference
	_jsii_.Get(
		j,
		"mySqlParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) MySqlParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mySqlParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) OracleParameters() QuicksightDataSourceAlternateDataSourceParametersOracleParametersOutputReference {
	var returns QuicksightDataSourceAlternateDataSourceParametersOracleParametersOutputReference
	_jsii_.Get(
		j,
		"oracleParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) OracleParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"oracleParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) PostgreSqlParameters() QuicksightDataSourceAlternateDataSourceParametersPostgreSqlParametersOutputReference {
	var returns QuicksightDataSourceAlternateDataSourceParametersPostgreSqlParametersOutputReference
	_jsii_.Get(
		j,
		"postgreSqlParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) PostgreSqlParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"postgreSqlParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) PrestoParameters() QuicksightDataSourceAlternateDataSourceParametersPrestoParametersOutputReference {
	var returns QuicksightDataSourceAlternateDataSourceParametersPrestoParametersOutputReference
	_jsii_.Get(
		j,
		"prestoParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) PrestoParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"prestoParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) RdsParameters() QuicksightDataSourceAlternateDataSourceParametersRdsParametersOutputReference {
	var returns QuicksightDataSourceAlternateDataSourceParametersRdsParametersOutputReference
	_jsii_.Get(
		j,
		"rdsParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) RdsParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rdsParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) RedshiftParameters() QuicksightDataSourceAlternateDataSourceParametersRedshiftParametersOutputReference {
	var returns QuicksightDataSourceAlternateDataSourceParametersRedshiftParametersOutputReference
	_jsii_.Get(
		j,
		"redshiftParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) RedshiftParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"redshiftParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) S3Parameters() QuicksightDataSourceAlternateDataSourceParametersS3ParametersOutputReference {
	var returns QuicksightDataSourceAlternateDataSourceParametersS3ParametersOutputReference
	_jsii_.Get(
		j,
		"s3Parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) S3ParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3ParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) S3TablesParameters() QuicksightDataSourceAlternateDataSourceParametersS3TablesParametersOutputReference {
	var returns QuicksightDataSourceAlternateDataSourceParametersS3TablesParametersOutputReference
	_jsii_.Get(
		j,
		"s3TablesParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) S3TablesParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3TablesParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) SnowflakeParameters() QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOutputReference {
	var returns QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOutputReference
	_jsii_.Get(
		j,
		"snowflakeParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) SnowflakeParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"snowflakeParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) SparkParameters() QuicksightDataSourceAlternateDataSourceParametersSparkParametersOutputReference {
	var returns QuicksightDataSourceAlternateDataSourceParametersSparkParametersOutputReference
	_jsii_.Get(
		j,
		"sparkParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) SparkParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sparkParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) SqlServerParameters() QuicksightDataSourceAlternateDataSourceParametersSqlServerParametersOutputReference {
	var returns QuicksightDataSourceAlternateDataSourceParametersSqlServerParametersOutputReference
	_jsii_.Get(
		j,
		"sqlServerParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) SqlServerParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sqlServerParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) StarburstParameters() QuicksightDataSourceAlternateDataSourceParametersStarburstParametersOutputReference {
	var returns QuicksightDataSourceAlternateDataSourceParametersStarburstParametersOutputReference
	_jsii_.Get(
		j,
		"starburstParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) StarburstParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"starburstParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) TeradataParameters() QuicksightDataSourceAlternateDataSourceParametersTeradataParametersOutputReference {
	var returns QuicksightDataSourceAlternateDataSourceParametersTeradataParametersOutputReference
	_jsii_.Get(
		j,
		"teradataParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) TeradataParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"teradataParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) TrinoParameters() QuicksightDataSourceAlternateDataSourceParametersTrinoParametersOutputReference {
	var returns QuicksightDataSourceAlternateDataSourceParametersTrinoParametersOutputReference
	_jsii_.Get(
		j,
		"trinoParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) TrinoParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trinoParametersInput",
		&returns,
	)
	return returns
}


func NewQuicksightDataSourceAlternateDataSourceParametersOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) QuicksightDataSourceAlternateDataSourceParametersOutputReference {
	_init_.Initialize()

	if err := validateNewQuicksightDataSourceAlternateDataSourceParametersOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.quicksightDataSource.QuicksightDataSourceAlternateDataSourceParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewQuicksightDataSourceAlternateDataSourceParametersOutputReference_Override(q QuicksightDataSourceAlternateDataSourceParametersOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.quicksightDataSource.QuicksightDataSourceAlternateDataSourceParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		q,
	)
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) PutAmazonElasticsearchParameters(value *QuicksightDataSourceAlternateDataSourceParametersAmazonElasticsearchParameters) {
	if err := q.validatePutAmazonElasticsearchParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putAmazonElasticsearchParameters",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) PutAmazonOpenSearchParameters(value *QuicksightDataSourceAlternateDataSourceParametersAmazonOpenSearchParameters) {
	if err := q.validatePutAmazonOpenSearchParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putAmazonOpenSearchParameters",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) PutAthenaParameters(value *QuicksightDataSourceAlternateDataSourceParametersAthenaParameters) {
	if err := q.validatePutAthenaParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putAthenaParameters",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) PutAuroraParameters(value *QuicksightDataSourceAlternateDataSourceParametersAuroraParameters) {
	if err := q.validatePutAuroraParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putAuroraParameters",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) PutAuroraPostgreSqlParameters(value *QuicksightDataSourceAlternateDataSourceParametersAuroraPostgreSqlParameters) {
	if err := q.validatePutAuroraPostgreSqlParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putAuroraPostgreSqlParameters",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) PutDatabricksParameters(value *QuicksightDataSourceAlternateDataSourceParametersDatabricksParameters) {
	if err := q.validatePutDatabricksParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putDatabricksParameters",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) PutMariaDbParameters(value *QuicksightDataSourceAlternateDataSourceParametersMariaDbParameters) {
	if err := q.validatePutMariaDbParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putMariaDbParameters",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) PutMySqlParameters(value *QuicksightDataSourceAlternateDataSourceParametersMySqlParameters) {
	if err := q.validatePutMySqlParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putMySqlParameters",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) PutOracleParameters(value *QuicksightDataSourceAlternateDataSourceParametersOracleParameters) {
	if err := q.validatePutOracleParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putOracleParameters",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) PutPostgreSqlParameters(value *QuicksightDataSourceAlternateDataSourceParametersPostgreSqlParameters) {
	if err := q.validatePutPostgreSqlParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putPostgreSqlParameters",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) PutPrestoParameters(value *QuicksightDataSourceAlternateDataSourceParametersPrestoParameters) {
	if err := q.validatePutPrestoParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putPrestoParameters",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) PutRdsParameters(value *QuicksightDataSourceAlternateDataSourceParametersRdsParameters) {
	if err := q.validatePutRdsParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putRdsParameters",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) PutRedshiftParameters(value *QuicksightDataSourceAlternateDataSourceParametersRedshiftParameters) {
	if err := q.validatePutRedshiftParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putRedshiftParameters",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) PutS3Parameters(value *QuicksightDataSourceAlternateDataSourceParametersS3Parameters) {
	if err := q.validatePutS3ParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putS3Parameters",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) PutS3TablesParameters(value *QuicksightDataSourceAlternateDataSourceParametersS3TablesParameters) {
	if err := q.validatePutS3TablesParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putS3TablesParameters",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) PutSnowflakeParameters(value *QuicksightDataSourceAlternateDataSourceParametersSnowflakeParameters) {
	if err := q.validatePutSnowflakeParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putSnowflakeParameters",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) PutSparkParameters(value *QuicksightDataSourceAlternateDataSourceParametersSparkParameters) {
	if err := q.validatePutSparkParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putSparkParameters",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) PutSqlServerParameters(value *QuicksightDataSourceAlternateDataSourceParametersSqlServerParameters) {
	if err := q.validatePutSqlServerParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putSqlServerParameters",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) PutStarburstParameters(value *QuicksightDataSourceAlternateDataSourceParametersStarburstParameters) {
	if err := q.validatePutStarburstParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putStarburstParameters",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) PutTeradataParameters(value *QuicksightDataSourceAlternateDataSourceParametersTeradataParameters) {
	if err := q.validatePutTeradataParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putTeradataParameters",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) PutTrinoParameters(value *QuicksightDataSourceAlternateDataSourceParametersTrinoParameters) {
	if err := q.validatePutTrinoParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putTrinoParameters",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) ResetAmazonElasticsearchParameters() {
	_jsii_.InvokeVoid(
		q,
		"resetAmazonElasticsearchParameters",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) ResetAmazonOpenSearchParameters() {
	_jsii_.InvokeVoid(
		q,
		"resetAmazonOpenSearchParameters",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) ResetAthenaParameters() {
	_jsii_.InvokeVoid(
		q,
		"resetAthenaParameters",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) ResetAuroraParameters() {
	_jsii_.InvokeVoid(
		q,
		"resetAuroraParameters",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) ResetAuroraPostgreSqlParameters() {
	_jsii_.InvokeVoid(
		q,
		"resetAuroraPostgreSqlParameters",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) ResetDatabricksParameters() {
	_jsii_.InvokeVoid(
		q,
		"resetDatabricksParameters",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) ResetMariaDbParameters() {
	_jsii_.InvokeVoid(
		q,
		"resetMariaDbParameters",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) ResetMySqlParameters() {
	_jsii_.InvokeVoid(
		q,
		"resetMySqlParameters",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) ResetOracleParameters() {
	_jsii_.InvokeVoid(
		q,
		"resetOracleParameters",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) ResetPostgreSqlParameters() {
	_jsii_.InvokeVoid(
		q,
		"resetPostgreSqlParameters",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) ResetPrestoParameters() {
	_jsii_.InvokeVoid(
		q,
		"resetPrestoParameters",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) ResetRdsParameters() {
	_jsii_.InvokeVoid(
		q,
		"resetRdsParameters",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) ResetRedshiftParameters() {
	_jsii_.InvokeVoid(
		q,
		"resetRedshiftParameters",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) ResetS3Parameters() {
	_jsii_.InvokeVoid(
		q,
		"resetS3Parameters",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) ResetS3TablesParameters() {
	_jsii_.InvokeVoid(
		q,
		"resetS3TablesParameters",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) ResetSnowflakeParameters() {
	_jsii_.InvokeVoid(
		q,
		"resetSnowflakeParameters",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) ResetSparkParameters() {
	_jsii_.InvokeVoid(
		q,
		"resetSparkParameters",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) ResetSqlServerParameters() {
	_jsii_.InvokeVoid(
		q,
		"resetSqlServerParameters",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) ResetStarburstParameters() {
	_jsii_.InvokeVoid(
		q,
		"resetStarburstParameters",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) ResetTeradataParameters() {
	_jsii_.InvokeVoid(
		q,
		"resetTeradataParameters",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) ResetTrinoParameters() {
	_jsii_.InvokeVoid(
		q,
		"resetTrinoParameters",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

