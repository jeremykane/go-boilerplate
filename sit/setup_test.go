package sit

import (
	"fmt"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/google/uuid"
	init_module "github.com/jeremykane/go-boilerplate/internal/app/init-module"
	"github.com/jeremykane/go-boilerplate/internal/config"
	"github.com/jeremykane/go-boilerplate/pkg/database"
	"github.com/jeremykane/go-boilerplate/pkg/logger"
	"github.com/jeremykane/go-boilerplate/pkg/logger/logrus"
	"github.com/stretchr/testify/suite"
	"github.com/walkerus/go-wiremock"
)

func TestSit(t *testing.T) {
	suite.Run(t, new(SitTestSuite))
}

type SitTestSuite struct {
	suite.Suite
	wiremockClient *wiremock.Client
	db             map[string]*database.Replication
	config         *config.Config
}

func (suite *SitTestSuite) SetupSuite() {
	_, b, _, _ := runtime.Caller(0)
	configPath := filepath.Join(filepath.Dir(b), "../", "config")
	fmt.Println(configPath)
	cfg, err := config.Load(configPath, "sit_config")
	if err != nil {
		panic("Failed to read config: " + err.Error())
	}

	suite.db = init_module.InitializeDB(cfg.Database)
	host := config.ConfigMap.GetString("WIREMOCK_HOST")
	suite.wiremockClient = wiremock.NewClient(host)
	suite.config = cfg

	paramLogs := logrus.LogrusParam{
		Level: cfg.Server.ApiLogLevel,
	}
	logs := logrus.NewLogrus(paramLogs)
	logger.NewLogger(logs)
}

func (suite *SitTestSuite) SetupTest() {
	suite.db[config.DatabaseGo].Master.Exec(`truncate table facility_types cascade`)
}

func (suite *SitTestSuite) AfterTest() {
	suite.wiremockClient.Reset()
	suite.wiremockClient.ResetAllScenarios()
}

func setupFixture(suite *SitTestSuite, randomUuid string) {
	var (
		facilityTypeId1         int64 = 1000
		facilityTypeName1             = "SIT Facility Type 1"
		facilityTypeBahasaName1       = "SIT Tipe Fasiliti 1"
		facilityTypeId2         int64 = 1001
		facilityTypeName2             = "SIT Facility Type 2"
		facilityTypeBahasaName2       = "SIT Tipe Fasiliti 2"
		facilityTypeId3         int64 = 1002
		facilityTypeName3             = "SIT Facility Type 3"
		facilityTypeBahasaName3       = "SIT Tipe Fasiliti 3"
	)

	err := suite.db[config.DatabaseGo].Master.Exec("INSERT INTO facility_types (id, uuid, name, bahasa_name) VALUES($1, $2, $3, $4) RETURNING id", facilityTypeId1, uuid.NewString(), facilityTypeName1, facilityTypeBahasaName1)
	suite.Assert().Equal(int64(1), err.RowsAffected)

	err = suite.db[config.DatabaseGo].Master.Exec("INSERT INTO facility_types (id, uuid, name, bahasa_name) VALUES($1, $2, $3, $4) RETURNING id", facilityTypeId2, uuid.NewString(), facilityTypeName2, facilityTypeBahasaName2)
	suite.Assert().Equal(int64(1), err.RowsAffected)

	err = suite.db[config.DatabaseGo].Master.Exec("INSERT INTO facility_types (id, uuid, name, bahasa_name) VALUES($1, $2, $3, $4) RETURNING id", facilityTypeId3, uuid.NewString(), facilityTypeName3, facilityTypeBahasaName3)
	suite.Assert().Equal(int64(1), err.RowsAffected)

}
