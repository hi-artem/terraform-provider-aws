// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package servicecatalog_test

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/servicecatalog"
	"github.com/hashicorp/terraform-plugin-testing/config"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// WARNING: This file is *NOT* generated. It may need to be updated manually
// to keep it in sync with generated tests.

func TestAccServiceCatalogProvisionedProduct_tags(t *testing.T) {
	ctx := acctest.Context(t)
	var v servicecatalog.ProvisionedProductDetail
	resourceName := "aws_servicecatalog_provisioned_product.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.ServiceCatalogServiceID),
		CheckDestroy:             testAccCheckProvisionedProductDestroy(ctx),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		Steps: []resource.TestStep{
			{
				ConfigDirectory: config.StaticDirectory("testdata/ProvisionedProduct/tags/"),
				ConfigVariables: config.Variables{
					"rName": config.StringVariable(rName),
					"resource_tags": config.MapVariable(map[string]config.Variable{
						"key1": config.StringVariable("value1"),
					}),
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckProvisionedProductExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
				),
			},
			{
				ConfigDirectory: config.StaticDirectory("testdata/ProvisionedProduct/tags/"),
				ConfigVariables: config.Variables{
					"rName": config.StringVariable(rName),
					"resource_tags": config.MapVariable(map[string]config.Variable{
						"key1": config.StringVariable("value1"),
					}),
				},
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"accept_language", "ignore_errors", "provisioning_artifact_name", "provisioning_parameters", "retain_physical_resources",
				},
			},
			{
				ConfigDirectory: config.StaticDirectory("testdata/ProvisionedProduct/tags/"),
				ConfigVariables: config.Variables{
					"rName": config.StringVariable(rName),
					"resource_tags": config.MapVariable(map[string]config.Variable{
						"key1": config.StringVariable("value1updated"),
						"key2": config.StringVariable("value2"),
					}),
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckProvisionedProductExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1updated"),
					resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
				),
			},
			{
				ConfigDirectory: config.StaticDirectory("testdata/ProvisionedProduct/tags/"),
				ConfigVariables: config.Variables{
					"rName": config.StringVariable(rName),
					"resource_tags": config.MapVariable(map[string]config.Variable{
						"key1": config.StringVariable("value1updated"),
						"key2": config.StringVariable("value2"),
					}),
				},
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"accept_language", "ignore_errors", "provisioning_artifact_name", "provisioning_parameters", "retain_physical_resources",
				},
			},
			// {
			// 	ConfigDirectory: config.StaticDirectory("testdata/ProvisionedProduct/tags/"),
			// 	ConfigVariables: config.Variables{
			// 		"rName": config.StringVariable(rName),
			// 		"resource_tags": config.MapVariable(map[string]config.Variable{
			// 			"key2": config.StringVariable("value2"),
			// 		}),
			// 	},
			// 	Check: resource.ComposeAggregateTestCheckFunc(
			// 		testAccCheckProvisionedProductExists(ctx, resourceName, &v),
			// 		resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
			// 		resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
			// 	),
			// },
			// {
			// 	ConfigDirectory: config.StaticDirectory("testdata/ProvisionedProduct/tags/"),
			// 	ConfigVariables: config.Variables{
			// 		"rName": config.StringVariable(rName),
			// 		"resource_tags": config.MapVariable(map[string]config.Variable{
			// 			"key2": config.StringVariable("value2"),
			// 		}),
			// 	},
			// 	ResourceName:      resourceName,
			// 	ImportState:       true,
			// 	ImportStateVerify: true,
			// 	ImportStateVerifyIgnore: []string{
			// 		"accept_language", "ignore_errors", "provisioning_artifact_name", "provisioning_parameters", "retain_physical_resources",
			// 	},
			// },
			// {
			// 	ConfigDirectory: config.StaticDirectory("testdata/ProvisionedProduct/tags/"),
			// 	ConfigVariables: config.Variables{
			// 		"rName":        config.StringVariable(rName),
			// 		"resource_tags": nil,
			// 	},
			// 	Check: resource.ComposeAggregateTestCheckFunc(
			// 		testAccCheckProvisionedProductExists(ctx, resourceName, &v),
			// 		resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
			// 	),
			// },
			// {
			// 	ConfigDirectory: config.StaticDirectory("testdata/ProvisionedProduct/tags/"),
			// 	ConfigVariables: config.Variables{
			// 		"rName":        config.StringVariable(rName),
			// 		"resource_tags": nil,
			// 	},
			// 	ResourceName:      resourceName,
			// 	ImportState:       true,
			// 	ImportStateVerify: true,
			// 	ImportStateVerifyIgnore: []string{
			// 		"accept_language", "ignore_errors", "provisioning_artifact_name", "provisioning_parameters", "retain_physical_resources",
			// 	},
			// },
		},
	})
}

