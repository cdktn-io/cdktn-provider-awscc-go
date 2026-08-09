// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package gameliftscript

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GameliftScriptTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (g *jsiiProxy_GameliftScriptTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (g *jsiiProxy_GameliftScriptTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_GameliftScriptTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_GameliftScriptTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GameliftScriptTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_GameliftScriptTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewGameliftScriptTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

