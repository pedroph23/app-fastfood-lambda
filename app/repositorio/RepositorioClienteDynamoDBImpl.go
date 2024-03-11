package repositorio

import (
	"fmt"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/pedroph23/app-fastfood-lambda/app/dominio"
)

type RepositorioClienteDynamoDBImpl struct {
	svc *dynamodb.DynamoDB
}

func NewRepositorioClienteImpl() *RepositorioClienteDynamoDBImpl {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	return &RepositorioClienteDynamoDBImpl{svc: dynamodb.New(sess)}
}

func (r *RepositorioClienteDynamoDBImpl) SalvarCliente(cliente *dominio.Cliente) error {
	av, err := dynamodbattribute.MarshalMap(cliente)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("ClienteAppFastfood"),
	}

	_, err = r.svc.PutItem(input)
	return err
}

func (r *RepositorioClienteDynamoDBImpl) AtualizarCliente(cliente *dominio.Cliente) error {
	av, err := dynamodbattribute.MarshalMap(cliente)
	if err != nil {
		return err
	}

	// Converte o mapa de atributos para um mapa de expressões de atualização
	updateExpression := "SET "
	expressionAttributeValues := make(map[string]*dynamodb.AttributeValue)
	expressionAttributeNames := make(map[string]*string)
	for k, v := range av {

		if k != "ID" {
			if k == "Status" { // Substitui o nome do atributo reservado por um nome alternativo
				expressionAttributeNames["#s"] = aws.String("Status")
				updateExpression += "#s = :" + k + ", "
			} else {
				updateExpression += k + " = :" + k + ", "
			}
			expressionAttributeValues[":"+k] = v
		}
	}
	// Remove a última vírgula e espaço da expressão de atualização
	updateExpression = strings.TrimSuffix(updateExpression, ", ")

	input := &dynamodb.UpdateItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {S: aws.String(cliente.ID)}, // Supondo que "ID" é a chave primária
		},
		TableName:                 aws.String("ClienteAppFastfood"),
		UpdateExpression:          aws.String(updateExpression),
		ExpressionAttributeValues: expressionAttributeValues,
		ExpressionAttributeNames:  expressionAttributeNames, // Adiciona os nomes de atributos alternativos
	}

	log.Printf("Input para UpdateItem: %+v\n", input)
	_, err = r.svc.UpdateItem(input)
	return err
}

func (r *RepositorioClienteDynamoDBImpl) BuscarClientePorID(idCliente string) (*dominio.Cliente, error) {
	input := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {
				S: aws.String(idCliente),
			},
		},
		TableName: aws.String("ClienteAppFastfood"),
	}

	result, err := r.svc.GetItem(input)
	if err != nil {
		return nil, err
	}
	fmt.Printf("result: %v\n", result)
	cliente := &dominio.Cliente{}
	err = dynamodbattribute.UnmarshalMap(result.Item, cliente)
	fmt.Printf("cliente: %v\n", cliente)
	return cliente, err
}
