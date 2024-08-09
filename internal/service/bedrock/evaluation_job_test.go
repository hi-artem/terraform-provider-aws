// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrock_test

// **PLEASE DELETE THIS AND ALL TIP COMMENTS BEFORE SUBMITTING A PR FOR REVIEW!**
//
// TIP: ==== INTRODUCTION ====
// Thank you for trying the skaff tool!
//
// You have opted to include these helpful comments. They all include "TIP:"
// to help you find and remove them when you're done with them.
//
// While some aspects of this file are customized to your input, the
// scaffold tool does *not* look at the AWS API and ensure it has correct
// function, structure, and variable names. It makes guesses based on
// commonalities. You will need to make significant adjustments.
//
// In other words, as generated, this is a rough outline of the work you will
// need to do. If something doesn't make sense for your situation, get rid of
// it.

import (
	// TIP: ==== IMPORTS ====
	// This is a common set of imports but not customized to your code since
	// your code hasn't been written yet. Make sure you, your IDE, or
	// goimports -w <file> fixes these imports.
	//
	// The provider linter wants your imports to be in two groups: first,
	// standard library (i.e., "fmt" or "strings"), second, everything else.
	//
	// Also, AWS Go SDK v2 may handle nested structures differently than v1,
	// using the services/bedrock/types package. If so, you'll
	// need to import types and reference the nested types, e.g., as
	// types.<Type Name>.
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/bedrock"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/names"
	// TIP: You will often need to import the package that this test file lives
	// in. Since it is in the "test" context, it must import the package to use
	// any normal context constants, variables, or functions.
)

