// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appsyncdatasource

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/appsyncdatasource/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AppsyncDataSourceDynamoDbConfigOutputReference interface {
	cdktn.ComplexObject
	AwsRegion() *string
	SetAwsRegion(val *string)
	AwsRegionInput() *string
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
	DeltaSyncConfig() AppsyncDataSourceDynamoDbConfigDeltaSyncConfigOutputReference
	DeltaSyncConfigInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	TableName() *string
	SetTableName(val *string)
	TableNameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UseCallerCredentials() interface{}
	SetUseCallerCredentials(val interface{})
	UseCallerCredentialsInput() interface{}
	Versioned() interface{}
	SetVersioned(val interface{})
	VersionedInput() interface{}
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
	PutDeltaSyncConfig(value *AppsyncDataSourceDynamoDbConfigDeltaSyncConfig)
	ResetAwsRegion()
	ResetDeltaSyncConfig()
	ResetTableName()
	ResetUseCallerCredentials()
	ResetVersioned()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppsyncDataSourceDynamoDbConfigOutputReference
type jsiiProxy_AppsyncDataSourceDynamoDbConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AppsyncDataSourceDynamoDbConfigOutputReference) AwsRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDataSourceDynamoDbConfigOutputReference) AwsRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDataSourceDynamoDbConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDataSourceDynamoDbConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDataSourceDynamoDbConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDataSourceDynamoDbConfigOutputReference) DeltaSyncConfig() AppsyncDataSourceDynamoDbConfigDeltaSyncConfigOutputReference {
	var returns AppsyncDataSourceDynamoDbConfigDeltaSyncConfigOutputReference
	_jsii_.Get(
		j,
		"deltaSyncConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDataSourceDynamoDbConfigOutputReference) DeltaSyncConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deltaSyncConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDataSourceDynamoDbConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDataSourceDynamoDbConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDataSourceDynamoDbConfigOutputReference) TableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDataSourceDynamoDbConfigOutputReference) TableNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDataSourceDynamoDbConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDataSourceDynamoDbConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDataSourceDynamoDbConfigOutputReference) UseCallerCredentials() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useCallerCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDataSourceDynamoDbConfigOutputReference) UseCallerCredentialsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useCallerCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDataSourceDynamoDbConfigOutputReference) Versioned() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"versioned",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDataSourceDynamoDbConfigOutputReference) VersionedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"versionedInput",
		&returns,
	)
	return returns
}


func NewAppsyncDataSourceDynamoDbConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AppsyncDataSourceDynamoDbConfigOutputReference {
	_init_.Initialize()

	if err := validateNewAppsyncDataSourceDynamoDbConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppsyncDataSourceDynamoDbConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.appsyncDataSource.AppsyncDataSourceDynamoDbConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppsyncDataSourceDynamoDbConfigOutputReference_Override(a AppsyncDataSourceDynamoDbConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.appsyncDataSource.AppsyncDataSourceDynamoDbConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppsyncDataSourceDynamoDbConfigOutputReference)SetAwsRegion(val *string) {
	if err := j.validateSetAwsRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"awsRegion",
		val,
	)
}

func (j *jsiiProxy_AppsyncDataSourceDynamoDbConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppsyncDataSourceDynamoDbConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppsyncDataSourceDynamoDbConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppsyncDataSourceDynamoDbConfigOutputReference)SetTableName(val *string) {
	if err := j.validateSetTableNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tableName",
		val,
	)
}

func (j *jsiiProxy_AppsyncDataSourceDynamoDbConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppsyncDataSourceDynamoDbConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AppsyncDataSourceDynamoDbConfigOutputReference)SetUseCallerCredentials(val interface{}) {
	if err := j.validateSetUseCallerCredentialsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useCallerCredentials",
		val,
	)
}

func (j *jsiiProxy_AppsyncDataSourceDynamoDbConfigOutputReference)SetVersioned(val interface{}) {
	if err := j.validateSetVersionedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"versioned",
		val,
	)
}

func (a *jsiiProxy_AppsyncDataSourceDynamoDbConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncDataSourceDynamoDbConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncDataSourceDynamoDbConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncDataSourceDynamoDbConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncDataSourceDynamoDbConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncDataSourceDynamoDbConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncDataSourceDynamoDbConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncDataSourceDynamoDbConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncDataSourceDynamoDbConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncDataSourceDynamoDbConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncDataSourceDynamoDbConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncDataSourceDynamoDbConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncDataSourceDynamoDbConfigOutputReference) PutDeltaSyncConfig(value *AppsyncDataSourceDynamoDbConfigDeltaSyncConfig) {
	if err := a.validatePutDeltaSyncConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDeltaSyncConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsyncDataSourceDynamoDbConfigOutputReference) ResetAwsRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetAwsRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncDataSourceDynamoDbConfigOutputReference) ResetDeltaSyncConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetDeltaSyncConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncDataSourceDynamoDbConfigOutputReference) ResetTableName() {
	_jsii_.InvokeVoid(
		a,
		"resetTableName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncDataSourceDynamoDbConfigOutputReference) ResetUseCallerCredentials() {
	_jsii_.InvokeVoid(
		a,
		"resetUseCallerCredentials",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncDataSourceDynamoDbConfigOutputReference) ResetVersioned() {
	_jsii_.InvokeVoid(
		a,
		"resetVersioned",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncDataSourceDynamoDbConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncDataSourceDynamoDbConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