func TestAccServiceCatalogProvisionedProduct_tags_null(t *testing.T) {
	ctx := acctest.Context(t)
	var v servicecatalog.ProvisionedProductDetail
	resourceName := "aws_servicecatalog_provisioned_product.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.ServiceCatalogServiceID),
		CheckDestroy:             testAccCheckProvisionedProductDestroy(ctx),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		Steps: []resource.TestStep{
			{
				ConfigDirectory: config.StaticDirectory("testdata/ProvisionedProduct/tags/"),
				ConfigVariables: config.Variables{
					"rName": config.StringVariable(rName),
					"resource_tags": config.MapVariable(map[string]config.Variable{
						"key1": nil,
					}),
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckProvisionedProductExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
				),
			},
			{
				ConfigDirectory: config.StaticDirectory("testdata/ProvisionedProduct/tags/"),
				ConfigVariables: config.Variables{
					"rName": config.StringVariable(rName),
					"resource_tags": config.MapVariable(map[string]config.Variable{
						"key1": nil,
					}),
				},
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"accept_language", "ignore_errors", "provisioning_artifact_name", "provisioning_parameters", "retain_physical_resources",
				},
			},
			{
				ConfigDirectory: config.StaticDirectory("testdata/ProvisionedProduct/tags/"),
				ConfigVariables: config.Variables{
					"rName":         config.StringVariable(rName),
					"resource_tags": nil,
				},
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
		},
	})
}

func TestAccServiceCatalogProvisionedProduct_tags_AddOnUpdate(t *testing.T) {
	ctx := acctest.Context(t)
	var v servicecatalog.ProvisionedProductDetail
	resourceName := "aws_servicecatalog_provisioned_product.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.ServiceCatalogServiceID),
		CheckDestroy:             testAccCheckProvisionedProductDestroy(ctx),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		Steps: []resource.TestStep{
			{
				ConfigDirectory: config.StaticDirectory("testdata/ProvisionedProduct/tags/"),
				ConfigVariables: config.Variables{
					"rName":         config.StringVariable(rName),
					"resource_tags": nil,
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckProvisionedProductExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
				),
			},
			{
				ConfigDirectory: config.StaticDirectory("testdata/ProvisionedProduct/tags/"),
				ConfigVariables: config.Variables{
					"rName": config.StringVariable(rName),
					"resource_tags": config.MapVariable(map[string]config.Variable{
						"key1": config.StringVariable("value1"),
					}),
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckProvisionedProductExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
				),
			},
			{
				ConfigDirectory: config.StaticDirectory("testdata/ProvisionedProduct/tags/"),
				ConfigVariables: config.Variables{
					"rName": config.StringVariable(rName),
					"resource_tags": config.MapVariable(map[string]config.Variable{
						"key1": config.StringVariable("value1"),
					}),
				},
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"accept_language", "ignore_errors", "provisioning_artifact_name", "provisioning_parameters", "retain_physical_resources",
				},
			},
		},
	})
}

func TestAccServiceCatalogProvisionedProduct_tags_EmptyTag_OnCreate(t *testing.T) {
	t.Skip("Resource ProvisionedProduct does not support empty tags")

	ctx := acctest.Context(t)
	var v servicecatalog.ProvisionedProductDetail
	resourceName := "aws_servicecatalog_provisioned_product.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.ServiceCatalogServiceID),
		CheckDestroy:             testAccCheckProvisionedProductDestroy(ctx),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		Steps: []resource.TestStep{
			{
				ConfigDirectory: config.StaticDirectory("testdata/ProvisionedProduct/tags/"),
				ConfigVariables: config.Variables{
					"rName": config.StringVariable(rName),
					"resource_tags": config.MapVariable(map[string]config.Variable{
						"key1": config.StringVariable(""),
					}),
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckProvisionedProductExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", ""),
				),
			},
			{
				ConfigDirectory: config.StaticDirectory("testdata/ProvisionedProduct/tags/"),
				ConfigVariables: config.Variables{
					"rName": config.StringVariable(rName),
					"resource_tags": config.MapVariable(map[string]config.Variable{
						"key1": config.StringVariable(""),
					}),
				},
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"accept_language", "ignore_errors", "provisioning_artifact_name", "provisioning_parameters", "retain_physical_resources",
				},
			},
			{
				ConfigDirectory: config.StaticDirectory("testdata/ProvisionedProduct/tags/"),
				ConfigVariables: config.Variables{
					"rName":         config.StringVariable(rName),
					"resource_tags": nil,
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckProvisionedProductExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
				),
			},
			{
				ConfigDirectory: config.StaticDirectory("testdata/ProvisionedProduct/tags/"),
				ConfigVariables: config.Variables{
					"rName":         config.StringVariable(rName),
					"resource_tags": nil,
				},
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"accept_language", "ignore_errors", "provisioning_artifact_name", "provisioning_parameters", "retain_physical_resources",
				},
			},
		},
	})
}

