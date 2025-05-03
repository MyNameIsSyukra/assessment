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
	assesmentRepository := repository.NewAssessmentRepository(db)
	questionService := service.NewQuestionService(questionRepository,assesmentRepository)

	do.Provide(injector, func (i *do.Injector) (controller.QuestionController, error){
		return controller.NewQuestionController(questionService),nil
	})
}

func ProvideAnswerDependencies(injector *do.Injector){
	db := do.MustInvokeNamed[*gorm.DB](injector, constants.DB)


	answerRepository := repository.NewAnswerRepository(db)
	submissionRepository := repository.NewSubmissionRepository(db)
	assesmentRepository := repository.NewAssessmentRepository(db)
	answerService := service.NewAnswerService(answerRepository, submissionRepository, assesmentRepository)
	do.Provide(injector, func (i *do.Injector) (controller.AnswerController, error){
		return controller.NewAnswerController(answerService),nil
	})
}

func ProvideChoiceDependencies(injector *do.Injector){
	db := do.MustInvokeNamed[*gorm.DB](injector, constants.DB)

	choiceRepository := repository.NewChoiceRepository(db)
	choiceService := service.NewChoiceService(choiceRepository)

	do.Provide(injector, func (i *do.Injector) (controller.ChoiceController, error){
		return controller.NewChoiceController(choiceService),nil
	})
}

func ProvideSubmissionDependencies(injector *do.Injector){
	db := do.MustInvokeNamed[*gorm.DB](injector, constants.DB)

	submissionRepository := repository.NewSubmissionRepository(db)
	questionRepository := repository.NewQuestionRepository(db)
	submissionService := service.NewSubmissionService(submissionRepository, questionRepository)

	do.Provide(injector, func (i *do.Injector) (controller.SubmissionController, error){
		return controller.NewSubmissionController(submissionService),nil
	})
}

func RegisterDependencies(injector *do.Injector) {
	InitDatabase(injector)
	ProvideEvaluationDependencies(injector)
	ProvideQuestionDependencies(injector)
	ProvideAnswerDependencies(injector)
	ProvideChoiceDependencies(injector)
	ProvideSubmissionDependencies(injector)
}
