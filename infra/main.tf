provider "aws" {
  region = "us-east-1" # Substitua pela sua região
}


variable "lambda_zip_path" {
  description = "Caminho para o arquivo zip da Lambda"
  type        = string
  default     = "${github.workspace}/app/lambda-deployment-package.zip"
}

resource "aws_lambda_function" "example" {
  function_name = "example"
  role         = aws_iam_role.example.arn
  handler      = "main"
  runtime      = "go1.x"
  filename     = lambda_zip_path


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