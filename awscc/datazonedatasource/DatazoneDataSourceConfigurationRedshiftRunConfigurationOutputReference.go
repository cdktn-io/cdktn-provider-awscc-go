// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datazonedatasource

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/datazonedatasource/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DatazoneDataSourceConfigurationRedshiftRunConfigurationOutputReference interface {
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
	DataAccessRole() *string
	SetDataAccessRole(val *string)
	DataAccessRoleInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	RedshiftCredentialConfiguration() DatazoneDataSourceConfigurationRedshiftRunConfigurationRedshiftCredentialConfigurationOutputReference
	RedshiftCredentialConfigurationInput() interface{}
	RedshiftStorage() DatazoneDataSourceConfigurationRedshiftRunConfigurationRedshiftStorageOutputReference
	RedshiftStorageInput() interface{}
	RelationalFilterConfigurations() DatazoneDataSourceConfigurationRedshiftRunConfigurationRelationalFilterConfigurationsList
	RelationalFilterConfigurationsInput() interface{}
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
	PutRedshiftCredentialConfiguration(value *DatazoneDataSourceConfigurationRedshiftRunConfigurationRedshiftCredentialConfiguration)
	PutRedshiftStorage(value *DatazoneDataSourceConfigurationRedshiftRunConfigurationRedshiftStorage)
	PutRelationalFilterConfigurations(value interface{})
	ResetDataAccessRole()
	ResetRedshiftCredentialConfiguration()
	ResetRedshiftStorage()
	ResetRelationalFilterConfigurations()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DatazoneDataSourceConfigurationRedshiftRunConfigurationOutputReference
type jsiiProxy_DatazoneDataSourceConfigurationRedshiftRunConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DatazoneDataSourceConfigurationRedshiftRunConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneDataSourceConfigurationRedshiftRunConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneDataSourceConfigurationRedshiftRunConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneDataSourceConfigurationRedshiftRunConfigurationOutputReference) DataAccessRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataAccessRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneDataSourceConfigurationRedshiftRunConfigurationOutputReference) DataAccessRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataAccessRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneDataSourceConfigurationRedshiftRunConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneDataSourceConfigurationRedshiftRunConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneDataSourceConfigurationRedshiftRunConfigurationOutputReference) RedshiftCredentialConfiguration() DatazoneDataSourceConfigurationRedshiftRunConfigurationRedshiftCredentialConfigurationOutputReference {
	var returns DatazoneDataSourceConfigurationRedshiftRunConfigurationRedshiftCredentialConfigurationOutputReference
	_jsii_.Get(
		j,
		"redshiftCredentialConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneDataSourceConfigurationRedshiftRunConfigurationOutputReference) RedshiftCredentialConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"redshiftCredentialConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneDataSourceConfigurationRedshiftRunConfigurationOutputReference) RedshiftStorage() DatazoneDataSourceConfigurationRedshiftRunConfigurationRedshiftStorageOutputReference {
	var returns DatazoneDataSourceConfigurationRedshiftRunConfigurationRedshiftStorageOutputReference
	_jsii_.Get(
		j,
		"redshiftStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneDataSourceConfigurationRedshiftRunConfigurationOutputReference) RedshiftStorageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"redshiftStorageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneDataSourceConfigurationRedshiftRunConfigurationOutputReference) RelationalFilterConfigurations() DatazoneDataSourceConfigurationRedshiftRunConfigurationRelationalFilterConfigurationsList {
	var returns DatazoneDataSourceConfigurationRedshiftRunConfigurationRelationalFilterConfigurationsList
	_jsii_.Get(
		j,
		"relationalFilterConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneDataSourceConfigurationRedshiftRunConfigurationOutputReference) RelationalFilterConfigurationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"relationalFilterConfigurationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneDataSourceConfigurationRedshiftRunConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneDataSourceConfigurationRedshiftRunConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDatazoneDataSourceConfigurationRedshiftRunConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DatazoneDataSourceConfigurationRedshiftRunConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewDatazoneDataSourceConfigurationRedshiftRunConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatazoneDataSourceConfigurationRedshiftRunConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.datazoneDataSource.DatazoneDataSourceConfigurationRedshiftRunConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDatazoneDataSourceConfigurationRedshiftRunConfigurationOutputReference_Override(d DatazoneDataSourceConfigurationRedshiftRunConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.datazoneDataSource.DatazoneDataSourceConfigurationRedshiftRunConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DatazoneDataSourceConfigurationRedshiftRunConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DatazoneDataSourceConfigurationRedshiftRunConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DatazoneDataSourceConfigurationRedshiftRunConfigurationOutputReference)SetDataAccessRole(val *string) {
	if err := j.validateSetDataAccessRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataAccessRole",
		val,
	)
}

func (j *jsiiProxy_DatazoneDataSourceConfigurationRedshiftRunConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatazoneDataSourceConfigurationRedshiftRunConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatazoneDataSourceConfigurationRedshiftRunConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DatazoneDataSourceConfigurationRedshiftRunConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatazoneDataSourceConfigurationRedshiftRunConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DatazoneDataSourceConfigurationRedshiftRunConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DatazoneDataSourceConfigurationRedshiftRunConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DatazoneDataSourceConfigurationRedshiftRunConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DatazoneDataSourceConfigurationRedshiftRunConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DatazoneDataSourceConfigurationRedshiftRunConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DatazoneDataSourceConfigurationRedshiftRunConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DatazoneDataSourceConfigurationRedshiftRunConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DatazoneDataSourceConfigurationRedshiftRunConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DatazoneDataSourceConfigurationRedshiftRunConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatazoneDataSourceConfigurationRedshiftRunConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DatazoneDataSourceConfigurationRedshiftRunConfigurationOutputReference) PutRedshiftCredentialConfiguration(value *DatazoneDataSourceConfigurationRedshiftRunConfigurationRedshiftCredentialConfiguration) {
	if err := d.validatePutRedshiftCredentialConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRedshiftCredentialConfiguration",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatazoneDataSourceConfigurationRedshiftRunConfigurationOutputReference) PutRedshiftStorage(value *DatazoneDataSourceConfigurationRedshiftRunConfigurationRedshiftStorage) {
	if err := d.validatePutRedshiftStorageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRedshiftStorage",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatazoneDataSourceConfigurationRedshiftRunConfigurationOutputReference) PutRelationalFilterConfigurations(value interface{}) {
	if err := d.validatePutRelationalFilterConfigurationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRelationalFilterConfigurations",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatazoneDataSourceConfigurationRedshiftRunConfigurationOutputReference) ResetDataAccessRole() {
	_jsii_.InvokeVoid(
		d,
		"resetDataAccessRole",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazoneDataSourceConfigurationRedshiftRunConfigurationOutputReference) ResetRedshiftCredentialConfiguration() {
	_jsii_.InvokeVoid(
		d,
		"resetRedshiftCredentialConfiguration",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazoneDataSourceConfigurationRedshiftRunConfigurationOutputReference) ResetRedshiftStorage() {
	_jsii_.InvokeVoid(
		d,
		"resetRedshiftStorage",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazoneDataSourceConfigurationRedshiftRunConfigurationOutputReference) ResetRelationalFilterConfigurations() {
	_jsii_.InvokeVoid(
		d,
		"resetRelationalFilterConfigurations",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazoneDataSourceConfigurationRedshiftRunConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DatazoneDataSourceConfigurationRedshiftRunConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

