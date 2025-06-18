package migration

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	database "assesment/config"
	entities "assesment/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func Seeder()error{
	db := database.SetUpDatabaseConnection()
	db.AutoMigrate(&entities.Answer{}, &entities.Assessment{}, &entities.Choice{}, &entities.Question{}, &entities.Submission{})

		// Seed Assessment service
	SeedAssessmentData(db)

	// Print summary
	PrintAssessmentSummary(db)

	fmt.Println("\n========== ASSESSMENT SEEDING COMPLETED ==========")
	fmt.Println("Assessment data seeded successfully!")
	fmt.Println("Data includes:")
	fmt.Println("- Multiple assessments per class (Web Programming, Database, Algorithm)")
	fmt.Println("- Questions with multiple choice answers")
	fmt.Println("- Student submissions with realistic completion rates")
	fmt.Println("- Calculated scores based on correct answers")
	fmt.Println("- Various submission statuses (todo, in_progress, submitted)")
	return nil
}

var (
	// Class IDs (from Class Control service)
	ClassWebProgID   = uuid.MustParse("550e8400-e29b-41d4-a716-446655440001")
	ClassDatabaseID  = uuid.MustParse("550e8400-e29b-41d4-a716-446655440002")
	ClassAlgorithmID = uuid.MustParse("550e8400-e29b-41d4-a716-446655440003")

	// Student IDs (from Class Control service - Member.User_userID)
	StudentAliceID    = uuid.MustParse("550e8400-e29b-41d4-a716-446655440201")
	StudentBobID      = uuid.MustParse("550e8400-e29b-41d4-a716-446655440202")
	StudentCharlieID  = uuid.MustParse("550e8400-e29b-41d4-a716-446655440203")
	StudentDianaID    = uuid.MustParse("550e8400-e29b-41d4-a716-446655440204")
	StudentEdwardID   = uuid.MustParse("550e8400-e29b-41d4-a716-446655440205")
	StudentFionaID    = uuid.MustParse("550e8400-e29b-41d4-a716-446655440206")
	StudentGeorgeID   = uuid.MustParse("550e8400-e29b-41d4-a716-446655440207")
	StudentHannahID   = uuid.MustParse("550e8400-e29b-41d4-a716-446655440208")
	StudentIvanID     = uuid.MustParse("550e8400-e29b-41d4-a716-446655440209")
	StudentJuliaID    = uuid.MustParse("550e8400-e29b-41d4-a716-446655440210")
	StudentKevinID    = uuid.MustParse("550e8400-e29b-41d4-a716-446655440211")
	StudentLindaID    = uuid.MustParse("550e8400-e29b-41d4-a716-446655440212")
	StudentMichaelID  = uuid.MustParse("550e8400-e29b-41d4-a716-446655440213")
	StudentNancyID    = uuid.MustParse("550e8400-e29b-41d4-a716-446655440214")
	StudentOscarID    = uuid.MustParse("550e8400-e29b-41d4-a716-446655440215")
)

// ========== STUDENT DATA MAPPING ==========
type StudentClassMapping struct {
	UserID   uuid.UUID
	Username string
	ClassID  uuid.UUID
}

func GetStudentClassMappings() []StudentClassMapping {
	return []StudentClassMapping{
		// Web Programming Students
		{StudentAliceID, "Alice Johnson", ClassWebProgID},
		{StudentBobID, "Bob Smith", ClassWebProgID},
		{StudentCharlieID, "Charlie Brown", ClassWebProgID},
		{StudentDianaID, "Diana Prince", ClassWebProgID},
		{StudentEdwardID, "Edward Norton", ClassWebProgID},

		// Database Students
		{StudentFionaID, "Fiona Green", ClassDatabaseID},
		{StudentGeorgeID, "George Wilson", ClassDatabaseID},
		{StudentHannahID, "Hannah Davis", ClassDatabaseID},
		{StudentIvanID, "Ivan Petrov", ClassDatabaseID},
		{StudentJuliaID, "Julia Roberts", ClassDatabaseID},

		// Algorithm Students
		{StudentKevinID, "Kevin Hart", ClassAlgorithmID},
		{StudentLindaID, "Linda Carter", ClassAlgorithmID},
		{StudentMichaelID, "Michael Jordan", ClassAlgorithmID},
		{StudentNancyID, "Nancy Drew", ClassAlgorithmID},
		{StudentOscarID, "Oscar Wilde", ClassAlgorithmID},
	}
}

