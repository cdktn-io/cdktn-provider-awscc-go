// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kinesisfirehosedeliverystream

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/kinesisfirehosedeliverystream/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationDestinationTableConfigurationListStructOutputReference interface {
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
	DestinationDatabaseName() *string
	SetDestinationDatabaseName(val *string)
	DestinationDatabaseNameInput() *string
	DestinationTableName() *string
	SetDestinationTableName(val *string)
	DestinationTableNameInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PartitionSpec() KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationDestinationTableConfigurationListPartitionSpecOutputReference
	PartitionSpecInput() interface{}
	S3ErrorOutputPrefix() *string
	SetS3ErrorOutputPrefix(val *string)
	S3ErrorOutputPrefixInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UniqueKeys() *[]*string
	SetUniqueKeys(val *[]*string)
	UniqueKeysInput() *[]*string
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
	PutPartitionSpec(value *KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationDestinationTableConfigurationListPartitionSpec)
	ResetDestinationDatabaseName()
	ResetDestinationTableName()
	ResetPartitionSpec()
	ResetS3ErrorOutputPrefix()
	ResetUniqueKeys()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationDestinationTableConfigurationListStructOutputReference
type jsiiProxy_KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationDestinationTableConfigurationListStructOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationDestinationTableConfigurationListStructOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationDestinationTableConfigurationListStructOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationDestinationTableConfigurationListStructOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationDestinationTableConfigurationListStructOutputReference) DestinationDatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationDatabaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationDestinationTableConfigurationListStructOutputReference) DestinationDatabaseNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationDatabaseNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationDestinationTableConfigurationListStructOutputReference) DestinationTableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationTableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationDestinationTableConfigurationListStructOutputReference) DestinationTableNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationTableNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationDestinationTableConfigurationListStructOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationDestinationTableConfigurationListStructOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationDestinationTableConfigurationListStructOutputReference) PartitionSpec() KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationDestinationTableConfigurationListPartitionSpecOutputReference {
	var returns KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationDestinationTableConfigurationListPartitionSpecOutputReference
	_jsii_.Get(
		j,
		"partitionSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationDestinationTableConfigurationListStructOutputReference) PartitionSpecInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"partitionSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationDestinationTableConfigurationListStructOutputReference) S3ErrorOutputPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3ErrorOutputPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationDestinationTableConfigurationListStructOutputReference) S3ErrorOutputPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3ErrorOutputPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationDestinationTableConfigurationListStructOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationDestinationTableConfigurationListStructOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationDestinationTableConfigurationListStructOutputReference) UniqueKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"uniqueKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationDestinationTableConfigurationListStructOutputReference) UniqueKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"uniqueKeysInput",
		&returns,
	)
	return returns
}


func NewKinesisfirehoseDeliveryStreamIcebergDestinationConfigurationDestinationTableConfigurationListStructOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationDestinationTableConfigurationListStructOutputReference {
	_init_.Initialize()

	if err := validateNewKinesisfirehoseDeliveryStreamIcebergDestinationConfigurationDestinationTableConfigurationListStructOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationDestinationTableConfigurationListStructOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationDestinationTableConfigurationListStructOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewKinesisfirehoseDeliveryStreamIcebergDestinationConfigurationDestinationTableConfigurationListStructOutputReference_Override(k KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationDestinationTableConfigurationListStructOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationDestinationTableConfigurationListStructOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		k,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationDestinationTableConfigurationListStructOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationDestinationTableConfigurationListStructOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationDestinationTableConfigurationListStructOutputReference)SetDestinationDatabaseName(val *string) {
	if err := j.validateSetDestinationDatabaseNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationDatabaseName",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationDestinationTableConfigurationListStructOutputReference)SetDestinationTableName(val *string) {
	if err := j.validateSetDestinationTableNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationTableName",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationDestinationTableConfigurationListStructOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationDestinationTableConfigurationListStructOutputReference)SetS3ErrorOutputPrefix(val *string) {
	if err := j.validateSetS3ErrorOutputPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3ErrorOutputPrefix",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationDestinationTableConfigurationListStructOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationDestinationTableConfigurationListStructOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationDestinationTableConfigurationListStructOutputReference)SetUniqueKeys(val *[]*string) {
	if err := j.validateSetUniqueKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uniqueKeys",
		val,
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationDestinationTableConfigurationListStructOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationDestinationTableConfigurationListStructOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := k.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationDestinationTableConfigurationListStructOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := k.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationDestinationTableConfigurationListStructOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := k.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationDestinationTableConfigurationListStructOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := k.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationDestinationTableConfigurationListStructOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := k.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationDestinationTableConfigurationListStructOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := k.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationDestinationTableConfigurationListStructOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := k.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationDestinationTableConfigurationListStructOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := k.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationDestinationTableConfigurationListStructOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := k.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationDestinationTableConfigurationListStructOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationDestinationTableConfigurationListStructOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := k.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationDestinationTableConfigurationListStructOutputReference) PutPartitionSpec(value *KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationDestinationTableConfigurationListPartitionSpec) {
	if err := k.validatePutPartitionSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putPartitionSpec",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationDestinationTableConfigurationListStructOutputReference) ResetDestinationDatabaseName() {
	_jsii_.InvokeVoid(
		k,
		"resetDestinationDatabaseName",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationDestinationTableConfigurationListStructOutputReference) ResetDestinationTableName() {
	_jsii_.InvokeVoid(
		k,
		"resetDestinationTableName",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationDestinationTableConfigurationListStructOutputReference) ResetPartitionSpec() {
	_jsii_.InvokeVoid(
		k,
		"resetPartitionSpec",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationDestinationTableConfigurationListStructOutputReference) ResetS3ErrorOutputPrefix() {
	_jsii_.InvokeVoid(
		k,
		"resetS3ErrorOutputPrefix",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationDestinationTableConfigurationListStructOutputReference) ResetUniqueKeys() {
	_jsii_.InvokeVoid(
		k,
		"resetUniqueKeys",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationDestinationTableConfigurationListStructOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := k.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationDestinationTableConfigurationListStructOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

