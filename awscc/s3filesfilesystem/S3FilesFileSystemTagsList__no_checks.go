// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package s3filesfilesystem

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_S3FilesFileSystemTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_S3FilesFileSystemTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_S3FilesFileSystemTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_S3FilesFileSystemTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_S3FilesFileSystemTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_S3FilesFileSystemTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_S3FilesFileSystemTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewS3FilesFileSystemTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

