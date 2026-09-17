// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package neptunegraphgraph

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/neptunegraphgraph/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type NeptunegraphGraphImportTaskOutputReference interface {
	cdktn.ComplexObject
	BlankNodeHandling() *string
	SetBlankNodeHandling(val *string)
	BlankNodeHandlingInput() *string
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
	FailOnError() interface{}
	SetFailOnError(val interface{})
	FailOnErrorInput() interface{}
	Format() *string
	SetFormat(val *string)
	FormatInput() *string
	// Experimental.
	Fqn() *string
	ImportOptions() NeptunegraphGraphImportTaskImportOptionsOutputReference
	ImportOptionsInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MaxProvisionedMemory() *float64
	SetMaxProvisionedMemory(val *float64)
	MaxProvisionedMemoryInput() *float64
	MinProvisionedMemory() *float64
	SetMinProvisionedMemory(val *float64)
	MinProvisionedMemoryInput() *float64
	ParquetType() *string
	SetParquetType(val *string)
	ParquetTypeInput() *string
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
	Source() *string
	SetSource(val *string)
	SourceInput() *string
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
	PutImportOptions(value *NeptunegraphGraphImportTaskImportOptions)
	ResetBlankNodeHandling()
	ResetFailOnError()
	ResetFormat()
	ResetImportOptions()
	ResetMaxProvisionedMemory()
	ResetMinProvisionedMemory()
	ResetParquetType()
	ResetRoleArn()
	ResetSource()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NeptunegraphGraphImportTaskOutputReference
type jsiiProxy_NeptunegraphGraphImportTaskOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_NeptunegraphGraphImportTaskOutputReference) BlankNodeHandling() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blankNodeHandling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraphImportTaskOutputReference) BlankNodeHandlingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blankNodeHandlingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraphImportTaskOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraphImportTaskOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraphImportTaskOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraphImportTaskOutputReference) FailOnError() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failOnError",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraphImportTaskOutputReference) FailOnErrorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failOnErrorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraphImportTaskOutputReference) Format() *string {
	var returns *string
	_jsii_.Get(
		j,
		"format",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraphImportTaskOutputReference) FormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"formatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraphImportTaskOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraphImportTaskOutputReference) ImportOptions() NeptunegraphGraphImportTaskImportOptionsOutputReference {
	var returns NeptunegraphGraphImportTaskImportOptionsOutputReference
	_jsii_.Get(
		j,
		"importOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraphImportTaskOutputReference) ImportOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"importOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraphImportTaskOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraphImportTaskOutputReference) MaxProvisionedMemory() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxProvisionedMemory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraphImportTaskOutputReference) MaxProvisionedMemoryInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxProvisionedMemoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraphImportTaskOutputReference) MinProvisionedMemory() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minProvisionedMemory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraphImportTaskOutputReference) MinProvisionedMemoryInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minProvisionedMemoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraphImportTaskOutputReference) ParquetType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parquetType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraphImportTaskOutputReference) ParquetTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parquetTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraphImportTaskOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraphImportTaskOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraphImportTaskOutputReference) Source() *string {
	var returns *string
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraphImportTaskOutputReference) SourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraphImportTaskOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraphImportTaskOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewNeptunegraphGraphImportTaskOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) NeptunegraphGraphImportTaskOutputReference {
	_init_.Initialize()

	if err := validateNewNeptunegraphGraphImportTaskOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_NeptunegraphGraphImportTaskOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.neptunegraphGraph.NeptunegraphGraphImportTaskOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewNeptunegraphGraphImportTaskOutputReference_Override(n NeptunegraphGraphImportTaskOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.neptunegraphGraph.NeptunegraphGraphImportTaskOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NeptunegraphGraphImportTaskOutputReference)SetBlankNodeHandling(val *string) {
	if err := j.validateSetBlankNodeHandlingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blankNodeHandling",
		val,
	)
}

