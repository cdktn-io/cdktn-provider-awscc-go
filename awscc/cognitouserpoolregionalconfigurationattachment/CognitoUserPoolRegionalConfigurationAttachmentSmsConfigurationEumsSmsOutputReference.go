// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cognitouserpoolregionalconfigurationattachment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/cognitouserpoolregionalconfigurationattachment/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationEumsSmsOutputReference interface {
	cdktn.ComplexObject
	CallerArn() *string
	SetCallerArn(val *string)
	CallerArnInput() *string
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
	ConfigurationSetName() *string
	SetConfigurationSetName(val *string)
	ConfigurationSetNameInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	ExternalId() *string
	SetExternalId(val *string)
	ExternalIdInput() *string
	// Experimental.
	Fqn() *string
	InEntityId() *string
	SetInEntityId(val *string)
	InEntityIdInput() *string
	InTemplateId() *string
	SetInTemplateId(val *string)
	InTemplateIdInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	OriginationIdentity() *string
	SetOriginationIdentity(val *string)
	OriginationIdentityInput() *string
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
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
	ResetCallerArn()
	ResetConfigurationSetName()
	ResetExternalId()
	ResetInEntityId()
	ResetInTemplateId()
	ResetOriginationIdentity()
	ResetRegion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationEumsSmsOutputReference
type jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationEumsSmsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationEumsSmsOutputReference) CallerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"callerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationEumsSmsOutputReference) CallerArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"callerArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationEumsSmsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationEumsSmsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationEumsSmsOutputReference) ConfigurationSetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationSetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationEumsSmsOutputReference) ConfigurationSetNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationSetNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationEumsSmsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationEumsSmsOutputReference) ExternalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationEumsSmsOutputReference) ExternalIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationEumsSmsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationEumsSmsOutputReference) InEntityId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inEntityId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationEumsSmsOutputReference) InEntityIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inEntityIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationEumsSmsOutputReference) InTemplateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inTemplateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationEumsSmsOutputReference) InTemplateIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inTemplateIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationEumsSmsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationEumsSmsOutputReference) OriginationIdentity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originationIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationEumsSmsOutputReference) OriginationIdentityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originationIdentityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationEumsSmsOutputReference) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationEumsSmsOutputReference) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationEumsSmsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationEumsSmsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationEumsSmsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationEumsSmsOutputReference {
	_init_.Initialize()

	if err := validateNewCognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationEumsSmsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationEumsSmsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.cognitoUserPoolRegionalConfigurationAttachment.CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationEumsSmsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationEumsSmsOutputReference_Override(c CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationEumsSmsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.cognitoUserPoolRegionalConfigurationAttachment.CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationEumsSmsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationEumsSmsOutputReference)SetCallerArn(val *string) {
	if err := j.validateSetCallerArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"callerArn",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationEumsSmsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationEumsSmsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationEumsSmsOutputReference)SetConfigurationSetName(val *string) {
	if err := j.validateSetConfigurationSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configurationSetName",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationEumsSmsOutputReference)SetExternalId(val *string) {
	if err := j.validateSetExternalIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalId",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationEumsSmsOutputReference)SetInEntityId(val *string) {
	if err := j.validateSetInEntityIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inEntityId",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationEumsSmsOutputReference)SetInTemplateId(val *string) {
	if err := j.validateSetInTemplateIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inTemplateId",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationEumsSmsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationEumsSmsOutputReference)SetOriginationIdentity(val *string) {
	if err := j.validateSetOriginationIdentityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"originationIdentity",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationEumsSmsOutputReference)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationEumsSmsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationEumsSmsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationEumsSmsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationEumsSmsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationEumsSmsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationEumsSmsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationEumsSmsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationEumsSmsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationEumsSmsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationEumsSmsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationEumsSmsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationEumsSmsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationEumsSmsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationEumsSmsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationEumsSmsOutputReference) ResetCallerArn() {
	_jsii_.InvokeVoid(
		c,
		"resetCallerArn",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationEumsSmsOutputReference) ResetConfigurationSetName() {
	_jsii_.InvokeVoid(
		c,
		"resetConfigurationSetName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationEumsSmsOutputReference) ResetExternalId() {
	_jsii_.InvokeVoid(
		c,
		"resetExternalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationEumsSmsOutputReference) ResetInEntityId() {
	_jsii_.InvokeVoid(
		c,
		"resetInEntityId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationEumsSmsOutputReference) ResetInTemplateId() {
	_jsii_.InvokeVoid(
		c,
		"resetInTemplateId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationEumsSmsOutputReference) ResetOriginationIdentity() {
	_jsii_.InvokeVoid(
		c,
		"resetOriginationIdentity",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationEumsSmsOutputReference) ResetRegion() {
	_jsii_.InvokeVoid(
		c,
		"resetRegion",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationEumsSmsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationEumsSmsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

