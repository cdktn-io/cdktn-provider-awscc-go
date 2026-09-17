// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kinesischannel

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/kinesischannel/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructOutputReference interface {
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
	CompressionType() *string
	SetCompressionType(val *string)
	CompressionTypeInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Namespace() *string
	SetNamespace(val *string)
	NamespaceInput() *string
	PartitionSpec() KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListPartitionSpecOutputReference
	PartitionSpecInput() interface{}
	TableBucketArn() *string
	SetTableBucketArn(val *string)
	TableBucketArnInput() *string
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
	PutPartitionSpec(value *KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListPartitionSpec)
	ResetCompressionType()
	ResetNamespace()
	ResetPartitionSpec()
	ResetTableBucketArn()
	ResetTableName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructOutputReference
type jsiiProxy_KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructOutputReference) CompressionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructOutputReference) CompressionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructOutputReference) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructOutputReference) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructOutputReference) PartitionSpec() KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListPartitionSpecOutputReference {
	var returns KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListPartitionSpecOutputReference
	_jsii_.Get(
		j,
		"partitionSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructOutputReference) PartitionSpecInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"partitionSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructOutputReference) TableBucketArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableBucketArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructOutputReference) TableBucketArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableBucketArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructOutputReference) TableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructOutputReference) TableNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructOutputReference {
	_init_.Initialize()

	if err := validateNewKinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.kinesisChannel.KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewKinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructOutputReference_Override(k KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.kinesisChannel.KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		k,
	)
}

func (j *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructOutputReference)SetCompressionType(val *string) {
	if err := j.validateSetCompressionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compressionType",
		val,
	)
}

func (j *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructOutputReference)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructOutputReference)SetTableBucketArn(val *string) {
	if err := j.validateSetTableBucketArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tableBucketArn",
		val,
	)
}

func (j *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructOutputReference)SetTableName(val *string) {
	if err := j.validateSetTableNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tableName",
		val,
	)
}

func (j *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (k *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (k *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (k *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (k *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (k *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (k *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (k *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (k *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (k *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (k *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructOutputReference) PutPartitionSpec(value *KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListPartitionSpec) {
	if err := k.validatePutPartitionSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putPartitionSpec",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructOutputReference) ResetCompressionType() {
	_jsii_.InvokeVoid(
		k,
		"resetCompressionType",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructOutputReference) ResetNamespace() {
	_jsii_.InvokeVoid(
		k,
		"resetNamespace",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructOutputReference) ResetPartitionSpec() {
	_jsii_.InvokeVoid(
		k,
		"resetPartitionSpec",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructOutputReference) ResetTableBucketArn() {
	_jsii_.InvokeVoid(
		k,
		"resetTableBucketArn",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructOutputReference) ResetTableName() {
	_jsii_.InvokeVoid(
		k,
		"resetTableName",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (k *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

