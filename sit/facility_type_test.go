package sit

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/jeremykane/go-boilerplate/internal/app"
)

func (suite *SitTestSuite) TestGetAllFacilityType() {
	setupFixture(suite, uuid.New().String())

	httpServer := http.Server{}
	go app.RunServer(&httpServer, suite.config)
	defer func() {
		httpServer.Shutdown(context.Background())
	}()

	fmt.Println("Waiting for server to start...")
	time.Sleep(1 * time.Second)

	// Test
	req, err := http.NewRequest("GET", "http://localhost:9000/facility-types", nil)
	suite.Assert().Nil(err)

	client := &http.Client{}

	resp, err := client.Do(req)
	suite.Assert().Nil(err)
	suite.Assert().Equal(http.StatusOK, resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	suite.Assert().Nil(err)

	respBody := make(map[string]json.RawMessage)
	err = json.Unmarshal(body, &respBody)
	suite.Assert().Nil(err)
	suite.Assert().NotEmpty(respBody["data"])
}
