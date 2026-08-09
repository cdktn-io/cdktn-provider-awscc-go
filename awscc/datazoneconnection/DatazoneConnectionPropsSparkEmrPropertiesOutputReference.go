// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datazoneconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/datazoneconnection/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DatazoneConnectionPropsSparkEmrPropertiesOutputReference interface {
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
	ComputeArn() *string
	SetComputeArn(val *string)
	ComputeArnInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InstanceProfileArn() *string
	SetInstanceProfileArn(val *string)
	InstanceProfileArnInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	JavaVirtualEnv() *string
	SetJavaVirtualEnv(val *string)
	JavaVirtualEnvInput() *string
	LogUri() *string
	SetLogUri(val *string)
	LogUriInput() *string
	ManagedEndpointArn() *string
	SetManagedEndpointArn(val *string)
	ManagedEndpointArnInput() *string
	PythonVirtualEnv() *string
	SetPythonVirtualEnv(val *string)
	PythonVirtualEnvInput() *string
	RuntimeRole() *string
	SetRuntimeRole(val *string)
	RuntimeRoleInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TrustedCertificatesS3Uri() *string
	SetTrustedCertificatesS3Uri(val *string)
	TrustedCertificatesS3UriInput() *string
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
	ResetComputeArn()
	ResetInstanceProfileArn()
	ResetJavaVirtualEnv()
	ResetLogUri()
	ResetManagedEndpointArn()
	ResetPythonVirtualEnv()
	ResetRuntimeRole()
	ResetTrustedCertificatesS3Uri()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DatazoneConnectionPropsSparkEmrPropertiesOutputReference
type jsiiProxy_DatazoneConnectionPropsSparkEmrPropertiesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DatazoneConnectionPropsSparkEmrPropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsSparkEmrPropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsSparkEmrPropertiesOutputReference) ComputeArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsSparkEmrPropertiesOutputReference) ComputeArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsSparkEmrPropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsSparkEmrPropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsSparkEmrPropertiesOutputReference) InstanceProfileArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceProfileArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsSparkEmrPropertiesOutputReference) InstanceProfileArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceProfileArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsSparkEmrPropertiesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsSparkEmrPropertiesOutputReference) JavaVirtualEnv() *string {
	var returns *string
	_jsii_.Get(
		j,
		"javaVirtualEnv",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsSparkEmrPropertiesOutputReference) JavaVirtualEnvInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"javaVirtualEnvInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsSparkEmrPropertiesOutputReference) LogUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsSparkEmrPropertiesOutputReference) LogUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsSparkEmrPropertiesOutputReference) ManagedEndpointArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedEndpointArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsSparkEmrPropertiesOutputReference) ManagedEndpointArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedEndpointArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsSparkEmrPropertiesOutputReference) PythonVirtualEnv() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pythonVirtualEnv",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsSparkEmrPropertiesOutputReference) PythonVirtualEnvInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pythonVirtualEnvInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsSparkEmrPropertiesOutputReference) RuntimeRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsSparkEmrPropertiesOutputReference) RuntimeRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsSparkEmrPropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsSparkEmrPropertiesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsSparkEmrPropertiesOutputReference) TrustedCertificatesS3Uri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustedCertificatesS3Uri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsSparkEmrPropertiesOutputReference) TrustedCertificatesS3UriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustedCertificatesS3UriInput",
		&returns,
	)
	return returns
}


func NewDatazoneConnectionPropsSparkEmrPropertiesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DatazoneConnectionPropsSparkEmrPropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewDatazoneConnectionPropsSparkEmrPropertiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatazoneConnectionPropsSparkEmrPropertiesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.datazoneConnection.DatazoneConnectionPropsSparkEmrPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDatazoneConnectionPropsSparkEmrPropertiesOutputReference_Override(d DatazoneConnectionPropsSparkEmrPropertiesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.datazoneConnection.DatazoneConnectionPropsSparkEmrPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DatazoneConnectionPropsSparkEmrPropertiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DatazoneConnectionPropsSparkEmrPropertiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DatazoneConnectionPropsSparkEmrPropertiesOutputReference)SetComputeArn(val *string) {
	if err := j.validateSetComputeArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"computeArn",
		val,
	)
}

