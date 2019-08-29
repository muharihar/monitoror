package jenkins

import (
	. "github.com/monitoror/monitoror/models/tiles"
	"github.com/monitoror/monitoror/monitorable/jenkins/models"
	"github.com/monitoror/monitoror/pkg/monitoror/builder"
)

const (
	JenkinsBuildTileType       TileType = "JENKINS-BUILD"
	JenkinsMultiBranchTileType TileType = "JENKINS-MULTIBRANCH"
)

// Usecase represent the jenkins's usecases
type (
	Usecase interface {
		Build(params *models.BuildParams) (*BuildTile, error)
		ListDynamicTile(params interface{}) ([]builder.Result, error)
	}
)