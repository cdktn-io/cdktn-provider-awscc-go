// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluetable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/gluetable/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GlueTableTableInputOutputReference interface {
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	Owner() *string
	SetOwner(val *string)
	OwnerInput() *string
	Parameters() *string
	SetParameters(val *string)
	ParametersInput() *string
	PartitionKeys() GlueTableTableInputPartitionKeysList
	PartitionKeysInput() interface{}
	Retention() *float64
	SetRetention(val *float64)
	RetentionInput() *float64
	StorageDescriptor() GlueTableTableInputStorageDescriptorOutputReference
	StorageDescriptorInput() interface{}
	TableType() *string
	SetTableType(val *string)
	TableTypeInput() *string
	TargetTable() GlueTableTableInputTargetTableOutputReference
	TargetTableInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	ViewDefinition() GlueTableTableInputViewDefinitionOutputReference
	ViewDefinitionInput() interface{}
	ViewExpandedText() *string
	SetViewExpandedText(val *string)
	ViewExpandedTextInput() *string
	ViewOriginalText() *string
	SetViewOriginalText(val *string)
	ViewOriginalTextInput() *string
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
	PutPartitionKeys(value interface{})
	PutStorageDescriptor(value *GlueTableTableInputStorageDescriptor)
	PutTargetTable(value *GlueTableTableInputTargetTable)
	PutViewDefinition(value *GlueTableTableInputViewDefinition)
	ResetDescription()
	ResetName()
	ResetOwner()
	ResetParameters()
	ResetPartitionKeys()
	ResetRetention()
	ResetStorageDescriptor()
	ResetTableType()
	ResetTargetTable()
	ResetViewDefinition()
	ResetViewExpandedText()
	ResetViewOriginalText()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GlueTableTableInputOutputReference
type jsiiProxy_GlueTableTableInputOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GlueTableTableInputOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputOutputReference) Owner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"owner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputOutputReference) OwnerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputOutputReference) Parameters() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputOutputReference) ParametersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputOutputReference) PartitionKeys() GlueTableTableInputPartitionKeysList {
	var returns GlueTableTableInputPartitionKeysList
	_jsii_.Get(
		j,
		"partitionKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputOutputReference) PartitionKeysInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"partitionKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputOutputReference) Retention() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retention",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputOutputReference) RetentionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputOutputReference) StorageDescriptor() GlueTableTableInputStorageDescriptorOutputReference {
	var returns GlueTableTableInputStorageDescriptorOutputReference
	_jsii_.Get(
		j,
		"storageDescriptor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputOutputReference) StorageDescriptorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageDescriptorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputOutputReference) TableType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputOutputReference) TableTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputOutputReference) TargetTable() GlueTableTableInputTargetTableOutputReference {
	var returns GlueTableTableInputTargetTableOutputReference
	_jsii_.Get(
		j,
		"targetTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputOutputReference) TargetTableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targetTableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputOutputReference) ViewDefinition() GlueTableTableInputViewDefinitionOutputReference {
	var returns GlueTableTableInputViewDefinitionOutputReference
	_jsii_.Get(
		j,
		"viewDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputOutputReference) ViewDefinitionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"viewDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputOutputReference) ViewExpandedText() *string {
	var returns *string
	_jsii_.Get(
		j,
		"viewExpandedText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputOutputReference) ViewExpandedTextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"viewExpandedTextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputOutputReference) ViewOriginalText() *string {
	var returns *string
	_jsii_.Get(
		j,
		"viewOriginalText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputOutputReference) ViewOriginalTextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"viewOriginalTextInput",
		&returns,
	)
	return returns
}


func NewGlueTableTableInputOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GlueTableTableInputOutputReference {
	_init_.Initialize()

	if err := validateNewGlueTableTableInputOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GlueTableTableInputOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.glueTable.GlueTableTableInputOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGlueTableTableInputOutputReference_Override(g GlueTableTableInputOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.glueTable.GlueTableTableInputOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GlueTableTableInputOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GlueTableTableInputOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GlueTableTableInputOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GlueTableTableInputOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GlueTableTableInputOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GlueTableTableInputOutputReference)SetOwner(val *string) {
	if err := j.validateSetOwnerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"owner",
		val,
	)
}

func (j *jsiiProxy_GlueTableTableInputOutputReference)SetParameters(val *string) {
	if err := j.validateSetParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_GlueTableTableInputOutputReference)SetRetention(val *float64) {
	if err := j.validateSetRetentionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retention",
		val,
	)
}

func (j *jsiiProxy_GlueTableTableInputOutputReference)SetTableType(val *string) {
	if err := j.validateSetTableTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tableType",
		val,
	)
}

func (j *jsiiProxy_GlueTableTableInputOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlueTableTableInputOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GlueTableTableInputOutputReference)SetViewExpandedText(val *string) {
	if err := j.validateSetViewExpandedTextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"viewExpandedText",
		val,
	)
}

func (j *jsiiProxy_GlueTableTableInputOutputReference)SetViewOriginalText(val *string) {
	if err := j.validateSetViewOriginalTextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"viewOriginalText",
		val,
	)
}

func (g *jsiiProxy_GlueTableTableInputOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueTableTableInputOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueTableTableInputOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueTableTableInputOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueTableTableInputOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueTableTableInputOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueTableTableInputOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueTableTableInputOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueTableTableInputOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueTableTableInputOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueTableTableInputOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueTableTableInputOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueTableTableInputOutputReference) PutPartitionKeys(value interface{}) {
	if err := g.validatePutPartitionKeysParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPartitionKeys",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueTableTableInputOutputReference) PutStorageDescriptor(value *GlueTableTableInputStorageDescriptor) {
	if err := g.validatePutStorageDescriptorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putStorageDescriptor",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueTableTableInputOutputReference) PutTargetTable(value *GlueTableTableInputTargetTable) {
	if err := g.validatePutTargetTableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTargetTable",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueTableTableInputOutputReference) PutViewDefinition(value *GlueTableTableInputViewDefinition) {
	if err := g.validatePutViewDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putViewDefinition",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueTableTableInputOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueTableTableInputOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		g,
		"resetName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueTableTableInputOutputReference) ResetOwner() {
	_jsii_.InvokeVoid(
		g,
		"resetOwner",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueTableTableInputOutputReference) ResetParameters() {
	_jsii_.InvokeVoid(
		g,
		"resetParameters",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueTableTableInputOutputReference) ResetPartitionKeys() {
	_jsii_.InvokeVoid(
		g,
		"resetPartitionKeys",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueTableTableInputOutputReference) ResetRetention() {
	_jsii_.InvokeVoid(
		g,
		"resetRetention",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueTableTableInputOutputReference) ResetStorageDescriptor() {
	_jsii_.InvokeVoid(
		g,
		"resetStorageDescriptor",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueTableTableInputOutputReference) ResetTableType() {
	_jsii_.InvokeVoid(
		g,
		"resetTableType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueTableTableInputOutputReference) ResetTargetTable() {
	_jsii_.InvokeVoid(
		g,
		"resetTargetTable",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueTableTableInputOutputReference) ResetViewDefinition() {
	_jsii_.InvokeVoid(
		g,
		"resetViewDefinition",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueTableTableInputOutputReference) ResetViewExpandedText() {
	_jsii_.InvokeVoid(
		g,
		"resetViewExpandedText",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueTableTableInputOutputReference) ResetViewOriginalText() {
	_jsii_.InvokeVoid(
		g,
		"resetViewOriginalText",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueTableTableInputOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := g.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueTableTableInputOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

