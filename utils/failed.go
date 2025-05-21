package utils

import "errors"

const (
	FailedGetDataFromBody		 = "failed to get data from body"
	FailedCreateAssessment         = "failed to create assessment"
	FailedGetAllAssessments        = "failed to get all assessments"
	FailedGetAssessmentByID        = "failed to get assessment by ID"
	FailedGetAllAssesmentByClassID = "failed to get all assessments by class ID"
	FailedUpdateAssessment         = "failed to update assessment"
	FailedDeleteAssessment         = "failed to delete assessment"
	FailedCreateSubmission         = "failed to create submission"
	FailedGetSubmissionByID        = "failed to get submission by ID"
	FailedGetAllSubmissions        = "failed to get all submissions"
	FailedGetSubmissionByStudentID = "failed to get submission by student ID"
	FailedUpdateSubmission         = "failed to update submission"
	FailedDeleteSubmission         = "failed to delete submission"

	FailedCreateQuestion  = "failed to create question"
	FailedGetAllQuestions = "failed to get all questions"
	FailedGetQuestionByID = "failed to get question by ID"
	FailedUpdateQuestion  = "failed to update question"
	FailedDeleteQuestion  = "failed to delete question"

	FailedCreateChoice  = "failed to create choice"
	FailedGetAllChoices = "failed to get all choices"
	FailedGetChoiceByID = "failed to get choice by ID"
	FailedUpdateChoice  = "failed to update choice"
	FailedDeleteChoice  = "failed to delete choice"

	FailedCreateAnswer  = "failed to create answer"
	FailedGetAllAnswers = "failed to get all answers"
	FailedGetAnswerByID = "failed to get answer by ID"
	FailedUpdateAnswer  = "failed to update answer"
	FailedDeleteAnswer  = "failed to delete answer"

	FailedGetAnswerByStudentID    = "failed to get answer by student ID"
	FailedGetAnswerBySubmissionID = "failed to get answer by submission ID"
	FailedGetAnswerByQuestionID   = "failed to get answer by question ID"

	FailedGetAnswerByChoiceID               = "failed to get answer by choice ID"
	FailedGetAnswerByQuestionIDAndStudentID = "failed to get answer by question ID and student ID"

	FailedGetAnswerBySubmissionIDAndStudentID  = "failed to get answer by submission ID and student ID"
	FailedGetAnswerByQuestionIDAndSubmissionID = "failed to get answer by question ID and submission ID"

	FailedGetAnswerByChoiceIDAndSubmissionID = "failed to get answer by choice ID and submission ID"
	FailedGetAnswerByQuestionIDAndChoiceID   = "failed to get answer by question ID and choice ID"

	FailedGetAnswerByQuestionIDAndChoiceIDAndStudentID    = "failed to get answer by question ID, choice ID and student ID"
	FailedGetAnswerByQuestionIDAndChoiceIDAndSubmissionID = "failed to get answer by question ID, choice ID and submission ID"
)

var (
	ErrCreateAssesment          = errors.New(FailedCreateAssessment)
	ErrGetAllAssesments         = errors.New(FailedGetAllAssessments)
	ErrGetAssesmentByID         = errors.New(FailedGetAssessmentByID)
	ErrGetAllAssesmentByClassID = errors.New(FailedGetAllAssesmentByClassID)
	ErrUpdateAssesment          = errors.New(FailedUpdateAssessment)
	ErrDeleteAssesment          = errors.New(FailedDeleteAssessment)

	ErrCreateSubmission         = errors.New(FailedCreateSubmission)
	ErrGetSubmissionByID        = errors.New(FailedGetSubmissionByID)
	ErrGetAllSubmissions        = errors.New(FailedGetAllSubmissions)
	ErrGetSubmissionByStudentID = errors.New(FailedGetSubmissionByStudentID)
	ErrUpdateSubmission         = errors.New(FailedUpdateSubmission)
	ErrDeleteSubmission         = errors.New(FailedDeleteSubmission)

	ErrCreateQuestion  = errors.New(FailedCreateQuestion)
	ErrGetAllQuestions = errors.New(FailedGetAllQuestions)
	ErrGetQuestionByID = errors.New(FailedGetQuestionByID)
	ErrUpdateQuestion  = errors.New(FailedUpdateQuestion)
	ErrDeleteQuestion  = errors.New(FailedDeleteQuestion)

	ErrCreateChoice  = errors.New(FailedCreateChoice)
	ErrGetAllChoices = errors.New(FailedGetAllChoices)
	ErrGetChoiceByID = errors.New(FailedGetChoiceByID)
	ErrUpdateChoice  = errors.New(FailedUpdateChoice)
	ErrDeleteChoice  = errors.New(FailedDeleteChoice)

	ErrCreateAnswer  = errors.New(FailedCreateAnswer)
	ErrGetAllAnswers = errors.New(FailedGetAllAnswers)
	ErrGetAnswerByID = errors.New(FailedGetAnswerByID)
	ErrUpdateAnswer  = errors.New(FailedUpdateAnswer)
	ErrDeleteAnswer  = errors.New(FailedDeleteAnswer)

	ErrGetAnswerByStudentID                 = errors.New(FailedGetAnswerByStudentID)
	ErrGetAnswerBySubmissionID              = errors.New(FailedGetAnswerBySubmissionID)
	ErrGetAnswerByQuestionID                = errors.New(FailedGetAnswerByQuestionID)
	ErrGetAnswerByChoiceID                  = errors.New(FailedGetAnswerByChoiceID)
	ErrGetAnswerByQuestionIDAndStudentID    = errors.New(FailedGetAnswerByQuestionIDAndStudentID)
	ErrGetAnswerBySubmissionIDAndStudentID  = errors.New(FailedGetAnswerBySubmissionIDAndStudentID)
	ErrGetAnswerByQuestionIDAndSubmissionID = errors.New(FailedGetAnswerByQuestionIDAndSubmissionID)

	ErrGetAnswerByChoiceIDAndSubmissionID           = errors.New(FailedGetAnswerByChoiceIDAndSubmissionID)
	ErrGetAnswerByQuestionIDAndChoiceID             = errors.New(FailedGetAnswerByQuestionIDAndChoiceID)
	ErrGetAnswerByQuestionIDAndChoiceIDAndStudentID = errors.New(FailedGetAnswerByQuestionIDAndChoiceIDAndStudentID)

	ErrGetAnswerByQuestionIDAndChoiceIDAndSubmissionID = errors.New(FailedGetAnswerByQuestionIDAndChoiceIDAndSubmissionID)
)