func TestAccServiceCatalogProvisionedProduct_tags_EmptyTag_OnUpdate_Add(t *testing.T) {
	t.Skip("Resource ProvisionedProduct does not support empty tags")

	ctx := acctest.Context(t)
	var v servicecatalog.ProvisionedProductDetail
	resourceName := "aws_servicecatalog_provisioned_product.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.ServiceCatalogServiceID),
		CheckDestroy:             testAccCheckProvisionedProductDestroy(ctx),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		Steps: []resource.TestStep{
			{
				ConfigDirectory: config.StaticDirectory("testdata/ProvisionedProduct/tags/"),
				ConfigVariables: config.Variables{
					"rName": config.StringVariable(rName),
					"resource_tags": config.MapVariable(map[string]config.Variable{
						"key1": config.StringVariable("value1"),
					}),
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckProvisionedProductExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
				),
			},
			{
				ConfigDirectory: config.StaticDirectory("testdata/ProvisionedProduct/tags/"),
				ConfigVariables: config.Variables{
					"rName": config.StringVariable(rName),
					"resource_tags": config.MapVariable(map[string]config.Variable{
						"key1": config.StringVariable("value1"),
						"key2": config.StringVariable(""),
					}),
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckProvisionedProductExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key2", ""),
				),
			},
			{
				ConfigDirectory: config.StaticDirectory("testdata/ProvisionedProduct/tags/"),
				ConfigVariables: config.Variables{
					"rName": config.StringVariable(rName),
					"resource_tags": config.MapVariable(map[string]config.Variable{
						"key1": config.StringVariable("value1"),
						"key2": config.StringVariable(""),
					}),
				},
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"accept_language", "ignore_errors", "provisioning_artifact_name", "provisioning_parameters", "retain_physical_resources",
				},
			},
			{
				ConfigDirectory: config.StaticDirectory("testdata/ProvisionedProduct/tags/"),
				ConfigVariables: config.Variables{
					"rName": config.StringVariable(rName),
					"resource_tags": config.MapVariable(map[string]config.Variable{
						"key1": config.StringVariable("value1"),
					}),
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckProvisionedProductExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
				),
			},
			{
				ConfigDirectory: config.StaticDirectory("testdata/ProvisionedProduct/tags/"),
				ConfigVariables: config.Variables{
					"rName": config.StringVariable(rName),
					"resource_tags": config.MapVariable(map[string]config.Variable{
						"key1": config.StringVariable("value1"),
					}),
				},
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"accept_language", "ignore_errors", "provisioning_artifact_name", "provisioning_parameters", "retain_physical_resources",
				},
			},
		},
	})
}

func TestAccServiceCatalogProvisionedProduct_tags_EmptyTag_OnUpdate_Replace(t *testing.T) {
	t.Skip("Resource ProvisionedProduct does not support empty tags")

	ctx := acctest.Context(t)
	var v servicecatalog.ProvisionedProductDetail
	resourceName := "aws_servicecatalog_provisioned_product.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.ServiceCatalogServiceID),
		CheckDestroy:             testAccCheckProvisionedProductDestroy(ctx),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		Steps: []resource.TestStep{
			{
				ConfigDirectory: config.StaticDirectory("testdata/ProvisionedProduct/tags/"),
				ConfigVariables: config.Variables{
					"rName": config.StringVariable(rName),
					"resource_tags": config.MapVariable(map[string]config.Variable{
						"key1": config.StringVariable("value1"),
					}),
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckProvisionedProductExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
				),
			},
			{
				ConfigDirectory: config.StaticDirectory("testdata/ProvisionedProduct/tags/"),
				ConfigVariables: config.Variables{
					"rName": config.StringVariable(rName),
					"resource_tags": config.MapVariable(map[string]config.Variable{
						"key1": config.StringVariable(""),
					}),
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckProvisionedProductExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", ""),
				),
			},
			{
				ConfigDirectory: config.StaticDirectory("testdata/ProvisionedProduct/tags/"),
				ConfigVariables: config.Variables{
					"rName": config.StringVariable(rName),
					"resource_tags": config.MapVariable(map[string]config.Variable{
						"key1": config.StringVariable(""),
					}),
				},
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"accept_language", "ignore_errors", "provisioning_artifact_name", "provisioning_parameters", "retain_physical_resources",
				},
			},
		},
	})
}

