---
subcategory: "Bedrock AgentCore"
layout: "aws"
page_title: "AWS: aws_bedrockagentcore_resource_policy"
description: |-
  Manages an AWS Bedrock AgentCore Resource Policy.
---

# Resource: aws_bedrockagentcore_resource_policy

Manages a resource-based policy for an Amazon Bedrock AgentCore resource. Resource-based policies allow you to control which principals (AWS accounts, IAM users, or IAM roles) can invoke and manage your Bedrock AgentCore resources, such as Agent Runtimes, Agent Runtime Endpoints, and Gateways.

## Example Usage

### Gateway Resource Policy (AWS IAM)

```terraform
data "aws_caller_identity" "current" {}

data "aws_iam_policy_document" "example" {
  statement {
    effect = "Allow"
    principals {
      type        = "AWS"
      identifiers = [data.aws_caller_identity.current.account_id]
    }
    actions   = ["bedrock-agentcore:InvokeGateway"]
    resources = [aws_bedrockagentcore_gateway.example.gateway_arn]
  }
}

resource "aws_bedrockagentcore_resource_policy" "example" {
  resource_arn = aws_bedrockagentcore_gateway.example.gateway_arn
  policy       = data.aws_iam_policy_document.example.json
}
```

### Agent Runtime Resource Policy (Cross-Account Access)

```terraform
data "aws_iam_policy_document" "example" {
  statement {
    effect = "Allow"
    principals {
      type        = "AWS"
      identifiers = ["arn:aws:iam::123456789012:role/CrossAccountRole"]
    }
    actions   = ["bedrock-agentcore:InvokeAgentRuntime"]
    resources = [aws_bedrockagentcore_agent_runtime.example.agent_runtime_arn]
  }
}

resource "aws_bedrockagentcore_resource_policy" "example" {
  resource_arn = aws_bedrockagentcore_agent_runtime.example.agent_runtime_arn
  policy       = data.aws_iam_policy_document.example.json
}
```

### Gateway Resource Policy (OAuth Authentication)

When the Gateway is configured with OAuth authentication, you must use a wildcard principal:

```terraform
data "aws_iam_policy_document" "example" {
  statement {
    sid    = "AllowOAuthFromVPC"
    effect = "Allow"
    principals {
      type        = "*"
      identifiers = ["*"]
    }
    actions   = ["bedrock-agentcore:InvokeGateway"]
    resources = [aws_bedrockagentcore_gateway.example.gateway_arn]
    condition {
      test     = "StringEquals"
      variable = "aws:SourceVpc"
      values   = ["vpc-1a2b3c4d"]
    }
  }
}

resource "aws_bedrockagentcore_resource_policy" "example" {
  resource_arn = aws_bedrockagentcore_gateway.example.gateway_arn
  policy       = data.aws_iam_policy_document.example.json
}
```

## Argument Reference

The following arguments are required:

* `resource_arn` - (Required, Forces new resource) ARN of the Bedrock AgentCore resource to attach the policy to. This can be an Agent Runtime ARN, Agent Runtime Endpoint ARN, or Gateway ARN.
* `policy` - (Required) JSON policy document to attach to the resource. Use the [`aws_iam_policy_document`](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/iam_policy_document) data source to define the policy.

The following arguments are optional:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - ARN of the resource the policy is attached to (same as `resource_arn`).

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Bedrock AgentCore Resource Policy using the resource ARN. For example:

```terraform
import {
  to = aws_bedrockagentcore_resource_policy.example
  id = "arn:aws:bedrock-agentcore:us-east-1:123456789012:gateway/example-gateway-id"
}
```

Using `terraform import`, import Bedrock AgentCore Resource Policy using the resource ARN. For example:

```console
% terraform import aws_bedrockagentcore_resource_policy.example arn:aws:bedrock-agentcore:us-east-1:123456789012:gateway/example-gateway-id
```