func (j *jsiiProxy_DatazoneConnectionPropsSparkEmrPropertiesOutputReference)SetInstanceProfileArn(val *string) {
	if err := j.validateSetInstanceProfileArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceProfileArn",
		val,
	)
}

func (j *jsiiProxy_DatazoneConnectionPropsSparkEmrPropertiesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatazoneConnectionPropsSparkEmrPropertiesOutputReference)SetJavaVirtualEnv(val *string) {
	if err := j.validateSetJavaVirtualEnvParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"javaVirtualEnv",
		val,
	)
}

func (j *jsiiProxy_DatazoneConnectionPropsSparkEmrPropertiesOutputReference)SetLogUri(val *string) {
	if err := j.validateSetLogUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logUri",
		val,
	)
}

func (j *jsiiProxy_DatazoneConnectionPropsSparkEmrPropertiesOutputReference)SetManagedEndpointArn(val *string) {
	if err := j.validateSetManagedEndpointArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"managedEndpointArn",
		val,
	)
}

func (j *jsiiProxy_DatazoneConnectionPropsSparkEmrPropertiesOutputReference)SetPythonVirtualEnv(val *string) {
	if err := j.validateSetPythonVirtualEnvParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pythonVirtualEnv",
		val,
	)
}

func (j *jsiiProxy_DatazoneConnectionPropsSparkEmrPropertiesOutputReference)SetRuntimeRole(val *string) {
	if err := j.validateSetRuntimeRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtimeRole",
		val,
	)
}

func (j *jsiiProxy_DatazoneConnectionPropsSparkEmrPropertiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatazoneConnectionPropsSparkEmrPropertiesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DatazoneConnectionPropsSparkEmrPropertiesOutputReference)SetTrustedCertificatesS3Uri(val *string) {
	if err := j.validateSetTrustedCertificatesS3UriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trustedCertificatesS3Uri",
		val,
	)
}

func (d *jsiiProxy_DatazoneConnectionPropsSparkEmrPropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatazoneConnectionPropsSparkEmrPropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DatazoneConnectionPropsSparkEmrPropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DatazoneConnectionPropsSparkEmrPropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DatazoneConnectionPropsSparkEmrPropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DatazoneConnectionPropsSparkEmrPropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DatazoneConnectionPropsSparkEmrPropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DatazoneConnectionPropsSparkEmrPropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DatazoneConnectionPropsSparkEmrPropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DatazoneConnectionPropsSparkEmrPropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DatazoneConnectionPropsSparkEmrPropertiesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatazoneConnectionPropsSparkEmrPropertiesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DatazoneConnectionPropsSparkEmrPropertiesOutputReference) ResetComputeArn() {
	_jsii_.InvokeVoid(
		d,
		"resetComputeArn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazoneConnectionPropsSparkEmrPropertiesOutputReference) ResetInstanceProfileArn() {
	_jsii_.InvokeVoid(
		d,
		"resetInstanceProfileArn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazoneConnectionPropsSparkEmrPropertiesOutputReference) ResetJavaVirtualEnv() {
	_jsii_.InvokeVoid(
		d,
		"resetJavaVirtualEnv",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazoneConnectionPropsSparkEmrPropertiesOutputReference) ResetLogUri() {
	_jsii_.InvokeVoid(
		d,
		"resetLogUri",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazoneConnectionPropsSparkEmrPropertiesOutputReference) ResetManagedEndpointArn() {
	_jsii_.InvokeVoid(
		d,
		"resetManagedEndpointArn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazoneConnectionPropsSparkEmrPropertiesOutputReference) ResetPythonVirtualEnv() {
	_jsii_.InvokeVoid(
		d,
		"resetPythonVirtualEnv",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazoneConnectionPropsSparkEmrPropertiesOutputReference) ResetRuntimeRole() {
	_jsii_.InvokeVoid(
		d,
		"resetRuntimeRole",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazoneConnectionPropsSparkEmrPropertiesOutputReference) ResetTrustedCertificatesS3Uri() {
	_jsii_.InvokeVoid(
		d,
		"resetTrustedCertificatesS3Uri",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazoneConnectionPropsSparkEmrPropertiesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DatazoneConnectionPropsSparkEmrPropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

