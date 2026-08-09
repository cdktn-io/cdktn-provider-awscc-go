// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package bedrockblueprint

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BedrockBlueprintTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (b *jsiiProxy_BedrockBlueprintTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (b *jsiiProxy_BedrockBlueprintTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_BedrockBlueprintTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_BedrockBlueprintTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_BedrockBlueprintTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_BedrockBlueprintTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewBedrockBlueprintTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