func TestAccServiceCatalogProvisionedProduct_tags_DefaultTags_providerOnly(t *testing.T) {
	ctx := acctest.Context(t)
	var v servicecatalog.ProvisionedProductDetail
	resourceName := "aws_servicecatalog_provisioned_product.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:   acctest.ErrorCheck(t, names.ServiceCatalogServiceID),
		CheckDestroy: testAccCheckProvisionedProductDestroy(ctx),
		Steps: []resource.TestStep{
			{
				ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
				ConfigDirectory:          config.StaticDirectory("testdata/ProvisionedProduct/tags_defaults/"),
				ConfigVariables: config.Variables{
					"rName": config.StringVariable(rName),
					"provider_tags": config.MapVariable(map[string]config.Variable{
						"key1": config.StringVariable("value1"),
					}),
					"resource_tags": nil,
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckProvisionedProductExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.key1", "value1"),
				),
			},
			{
				ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
				ConfigDirectory:          config.StaticDirectory("testdata/ProvisionedProduct/tags_defaults/"),
				ConfigVariables: config.Variables{
					"rName": config.StringVariable(rName),
					"provider_tags": config.MapVariable(map[string]config.Variable{
						"key1": config.StringVariable("value1"),
					}),
					"resource_tags": nil,
				},
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"accept_language", "ignore_errors", "provisioning_artifact_name", "provisioning_parameters", "retain_physical_resources",
				},
			},
			{
				ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
				ConfigDirectory:          config.StaticDirectory("testdata/ProvisionedProduct/tags_defaults/"),
				ConfigVariables: config.Variables{
					"rName": config.StringVariable(rName),
					"provider_tags": config.MapVariable(map[string]config.Variable{
						"key1": config.StringVariable("value1updated"),
						"key2": config.StringVariable("value2"),
					}),
					"resource_tags": nil,
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckProvisionedProductExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.key1", "value1updated"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.key2", "value2"),
				),
			},
			{
				ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
				ConfigDirectory:          config.StaticDirectory("testdata/ProvisionedProduct/tags_defaults/"),
				ConfigVariables: config.Variables{
					"rName": config.StringVariable(rName),
					"provider_tags": config.MapVariable(map[string]config.Variable{
						"key1": config.StringVariable("value1updated"),
						"key2": config.StringVariable("value2"),
					}),
					"resource_tags": nil,
				},
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"accept_language", "ignore_errors", "provisioning_artifact_name", "provisioning_parameters", "retain_physical_resources",
				},
			},
			// {
			// 	ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
			// 	ConfigDirectory:          config.StaticDirectory("testdata/ProvisionedProduct/tags_defaults/"),
			// 	ConfigVariables: config.Variables{
			// 		"rName": config.StringVariable(rName),
			// 		"provider_tags": config.MapVariable(map[string]config.Variable{
			// 			"key2": config.StringVariable("value2"),
			// 		}),
			// 		"resource_tags": nil,
			// 	},
			// 	Check: resource.ComposeAggregateTestCheckFunc(
			// 		testAccCheckProvisionedProductExists(ctx, resourceName, &v),
			// 		resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
			// 		resource.TestCheckResourceAttr(resourceName, "tags_all.%", "1"),
			// 		resource.TestCheckResourceAttr(resourceName, "tags_all.key2", "value2"),
			// 	),
			// },
			// {
			// 	ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
			// 	ConfigDirectory:          config.StaticDirectory("testdata/ProvisionedProduct/tags_defaults/"),
			// 	ConfigVariables: config.Variables{
			// 		"rName": config.StringVariable(rName),
			// 		"provider_tags": config.MapVariable(map[string]config.Variable{
			// 			"key2": config.StringVariable("value2"),
			// 		}),
			// 		"resource_tags": nil,
			// 	},
			// 	ResourceName:      resourceName,
			// 	ImportState:       true,
			// 	ImportStateVerify: true,
			// 	ImportStateVerifyIgnore: []string{
			// 		"accept_language", "ignore_errors", "provisioning_artifact_name", "provisioning_parameters", "retain_physical_resources",
			// 	},
			// },
			// {
			// 	ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
			// 	ConfigDirectory:          config.StaticDirectory("testdata/ProvisionedProduct/tags/"),
			// 	ConfigVariables: config.Variables{
			// 		"rName":        config.StringVariable(rName),
			// 		"resource_tags": nil,
			// 	},
			// 	Check: resource.ComposeAggregateTestCheckFunc(
			// 		testAccCheckProvisionedProductExists(ctx, resourceName, &v),
			// 		resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
			// 		resource.TestCheckResourceAttr(resourceName, "tags_all.%", "0"),
			// 	),
			// },
			// {
			// 	ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
			// 	ConfigDirectory:          config.StaticDirectory("testdata/ProvisionedProduct/tags/"),
			// 	ConfigVariables: config.Variables{
			// 		"rName":        config.StringVariable(rName),
			// 		"resource_tags": nil,
			// 	},
			// 	ResourceName:      resourceName,
			// 	ImportState:       true,
			// 	ImportStateVerify: true,
			// 	ImportStateVerifyIgnore: []string{
			// 		"accept_language", "ignore_errors", "provisioning_artifact_name", "provisioning_parameters", "retain_physical_resources",
			// 	},
			// },
		},
	})
}