func (j *jsiiProxy_NeptunegraphGraphImportTaskOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NeptunegraphGraphImportTaskOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NeptunegraphGraphImportTaskOutputReference)SetFailOnError(val interface{}) {
	if err := j.validateSetFailOnErrorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failOnError",
		val,
	)
}

func (j *jsiiProxy_NeptunegraphGraphImportTaskOutputReference)SetFormat(val *string) {
	if err := j.validateSetFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"format",
		val,
	)
}

func (j *jsiiProxy_NeptunegraphGraphImportTaskOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NeptunegraphGraphImportTaskOutputReference)SetMaxProvisionedMemory(val *float64) {
	if err := j.validateSetMaxProvisionedMemoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxProvisionedMemory",
		val,
	)
}

func (j *jsiiProxy_NeptunegraphGraphImportTaskOutputReference)SetMinProvisionedMemory(val *float64) {
	if err := j.validateSetMinProvisionedMemoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minProvisionedMemory",
		val,
	)
}

func (j *jsiiProxy_NeptunegraphGraphImportTaskOutputReference)SetParquetType(val *string) {
	if err := j.validateSetParquetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parquetType",
		val,
	)
}

func (j *jsiiProxy_NeptunegraphGraphImportTaskOutputReference)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_NeptunegraphGraphImportTaskOutputReference)SetSource(val *string) {
	if err := j.validateSetSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"source",
		val,
	)
}

func (j *jsiiProxy_NeptunegraphGraphImportTaskOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NeptunegraphGraphImportTaskOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (n *jsiiProxy_NeptunegraphGraphImportTaskOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NeptunegraphGraphImportTaskOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (n *jsiiProxy_NeptunegraphGraphImportTaskOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (n *jsiiProxy_NeptunegraphGraphImportTaskOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (n *jsiiProxy_NeptunegraphGraphImportTaskOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (n *jsiiProxy_NeptunegraphGraphImportTaskOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (n *jsiiProxy_NeptunegraphGraphImportTaskOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (n *jsiiProxy_NeptunegraphGraphImportTaskOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (n *jsiiProxy_NeptunegraphGraphImportTaskOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (n *jsiiProxy_NeptunegraphGraphImportTaskOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (n *jsiiProxy_NeptunegraphGraphImportTaskOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NeptunegraphGraphImportTaskOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (n *jsiiProxy_NeptunegraphGraphImportTaskOutputReference) PutImportOptions(value *NeptunegraphGraphImportTaskImportOptions) {
	if err := n.validatePutImportOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putImportOptions",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NeptunegraphGraphImportTaskOutputReference) ResetBlankNodeHandling() {
	_jsii_.InvokeVoid(
		n,
		"resetBlankNodeHandling",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NeptunegraphGraphImportTaskOutputReference) ResetFailOnError() {
	_jsii_.InvokeVoid(
		n,
		"resetFailOnError",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NeptunegraphGraphImportTaskOutputReference) ResetFormat() {
	_jsii_.InvokeVoid(
		n,
		"resetFormat",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NeptunegraphGraphImportTaskOutputReference) ResetImportOptions() {
	_jsii_.InvokeVoid(
		n,
		"resetImportOptions",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NeptunegraphGraphImportTaskOutputReference) ResetMaxProvisionedMemory() {
	_jsii_.InvokeVoid(
		n,
		"resetMaxProvisionedMemory",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NeptunegraphGraphImportTaskOutputReference) ResetMinProvisionedMemory() {
	_jsii_.InvokeVoid(
		n,
		"resetMinProvisionedMemory",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NeptunegraphGraphImportTaskOutputReference) ResetParquetType() {
	_jsii_.InvokeVoid(
		n,
		"resetParquetType",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NeptunegraphGraphImportTaskOutputReference) ResetRoleArn() {
	_jsii_.InvokeVoid(
		n,
		"resetRoleArn",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NeptunegraphGraphImportTaskOutputReference) ResetSource() {
	_jsii_.InvokeVoid(
		n,
		"resetSource",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NeptunegraphGraphImportTaskOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (n *jsiiProxy_NeptunegraphGraphImportTaskOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

