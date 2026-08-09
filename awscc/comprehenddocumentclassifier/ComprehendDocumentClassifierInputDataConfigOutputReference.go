// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package comprehenddocumentclassifier

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/comprehenddocumentclassifier/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ComprehendDocumentClassifierInputDataConfigOutputReference interface {
	cdktn.ComplexObject
	AugmentedManifests() ComprehendDocumentClassifierInputDataConfigAugmentedManifestsList
	AugmentedManifestsInput() interface{}
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
	DataFormat() *string
	SetDataFormat(val *string)
	DataFormatInput() *string
	DocumentReaderConfig() ComprehendDocumentClassifierInputDataConfigDocumentReaderConfigOutputReference
	DocumentReaderConfigInput() interface{}
	Documents() ComprehendDocumentClassifierInputDataConfigDocumentsOutputReference
	DocumentsInput() interface{}
	DocumentType() *string
	SetDocumentType(val *string)
	DocumentTypeInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LabelDelimiter() *string
	SetLabelDelimiter(val *string)
	LabelDelimiterInput() *string
	S3Uri() *string
	SetS3Uri(val *string)
	S3UriInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TestS3Uri() *string
	SetTestS3Uri(val *string)
	TestS3UriInput() *string
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
	PutAugmentedManifests(value interface{})
	PutDocumentReaderConfig(value *ComprehendDocumentClassifierInputDataConfigDocumentReaderConfig)
	PutDocuments(value *ComprehendDocumentClassifierInputDataConfigDocuments)
	ResetAugmentedManifests()
	ResetDataFormat()
	ResetDocumentReaderConfig()
	ResetDocuments()
	ResetDocumentType()
	ResetLabelDelimiter()
	ResetS3Uri()
	ResetTestS3Uri()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComprehendDocumentClassifierInputDataConfigOutputReference
type jsiiProxy_ComprehendDocumentClassifierInputDataConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigOutputReference) AugmentedManifests() ComprehendDocumentClassifierInputDataConfigAugmentedManifestsList {
	var returns ComprehendDocumentClassifierInputDataConfigAugmentedManifestsList
	_jsii_.Get(
		j,
		"augmentedManifests",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigOutputReference) AugmentedManifestsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"augmentedManifestsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigOutputReference) DataFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigOutputReference) DataFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigOutputReference) DocumentReaderConfig() ComprehendDocumentClassifierInputDataConfigDocumentReaderConfigOutputReference {
	var returns ComprehendDocumentClassifierInputDataConfigDocumentReaderConfigOutputReference
	_jsii_.Get(
		j,
		"documentReaderConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigOutputReference) DocumentReaderConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"documentReaderConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigOutputReference) Documents() ComprehendDocumentClassifierInputDataConfigDocumentsOutputReference {
	var returns ComprehendDocumentClassifierInputDataConfigDocumentsOutputReference
	_jsii_.Get(
		j,
		"documents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigOutputReference) DocumentsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"documentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigOutputReference) DocumentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigOutputReference) DocumentTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigOutputReference) LabelDelimiter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelDelimiter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigOutputReference) LabelDelimiterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelDelimiterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigOutputReference) S3Uri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3Uri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigOutputReference) S3UriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3UriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigOutputReference) TestS3Uri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"testS3Uri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigOutputReference) TestS3UriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"testS3UriInput",
		&returns,
	)
	return returns
}


func NewComprehendDocumentClassifierInputDataConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ComprehendDocumentClassifierInputDataConfigOutputReference {
	_init_.Initialize()

	if err := validateNewComprehendDocumentClassifierInputDataConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComprehendDocumentClassifierInputDataConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.comprehendDocumentClassifier.ComprehendDocumentClassifierInputDataConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewComprehendDocumentClassifierInputDataConfigOutputReference_Override(c ComprehendDocumentClassifierInputDataConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.comprehendDocumentClassifier.ComprehendDocumentClassifierInputDataConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigOutputReference)SetDataFormat(val *string) {
	if err := j.validateSetDataFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataFormat",
		val,
	)
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigOutputReference)SetDocumentType(val *string) {
	if err := j.validateSetDocumentTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"documentType",
		val,
	)
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigOutputReference)SetLabelDelimiter(val *string) {
	if err := j.validateSetLabelDelimiterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labelDelimiter",
		val,
	)
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigOutputReference)SetS3Uri(val *string) {
	if err := j.validateSetS3UriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3Uri",
		val,
	)
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigOutputReference)SetTestS3Uri(val *string) {
	if err := j.validateSetTestS3UriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"testS3Uri",
		val,
	)
}

func (c *jsiiProxy_ComprehendDocumentClassifierInputDataConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComprehendDocumentClassifierInputDataConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComprehendDocumentClassifierInputDataConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ComprehendDocumentClassifierInputDataConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComprehendDocumentClassifierInputDataConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComprehendDocumentClassifierInputDataConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComprehendDocumentClassifierInputDataConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComprehendDocumentClassifierInputDataConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComprehendDocumentClassifierInputDataConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComprehendDocumentClassifierInputDataConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComprehendDocumentClassifierInputDataConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComprehendDocumentClassifierInputDataConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ComprehendDocumentClassifierInputDataConfigOutputReference) PutAugmentedManifests(value interface{}) {
	if err := c.validatePutAugmentedManifestsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAugmentedManifests",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComprehendDocumentClassifierInputDataConfigOutputReference) PutDocumentReaderConfig(value *ComprehendDocumentClassifierInputDataConfigDocumentReaderConfig) {
	if err := c.validatePutDocumentReaderConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDocumentReaderConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComprehendDocumentClassifierInputDataConfigOutputReference) PutDocuments(value *ComprehendDocumentClassifierInputDataConfigDocuments) {
	if err := c.validatePutDocumentsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDocuments",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComprehendDocumentClassifierInputDataConfigOutputReference) ResetAugmentedManifests() {
	_jsii_.InvokeVoid(
		c,
		"resetAugmentedManifests",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComprehendDocumentClassifierInputDataConfigOutputReference) ResetDataFormat() {
	_jsii_.InvokeVoid(
		c,
		"resetDataFormat",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComprehendDocumentClassifierInputDataConfigOutputReference) ResetDocumentReaderConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetDocumentReaderConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComprehendDocumentClassifierInputDataConfigOutputReference) ResetDocuments() {
	_jsii_.InvokeVoid(
		c,
		"resetDocuments",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComprehendDocumentClassifierInputDataConfigOutputReference) ResetDocumentType() {
	_jsii_.InvokeVoid(
		c,
		"resetDocumentType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComprehendDocumentClassifierInputDataConfigOutputReference) ResetLabelDelimiter() {
	_jsii_.InvokeVoid(
		c,
		"resetLabelDelimiter",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComprehendDocumentClassifierInputDataConfigOutputReference) ResetS3Uri() {
	_jsii_.InvokeVoid(
		c,
		"resetS3Uri",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComprehendDocumentClassifierInputDataConfigOutputReference) ResetTestS3Uri() {
	_jsii_.InvokeVoid(
		c,
		"resetTestS3Uri",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComprehendDocumentClassifierInputDataConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_ComprehendDocumentClassifierInputDataConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