type AssessmentTemplate struct {
	Name        string
	Description string
	ClassID     uuid.UUID
	Duration    int // in minutes
	Questions   []QuestionTemplate
}

type QuestionTemplate struct {
	Text    string
	Choices []ChoiceTemplate
}

type ChoiceTemplate struct {
	Text      string
	IsCorrect bool
}

func GetAssessmentTemplates() []AssessmentTemplate {
	return []AssessmentTemplate{
		// English Grammar Assessments
		{
			Name:        "Quiz Basic Grammar",
			Description: "Quiz tentang tata bahasa dasar bahasa Inggris",
			ClassID:     ClassWebProgID,
			Duration:    3600,
			Questions: []QuestionTemplate{
				{
					Text: "Which sentence uses the correct present tense?",
					Choices: []ChoiceTemplate{
						{Text: "She go to school every day", IsCorrect: false},
						{Text: "She goes to school every day", IsCorrect: true},
						{Text: "She going to school every day", IsCorrect: false},
						{Text: "She went to school every day", IsCorrect: false},
					},
				},
				{
					Text: "What is the plural form of 'child'?",
					Choices: []ChoiceTemplate{
						{Text: "childs", IsCorrect: false},
						{Text: "children", IsCorrect: true},
						{Text: "childes", IsCorrect: false},
						{Text: "child", IsCorrect: false},
					},
				},
				{
					Text: "Choose the correct article: '___ apple is red'",
					Choices: []ChoiceTemplate{
						{Text: "A", IsCorrect: false},
						{Text: "An", IsCorrect: true},
						{Text: "The", IsCorrect: false},
						{Text: "No article needed", IsCorrect: false},
					},
				},
			},
		},
		{
			Name:        "Quiz Tenses & Verb Forms",
			Description: "Quiz tentang berbagai bentuk waktu dan kata kerja dalam bahasa Inggris",
			ClassID:     ClassWebProgID,
			Duration:    3600,
			Questions: []QuestionTemplate{
				{
					Text: "Which sentence is in the past continuous tense?",
					Choices: []ChoiceTemplate{
						{Text: "I was reading a book", IsCorrect: true},
						{Text: "I read a book", IsCorrect: false},
						{Text: "I have read a book", IsCorrect: false},
						{Text: "I will read a book", IsCorrect: false},
					},
				},
				{
					Text: "What is the past participle of 'eat'?",
					Choices: []ChoiceTemplate{
						{Text: "ate", IsCorrect: false},
						{Text: "eating", IsCorrect: false},
						{Text: "eaten", IsCorrect: true},
						{Text: "eats", IsCorrect: false},
					},
				},
			},
		},

		// English Conversation Assessments
		{
			Name:        "Quiz Daily Conversation",
			Description: "Quiz tentang percakapan sehari-hari dalam bahasa Inggris",
			ClassID:     ClassDatabaseID,
			Duration:    3600,
			Questions: []QuestionTemplate{
				{
					Text: "How do you greet someone in the morning?",
					Choices: []ChoiceTemplate{
						{Text: "Good night", IsCorrect: false},
						{Text: "Good morning", IsCorrect: true},
						{Text: "Good afternoon", IsCorrect: false},
						{Text: "Good evening", IsCorrect: false},
					},
				},
				{
					Text: "What is the polite way to ask for help?",
					Choices: []ChoiceTemplate{
						{Text: "Help me!", IsCorrect: false},
						{Text: "Give me help!", IsCorrect: false},
						{Text: "Could you please help me?", IsCorrect: true},
						{Text: "You must help me!", IsCorrect: false},
					},
				},
				{
					Text: "How do you express disagreement politely?",
					Choices: []ChoiceTemplate{
						{Text: "You are wrong", IsCorrect: false},
						{Text: "I'm afraid I disagree", IsCorrect: true},
						{Text: "That's stupid", IsCorrect: false},
						{Text: "No way", IsCorrect: false},
					},
				},
			},
		},
		{
			Name:        "Quiz Expressions & Idioms",
			Description: "Quiz tentang ungkapan dan idiom dalam bahasa Inggris",
			ClassID:     ClassDatabaseID,
			Duration:    3600,
			Questions: []QuestionTemplate{
				{
					Text: "What does 'break a leg' mean?",
					Choices: []ChoiceTemplate{
						{Text: "To hurt yourself", IsCorrect: false},
						{Text: "Good luck", IsCorrect: true},
						{Text: "To run fast", IsCorrect: false},
						{Text: "To dance", IsCorrect: false},
					},
				},
				{
					Text: "What does 'it's raining cats and dogs' mean?",
					Choices: []ChoiceTemplate{
						{Text: "Animals are falling", IsCorrect: false},
						{Text: "It's raining heavily", IsCorrect: true},
						{Text: "It's sunny", IsCorrect: false},
						{Text: "There are pets outside", IsCorrect: false},
					},
				},
			},
		},

		// English Literature Assessments
		{
			Name:        "Quiz Classic Literature",
			Description: "Quiz tentang karya sastra klasik bahasa Inggris",
			ClassID:     ClassAlgorithmID,
			Duration:    3600,
			Questions: []QuestionTemplate{
				{
					Text: "Who wrote 'Romeo and Juliet'?",
					Choices: []ChoiceTemplate{
						{Text: "Charles Dickens", IsCorrect: false},
						{Text: "William Shakespeare", IsCorrect: true},
						{Text: "Jane Austen", IsCorrect: false},
						{Text: "Mark Twain", IsCorrect: false},
					},
				},
				{
					Text: "What is the main theme of 'Pride and Prejudice'?",
					Choices: []ChoiceTemplate{
						{Text: "War and peace", IsCorrect: false},
						{Text: "Love and social class", IsCorrect: true},
						{Text: "Adventure and travel", IsCorrect: false},
						{Text: "Science and technology", IsCorrect: false},
					},
				},
				{
					Text: "In which century did Shakespeare live?",
					Choices: []ChoiceTemplate{
						{Text: "15th century", IsCorrect: false},
						{Text: "16th-17th century", IsCorrect: true},
						{Text: "18th century", IsCorrect: false},
						{Text: "19th century", IsCorrect: false},
					},
				},
			},
		},
		{
			Name:        "Quiz Poetry & Literary Devices",
			Description: "Quiz tentang puisi dan perangkat sastra dalam bahasa Inggris",
			ClassID:     ClassAlgorithmID,
			Duration:    3600,
			Questions: []QuestionTemplate{
				{
					Text: "What is a metaphor?",
					Choices: []ChoiceTemplate{
						{Text: "A direct comparison using 'like' or 'as'", IsCorrect: false},
						{Text: "A direct comparison without using 'like' or 'as'", IsCorrect: true},
						{Text: "A sound device", IsCorrect: false},
						{Text: "A rhyme scheme", IsCorrect: false},
					},
				},
				{
					Text: "What is alliteration?",
					Choices: []ChoiceTemplate{
						{Text: "Repetition of ending sounds", IsCorrect: false},
						{Text: "Repetition of beginning consonant sounds", IsCorrect: true},
						{Text: "Repetition of vowel sounds", IsCorrect: false},
						{Text: "Repetition of words", IsCorrect: false},
					},
				},
			},
		},
	}
}

