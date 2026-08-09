// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package databrewdataset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/databrewdataset/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DatabrewDatasetInputDatabaseInputDefinitionOutputReference interface {
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
	DatabaseTableName() *string
	SetDatabaseTableName(val *string)
	DatabaseTableNameInput() *string
	// Experimental.
	Fqn() *string
	GlueConnectionName() *string
	SetGlueConnectionName(val *string)
	GlueConnectionNameInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	QueryString() *string
	SetQueryString(val *string)
	QueryStringInput() *string
	TempDirectory() DatabrewDatasetInputDatabaseInputDefinitionTempDirectoryOutputReference
	TempDirectoryInput() interface{}
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
	PutTempDirectory(value *DatabrewDatasetInputDatabaseInputDefinitionTempDirectory)
	ResetDatabaseTableName()
	ResetGlueConnectionName()
	ResetQueryString()
	ResetTempDirectory()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DatabrewDatasetInputDatabaseInputDefinitionOutputReference
type jsiiProxy_DatabrewDatasetInputDatabaseInputDefinitionOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DatabrewDatasetInputDatabaseInputDefinitionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewDatasetInputDatabaseInputDefinitionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewDatasetInputDatabaseInputDefinitionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewDatasetInputDatabaseInputDefinitionOutputReference) DatabaseTableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseTableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewDatasetInputDatabaseInputDefinitionOutputReference) DatabaseTableNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseTableNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewDatasetInputDatabaseInputDefinitionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewDatasetInputDatabaseInputDefinitionOutputReference) GlueConnectionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"glueConnectionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewDatasetInputDatabaseInputDefinitionOutputReference) GlueConnectionNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"glueConnectionNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewDatasetInputDatabaseInputDefinitionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewDatasetInputDatabaseInputDefinitionOutputReference) QueryString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewDatasetInputDatabaseInputDefinitionOutputReference) QueryStringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewDatasetInputDatabaseInputDefinitionOutputReference) TempDirectory() DatabrewDatasetInputDatabaseInputDefinitionTempDirectoryOutputReference {
	var returns DatabrewDatasetInputDatabaseInputDefinitionTempDirectoryOutputReference
	_jsii_.Get(
		j,
		"tempDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewDatasetInputDatabaseInputDefinitionOutputReference) TempDirectoryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tempDirectoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewDatasetInputDatabaseInputDefinitionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewDatasetInputDatabaseInputDefinitionOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDatabrewDatasetInputDatabaseInputDefinitionOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DatabrewDatasetInputDatabaseInputDefinitionOutputReference {
	_init_.Initialize()

	if err := validateNewDatabrewDatasetInputDatabaseInputDefinitionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatabrewDatasetInputDatabaseInputDefinitionOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.databrewDataset.DatabrewDatasetInputDatabaseInputDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDatabrewDatasetInputDatabaseInputDefinitionOutputReference_Override(d DatabrewDatasetInputDatabaseInputDefinitionOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.databrewDataset.DatabrewDatasetInputDatabaseInputDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DatabrewDatasetInputDatabaseInputDefinitionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DatabrewDatasetInputDatabaseInputDefinitionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DatabrewDatasetInputDatabaseInputDefinitionOutputReference)SetDatabaseTableName(val *string) {
	if err := j.validateSetDatabaseTableNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseTableName",
		val,
	)
}

func (j *jsiiProxy_DatabrewDatasetInputDatabaseInputDefinitionOutputReference)SetGlueConnectionName(val *string) {
	if err := j.validateSetGlueConnectionNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"glueConnectionName",
		val,
	)
}

func (j *jsiiProxy_DatabrewDatasetInputDatabaseInputDefinitionOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatabrewDatasetInputDatabaseInputDefinitionOutputReference)SetQueryString(val *string) {
	if err := j.validateSetQueryStringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryString",
		val,
	)
}

func (j *jsiiProxy_DatabrewDatasetInputDatabaseInputDefinitionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatabrewDatasetInputDatabaseInputDefinitionOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DatabrewDatasetInputDatabaseInputDefinitionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabrewDatasetInputDatabaseInputDefinitionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DatabrewDatasetInputDatabaseInputDefinitionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DatabrewDatasetInputDatabaseInputDefinitionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DatabrewDatasetInputDatabaseInputDefinitionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DatabrewDatasetInputDatabaseInputDefinitionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DatabrewDatasetInputDatabaseInputDefinitionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DatabrewDatasetInputDatabaseInputDefinitionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DatabrewDatasetInputDatabaseInputDefinitionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DatabrewDatasetInputDatabaseInputDefinitionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DatabrewDatasetInputDatabaseInputDefinitionOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabrewDatasetInputDatabaseInputDefinitionOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DatabrewDatasetInputDatabaseInputDefinitionOutputReference) PutTempDirectory(value *DatabrewDatasetInputDatabaseInputDefinitionTempDirectory) {
	if err := d.validatePutTempDirectoryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTempDirectory",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabrewDatasetInputDatabaseInputDefinitionOutputReference) ResetDatabaseTableName() {
	_jsii_.InvokeVoid(
		d,
		"resetDatabaseTableName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabrewDatasetInputDatabaseInputDefinitionOutputReference) ResetGlueConnectionName() {
	_jsii_.InvokeVoid(
		d,
		"resetGlueConnectionName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabrewDatasetInputDatabaseInputDefinitionOutputReference) ResetQueryString() {
	_jsii_.InvokeVoid(
		d,
		"resetQueryString",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabrewDatasetInputDatabaseInputDefinitionOutputReference) ResetTempDirectory() {
	_jsii_.InvokeVoid(
		d,
		"resetTempDirectory",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabrewDatasetInputDatabaseInputDefinitionOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DatabrewDatasetInputDatabaseInputDefinitionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