func TestAccServiceCatalogProvisionedProduct_tags_DefaultTags_nonOverlapping(t *testing.T) {
	ctx := acctest.Context(t)
	var v servicecatalog.ProvisionedProductDetail
	resourceName := "aws_servicecatalog_provisioned_product.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:   acctest.ErrorCheck(t, names.ServiceCatalogServiceID),
		CheckDestroy: testAccCheckProvisionedProductDestroy(ctx),
		Steps: []resource.TestStep{
			{
				ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
				ConfigDirectory:          config.StaticDirectory("testdata/ProvisionedProduct/tags_defaults/"),
				ConfigVariables: config.Variables{
					"rName": config.StringVariable(rName),
					"provider_tags": config.MapVariable(map[string]config.Variable{
						"providerkey1": config.StringVariable("providervalue1"),
					}),
					"resource_tags": config.MapVariable(map[string]config.Variable{
						"resourcekey1": config.StringVariable("resourcevalue1"),
					}),
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckProvisionedProductExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.resourcekey1", "resourcevalue1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.providerkey1", "providervalue1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.resourcekey1", "resourcevalue1"),
				),
			},
			{
				ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
				ConfigDirectory:          config.StaticDirectory("testdata/ProvisionedProduct/tags_defaults/"),
				ConfigVariables: config.Variables{
					"rName": config.StringVariable(rName),
					"provider_tags": config.MapVariable(map[string]config.Variable{
						"providerkey1": config.StringVariable("providervalue1"),
					}),
					"resource_tags": config.MapVariable(map[string]config.Variable{
						"resourcekey1": config.StringVariable("resourcevalue1"),
					}),
				},
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"accept_language", "ignore_errors", "provisioning_artifact_name", "provisioning_parameters", "retain_physical_resources",
				},
			},
			{
				ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
				ConfigDirectory:          config.StaticDirectory("testdata/ProvisionedProduct/tags_defaults/"),
				ConfigVariables: config.Variables{
					"rName": config.StringVariable(rName),
					"provider_tags": config.MapVariable(map[string]config.Variable{
						"providerkey1": config.StringVariable("providervalue1updated"),
					}),
					"resource_tags": config.MapVariable(map[string]config.Variable{
						"resourcekey1": config.StringVariable("resourcevalue1updated"),
						"resourcekey2": config.StringVariable("resourcevalue2"),
					}),
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckProvisionedProductExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "tags.resourcekey1", "resourcevalue1updated"),
					resource.TestCheckResourceAttr(resourceName, "tags.resourcekey2", "resourcevalue2"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.%", "3"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.providerkey1", "providervalue1updated"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.resourcekey1", "resourcevalue1updated"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.resourcekey2", "resourcevalue2"),
				),
			},
			{
				ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
				ConfigDirectory:          config.StaticDirectory("testdata/ProvisionedProduct/tags_defaults/"),
				ConfigVariables: config.Variables{
					"rName": config.StringVariable(rName),
					"provider_tags": config.MapVariable(map[string]config.Variable{
						"providerkey1": config.StringVariable("providervalue1updated"),
					}),
					"resource_tags": config.MapVariable(map[string]config.Variable{
						"resourcekey1": config.StringVariable("resourcevalue1updated"),
						"resourcekey2": config.StringVariable("resourcevalue2"),
					}),
				},
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"accept_language", "ignore_errors", "provisioning_artifact_name", "provisioning_parameters", "retain_physical_resources",
				},
			},
			// {
			// 	ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
			// 	ConfigDirectory:          config.StaticDirectory("testdata/ProvisionedProduct/tags/"),
			// 	ConfigVariables: config.Variables{
			// 		"rName":        config.StringVariable(rName),
			// 		"resource_tags": nil,
			// 	},
			// 	Check: resource.ComposeAggregateTestCheckFunc(
			// 		testAccCheckProvisionedProductExists(ctx, resourceName, &v),
			// 		resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
			// 		resource.TestCheckResourceAttr(resourceName, "tags_all.%", "0"),
			// 	),
			// },
			// {
			// 	ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
			// 	ConfigDirectory:          config.StaticDirectory("testdata/ProvisionedProduct/tags/"),
			// 	ConfigVariables: config.Variables{
			// 		"rName":        config.StringVariable(rName),
			// 		"resource_tags": nil,
			// 	},
			// 	ResourceName:      resourceName,
			// 	ImportState:       true,
			// 	ImportStateVerify: true,
			// 	ImportStateVerifyIgnore: []string{
			// 		"accept_language", "ignore_errors", "provisioning_artifact_name", "provisioning_parameters", "retain_physical_resources",
			// 	},
			// },
		},
	})
}

func TestAccServiceCatalogProvisionedProduct_tags_DefaultTags_overlapping(t *testing.T) {
	ctx := acctest.Context(t)
	var v servicecatalog.ProvisionedProductDetail
	resourceName := "aws_servicecatalog_provisioned_product.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:   acctest.ErrorCheck(t, names.ServiceCatalogServiceID),
		CheckDestroy: testAccCheckProvisionedProductDestroy(ctx),
		Steps: []resource.TestStep{
			{
				ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
				ConfigDirectory:          config.StaticDirectory("testdata/ProvisionedProduct/tags_defaults/"),
				ConfigVariables: config.Variables{
					"rName": config.StringVariable(rName),
					"provider_tags": config.MapVariable(map[string]config.Variable{
						"overlapkey1": config.StringVariable("providervalue1"),
					}),
					"resource_tags": config.MapVariable(map[string]config.Variable{
						"overlapkey1": config.StringVariable("resourcevalue1"),
					}),
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckProvisionedProductExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.overlapkey1", "resourcevalue1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.overlapkey1", "resourcevalue1"),
				),
			},
			{
				ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
				ConfigDirectory:          config.StaticDirectory("testdata/ProvisionedProduct/tags_defaults/"),
				ConfigVariables: config.Variables{
					"rName": config.StringVariable(rName),
					"provider_tags": config.MapVariable(map[string]config.Variable{
						"overlapkey1": config.StringVariable("providervalue1"),
					}),
					"resource_tags": config.MapVariable(map[string]config.Variable{
						"overlapkey1": config.StringVariable("resourcevalue1"),
					}),
				},
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"accept_language", "ignore_errors", "provisioning_artifact_name", "provisioning_parameters", "retain_physical_resources",
				},
			},
			{
				ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
				ConfigDirectory:          config.StaticDirectory("testdata/ProvisionedProduct/tags_defaults/"),
				ConfigVariables: config.Variables{
					"rName": config.StringVariable(rName),
					"provider_tags": config.MapVariable(map[string]config.Variable{
						"overlapkey1": config.StringVariable("providervalue1"),
						"overlapkey2": config.StringVariable("providervalue2"),
					}),
					"resource_tags": config.MapVariable(map[string]config.Variable{
						"overlapkey1": config.StringVariable("resourcevalue1"),
						"overlapkey2": config.StringVariable("resourcevalue2"),
					}),
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckProvisionedProductExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "tags.overlapkey1", "resourcevalue1"),
					resource.TestCheckResourceAttr(resourceName, "tags.overlapkey2", "resourcevalue2"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.overlapkey1", "resourcevalue1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.overlapkey2", "resourcevalue2"),
				),
			},
			{
				ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
				ConfigDirectory:          config.StaticDirectory("testdata/ProvisionedProduct/tags_defaults/"),
				ConfigVariables: config.Variables{
					"rName": config.StringVariable(rName),
					"provider_tags": config.MapVariable(map[string]config.Variable{
						"overlapkey1": config.StringVariable("providervalue1"),
						"overlapkey2": config.StringVariable("providervalue2"),
					}),
					"resource_tags": config.MapVariable(map[string]config.Variable{
						"overlapkey1": config.StringVariable("resourcevalue1"),
						"overlapkey2": config.StringVariable("resourcevalue2"),
					}),
				},
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"accept_language", "ignore_errors", "provisioning_artifact_name", "provisioning_parameters", "retain_physical_resources",
				},
			},
			{
				ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
				ConfigDirectory:          config.StaticDirectory("testdata/ProvisionedProduct/tags_defaults/"),
				ConfigVariables: config.Variables{
					"rName": config.StringVariable(rName),
					"provider_tags": config.MapVariable(map[string]config.Variable{
						"overlapkey1": config.StringVariable("providervalue1"),
					}),
					"resource_tags": config.MapVariable(map[string]config.Variable{
						"overlapkey1": config.StringVariable("resourcevalue2"),
					}),
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckProvisionedProductExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.overlapkey1", "resourcevalue2"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.overlapkey1", "resourcevalue2"),
				),
			},
			{
				ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
				ConfigDirectory:          config.StaticDirectory("testdata/ProvisionedProduct/tags_defaults/"),
				ConfigVariables: config.Variables{
					"rName": config.StringVariable(rName),
					"provider_tags": config.MapVariable(map[string]config.Variable{
						"overlapkey1": config.StringVariable("providervalue1"),
					}),
					"resource_tags": config.MapVariable(map[string]config.Variable{
						"overlapkey1": config.StringVariable("resourcevalue2"),
					}),
				},
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"accept_language", "ignore_errors", "provisioning_artifact_name", "provisioning_parameters", "retain_physical_resources",
				},
			},
		},
	})
}

