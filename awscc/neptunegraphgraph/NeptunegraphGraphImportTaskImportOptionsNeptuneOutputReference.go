// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package neptunegraphgraph

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/neptunegraphgraph/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type NeptunegraphGraphImportTaskImportOptionsNeptuneOutputReference interface {
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
	PreserveDefaultVertexLabels() interface{}
	SetPreserveDefaultVertexLabels(val interface{})
	PreserveDefaultVertexLabelsInput() interface{}
	PreserveEdgeIds() interface{}
	SetPreserveEdgeIds(val interface{})
	PreserveEdgeIdsInput() interface{}
	S3ExportKmsKeyId() *string
	SetS3ExportKmsKeyId(val *string)
	S3ExportKmsKeyIdInput() *string
	S3ExportPath() *string
	SetS3ExportPath(val *string)
	S3ExportPathInput() *string
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
	ResetPreserveDefaultVertexLabels()
	ResetPreserveEdgeIds()
	ResetS3ExportKmsKeyId()
	ResetS3ExportPath()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NeptunegraphGraphImportTaskImportOptionsNeptuneOutputReference
type jsiiProxy_NeptunegraphGraphImportTaskImportOptionsNeptuneOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_NeptunegraphGraphImportTaskImportOptionsNeptuneOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraphImportTaskImportOptionsNeptuneOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraphImportTaskImportOptionsNeptuneOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraphImportTaskImportOptionsNeptuneOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraphImportTaskImportOptionsNeptuneOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraphImportTaskImportOptionsNeptuneOutputReference) PreserveDefaultVertexLabels() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preserveDefaultVertexLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraphImportTaskImportOptionsNeptuneOutputReference) PreserveDefaultVertexLabelsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preserveDefaultVertexLabelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraphImportTaskImportOptionsNeptuneOutputReference) PreserveEdgeIds() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preserveEdgeIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraphImportTaskImportOptionsNeptuneOutputReference) PreserveEdgeIdsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preserveEdgeIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraphImportTaskImportOptionsNeptuneOutputReference) S3ExportKmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3ExportKmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraphImportTaskImportOptionsNeptuneOutputReference) S3ExportKmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3ExportKmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraphImportTaskImportOptionsNeptuneOutputReference) S3ExportPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3ExportPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraphImportTaskImportOptionsNeptuneOutputReference) S3ExportPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3ExportPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraphImportTaskImportOptionsNeptuneOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraphImportTaskImportOptionsNeptuneOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewNeptunegraphGraphImportTaskImportOptionsNeptuneOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) NeptunegraphGraphImportTaskImportOptionsNeptuneOutputReference {
	_init_.Initialize()

	if err := validateNewNeptunegraphGraphImportTaskImportOptionsNeptuneOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_NeptunegraphGraphImportTaskImportOptionsNeptuneOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.neptunegraphGraph.NeptunegraphGraphImportTaskImportOptionsNeptuneOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewNeptunegraphGraphImportTaskImportOptionsNeptuneOutputReference_Override(n NeptunegraphGraphImportTaskImportOptionsNeptuneOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.neptunegraphGraph.NeptunegraphGraphImportTaskImportOptionsNeptuneOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NeptunegraphGraphImportTaskImportOptionsNeptuneOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NeptunegraphGraphImportTaskImportOptionsNeptuneOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NeptunegraphGraphImportTaskImportOptionsNeptuneOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NeptunegraphGraphImportTaskImportOptionsNeptuneOutputReference)SetPreserveDefaultVertexLabels(val interface{}) {
	if err := j.validateSetPreserveDefaultVertexLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preserveDefaultVertexLabels",
		val,
	)
}

func (j *jsiiProxy_NeptunegraphGraphImportTaskImportOptionsNeptuneOutputReference)SetPreserveEdgeIds(val interface{}) {
	if err := j.validateSetPreserveEdgeIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preserveEdgeIds",
		val,
	)
}

func (j *jsiiProxy_NeptunegraphGraphImportTaskImportOptionsNeptuneOutputReference)SetS3ExportKmsKeyId(val *string) {
	if err := j.validateSetS3ExportKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3ExportKmsKeyId",
		val,
	)
}

func (j *jsiiProxy_NeptunegraphGraphImportTaskImportOptionsNeptuneOutputReference)SetS3ExportPath(val *string) {
	if err := j.validateSetS3ExportPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3ExportPath",
		val,
	)
}

func (j *jsiiProxy_NeptunegraphGraphImportTaskImportOptionsNeptuneOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NeptunegraphGraphImportTaskImportOptionsNeptuneOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (n *jsiiProxy_NeptunegraphGraphImportTaskImportOptionsNeptuneOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NeptunegraphGraphImportTaskImportOptionsNeptuneOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (n *jsiiProxy_NeptunegraphGraphImportTaskImportOptionsNeptuneOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (n *jsiiProxy_NeptunegraphGraphImportTaskImportOptionsNeptuneOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (n *jsiiProxy_NeptunegraphGraphImportTaskImportOptionsNeptuneOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (n *jsiiProxy_NeptunegraphGraphImportTaskImportOptionsNeptuneOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (n *jsiiProxy_NeptunegraphGraphImportTaskImportOptionsNeptuneOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (n *jsiiProxy_NeptunegraphGraphImportTaskImportOptionsNeptuneOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (n *jsiiProxy_NeptunegraphGraphImportTaskImportOptionsNeptuneOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (n *jsiiProxy_NeptunegraphGraphImportTaskImportOptionsNeptuneOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (n *jsiiProxy_NeptunegraphGraphImportTaskImportOptionsNeptuneOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NeptunegraphGraphImportTaskImportOptionsNeptuneOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (n *jsiiProxy_NeptunegraphGraphImportTaskImportOptionsNeptuneOutputReference) ResetPreserveDefaultVertexLabels() {
	_jsii_.InvokeVoid(
		n,
		"resetPreserveDefaultVertexLabels",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NeptunegraphGraphImportTaskImportOptionsNeptuneOutputReference) ResetPreserveEdgeIds() {
	_jsii_.InvokeVoid(
		n,
		"resetPreserveEdgeIds",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NeptunegraphGraphImportTaskImportOptionsNeptuneOutputReference) ResetS3ExportKmsKeyId() {
	_jsii_.InvokeVoid(
		n,
		"resetS3ExportKmsKeyId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NeptunegraphGraphImportTaskImportOptionsNeptuneOutputReference) ResetS3ExportPath() {
	_jsii_.InvokeVoid(
		n,
		"resetS3ExportPath",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NeptunegraphGraphImportTaskImportOptionsNeptuneOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (n *jsiiProxy_NeptunegraphGraphImportTaskImportOptionsNeptuneOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