// ========== SEEDER FUNCTIONS ==========
func SeedAssessmentData(db *gorm.DB) {
	fmt.Println("Seeding Assessment data...")

	templates := GetAssessmentTemplates()
	studentMappings := GetStudentClassMappings()

	// Seed Assessments with Questions and Choices
	var createdAssessments []entities.Assessment
	for _, template := range templates {
		// Create Assessment
		assessment := entities.Assessment{
			Name:        template.Name,
			Description: template.Description,
			StartTime:   time.Now().AddDate(0, 0, -7), // Started 1 week ago
			EndTime:     time.Now().AddDate(0, 0, 14), // Ends in 2 weeks
			Duration:    template.Duration,
			ClassID:     template.ClassID,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}
		
		if err := db.Create(&assessment).Error; err != nil {
			log.Printf("Error creating assessment %s: %v", template.Name, err)
			continue
		}
		createdAssessments = append(createdAssessments, assessment)

		fmt.Printf("Created assessment: %s (ID: %s)\n", assessment.Name, assessment.ID)

		// Create Questions and Choices for this Assessment
		for _, questionTemplate := range template.Questions {
			question := entities.Question{
				QuestionText: questionTemplate.Text,
				AssessmentID: assessment.ID,
				CreatedAt:    time.Now(),
				UpdatedAt:    time.Now(),
			}
			
			if err := db.Create(&question).Error; err != nil {
				log.Printf("Error creating question: %v", err)
				continue
			}

			// Create Choices for this Question
			for _, choiceTemplate := range questionTemplate.Choices {
				choice := entities.Choice{
					ChoiceText: choiceTemplate.Text,
					QuestionID: question.ID,
					IsCorrect:  choiceTemplate.IsCorrect,
					CreatedAt:  time.Now(),
					UpdatedAt:  time.Now(),
				}
				
				if err := db.Create(&choice).Error; err != nil {
					log.Printf("Error creating choice: %v", err)
				}
			}
		}
	}

	fmt.Printf("Created %d assessments with questions and choices\n", len(createdAssessments))

	// Seed Submissions and Answers
	SeedSubmissionsAndAnswers(db, createdAssessments, studentMappings)

	fmt.Println("Assessment data seeded successfully!")
}

