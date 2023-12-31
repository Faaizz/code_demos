package types

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	dynamodbtypes "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	"github.com/Faaizz/code_demos/go_simple_chat_app_api/misc"
)

var ddbSvc *dynamodb.Client

func init() {
	logger = misc.Logger()

	epUrl := os.Getenv("DYNAMODB_ENDPOINT_URL")

	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		if epUrl != "" {
			return aws.Endpoint{
				URL: epUrl,
			}, nil
		}

		return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	})

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithEndpointResolverWithOptions(customResolver))
	if err != nil {
		logger.Fatalf("could not initialize AWS client %v", err)
	}

	ddbSvc = dynamodb.NewFromConfig(cfg)
}

// A DynamoDBAdapter provides a layer of abstraction for interaction an underlying AWS DynamoDB database
// It expects a DynamoDB table with a string-valued partition key "connectionId".
type DynamoDBAdapter struct {
	TableName string
}

func (dba *DynamoDBAdapter) SetTableName(tn string) {
	dba.TableName = tn
}

func (dba *DynamoDBAdapter) CheckExists(ctx context.Context) error {
	in := dynamodb.DescribeTableInput{
		TableName: aws.String(dba.TableName),
	}
	_, err := ddbSvc.DescribeTable(context.TODO(), &in)
	if err != nil {
		return err
	}

	return nil
}

// PutConn inserts a connectionId in the underlying DynamoDB table
func (dba *DynamoDBAdapter) PutConn(ctx context.Context, pcIn Connection) error {

	in := dynamodb.PutItemInput{
		TableName: aws.String(dba.TableName),
		Item: map[string]dynamodbtypes.AttributeValue{
			"connectionId": &dynamodbtypes.AttributeValueMemberS{
				Value: pcIn.ConnectionID,
			},
		},
	}

	_, err := ddbSvc.PutItem(
		ctx,
		&in,
	)
	if err != nil {
		return err
	}

	return nil
}

// ConnectionID gets the connection ID associated with the specified username
func (dba *DynamoDBAdapter) ConnectionID(ctx context.Context, un string) (string, error) {
	in := dynamodb.ScanInput{
		TableName:        &dba.TableName,
		FilterExpression: aws.String("username = :val"),
		ExpressionAttributeValues: map[string]dynamodbtypes.AttributeValue{
			":val": &dynamodbtypes.AttributeValueMemberS{
				Value: un,
			},
		},
		ConsistentRead: aws.Bool(true),
	}

	out, err := ddbSvc.Scan(ctx, &in)
	if err != nil {
		logger.Errorln(err)
		return "", err
	}
	if out.Count <= 0 {
		err := errors.New("user not found")
		logger.Errorln(err)
		return "", err
	}

	connIDAtt := out.Items[0]["connectionId"]
	connID, ok := connIDAtt.(*dynamodbtypes.AttributeValueMemberS)
	if !ok {
		err := errors.New("could not obtain connectionId")
		logger.Errorln(err)
		return "", err
	}

	return connID.Value, nil
}

// SetUsername sets the username of an existing connection
func (dba *DynamoDBAdapter) SetUsername(ctx context.Context, pcIn User) error {
	err := dba.CheckUsername(ctx, pcIn.Username)
	if err != nil {
		return err
	}

	in := dynamodb.PutItemInput{
		TableName: aws.String(dba.TableName),
		Item: map[string]dynamodbtypes.AttributeValue{
			"connectionId": &dynamodbtypes.AttributeValueMemberS{
				Value: pcIn.ConnectionID,
			},
			"username": &dynamodbtypes.AttributeValueMemberS{
				Value: pcIn.Username,
			},
		},
	}

	_, err = ddbSvc.PutItem(
		ctx,
		&in,
	)
	if err != nil {
		return err
	}

	return nil
}

// Username gets the username associated with connID
func (dba *DynamoDBAdapter) Username(ctx context.Context, connID string) (string, error) {
	in := dynamodb.GetItemInput{
		TableName: aws.String(dba.TableName),
		Key: map[string]dynamodbtypes.AttributeValue{
			"connectionId": &dynamodbtypes.AttributeValueMemberS{
				Value: connID,
			},
		},
	}

	out, err := ddbSvc.GetItem(
		ctx,
		&in,
	)
	if err != nil {
		logger.Errorln(err)
		return "", err
	}

	unAtt, ok := out.Item["username"].(*dynamodbtypes.AttributeValueMemberS)
	if !ok {
		err := errors.New("could not obtain username")
		logger.Errorln(err)
		return "", err
	}

	return unAtt.Value, nil
}

// CheckUsername checks if username already exists on DynamDB table
func (dba *DynamoDBAdapter) CheckUsername(ctx context.Context, username string) error {
	in := dynamodb.ScanInput{
		TableName:        &dba.TableName,
		FilterExpression: aws.String("username = :val"),
		ExpressionAttributeValues: map[string]dynamodbtypes.AttributeValue{
			":val": &dynamodbtypes.AttributeValueMemberS{
				Value: username,
			},
		},
		ConsistentRead: aws.Bool(true),
	}

	out, err := ddbSvc.Scan(ctx, &in)
	if err != nil {
		return err
	}

	if out.Count <= 0 {
		return nil
	}

	return fmt.Errorf("username '%s' already exists", username)
}

// AvailableUsers lists available users and their connection IDs
// Possible bug: return payload exceeds maximum dataset size limit of 1 MB
func (dba *DynamoDBAdapter) AvailableUsers(ctx context.Context, u User) ([]User, error) {
	in := &dynamodb.ScanInput{
		TableName: &dba.TableName,
	}

	if u.Username != "" {
		in.FilterExpression = aws.String(
			"username <> :val",
		)
		in.ExpressionAttributeValues = map[string]dynamodbtypes.AttributeValue{
			":val": &dynamodbtypes.AttributeValueMemberS{
				Value: u.Username,
			},
		}
	}

	out, err := ddbSvc.Scan(
		ctx,
		in,
	)
	if err != nil {
		return []User{}, err
	}

	au := make([]User, out.Count)
	for idx, item := range out.Items {

		connId := item["connectionId"]
		var connIdStr string
		switch v := connId.(type) {
		case *dynamodbtypes.AttributeValueMemberS:
			connIdStr = v.Value
		default:
			connIdStr = ""
		}

		username := item["username"]
		var usernameStr string
		switch v := username.(type) {
		case *dynamodbtypes.AttributeValueMemberS:
			usernameStr = v.Value
		default:
			usernameStr = ""
		}

		if connIdStr == "" {
			return []User{}, errors.New("could not decode response")
		}

		au[idx] = User{
			ConnectionID: connIdStr,
			Username:     usernameStr,
		}
	}

	return au, nil
}

// Disconnect disconnects current User by deleting the user from DB
func (dba *DynamoDBAdapter) Disconnect(ctx context.Context, u User) error {
	in := &dynamodb.DeleteItemInput{
		TableName: &dba.TableName,
		Key: map[string]dynamodbtypes.AttributeValue{
			"connectionId": &dynamodbtypes.AttributeValueMemberS{
				Value: u.ConnectionID,
			},
		},
	}

	_, err := ddbSvc.DeleteItem(ctx, in)
	if err != nil {
		return err
	}

	return nil
}