func TestAccServiceCatalogProvisionedProduct_tags_DefaultTags_updateToProviderOnly(t *testing.T) {
	ctx := acctest.Context(t)
	var v servicecatalog.ProvisionedProductDetail
	resourceName := "aws_servicecatalog_provisioned_product.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:   acctest.ErrorCheck(t, names.ServiceCatalogServiceID),
		CheckDestroy: testAccCheckProvisionedProductDestroy(ctx),
		Steps: []resource.TestStep{
			{
				ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
				ConfigDirectory:          config.StaticDirectory("testdata/ProvisionedProduct/tags/"),
				ConfigVariables: config.Variables{
					"rName": config.StringVariable(rName),
					"resource_tags": config.MapVariable(map[string]config.Variable{
						"key1": config.StringVariable("value1"),
					}),
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckProvisionedProductExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.key1", "value1"),
				),
			},
			{
				ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
				ConfigDirectory:          config.StaticDirectory("testdata/ProvisionedProduct/tags_defaults/"),
				ConfigVariables: config.Variables{
					"rName": config.StringVariable(rName),
					"provider_tags": config.MapVariable(map[string]config.Variable{
						"key1": config.StringVariable("value1"),
					}),
					"resource_tags": nil,
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckProvisionedProductExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.key1", "value1"),
				),
			},
			{
				ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
				ConfigDirectory:          config.StaticDirectory("testdata/ProvisionedProduct/tags_defaults/"),
				ConfigVariables: config.Variables{
					"rName": config.StringVariable(rName),
					"provider_tags": config.MapVariable(map[string]config.Variable{
						"key1": config.StringVariable("value1"),
					}),
					"resource_tags": nil,
				},
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"accept_language", "ignore_errors", "provisioning_artifact_name", "provisioning_parameters", "retain_physical_resources",
				},
			},
		},
	})
}

func TestAccServiceCatalogProvisionedProduct_tags_DefaultTags_updateToResourceOnly(t *testing.T) {
	ctx := acctest.Context(t)
	var v servicecatalog.ProvisionedProductDetail
	resourceName := "aws_servicecatalog_provisioned_product.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:   acctest.ErrorCheck(t, names.ServiceCatalogServiceID),
		CheckDestroy: testAccCheckProvisionedProductDestroy(ctx),
		Steps: []resource.TestStep{
			{
				ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
				ConfigDirectory:          config.StaticDirectory("testdata/ProvisionedProduct/tags_defaults/"),
				ConfigVariables: config.Variables{
					"rName": config.StringVariable(rName),
					"provider_tags": config.MapVariable(map[string]config.Variable{
						"key1": config.StringVariable("value1"),
					}),
					"resource_tags": nil,
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckProvisionedProductExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.key1", "value1"),
				),
			},
			{
				ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
				ConfigDirectory:          config.StaticDirectory("testdata/ProvisionedProduct/tags/"),
				ConfigVariables: config.Variables{
					"rName": config.StringVariable(rName),
					"resource_tags": config.MapVariable(map[string]config.Variable{
						"key1": config.StringVariable("value1"),
					}),
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckProvisionedProductExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.key1", "value1"),
				),
			},
			{
				ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
				ConfigDirectory:          config.StaticDirectory("testdata/ProvisionedProduct/tags/"),
				ConfigVariables: config.Variables{
					"rName": config.StringVariable(rName),
					"resource_tags": config.MapVariable(map[string]config.Variable{
						"key1": config.StringVariable("value1"),
					}),
				},
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"accept_language", "ignore_errors", "provisioning_artifact_name", "provisioning_parameters", "retain_physical_resources",
				},
			},
		},
	})
}

