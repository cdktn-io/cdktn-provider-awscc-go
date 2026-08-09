// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kinesisfirehosedeliverystream

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/kinesisfirehosedeliverystream/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference interface {
	cdktn.ComplexObject
	Columns() KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationColumnsOutputReference
	ColumnsInput() interface{}
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
	Databases() KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationDatabasesOutputReference
	DatabasesInput() interface{}
	DatabaseSourceAuthenticationConfiguration() KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationDatabaseSourceAuthenticationConfigurationOutputReference
	DatabaseSourceAuthenticationConfigurationInput() interface{}
	DatabaseSourceVpcConfiguration() KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationDatabaseSourceVpcConfigurationOutputReference
	DatabaseSourceVpcConfigurationInput() interface{}
	Digest() *string
	SetDigest(val *string)
	DigestInput() *string
	Endpoint() *string
	SetEndpoint(val *string)
	EndpointInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
	PublicCertificate() *string
	SetPublicCertificate(val *string)
	PublicCertificateInput() *string
	SnapshotWatermarkTable() *string
	SetSnapshotWatermarkTable(val *string)
	SnapshotWatermarkTableInput() *string
	SslMode() *string
	SetSslMode(val *string)
	SslModeInput() *string
	SurrogateKeys() *[]*string
	SetSurrogateKeys(val *[]*string)
	SurrogateKeysInput() *[]*string
	Tables() KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationTablesOutputReference
	TablesInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	PutColumns(value *KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationColumns)
	PutDatabases(value *KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationDatabases)
	PutDatabaseSourceAuthenticationConfiguration(value *KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationDatabaseSourceAuthenticationConfiguration)
	PutDatabaseSourceVpcConfiguration(value *KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationDatabaseSourceVpcConfiguration)
	PutTables(value *KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationTables)
	ResetColumns()
	ResetDatabases()
	ResetDatabaseSourceAuthenticationConfiguration()
	ResetDatabaseSourceVpcConfiguration()
	ResetDigest()
	ResetEndpoint()
	ResetPort()
	ResetPublicCertificate()
	ResetSnapshotWatermarkTable()
	ResetSslMode()
	ResetSurrogateKeys()
	ResetTables()
	ResetType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference
type jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference) Columns() KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationColumnsOutputReference {
	var returns KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationColumnsOutputReference
	_jsii_.Get(
		j,
		"columns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference) ColumnsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"columnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference) Databases() KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationDatabasesOutputReference {
	var returns KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationDatabasesOutputReference
	_jsii_.Get(
		j,
		"databases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference) DatabasesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"databasesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference) DatabaseSourceAuthenticationConfiguration() KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationDatabaseSourceAuthenticationConfigurationOutputReference {
	var returns KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationDatabaseSourceAuthenticationConfigurationOutputReference
	_jsii_.Get(
		j,
		"databaseSourceAuthenticationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference) DatabaseSourceAuthenticationConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"databaseSourceAuthenticationConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference) DatabaseSourceVpcConfiguration() KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationDatabaseSourceVpcConfigurationOutputReference {
	var returns KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationDatabaseSourceVpcConfigurationOutputReference
	_jsii_.Get(
		j,
		"databaseSourceVpcConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference) DatabaseSourceVpcConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"databaseSourceVpcConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference) Digest() *string {
	var returns *string
	_jsii_.Get(
		j,
		"digest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference) DigestInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"digestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference) EndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference) PublicCertificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference) PublicCertificateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference) SnapshotWatermarkTable() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotWatermarkTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference) SnapshotWatermarkTableInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotWatermarkTableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference) SslMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference) SslModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference) SurrogateKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"surrogateKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference) SurrogateKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"surrogateKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference) Tables() KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationTablesOutputReference {
	var returns KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationTablesOutputReference
	_jsii_.Get(
		j,
		"tables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference) TablesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewKinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewKinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference_Override(k KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference)SetDigest(val *string) {
	if err := j.validateSetDigestParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"digest",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference)SetEndpoint(val *string) {
	if err := j.validateSetEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpoint",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference)SetPublicCertificate(val *string) {
	if err := j.validateSetPublicCertificateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicCertificate",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference)SetSnapshotWatermarkTable(val *string) {
	if err := j.validateSetSnapshotWatermarkTableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotWatermarkTable",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference)SetSslMode(val *string) {
	if err := j.validateSetSslModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslMode",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference)SetSurrogateKeys(val *[]*string) {
	if err := j.validateSetSurrogateKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"surrogateKeys",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference) PutColumns(value *KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationColumns) {
	if err := k.validatePutColumnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putColumns",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference) PutDatabases(value *KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationDatabases) {
	if err := k.validatePutDatabasesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putDatabases",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference) PutDatabaseSourceAuthenticationConfiguration(value *KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationDatabaseSourceAuthenticationConfiguration) {
	if err := k.validatePutDatabaseSourceAuthenticationConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putDatabaseSourceAuthenticationConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference) PutDatabaseSourceVpcConfiguration(value *KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationDatabaseSourceVpcConfiguration) {
	if err := k.validatePutDatabaseSourceVpcConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putDatabaseSourceVpcConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference) PutTables(value *KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationTables) {
	if err := k.validatePutTablesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putTables",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference) ResetColumns() {
	_jsii_.InvokeVoid(
		k,
		"resetColumns",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference) ResetDatabases() {
	_jsii_.InvokeVoid(
		k,
		"resetDatabases",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference) ResetDatabaseSourceAuthenticationConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetDatabaseSourceAuthenticationConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference) ResetDatabaseSourceVpcConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetDatabaseSourceVpcConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference) ResetDigest() {
	_jsii_.InvokeVoid(
		k,
		"resetDigest",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference) ResetEndpoint() {
	_jsii_.InvokeVoid(
		k,
		"resetEndpoint",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference) ResetPort() {
	_jsii_.InvokeVoid(
		k,
		"resetPort",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference) ResetPublicCertificate() {
	_jsii_.InvokeVoid(
		k,
		"resetPublicCertificate",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference) ResetSnapshotWatermarkTable() {
	_jsii_.InvokeVoid(
		k,
		"resetSnapshotWatermarkTable",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference) ResetSslMode() {
	_jsii_.InvokeVoid(
		k,
		"resetSslMode",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference) ResetSurrogateKeys() {
	_jsii_.InvokeVoid(
		k,
		"resetSurrogateKeys",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference) ResetTables() {
	_jsii_.InvokeVoid(
		k,
		"resetTables",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		k,
		"resetType",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

