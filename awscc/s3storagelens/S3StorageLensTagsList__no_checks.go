// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package s3storagelens

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_S3StorageLensTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_S3StorageLensTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_S3StorageLensTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_S3StorageLensTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_S3StorageLensTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_S3StorageLensTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_S3StorageLensTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewS3StorageLensTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

