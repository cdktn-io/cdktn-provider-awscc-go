// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccquicksightdatasource

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawsccquicksightdatasource/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccQuicksightDataSourceDataSourceParametersOutputReference interface {
	cdktn.ComplexObject
	AmazonElasticsearchParameters() DataAwsccQuicksightDataSourceDataSourceParametersAmazonElasticsearchParametersOutputReference
	AmazonOpenSearchParameters() DataAwsccQuicksightDataSourceDataSourceParametersAmazonOpenSearchParametersOutputReference
	AthenaParameters() DataAwsccQuicksightDataSourceDataSourceParametersAthenaParametersOutputReference
	AuroraParameters() DataAwsccQuicksightDataSourceDataSourceParametersAuroraParametersOutputReference
	AuroraPostgreSqlParameters() DataAwsccQuicksightDataSourceDataSourceParametersAuroraPostgreSqlParametersOutputReference
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
	DatabricksParameters() DataAwsccQuicksightDataSourceDataSourceParametersDatabricksParametersOutputReference
	// Experimental.
	Fqn() *string
	InternalValue() *DataAwsccQuicksightDataSourceDataSourceParameters
	SetInternalValue(val *DataAwsccQuicksightDataSourceDataSourceParameters)
	MariaDbParameters() DataAwsccQuicksightDataSourceDataSourceParametersMariaDbParametersOutputReference
	MySqlParameters() DataAwsccQuicksightDataSourceDataSourceParametersMySqlParametersOutputReference
	OracleParameters() DataAwsccQuicksightDataSourceDataSourceParametersOracleParametersOutputReference
	PostgreSqlParameters() DataAwsccQuicksightDataSourceDataSourceParametersPostgreSqlParametersOutputReference
	PrestoParameters() DataAwsccQuicksightDataSourceDataSourceParametersPrestoParametersOutputReference
	RdsParameters() DataAwsccQuicksightDataSourceDataSourceParametersRdsParametersOutputReference
	RedshiftParameters() DataAwsccQuicksightDataSourceDataSourceParametersRedshiftParametersOutputReference
	S3Parameters() DataAwsccQuicksightDataSourceDataSourceParametersS3ParametersOutputReference
	S3TablesParameters() DataAwsccQuicksightDataSourceDataSourceParametersS3TablesParametersOutputReference
	SnowflakeParameters() DataAwsccQuicksightDataSourceDataSourceParametersSnowflakeParametersOutputReference
	SparkParameters() DataAwsccQuicksightDataSourceDataSourceParametersSparkParametersOutputReference
	SqlServerParameters() DataAwsccQuicksightDataSourceDataSourceParametersSqlServerParametersOutputReference
	StarburstParameters() DataAwsccQuicksightDataSourceDataSourceParametersStarburstParametersOutputReference
	TeradataParameters() DataAwsccQuicksightDataSourceDataSourceParametersTeradataParametersOutputReference
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TrinoParameters() DataAwsccQuicksightDataSourceDataSourceParametersTrinoParametersOutputReference
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsccQuicksightDataSourceDataSourceParametersOutputReference
type jsiiProxy_DataAwsccQuicksightDataSourceDataSourceParametersOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceDataSourceParametersOutputReference) AmazonElasticsearchParameters() DataAwsccQuicksightDataSourceDataSourceParametersAmazonElasticsearchParametersOutputReference {
	var returns DataAwsccQuicksightDataSourceDataSourceParametersAmazonElasticsearchParametersOutputReference
	_jsii_.Get(
		j,
		"amazonElasticsearchParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceDataSourceParametersOutputReference) AmazonOpenSearchParameters() DataAwsccQuicksightDataSourceDataSourceParametersAmazonOpenSearchParametersOutputReference {
	var returns DataAwsccQuicksightDataSourceDataSourceParametersAmazonOpenSearchParametersOutputReference
	_jsii_.Get(
		j,
		"amazonOpenSearchParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceDataSourceParametersOutputReference) AthenaParameters() DataAwsccQuicksightDataSourceDataSourceParametersAthenaParametersOutputReference {
	var returns DataAwsccQuicksightDataSourceDataSourceParametersAthenaParametersOutputReference
	_jsii_.Get(
		j,
		"athenaParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceDataSourceParametersOutputReference) AuroraParameters() DataAwsccQuicksightDataSourceDataSourceParametersAuroraParametersOutputReference {
	var returns DataAwsccQuicksightDataSourceDataSourceParametersAuroraParametersOutputReference
	_jsii_.Get(
		j,
		"auroraParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceDataSourceParametersOutputReference) AuroraPostgreSqlParameters() DataAwsccQuicksightDataSourceDataSourceParametersAuroraPostgreSqlParametersOutputReference {
	var returns DataAwsccQuicksightDataSourceDataSourceParametersAuroraPostgreSqlParametersOutputReference
	_jsii_.Get(
		j,
		"auroraPostgreSqlParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceDataSourceParametersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceDataSourceParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceDataSourceParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceDataSourceParametersOutputReference) DatabricksParameters() DataAwsccQuicksightDataSourceDataSourceParametersDatabricksParametersOutputReference {
	var returns DataAwsccQuicksightDataSourceDataSourceParametersDatabricksParametersOutputReference
	_jsii_.Get(
		j,
		"databricksParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceDataSourceParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceDataSourceParametersOutputReference) InternalValue() *DataAwsccQuicksightDataSourceDataSourceParameters {
	var returns *DataAwsccQuicksightDataSourceDataSourceParameters
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceDataSourceParametersOutputReference) MariaDbParameters() DataAwsccQuicksightDataSourceDataSourceParametersMariaDbParametersOutputReference {
	var returns DataAwsccQuicksightDataSourceDataSourceParametersMariaDbParametersOutputReference
	_jsii_.Get(
		j,
		"mariaDbParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceDataSourceParametersOutputReference) MySqlParameters() DataAwsccQuicksightDataSourceDataSourceParametersMySqlParametersOutputReference {
	var returns DataAwsccQuicksightDataSourceDataSourceParametersMySqlParametersOutputReference
	_jsii_.Get(
		j,
		"mySqlParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceDataSourceParametersOutputReference) OracleParameters() DataAwsccQuicksightDataSourceDataSourceParametersOracleParametersOutputReference {
	var returns DataAwsccQuicksightDataSourceDataSourceParametersOracleParametersOutputReference
	_jsii_.Get(
		j,
		"oracleParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceDataSourceParametersOutputReference) PostgreSqlParameters() DataAwsccQuicksightDataSourceDataSourceParametersPostgreSqlParametersOutputReference {
	var returns DataAwsccQuicksightDataSourceDataSourceParametersPostgreSqlParametersOutputReference
	_jsii_.Get(
		j,
		"postgreSqlParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceDataSourceParametersOutputReference) PrestoParameters() DataAwsccQuicksightDataSourceDataSourceParametersPrestoParametersOutputReference {
	var returns DataAwsccQuicksightDataSourceDataSourceParametersPrestoParametersOutputReference
	_jsii_.Get(
		j,
		"prestoParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceDataSourceParametersOutputReference) RdsParameters() DataAwsccQuicksightDataSourceDataSourceParametersRdsParametersOutputReference {
	var returns DataAwsccQuicksightDataSourceDataSourceParametersRdsParametersOutputReference
	_jsii_.Get(
		j,
		"rdsParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceDataSourceParametersOutputReference) RedshiftParameters() DataAwsccQuicksightDataSourceDataSourceParametersRedshiftParametersOutputReference {
	var returns DataAwsccQuicksightDataSourceDataSourceParametersRedshiftParametersOutputReference
	_jsii_.Get(
		j,
		"redshiftParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceDataSourceParametersOutputReference) S3Parameters() DataAwsccQuicksightDataSourceDataSourceParametersS3ParametersOutputReference {
	var returns DataAwsccQuicksightDataSourceDataSourceParametersS3ParametersOutputReference
	_jsii_.Get(
		j,
		"s3Parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceDataSourceParametersOutputReference) S3TablesParameters() DataAwsccQuicksightDataSourceDataSourceParametersS3TablesParametersOutputReference {
	var returns DataAwsccQuicksightDataSourceDataSourceParametersS3TablesParametersOutputReference
	_jsii_.Get(
		j,
		"s3TablesParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceDataSourceParametersOutputReference) SnowflakeParameters() DataAwsccQuicksightDataSourceDataSourceParametersSnowflakeParametersOutputReference {
	var returns DataAwsccQuicksightDataSourceDataSourceParametersSnowflakeParametersOutputReference
	_jsii_.Get(
		j,
		"snowflakeParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceDataSourceParametersOutputReference) SparkParameters() DataAwsccQuicksightDataSourceDataSourceParametersSparkParametersOutputReference {
	var returns DataAwsccQuicksightDataSourceDataSourceParametersSparkParametersOutputReference
	_jsii_.Get(
		j,
		"sparkParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceDataSourceParametersOutputReference) SqlServerParameters() DataAwsccQuicksightDataSourceDataSourceParametersSqlServerParametersOutputReference {
	var returns DataAwsccQuicksightDataSourceDataSourceParametersSqlServerParametersOutputReference
	_jsii_.Get(
		j,
		"sqlServerParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceDataSourceParametersOutputReference) StarburstParameters() DataAwsccQuicksightDataSourceDataSourceParametersStarburstParametersOutputReference {
	var returns DataAwsccQuicksightDataSourceDataSourceParametersStarburstParametersOutputReference
	_jsii_.Get(
		j,
		"starburstParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceDataSourceParametersOutputReference) TeradataParameters() DataAwsccQuicksightDataSourceDataSourceParametersTeradataParametersOutputReference {
	var returns DataAwsccQuicksightDataSourceDataSourceParametersTeradataParametersOutputReference
	_jsii_.Get(
		j,
		"teradataParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceDataSourceParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceDataSourceParametersOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceDataSourceParametersOutputReference) TrinoParameters() DataAwsccQuicksightDataSourceDataSourceParametersTrinoParametersOutputReference {
	var returns DataAwsccQuicksightDataSourceDataSourceParametersTrinoParametersOutputReference
	_jsii_.Get(
		j,
		"trinoParameters",
		&returns,
	)
	return returns
}


func NewDataAwsccQuicksightDataSourceDataSourceParametersOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataAwsccQuicksightDataSourceDataSourceParametersOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccQuicksightDataSourceDataSourceParametersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccQuicksightDataSourceDataSourceParametersOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccQuicksightDataSource.DataAwsccQuicksightDataSourceDataSourceParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccQuicksightDataSourceDataSourceParametersOutputReference_Override(d DataAwsccQuicksightDataSourceDataSourceParametersOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccQuicksightDataSource.DataAwsccQuicksightDataSourceDataSourceParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceDataSourceParametersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceDataSourceParametersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceDataSourceParametersOutputReference)SetInternalValue(val *DataAwsccQuicksightDataSourceDataSourceParameters) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceDataSourceParametersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceDataSourceParametersOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccQuicksightDataSourceDataSourceParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccQuicksightDataSourceDataSourceParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccQuicksightDataSourceDataSourceParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccQuicksightDataSourceDataSourceParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccQuicksightDataSourceDataSourceParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccQuicksightDataSourceDataSourceParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccQuicksightDataSourceDataSourceParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccQuicksightDataSourceDataSourceParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccQuicksightDataSourceDataSourceParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccQuicksightDataSourceDataSourceParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccQuicksightDataSourceDataSourceParametersOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccQuicksightDataSourceDataSourceParametersOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccQuicksightDataSourceDataSourceParametersOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccQuicksightDataSourceDataSourceParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