func TestAccServiceCatalogProvisionedProduct_tags_DefaultTags_emptyResourceTag(t *testing.T) {
	t.Skip("Resource ProvisionedProduct does not support empty tags")

	ctx := acctest.Context(t)
	var v servicecatalog.ProvisionedProductDetail
	resourceName := "aws_servicecatalog_provisioned_product.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:   acctest.ErrorCheck(t, names.ServiceCatalogServiceID),
		CheckDestroy: testAccCheckProvisionedProductDestroy(ctx),
		Steps: []resource.TestStep{
			{
				ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
				ConfigDirectory:          config.StaticDirectory("testdata/ProvisionedProduct/tags_defaults/"),
				ConfigVariables: config.Variables{
					"rName": config.StringVariable(rName),
					"provider_tags": config.MapVariable(map[string]config.Variable{
						"key1": config.StringVariable("value1"),
					}),
					"resource_tags": config.MapVariable(map[string]config.Variable{
						"key1": config.StringVariable(""),
					}),
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckProvisionedProductExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", ""),
					resource.TestCheckResourceAttr(resourceName, "tags_all.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.key1", ""),
				),
			},
			{
				ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
				ConfigDirectory:          config.StaticDirectory("testdata/ProvisionedProduct/tags_defaults/"),
				ConfigVariables: config.Variables{
					"rName": config.StringVariable(rName),
					"provider_tags": config.MapVariable(map[string]config.Variable{
						"key1": config.StringVariable("value1"),
					}),
					"resource_tags": config.MapVariable(map[string]config.Variable{
						"key1": config.StringVariable(""),
					}),
				},
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"accept_language", "ignore_errors", "provisioning_artifact_name", "provisioning_parameters", "retain_physical_resources",
				},
			},
		},
	})
}

func TestAccServiceCatalogProvisionedProduct_tags_DefaultTags_nullOverlappingResourceTag(t *testing.T) {
	ctx := acctest.Context(t)
	var v servicecatalog.ProvisionedProductDetail
	resourceName := "aws_servicecatalog_provisioned_product.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:   acctest.ErrorCheck(t, names.ServiceCatalogServiceID),
		CheckDestroy: testAccCheckProvisionedProductDestroy(ctx),
		Steps: []resource.TestStep{
			{
				ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
				ConfigDirectory:          config.StaticDirectory("testdata/ProvisionedProduct/tags_defaults/"),
				ConfigVariables: config.Variables{
					"rName": config.StringVariable(rName),
					"provider_tags": config.MapVariable(map[string]config.Variable{
						"key1": config.StringVariable("providervalue1"),
					}),
					"resource_tags": config.MapVariable(map[string]config.Variable{
						"key1": nil,
					}),
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckProvisionedProductExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.key1", "providervalue1"),
				),
			},
			{
				ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
				ConfigDirectory:          config.StaticDirectory("testdata/ProvisionedProduct/tags_defaults/"),
				ConfigVariables: config.Variables{
					"rName": config.StringVariable(rName),
					"provider_tags": config.MapVariable(map[string]config.Variable{
						"key1": config.StringVariable("providervalue1"),
					}),
					"resource_tags": config.MapVariable(map[string]config.Variable{
						"key1": nil,
					}),
				},
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"accept_language", "ignore_errors", "provisioning_artifact_name", "provisioning_parameters", "retain_physical_resources",
				},
			},
		},
	})
}

func TestAccServiceCatalogProvisionedProduct_tags_DefaultTags_nullNonOverlappingResourceTag(t *testing.T) {
	ctx := acctest.Context(t)
	var v servicecatalog.ProvisionedProductDetail
	resourceName := "aws_servicecatalog_provisioned_product.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:   acctest.ErrorCheck(t, names.ServiceCatalogServiceID),
		CheckDestroy: testAccCheckProvisionedProductDestroy(ctx),
		Steps: []resource.TestStep{
			{
				ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
				ConfigDirectory:          config.StaticDirectory("testdata/ProvisionedProduct/tags_defaults/"),
				ConfigVariables: config.Variables{
					"rName": config.StringVariable(rName),
					"provider_tags": config.MapVariable(map[string]config.Variable{
						"providerkey1": config.StringVariable("providervalue1"),
					}),
					"resource_tags": config.MapVariable(map[string]config.Variable{
						"resourcekey1": nil,
					}),
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckProvisionedProductExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.providerkey1", "providervalue1"),
				),
			},
			{
				ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
				ConfigDirectory:          config.StaticDirectory("testdata/ProvisionedProduct/tags_defaults/"),
				ConfigVariables: config.Variables{
					"rName": config.StringVariable(rName),
					"provider_tags": config.MapVariable(map[string]config.Variable{
						"providerkey1": config.StringVariable("providervalue1"),
					}),
					"resource_tags": config.MapVariable(map[string]config.Variable{
						"resourcekey1": nil,
					}),
				},
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"accept_language", "ignore_errors", "provisioning_artifact_name", "provisioning_parameters", "retain_physical_resources",
				},
			},
		},
	})
}

