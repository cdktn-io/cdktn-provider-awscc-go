// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightdatasource

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/quicksightdatasource/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type QuicksightDataSourceDataSourceParametersOutputReference interface {
	cdktn.ComplexObject
	AmazonElasticsearchParameters() QuicksightDataSourceDataSourceParametersAmazonElasticsearchParametersOutputReference
	AmazonElasticsearchParametersInput() interface{}
	AmazonOpenSearchParameters() QuicksightDataSourceDataSourceParametersAmazonOpenSearchParametersOutputReference
	AmazonOpenSearchParametersInput() interface{}
	AthenaParameters() QuicksightDataSourceDataSourceParametersAthenaParametersOutputReference
	AthenaParametersInput() interface{}
	AuroraParameters() QuicksightDataSourceDataSourceParametersAuroraParametersOutputReference
	AuroraParametersInput() interface{}
	AuroraPostgreSqlParameters() QuicksightDataSourceDataSourceParametersAuroraPostgreSqlParametersOutputReference
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
	DatabricksParameters() QuicksightDataSourceDataSourceParametersDatabricksParametersOutputReference
	DatabricksParametersInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MariaDbParameters() QuicksightDataSourceDataSourceParametersMariaDbParametersOutputReference
	MariaDbParametersInput() interface{}
	MySqlParameters() QuicksightDataSourceDataSourceParametersMySqlParametersOutputReference
	MySqlParametersInput() interface{}
	OracleParameters() QuicksightDataSourceDataSourceParametersOracleParametersOutputReference
	OracleParametersInput() interface{}
	PostgreSqlParameters() QuicksightDataSourceDataSourceParametersPostgreSqlParametersOutputReference
	PostgreSqlParametersInput() interface{}
	PrestoParameters() QuicksightDataSourceDataSourceParametersPrestoParametersOutputReference
	PrestoParametersInput() interface{}
	RdsParameters() QuicksightDataSourceDataSourceParametersRdsParametersOutputReference
	RdsParametersInput() interface{}
	RedshiftParameters() QuicksightDataSourceDataSourceParametersRedshiftParametersOutputReference
	RedshiftParametersInput() interface{}
	S3Parameters() QuicksightDataSourceDataSourceParametersS3ParametersOutputReference
	S3ParametersInput() interface{}
	S3TablesParameters() QuicksightDataSourceDataSourceParametersS3TablesParametersOutputReference
	S3TablesParametersInput() interface{}
	SnowflakeParameters() QuicksightDataSourceDataSourceParametersSnowflakeParametersOutputReference
	SnowflakeParametersInput() interface{}
	SparkParameters() QuicksightDataSourceDataSourceParametersSparkParametersOutputReference
	SparkParametersInput() interface{}
	SqlServerParameters() QuicksightDataSourceDataSourceParametersSqlServerParametersOutputReference
	SqlServerParametersInput() interface{}
	StarburstParameters() QuicksightDataSourceDataSourceParametersStarburstParametersOutputReference
	StarburstParametersInput() interface{}
	TeradataParameters() QuicksightDataSourceDataSourceParametersTeradataParametersOutputReference
	TeradataParametersInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TrinoParameters() QuicksightDataSourceDataSourceParametersTrinoParametersOutputReference
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
	PutAmazonElasticsearchParameters(value *QuicksightDataSourceDataSourceParametersAmazonElasticsearchParameters)
	PutAmazonOpenSearchParameters(value *QuicksightDataSourceDataSourceParametersAmazonOpenSearchParameters)
	PutAthenaParameters(value *QuicksightDataSourceDataSourceParametersAthenaParameters)
	PutAuroraParameters(value *QuicksightDataSourceDataSourceParametersAuroraParameters)
	PutAuroraPostgreSqlParameters(value *QuicksightDataSourceDataSourceParametersAuroraPostgreSqlParameters)
	PutDatabricksParameters(value *QuicksightDataSourceDataSourceParametersDatabricksParameters)
	PutMariaDbParameters(value *QuicksightDataSourceDataSourceParametersMariaDbParameters)
	PutMySqlParameters(value *QuicksightDataSourceDataSourceParametersMySqlParameters)
	PutOracleParameters(value *QuicksightDataSourceDataSourceParametersOracleParameters)
	PutPostgreSqlParameters(value *QuicksightDataSourceDataSourceParametersPostgreSqlParameters)
	PutPrestoParameters(value *QuicksightDataSourceDataSourceParametersPrestoParameters)
	PutRdsParameters(value *QuicksightDataSourceDataSourceParametersRdsParameters)
	PutRedshiftParameters(value *QuicksightDataSourceDataSourceParametersRedshiftParameters)
	PutS3Parameters(value *QuicksightDataSourceDataSourceParametersS3Parameters)
	PutS3TablesParameters(value *QuicksightDataSourceDataSourceParametersS3TablesParameters)
	PutSnowflakeParameters(value *QuicksightDataSourceDataSourceParametersSnowflakeParameters)
	PutSparkParameters(value *QuicksightDataSourceDataSourceParametersSparkParameters)
	PutSqlServerParameters(value *QuicksightDataSourceDataSourceParametersSqlServerParameters)
	PutStarburstParameters(value *QuicksightDataSourceDataSourceParametersStarburstParameters)
	PutTeradataParameters(value *QuicksightDataSourceDataSourceParametersTeradataParameters)
	PutTrinoParameters(value *QuicksightDataSourceDataSourceParametersTrinoParameters)
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

// The jsii proxy struct for QuicksightDataSourceDataSourceParametersOutputReference
type jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) AmazonElasticsearchParameters() QuicksightDataSourceDataSourceParametersAmazonElasticsearchParametersOutputReference {
	var returns QuicksightDataSourceDataSourceParametersAmazonElasticsearchParametersOutputReference
	_jsii_.Get(
		j,
		"amazonElasticsearchParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) AmazonElasticsearchParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"amazonElasticsearchParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) AmazonOpenSearchParameters() QuicksightDataSourceDataSourceParametersAmazonOpenSearchParametersOutputReference {
	var returns QuicksightDataSourceDataSourceParametersAmazonOpenSearchParametersOutputReference
	_jsii_.Get(
		j,
		"amazonOpenSearchParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) AmazonOpenSearchParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"amazonOpenSearchParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) AthenaParameters() QuicksightDataSourceDataSourceParametersAthenaParametersOutputReference {
	var returns QuicksightDataSourceDataSourceParametersAthenaParametersOutputReference
	_jsii_.Get(
		j,
		"athenaParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) AthenaParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"athenaParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) AuroraParameters() QuicksightDataSourceDataSourceParametersAuroraParametersOutputReference {
	var returns QuicksightDataSourceDataSourceParametersAuroraParametersOutputReference
	_jsii_.Get(
		j,
		"auroraParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) AuroraParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"auroraParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) AuroraPostgreSqlParameters() QuicksightDataSourceDataSourceParametersAuroraPostgreSqlParametersOutputReference {
	var returns QuicksightDataSourceDataSourceParametersAuroraPostgreSqlParametersOutputReference
	_jsii_.Get(
		j,
		"auroraPostgreSqlParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) AuroraPostgreSqlParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"auroraPostgreSqlParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) DatabricksParameters() QuicksightDataSourceDataSourceParametersDatabricksParametersOutputReference {
	var returns QuicksightDataSourceDataSourceParametersDatabricksParametersOutputReference
	_jsii_.Get(
		j,
		"databricksParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) DatabricksParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"databricksParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) MariaDbParameters() QuicksightDataSourceDataSourceParametersMariaDbParametersOutputReference {
	var returns QuicksightDataSourceDataSourceParametersMariaDbParametersOutputReference
	_jsii_.Get(
		j,
		"mariaDbParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) MariaDbParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mariaDbParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) MySqlParameters() QuicksightDataSourceDataSourceParametersMySqlParametersOutputReference {
	var returns QuicksightDataSourceDataSourceParametersMySqlParametersOutputReference
	_jsii_.Get(
		j,
		"mySqlParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) MySqlParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mySqlParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) OracleParameters() QuicksightDataSourceDataSourceParametersOracleParametersOutputReference {
	var returns QuicksightDataSourceDataSourceParametersOracleParametersOutputReference
	_jsii_.Get(
		j,
		"oracleParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) OracleParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"oracleParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) PostgreSqlParameters() QuicksightDataSourceDataSourceParametersPostgreSqlParametersOutputReference {
	var returns QuicksightDataSourceDataSourceParametersPostgreSqlParametersOutputReference
	_jsii_.Get(
		j,
		"postgreSqlParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) PostgreSqlParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"postgreSqlParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) PrestoParameters() QuicksightDataSourceDataSourceParametersPrestoParametersOutputReference {
	var returns QuicksightDataSourceDataSourceParametersPrestoParametersOutputReference
	_jsii_.Get(
		j,
		"prestoParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) PrestoParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"prestoParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) RdsParameters() QuicksightDataSourceDataSourceParametersRdsParametersOutputReference {
	var returns QuicksightDataSourceDataSourceParametersRdsParametersOutputReference
	_jsii_.Get(
		j,
		"rdsParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) RdsParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rdsParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) RedshiftParameters() QuicksightDataSourceDataSourceParametersRedshiftParametersOutputReference {
	var returns QuicksightDataSourceDataSourceParametersRedshiftParametersOutputReference
	_jsii_.Get(
		j,
		"redshiftParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) RedshiftParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"redshiftParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) S3Parameters() QuicksightDataSourceDataSourceParametersS3ParametersOutputReference {
	var returns QuicksightDataSourceDataSourceParametersS3ParametersOutputReference
	_jsii_.Get(
		j,
		"s3Parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) S3ParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3ParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) S3TablesParameters() QuicksightDataSourceDataSourceParametersS3TablesParametersOutputReference {
	var returns QuicksightDataSourceDataSourceParametersS3TablesParametersOutputReference
	_jsii_.Get(
		j,
		"s3TablesParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) S3TablesParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3TablesParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) SnowflakeParameters() QuicksightDataSourceDataSourceParametersSnowflakeParametersOutputReference {
	var returns QuicksightDataSourceDataSourceParametersSnowflakeParametersOutputReference
	_jsii_.Get(
		j,
		"snowflakeParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) SnowflakeParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"snowflakeParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) SparkParameters() QuicksightDataSourceDataSourceParametersSparkParametersOutputReference {
	var returns QuicksightDataSourceDataSourceParametersSparkParametersOutputReference
	_jsii_.Get(
		j,
		"sparkParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) SparkParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sparkParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) SqlServerParameters() QuicksightDataSourceDataSourceParametersSqlServerParametersOutputReference {
	var returns QuicksightDataSourceDataSourceParametersSqlServerParametersOutputReference
	_jsii_.Get(
		j,
		"sqlServerParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) SqlServerParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sqlServerParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) StarburstParameters() QuicksightDataSourceDataSourceParametersStarburstParametersOutputReference {
	var returns QuicksightDataSourceDataSourceParametersStarburstParametersOutputReference
	_jsii_.Get(
		j,
		"starburstParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) StarburstParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"starburstParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) TeradataParameters() QuicksightDataSourceDataSourceParametersTeradataParametersOutputReference {
	var returns QuicksightDataSourceDataSourceParametersTeradataParametersOutputReference
	_jsii_.Get(
		j,
		"teradataParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) TeradataParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"teradataParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) TrinoParameters() QuicksightDataSourceDataSourceParametersTrinoParametersOutputReference {
	var returns QuicksightDataSourceDataSourceParametersTrinoParametersOutputReference
	_jsii_.Get(
		j,
		"trinoParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) TrinoParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trinoParametersInput",
		&returns,
	)
	return returns
}


func NewQuicksightDataSourceDataSourceParametersOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) QuicksightDataSourceDataSourceParametersOutputReference {
	_init_.Initialize()

	if err := validateNewQuicksightDataSourceDataSourceParametersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.quicksightDataSource.QuicksightDataSourceDataSourceParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewQuicksightDataSourceDataSourceParametersOutputReference_Override(q QuicksightDataSourceDataSourceParametersOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.quicksightDataSource.QuicksightDataSourceDataSourceParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		q,
	)
}

func (j *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (q *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (q *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (q *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (q *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (q *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (q *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (q *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (q *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (q *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (q *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (q *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) PutAmazonElasticsearchParameters(value *QuicksightDataSourceDataSourceParametersAmazonElasticsearchParameters) {
	if err := q.validatePutAmazonElasticsearchParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putAmazonElasticsearchParameters",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) PutAmazonOpenSearchParameters(value *QuicksightDataSourceDataSourceParametersAmazonOpenSearchParameters) {
	if err := q.validatePutAmazonOpenSearchParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putAmazonOpenSearchParameters",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) PutAthenaParameters(value *QuicksightDataSourceDataSourceParametersAthenaParameters) {
	if err := q.validatePutAthenaParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putAthenaParameters",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) PutAuroraParameters(value *QuicksightDataSourceDataSourceParametersAuroraParameters) {
	if err := q.validatePutAuroraParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putAuroraParameters",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) PutAuroraPostgreSqlParameters(value *QuicksightDataSourceDataSourceParametersAuroraPostgreSqlParameters) {
	if err := q.validatePutAuroraPostgreSqlParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putAuroraPostgreSqlParameters",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) PutDatabricksParameters(value *QuicksightDataSourceDataSourceParametersDatabricksParameters) {
	if err := q.validatePutDatabricksParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putDatabricksParameters",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) PutMariaDbParameters(value *QuicksightDataSourceDataSourceParametersMariaDbParameters) {
	if err := q.validatePutMariaDbParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putMariaDbParameters",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) PutMySqlParameters(value *QuicksightDataSourceDataSourceParametersMySqlParameters) {
	if err := q.validatePutMySqlParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putMySqlParameters",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) PutOracleParameters(value *QuicksightDataSourceDataSourceParametersOracleParameters) {
	if err := q.validatePutOracleParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putOracleParameters",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) PutPostgreSqlParameters(value *QuicksightDataSourceDataSourceParametersPostgreSqlParameters) {
	if err := q.validatePutPostgreSqlParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putPostgreSqlParameters",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) PutPrestoParameters(value *QuicksightDataSourceDataSourceParametersPrestoParameters) {
	if err := q.validatePutPrestoParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putPrestoParameters",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) PutRdsParameters(value *QuicksightDataSourceDataSourceParametersRdsParameters) {
	if err := q.validatePutRdsParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putRdsParameters",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) PutRedshiftParameters(value *QuicksightDataSourceDataSourceParametersRedshiftParameters) {
	if err := q.validatePutRedshiftParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putRedshiftParameters",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) PutS3Parameters(value *QuicksightDataSourceDataSourceParametersS3Parameters) {
	if err := q.validatePutS3ParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putS3Parameters",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) PutS3TablesParameters(value *QuicksightDataSourceDataSourceParametersS3TablesParameters) {
	if err := q.validatePutS3TablesParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putS3TablesParameters",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) PutSnowflakeParameters(value *QuicksightDataSourceDataSourceParametersSnowflakeParameters) {
	if err := q.validatePutSnowflakeParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putSnowflakeParameters",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) PutSparkParameters(value *QuicksightDataSourceDataSourceParametersSparkParameters) {
	if err := q.validatePutSparkParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putSparkParameters",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) PutSqlServerParameters(value *QuicksightDataSourceDataSourceParametersSqlServerParameters) {
	if err := q.validatePutSqlServerParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putSqlServerParameters",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) PutStarburstParameters(value *QuicksightDataSourceDataSourceParametersStarburstParameters) {
	if err := q.validatePutStarburstParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putStarburstParameters",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) PutTeradataParameters(value *QuicksightDataSourceDataSourceParametersTeradataParameters) {
	if err := q.validatePutTeradataParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putTeradataParameters",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) PutTrinoParameters(value *QuicksightDataSourceDataSourceParametersTrinoParameters) {
	if err := q.validatePutTrinoParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putTrinoParameters",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) ResetAmazonElasticsearchParameters() {
	_jsii_.InvokeVoid(
		q,
		"resetAmazonElasticsearchParameters",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) ResetAmazonOpenSearchParameters() {
	_jsii_.InvokeVoid(
		q,
		"resetAmazonOpenSearchParameters",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) ResetAthenaParameters() {
	_jsii_.InvokeVoid(
		q,
		"resetAthenaParameters",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) ResetAuroraParameters() {
	_jsii_.InvokeVoid(
		q,
		"resetAuroraParameters",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) ResetAuroraPostgreSqlParameters() {
	_jsii_.InvokeVoid(
		q,
		"resetAuroraPostgreSqlParameters",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) ResetDatabricksParameters() {
	_jsii_.InvokeVoid(
		q,
		"resetDatabricksParameters",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) ResetMariaDbParameters() {
	_jsii_.InvokeVoid(
		q,
		"resetMariaDbParameters",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) ResetMySqlParameters() {
	_jsii_.InvokeVoid(
		q,
		"resetMySqlParameters",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) ResetOracleParameters() {
	_jsii_.InvokeVoid(
		q,
		"resetOracleParameters",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) ResetPostgreSqlParameters() {
	_jsii_.InvokeVoid(
		q,
		"resetPostgreSqlParameters",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) ResetPrestoParameters() {
	_jsii_.InvokeVoid(
		q,
		"resetPrestoParameters",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) ResetRdsParameters() {
	_jsii_.InvokeVoid(
		q,
		"resetRdsParameters",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) ResetRedshiftParameters() {
	_jsii_.InvokeVoid(
		q,
		"resetRedshiftParameters",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) ResetS3Parameters() {
	_jsii_.InvokeVoid(
		q,
		"resetS3Parameters",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) ResetS3TablesParameters() {
	_jsii_.InvokeVoid(
		q,
		"resetS3TablesParameters",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) ResetSnowflakeParameters() {
	_jsii_.InvokeVoid(
		q,
		"resetSnowflakeParameters",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) ResetSparkParameters() {
	_jsii_.InvokeVoid(
		q,
		"resetSparkParameters",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) ResetSqlServerParameters() {
	_jsii_.InvokeVoid(
		q,
		"resetSqlServerParameters",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) ResetStarburstParameters() {
	_jsii_.InvokeVoid(
		q,
		"resetStarburstParameters",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) ResetTeradataParameters() {
	_jsii_.InvokeVoid(
		q,
		"resetTeradataParameters",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) ResetTrinoParameters() {
	_jsii_.InvokeVoid(
		q,
		"resetTrinoParameters",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (q *jsiiProxy_QuicksightDataSourceDataSourceParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

