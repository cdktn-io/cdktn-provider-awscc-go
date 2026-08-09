// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datazoneconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/datazoneconnection/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DatazoneConnectionPropsSparkGluePropertiesOutputReference interface {
	cdktn.ComplexObject
	AdditionalArgs() DatazoneConnectionPropsSparkGluePropertiesAdditionalArgsOutputReference
	AdditionalArgsInput() interface{}
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
	GlueConnectionName() *string
	SetGlueConnectionName(val *string)
	GlueConnectionNameInput() *string
	GlueVersion() *string
	SetGlueVersion(val *string)
	GlueVersionInput() *string
	IdleTimeout() *float64
	SetIdleTimeout(val *float64)
	IdleTimeoutInput() *float64
	InternalValue() interface{}
	SetInternalValue(val interface{})
	JavaVirtualEnv() *string
	SetJavaVirtualEnv(val *string)
	JavaVirtualEnvInput() *string
	NumberOfWorkers() *float64
	SetNumberOfWorkers(val *float64)
	NumberOfWorkersInput() *float64
	PythonVirtualEnv() *string
	SetPythonVirtualEnv(val *string)
	PythonVirtualEnvInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	WorkerType() *string
	SetWorkerType(val *string)
	WorkerTypeInput() *string
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
	PutAdditionalArgs(value *DatazoneConnectionPropsSparkGluePropertiesAdditionalArgs)
	ResetAdditionalArgs()
	ResetGlueConnectionName()
	ResetGlueVersion()
	ResetIdleTimeout()
	ResetJavaVirtualEnv()
	ResetNumberOfWorkers()
	ResetPythonVirtualEnv()
	ResetWorkerType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DatazoneConnectionPropsSparkGluePropertiesOutputReference
type jsiiProxy_DatazoneConnectionPropsSparkGluePropertiesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DatazoneConnectionPropsSparkGluePropertiesOutputReference) AdditionalArgs() DatazoneConnectionPropsSparkGluePropertiesAdditionalArgsOutputReference {
	var returns DatazoneConnectionPropsSparkGluePropertiesAdditionalArgsOutputReference
	_jsii_.Get(
		j,
		"additionalArgs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsSparkGluePropertiesOutputReference) AdditionalArgsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"additionalArgsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsSparkGluePropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsSparkGluePropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsSparkGluePropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsSparkGluePropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsSparkGluePropertiesOutputReference) GlueConnectionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"glueConnectionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsSparkGluePropertiesOutputReference) GlueConnectionNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"glueConnectionNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsSparkGluePropertiesOutputReference) GlueVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"glueVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsSparkGluePropertiesOutputReference) GlueVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"glueVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsSparkGluePropertiesOutputReference) IdleTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsSparkGluePropertiesOutputReference) IdleTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsSparkGluePropertiesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsSparkGluePropertiesOutputReference) JavaVirtualEnv() *string {
	var returns *string
	_jsii_.Get(
		j,
		"javaVirtualEnv",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsSparkGluePropertiesOutputReference) JavaVirtualEnvInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"javaVirtualEnvInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsSparkGluePropertiesOutputReference) NumberOfWorkers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfWorkers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsSparkGluePropertiesOutputReference) NumberOfWorkersInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfWorkersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsSparkGluePropertiesOutputReference) PythonVirtualEnv() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pythonVirtualEnv",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsSparkGluePropertiesOutputReference) PythonVirtualEnvInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pythonVirtualEnvInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsSparkGluePropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsSparkGluePropertiesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsSparkGluePropertiesOutputReference) WorkerType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workerType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsSparkGluePropertiesOutputReference) WorkerTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workerTypeInput",
		&returns,
	)
	return returns
}


func NewDatazoneConnectionPropsSparkGluePropertiesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DatazoneConnectionPropsSparkGluePropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewDatazoneConnectionPropsSparkGluePropertiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatazoneConnectionPropsSparkGluePropertiesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.datazoneConnection.DatazoneConnectionPropsSparkGluePropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDatazoneConnectionPropsSparkGluePropertiesOutputReference_Override(d DatazoneConnectionPropsSparkGluePropertiesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.datazoneConnection.DatazoneConnectionPropsSparkGluePropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DatazoneConnectionPropsSparkGluePropertiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DatazoneConnectionPropsSparkGluePropertiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DatazoneConnectionPropsSparkGluePropertiesOutputReference)SetGlueConnectionName(val *string) {
	if err := j.validateSetGlueConnectionNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"glueConnectionName",
		val,
	)
}

