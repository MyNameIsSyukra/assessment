package provider

import (
	config "assesment/config"
	constants "assesment/constants"
	controller "assesment/controller"
	repository "assesment/repository"
	service "assesment/service"

	"github.com/samber/do"
	"gorm.io/gorm"
)

func InitDatabase(injector *do.Injector) {
	do.ProvideNamed(injector, constants.DB, func (i *do.Injector) (*gorm.DB, error) {
		return config.SetUpDatabaseConnection(), nil
	})
}

func ProvideEvaluationDependencies(injector *do.Injector){
	db := do.MustInvokeNamed[*gorm.DB](injector, constants.DB)

	assesmentRepository := repository.NewAssessmentRepository(db)
	assesmentService := service.NewAssessmentService(assesmentRepository)

	do.Provide(injector, func (i *do.Injector) (controller.AssessmentController, error){
		return controller.NewAssessmentController(assesmentService),nil
	})
}

func ProvideQuestionDependencies(injector *do.Injector){
	db := do.MustInvokeNamed[*gorm.DB](injector, constants.DB)

	questionRepository := repository.NewQuestionRepository(db)
	questionService := service.NewQuestionService(questionRepository)

	do.Provide(injector, func (i *do.Injector) (controller.QuestionController, error){
		return controller.NewQuestionController(questionService),nil
	})
}


func RegisterDependencies(injector *do.Injector) {
	InitDatabase(injector)
	ProvideEvaluationDependencies(injector)
	ProvideQuestionDependencies(injector)
}
