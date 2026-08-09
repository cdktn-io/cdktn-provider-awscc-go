// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package backuplegalhold

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/backuplegalhold/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BackupLegalHoldRecoveryPointSelectionOutputReference interface {
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
	DateRange() BackupLegalHoldRecoveryPointSelectionDateRangeOutputReference
	DateRangeInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ResourceIdentifiers() *[]*string
	SetResourceIdentifiers(val *[]*string)
	ResourceIdentifiersInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	VaultNames() *[]*string
	SetVaultNames(val *[]*string)
	VaultNamesInput() *[]*string
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
	PutDateRange(value *BackupLegalHoldRecoveryPointSelectionDateRange)
	ResetDateRange()
	ResetResourceIdentifiers()
	ResetVaultNames()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BackupLegalHoldRecoveryPointSelectionOutputReference
type jsiiProxy_BackupLegalHoldRecoveryPointSelectionOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_BackupLegalHoldRecoveryPointSelectionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupLegalHoldRecoveryPointSelectionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupLegalHoldRecoveryPointSelectionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupLegalHoldRecoveryPointSelectionOutputReference) DateRange() BackupLegalHoldRecoveryPointSelectionDateRangeOutputReference {
	var returns BackupLegalHoldRecoveryPointSelectionDateRangeOutputReference
	_jsii_.Get(
		j,
		"dateRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupLegalHoldRecoveryPointSelectionOutputReference) DateRangeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dateRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupLegalHoldRecoveryPointSelectionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupLegalHoldRecoveryPointSelectionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupLegalHoldRecoveryPointSelectionOutputReference) ResourceIdentifiers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceIdentifiers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupLegalHoldRecoveryPointSelectionOutputReference) ResourceIdentifiersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceIdentifiersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupLegalHoldRecoveryPointSelectionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupLegalHoldRecoveryPointSelectionOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupLegalHoldRecoveryPointSelectionOutputReference) VaultNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vaultNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupLegalHoldRecoveryPointSelectionOutputReference) VaultNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vaultNamesInput",
		&returns,
	)
	return returns
}


func NewBackupLegalHoldRecoveryPointSelectionOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) BackupLegalHoldRecoveryPointSelectionOutputReference {
	_init_.Initialize()

	if err := validateNewBackupLegalHoldRecoveryPointSelectionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BackupLegalHoldRecoveryPointSelectionOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.backupLegalHold.BackupLegalHoldRecoveryPointSelectionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBackupLegalHoldRecoveryPointSelectionOutputReference_Override(b BackupLegalHoldRecoveryPointSelectionOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.backupLegalHold.BackupLegalHoldRecoveryPointSelectionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BackupLegalHoldRecoveryPointSelectionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BackupLegalHoldRecoveryPointSelectionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BackupLegalHoldRecoveryPointSelectionOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BackupLegalHoldRecoveryPointSelectionOutputReference)SetResourceIdentifiers(val *[]*string) {
	if err := j.validateSetResourceIdentifiersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceIdentifiers",
		val,
	)
}

func (j *jsiiProxy_BackupLegalHoldRecoveryPointSelectionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BackupLegalHoldRecoveryPointSelectionOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BackupLegalHoldRecoveryPointSelectionOutputReference)SetVaultNames(val *[]*string) {
	if err := j.validateSetVaultNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vaultNames",
		val,
	)
}

func (b *jsiiProxy_BackupLegalHoldRecoveryPointSelectionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupLegalHoldRecoveryPointSelectionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupLegalHoldRecoveryPointSelectionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupLegalHoldRecoveryPointSelectionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupLegalHoldRecoveryPointSelectionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupLegalHoldRecoveryPointSelectionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupLegalHoldRecoveryPointSelectionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupLegalHoldRecoveryPointSelectionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupLegalHoldRecoveryPointSelectionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupLegalHoldRecoveryPointSelectionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupLegalHoldRecoveryPointSelectionOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupLegalHoldRecoveryPointSelectionOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupLegalHoldRecoveryPointSelectionOutputReference) PutDateRange(value *BackupLegalHoldRecoveryPointSelectionDateRange) {
	if err := b.validatePutDateRangeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putDateRange",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BackupLegalHoldRecoveryPointSelectionOutputReference) ResetDateRange() {
	_jsii_.InvokeVoid(
		b,
		"resetDateRange",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupLegalHoldRecoveryPointSelectionOutputReference) ResetResourceIdentifiers() {
	_jsii_.InvokeVoid(
		b,
		"resetResourceIdentifiers",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupLegalHoldRecoveryPointSelectionOutputReference) ResetVaultNames() {
	_jsii_.InvokeVoid(
		b,
		"resetVaultNames",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupLegalHoldRecoveryPointSelectionOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := b.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupLegalHoldRecoveryPointSelectionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

