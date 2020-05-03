package main

import (
	"context"
	"encoding/json"
	"github.com/Development/svg-responder/pkg/svgengine"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (Response, error) {
	var req svgengine.SvgRequest
	var err = json.Unmarshal([]byte(request.Body), &req)
	if err != nil {
		return Response{StatusCode: 500}, err
	}
	file, err := svgengine.Render(req)
	if err != nil {
		return Response{StatusCode: 500}, err
	}
	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            file.String(),
		Headers: map[string]string{
			"Content-Type":        "image/svg+xml",
			"Content-Disposition": "attachment; filename=responsive-svg.svg",
		},
	}
	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
