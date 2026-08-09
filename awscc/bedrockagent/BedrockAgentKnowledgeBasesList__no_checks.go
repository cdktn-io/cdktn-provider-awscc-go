// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package bedrockagent

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BedrockAgentKnowledgeBasesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (b *jsiiProxy_BedrockAgentKnowledgeBasesList) validateGetParameters(index *float64) error {
	return nil
}

func (b *jsiiProxy_BedrockAgentKnowledgeBasesList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_BedrockAgentKnowledgeBasesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_BedrockAgentKnowledgeBasesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_BedrockAgentKnowledgeBasesList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_BedrockAgentKnowledgeBasesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewBedrockAgentKnowledgeBasesListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

