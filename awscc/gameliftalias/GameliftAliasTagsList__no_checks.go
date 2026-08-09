// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package gameliftalias

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GameliftAliasTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (g *jsiiProxy_GameliftAliasTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (g *jsiiProxy_GameliftAliasTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_GameliftAliasTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_GameliftAliasTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GameliftAliasTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_GameliftAliasTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewGameliftAliasTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

