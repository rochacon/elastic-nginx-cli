package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"io"
	"net/http"
	"os"
)

func main(){
	payload := struct{
		TopicArn string
		Message  string
	}{
		"arn:aws:sns:autoscaling",
		fmt.Sprintf(`{"AutoScalingGroupARN":"arn:aws:autoscaling:backends-1",`+
					`"Event":"autoscaling:EC2_INSTANCE_%s","EC2InstanceId":"%s"}`,
					os.Args[1], os.Args[2]),
	}

	buf, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}

	r, err := http.Post("http://localhost:5000", "application/json", bytes.NewReader(buf))
	if err != nil {
		log.Fatal(err)
	}

	io.Copy(os.Stdout, r.Body)
}