func TestAccBedrockEvaluationJob_basic(t *testing.T) {
	ctx := acctest.Context(t)
	// TIP: This is a long-running test guard for tests that run longer than
	// 300s (5 min) generally.
	if testing.Short() {
		t.Skip("skipping long-running test in short mode")
	}

	//var evaluationjob bedrock.DescribeEvaluationJobResponse
	rName := name_regex(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix))
	resourceName := "aws_bedrock_evaluation_job.test"
	bName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	//bucketName := "aws_s3_bucket.test"
	modelName := "aws_bedrock_foundation_model.test"
	iamName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	bucketAccessName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	modelAccessName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	transferRoleName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	sagemakerName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	sagemakerCreate := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
			acctest.PreCheckPartitionHasService(t, names.BedrockEndpointID)
			testAccPreCheck(ctx, t)
		},
		ErrorCheck:               acctest.ErrorCheck(t, names.BedrockServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccEvaluationJobConfig_basic(iamName, rName, bName, bucketAccessName, modelAccessName, transferRoleName, sagemakerName, sagemakerCreate, acctest.CtKey1, acctest.CtValue1, acctest.CtKey2, acctest.CtValue2),
				Check: resource.ComposeTestCheckFunc(
					//testAccCheckEvaluationJobExists(ctx, resourceName, &evaluationjob),
					resource.TestCheckResourceAttrSet(resourceName, "creation_time"),
					//resource.TestCheckResourceAttrPair(resourceName, "customer_cencryption_key_id", ""),
					// eval config
					resource.TestCheckResourceAttr(resourceName, "evaluation_config.automated.dataset_metric_configs.metric_names.0", "Builtin.Accuracy"),
					resource.TestCheckResourceAttr(resourceName, "evaluation_config.automated.dataset_metric_configs.task_type", "Question and Answer"),
					resource.TestCheckResourceAttr(resourceName, "evaluation_config.automated.dataset_metric_configs.dataset.name", "Builtin.BoolQ"),
					resource.TestCheckResourceAttrSet(resourceName, "evaluation_config.automated.dataset_metric_configs.dataset.dataset_location.s3_uri"),
					// eval config end
					// inf config
					resource.TestCheckResourceAttrPair(resourceName, "inference_config.models.bedrock_model.inference_params", modelName, "inference_types_supported.0"),
					resource.TestCheckResourceAttrPair(resourceName, "inference_config.models.bedrock_model.model_identifier", modelName, names.AttrID),
					// inf config end
					resource.TestCheckResourceAttrSet(resourceName, names.AttrARN),
					resource.TestCheckResourceAttr(resourceName, names.AttrName, rName),
					resource.TestCheckResourceAttr(resourceName, names.AttrType, "Summarization"),
					resource.TestCheckResourceAttrSet(resourceName, "output_data_config.0.s3_uri"),
					resource.TestCheckResourceAttrSet(resourceName, "role_arn"),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"apply_immediately", "user"},
			},
		},
	})
}
func TestAccBedrockEvaluationJob_encryption(t *testing.T) {
	ctx := acctest.Context(t)
	// TIP: This is a long-running test guard for tests that run longer than
	// 300s (5 min) generally.
	if testing.Short() {
		t.Skip("skipping long-running test in short mode")
	}

	//var evaluationjob bedrock.DescribeEvaluationJobResponse
	rName := name_regex(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix))
	resourceName := "aws_bedrock_evaluation_job.test"
	bName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	//bucketName := "aws_s3_bucket.test"
	modelName := "aws_bedrock_foundation_model.test"
	iamName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	bucketAccessName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	modelAccessName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	transferRoleName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	sagemakerName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	sagemakerCreate := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	keyName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
			acctest.PreCheckPartitionHasService(t, names.BedrockEndpointID)
			testAccPreCheck(ctx, t)
		},
		ErrorCheck:               acctest.ErrorCheck(t, names.BedrockServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccEvaluationJobConfig_encryption(iamName, rName, bName, bucketAccessName, modelAccessName, transferRoleName, sagemakerName, sagemakerCreate, keyName, acctest.CtKey1, acctest.CtValue1, acctest.CtKey2, acctest.CtValue2),
				Check: resource.ComposeTestCheckFunc(
					//testAccCheckEvaluationJobExists(ctx, resourceName, &evaluationjob),
					resource.TestCheckResourceAttrSet(resourceName, "creation_time"),
					resource.TestCheckResourceAttrPair(resourceName, "customer_encryption_key_id", keyName, names.AttrID),
					//resource.TestCheckResourceAttrPair(resourceName, "customer_cencryption_key_id", ""),
					// eval config
					resource.TestCheckResourceAttr(resourceName, "evaluation_config.automated.dataset_metric_configs.metric_names.0", "Builtin.Accuracy"),
					resource.TestCheckResourceAttr(resourceName, "evaluation_config.automated.dataset_metric_configs.task_type", "Question and Answer"),
					resource.TestCheckResourceAttr(resourceName, "evaluation_config.automated.dataset_metric_configs.dataset.name", "Builtin.BoolQ"),
					resource.TestCheckResourceAttrSet(resourceName, "evaluation_config.automated.dataset_metric_configs.dataset.dataset_location.s3_uri"),
					// eval config end
					// inf config
					resource.TestCheckResourceAttrPair(resourceName, "inference_config.models.bedrock_model.inference_params", modelName, "inference_types_supported.0"),
					resource.TestCheckResourceAttrPair(resourceName, "inference_config.models.bedrock_model.model_identifier", modelName, names.AttrID),
					// inf config end
					resource.TestCheckResourceAttrSet(resourceName, names.AttrARN),
					resource.TestCheckResourceAttr(resourceName, names.AttrName, rName),
					resource.TestCheckResourceAttr(resourceName, names.AttrType, "Summarization"),
					resource.TestCheckResourceAttrSet(resourceName, "output_data_config.0.s3_uri"),
					resource.TestCheckResourceAttrSet(resourceName, "role_arn"),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"apply_immediately", "user"},
			},
		},
	})
}