func TestAccServiceCatalogProvisionedProduct_tags_ComputedTag_OnCreate(t *testing.T) {
	ctx := acctest.Context(t)
	var v servicecatalog.ProvisionedProductDetail
	resourceName := "aws_servicecatalog_provisioned_product.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:   acctest.ErrorCheck(t, names.ServiceCatalogServiceID),
		CheckDestroy: testAccCheckProvisionedProductDestroy(ctx),
		Steps: []resource.TestStep{
			{
				ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
				ConfigDirectory:          config.StaticDirectory("testdata/ProvisionedProduct/tagsComputed1/"),
				ConfigVariables: config.Variables{
					"rName":         config.StringVariable(rName),
					"unknownTagKey": config.StringVariable("computedkey1"),
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckProvisionedProductExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttrPair(resourceName, "tags.computedkey1", "null_resource.test", names.AttrID),
				),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionCreate),
						plancheck.ExpectUnknownValue(resourceName, tfjsonpath.New(names.AttrTags)),
					},
					PostApplyPreRefresh: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionNoop),
					},
					PostApplyPostRefresh: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionNoop),
					},
				},
			},
			{
				ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
				ConfigDirectory:          config.StaticDirectory("testdata/ProvisionedProduct/tagsComputed1/"),
				ConfigVariables: config.Variables{
					"rName":         config.StringVariable(rName),
					"unknownTagKey": config.StringVariable("computedkey1"),
				},
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"accept_language", "ignore_errors", "provisioning_artifact_name", "provisioning_parameters", "retain_physical_resources",
				},
			},
		},
	})
}

func TestAccServiceCatalogProvisionedProduct_tags_ComputedTag_OnUpdate_Add(t *testing.T) {
	ctx := acctest.Context(t)
	var v servicecatalog.ProvisionedProductDetail
	resourceName := "aws_servicecatalog_provisioned_product.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:   acctest.ErrorCheck(t, names.ServiceCatalogServiceID),
		CheckDestroy: testAccCheckProvisionedProductDestroy(ctx),
		Steps: []resource.TestStep{
			{
				ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
				ConfigDirectory:          config.StaticDirectory("testdata/ProvisionedProduct/tags/"),
				ConfigVariables: config.Variables{
					"rName": config.StringVariable(rName),
					"resource_tags": config.MapVariable(map[string]config.Variable{
						"key1": config.StringVariable("value1"),
					}),
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckProvisionedProductExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
				),
			},
			{
				ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
				ConfigDirectory:          config.StaticDirectory("testdata/ProvisionedProduct/tagsComputed2/"),
				ConfigVariables: config.Variables{
					"rName":         config.StringVariable(rName),
					"unknownTagKey": config.StringVariable("computedkey1"),
					"knownTagKey":   config.StringVariable("key1"),
					"knownTagValue": config.StringVariable("value1"),
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckProvisionedProductExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
					resource.TestCheckResourceAttrPair(resourceName, "tags.computedkey1", "null_resource.test", names.AttrID),
				),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
						plancheck.ExpectUnknownValue(resourceName, tfjsonpath.New(names.AttrTags)),
					},
					PostApplyPreRefresh: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionNoop),
					},
					PostApplyPostRefresh: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionNoop),
					},
				},
			},
			{
				ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
				ConfigDirectory:          config.StaticDirectory("testdata/ProvisionedProduct/tagsComputed2/"),
				ConfigVariables: config.Variables{
					"rName":         config.StringVariable(rName),
					"unknownTagKey": config.StringVariable("computedkey1"),
					"knownTagKey":   config.StringVariable("key1"),
					"knownTagValue": config.StringVariable("value1"),
				},
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"accept_language", "ignore_errors", "provisioning_artifact_name", "provisioning_parameters", "retain_physical_resources",
				},
			},
		},
	})
}

func TestAccServiceCatalogProvisionedProduct_tags_ComputedTag_OnUpdate_Replace(t *testing.T) {
	ctx := acctest.Context(t)
	var v servicecatalog.ProvisionedProductDetail
	resourceName := "aws_servicecatalog_provisioned_product.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:   acctest.ErrorCheck(t, names.ServiceCatalogServiceID),
		CheckDestroy: testAccCheckProvisionedProductDestroy(ctx),
		Steps: []resource.TestStep{
			{
				ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
				ConfigDirectory:          config.StaticDirectory("testdata/ProvisionedProduct/tags/"),
				ConfigVariables: config.Variables{
					"rName": config.StringVariable(rName),
					"resource_tags": config.MapVariable(map[string]config.Variable{
						"key1": config.StringVariable("value1"),
					}),
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckProvisionedProductExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
				),
			},
			{
				ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
				ConfigDirectory:          config.StaticDirectory("testdata/ProvisionedProduct/tagsComputed1/"),
				ConfigVariables: config.Variables{
					"rName":         config.StringVariable(rName),
					"unknownTagKey": config.StringVariable("key1"),
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckProvisionedProductExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttrPair(resourceName, "tags.key1", "null_resource.test", names.AttrID),
				),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
						plancheck.ExpectUnknownValue(resourceName, tfjsonpath.New(names.AttrTags)),
					},
					PostApplyPreRefresh: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionNoop),
					},
					PostApplyPostRefresh: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionNoop),
					},
				},
			},
			{
				ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
				ConfigDirectory:          config.StaticDirectory("testdata/ProvisionedProduct/tagsComputed1/"),
				ConfigVariables: config.Variables{
					"rName":         config.StringVariable(rName),
					"unknownTagKey": config.StringVariable("key1"),
				},
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"accept_language", "ignore_errors", "provisioning_artifact_name", "provisioning_parameters", "retain_physical_resources",
				},
			},
		},
	})
}
