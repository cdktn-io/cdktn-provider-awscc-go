// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iototaupdate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/iototaupdate/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IotOtaUpdateFilesCodeSigningStartSigningJobParameterSigningProfileParameterOutputReference interface {
	cdktn.ComplexObject
	CertificateArn() *string
	SetCertificateArn(val *string)
	CertificateArnInput() *string
	CertificatePathOnDevice() *string
	SetCertificatePathOnDevice(val *string)
	CertificatePathOnDeviceInput() *string
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Platform() *string
	SetPlatform(val *string)
	PlatformInput() *string
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
	ResetCertificateArn()
	ResetCertificatePathOnDevice()
	ResetPlatform()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IotOtaUpdateFilesCodeSigningStartSigningJobParameterSigningProfileParameterOutputReference
type jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterSigningProfileParameterOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterSigningProfileParameterOutputReference) CertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterSigningProfileParameterOutputReference) CertificateArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterSigningProfileParameterOutputReference) CertificatePathOnDevice() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificatePathOnDevice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterSigningProfileParameterOutputReference) CertificatePathOnDeviceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificatePathOnDeviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterSigningProfileParameterOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterSigningProfileParameterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterSigningProfileParameterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterSigningProfileParameterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterSigningProfileParameterOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterSigningProfileParameterOutputReference) Platform() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterSigningProfileParameterOutputReference) PlatformInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterSigningProfileParameterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterSigningProfileParameterOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewIotOtaUpdateFilesCodeSigningStartSigningJobParameterSigningProfileParameterOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) IotOtaUpdateFilesCodeSigningStartSigningJobParameterSigningProfileParameterOutputReference {
	_init_.Initialize()

	if err := validateNewIotOtaUpdateFilesCodeSigningStartSigningJobParameterSigningProfileParameterOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterSigningProfileParameterOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.iotOtaUpdate.IotOtaUpdateFilesCodeSigningStartSigningJobParameterSigningProfileParameterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIotOtaUpdateFilesCodeSigningStartSigningJobParameterSigningProfileParameterOutputReference_Override(i IotOtaUpdateFilesCodeSigningStartSigningJobParameterSigningProfileParameterOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.iotOtaUpdate.IotOtaUpdateFilesCodeSigningStartSigningJobParameterSigningProfileParameterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterSigningProfileParameterOutputReference)SetCertificateArn(val *string) {
	if err := j.validateSetCertificateArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificateArn",
		val,
	)
}

func (j *jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterSigningProfileParameterOutputReference)SetCertificatePathOnDevice(val *string) {
	if err := j.validateSetCertificatePathOnDeviceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificatePathOnDevice",
		val,
	)
}

func (j *jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterSigningProfileParameterOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterSigningProfileParameterOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterSigningProfileParameterOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterSigningProfileParameterOutputReference)SetPlatform(val *string) {
	if err := j.validateSetPlatformParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"platform",
		val,
	)
}

func (j *jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterSigningProfileParameterOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterSigningProfileParameterOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterSigningProfileParameterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterSigningProfileParameterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterSigningProfileParameterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterSigningProfileParameterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterSigningProfileParameterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterSigningProfileParameterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterSigningProfileParameterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterSigningProfileParameterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterSigningProfileParameterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterSigningProfileParameterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterSigningProfileParameterOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterSigningProfileParameterOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterSigningProfileParameterOutputReference) ResetCertificateArn() {
	_jsii_.InvokeVoid(
		i,
		"resetCertificateArn",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterSigningProfileParameterOutputReference) ResetCertificatePathOnDevice() {
	_jsii_.InvokeVoid(
		i,
		"resetCertificatePathOnDevice",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterSigningProfileParameterOutputReference) ResetPlatform() {
	_jsii_.InvokeVoid(
		i,
		"resetPlatform",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterSigningProfileParameterOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := i.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterSigningProfileParameterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

