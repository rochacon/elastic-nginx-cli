package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"
)

func main() {
	var topicArn string
	var autoScalingGroupARN string
	var event string
	var host string

	flag.StringVar(&topicArn, "topic-arn", "", "Topic ARN")
	flag.StringVar(&autoScalingGroupARN, "asg-arn", "", "Auto Scaling Group ARN")
	flag.StringVar(&event, "event", "launch", "Auto Scaling Event (launch/terminate)")
	flag.StringVar(&host, "host", "127.0.0.1:5000", "Elastic NGINX host (may include port, e.g. 127.0.0.1:5000)")
	flag.Parse()

	if topicArn == "" || autoScalingGroupARN == "" {
		flag.Usage()
		return
	}

	event = strings.ToUpper(event)

	for _, instanceId := range flag.Args() {
		payload := struct {
			TopicArn string
			Message  string
		}{
			topicArn,
			fmt.Sprintf(`{"AutoScalingGroupARN":"%s","Event":"autoscaling:EC2_INSTANCE_%s","EC2InstanceId":"%s"}`, autoScalingGroupARN, event, instanceId),
		}

		buf, err := json.Marshal(payload)
		if err != nil {
			log.Fatal(err)
		}

		r, err := http.Post("http://"+host, "application/json", bytes.NewReader(buf))
		if err != nil {
			log.Fatal(err)
		}

		log.Println("--- InstanceID:", instanceId)
		out, _ := httputil.DumpResponse(r, true)
		fmt.Println(string(out))
		fmt.Println("\n")
	}
}
