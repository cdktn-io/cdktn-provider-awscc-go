// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package bedrockprompt

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BedrockPromptVariantsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (b *jsiiProxy_BedrockPromptVariantsList) validateGetParameters(index *float64) error {
	return nil
}

func (b *jsiiProxy_BedrockPromptVariantsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_BedrockPromptVariantsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_BedrockPromptVariantsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_BedrockPromptVariantsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_BedrockPromptVariantsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewBedrockPromptVariantsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

