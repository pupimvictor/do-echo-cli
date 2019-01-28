package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-openapi/strfmt"
	"github.com/pupimvictor/do-echo-cli/client"
	"github.com/pupimvictor/do-echo-cli/models"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	httptransport "github.com/go-openapi/runtime/client"
)

func TestYell(t *testing.T) {
	tests := []struct{
		name string
		msg string
		expectedOutput string
	}{
		{
			name: "test1",
			msg: "msg1",
			expectedOutput: "msg1\n",
		},
		{
			name: "test2",
			msg: "msg2 msg2",
			expectedOutput: "msg2 msg2\n",
		},
		{
			name: "testEmpty",
			msg: "",
			expectedOutput: "\n",
		},
	}

	for _, test := range tests{
		t.Run(test.name, func(t *testing.T) {
			echoServerMock := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")

				resp := &models.Echo{
					Echo: test.msg,
				}
				respbytes, _ := json.Marshal(resp)
				_, err := w.Write(respbytes)
				if err != nil {
					t.Errorf("err on mock write: %v\n", err)
				}
				return
			}))

			msgs := []models.Message{
				{
					Msg: &test.msg,
				},
			}

			stdout := bytes.NewBuffer([]byte{})

			host, _ := url.Parse(echoServerMock.URL)
			hostUrl := host.Host
			transport := httptransport.New(hostUrl, "", []string{"http"})
			echoer := client.New(transport, strfmt.Default)

			cli := echoClient{
				messages: &msgs,
				delay: 0,
				stdOut: stdout,
				echoer: echoer,
			}
			err := cli.yell()
			if err != nil {
				t.Error(err)
			}
            r := cli.stdOut.(*bytes.Buffer).String()
			if r != test.expectedOutput {
				t.Error(fmt.Sprintf("expected %s got %s", test.expectedOutput, r))
			}
		})
	}
}

