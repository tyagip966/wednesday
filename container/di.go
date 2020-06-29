package container

import (
	"github.com/jinzhu/gorm"
	"log"
	"wednesday/config"
	"wednesday/constants"
	"wednesday/controller"
	"wednesday/models"
	"wednesday/models/cab/repository/postgres"
	"wednesday/models/cab/usecase"
	"wednesday/models/pg_utils"
	"wednesday/utils"
)

type Container struct {
	Profile       string
	Injected      bool `json:"injected"`
	CabController *controller.CabController
	CabUseCase    models.CabUseCase
	CabRepo       models.CabRepo
	DB            *gorm.DB
	YamlConfig    *config.YamlConfig
}

func (container *Container) TriggerDI() {
	if container.Injected {
		log.Println("Dependency Already Injected !!!")
		return
	}
	log.Println("Started Injecting Dependencies !!!")
	container.GetCabController()
	log.Println("DI over !!!")
}

func (container *Container) GetCabController() *controller.CabController {
	if container.CabController == nil {
		container.CabController = controller.NewCabController(container.GetCabUseCase())
	}
	return container.CabController
}

func (container *Container) GetCabUseCase() models.CabUseCase {
	if container.CabUseCase == nil {
		tx := pg_utils.NewTransaction(container.GetDB())
		container.CabUseCase = usecase.NewUseCaseCab(container.GetCabRepo(),tx)
	}
	return container.CabUseCase
}

func (container *Container) GetCabRepo() models.CabRepo {
	if container.CabRepo == nil {
		container.CabRepo = postgres.NewCabRepository()
	}
	return container.CabRepo
}

func (container *Container) GetDB() *gorm.DB {
	var err error
	if container.DB != nil {
		return container.DB
	}
	container.DB, err = utils.StartUp(&container.GetYamlConfig().DataSource)
	if err != nil {
		log.Fatal(constants.DBConnectionErr, err)
	}
	return container.DB
}

func (container *Container) GetYamlConfig() *config.YamlConfig {
	if container.YamlConfig == nil {
		configYaml := config.GetServiceConfig(container.Profile)
		if configYaml != nil {
			container.YamlConfig = configYaml
		}
	}
	return container.YamlConfig
}