func name_regex(in string) string {
	return in
	var out string
	for index, char := range in {
		if index == 1 {
			out += "-"
		} else if index == 2 {
			out += "*"
		} else {
			out += string(char)
		}
	}
	for len(in) < 61 {
		in += "e"
	}
	return out
}

/*
func TestAccBedrockEvaluationJob_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	if testing.Short() {
		t.Skip("skipping long-running test in short mode")
	}

	var evaluationjob bedrock.DescribeEvaluationJobResponse
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_bedrock_evaluation_job.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
			acctest.PreCheckPartitionHasService(t, names.BedrockEndpointID)
			testAccPreCheck(t)
		},
		ErrorCheck:               acctest.ErrorCheck(t, names.BedrockServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckEvaluationJobDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccEvaluationJobConfig_basic(rName, testAccEvaluationJobVersionNewer),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEvaluationJobExists(ctx, resourceName, &evaluationjob),
					// TIP: The Plugin-Framework disappears helper is similar to the Plugin-SDK version,
					// but expects a new resource factory function as the third argument. To expose this
					// private function to the testing package, you may need to add a line like the following
					// to exports_test.go:
					//
					//   var ResourceEvaluationJob = newResourceEvaluationJob
					acctest.CheckFrameworkResourceDisappears(ctx, acctest.Provider, tfbedrock.ResourceEvaluationJob, resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}
*/
/*
func testAccCheckEvaluationJobExists(ctx context.Context, name string, evaluationjob *bedrock.DescribeEvaluationJobResponse) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return create.Error(names.Bedrock, create.ErrActionCheckingExistence, tfbedrock.ResNameEvaluationJob, name, errors.New("not found"))
		}

		if rs.Primary.ID == "" {
			return create.Error(names.Bedrock, create.ErrActionCheckingExistence, tfbedrock.ResNameEvaluationJob, name, errors.New("not set"))
		}

		conn := acctest.Provider.Meta().(*conns.AWSClient).BedrockClient(ctx)
		resp, err := conn.DescribeEvaluationJob(ctx, &bedrock.DescribeEvaluationJobInput{
			EvaluationJobId: aws.String(rs.Primary.ID),
		})

		if err != nil {
			return create.Error(names.Bedrock, create.ErrActionCheckingExistence, tfbedrock.ResNameEvaluationJob, rs.Primary.ID, err)
		}

		*evaluationjob = *resp

		return nil
	}
}
*/

func testAccPreCheck(ctx context.Context, t *testing.T) {
	conn := acctest.Provider.Meta().(*conns.AWSClient).BedrockClient(ctx)

	input := &bedrock.ListEvaluationJobsInput{}
	_, err := conn.ListEvaluationJobs(ctx, input)

	if acctest.PreCheckSkipError(err) {
		t.Skipf("skipping acceptance testing: %s", err)
	}
	if err != nil {
		t.Fatalf("unexpected PreCheck error: %s", err)
	}
}

