// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package nimblestudiostudiocomponent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/nimblestudiostudiocomponent/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type NimblestudioStudioComponentConfigurationOutputReference interface {
	cdktn.ComplexObject
	ActiveDirectoryConfiguration() NimblestudioStudioComponentConfigurationActiveDirectoryConfigurationOutputReference
	ActiveDirectoryConfigurationInput() interface{}
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
	ComputeFarmConfiguration() NimblestudioStudioComponentConfigurationComputeFarmConfigurationOutputReference
	ComputeFarmConfigurationInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LicenseServiceConfiguration() NimblestudioStudioComponentConfigurationLicenseServiceConfigurationOutputReference
	LicenseServiceConfigurationInput() interface{}
	SharedFileSystemConfiguration() NimblestudioStudioComponentConfigurationSharedFileSystemConfigurationOutputReference
	SharedFileSystemConfigurationInput() interface{}
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
	PutActiveDirectoryConfiguration(value *NimblestudioStudioComponentConfigurationActiveDirectoryConfiguration)
	PutComputeFarmConfiguration(value *NimblestudioStudioComponentConfigurationComputeFarmConfiguration)
	PutLicenseServiceConfiguration(value *NimblestudioStudioComponentConfigurationLicenseServiceConfiguration)
	PutSharedFileSystemConfiguration(value *NimblestudioStudioComponentConfigurationSharedFileSystemConfiguration)
	ResetActiveDirectoryConfiguration()
	ResetComputeFarmConfiguration()
	ResetLicenseServiceConfiguration()
	ResetSharedFileSystemConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NimblestudioStudioComponentConfigurationOutputReference
type jsiiProxy_NimblestudioStudioComponentConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_NimblestudioStudioComponentConfigurationOutputReference) ActiveDirectoryConfiguration() NimblestudioStudioComponentConfigurationActiveDirectoryConfigurationOutputReference {
	var returns NimblestudioStudioComponentConfigurationActiveDirectoryConfigurationOutputReference
	_jsii_.Get(
		j,
		"activeDirectoryConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioStudioComponentConfigurationOutputReference) ActiveDirectoryConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"activeDirectoryConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioStudioComponentConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioStudioComponentConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioStudioComponentConfigurationOutputReference) ComputeFarmConfiguration() NimblestudioStudioComponentConfigurationComputeFarmConfigurationOutputReference {
	var returns NimblestudioStudioComponentConfigurationComputeFarmConfigurationOutputReference
	_jsii_.Get(
		j,
		"computeFarmConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioStudioComponentConfigurationOutputReference) ComputeFarmConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"computeFarmConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioStudioComponentConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioStudioComponentConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioStudioComponentConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioStudioComponentConfigurationOutputReference) LicenseServiceConfiguration() NimblestudioStudioComponentConfigurationLicenseServiceConfigurationOutputReference {
	var returns NimblestudioStudioComponentConfigurationLicenseServiceConfigurationOutputReference
	_jsii_.Get(
		j,
		"licenseServiceConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioStudioComponentConfigurationOutputReference) LicenseServiceConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"licenseServiceConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioStudioComponentConfigurationOutputReference) SharedFileSystemConfiguration() NimblestudioStudioComponentConfigurationSharedFileSystemConfigurationOutputReference {
	var returns NimblestudioStudioComponentConfigurationSharedFileSystemConfigurationOutputReference
	_jsii_.Get(
		j,
		"sharedFileSystemConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioStudioComponentConfigurationOutputReference) SharedFileSystemConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sharedFileSystemConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioStudioComponentConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioStudioComponentConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewNimblestudioStudioComponentConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) NimblestudioStudioComponentConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewNimblestudioStudioComponentConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_NimblestudioStudioComponentConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.nimblestudioStudioComponent.NimblestudioStudioComponentConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewNimblestudioStudioComponentConfigurationOutputReference_Override(n NimblestudioStudioComponentConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.nimblestudioStudioComponent.NimblestudioStudioComponentConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NimblestudioStudioComponentConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NimblestudioStudioComponentConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NimblestudioStudioComponentConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NimblestudioStudioComponentConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NimblestudioStudioComponentConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (n *jsiiProxy_NimblestudioStudioComponentConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NimblestudioStudioComponentConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (n *jsiiProxy_NimblestudioStudioComponentConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (n *jsiiProxy_NimblestudioStudioComponentConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (n *jsiiProxy_NimblestudioStudioComponentConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (n *jsiiProxy_NimblestudioStudioComponentConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (n *jsiiProxy_NimblestudioStudioComponentConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (n *jsiiProxy_NimblestudioStudioComponentConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (n *jsiiProxy_NimblestudioStudioComponentConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (n *jsiiProxy_NimblestudioStudioComponentConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (n *jsiiProxy_NimblestudioStudioComponentConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NimblestudioStudioComponentConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (n *jsiiProxy_NimblestudioStudioComponentConfigurationOutputReference) PutActiveDirectoryConfiguration(value *NimblestudioStudioComponentConfigurationActiveDirectoryConfiguration) {
	if err := n.validatePutActiveDirectoryConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putActiveDirectoryConfiguration",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NimblestudioStudioComponentConfigurationOutputReference) PutComputeFarmConfiguration(value *NimblestudioStudioComponentConfigurationComputeFarmConfiguration) {
	if err := n.validatePutComputeFarmConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putComputeFarmConfiguration",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NimblestudioStudioComponentConfigurationOutputReference) PutLicenseServiceConfiguration(value *NimblestudioStudioComponentConfigurationLicenseServiceConfiguration) {
	if err := n.validatePutLicenseServiceConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putLicenseServiceConfiguration",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NimblestudioStudioComponentConfigurationOutputReference) PutSharedFileSystemConfiguration(value *NimblestudioStudioComponentConfigurationSharedFileSystemConfiguration) {
	if err := n.validatePutSharedFileSystemConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putSharedFileSystemConfiguration",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NimblestudioStudioComponentConfigurationOutputReference) ResetActiveDirectoryConfiguration() {
	_jsii_.InvokeVoid(
		n,
		"resetActiveDirectoryConfiguration",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NimblestudioStudioComponentConfigurationOutputReference) ResetComputeFarmConfiguration() {
	_jsii_.InvokeVoid(
		n,
		"resetComputeFarmConfiguration",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NimblestudioStudioComponentConfigurationOutputReference) ResetLicenseServiceConfiguration() {
	_jsii_.InvokeVoid(
		n,
		"resetLicenseServiceConfiguration",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NimblestudioStudioComponentConfigurationOutputReference) ResetSharedFileSystemConfiguration() {
	_jsii_.InvokeVoid(
		n,
		"resetSharedFileSystemConfiguration",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NimblestudioStudioComponentConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (n *jsiiProxy_NimblestudioStudioComponentConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

