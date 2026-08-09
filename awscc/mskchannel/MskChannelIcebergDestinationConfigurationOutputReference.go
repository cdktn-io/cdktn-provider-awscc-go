// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mskchannel

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/mskchannel/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MskChannelIcebergDestinationConfigurationOutputReference interface {
	cdktn.ComplexObject
	AppendOnly() interface{}
	SetAppendOnly(val interface{})
	AppendOnlyInput() interface{}
	Catalog() MskChannelIcebergDestinationConfigurationCatalogOutputReference
	CatalogInput() interface{}
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
	DataFreshnessInSeconds() *float64
	SetDataFreshnessInSeconds(val *float64)
	DataFreshnessInSecondsInput() *float64
	DeadLetterQueueS3() MskChannelIcebergDestinationConfigurationDeadLetterQueueS3OutputReference
	DeadLetterQueueS3Input() interface{}
	DestinationTableList() MskChannelIcebergDestinationConfigurationDestinationTableListStructList
	DestinationTableListInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	SchemaEvolution() MskChannelIcebergDestinationConfigurationSchemaEvolutionOutputReference
	SchemaEvolutionInput() interface{}
	ServiceExecutionRoleArn() *string
	SetServiceExecutionRoleArn(val *string)
	ServiceExecutionRoleArnInput() *string
	TableCreation() MskChannelIcebergDestinationConfigurationTableCreationOutputReference
	TableCreationInput() interface{}
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
	PutCatalog(value *MskChannelIcebergDestinationConfigurationCatalog)
	PutDeadLetterQueueS3(value *MskChannelIcebergDestinationConfigurationDeadLetterQueueS3)
	PutDestinationTableList(value interface{})
	PutSchemaEvolution(value *MskChannelIcebergDestinationConfigurationSchemaEvolution)
	PutTableCreation(value *MskChannelIcebergDestinationConfigurationTableCreation)
	ResetAppendOnly()
	ResetCatalog()
	ResetCompressionType()
	ResetDataFreshnessInSeconds()
	ResetDeadLetterQueueS3()
	ResetDestinationTableList()
	ResetSchemaEvolution()
	ResetServiceExecutionRoleArn()
	ResetTableCreation()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MskChannelIcebergDestinationConfigurationOutputReference
type jsiiProxy_MskChannelIcebergDestinationConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_MskChannelIcebergDestinationConfigurationOutputReference) AppendOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"appendOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskChannelIcebergDestinationConfigurationOutputReference) AppendOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"appendOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskChannelIcebergDestinationConfigurationOutputReference) Catalog() MskChannelIcebergDestinationConfigurationCatalogOutputReference {
	var returns MskChannelIcebergDestinationConfigurationCatalogOutputReference
	_jsii_.Get(
		j,
		"catalog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskChannelIcebergDestinationConfigurationOutputReference) CatalogInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"catalogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskChannelIcebergDestinationConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskChannelIcebergDestinationConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskChannelIcebergDestinationConfigurationOutputReference) CompressionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskChannelIcebergDestinationConfigurationOutputReference) CompressionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskChannelIcebergDestinationConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskChannelIcebergDestinationConfigurationOutputReference) DataFreshnessInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataFreshnessInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskChannelIcebergDestinationConfigurationOutputReference) DataFreshnessInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataFreshnessInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskChannelIcebergDestinationConfigurationOutputReference) DeadLetterQueueS3() MskChannelIcebergDestinationConfigurationDeadLetterQueueS3OutputReference {
	var returns MskChannelIcebergDestinationConfigurationDeadLetterQueueS3OutputReference
	_jsii_.Get(
		j,
		"deadLetterQueueS3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskChannelIcebergDestinationConfigurationOutputReference) DeadLetterQueueS3Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deadLetterQueueS3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskChannelIcebergDestinationConfigurationOutputReference) DestinationTableList() MskChannelIcebergDestinationConfigurationDestinationTableListStructList {
	var returns MskChannelIcebergDestinationConfigurationDestinationTableListStructList
	_jsii_.Get(
		j,
		"destinationTableList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskChannelIcebergDestinationConfigurationOutputReference) DestinationTableListInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"destinationTableListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskChannelIcebergDestinationConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskChannelIcebergDestinationConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskChannelIcebergDestinationConfigurationOutputReference) SchemaEvolution() MskChannelIcebergDestinationConfigurationSchemaEvolutionOutputReference {
	var returns MskChannelIcebergDestinationConfigurationSchemaEvolutionOutputReference
	_jsii_.Get(
		j,
		"schemaEvolution",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskChannelIcebergDestinationConfigurationOutputReference) SchemaEvolutionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"schemaEvolutionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskChannelIcebergDestinationConfigurationOutputReference) ServiceExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceExecutionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskChannelIcebergDestinationConfigurationOutputReference) ServiceExecutionRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceExecutionRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskChannelIcebergDestinationConfigurationOutputReference) TableCreation() MskChannelIcebergDestinationConfigurationTableCreationOutputReference {
	var returns MskChannelIcebergDestinationConfigurationTableCreationOutputReference
	_jsii_.Get(
		j,
		"tableCreation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskChannelIcebergDestinationConfigurationOutputReference) TableCreationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tableCreationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskChannelIcebergDestinationConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskChannelIcebergDestinationConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMskChannelIcebergDestinationConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) MskChannelIcebergDestinationConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewMskChannelIcebergDestinationConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MskChannelIcebergDestinationConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.mskChannel.MskChannelIcebergDestinationConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMskChannelIcebergDestinationConfigurationOutputReference_Override(m MskChannelIcebergDestinationConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.mskChannel.MskChannelIcebergDestinationConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MskChannelIcebergDestinationConfigurationOutputReference)SetAppendOnly(val interface{}) {
	if err := j.validateSetAppendOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appendOnly",
		val,
	)
}

func (j *jsiiProxy_MskChannelIcebergDestinationConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MskChannelIcebergDestinationConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MskChannelIcebergDestinationConfigurationOutputReference)SetCompressionType(val *string) {
	if err := j.validateSetCompressionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compressionType",
		val,
	)
}

func (j *jsiiProxy_MskChannelIcebergDestinationConfigurationOutputReference)SetDataFreshnessInSeconds(val *float64) {
	if err := j.validateSetDataFreshnessInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataFreshnessInSeconds",
		val,
	)
}

