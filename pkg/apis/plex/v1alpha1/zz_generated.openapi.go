// +build !

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"./pkg/apis/plex/v1alpha1.PlexTranscodeJob":       schema_pkg_apis_plex_v1alpha1_PlexTranscodeJob(ref),
		"./pkg/apis/plex/v1alpha1.PlexTranscodeJobSpec":   schema_pkg_apis_plex_v1alpha1_PlexTranscodeJobSpec(ref),
		"./pkg/apis/plex/v1alpha1.PlexTranscodeJobStatus": schema_pkg_apis_plex_v1alpha1_PlexTranscodeJobStatus(ref),
	}
}

func schema_pkg_apis_plex_v1alpha1_PlexTranscodeJob(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "PlexTranscodeJob is the Schema for the plextranscodejobs API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/plex/v1alpha1.PlexTranscodeJobSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/plex/v1alpha1.PlexTranscodeJobStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/plex/v1alpha1.PlexTranscodeJobSpec", "./pkg/apis/plex/v1alpha1.PlexTranscodeJobStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_plex_v1alpha1_PlexTranscodeJobSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "PlexTranscodeJobSpec defines the desired state of PlexTranscodeJob",
				Properties: map[string]spec.Schema{
					"args": {
						SchemaProps: spec.SchemaProps{
							Description: "An array of arguments to pass to the real plex transcode binary",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"env": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"cwd": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"args", "env", "cwd"},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_plex_v1alpha1_PlexTranscodeJobStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "PlexTranscodeJobStatus defines the observed state of PlexTranscodeJob",
				Properties: map[string]spec.Schema{
					"transcoder": {
						SchemaProps: spec.SchemaProps{
							Description: "Name of the transcoder pod assigned the transcode job",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"state": {
						SchemaProps: spec.SchemaProps{
							Description: "The state of the job, one of: CREATED ASSIGNED STARTED FAILED COMPLETED",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"error": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"transcoder", "state", "error"},
			},
		},
		Dependencies: []string{},
	}
}