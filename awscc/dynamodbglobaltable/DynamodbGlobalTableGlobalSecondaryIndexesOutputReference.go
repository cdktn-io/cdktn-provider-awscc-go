// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dynamodbglobaltable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dynamodbglobaltable/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DynamodbGlobalTableGlobalSecondaryIndexesOutputReference interface {
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
	// Experimental.
	Fqn() *string
	IndexName() *string
	SetIndexName(val *string)
	IndexNameInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	KeySchema() DynamodbGlobalTableGlobalSecondaryIndexesKeySchemaList
	KeySchemaInput() interface{}
	Projection() DynamodbGlobalTableGlobalSecondaryIndexesProjectionOutputReference
	ProjectionInput() interface{}
	ReadOnDemandThroughputSettings() DynamodbGlobalTableGlobalSecondaryIndexesReadOnDemandThroughputSettingsOutputReference
	ReadOnDemandThroughputSettingsInput() interface{}
	ReadProvisionedThroughputSettings() DynamodbGlobalTableGlobalSecondaryIndexesReadProvisionedThroughputSettingsOutputReference
	ReadProvisionedThroughputSettingsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	WarmThroughput() DynamodbGlobalTableGlobalSecondaryIndexesWarmThroughputOutputReference
	WarmThroughputInput() interface{}
	WriteOnDemandThroughputSettings() DynamodbGlobalTableGlobalSecondaryIndexesWriteOnDemandThroughputSettingsOutputReference
	WriteOnDemandThroughputSettingsInput() interface{}
	WriteProvisionedThroughputSettings() DynamodbGlobalTableGlobalSecondaryIndexesWriteProvisionedThroughputSettingsOutputReference
	WriteProvisionedThroughputSettingsInput() interface{}
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
	PutKeySchema(value interface{})
	PutProjection(value *DynamodbGlobalTableGlobalSecondaryIndexesProjection)
	PutReadOnDemandThroughputSettings(value *DynamodbGlobalTableGlobalSecondaryIndexesReadOnDemandThroughputSettings)
	PutReadProvisionedThroughputSettings(value *DynamodbGlobalTableGlobalSecondaryIndexesReadProvisionedThroughputSettings)
	PutWarmThroughput(value *DynamodbGlobalTableGlobalSecondaryIndexesWarmThroughput)
	PutWriteOnDemandThroughputSettings(value *DynamodbGlobalTableGlobalSecondaryIndexesWriteOnDemandThroughputSettings)
	PutWriteProvisionedThroughputSettings(value *DynamodbGlobalTableGlobalSecondaryIndexesWriteProvisionedThroughputSettings)
	ResetIndexName()
	ResetKeySchema()
	ResetProjection()
	ResetReadOnDemandThroughputSettings()
	ResetReadProvisionedThroughputSettings()
	ResetWarmThroughput()
	ResetWriteOnDemandThroughputSettings()
	ResetWriteProvisionedThroughputSettings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DynamodbGlobalTableGlobalSecondaryIndexesOutputReference