func (j *jsiiProxy_MskChannelIcebergDestinationConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MskChannelIcebergDestinationConfigurationOutputReference)SetServiceExecutionRoleArn(val *string) {
	if err := j.validateSetServiceExecutionRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceExecutionRoleArn",
		val,
	)
}

func (j *jsiiProxy_MskChannelIcebergDestinationConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MskChannelIcebergDestinationConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MskChannelIcebergDestinationConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskChannelIcebergDestinationConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskChannelIcebergDestinationConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskChannelIcebergDestinationConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskChannelIcebergDestinationConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskChannelIcebergDestinationConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskChannelIcebergDestinationConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskChannelIcebergDestinationConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskChannelIcebergDestinationConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskChannelIcebergDestinationConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskChannelIcebergDestinationConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskChannelIcebergDestinationConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskChannelIcebergDestinationConfigurationOutputReference) PutCatalog(value *MskChannelIcebergDestinationConfigurationCatalog) {
	if err := m.validatePutCatalogParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putCatalog",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MskChannelIcebergDestinationConfigurationOutputReference) PutDeadLetterQueueS3(value *MskChannelIcebergDestinationConfigurationDeadLetterQueueS3) {
	if err := m.validatePutDeadLetterQueueS3Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putDeadLetterQueueS3",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MskChannelIcebergDestinationConfigurationOutputReference) PutDestinationTableList(value interface{}) {
	if err := m.validatePutDestinationTableListParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putDestinationTableList",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MskChannelIcebergDestinationConfigurationOutputReference) PutSchemaEvolution(value *MskChannelIcebergDestinationConfigurationSchemaEvolution) {
	if err := m.validatePutSchemaEvolutionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putSchemaEvolution",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MskChannelIcebergDestinationConfigurationOutputReference) PutTableCreation(value *MskChannelIcebergDestinationConfigurationTableCreation) {
	if err := m.validatePutTableCreationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putTableCreation",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MskChannelIcebergDestinationConfigurationOutputReference) ResetAppendOnly() {
	_jsii_.InvokeVoid(
		m,
		"resetAppendOnly",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MskChannelIcebergDestinationConfigurationOutputReference) ResetCatalog() {
	_jsii_.InvokeVoid(
		m,
		"resetCatalog",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MskChannelIcebergDestinationConfigurationOutputReference) ResetCompressionType() {
	_jsii_.InvokeVoid(
		m,
		"resetCompressionType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MskChannelIcebergDestinationConfigurationOutputReference) ResetDataFreshnessInSeconds() {
	_jsii_.InvokeVoid(
		m,
		"resetDataFreshnessInSeconds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MskChannelIcebergDestinationConfigurationOutputReference) ResetDeadLetterQueueS3() {
	_jsii_.InvokeVoid(
		m,
		"resetDeadLetterQueueS3",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MskChannelIcebergDestinationConfigurationOutputReference) ResetDestinationTableList() {
	_jsii_.InvokeVoid(
		m,
		"resetDestinationTableList",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MskChannelIcebergDestinationConfigurationOutputReference) ResetSchemaEvolution() {
	_jsii_.InvokeVoid(
		m,
		"resetSchemaEvolution",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MskChannelIcebergDestinationConfigurationOutputReference) ResetServiceExecutionRoleArn() {
	_jsii_.InvokeVoid(
		m,
		"resetServiceExecutionRoleArn",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MskChannelIcebergDestinationConfigurationOutputReference) ResetTableCreation() {
	_jsii_.InvokeVoid(
		m,
		"resetTableCreation",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MskChannelIcebergDestinationConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := m.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskChannelIcebergDestinationConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

