package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/shortedapp/shortedfunctions/internal/handlerhelper/topshortsingestor"
	"github.com/shortedapp/shortedfunctions/pkg/awsutils"
	log "github.com/shortedapp/shortedfunctions/pkg/loggingutil"
)

//Handler - the main function handler, triggered by a SNS event
func Handler(request events.SNSEvent) {
	//Generate Clients
	clients := awsutils.GenerateAWSClients("dynamoDB", "s3")

	//Create topshortslist struct
	t := topshortsingestor.Topshortsingestor{Clients: clients}

	//Run the topshorts fetch routine
	t.IngestTopShorted("testTopShorts")
}

func main() {
	log.SetAppName("ShortedApp")
	lambda.Start(Handler)
}
