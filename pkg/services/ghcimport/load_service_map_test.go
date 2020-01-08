package ghcimport

import (
	"testing"

	"github.com/gofrs/uuid"

	"github.com/transcom/mymove/pkg/models"
)

func (suite *GHCRateEngineImportSuite) Test_loadServiceMap() {
	gre := &GHCRateEngineImporter{
		Logger:       suite.logger,
		ContractCode: testContractCode,
	}

	suite.T().Run("load success", func(t *testing.T) {
		err := gre.loadServiceMap(suite.DB())
		suite.NoError(err)

		suite.NotNil(gre.serviceToIDMap)

		count, err := suite.DB().Count(&models.ReService{})
		suite.NoError(err)
		suite.Greater(count, 0)
		suite.Equal(count, len(gre.serviceToIDMap))

		// Spot-check a service code
		testServiceCode := "DOASIT"
		if suite.Contains(gre.serviceToIDMap, testServiceCode) {
			suite.NotEqual(uuid.Nil, gre.serviceToIDMap[testServiceCode])
		}
	})
}
