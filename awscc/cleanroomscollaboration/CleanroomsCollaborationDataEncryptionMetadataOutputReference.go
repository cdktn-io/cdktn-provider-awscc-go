// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cleanroomscollaboration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/cleanroomscollaboration/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CleanroomsCollaborationDataEncryptionMetadataOutputReference interface {
	cdktn.ComplexObject
	AllowCleartext() interface{}
	SetAllowCleartext(val interface{})
	AllowCleartextInput() interface{}
	AllowDuplicates() interface{}
	SetAllowDuplicates(val interface{})
	AllowDuplicatesInput() interface{}
	AllowJoinsOnColumnsWithDifferentNames() interface{}
	SetAllowJoinsOnColumnsWithDifferentNames(val interface{})
	AllowJoinsOnColumnsWithDifferentNamesInput() interface{}
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
	PreserveNulls() interface{}
	SetPreserveNulls(val interface{})
	PreserveNullsInput() interface{}
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
	ResetAllowCleartext()
	ResetAllowDuplicates()
	ResetAllowJoinsOnColumnsWithDifferentNames()
	ResetPreserveNulls()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CleanroomsCollaborationDataEncryptionMetadataOutputReference
type jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference) AllowCleartext() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowCleartext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference) AllowCleartextInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowCleartextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference) AllowDuplicates() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowDuplicates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference) AllowDuplicatesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowDuplicatesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference) AllowJoinsOnColumnsWithDifferentNames() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowJoinsOnColumnsWithDifferentNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference) AllowJoinsOnColumnsWithDifferentNamesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowJoinsOnColumnsWithDifferentNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference) PreserveNulls() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preserveNulls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference) PreserveNullsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preserveNullsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCleanroomsCollaborationDataEncryptionMetadataOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) CleanroomsCollaborationDataEncryptionMetadataOutputReference {
	_init_.Initialize()

	if err := validateNewCleanroomsCollaborationDataEncryptionMetadataOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.cleanroomsCollaboration.CleanroomsCollaborationDataEncryptionMetadataOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCleanroomsCollaborationDataEncryptionMetadataOutputReference_Override(c CleanroomsCollaborationDataEncryptionMetadataOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.cleanroomsCollaboration.CleanroomsCollaborationDataEncryptionMetadataOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference)SetAllowCleartext(val interface{}) {
	if err := j.validateSetAllowCleartextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowCleartext",
		val,
	)
}

func (j *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference)SetAllowDuplicates(val interface{}) {
	if err := j.validateSetAllowDuplicatesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowDuplicates",
		val,
	)
}

func (j *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference)SetAllowJoinsOnColumnsWithDifferentNames(val interface{}) {
	if err := j.validateSetAllowJoinsOnColumnsWithDifferentNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowJoinsOnColumnsWithDifferentNames",
		val,
	)
}

func (j *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference)SetPreserveNulls(val interface{}) {
	if err := j.validateSetPreserveNullsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preserveNulls",
		val,
	)
}

func (j *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference) ResetAllowCleartext() {
	_jsii_.InvokeVoid(
		c,
		"resetAllowCleartext",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference) ResetAllowDuplicates() {
	_jsii_.InvokeVoid(
		c,
		"resetAllowDuplicates",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference) ResetAllowJoinsOnColumnsWithDifferentNames() {
	_jsii_.InvokeVoid(
		c,
		"resetAllowJoinsOnColumnsWithDifferentNames",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference) ResetPreserveNulls() {
	_jsii_.InvokeVoid(
		c,
		"resetPreserveNulls",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := c.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

