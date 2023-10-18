provider "aws" {
  region = "us-east-1"  # Substitua pela região desejada
}

resource "aws_lambda_function" "example" {
  function_name = "example-lambda"
  filename      = "example_lambda.zip"  # O arquivo ZIP da sua função
  role         = aws_iam_role.lambda_exec_role.arn
  handler      = "com.example.LambdaHandler::handleRequest"  # Substitua com o nome da sua classe e método

  runtime = "java17"  # Define a versão do JDK

  environment {
    variables = {
      key1 = "value1",
      key2 = "value2",
    }
  }
}

resource "aws_iam_role" "lambda_exec_role" {
  name = "example-lambda-role"

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

resource "aws_iam_policy_attachment" "lambda_execution" {
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
  roles      = [aws_iam_role.lambda_exec_role.name]
}

data "archive_file" "example_lambda" {
  type        = "zip"
  source_dir  = "example_lambda_code"  # Diretório contendo o código da função Lambda
  output_path = "example_lambda.zip"
}