func testAccEvaluationJobConfig_base(iamName, bucketName, bucketAccessName, modelAccessName, transferRoleName, sagemakerName, sagemakerCreate string) string {
	return fmt.Sprintf(`

data "aws_caller_identity" "test" {}
data "aws_region" "test" {}
data "aws_partition" "test" {}
data "aws_bedrock_foundation_models" "test" {}
data "aws_bedrock_foundation_model" "test" {
  model_id = "amazon.titan-text-lite-v1"
}

resource "aws_s3_bucket" "test" {
  bucket = %[2]q
  cors_rule {
    allowed_headers = ["*"]
    allowed_methods = ["GET", "PUT", "POST", "DELETE"]
    allowed_origins = ["*"]
    expose_headers  = ["Access-Control-Allow-Origin"]
  }
}

resource "aws_s3_object" "output" {
	bucket = aws_s3_bucket.test.id
	key = "data/dataset.jsonl"
	source = "dataset.jsonl"	
}

resource "aws_s3_object" "dataset" {
	bucket = aws_s3_bucket.test.id
	key = "data/dataset.jsonl"
	source = "dataset.jsonl"	
}
	resource "aws_iam_role" "test" {
  name = %[1]q

  assume_role_policy = jsonencode(
  {
    "Version": "2012-10-17",
    "Statement": [{
        "Sid": "AllowBedrockToAssumeRole",
        "Effect": "Allow",
        "Principal": {
            "Service": "bedrock.amazonaws.com"
        },
        "Action": "sts:AssumeRole",
        "Condition": {
            "StringEquals": {
        		"aws:SourceAccount": "${data.aws_caller_identity.test.user_id}"
            },
            "ArnEquals": {
				"aws:SourceArn" : "arn:aws:bedrock:${data.aws_region.test.name}:${data.aws_caller_identity.test.user_id}:evaluation-job/*"
            }
        }
    }]
}
)
}



resource "aws_iam_policy" "test" {
  name = %[1]q
  policy = jsonencode({
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "BedrockConsole",
            "Effect": "Allow",
            "Action": [
               "bedrock:CreateEvaluationJob",
               "bedrock:GetEvaluationJob",
               "bedrock:ListEvaluationJobs",
               "bedrock:StopEvaluationJob",
               "bedrock:GetCustomModel",
               "bedrock:ListCustomModels",
               "bedrock:CreateProvisionedModelThroughput",
               "bedrock:UpdateProvisionedModelThroughput",
               "bedrock:GetProvisionedModelThroughput",
               "bedrock:ListProvisionedModelThroughputs",
               "bedrock:ListTagsForResource",
               "bedrock:UntagResource",
               "bedrock:TagResource"
            ],
            "Resource": "*"
        },
        {
            "Sid": "AllowConsoleS3AccessForModelEvaluation",
            "Effect": "Allow",
            "Action": [
              "s3:GetObject",
              "s3:GetBucketCORS",
              "s3:ListBucket",
              "s3:ListBucketVersions",
              "s3:GetBucketLocation"
            ],
            "Resource": "*"
        }
    ]
  })
}
resource "aws_iam_role_policy_attachment" "test" {
	role = aws_iam_role.test.name
	policy_arn = aws_iam_policy.test.arn
}

resource "aws_iam_policy" "bucket_access" {
  name = %[3]q

  policy = jsonencode({
"Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "AllowAccessToCustomDatasets",
            "Effect": "Allow",
            "Action": [
                "s3:GetObject",
                "s3:ListBucket"
            ],
            "Resource": [
            "arn:aws:s3:::${aws_s3_bucket.test.bucket}",
            "arn:aws:s3:::${aws_s3_bucket.test.bucket}/${aws_s3_object.dataset.key}",
			"*"
			]
        },
        {
            "Sid": "AllowAccessToOutputBucket",
            "Effect": "Allow",
            "Action": [
                "s3:GetObject",
                "s3:ListBucket",
                "s3:PutObject",
                "s3:GetBucketLocation",
                "s3:AbortMultipartUpload",
                "s3:ListBucketMultipartUploads"
            ],
            "Resource": [
            "arn:aws:s3:::${aws_s3_bucket.test.bucket}",
			"arn:aws:s3:::${aws_s3_bucket.test.bucket}/${aws_s3_object.output.key}",
			"*"            
			]
        }
    ]
  })
}
resource "aws_iam_role_policy_attachment" "bucket_access" {
	role = aws_iam_role.test.name
	policy_arn = aws_iam_policy.bucket_access.arn
}

  resource "aws_iam_policy" "model_access" {
  name = %[4]q

  policy = jsonencode({
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "AllowSpecificModels",
            "Effect": "Allow",
            "Action": [
                "bedrock:InvokeModel",
                "bedrock:InvokeModelWithResponseStream",
				"bedrock:CreateModelInvocationJob"
            ],
            "Resource": [
				"arn:aws:bedrock:${data.aws_region.test.name}::foundation-model/amazon.titan-text-lite-v1",
				"*"
            ]
        }
    ]
  })
}
resource "aws_iam_role_policy_attachment" "model_access" {
	role = aws_iam_role.test.name
	policy_arn = aws_iam_policy.model_access.arn
}

resource "aws_iam_policy" "pass" {
name = %[5]q

policy = jsonencode({
    "Version": "2012-10-17",
    "Statement": [{
        "Effect": "Allow",
        "Action": [
            "iam:GetRole",
            "iam:PassRole"
        ],
        "Resource": [
		"arn:aws:iam::${data.aws_caller_identity.test.account_id}:role/bedrock-*",
		"*"
		]
    }]
})
}
resource "aws_iam_role_policy_attachment" "pass" {
	role = aws_iam_role.test.name
	policy_arn = aws_iam_policy.pass.arn
}

resource "aws_iam_policy" "sagemaker_create" {
	name = %[7]q
	policy = jsonencode ({
	"Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "ManageHumanLoops",
            "Effect": "Allow",
            "Action": [
                "sagemaker:StartHumanLoop",
                "sagemaker:DescribeFlowDefinition",
                "sagemaker:DescribeHumanLoop",
                "sagemaker:StopHumanLoop",
                "sagemaker:DeleteHumanLoop"
            ],
            "Resource": "*"
        }
    ]	
})
}

resource "aws_iam_role_policy_attachment" "sagemaker_create" {
	role = aws_iam_role.test.name
	policy_arn = aws_iam_policy.sagemaker_create.arn
}





	`, iamName, bucketName, bucketAccessName, modelAccessName, transferRoleName, sagemakerName, sagemakerCreate)
}

