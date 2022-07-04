package aws

import (
	"context"
	"errors"
	"log"

	appConfig "rohandhamapurkar/code-executor/core/config"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

var cognitoClient *cognitoidentityprovider.Client
var cognitoClientId string

func initCognitoClient() {

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(appConfig.AwsCognitoRegion))
	if err != nil {
		log.Fatalf("unable to load AWS Cognito SDK config, %v", err)
	}

	cognitoClient = cognitoidentityprovider.NewFromConfig(cfg)
	cognitoClientId = appConfig.AwsCognitoClientId

}

func SignUpUser(username *string, password *string) error {
	context := context.Background()
	output, err := cognitoClient.SignUp(context, &cognitoidentityprovider.SignUpInput{
		ClientId: &cognitoClientId,

		Username: username,
		Password: password,
	})
	if err != nil {
		return err
	}
	if output.UserConfirmed != false {
		return errors.New("User signUp failed")
	}
	return nil
}
