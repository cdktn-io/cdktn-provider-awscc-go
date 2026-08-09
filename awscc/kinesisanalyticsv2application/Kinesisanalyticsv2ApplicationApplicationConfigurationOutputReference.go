// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kinesisanalyticsv2application

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/kinesisanalyticsv2application/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference interface {
	cdktn.ComplexObject
	ApplicationCodeConfiguration() Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfigurationOutputReference
	ApplicationCodeConfigurationInput() interface{}
	ApplicationEncryptionConfiguration() Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationEncryptionConfigurationOutputReference
	ApplicationEncryptionConfigurationInput() interface{}
	ApplicationSnapshotConfiguration() Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationSnapshotConfigurationOutputReference
	ApplicationSnapshotConfigurationInput() interface{}
	ApplicationSystemRollbackConfiguration() Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationSystemRollbackConfigurationOutputReference
	ApplicationSystemRollbackConfigurationInput() interface{}
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
	EnvironmentProperties() Kinesisanalyticsv2ApplicationApplicationConfigurationEnvironmentPropertiesOutputReference
	EnvironmentPropertiesInput() interface{}
	FlinkApplicationConfiguration() Kinesisanalyticsv2ApplicationApplicationConfigurationFlinkApplicationConfigurationOutputReference
	FlinkApplicationConfigurationInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	SqlApplicationConfiguration() Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationOutputReference
	SqlApplicationConfigurationInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	VpcConfigurations() Kinesisanalyticsv2ApplicationApplicationConfigurationVpcConfigurationsList
	VpcConfigurationsInput() interface{}
	ZeppelinApplicationConfiguration() Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationOutputReference
	ZeppelinApplicationConfigurationInput() interface{}
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
	PutApplicationCodeConfiguration(value *Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfiguration)
	PutApplicationEncryptionConfiguration(value *Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationEncryptionConfiguration)
	PutApplicationSnapshotConfiguration(value *Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationSnapshotConfiguration)
	PutApplicationSystemRollbackConfiguration(value *Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationSystemRollbackConfiguration)
	PutEnvironmentProperties(value *Kinesisanalyticsv2ApplicationApplicationConfigurationEnvironmentProperties)
	PutFlinkApplicationConfiguration(value *Kinesisanalyticsv2ApplicationApplicationConfigurationFlinkApplicationConfiguration)
	PutSqlApplicationConfiguration(value *Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfiguration)
	PutVpcConfigurations(value interface{})
	PutZeppelinApplicationConfiguration(value *Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfiguration)
	ResetApplicationCodeConfiguration()
	ResetApplicationEncryptionConfiguration()
	ResetApplicationSnapshotConfiguration()
	ResetApplicationSystemRollbackConfiguration()
	ResetEnvironmentProperties()
	ResetFlinkApplicationConfiguration()
	ResetSqlApplicationConfiguration()
	ResetVpcConfigurations()
	ResetZeppelinApplicationConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference
type jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) ApplicationCodeConfiguration() Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfigurationOutputReference {
	var returns Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfigurationOutputReference
	_jsii_.Get(
		j,
		"applicationCodeConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) ApplicationCodeConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applicationCodeConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) ApplicationEncryptionConfiguration() Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationEncryptionConfigurationOutputReference {
	var returns Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationEncryptionConfigurationOutputReference
	_jsii_.Get(
		j,
		"applicationEncryptionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) ApplicationEncryptionConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applicationEncryptionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) ApplicationSnapshotConfiguration() Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationSnapshotConfigurationOutputReference {
	var returns Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationSnapshotConfigurationOutputReference
	_jsii_.Get(
		j,
		"applicationSnapshotConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) ApplicationSnapshotConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applicationSnapshotConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) ApplicationSystemRollbackConfiguration() Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationSystemRollbackConfigurationOutputReference {
	var returns Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationSystemRollbackConfigurationOutputReference
	_jsii_.Get(
		j,
		"applicationSystemRollbackConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) ApplicationSystemRollbackConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applicationSystemRollbackConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) EnvironmentProperties() Kinesisanalyticsv2ApplicationApplicationConfigurationEnvironmentPropertiesOutputReference {
	var returns Kinesisanalyticsv2ApplicationApplicationConfigurationEnvironmentPropertiesOutputReference
	_jsii_.Get(
		j,
		"environmentProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) EnvironmentPropertiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"environmentPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) FlinkApplicationConfiguration() Kinesisanalyticsv2ApplicationApplicationConfigurationFlinkApplicationConfigurationOutputReference {
	var returns Kinesisanalyticsv2ApplicationApplicationConfigurationFlinkApplicationConfigurationOutputReference
	_jsii_.Get(
		j,
		"flinkApplicationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) FlinkApplicationConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"flinkApplicationConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) SqlApplicationConfiguration() Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationOutputReference {
	var returns Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationOutputReference
	_jsii_.Get(
		j,
		"sqlApplicationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) SqlApplicationConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sqlApplicationConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) VpcConfigurations() Kinesisanalyticsv2ApplicationApplicationConfigurationVpcConfigurationsList {
	var returns Kinesisanalyticsv2ApplicationApplicationConfigurationVpcConfigurationsList
	_jsii_.Get(
		j,
		"vpcConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) VpcConfigurationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vpcConfigurationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) ZeppelinApplicationConfiguration() Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationOutputReference {
	var returns Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationOutputReference
	_jsii_.Get(
		j,
		"zeppelinApplicationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) ZeppelinApplicationConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"zeppelinApplicationConfigurationInput",
		&returns,
	)
	return returns
}


func NewKinesisanalyticsv2ApplicationApplicationConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewKinesisanalyticsv2ApplicationApplicationConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.kinesisanalyticsv2Application.Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKinesisanalyticsv2ApplicationApplicationConfigurationOutputReference_Override(k Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.kinesisanalyticsv2Application.Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) PutApplicationCodeConfiguration(value *Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfiguration) {
	if err := k.validatePutApplicationCodeConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putApplicationCodeConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) PutApplicationEncryptionConfiguration(value *Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationEncryptionConfiguration) {
	if err := k.validatePutApplicationEncryptionConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putApplicationEncryptionConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) PutApplicationSnapshotConfiguration(value *Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationSnapshotConfiguration) {
	if err := k.validatePutApplicationSnapshotConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putApplicationSnapshotConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) PutApplicationSystemRollbackConfiguration(value *Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationSystemRollbackConfiguration) {
	if err := k.validatePutApplicationSystemRollbackConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putApplicationSystemRollbackConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) PutEnvironmentProperties(value *Kinesisanalyticsv2ApplicationApplicationConfigurationEnvironmentProperties) {
	if err := k.validatePutEnvironmentPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putEnvironmentProperties",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) PutFlinkApplicationConfiguration(value *Kinesisanalyticsv2ApplicationApplicationConfigurationFlinkApplicationConfiguration) {
	if err := k.validatePutFlinkApplicationConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putFlinkApplicationConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) PutSqlApplicationConfiguration(value *Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfiguration) {
	if err := k.validatePutSqlApplicationConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putSqlApplicationConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) PutVpcConfigurations(value interface{}) {
	if err := k.validatePutVpcConfigurationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putVpcConfigurations",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) PutZeppelinApplicationConfiguration(value *Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfiguration) {
	if err := k.validatePutZeppelinApplicationConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putZeppelinApplicationConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) ResetApplicationCodeConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetApplicationCodeConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) ResetApplicationEncryptionConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetApplicationEncryptionConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) ResetApplicationSnapshotConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetApplicationSnapshotConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) ResetApplicationSystemRollbackConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetApplicationSystemRollbackConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) ResetEnvironmentProperties() {
	_jsii_.InvokeVoid(
		k,
		"resetEnvironmentProperties",
		nil, // no parameters
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) ResetFlinkApplicationConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetFlinkApplicationConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) ResetSqlApplicationConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetSqlApplicationConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) ResetVpcConfigurations() {
	_jsii_.InvokeVoid(
		k,
		"resetVpcConfigurations",
		nil, // no parameters
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) ResetZeppelinApplicationConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetZeppelinApplicationConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

