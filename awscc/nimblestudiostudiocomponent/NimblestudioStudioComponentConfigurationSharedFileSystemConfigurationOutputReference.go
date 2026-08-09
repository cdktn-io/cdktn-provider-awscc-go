// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package nimblestudiostudiocomponent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/nimblestudiostudiocomponent/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type NimblestudioStudioComponentConfigurationSharedFileSystemConfigurationOutputReference interface {
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
	Endpoint() *string
	SetEndpoint(val *string)
	EndpointInput() *string
	FileSystemId() *string
	SetFileSystemId(val *string)
	FileSystemIdInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LinuxMountPoint() *string
	SetLinuxMountPoint(val *string)
	LinuxMountPointInput() *string
	ShareName() *string
	SetShareName(val *string)
	ShareNameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	WindowsMountDrive() *string
	SetWindowsMountDrive(val *string)
	WindowsMountDriveInput() *string
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
	ResetEndpoint()
	ResetFileSystemId()
	ResetLinuxMountPoint()
	ResetShareName()
	ResetWindowsMountDrive()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NimblestudioStudioComponentConfigurationSharedFileSystemConfigurationOutputReference
type jsiiProxy_NimblestudioStudioComponentConfigurationSharedFileSystemConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_NimblestudioStudioComponentConfigurationSharedFileSystemConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioStudioComponentConfigurationSharedFileSystemConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioStudioComponentConfigurationSharedFileSystemConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioStudioComponentConfigurationSharedFileSystemConfigurationOutputReference) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioStudioComponentConfigurationSharedFileSystemConfigurationOutputReference) EndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioStudioComponentConfigurationSharedFileSystemConfigurationOutputReference) FileSystemId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileSystemId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioStudioComponentConfigurationSharedFileSystemConfigurationOutputReference) FileSystemIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileSystemIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioStudioComponentConfigurationSharedFileSystemConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioStudioComponentConfigurationSharedFileSystemConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioStudioComponentConfigurationSharedFileSystemConfigurationOutputReference) LinuxMountPoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"linuxMountPoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioStudioComponentConfigurationSharedFileSystemConfigurationOutputReference) LinuxMountPointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"linuxMountPointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioStudioComponentConfigurationSharedFileSystemConfigurationOutputReference) ShareName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioStudioComponentConfigurationSharedFileSystemConfigurationOutputReference) ShareNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioStudioComponentConfigurationSharedFileSystemConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioStudioComponentConfigurationSharedFileSystemConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioStudioComponentConfigurationSharedFileSystemConfigurationOutputReference) WindowsMountDrive() *string {
	var returns *string
	_jsii_.Get(
		j,
		"windowsMountDrive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioStudioComponentConfigurationSharedFileSystemConfigurationOutputReference) WindowsMountDriveInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"windowsMountDriveInput",
		&returns,
	)
	return returns
}


func NewNimblestudioStudioComponentConfigurationSharedFileSystemConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) NimblestudioStudioComponentConfigurationSharedFileSystemConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewNimblestudioStudioComponentConfigurationSharedFileSystemConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_NimblestudioStudioComponentConfigurationSharedFileSystemConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.nimblestudioStudioComponent.NimblestudioStudioComponentConfigurationSharedFileSystemConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewNimblestudioStudioComponentConfigurationSharedFileSystemConfigurationOutputReference_Override(n NimblestudioStudioComponentConfigurationSharedFileSystemConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.nimblestudioStudioComponent.NimblestudioStudioComponentConfigurationSharedFileSystemConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NimblestudioStudioComponentConfigurationSharedFileSystemConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NimblestudioStudioComponentConfigurationSharedFileSystemConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NimblestudioStudioComponentConfigurationSharedFileSystemConfigurationOutputReference)SetEndpoint(val *string) {
	if err := j.validateSetEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpoint",
		val,
	)
}

func (j *jsiiProxy_NimblestudioStudioComponentConfigurationSharedFileSystemConfigurationOutputReference)SetFileSystemId(val *string) {
	if err := j.validateSetFileSystemIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileSystemId",
		val,
	)
}

func (j *jsiiProxy_NimblestudioStudioComponentConfigurationSharedFileSystemConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NimblestudioStudioComponentConfigurationSharedFileSystemConfigurationOutputReference)SetLinuxMountPoint(val *string) {
	if err := j.validateSetLinuxMountPointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"linuxMountPoint",
		val,
	)
}

func (j *jsiiProxy_NimblestudioStudioComponentConfigurationSharedFileSystemConfigurationOutputReference)SetShareName(val *string) {
	if err := j.validateSetShareNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareName",
		val,
	)
}

func (j *jsiiProxy_NimblestudioStudioComponentConfigurationSharedFileSystemConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NimblestudioStudioComponentConfigurationSharedFileSystemConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_NimblestudioStudioComponentConfigurationSharedFileSystemConfigurationOutputReference)SetWindowsMountDrive(val *string) {
	if err := j.validateSetWindowsMountDriveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"windowsMountDrive",
		val,
	)
}

func (n *jsiiProxy_NimblestudioStudioComponentConfigurationSharedFileSystemConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NimblestudioStudioComponentConfigurationSharedFileSystemConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := n.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NimblestudioStudioComponentConfigurationSharedFileSystemConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := n.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NimblestudioStudioComponentConfigurationSharedFileSystemConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := n.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NimblestudioStudioComponentConfigurationSharedFileSystemConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := n.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NimblestudioStudioComponentConfigurationSharedFileSystemConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := n.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NimblestudioStudioComponentConfigurationSharedFileSystemConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := n.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NimblestudioStudioComponentConfigurationSharedFileSystemConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := n.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NimblestudioStudioComponentConfigurationSharedFileSystemConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := n.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NimblestudioStudioComponentConfigurationSharedFileSystemConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := n.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NimblestudioStudioComponentConfigurationSharedFileSystemConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NimblestudioStudioComponentConfigurationSharedFileSystemConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NimblestudioStudioComponentConfigurationSharedFileSystemConfigurationOutputReference) ResetEndpoint() {
	_jsii_.InvokeVoid(
		n,
		"resetEndpoint",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NimblestudioStudioComponentConfigurationSharedFileSystemConfigurationOutputReference) ResetFileSystemId() {
	_jsii_.InvokeVoid(
		n,
		"resetFileSystemId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NimblestudioStudioComponentConfigurationSharedFileSystemConfigurationOutputReference) ResetLinuxMountPoint() {
	_jsii_.InvokeVoid(
		n,
		"resetLinuxMountPoint",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NimblestudioStudioComponentConfigurationSharedFileSystemConfigurationOutputReference) ResetShareName() {
	_jsii_.InvokeVoid(
		n,
		"resetShareName",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NimblestudioStudioComponentConfigurationSharedFileSystemConfigurationOutputReference) ResetWindowsMountDrive() {
	_jsii_.InvokeVoid(
		n,
		"resetWindowsMountDrive",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NimblestudioStudioComponentConfigurationSharedFileSystemConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := n.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		n,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NimblestudioStudioComponentConfigurationSharedFileSystemConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

