package tests

import (
  "fmt"
  "net/http"
  "net"
  "bytes"
  "errors"
  "time"
)

func makeRequest(host net.IP) (*http.Request, error) {
  url := fmt.Sprintf("http://%s/test", host.String())
  return http.NewRequest(http.MethodGet, url, bytes.NewBuffer([]byte{}))
}

func HttpTest(host net.IP) error {
  req, err := makeRequest(host)
  if err != nil { return err }

  client := &http.Client{Timeout: 3 * time.Second}

  resp, err :=client.Do(req)
  if err != nil { return err }

  if resp.StatusCode == 200 {
    /* test succesful */
    return nil
  } else {
    return errors.New("http test failed")
  }
}

