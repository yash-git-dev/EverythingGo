package Controllers

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

func SendMail() {
	region := "sotheast-asia-2"
	AWSAccessKeyID := "somerandomstring"
	AWSSecretAccessKey := "something"
	creds := credentials.NewStaticCredentials(AWSAccessKeyID, AWSSecretAccessKey, "")
	sess, _ := session.NewSession(&aws.Config{Region: &region, Credentials: creds})
	ses1 := ses.New(sess)

	_, err := ses1.SendEmail(&ses.SendEmailInput{Destination: &ses.Destination{
		ToAddresses: []*string{aws.String("yash@gmail.com")},
	},
		Message: &ses.Message{
			Body: &ses.Body{
				Text: &ses.Content{
					Data: aws.String("hello world"),
				},
			},
			Subject: &ses.Content{
				Data: aws.String("This is a test email."),
			},
		},
		Source: aws.String("yash@gmail.com"),
	})

	if err != nil {
		fmt.Println(err)
	}
}
