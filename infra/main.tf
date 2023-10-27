provider "aws" {
  region = "us-east-1" # Substitua pela sua região
}

resource "aws_lambda_function" "example" {
  function_name = "example"
  role         = aws_iam_role.example.arn
  handler      = "main"
  runtime      = "go1.x"

  filename     = "${path.module}/lambda-deployment-package.zip" # Recupera o zip da lambda disponibilizado pela esteira

  source_code_hash = filebase64sha256("${path.module}/lambda-deployment-package.zip")

  environment {
    variables = {
      EXAMPLE_ENV_VAR = "example"
    }
  }
}

resource "aws_iam_role" "example" {
  name = "example"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "lambda.amazonaws.com"
        }
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "example" {
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
  role       = aws_iam_role.example.name
}