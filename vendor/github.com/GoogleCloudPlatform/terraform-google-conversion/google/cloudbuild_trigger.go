// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func stepTimeoutCustomizeDiff(_ context.Context, diff *schema.ResourceDiff, v interface{}) error {
	buildList := diff.Get("build").([]interface{})
	if len(buildList) == 0 || buildList[0] == nil {
		return nil
	}
	build := buildList[0].(map[string]interface{})
	buildTimeoutString := build["timeout"].(string)

	buildTimeout, err := time.ParseDuration(buildTimeoutString)
	if err != nil {
		return fmt.Errorf("Error parsing build timeout : %s", err)
	}

	var stepTimeoutSum time.Duration = 0
	steps := build["step"].([]interface{})
	for _, rawstep := range steps {
		if rawstep == nil {
			continue
		}
		step := rawstep.(map[string]interface{})
		timeoutString := step["timeout"].(string)
		if len(timeoutString) == 0 {
			continue
		}

		timeout, err := time.ParseDuration(timeoutString)
		if err != nil {
			return fmt.Errorf("Error parsing build step timeout: %s", err)
		}
		stepTimeoutSum += timeout
	}
	if stepTimeoutSum > buildTimeout {
		return fmt.Errorf("Step timeout sum (%v) cannot be greater than build timeout (%v)", stepTimeoutSum, buildTimeout)
	}
	return nil
}

func GetCloudBuildTriggerCaiObject(d TerraformResourceData, config *Config) (Asset, error) {
	name, err := assetName(d, config, "//cloudbuild.googleapis.com/projects/{{project}}/triggers/{{trigger_id}}")
	if err != nil {
		return Asset{}, err
	}
	if obj, err := GetCloudBuildTriggerApiObject(d, config); err == nil {
		return Asset{
			Name: name,
			Type: "cloudbuild.googleapis.com/Trigger",
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/cloudbuild/v1/rest",
				DiscoveryName:        "Trigger",
				Data:                 obj,
			},
		}, nil
	} else {
		return Asset{}, err
	}
}

func GetCloudBuildTriggerApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandCloudBuildTriggerName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandCloudBuildTriggerDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	tagsProp, err := expandCloudBuildTriggerTags(d.Get("tags"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("tags"); !isEmptyValue(reflect.ValueOf(tagsProp)) && (ok || !reflect.DeepEqual(v, tagsProp)) {
		obj["tags"] = tagsProp
	}
	disabledProp, err := expandCloudBuildTriggerDisabled(d.Get("disabled"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("disabled"); !isEmptyValue(reflect.ValueOf(disabledProp)) && (ok || !reflect.DeepEqual(v, disabledProp)) {
		obj["disabled"] = disabledProp
	}
	substitutionsProp, err := expandCloudBuildTriggerSubstitutions(d.Get("substitutions"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("substitutions"); !isEmptyValue(reflect.ValueOf(substitutionsProp)) && (ok || !reflect.DeepEqual(v, substitutionsProp)) {
		obj["substitutions"] = substitutionsProp
	}
	filenameProp, err := expandCloudBuildTriggerFilename(d.Get("filename"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("filename"); !isEmptyValue(reflect.ValueOf(filenameProp)) && (ok || !reflect.DeepEqual(v, filenameProp)) {
		obj["filename"] = filenameProp
	}
	ignoredFilesProp, err := expandCloudBuildTriggerIgnoredFiles(d.Get("ignored_files"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("ignored_files"); !isEmptyValue(reflect.ValueOf(ignoredFilesProp)) && (ok || !reflect.DeepEqual(v, ignoredFilesProp)) {
		obj["ignoredFiles"] = ignoredFilesProp
	}
	includedFilesProp, err := expandCloudBuildTriggerIncludedFiles(d.Get("included_files"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("included_files"); !isEmptyValue(reflect.ValueOf(includedFilesProp)) && (ok || !reflect.DeepEqual(v, includedFilesProp)) {
		obj["includedFiles"] = includedFilesProp
	}
	triggerTemplateProp, err := expandCloudBuildTriggerTriggerTemplate(d.Get("trigger_template"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("trigger_template"); !isEmptyValue(reflect.ValueOf(triggerTemplateProp)) && (ok || !reflect.DeepEqual(v, triggerTemplateProp)) {
		obj["triggerTemplate"] = triggerTemplateProp
	}
	buildProp, err := expandCloudBuildTriggerBuild(d.Get("build"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("build"); !isEmptyValue(reflect.ValueOf(buildProp)) && (ok || !reflect.DeepEqual(v, buildProp)) {
		obj["build"] = buildProp
	}

	return obj, nil
}

func expandCloudBuildTriggerName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerTags(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerDisabled(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerSubstitutions(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandCloudBuildTriggerFilename(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerIgnoredFiles(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerIncludedFiles(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerTriggerTemplate(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedProjectId, err := expandCloudBuildTriggerTriggerTemplateProjectId(original["project_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProjectId); val.IsValid() && !isEmptyValue(val) {
		transformed["projectId"] = transformedProjectId
	}

	transformedRepoName, err := expandCloudBuildTriggerTriggerTemplateRepoName(original["repo_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRepoName); val.IsValid() && !isEmptyValue(val) {
		transformed["repoName"] = transformedRepoName
	}

	transformedDir, err := expandCloudBuildTriggerTriggerTemplateDir(original["dir"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDir); val.IsValid() && !isEmptyValue(val) {
		transformed["dir"] = transformedDir
	}

	transformedInvertRegex, err := expandCloudBuildTriggerTriggerTemplateInvertRegex(original["invert_regex"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedInvertRegex); val.IsValid() && !isEmptyValue(val) {
		transformed["invertRegex"] = transformedInvertRegex
	}

	transformedBranchName, err := expandCloudBuildTriggerTriggerTemplateBranchName(original["branch_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedBranchName); val.IsValid() && !isEmptyValue(val) {
		transformed["branchName"] = transformedBranchName
	}

	transformedTagName, err := expandCloudBuildTriggerTriggerTemplateTagName(original["tag_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTagName); val.IsValid() && !isEmptyValue(val) {
		transformed["tagName"] = transformedTagName
	}

	transformedCommitSha, err := expandCloudBuildTriggerTriggerTemplateCommitSha(original["commit_sha"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCommitSha); val.IsValid() && !isEmptyValue(val) {
		transformed["commitSha"] = transformedCommitSha
	}

	return transformed, nil
}

func expandCloudBuildTriggerTriggerTemplateProjectId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerTriggerTemplateRepoName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerTriggerTemplateDir(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerTriggerTemplateInvertRegex(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerTriggerTemplateBranchName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerTriggerTemplateTagName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerTriggerTemplateCommitSha(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerBuild(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedSource, err := expandCloudBuildTriggerBuildSource(original["source"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSource); val.IsValid() && !isEmptyValue(val) {
		transformed["source"] = transformedSource
	}

	transformedTags, err := expandCloudBuildTriggerBuildTags(original["tags"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTags); val.IsValid() && !isEmptyValue(val) {
		transformed["tags"] = transformedTags
	}

	transformedImages, err := expandCloudBuildTriggerBuildImages(original["images"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedImages); val.IsValid() && !isEmptyValue(val) {
		transformed["images"] = transformedImages
	}

	transformedSubstitutions, err := expandCloudBuildTriggerBuildSubstitutions(original["substitutions"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSubstitutions); val.IsValid() && !isEmptyValue(val) {
		transformed["substitutions"] = transformedSubstitutions
	}

	transformedQueueTtl, err := expandCloudBuildTriggerBuildQueueTtl(original["queue_ttl"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedQueueTtl); val.IsValid() && !isEmptyValue(val) {
		transformed["queueTtl"] = transformedQueueTtl
	}

	transformedLogsBucket, err := expandCloudBuildTriggerBuildLogsBucket(original["logs_bucket"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLogsBucket); val.IsValid() && !isEmptyValue(val) {
		transformed["logsBucket"] = transformedLogsBucket
	}

	transformedTimeout, err := expandCloudBuildTriggerBuildTimeout(original["timeout"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTimeout); val.IsValid() && !isEmptyValue(val) {
		transformed["timeout"] = transformedTimeout
	}

	transformedSecret, err := expandCloudBuildTriggerBuildSecret(original["secret"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSecret); val.IsValid() && !isEmptyValue(val) {
		transformed["secrets"] = transformedSecret
	}

	transformedStep, err := expandCloudBuildTriggerBuildStep(original["step"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedStep); val.IsValid() && !isEmptyValue(val) {
		transformed["steps"] = transformedStep
	}

	transformedArtifacts, err := expandCloudBuildTriggerBuildArtifacts(original["artifacts"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedArtifacts); val.IsValid() && !isEmptyValue(val) {
		transformed["artifacts"] = transformedArtifacts
	}

	transformedOptions, err := expandCloudBuildTriggerBuildOptions(original["options"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedOptions); val.IsValid() && !isEmptyValue(val) {
		transformed["options"] = transformedOptions
	}

	return transformed, nil
}

func expandCloudBuildTriggerBuildSource(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedStorageSource, err := expandCloudBuildTriggerBuildSourceStorageSource(original["storage_source"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedStorageSource); val.IsValid() && !isEmptyValue(val) {
		transformed["storageSource"] = transformedStorageSource
	}

	transformedRepoSource, err := expandCloudBuildTriggerBuildSourceRepoSource(original["repo_source"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRepoSource); val.IsValid() && !isEmptyValue(val) {
		transformed["repoSource"] = transformedRepoSource
	}

	return transformed, nil
}

func expandCloudBuildTriggerBuildSourceStorageSource(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedBucket, err := expandCloudBuildTriggerBuildSourceStorageSourceBucket(original["bucket"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedBucket); val.IsValid() && !isEmptyValue(val) {
		transformed["bucket"] = transformedBucket
	}

	transformedObject, err := expandCloudBuildTriggerBuildSourceStorageSourceObject(original["object"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedObject); val.IsValid() && !isEmptyValue(val) {
		transformed["object"] = transformedObject
	}

	transformedGeneration, err := expandCloudBuildTriggerBuildSourceStorageSourceGeneration(original["generation"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedGeneration); val.IsValid() && !isEmptyValue(val) {
		transformed["generation"] = transformedGeneration
	}

	return transformed, nil
}

func expandCloudBuildTriggerBuildSourceStorageSourceBucket(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerBuildSourceStorageSourceObject(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerBuildSourceStorageSourceGeneration(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerBuildSourceRepoSource(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedProjectId, err := expandCloudBuildTriggerBuildSourceRepoSourceProjectId(original["project_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProjectId); val.IsValid() && !isEmptyValue(val) {
		transformed["projectId"] = transformedProjectId
	}

	transformedRepoName, err := expandCloudBuildTriggerBuildSourceRepoSourceRepoName(original["repo_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRepoName); val.IsValid() && !isEmptyValue(val) {
		transformed["repoName"] = transformedRepoName
	}

	transformedDir, err := expandCloudBuildTriggerBuildSourceRepoSourceDir(original["dir"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDir); val.IsValid() && !isEmptyValue(val) {
		transformed["dir"] = transformedDir
	}

	transformedInvertRegex, err := expandCloudBuildTriggerBuildSourceRepoSourceInvertRegex(original["invert_regex"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedInvertRegex); val.IsValid() && !isEmptyValue(val) {
		transformed["invertRegex"] = transformedInvertRegex
	}

	transformedSubstitutions, err := expandCloudBuildTriggerBuildSourceRepoSourceSubstitutions(original["substitutions"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSubstitutions); val.IsValid() && !isEmptyValue(val) {
		transformed["substitutions"] = transformedSubstitutions
	}

	transformedBranchName, err := expandCloudBuildTriggerBuildSourceRepoSourceBranchName(original["branch_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedBranchName); val.IsValid() && !isEmptyValue(val) {
		transformed["branchName"] = transformedBranchName
	}

	transformedTagName, err := expandCloudBuildTriggerBuildSourceRepoSourceTagName(original["tag_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTagName); val.IsValid() && !isEmptyValue(val) {
		transformed["tagName"] = transformedTagName
	}

	transformedCommitSha, err := expandCloudBuildTriggerBuildSourceRepoSourceCommitSha(original["commit_sha"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCommitSha); val.IsValid() && !isEmptyValue(val) {
		transformed["commitSha"] = transformedCommitSha
	}

	return transformed, nil
}

func expandCloudBuildTriggerBuildSourceRepoSourceProjectId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerBuildSourceRepoSourceRepoName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerBuildSourceRepoSourceDir(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerBuildSourceRepoSourceInvertRegex(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerBuildSourceRepoSourceSubstitutions(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandCloudBuildTriggerBuildSourceRepoSourceBranchName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerBuildSourceRepoSourceTagName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerBuildSourceRepoSourceCommitSha(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerBuildTags(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerBuildImages(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerBuildSubstitutions(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandCloudBuildTriggerBuildQueueTtl(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerBuildLogsBucket(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerBuildTimeout(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerBuildSecret(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedKmsKeyName, err := expandCloudBuildTriggerBuildSecretKmsKeyName(original["kms_key_name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedKmsKeyName); val.IsValid() && !isEmptyValue(val) {
			transformed["kmsKeyName"] = transformedKmsKeyName
		}

		transformedSecretEnv, err := expandCloudBuildTriggerBuildSecretSecretEnv(original["secret_env"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedSecretEnv); val.IsValid() && !isEmptyValue(val) {
			transformed["secretEnv"] = transformedSecretEnv
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandCloudBuildTriggerBuildSecretKmsKeyName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerBuildSecretSecretEnv(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandCloudBuildTriggerBuildStep(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedName, err := expandCloudBuildTriggerBuildStepName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !isEmptyValue(val) {
			transformed["name"] = transformedName
		}

		transformedArgs, err := expandCloudBuildTriggerBuildStepArgs(original["args"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedArgs); val.IsValid() && !isEmptyValue(val) {
			transformed["args"] = transformedArgs
		}

		transformedEnv, err := expandCloudBuildTriggerBuildStepEnv(original["env"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedEnv); val.IsValid() && !isEmptyValue(val) {
			transformed["env"] = transformedEnv
		}

		transformedId, err := expandCloudBuildTriggerBuildStepId(original["id"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedId); val.IsValid() && !isEmptyValue(val) {
			transformed["id"] = transformedId
		}

		transformedEntrypoint, err := expandCloudBuildTriggerBuildStepEntrypoint(original["entrypoint"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedEntrypoint); val.IsValid() && !isEmptyValue(val) {
			transformed["entrypoint"] = transformedEntrypoint
		}

		transformedDir, err := expandCloudBuildTriggerBuildStepDir(original["dir"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDir); val.IsValid() && !isEmptyValue(val) {
			transformed["dir"] = transformedDir
		}

		transformedSecretEnv, err := expandCloudBuildTriggerBuildStepSecretEnv(original["secret_env"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedSecretEnv); val.IsValid() && !isEmptyValue(val) {
			transformed["secretEnv"] = transformedSecretEnv
		}

		transformedTimeout, err := expandCloudBuildTriggerBuildStepTimeout(original["timeout"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedTimeout); val.IsValid() && !isEmptyValue(val) {
			transformed["timeout"] = transformedTimeout
		}

		transformedTiming, err := expandCloudBuildTriggerBuildStepTiming(original["timing"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedTiming); val.IsValid() && !isEmptyValue(val) {
			transformed["timing"] = transformedTiming
		}

		transformedVolumes, err := expandCloudBuildTriggerBuildStepVolumes(original["volumes"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedVolumes); val.IsValid() && !isEmptyValue(val) {
			transformed["volumes"] = transformedVolumes
		}

		transformedWaitFor, err := expandCloudBuildTriggerBuildStepWaitFor(original["wait_for"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedWaitFor); val.IsValid() && !isEmptyValue(val) {
			transformed["waitFor"] = transformedWaitFor
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandCloudBuildTriggerBuildStepName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerBuildStepArgs(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerBuildStepEnv(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerBuildStepId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerBuildStepEntrypoint(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerBuildStepDir(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerBuildStepSecretEnv(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerBuildStepTimeout(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerBuildStepTiming(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerBuildStepVolumes(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedName, err := expandCloudBuildTriggerBuildStepVolumesName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !isEmptyValue(val) {
			transformed["name"] = transformedName
		}

		transformedPath, err := expandCloudBuildTriggerBuildStepVolumesPath(original["path"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedPath); val.IsValid() && !isEmptyValue(val) {
			transformed["path"] = transformedPath
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandCloudBuildTriggerBuildStepVolumesName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerBuildStepVolumesPath(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerBuildStepWaitFor(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerBuildArtifacts(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedImages, err := expandCloudBuildTriggerBuildArtifactsImages(original["images"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedImages); val.IsValid() && !isEmptyValue(val) {
		transformed["images"] = transformedImages
	}

	transformedObjects, err := expandCloudBuildTriggerBuildArtifactsObjects(original["objects"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedObjects); val.IsValid() && !isEmptyValue(val) {
		transformed["objects"] = transformedObjects
	}

	return transformed, nil
}

func expandCloudBuildTriggerBuildArtifactsImages(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerBuildArtifactsObjects(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedLocation, err := expandCloudBuildTriggerBuildArtifactsObjectsLocation(original["location"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLocation); val.IsValid() && !isEmptyValue(val) {
		transformed["location"] = transformedLocation
	}

	transformedPaths, err := expandCloudBuildTriggerBuildArtifactsObjectsPaths(original["paths"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPaths); val.IsValid() && !isEmptyValue(val) {
		transformed["paths"] = transformedPaths
	}

	transformedTiming, err := expandCloudBuildTriggerBuildArtifactsObjectsTiming(original["timing"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTiming); val.IsValid() && !isEmptyValue(val) {
		transformed["timing"] = transformedTiming
	}

	return transformed, nil
}

func expandCloudBuildTriggerBuildArtifactsObjectsLocation(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerBuildArtifactsObjectsPaths(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerBuildArtifactsObjectsTiming(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedStartTime, err := expandCloudBuildTriggerBuildArtifactsObjectsTimingStartTime(original["start_time"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedStartTime); val.IsValid() && !isEmptyValue(val) {
		transformed["startTime"] = transformedStartTime
	}

	transformedEndTime, err := expandCloudBuildTriggerBuildArtifactsObjectsTimingEndTime(original["end_time"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEndTime); val.IsValid() && !isEmptyValue(val) {
		transformed["endTime"] = transformedEndTime
	}

	return transformed, nil
}

func expandCloudBuildTriggerBuildArtifactsObjectsTimingStartTime(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerBuildArtifactsObjectsTimingEndTime(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerBuildOptions(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedSourceProvenanceHash, err := expandCloudBuildTriggerBuildOptionsSourceProvenanceHash(original["source_provenance_hash"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSourceProvenanceHash); val.IsValid() && !isEmptyValue(val) {
		transformed["sourceProvenanceHash"] = transformedSourceProvenanceHash
	}

	transformedRequestedVerifyOption, err := expandCloudBuildTriggerBuildOptionsRequestedVerifyOption(original["requested_verify_option"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRequestedVerifyOption); val.IsValid() && !isEmptyValue(val) {
		transformed["requestedVerifyOption"] = transformedRequestedVerifyOption
	}

	transformedMachineType, err := expandCloudBuildTriggerBuildOptionsMachineType(original["machine_type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMachineType); val.IsValid() && !isEmptyValue(val) {
		transformed["machineType"] = transformedMachineType
	}

	transformedDiskSizeGb, err := expandCloudBuildTriggerBuildOptionsDiskSizeGb(original["disk_size_gb"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDiskSizeGb); val.IsValid() && !isEmptyValue(val) {
		transformed["diskSizeGb"] = transformedDiskSizeGb
	}

	transformedSubstitutionOption, err := expandCloudBuildTriggerBuildOptionsSubstitutionOption(original["substitution_option"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSubstitutionOption); val.IsValid() && !isEmptyValue(val) {
		transformed["substitutionOption"] = transformedSubstitutionOption
	}

	transformedDynamicSubstitutions, err := expandCloudBuildTriggerBuildOptionsDynamicSubstitutions(original["dynamic_substitutions"], d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["dynamicSubstitutions"] = transformedDynamicSubstitutions
	}

	transformedLogStreamingOption, err := expandCloudBuildTriggerBuildOptionsLogStreamingOption(original["log_streaming_option"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLogStreamingOption); val.IsValid() && !isEmptyValue(val) {
		transformed["logStreamingOption"] = transformedLogStreamingOption
	}

	transformedWorkerPool, err := expandCloudBuildTriggerBuildOptionsWorkerPool(original["worker_pool"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedWorkerPool); val.IsValid() && !isEmptyValue(val) {
		transformed["workerPool"] = transformedWorkerPool
	}

	transformedLogging, err := expandCloudBuildTriggerBuildOptionsLogging(original["logging"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLogging); val.IsValid() && !isEmptyValue(val) {
		transformed["logging"] = transformedLogging
	}

	transformedEnv, err := expandCloudBuildTriggerBuildOptionsEnv(original["env"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnv); val.IsValid() && !isEmptyValue(val) {
		transformed["env"] = transformedEnv
	}

	transformedSecretEnv, err := expandCloudBuildTriggerBuildOptionsSecretEnv(original["secret_env"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSecretEnv); val.IsValid() && !isEmptyValue(val) {
		transformed["secretEnv"] = transformedSecretEnv
	}

	transformedVolumes, err := expandCloudBuildTriggerBuildOptionsVolumes(original["volumes"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedVolumes); val.IsValid() && !isEmptyValue(val) {
		transformed["volumes"] = transformedVolumes
	}

	return transformed, nil
}

func expandCloudBuildTriggerBuildOptionsSourceProvenanceHash(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerBuildOptionsRequestedVerifyOption(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerBuildOptionsMachineType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerBuildOptionsDiskSizeGb(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerBuildOptionsSubstitutionOption(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerBuildOptionsDynamicSubstitutions(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerBuildOptionsLogStreamingOption(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerBuildOptionsWorkerPool(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerBuildOptionsLogging(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerBuildOptionsEnv(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerBuildOptionsSecretEnv(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerBuildOptionsVolumes(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedName, err := expandCloudBuildTriggerBuildOptionsVolumesName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !isEmptyValue(val) {
			transformed["name"] = transformedName
		}

		transformedPath, err := expandCloudBuildTriggerBuildOptionsVolumesPath(original["path"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedPath); val.IsValid() && !isEmptyValue(val) {
			transformed["path"] = transformedPath
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandCloudBuildTriggerBuildOptionsVolumesName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerBuildOptionsVolumesPath(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
