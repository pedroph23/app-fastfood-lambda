package repositorio

import (
	"fmt"

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

func (r *RepositorioClienteDynamoDBImpl) SalvarOuAtualizarCliente(cliente *dominio.Cliente) error {
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