func testAccEvaluationJobConfig_basic(iamName, jobName, bucketName, bucketAccessName, modelAccessName, transferRoleName, sagemakerName, sagemakerCreate, tagKey1, tagValue1, tagKey2, tagValue2 string) string {
	return acctest.ConfigCompose(testAccEvaluationJobConfig_base(iamName, bucketName, bucketAccessName, modelAccessName, transferRoleName, sagemakerName, sagemakerCreate), fmt.Sprintf(`


resource "aws_bedrock_evaluation_job" "test" {

  evaluation_config {
    automated {
        dataset_metric_configs {
          dataset {
		    name = "Builtin.BoolQ"
          }
          metric_names = ["Builtin.Accuracy"]
          task_type    = "QuestionAndAnswer"
        }
    }
  }

  inference_config {
    models {
      bedrock_model { 
        inference_params = tolist(data.aws_bedrock_foundation_model.test.inference_types_supported)[0]
        model_identifier = data.aws_bedrock_foundation_model.test.id
		}
    }
  }

  description = "test"
  name        = %[1]q

  output_data_config {
    s3_uri = "s3://${aws_s3_bucket.test.bucket}/bedrock/evaluation_jobs"
  }

  tags = {
	%[2]q = %[3]q
	%[4]q = %[5]q
  }

  role_arn = aws_iam_role.test.arn
}
`, jobName, tagKey1, tagValue1, tagKey2, tagValue2))
}

