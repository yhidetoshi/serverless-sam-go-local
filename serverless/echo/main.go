package main

import (
	"context"
	"log"
	"yhidetoshi/serverless-sam-go-local/api/healthcheck"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	echoadapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var echoLambda *echoadapter.EchoLambda // APIGW(REST)
//var echoLambda *echoadapter.EchoLambdaV2 // APIGW(HTTP)

func init() {
	log.Printf("echo cold start")

	e := echo.New()
	e.Use(middleware.Recover())
	//e.Use(auth.BasicAuth())
	e.GET("/api/healthcheck", healthcheck.Healthcheck)

	echoLambda = echoadapter.New(e) // APIGW(REST)
	//echoLambda = echoadapter.NewV2(e) // APIGW(HTTP)
}

func main() {
	lambda.Start(Handler)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) { // APIGW(REST)
	//func Handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) { // APIGW(HTTP)
	return echoLambda.ProxyWithContext(ctx, req)
}
