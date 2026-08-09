// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sesmailmanagerruleset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/sesmailmanagerruleset/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SesMailManagerRuleSetRulesActionsOutputReference interface {
	cdktn.ComplexObject
	AddHeader() SesMailManagerRuleSetRulesActionsAddHeaderOutputReference
	AddHeaderInput() interface{}
	Archive() SesMailManagerRuleSetRulesActionsArchiveOutputReference
	ArchiveInput() interface{}
	Bounce() SesMailManagerRuleSetRulesActionsBounceOutputReference
	BounceInput() interface{}
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
	DeliverToMailbox() SesMailManagerRuleSetRulesActionsDeliverToMailboxOutputReference
	DeliverToMailboxInput() interface{}
	DeliverToQBusiness() SesMailManagerRuleSetRulesActionsDeliverToQBusinessOutputReference
	DeliverToQBusinessInput() interface{}
	Drop() *string
	SetDrop(val *string)
	DropInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	InvokeLambda() SesMailManagerRuleSetRulesActionsInvokeLambdaOutputReference
	InvokeLambdaInput() interface{}
	PublishToSns() SesMailManagerRuleSetRulesActionsPublishToSnsOutputReference
	PublishToSnsInput() interface{}
	Relay() SesMailManagerRuleSetRulesActionsRelayOutputReference
	RelayInput() interface{}
	ReplaceRecipient() SesMailManagerRuleSetRulesActionsReplaceRecipientOutputReference
	ReplaceRecipientInput() interface{}
	Send() SesMailManagerRuleSetRulesActionsSendOutputReference
	SendInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	WriteToS3() SesMailManagerRuleSetRulesActionsWriteToS3OutputReference
	WriteToS3Input() interface{}
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
	PutAddHeader(value *SesMailManagerRuleSetRulesActionsAddHeader)
	PutArchive(value *SesMailManagerRuleSetRulesActionsArchive)
	PutBounce(value *SesMailManagerRuleSetRulesActionsBounce)
	PutDeliverToMailbox(value *SesMailManagerRuleSetRulesActionsDeliverToMailbox)
	PutDeliverToQBusiness(value *SesMailManagerRuleSetRulesActionsDeliverToQBusiness)
	PutInvokeLambda(value *SesMailManagerRuleSetRulesActionsInvokeLambda)
	PutPublishToSns(value *SesMailManagerRuleSetRulesActionsPublishToSns)
	PutRelay(value *SesMailManagerRuleSetRulesActionsRelay)
	PutReplaceRecipient(value *SesMailManagerRuleSetRulesActionsReplaceRecipient)
	PutSend(value *SesMailManagerRuleSetRulesActionsSend)
	PutWriteToS3(value *SesMailManagerRuleSetRulesActionsWriteToS3)
	ResetAddHeader()
	ResetArchive()
	ResetBounce()
	ResetDeliverToMailbox()
	ResetDeliverToQBusiness()
	ResetDrop()
	ResetInvokeLambda()
	ResetPublishToSns()
	ResetRelay()
	ResetReplaceRecipient()
	ResetSend()
	ResetWriteToS3()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SesMailManagerRuleSetRulesActionsOutputReference