func testAccEvaluationJobConfig_encryptionbase(iamName, bucketName, bucketAccessName, modelAccessName, transferRoleName, sagemakerName, sagemakerCreate, keyName string) string {
	return acctest.ConfigCompose(testAccEvaluationJobConfig_base(iamName, bucketName, bucketAccessName, modelAccessName, transferRoleName, sagemakerName, sagemakerCreate), fmt.Sprintf(
		`
		resource "aws_kms_key" "test" {
			description             = %[1]q
			deletion_window_in_days = 7
		}
				
		resource "aws_iam_policy" "kms" {
		name = %[1]q 
		policy = jsonencode({
		    "Version": "2012-10-17",
			"Statement": [
			{ 
				"Sid": "CustomKMSKeyProvidedToBedrock",
				"Effect": "Allow",
				"Action": [
					"kms:Decrypt",
					"kms:GenerateDataKey"
				],
				"Resource": [
				"arn:aws:kms:${data.aws_region.test.name}:${data.aws_caller_identity.test.account_id}:key/${aws_kms_key.test.id}"
				],
				"Condition": {
					"StringEquals": {
						"kms:ViaService": "s3.${data.aws_region.test.name}.amazonaws.com"
					}
				}
			},
			{ 
				"Sid": "CustomKMSDescribeKeyProvidedToBedrock",
				"Effect": "Allow",
				"Action": [
					"kms:DescribeKey"
				],
				"Resource": [
				"arn:aws:kms:${data.aws_region.test.name}:${data.aws_caller_identity.test.account_id}:key/${aws_kms_key.test.id}"
				]
			}
		]
	})
	}
	resource "aws_iam_role_policy_attachment" "kms" {
		role = aws_iam_role.test.name
		policy_arn = aws_iam_policy.kms.arn
	}
	resource "aws_kms_key_policy" "test" {
		key_id = aws_kms_key.test.id
		policy = jsonencode(
		{
			"Version": "2012-10-17",
			"Id": "key-consolepolicy-3",
			"Statement": [
				{
					"Sid": "EnableIAMUserPermissions",
					"Effect": "Allow",
					"Principal": {
						"AWS": "arn:aws:iam::${data.aws_caller_identity.test.account_id}:root"
					},
					"Action": "kms:*",
					"Resource": "*"
				},
				{
					"Effect": "Allow",
					"Principal": {
						"Service": "bedrock.amazonaws.com"
					},
					"Action": [
						"kms:GenerateDataKey",
						"kms:Decrypt",
						"kms:DescribeKey"
					],
					"Resource": "*",
					"Condition": {
						"StringLike": {
							"kms:EncryptionContext:evaluationJobArn": "arn:aws:bedrock:${data.aws_region.test.name}:${data.aws_caller_identity.test.account_id}:evaluation-job/*",
							"aws:SourceArn": "arn:aws:bedrock:${data.aws_region.test.name}:${data.aws_caller_identity.test.account_id}:evaluation-job/*"
						}
					}
				}
			]
			}
		)	
	}

	
	`, keyName))

}

func testAccEvaluationJobConfig_encryption(iamName, jobName, bucketName, bucketAccessName, modelAccessName, transferRoleName, sagemakerName, sagemakerCreate, keyName, tagKey1, tagValue1, tagKey2, tagValue2 string) string {
	return acctest.ConfigCompose(testAccEvaluationJobConfig_encryptionbase(iamName, bucketName, bucketAccessName, modelAccessName, transferRoleName, sagemakerName, sagemakerCreate, keyName), fmt.Sprintf(`


resource "aws_bedrock_evaluation_job" "test" {

  evaluation_config {
    automated {
        dataset_metric_configs {
          dataset {
		    name = "BoolQ"
			dataset_location {
				s3_uri = "s3://${aws_s3_bucket.test.bucket}/${aws_s3_object.dataset.key}"
			}
          }
          metric_names = ["Builtin.Accuracy"]
          task_type    = "QuestionAndAnswer"
        }
    }
  }

  inference_config {
    models {
      bedrock_model { 
        inference_params = tolist(data.aws_bedrock_foundation_model.test.inference_types_supported)[0]
        model_identifier = data.aws_bedrock_foundation_model.test.id
		}
    }
  }

  customer_encryption_key_id = aws_kms_key.test.arn
  description = "test"
  name        = %[1]q

  output_data_config {
    s3_uri = "s3://${aws_s3_bucket.test.bucket}/bedrock/evaluation_jobs"
  }

  tags = {
	%[2]q = %[3]q
	%[4]q = %[5]q
  }

  role_arn = aws_iam_role.test.arn
}
`, jobName, tagKey1, tagValue1, tagKey2, tagValue2))
}