provider "aws" {
  region = "us-east-1" # Substitua pela sua região
}

variable "TF_LAMBDA_ZIP_PATH" {
  type = string
}

resource "aws_lambda_function" "example" {
  function_name = "example"
  role         = aws_iam_role.example.arn
  handler      = "main"
  runtime      = "go1.x"

  filename     = var.TF_LAMBDA_ZIP_PATH # Recupera o zip da lambda disponibilizado pela esteira

  environment {
    variables = {
      EXAMPLE_ENV_VAR = "example"
    }
  }
}

resource "aws_iam_role" "example" {
  name = "example"

  assume_role_policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Action = "sts:AssumeRole",
        Effect = "Allow",
        Principal = {
          Service = "lambda.amazonaws.com"
        }
      }
    ]
  })
}

resource "aws_iam_role_policy" "example" {
  name        = "DynamoDBAccessPolicy"
  role        = aws_iam_role.example.name

  policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Action   = [
          "dynamodb:GetItem",
          "dynamodb:PutItem"
          // Adicione outras ações conforme necessário
        ],
        Effect   = "Allow",
        Resource = "arn:aws:dynamodb:us-east-1:101478099523:table/ClienteAppFastfood"
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "example" {
  policy_arn = aws_iam_role_policy.example.arn
  role       = aws_iam_role.example.name
}

resource "aws_lambda_permission" "example" {
  statement_id  = "AllowExecutionFromLambda"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.example.arn
  principal     = "dynamodb.amazonaws.com"
  source_arn    = "arn:aws:dynamodb:us-east-1:101478099523:table/ClienteAppFastfood"
}