type jsiiProxy_DynamodbGlobalTableGlobalSecondaryIndexesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DynamodbGlobalTableGlobalSecondaryIndexesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTableGlobalSecondaryIndexesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTableGlobalSecondaryIndexesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTableGlobalSecondaryIndexesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTableGlobalSecondaryIndexesOutputReference) IndexName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTableGlobalSecondaryIndexesOutputReference) IndexNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTableGlobalSecondaryIndexesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTableGlobalSecondaryIndexesOutputReference) KeySchema() DynamodbGlobalTableGlobalSecondaryIndexesKeySchemaList {
	var returns DynamodbGlobalTableGlobalSecondaryIndexesKeySchemaList
	_jsii_.Get(
		j,
		"keySchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTableGlobalSecondaryIndexesOutputReference) KeySchemaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keySchemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTableGlobalSecondaryIndexesOutputReference) Projection() DynamodbGlobalTableGlobalSecondaryIndexesProjectionOutputReference {
	var returns DynamodbGlobalTableGlobalSecondaryIndexesProjectionOutputReference
	_jsii_.Get(
		j,
		"projection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTableGlobalSecondaryIndexesOutputReference) ProjectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"projectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTableGlobalSecondaryIndexesOutputReference) ReadOnDemandThroughputSettings() DynamodbGlobalTableGlobalSecondaryIndexesReadOnDemandThroughputSettingsOutputReference {
	var returns DynamodbGlobalTableGlobalSecondaryIndexesReadOnDemandThroughputSettingsOutputReference
	_jsii_.Get(
		j,
		"readOnDemandThroughputSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTableGlobalSecondaryIndexesOutputReference) ReadOnDemandThroughputSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readOnDemandThroughputSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTableGlobalSecondaryIndexesOutputReference) ReadProvisionedThroughputSettings() DynamodbGlobalTableGlobalSecondaryIndexesReadProvisionedThroughputSettingsOutputReference {
	var returns DynamodbGlobalTableGlobalSecondaryIndexesReadProvisionedThroughputSettingsOutputReference
	_jsii_.Get(
		j,
		"readProvisionedThroughputSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTableGlobalSecondaryIndexesOutputReference) ReadProvisionedThroughputSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readProvisionedThroughputSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTableGlobalSecondaryIndexesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTableGlobalSecondaryIndexesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTableGlobalSecondaryIndexesOutputReference) WarmThroughput() DynamodbGlobalTableGlobalSecondaryIndexesWarmThroughputOutputReference {
	var returns DynamodbGlobalTableGlobalSecondaryIndexesWarmThroughputOutputReference
	_jsii_.Get(
		j,
		"warmThroughput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTableGlobalSecondaryIndexesOutputReference) WarmThroughputInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"warmThroughputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTableGlobalSecondaryIndexesOutputReference) WriteOnDemandThroughputSettings() DynamodbGlobalTableGlobalSecondaryIndexesWriteOnDemandThroughputSettingsOutputReference {
	var returns DynamodbGlobalTableGlobalSecondaryIndexesWriteOnDemandThroughputSettingsOutputReference
	_jsii_.Get(
		j,
		"writeOnDemandThroughputSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTableGlobalSecondaryIndexesOutputReference) WriteOnDemandThroughputSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"writeOnDemandThroughputSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTableGlobalSecondaryIndexesOutputReference) WriteProvisionedThroughputSettings() DynamodbGlobalTableGlobalSecondaryIndexesWriteProvisionedThroughputSettingsOutputReference {
	var returns DynamodbGlobalTableGlobalSecondaryIndexesWriteProvisionedThroughputSettingsOutputReference
	_jsii_.Get(
		j,
		"writeProvisionedThroughputSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTableGlobalSecondaryIndexesOutputReference) WriteProvisionedThroughputSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"writeProvisionedThroughputSettingsInput",
		&returns,
	)
	return returns
}


func NewDynamodbGlobalTableGlobalSecondaryIndexesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DynamodbGlobalTableGlobalSecondaryIndexesOutputReference {
	_init_.Initialize()

	if err := validateNewDynamodbGlobalTableGlobalSecondaryIndexesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DynamodbGlobalTableGlobalSecondaryIndexesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dynamodbGlobalTable.DynamodbGlobalTableGlobalSecondaryIndexesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDynamodbGlobalTableGlobalSecondaryIndexesOutputReference_Override(d DynamodbGlobalTableGlobalSecondaryIndexesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dynamodbGlobalTable.DynamodbGlobalTableGlobalSecondaryIndexesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DynamodbGlobalTableGlobalSecondaryIndexesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DynamodbGlobalTableGlobalSecondaryIndexesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DynamodbGlobalTableGlobalSecondaryIndexesOutputReference)SetIndexName(val *string) {
	if err := j.validateSetIndexNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"indexName",
		val,
	)
}

func (j *jsiiProxy_DynamodbGlobalTableGlobalSecondaryIndexesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DynamodbGlobalTableGlobalSecondaryIndexesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DynamodbGlobalTableGlobalSecondaryIndexesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DynamodbGlobalTableGlobalSecondaryIndexesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamodbGlobalTableGlobalSecondaryIndexesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DynamodbGlobalTableGlobalSecondaryIndexesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DynamodbGlobalTableGlobalSecondaryIndexesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DynamodbGlobalTableGlobalSecondaryIndexesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DynamodbGlobalTableGlobalSecondaryIndexesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DynamodbGlobalTableGlobalSecondaryIndexesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DynamodbGlobalTableGlobalSecondaryIndexesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DynamodbGlobalTableGlobalSecondaryIndexesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DynamodbGlobalTableGlobalSecondaryIndexesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DynamodbGlobalTableGlobalSecondaryIndexesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamodbGlobalTableGlobalSecondaryIndexesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DynamodbGlobalTableGlobalSecondaryIndexesOutputReference) PutKeySchema(value interface{}) {
	if err := d.validatePutKeySchemaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putKeySchema",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbGlobalTableGlobalSecondaryIndexesOutputReference) PutProjection(value *DynamodbGlobalTableGlobalSecondaryIndexesProjection) {
	if err := d.validatePutProjectionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putProjection",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbGlobalTableGlobalSecondaryIndexesOutputReference) PutReadOnDemandThroughputSettings(value *DynamodbGlobalTableGlobalSecondaryIndexesReadOnDemandThroughputSettings) {
	if err := d.validatePutReadOnDemandThroughputSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putReadOnDemandThroughputSettings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbGlobalTableGlobalSecondaryIndexesOutputReference) PutReadProvisionedThroughputSettings(value *DynamodbGlobalTableGlobalSecondaryIndexesReadProvisionedThroughputSettings) {
	if err := d.validatePutReadProvisionedThroughputSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putReadProvisionedThroughputSettings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbGlobalTableGlobalSecondaryIndexesOutputReference) PutWarmThroughput(value *DynamodbGlobalTableGlobalSecondaryIndexesWarmThroughput) {
	if err := d.validatePutWarmThroughputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putWarmThroughput",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbGlobalTableGlobalSecondaryIndexesOutputReference) PutWriteOnDemandThroughputSettings(value *DynamodbGlobalTableGlobalSecondaryIndexesWriteOnDemandThroughputSettings) {
	if err := d.validatePutWriteOnDemandThroughputSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putWriteOnDemandThroughputSettings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbGlobalTableGlobalSecondaryIndexesOutputReference) PutWriteProvisionedThroughputSettings(value *DynamodbGlobalTableGlobalSecondaryIndexesWriteProvisionedThroughputSettings) {
	if err := d.validatePutWriteProvisionedThroughputSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putWriteProvisionedThroughputSettings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbGlobalTableGlobalSecondaryIndexesOutputReference) ResetIndexName() {
	_jsii_.InvokeVoid(
		d,
		"resetIndexName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbGlobalTableGlobalSecondaryIndexesOutputReference) ResetKeySchema() {
	_jsii_.InvokeVoid(
		d,
		"resetKeySchema",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbGlobalTableGlobalSecondaryIndexesOutputReference) ResetProjection() {
	_jsii_.InvokeVoid(
		d,
		"resetProjection",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbGlobalTableGlobalSecondaryIndexesOutputReference) ResetReadOnDemandThroughputSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetReadOnDemandThroughputSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbGlobalTableGlobalSecondaryIndexesOutputReference) ResetReadProvisionedThroughputSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetReadProvisionedThroughputSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbGlobalTableGlobalSecondaryIndexesOutputReference) ResetWarmThroughput() {
	_jsii_.InvokeVoid(
		d,
		"resetWarmThroughput",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbGlobalTableGlobalSecondaryIndexesOutputReference) ResetWriteOnDemandThroughputSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetWriteOnDemandThroughputSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbGlobalTableGlobalSecondaryIndexesOutputReference) ResetWriteProvisionedThroughputSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetWriteProvisionedThroughputSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbGlobalTableGlobalSecondaryIndexesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DynamodbGlobalTableGlobalSecondaryIndexesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