func SeedSubmissionsAndAnswers(db *gorm.DB, assessments []entities.Assessment, studentMappings []StudentClassMapping) {
	fmt.Println("Seeding submissions and answers...")

	rand.Seed(time.Now().UnixNano())

	for _, assessment := range assessments {
		// Get students for this class
		var studentsInClass []StudentClassMapping
		for _, student := range studentMappings {
			if student.ClassID == assessment.ClassID {
				studentsInClass = append(studentsInClass, student)
			}
		}

		// Create submissions for 60-80% of students in class
		numSubmissions := len(studentsInClass)*3/4 + rand.Intn(len(studentsInClass)/4+1)
		if numSubmissions > len(studentsInClass) {
			numSubmissions = len(studentsInClass)
		}

		// Shuffle students to get random selection
		shuffledStudents := make([]StudentClassMapping, len(studentsInClass))
		copy(shuffledStudents, studentsInClass)
		for i := range shuffledStudents {
			j := rand.Intn(i + 1)
			shuffledStudents[i], shuffledStudents[j] = shuffledStudents[j], shuffledStudents[i]
		}

		for i := 0; i < numSubmissions; i++ {
			student := shuffledStudents[i]
			
			// Determine submission status randomly
			statusRand := rand.Float32()
			var status entities.ExamStatus
			var submittedAt time.Time
			var createdAt time.Time
			var endedTime time.Time

			// Set CreatedAt untuk submission (kapan submission dimulai)
			createdAt = time.Now().Add(-time.Duration(rand.Intn(7*24)) * time.Hour) // Random time in last week

			// Calculate EndedTime berdasarkan CreatedAt + Duration assessment
			expectedEndTime := createdAt.Add(time.Duration(assessment.Duration) * time.Minute)

			if statusRand < 0.7 { // 70% submitted
				status = entities.StatusSubmitted
				// SubmittedAt bisa sebelum atau pada expectedEndTime
				maxSubmissionDelay := int(time.Duration(assessment.Duration) * time.Minute / time.Hour) // Convert to hours
				if maxSubmissionDelay < 1 {
					maxSubmissionDelay = 1
				}
				submittedAt = createdAt.Add(time.Duration(rand.Intn(maxSubmissionDelay*60)) * time.Minute)
				
				// EndedTime adalah waktu ketika submission selesai (bisa sama dengan submittedAt atau expectedEndTime)
				if submittedAt.Before(expectedEndTime) {
					endedTime = submittedAt // Selesai lebih awal
				} else {
					endedTime = expectedEndTime // Selesai tepat waktu
				}
			} else if statusRand < 0.9 { // 20% in progress
				status = entities.StatusInProgress
				submittedAt = time.Time{} // Zero time for in progress
				endedTime = expectedEndTime // Expected end time
			} else { // 10% todo
				status = entities.StatusTodo
				submittedAt = time.Time{} // Zero time for todo
				endedTime = expectedEndTime // Expected end time
			}

			submission := entities.Submission{
				UserID:       student.UserID,
				AssessmentID: assessment.ID,
				EndedTime:    endedTime,
				SubmittedAt:  submittedAt,
				Status:       status,
				CreatedAt:    createdAt, // Tambahkan CreatedAt
				UpdatedAt:    time.Now(),
			}

			// Only calculate score for submitted assessments
			if status == entities.StatusSubmitted {
				// Get questions for this assessment
				var questions []entities.Question
				db.Where("assessment_id = ?", assessment.ID).Find(&questions)
				
				totalQuestions := len(questions)
				correctAnswers := 0

				// Prepare answers (but don't create them yet)
				var answersToCreate []entities.Answer

				// Create answer data for each question
				for _, question := range questions {
					// Get choices for this question
					var choices []entities.Choice
					db.Where("question_id = ?", question.ID).Find(&choices)
					
					if len(choices) == 0 {
						continue
					}

					// Simulate student answering (70% chance of correct answer)
					var selectedChoice entities.Choice
					if rand.Float32() < 0.7 {
						// Try to find correct answer
						for _, choice := range choices {
							if choice.IsCorrect {
								selectedChoice = choice
								correctAnswers++
								break
							}
						}
					}
					
					// If no correct answer selected, pick random
					if selectedChoice.ID == uuid.Nil {
						selectedChoice = choices[rand.Intn(len(choices))]
					}

					// Prepare answer record (don't create yet)
					answer := entities.Answer{
						QuestionID: question.ID,
						ChoiceID:   selectedChoice.ID,
						// SubmissionID will be set after submission is created
						CreatedAt:  time.Now(),
						UpdatedAt:  time.Now(),
					}
					
					answersToCreate = append(answersToCreate, answer)
				}

				// Calculate score
				if totalQuestions > 0 {
					submission.Score = float64(correctAnswers) / float64(totalQuestions) * 100
				}

				// Create submission first
				if err := db.Create(&submission).Error; err != nil {
					log.Printf("Error creating submission: %v", err)
					continue
				}

				// Now create answers with the correct submission ID
				for _, answer := range answersToCreate {
					answer.SubmissionID = submission.ID // Now we have the real submission ID
					if err := db.Create(&answer).Error; err != nil {
						log.Printf("Error creating answer: %v", err)
					}
				}

				fmt.Printf("Created submission for %s in %s (Score: %.1f, Status: %s, Started: %s, Ended: %s)\n", 
					student.Username, assessment.Name, submission.Score, submission.Status, 
					createdAt.Format("2006-01-02 15:04"), endedTime.Format("2006-01-02 15:04"))
			} else {
				// For non-submitted status, just create the submission
				if err := db.Create(&submission).Error; err != nil {
					log.Printf("Error creating submission: %v", err)
				} else {
					fmt.Printf("Created submission for %s in %s (Score: %.1f, Status: %s, Started: %s, Expected End: %s)\n", 
						student.Username, assessment.Name, submission.Score, submission.Status,
						createdAt.Format("2006-01-02 15:04"), endedTime.Format("2006-01-02 15:04"))
				}
			}
		}
	}

	fmt.Printf("Created submissions and answers for assessments\n")
}

