// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package bedrocksession

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BedrockSessionTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (b *jsiiProxy_BedrockSessionTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (b *jsiiProxy_BedrockSessionTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_BedrockSessionTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_BedrockSessionTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_BedrockSessionTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_BedrockSessionTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewBedrockSessionTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

