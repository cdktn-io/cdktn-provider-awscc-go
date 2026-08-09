// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectcampaignsv2campaign

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/connectcampaignsv2campaign/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyOutboundModeOutputReference interface {
	cdktn.ComplexObject
	AgentlessConfig() *string
	SetAgentlessConfig(val *string)
	AgentlessConfigInput() *string
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
	PredictiveConfig() Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyOutboundModePredictiveConfigOutputReference
	PredictiveConfigInput() interface{}
	PreviewConfig() Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyOutboundModePreviewConfigOutputReference
	PreviewConfigInput() interface{}
	ProgressiveConfig() Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyOutboundModeProgressiveConfigOutputReference
	ProgressiveConfigInput() interface{}
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
	PutPredictiveConfig(value *Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyOutboundModePredictiveConfig)
	PutPreviewConfig(value *Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyOutboundModePreviewConfig)
	PutProgressiveConfig(value *Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyOutboundModeProgressiveConfig)
	ResetAgentlessConfig()
	ResetPredictiveConfig()
	ResetPreviewConfig()
	ResetProgressiveConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyOutboundModeOutputReference
type jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyOutboundModeOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyOutboundModeOutputReference) AgentlessConfig() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentlessConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyOutboundModeOutputReference) AgentlessConfigInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentlessConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyOutboundModeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyOutboundModeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyOutboundModeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyOutboundModeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyOutboundModeOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyOutboundModeOutputReference) PredictiveConfig() Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyOutboundModePredictiveConfigOutputReference {
	var returns Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyOutboundModePredictiveConfigOutputReference
	_jsii_.Get(
		j,
		"predictiveConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyOutboundModeOutputReference) PredictiveConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"predictiveConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyOutboundModeOutputReference) PreviewConfig() Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyOutboundModePreviewConfigOutputReference {
	var returns Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyOutboundModePreviewConfigOutputReference
	_jsii_.Get(
		j,
		"previewConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyOutboundModeOutputReference) PreviewConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"previewConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyOutboundModeOutputReference) ProgressiveConfig() Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyOutboundModeProgressiveConfigOutputReference {
	var returns Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyOutboundModeProgressiveConfigOutputReference
	_jsii_.Get(
		j,
		"progressiveConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyOutboundModeOutputReference) ProgressiveConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"progressiveConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyOutboundModeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyOutboundModeOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewConnectcampaignsv2CampaignChannelSubtypeConfigTelephonyOutboundModeOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyOutboundModeOutputReference {
	_init_.Initialize()

	if err := validateNewConnectcampaignsv2CampaignChannelSubtypeConfigTelephonyOutboundModeOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyOutboundModeOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.connectcampaignsv2Campaign.Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyOutboundModeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewConnectcampaignsv2CampaignChannelSubtypeConfigTelephonyOutboundModeOutputReference_Override(c Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyOutboundModeOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.connectcampaignsv2Campaign.Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyOutboundModeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyOutboundModeOutputReference)SetAgentlessConfig(val *string) {
	if err := j.validateSetAgentlessConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"agentlessConfig",
		val,
	)
}

func (j *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyOutboundModeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyOutboundModeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyOutboundModeOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyOutboundModeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyOutboundModeOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyOutboundModeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyOutboundModeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyOutboundModeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyOutboundModeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyOutboundModeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyOutboundModeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyOutboundModeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyOutboundModeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyOutboundModeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyOutboundModeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyOutboundModeOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyOutboundModeOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyOutboundModeOutputReference) PutPredictiveConfig(value *Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyOutboundModePredictiveConfig) {
	if err := c.validatePutPredictiveConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPredictiveConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyOutboundModeOutputReference) PutPreviewConfig(value *Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyOutboundModePreviewConfig) {
	if err := c.validatePutPreviewConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPreviewConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyOutboundModeOutputReference) PutProgressiveConfig(value *Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyOutboundModeProgressiveConfig) {
	if err := c.validatePutProgressiveConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putProgressiveConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyOutboundModeOutputReference) ResetAgentlessConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetAgentlessConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyOutboundModeOutputReference) ResetPredictiveConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetPredictiveConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyOutboundModeOutputReference) ResetPreviewConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetPreviewConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyOutboundModeOutputReference) ResetProgressiveConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetProgressiveConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyOutboundModeOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyOutboundModeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

