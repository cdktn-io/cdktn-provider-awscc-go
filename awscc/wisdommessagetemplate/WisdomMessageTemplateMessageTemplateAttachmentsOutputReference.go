// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wisdommessagetemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/wisdommessagetemplate/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type WisdomMessageTemplateMessageTemplateAttachmentsOutputReference interface {
	cdktn.ComplexObject
	AttachmentId() *string
	SetAttachmentId(val *string)
	AttachmentIdInput() *string
	AttachmentName() *string
	SetAttachmentName(val *string)
	AttachmentNameInput() *string
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
	S3PresignedUrl() *string
	SetS3PresignedUrl(val *string)
	S3PresignedUrlInput() *string
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
	ResetAttachmentId()
	ResetAttachmentName()
	ResetS3PresignedUrl()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WisdomMessageTemplateMessageTemplateAttachmentsOutputReference
type jsiiProxy_WisdomMessageTemplateMessageTemplateAttachmentsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_WisdomMessageTemplateMessageTemplateAttachmentsOutputReference) AttachmentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attachmentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateMessageTemplateAttachmentsOutputReference) AttachmentIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attachmentIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateMessageTemplateAttachmentsOutputReference) AttachmentName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attachmentName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateMessageTemplateAttachmentsOutputReference) AttachmentNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attachmentNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateMessageTemplateAttachmentsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateMessageTemplateAttachmentsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateMessageTemplateAttachmentsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateMessageTemplateAttachmentsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateMessageTemplateAttachmentsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateMessageTemplateAttachmentsOutputReference) S3PresignedUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3PresignedUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateMessageTemplateAttachmentsOutputReference) S3PresignedUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3PresignedUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateMessageTemplateAttachmentsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateMessageTemplateAttachmentsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWisdomMessageTemplateMessageTemplateAttachmentsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) WisdomMessageTemplateMessageTemplateAttachmentsOutputReference {
	_init_.Initialize()

	if err := validateNewWisdomMessageTemplateMessageTemplateAttachmentsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_WisdomMessageTemplateMessageTemplateAttachmentsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.wisdomMessageTemplate.WisdomMessageTemplateMessageTemplateAttachmentsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewWisdomMessageTemplateMessageTemplateAttachmentsOutputReference_Override(w WisdomMessageTemplateMessageTemplateAttachmentsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.wisdomMessageTemplate.WisdomMessageTemplateMessageTemplateAttachmentsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		w,
	)
}

func (j *jsiiProxy_WisdomMessageTemplateMessageTemplateAttachmentsOutputReference)SetAttachmentId(val *string) {
	if err := j.validateSetAttachmentIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"attachmentId",
		val,
	)
}

func (j *jsiiProxy_WisdomMessageTemplateMessageTemplateAttachmentsOutputReference)SetAttachmentName(val *string) {
	if err := j.validateSetAttachmentNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"attachmentName",
		val,
	)
}

func (j *jsiiProxy_WisdomMessageTemplateMessageTemplateAttachmentsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WisdomMessageTemplateMessageTemplateAttachmentsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WisdomMessageTemplateMessageTemplateAttachmentsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WisdomMessageTemplateMessageTemplateAttachmentsOutputReference)SetS3PresignedUrl(val *string) {
	if err := j.validateSetS3PresignedUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3PresignedUrl",
		val,
	)
}

func (j *jsiiProxy_WisdomMessageTemplateMessageTemplateAttachmentsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WisdomMessageTemplateMessageTemplateAttachmentsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WisdomMessageTemplateMessageTemplateAttachmentsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomMessageTemplateMessageTemplateAttachmentsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := w.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomMessageTemplateMessageTemplateAttachmentsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := w.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomMessageTemplateMessageTemplateAttachmentsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := w.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomMessageTemplateMessageTemplateAttachmentsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := w.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomMessageTemplateMessageTemplateAttachmentsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := w.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomMessageTemplateMessageTemplateAttachmentsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := w.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomMessageTemplateMessageTemplateAttachmentsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := w.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomMessageTemplateMessageTemplateAttachmentsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := w.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomMessageTemplateMessageTemplateAttachmentsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := w.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomMessageTemplateMessageTemplateAttachmentsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomMessageTemplateMessageTemplateAttachmentsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomMessageTemplateMessageTemplateAttachmentsOutputReference) ResetAttachmentId() {
	_jsii_.InvokeVoid(
		w,
		"resetAttachmentId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomMessageTemplateMessageTemplateAttachmentsOutputReference) ResetAttachmentName() {
	_jsii_.InvokeVoid(
		w,
		"resetAttachmentName",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomMessageTemplateMessageTemplateAttachmentsOutputReference) ResetS3PresignedUrl() {
	_jsii_.InvokeVoid(
		w,
		"resetS3PresignedUrl",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomMessageTemplateMessageTemplateAttachmentsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := w.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomMessageTemplateMessageTemplateAttachmentsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