type jsiiProxy_SesMailManagerRuleSetRulesActionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsOutputReference) AddHeader() SesMailManagerRuleSetRulesActionsAddHeaderOutputReference {
	var returns SesMailManagerRuleSetRulesActionsAddHeaderOutputReference
	_jsii_.Get(
		j,
		"addHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsOutputReference) AddHeaderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"addHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsOutputReference) Archive() SesMailManagerRuleSetRulesActionsArchiveOutputReference {
	var returns SesMailManagerRuleSetRulesActionsArchiveOutputReference
	_jsii_.Get(
		j,
		"archive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsOutputReference) ArchiveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"archiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsOutputReference) Bounce() SesMailManagerRuleSetRulesActionsBounceOutputReference {
	var returns SesMailManagerRuleSetRulesActionsBounceOutputReference
	_jsii_.Get(
		j,
		"bounce",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsOutputReference) BounceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bounceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsOutputReference) DeliverToMailbox() SesMailManagerRuleSetRulesActionsDeliverToMailboxOutputReference {
	var returns SesMailManagerRuleSetRulesActionsDeliverToMailboxOutputReference
	_jsii_.Get(
		j,
		"deliverToMailbox",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsOutputReference) DeliverToMailboxInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deliverToMailboxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsOutputReference) DeliverToQBusiness() SesMailManagerRuleSetRulesActionsDeliverToQBusinessOutputReference {
	var returns SesMailManagerRuleSetRulesActionsDeliverToQBusinessOutputReference
	_jsii_.Get(
		j,
		"deliverToQBusiness",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsOutputReference) DeliverToQBusinessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deliverToQBusinessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsOutputReference) Drop() *string {
	var returns *string
	_jsii_.Get(
		j,
		"drop",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsOutputReference) DropInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dropInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsOutputReference) InvokeLambda() SesMailManagerRuleSetRulesActionsInvokeLambdaOutputReference {
	var returns SesMailManagerRuleSetRulesActionsInvokeLambdaOutputReference
	_jsii_.Get(
		j,
		"invokeLambda",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsOutputReference) InvokeLambdaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"invokeLambdaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsOutputReference) PublishToSns() SesMailManagerRuleSetRulesActionsPublishToSnsOutputReference {
	var returns SesMailManagerRuleSetRulesActionsPublishToSnsOutputReference
	_jsii_.Get(
		j,
		"publishToSns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsOutputReference) PublishToSnsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publishToSnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsOutputReference) Relay() SesMailManagerRuleSetRulesActionsRelayOutputReference {
	var returns SesMailManagerRuleSetRulesActionsRelayOutputReference
	_jsii_.Get(
		j,
		"relay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsOutputReference) RelayInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"relayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsOutputReference) ReplaceRecipient() SesMailManagerRuleSetRulesActionsReplaceRecipientOutputReference {
	var returns SesMailManagerRuleSetRulesActionsReplaceRecipientOutputReference
	_jsii_.Get(
		j,
		"replaceRecipient",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsOutputReference) ReplaceRecipientInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"replaceRecipientInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsOutputReference) Send() SesMailManagerRuleSetRulesActionsSendOutputReference {
	var returns SesMailManagerRuleSetRulesActionsSendOutputReference
	_jsii_.Get(
		j,
		"send",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsOutputReference) SendInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sendInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsOutputReference) WriteToS3() SesMailManagerRuleSetRulesActionsWriteToS3OutputReference {
	var returns SesMailManagerRuleSetRulesActionsWriteToS3OutputReference
	_jsii_.Get(
		j,
		"writeToS3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsOutputReference) WriteToS3Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"writeToS3Input",
		&returns,
	)
	return returns
}


func NewSesMailManagerRuleSetRulesActionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) SesMailManagerRuleSetRulesActionsOutputReference {
	_init_.Initialize()

	if err := validateNewSesMailManagerRuleSetRulesActionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_SesMailManagerRuleSetRulesActionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.sesMailManagerRuleSet.SesMailManagerRuleSetRulesActionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewSesMailManagerRuleSetRulesActionsOutputReference_Override(s SesMailManagerRuleSetRulesActionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.sesMailManagerRuleSet.SesMailManagerRuleSetRulesActionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsOutputReference)SetDrop(val *string) {
	if err := j.validateSetDropParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"drop",
		val,
	)
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesActionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesActionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesActionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesActionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesActionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesActionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesActionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesActionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesActionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesActionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesActionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesActionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesActionsOutputReference) PutAddHeader(value *SesMailManagerRuleSetRulesActionsAddHeader) {
	if err := s.validatePutAddHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAddHeader",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesActionsOutputReference) PutArchive(value *SesMailManagerRuleSetRulesActionsArchive) {
	if err := s.validatePutArchiveParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putArchive",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesActionsOutputReference) PutBounce(value *SesMailManagerRuleSetRulesActionsBounce) {
	if err := s.validatePutBounceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putBounce",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesActionsOutputReference) PutDeliverToMailbox(value *SesMailManagerRuleSetRulesActionsDeliverToMailbox) {
	if err := s.validatePutDeliverToMailboxParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putDeliverToMailbox",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesActionsOutputReference) PutDeliverToQBusiness(value *SesMailManagerRuleSetRulesActionsDeliverToQBusiness) {
	if err := s.validatePutDeliverToQBusinessParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putDeliverToQBusiness",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesActionsOutputReference) PutInvokeLambda(value *SesMailManagerRuleSetRulesActionsInvokeLambda) {
	if err := s.validatePutInvokeLambdaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putInvokeLambda",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesActionsOutputReference) PutPublishToSns(value *SesMailManagerRuleSetRulesActionsPublishToSns) {
	if err := s.validatePutPublishToSnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putPublishToSns",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesActionsOutputReference) PutRelay(value *SesMailManagerRuleSetRulesActionsRelay) {
	if err := s.validatePutRelayParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putRelay",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesActionsOutputReference) PutReplaceRecipient(value *SesMailManagerRuleSetRulesActionsReplaceRecipient) {
	if err := s.validatePutReplaceRecipientParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putReplaceRecipient",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesActionsOutputReference) PutSend(value *SesMailManagerRuleSetRulesActionsSend) {
	if err := s.validatePutSendParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSend",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesActionsOutputReference) PutWriteToS3(value *SesMailManagerRuleSetRulesActionsWriteToS3) {
	if err := s.validatePutWriteToS3Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putWriteToS3",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesActionsOutputReference) ResetAddHeader() {
	_jsii_.InvokeVoid(
		s,
		"resetAddHeader",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesActionsOutputReference) ResetArchive() {
	_jsii_.InvokeVoid(
		s,
		"resetArchive",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesActionsOutputReference) ResetBounce() {
	_jsii_.InvokeVoid(
		s,
		"resetBounce",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesActionsOutputReference) ResetDeliverToMailbox() {
	_jsii_.InvokeVoid(
		s,
		"resetDeliverToMailbox",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesActionsOutputReference) ResetDeliverToQBusiness() {
	_jsii_.InvokeVoid(
		s,
		"resetDeliverToQBusiness",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesActionsOutputReference) ResetDrop() {
	_jsii_.InvokeVoid(
		s,
		"resetDrop",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesActionsOutputReference) ResetInvokeLambda() {
	_jsii_.InvokeVoid(
		s,
		"resetInvokeLambda",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesActionsOutputReference) ResetPublishToSns() {
	_jsii_.InvokeVoid(
		s,
		"resetPublishToSns",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesActionsOutputReference) ResetRelay() {
	_jsii_.InvokeVoid(
		s,
		"resetRelay",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesActionsOutputReference) ResetReplaceRecipient() {
	_jsii_.InvokeVoid(
		s,
		"resetReplaceRecipient",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesActionsOutputReference) ResetSend() {
	_jsii_.InvokeVoid(
		s,
		"resetSend",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesActionsOutputReference) ResetWriteToS3() {
	_jsii_.InvokeVoid(
		s,
		"resetWriteToS3",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesActionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := s.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesActionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

