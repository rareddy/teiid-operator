// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"./pkg/apis/vdb/v1alpha1.Image":                      schema_pkg_apis_vdb_v1alpha1_Image(ref),
		"./pkg/apis/vdb/v1alpha1.S2i":                        schema_pkg_apis_vdb_v1alpha1_S2i(ref),
		"./pkg/apis/vdb/v1alpha1.Source":                     schema_pkg_apis_vdb_v1alpha1_Source(ref),
		"./pkg/apis/vdb/v1alpha1.VirtualDatabase":            schema_pkg_apis_vdb_v1alpha1_VirtualDatabase(ref),
		"./pkg/apis/vdb/v1alpha1.VirtualDatabaseBuildObject": schema_pkg_apis_vdb_v1alpha1_VirtualDatabaseBuildObject(ref),
		"./pkg/apis/vdb/v1alpha1.VirtualDatabaseSpec":        schema_pkg_apis_vdb_v1alpha1_VirtualDatabaseSpec(ref),
		"./pkg/apis/vdb/v1alpha1.VirtualDatabaseStatus":      schema_pkg_apis_vdb_v1alpha1_VirtualDatabaseStatus(ref),
	}
}

func schema_pkg_apis_vdb_v1alpha1_Image(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Image - image details",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"imageStreamName": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"imageStreamTag": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"imageStreamNamespace": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"imageRegistry": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"imageRepository": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"builderImage": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
				},
			},
		},
	}
}

func schema_pkg_apis_vdb_v1alpha1_S2i(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "S2i Git coordinates to locate the s2i image",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"registry": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"imagePrefix": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"imageName": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"tag": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
			},
		},
	}
}

func schema_pkg_apis_vdb_v1alpha1_Source(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Source VDB coordinates to locate the source code to build",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"ddl": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"maven": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"openapi": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"dependencies": {
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
					"mavenRepositories": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func schema_pkg_apis_vdb_v1alpha1_VirtualDatabase(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "VirtualDatabase is the Schema for the virtualdatabases API",
				Type:        []string{"object"},
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
							Ref: ref("./pkg/apis/vdb/v1alpha1.VirtualDatabaseSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/vdb/v1alpha1.VirtualDatabaseStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/vdb/v1alpha1.VirtualDatabaseSpec", "./pkg/apis/vdb/v1alpha1.VirtualDatabaseStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_vdb_v1alpha1_VirtualDatabaseBuildObject(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "VirtualDatabaseBuildObject Data to define how to build an application from source",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"incremental": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
					"env": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("k8s.io/api/core/v1.EnvVar"),
									},
								},
							},
						},
					},
					"source": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/vdb/v1alpha1.Source"),
						},
					},
					"s2i": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/vdb/v1alpha1.S2i"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/vdb/v1alpha1.S2i", "./pkg/apis/vdb/v1alpha1.Source", "k8s.io/api/core/v1.EnvVar"},
	}
}

func schema_pkg_apis_vdb_v1alpha1_VirtualDatabaseSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "VirtualDatabaseSpec defines the desired state of VirtualDatabase",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"replicas": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"exposeVia3scale": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
					"env": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("k8s.io/api/core/v1.EnvVar"),
									},
								},
							},
						},
					},
					"runtime": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/vdb/v1alpha1.RuntimeType"),
						},
					},
					"resources": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/api/core/v1.ResourceRequirements"),
						},
					},
					"build": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/vdb/v1alpha1.VirtualDatabaseBuildObject"),
						},
					},
				},
				Required: []string{"build"},
			},
		},
		Dependencies: []string{
			"./pkg/apis/vdb/v1alpha1.RuntimeType", "./pkg/apis/vdb/v1alpha1.VirtualDatabaseBuildObject", "k8s.io/api/core/v1.EnvVar", "k8s.io/api/core/v1.ResourceRequirements"},
	}
}

func schema_pkg_apis_vdb_v1alpha1_VirtualDatabaseStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "VirtualDatabaseStatus defines the observed state of VirtualDatabase",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"phase": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"digest": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"failure": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"route": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
			},
		},
	}
}