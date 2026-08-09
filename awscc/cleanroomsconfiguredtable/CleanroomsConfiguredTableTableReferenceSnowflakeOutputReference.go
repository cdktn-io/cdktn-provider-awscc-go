// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cleanroomsconfiguredtable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/cleanroomsconfiguredtable/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CleanroomsConfiguredTableTableReferenceSnowflakeOutputReference interface {
	cdktn.ComplexObject
	AccountIdentifier() *string
	SetAccountIdentifier(val *string)
	AccountIdentifierInput() *string
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
	DatabaseName() *string
	SetDatabaseName(val *string)
	DatabaseNameInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	SchemaName() *string
	SetSchemaName(val *string)
	SchemaNameInput() *string
	SecretArn() *string
	SetSecretArn(val *string)
	SecretArnInput() *string
	TableName() *string
	SetTableName(val *string)
	TableNameInput() *string
	TableSchema() CleanroomsConfiguredTableTableReferenceSnowflakeTableSchemaOutputReference
	TableSchemaInput() interface{}
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
	PutTableSchema(value *CleanroomsConfiguredTableTableReferenceSnowflakeTableSchema)
	ResetAccountIdentifier()
	ResetDatabaseName()
	ResetSchemaName()
	ResetSecretArn()
	ResetTableName()
	ResetTableSchema()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CleanroomsConfiguredTableTableReferenceSnowflakeOutputReference
type jsiiProxy_CleanroomsConfiguredTableTableReferenceSnowflakeOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_CleanroomsConfiguredTableTableReferenceSnowflakeOutputReference) AccountIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsConfiguredTableTableReferenceSnowflakeOutputReference) AccountIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsConfiguredTableTableReferenceSnowflakeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsConfiguredTableTableReferenceSnowflakeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsConfiguredTableTableReferenceSnowflakeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsConfiguredTableTableReferenceSnowflakeOutputReference) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsConfiguredTableTableReferenceSnowflakeOutputReference) DatabaseNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsConfiguredTableTableReferenceSnowflakeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsConfiguredTableTableReferenceSnowflakeOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsConfiguredTableTableReferenceSnowflakeOutputReference) SchemaName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsConfiguredTableTableReferenceSnowflakeOutputReference) SchemaNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsConfiguredTableTableReferenceSnowflakeOutputReference) SecretArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsConfiguredTableTableReferenceSnowflakeOutputReference) SecretArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsConfiguredTableTableReferenceSnowflakeOutputReference) TableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsConfiguredTableTableReferenceSnowflakeOutputReference) TableNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsConfiguredTableTableReferenceSnowflakeOutputReference) TableSchema() CleanroomsConfiguredTableTableReferenceSnowflakeTableSchemaOutputReference {
	var returns CleanroomsConfiguredTableTableReferenceSnowflakeTableSchemaOutputReference
	_jsii_.Get(
		j,
		"tableSchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsConfiguredTableTableReferenceSnowflakeOutputReference) TableSchemaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tableSchemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsConfiguredTableTableReferenceSnowflakeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsConfiguredTableTableReferenceSnowflakeOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCleanroomsConfiguredTableTableReferenceSnowflakeOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) CleanroomsConfiguredTableTableReferenceSnowflakeOutputReference {
	_init_.Initialize()

	if err := validateNewCleanroomsConfiguredTableTableReferenceSnowflakeOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CleanroomsConfiguredTableTableReferenceSnowflakeOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.cleanroomsConfiguredTable.CleanroomsConfiguredTableTableReferenceSnowflakeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCleanroomsConfiguredTableTableReferenceSnowflakeOutputReference_Override(c CleanroomsConfiguredTableTableReferenceSnowflakeOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.cleanroomsConfiguredTable.CleanroomsConfiguredTableTableReferenceSnowflakeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CleanroomsConfiguredTableTableReferenceSnowflakeOutputReference)SetAccountIdentifier(val *string) {
	if err := j.validateSetAccountIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountIdentifier",
		val,
	)
}

func (j *jsiiProxy_CleanroomsConfiguredTableTableReferenceSnowflakeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CleanroomsConfiguredTableTableReferenceSnowflakeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CleanroomsConfiguredTableTableReferenceSnowflakeOutputReference)SetDatabaseName(val *string) {
	if err := j.validateSetDatabaseNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseName",
		val,
	)
}

func (j *jsiiProxy_CleanroomsConfiguredTableTableReferenceSnowflakeOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CleanroomsConfiguredTableTableReferenceSnowflakeOutputReference)SetSchemaName(val *string) {
	if err := j.validateSetSchemaNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schemaName",
		val,
	)
}

func (j *jsiiProxy_CleanroomsConfiguredTableTableReferenceSnowflakeOutputReference)SetSecretArn(val *string) {
	if err := j.validateSetSecretArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretArn",
		val,
	)
}

func (j *jsiiProxy_CleanroomsConfiguredTableTableReferenceSnowflakeOutputReference)SetTableName(val *string) {
	if err := j.validateSetTableNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tableName",
		val,
	)
}

func (j *jsiiProxy_CleanroomsConfiguredTableTableReferenceSnowflakeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CleanroomsConfiguredTableTableReferenceSnowflakeOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CleanroomsConfiguredTableTableReferenceSnowflakeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsConfiguredTableTableReferenceSnowflakeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsConfiguredTableTableReferenceSnowflakeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsConfiguredTableTableReferenceSnowflakeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsConfiguredTableTableReferenceSnowflakeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsConfiguredTableTableReferenceSnowflakeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsConfiguredTableTableReferenceSnowflakeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsConfiguredTableTableReferenceSnowflakeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsConfiguredTableTableReferenceSnowflakeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsConfiguredTableTableReferenceSnowflakeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsConfiguredTableTableReferenceSnowflakeOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsConfiguredTableTableReferenceSnowflakeOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsConfiguredTableTableReferenceSnowflakeOutputReference) PutTableSchema(value *CleanroomsConfiguredTableTableReferenceSnowflakeTableSchema) {
	if err := c.validatePutTableSchemaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTableSchema",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CleanroomsConfiguredTableTableReferenceSnowflakeOutputReference) ResetAccountIdentifier() {
	_jsii_.InvokeVoid(
		c,
		"resetAccountIdentifier",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanroomsConfiguredTableTableReferenceSnowflakeOutputReference) ResetDatabaseName() {
	_jsii_.InvokeVoid(
		c,
		"resetDatabaseName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanroomsConfiguredTableTableReferenceSnowflakeOutputReference) ResetSchemaName() {
	_jsii_.InvokeVoid(
		c,
		"resetSchemaName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanroomsConfiguredTableTableReferenceSnowflakeOutputReference) ResetSecretArn() {
	_jsii_.InvokeVoid(
		c,
		"resetSecretArn",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanroomsConfiguredTableTableReferenceSnowflakeOutputReference) ResetTableName() {
	_jsii_.InvokeVoid(
		c,
		"resetTableName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanroomsConfiguredTableTableReferenceSnowflakeOutputReference) ResetTableSchema() {
	_jsii_.InvokeVoid(
		c,
		"resetTableSchema",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanroomsConfiguredTableTableReferenceSnowflakeOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := c.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsConfiguredTableTableReferenceSnowflakeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