func (j *jsiiProxy_DatazoneConnectionPropsSparkGluePropertiesOutputReference)SetGlueVersion(val *string) {
	if err := j.validateSetGlueVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"glueVersion",
		val,
	)
}

func (j *jsiiProxy_DatazoneConnectionPropsSparkGluePropertiesOutputReference)SetIdleTimeout(val *float64) {
	if err := j.validateSetIdleTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idleTimeout",
		val,
	)
}

func (j *jsiiProxy_DatazoneConnectionPropsSparkGluePropertiesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatazoneConnectionPropsSparkGluePropertiesOutputReference)SetJavaVirtualEnv(val *string) {
	if err := j.validateSetJavaVirtualEnvParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"javaVirtualEnv",
		val,
	)
}

func (j *jsiiProxy_DatazoneConnectionPropsSparkGluePropertiesOutputReference)SetNumberOfWorkers(val *float64) {
	if err := j.validateSetNumberOfWorkersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numberOfWorkers",
		val,
	)
}

func (j *jsiiProxy_DatazoneConnectionPropsSparkGluePropertiesOutputReference)SetPythonVirtualEnv(val *string) {
	if err := j.validateSetPythonVirtualEnvParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pythonVirtualEnv",
		val,
	)
}

func (j *jsiiProxy_DatazoneConnectionPropsSparkGluePropertiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatazoneConnectionPropsSparkGluePropertiesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DatazoneConnectionPropsSparkGluePropertiesOutputReference)SetWorkerType(val *string) {
	if err := j.validateSetWorkerTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workerType",
		val,
	)
}

func (d *jsiiProxy_DatazoneConnectionPropsSparkGluePropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatazoneConnectionPropsSparkGluePropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DatazoneConnectionPropsSparkGluePropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DatazoneConnectionPropsSparkGluePropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DatazoneConnectionPropsSparkGluePropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DatazoneConnectionPropsSparkGluePropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DatazoneConnectionPropsSparkGluePropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DatazoneConnectionPropsSparkGluePropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DatazoneConnectionPropsSparkGluePropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DatazoneConnectionPropsSparkGluePropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DatazoneConnectionPropsSparkGluePropertiesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatazoneConnectionPropsSparkGluePropertiesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DatazoneConnectionPropsSparkGluePropertiesOutputReference) PutAdditionalArgs(value *DatazoneConnectionPropsSparkGluePropertiesAdditionalArgs) {
	if err := d.validatePutAdditionalArgsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAdditionalArgs",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatazoneConnectionPropsSparkGluePropertiesOutputReference) ResetAdditionalArgs() {
	_jsii_.InvokeVoid(
		d,
		"resetAdditionalArgs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazoneConnectionPropsSparkGluePropertiesOutputReference) ResetGlueConnectionName() {
	_jsii_.InvokeVoid(
		d,
		"resetGlueConnectionName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazoneConnectionPropsSparkGluePropertiesOutputReference) ResetGlueVersion() {
	_jsii_.InvokeVoid(
		d,
		"resetGlueVersion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazoneConnectionPropsSparkGluePropertiesOutputReference) ResetIdleTimeout() {
	_jsii_.InvokeVoid(
		d,
		"resetIdleTimeout",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazoneConnectionPropsSparkGluePropertiesOutputReference) ResetJavaVirtualEnv() {
	_jsii_.InvokeVoid(
		d,
		"resetJavaVirtualEnv",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazoneConnectionPropsSparkGluePropertiesOutputReference) ResetNumberOfWorkers() {
	_jsii_.InvokeVoid(
		d,
		"resetNumberOfWorkers",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazoneConnectionPropsSparkGluePropertiesOutputReference) ResetPythonVirtualEnv() {
	_jsii_.InvokeVoid(
		d,
		"resetPythonVirtualEnv",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazoneConnectionPropsSparkGluePropertiesOutputReference) ResetWorkerType() {
	_jsii_.InvokeVoid(
		d,
		"resetWorkerType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazoneConnectionPropsSparkGluePropertiesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DatazoneConnectionPropsSparkGluePropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