// ========== SUMMARY FUNCTIONS ==========
func PrintAssessmentSummary(db *gorm.DB) {
	fmt.Println("\n========== ASSESSMENT DATA SUMMARY ==========")
	
	var assessments []entities.Assessment
	db.Preload("Questions").Preload("Submissions").Find(&assessments)
	
	for _, assessment := range assessments {
		fmt.Printf("\nAssessment: %s\n", assessment.Name)
		fmt.Printf("  - Questions: %d\n", len(assessment.Questions))
		fmt.Printf("  - Submissions: %d\n", len(assessment.Submissions))
		fmt.Printf("  - Duration: %d minutes\n", assessment.Duration)
		fmt.Printf("  - Class ID: %s\n", assessment.ClassID)
		
		// Count submissions by status
		statusCount := make(map[entities.ExamStatus]int)
		totalScore := 0.0
		submittedCount := 0
		
		for _, submission := range assessment.Submissions {
			statusCount[submission.Status]++
			if submission.Status == entities.StatusSubmitted {
				totalScore += submission.Score
				submittedCount++
			}
		}
		
		fmt.Printf("  - Status breakdown:\n")
		for status, count := range statusCount {
			fmt.Printf("    • %s: %d\n", status, count)
		}
		
		if submittedCount > 0 {
			avgScore := totalScore / float64(submittedCount)
			fmt.Printf("  - Average score: %.1f%%\n", avgScore)
		}
	}
}