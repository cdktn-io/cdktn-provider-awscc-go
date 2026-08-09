// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package bedrockagent

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BedrockAgentActionGroupsFunctionSchemaFunctionsParametersMap) validateGetParameters(key *string) error {
	return nil
}

func (b *jsiiProxy_BedrockAgentActionGroupsFunctionSchemaFunctionsParametersMap) validateInterpolationForAttributeParameters(property *string) error {
	return nil
}

func (b *jsiiProxy_BedrockAgentActionGroupsFunctionSchemaFunctionsParametersMap) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_BedrockAgentActionGroupsFunctionSchemaFunctionsParametersMap) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_BedrockAgentActionGroupsFunctionSchemaFunctionsParametersMap) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_BedrockAgentActionGroupsFunctionSchemaFunctionsParametersMap) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func validateNewBedrockAgentActionGroupsFunctionSchemaFunctionsParametersMapParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) error {
	return nil
}

