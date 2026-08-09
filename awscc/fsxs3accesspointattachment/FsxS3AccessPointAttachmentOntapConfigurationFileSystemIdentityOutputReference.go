// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package fsxs3accesspointattachment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/fsxs3accesspointattachment/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type FsxS3AccessPointAttachmentOntapConfigurationFileSystemIdentityOutputReference interface {
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
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
	UnixUser() FsxS3AccessPointAttachmentOntapConfigurationFileSystemIdentityUnixUserOutputReference
	UnixUserInput() interface{}
	WindowsUser() FsxS3AccessPointAttachmentOntapConfigurationFileSystemIdentityWindowsUserOutputReference
	WindowsUserInput() interface{}
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
	PutUnixUser(value *FsxS3AccessPointAttachmentOntapConfigurationFileSystemIdentityUnixUser)
	PutWindowsUser(value *FsxS3AccessPointAttachmentOntapConfigurationFileSystemIdentityWindowsUser)
	ResetType()
	ResetUnixUser()
	ResetWindowsUser()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FsxS3AccessPointAttachmentOntapConfigurationFileSystemIdentityOutputReference
type jsiiProxy_FsxS3AccessPointAttachmentOntapConfigurationFileSystemIdentityOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_FsxS3AccessPointAttachmentOntapConfigurationFileSystemIdentityOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxS3AccessPointAttachmentOntapConfigurationFileSystemIdentityOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxS3AccessPointAttachmentOntapConfigurationFileSystemIdentityOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxS3AccessPointAttachmentOntapConfigurationFileSystemIdentityOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxS3AccessPointAttachmentOntapConfigurationFileSystemIdentityOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxS3AccessPointAttachmentOntapConfigurationFileSystemIdentityOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxS3AccessPointAttachmentOntapConfigurationFileSystemIdentityOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxS3AccessPointAttachmentOntapConfigurationFileSystemIdentityOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxS3AccessPointAttachmentOntapConfigurationFileSystemIdentityOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxS3AccessPointAttachmentOntapConfigurationFileSystemIdentityOutputReference) UnixUser() FsxS3AccessPointAttachmentOntapConfigurationFileSystemIdentityUnixUserOutputReference {
	var returns FsxS3AccessPointAttachmentOntapConfigurationFileSystemIdentityUnixUserOutputReference
	_jsii_.Get(
		j,
		"unixUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxS3AccessPointAttachmentOntapConfigurationFileSystemIdentityOutputReference) UnixUserInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"unixUserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxS3AccessPointAttachmentOntapConfigurationFileSystemIdentityOutputReference) WindowsUser() FsxS3AccessPointAttachmentOntapConfigurationFileSystemIdentityWindowsUserOutputReference {
	var returns FsxS3AccessPointAttachmentOntapConfigurationFileSystemIdentityWindowsUserOutputReference
	_jsii_.Get(
		j,
		"windowsUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxS3AccessPointAttachmentOntapConfigurationFileSystemIdentityOutputReference) WindowsUserInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"windowsUserInput",
		&returns,
	)
	return returns
}


func NewFsxS3AccessPointAttachmentOntapConfigurationFileSystemIdentityOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) FsxS3AccessPointAttachmentOntapConfigurationFileSystemIdentityOutputReference {
	_init_.Initialize()

	if err := validateNewFsxS3AccessPointAttachmentOntapConfigurationFileSystemIdentityOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_FsxS3AccessPointAttachmentOntapConfigurationFileSystemIdentityOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.fsxS3AccessPointAttachment.FsxS3AccessPointAttachmentOntapConfigurationFileSystemIdentityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewFsxS3AccessPointAttachmentOntapConfigurationFileSystemIdentityOutputReference_Override(f FsxS3AccessPointAttachmentOntapConfigurationFileSystemIdentityOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.fsxS3AccessPointAttachment.FsxS3AccessPointAttachmentOntapConfigurationFileSystemIdentityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		f,
	)
}

func (j *jsiiProxy_FsxS3AccessPointAttachmentOntapConfigurationFileSystemIdentityOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FsxS3AccessPointAttachmentOntapConfigurationFileSystemIdentityOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FsxS3AccessPointAttachmentOntapConfigurationFileSystemIdentityOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FsxS3AccessPointAttachmentOntapConfigurationFileSystemIdentityOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FsxS3AccessPointAttachmentOntapConfigurationFileSystemIdentityOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FsxS3AccessPointAttachmentOntapConfigurationFileSystemIdentityOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (f *jsiiProxy_FsxS3AccessPointAttachmentOntapConfigurationFileSystemIdentityOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxS3AccessPointAttachmentOntapConfigurationFileSystemIdentityOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := f.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxS3AccessPointAttachmentOntapConfigurationFileSystemIdentityOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := f.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxS3AccessPointAttachmentOntapConfigurationFileSystemIdentityOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := f.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxS3AccessPointAttachmentOntapConfigurationFileSystemIdentityOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := f.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxS3AccessPointAttachmentOntapConfigurationFileSystemIdentityOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := f.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxS3AccessPointAttachmentOntapConfigurationFileSystemIdentityOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := f.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxS3AccessPointAttachmentOntapConfigurationFileSystemIdentityOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := f.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxS3AccessPointAttachmentOntapConfigurationFileSystemIdentityOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := f.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxS3AccessPointAttachmentOntapConfigurationFileSystemIdentityOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := f.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxS3AccessPointAttachmentOntapConfigurationFileSystemIdentityOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxS3AccessPointAttachmentOntapConfigurationFileSystemIdentityOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := f.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxS3AccessPointAttachmentOntapConfigurationFileSystemIdentityOutputReference) PutUnixUser(value *FsxS3AccessPointAttachmentOntapConfigurationFileSystemIdentityUnixUser) {
	if err := f.validatePutUnixUserParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putUnixUser",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FsxS3AccessPointAttachmentOntapConfigurationFileSystemIdentityOutputReference) PutWindowsUser(value *FsxS3AccessPointAttachmentOntapConfigurationFileSystemIdentityWindowsUser) {
	if err := f.validatePutWindowsUserParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putWindowsUser",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FsxS3AccessPointAttachmentOntapConfigurationFileSystemIdentityOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		f,
		"resetType",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxS3AccessPointAttachmentOntapConfigurationFileSystemIdentityOutputReference) ResetUnixUser() {
	_jsii_.InvokeVoid(
		f,
		"resetUnixUser",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxS3AccessPointAttachmentOntapConfigurationFileSystemIdentityOutputReference) ResetWindowsUser() {
	_jsii_.InvokeVoid(
		f,
		"resetWindowsUser",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxS3AccessPointAttachmentOntapConfigurationFileSystemIdentityOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := f.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxS3AccessPointAttachmentOntapConfigurationFileSystemIdentityOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

