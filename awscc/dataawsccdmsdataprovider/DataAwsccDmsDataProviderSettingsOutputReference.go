// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccdmsdataprovider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawsccdmsdataprovider/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccDmsDataProviderSettingsOutputReference interface {
	cdktn.ComplexObject
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
	DocDbSettings() DataAwsccDmsDataProviderSettingsDocDbSettingsOutputReference
	// Experimental.
	Fqn() *string
	IbmDb2LuwSettings() DataAwsccDmsDataProviderSettingsIbmDb2LuwSettingsOutputReference
	IbmDb2ZOsSettings() DataAwsccDmsDataProviderSettingsIbmDb2ZOsSettingsOutputReference
	InternalValue() *DataAwsccDmsDataProviderSettings
	SetInternalValue(val *DataAwsccDmsDataProviderSettings)
	MariaDbSettings() DataAwsccDmsDataProviderSettingsMariaDbSettingsOutputReference
	MicrosoftSqlServerSettings() DataAwsccDmsDataProviderSettingsMicrosoftSqlServerSettingsOutputReference
	MongoDbSettings() DataAwsccDmsDataProviderSettingsMongoDbSettingsOutputReference
	MySqlSettings() DataAwsccDmsDataProviderSettingsMySqlSettingsOutputReference
	OracleSettings() DataAwsccDmsDataProviderSettingsOracleSettingsOutputReference
	PostgreSqlSettings() DataAwsccDmsDataProviderSettingsPostgreSqlSettingsOutputReference
	RedshiftSettings() DataAwsccDmsDataProviderSettingsRedshiftSettingsOutputReference
	SybaseAseSettings() DataAwsccDmsDataProviderSettingsSybaseAseSettingsOutputReference
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
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

// The jsii proxy struct for DataAwsccDmsDataProviderSettingsOutputReference
type jsiiProxy_DataAwsccDmsDataProviderSettingsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccDmsDataProviderSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsDataProviderSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsDataProviderSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsDataProviderSettingsOutputReference) DocDbSettings() DataAwsccDmsDataProviderSettingsDocDbSettingsOutputReference {
	var returns DataAwsccDmsDataProviderSettingsDocDbSettingsOutputReference
	_jsii_.Get(
		j,
		"docDbSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsDataProviderSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsDataProviderSettingsOutputReference) IbmDb2LuwSettings() DataAwsccDmsDataProviderSettingsIbmDb2LuwSettingsOutputReference {
	var returns DataAwsccDmsDataProviderSettingsIbmDb2LuwSettingsOutputReference
	_jsii_.Get(
		j,
		"ibmDb2LuwSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsDataProviderSettingsOutputReference) IbmDb2ZOsSettings() DataAwsccDmsDataProviderSettingsIbmDb2ZOsSettingsOutputReference {
	var returns DataAwsccDmsDataProviderSettingsIbmDb2ZOsSettingsOutputReference
	_jsii_.Get(
		j,
		"ibmDb2ZOsSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsDataProviderSettingsOutputReference) InternalValue() *DataAwsccDmsDataProviderSettings {
	var returns *DataAwsccDmsDataProviderSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsDataProviderSettingsOutputReference) MariaDbSettings() DataAwsccDmsDataProviderSettingsMariaDbSettingsOutputReference {
	var returns DataAwsccDmsDataProviderSettingsMariaDbSettingsOutputReference
	_jsii_.Get(
		j,
		"mariaDbSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsDataProviderSettingsOutputReference) MicrosoftSqlServerSettings() DataAwsccDmsDataProviderSettingsMicrosoftSqlServerSettingsOutputReference {
	var returns DataAwsccDmsDataProviderSettingsMicrosoftSqlServerSettingsOutputReference
	_jsii_.Get(
		j,
		"microsoftSqlServerSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsDataProviderSettingsOutputReference) MongoDbSettings() DataAwsccDmsDataProviderSettingsMongoDbSettingsOutputReference {
	var returns DataAwsccDmsDataProviderSettingsMongoDbSettingsOutputReference
	_jsii_.Get(
		j,
		"mongoDbSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsDataProviderSettingsOutputReference) MySqlSettings() DataAwsccDmsDataProviderSettingsMySqlSettingsOutputReference {
	var returns DataAwsccDmsDataProviderSettingsMySqlSettingsOutputReference
	_jsii_.Get(
		j,
		"mySqlSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsDataProviderSettingsOutputReference) OracleSettings() DataAwsccDmsDataProviderSettingsOracleSettingsOutputReference {
	var returns DataAwsccDmsDataProviderSettingsOracleSettingsOutputReference
	_jsii_.Get(
		j,
		"oracleSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsDataProviderSettingsOutputReference) PostgreSqlSettings() DataAwsccDmsDataProviderSettingsPostgreSqlSettingsOutputReference {
	var returns DataAwsccDmsDataProviderSettingsPostgreSqlSettingsOutputReference
	_jsii_.Get(
		j,
		"postgreSqlSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsDataProviderSettingsOutputReference) RedshiftSettings() DataAwsccDmsDataProviderSettingsRedshiftSettingsOutputReference {
	var returns DataAwsccDmsDataProviderSettingsRedshiftSettingsOutputReference
	_jsii_.Get(
		j,
		"redshiftSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsDataProviderSettingsOutputReference) SybaseAseSettings() DataAwsccDmsDataProviderSettingsSybaseAseSettingsOutputReference {
	var returns DataAwsccDmsDataProviderSettingsSybaseAseSettingsOutputReference
	_jsii_.Get(
		j,
		"sybaseAseSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsDataProviderSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsDataProviderSettingsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataAwsccDmsDataProviderSettingsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataAwsccDmsDataProviderSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccDmsDataProviderSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccDmsDataProviderSettingsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccDmsDataProvider.DataAwsccDmsDataProviderSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccDmsDataProviderSettingsOutputReference_Override(d DataAwsccDmsDataProviderSettingsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccDmsDataProvider.DataAwsccDmsDataProviderSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccDmsDataProviderSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccDmsDataProviderSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccDmsDataProviderSettingsOutputReference)SetInternalValue(val *DataAwsccDmsDataProviderSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccDmsDataProviderSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccDmsDataProviderSettingsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccDmsDataProviderSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccDmsDataProviderSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccDmsDataProviderSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccDmsDataProviderSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccDmsDataProviderSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccDmsDataProviderSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccDmsDataProviderSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccDmsDataProviderSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccDmsDataProviderSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccDmsDataProviderSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccDmsDataProviderSettingsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccDmsDataProviderSettingsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccDmsDataProviderSettingsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccDmsDataProviderSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

