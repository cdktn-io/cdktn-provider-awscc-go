// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package fsxvolume

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/fsxvolume/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type FsxVolumeOntapConfigurationSnaplockConfigurationOutputReference interface {
	cdktn.ComplexObject
	AuditLogVolume() *string
	SetAuditLogVolume(val *string)
	AuditLogVolumeInput() *string
	AutocommitPeriod() FsxVolumeOntapConfigurationSnaplockConfigurationAutocommitPeriodOutputReference
	AutocommitPeriodInput() interface{}
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
	PrivilegedDelete() *string
	SetPrivilegedDelete(val *string)
	PrivilegedDeleteInput() *string
	RetentionPeriod() FsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriodOutputReference
	RetentionPeriodInput() interface{}
	SnaplockType() *string
	SetSnaplockType(val *string)
	SnaplockTypeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	VolumeAppendModeEnabled() *string
	SetVolumeAppendModeEnabled(val *string)
	VolumeAppendModeEnabledInput() *string
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
	PutAutocommitPeriod(value *FsxVolumeOntapConfigurationSnaplockConfigurationAutocommitPeriod)
	PutRetentionPeriod(value *FsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriod)
	ResetAuditLogVolume()
	ResetAutocommitPeriod()
	ResetPrivilegedDelete()
	ResetRetentionPeriod()
	ResetSnaplockType()
	ResetVolumeAppendModeEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FsxVolumeOntapConfigurationSnaplockConfigurationOutputReference
type jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationOutputReference) AuditLogVolume() *string {
	var returns *string
	_jsii_.Get(
		j,
		"auditLogVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationOutputReference) AuditLogVolumeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"auditLogVolumeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationOutputReference) AutocommitPeriod() FsxVolumeOntapConfigurationSnaplockConfigurationAutocommitPeriodOutputReference {
	var returns FsxVolumeOntapConfigurationSnaplockConfigurationAutocommitPeriodOutputReference
	_jsii_.Get(
		j,
		"autocommitPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationOutputReference) AutocommitPeriodInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autocommitPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationOutputReference) PrivilegedDelete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privilegedDelete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationOutputReference) PrivilegedDeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privilegedDeleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationOutputReference) RetentionPeriod() FsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriodOutputReference {
	var returns FsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriodOutputReference
	_jsii_.Get(
		j,
		"retentionPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationOutputReference) RetentionPeriodInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retentionPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationOutputReference) SnaplockType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snaplockType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationOutputReference) SnaplockTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snaplockTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationOutputReference) VolumeAppendModeEnabled() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeAppendModeEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationOutputReference) VolumeAppendModeEnabledInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeAppendModeEnabledInput",
		&returns,
	)
	return returns
}


func NewFsxVolumeOntapConfigurationSnaplockConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) FsxVolumeOntapConfigurationSnaplockConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewFsxVolumeOntapConfigurationSnaplockConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.fsxVolume.FsxVolumeOntapConfigurationSnaplockConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewFsxVolumeOntapConfigurationSnaplockConfigurationOutputReference_Override(f FsxVolumeOntapConfigurationSnaplockConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.fsxVolume.FsxVolumeOntapConfigurationSnaplockConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		f,
	)
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationOutputReference)SetAuditLogVolume(val *string) {
	if err := j.validateSetAuditLogVolumeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"auditLogVolume",
		val,
	)
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationOutputReference)SetPrivilegedDelete(val *string) {
	if err := j.validateSetPrivilegedDeleteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privilegedDelete",
		val,
	)
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationOutputReference)SetSnaplockType(val *string) {
	if err := j.validateSetSnaplockTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snaplockType",
		val,
	)
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationOutputReference)SetVolumeAppendModeEnabled(val *string) {
	if err := j.validateSetVolumeAppendModeEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumeAppendModeEnabled",
		val,
	)
}

func (f *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (f *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (f *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (f *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (f *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (f *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (f *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (f *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (f *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (f *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (f *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationOutputReference) PutAutocommitPeriod(value *FsxVolumeOntapConfigurationSnaplockConfigurationAutocommitPeriod) {
	if err := f.validatePutAutocommitPeriodParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putAutocommitPeriod",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationOutputReference) PutRetentionPeriod(value *FsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriod) {
	if err := f.validatePutRetentionPeriodParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putRetentionPeriod",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationOutputReference) ResetAuditLogVolume() {
	_jsii_.InvokeVoid(
		f,
		"resetAuditLogVolume",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationOutputReference) ResetAutocommitPeriod() {
	_jsii_.InvokeVoid(
		f,
		"resetAutocommitPeriod",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationOutputReference) ResetPrivilegedDelete() {
	_jsii_.InvokeVoid(
		f,
		"resetPrivilegedDelete",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationOutputReference) ResetRetentionPeriod() {
	_jsii_.InvokeVoid(
		f,
		"resetRetentionPeriod",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationOutputReference) ResetSnaplockType() {
	_jsii_.InvokeVoid(
		f,
		"resetSnaplockType",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationOutputReference) ResetVolumeAppendModeEnabled() {
	_jsii_.InvokeVoid(
		f,
		"resetVolumeAppendModeEnabled",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (f *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

