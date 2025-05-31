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
		// Web Programming Assessments
		{
			Name:        "Quiz HTML & CSS Dasar",
			Description: "Quiz tentang konsep dasar HTML dan CSS untuk pemrograman web",
			ClassID:     ClassWebProgID,
			Duration:    3600,
			Questions: []QuestionTemplate{
				{
					Text: "Apa fungsi utama dari tag HTML <div>?",
					Choices: []ChoiceTemplate{
						{Text: "Membuat tabel", IsCorrect: false},
						{Text: "Membuat container atau pembungkus elemen", IsCorrect: true},
						{Text: "Membuat link", IsCorrect: false},
						{Text: "Membuat form input", IsCorrect: false},
					},
				},
				{
					Text: "CSS property mana yang digunakan untuk mengatur warna background?",
					Choices: []ChoiceTemplate{
						{Text: "color", IsCorrect: false},
						{Text: "background-color", IsCorrect: true},
						{Text: "font-color", IsCorrect: false},
						{Text: "text-color", IsCorrect: false},
					},
				},
				{
					Text: "Apa kepanjangan dari CSS?",
					Choices: []ChoiceTemplate{
						{Text: "Computer Style Sheets", IsCorrect: false},
						{Text: "Creative Style Sheets", IsCorrect: false},
						{Text: "Cascading Style Sheets", IsCorrect: true},
						{Text: "Colorful Style Sheets", IsCorrect: false},
					},
				},
			},
		},
		{
			Name:        "Quiz JavaScript & DOM",
			Description: "Quiz tentang JavaScript dan manipulasi DOM",
			ClassID:     ClassWebProgID,
			Duration:    3600,
			Questions: []QuestionTemplate{
				{
					Text: "Method JavaScript mana yang digunakan untuk mengambil elemen berdasarkan ID?",
					Choices: []ChoiceTemplate{
						{Text: "getElementById()", IsCorrect: true},
						{Text: "getElementByClass()", IsCorrect: false},
						{Text: "querySelector()", IsCorrect: false},
						{Text: "findElement()", IsCorrect: false},
					},
				},
				{
					Text: "Apa output dari console.log(typeof null) di JavaScript?",
					Choices: []ChoiceTemplate{
						{Text: "null", IsCorrect: false},
						{Text: "undefined", IsCorrect: false},
						{Text: "object", IsCorrect: true},
						{Text: "string", IsCorrect: false},
					},
				},
			},
		},

		// Database Assessments
		{
			Name:        "Quiz SQL Dasar",
			Description: "Quiz tentang perintah SQL dasar dan relational database",
			ClassID:     ClassDatabaseID,
			Duration:    3600,
			Questions: []QuestionTemplate{
				{
					Text: "Perintah SQL mana yang digunakan untuk mengambil data dari database?",
					Choices: []ChoiceTemplate{
						{Text: "GET", IsCorrect: false},
						{Text: "SELECT", IsCorrect: true},
						{Text: "FETCH", IsCorrect: false},
						{Text: "RETRIEVE", IsCorrect: false},
					},
				},
				{
					Text: "Apa fungsi dari PRIMARY KEY dalam database?",
					Choices: []ChoiceTemplate{
						{Text: "Mengurutkan data", IsCorrect: false},
						{Text: "Mengenkripsi data", IsCorrect: false},
						{Text: "Mengidentifikasi record secara unik", IsCorrect: true},
						{Text: "Membuat backup data", IsCorrect: false},
					},
				},
				{
					Text: "Perintah SQL mana yang digunakan untuk menambah data baru?",
					Choices: []ChoiceTemplate{
						{Text: "ADD", IsCorrect: false},
						{Text: "INSERT", IsCorrect: true},
						{Text: "CREATE", IsCorrect: false},
						{Text: "PUT", IsCorrect: false},
					},
				},
			},
		},
		{
			Name:        "Quiz Database Design",
			Description: "Quiz tentang perancangan database dan normalisasi",
			ClassID:     ClassDatabaseID,
			Duration:    3600,
			Questions: []QuestionTemplate{
				{
					Text: "Apa tujuan utama dari normalisasi database?",
					Choices: []ChoiceTemplate{
						{Text: "Mempercepat query", IsCorrect: false},
						{Text: "Mengurangi redundansi data", IsCorrect: true},
						{Text: "Menambah storage", IsCorrect: false},
						{Text: "Meningkatkan keamanan", IsCorrect: false},
					},
				},
				{
					Text: "Apa itu FOREIGN KEY?",
					Choices: []ChoiceTemplate{
						{Text: "Key yang dienkripsi", IsCorrect: false},
						{Text: "Key yang mereferensikan PRIMARY KEY tabel lain", IsCorrect: true},
						{Text: "Key yang tidak boleh null", IsCorrect: false},
						{Text: "Key yang unique", IsCorrect: false},
					},
				},
			},
		},

		// Algorithm Assessments
		{
			Name:        "Quiz Algoritma Sorting",
			Description: "Quiz tentang algoritma pengurutan data",
			ClassID:     ClassAlgorithmID,
			Duration:    3600,
			Questions: []QuestionTemplate{
				{
					Text: "Kompleksitas waktu rata-rata dari algoritma Quick Sort adalah?",
					Choices: []ChoiceTemplate{
						{Text: "O(n)", IsCorrect: false},
						{Text: "O(n log n)", IsCorrect: true},
						{Text: "O(n²)", IsCorrect: false},
						{Text: "O(log n)", IsCorrect: false},
					},
				},
				{
					Text: "Algoritma sorting mana yang paling efisien untuk data yang sudah hampir terurut?",
					Choices: []ChoiceTemplate{
						{Text: "Bubble Sort", IsCorrect: false},
						{Text: "Insertion Sort", IsCorrect: true},
						{Text: "Selection Sort", IsCorrect: false},
						{Text: "Merge Sort", IsCorrect: false},
					},
				},
				{
					Text: "Apa yang dimaksud dengan 'stable sorting algorithm'?",
					Choices: []ChoiceTemplate{
						{Text: "Algoritma yang tidak pernah error", IsCorrect: false},
						{Text: "Algoritma yang mempertahankan urutan relatif elemen yang sama", IsCorrect: true},
						{Text: "Algoritma yang selalu cepat", IsCorrect: false},
						{Text: "Algoritma yang menggunakan sedikit memori", IsCorrect: false},
					},
				},
			},
		},
		{
			Name:        "Quiz Struktur Data",
			Description: "Quiz tentang berbagai struktur data dan penggunaannya",
			ClassID:     ClassAlgorithmID,
			Duration:    3600,
			Questions: []QuestionTemplate{
				{
					Text: "Struktur data mana yang menggunakan prinsip LIFO (Last In First Out)?",
					Choices: []ChoiceTemplate{
						{Text: "Queue", IsCorrect: false},
						{Text: "Stack", IsCorrect: true},
						{Text: "Array", IsCorrect: false},
						{Text: "Linked List", IsCorrect: false},
					},
				},
				{
					Text: "Operasi apa yang paling efisien pada Binary Search Tree yang seimbang?",
					Choices: []ChoiceTemplate{
						{Text: "Insert dengan kompleksitas O(1)", IsCorrect: false},
						{Text: "Search dengan kompleksitas O(log n)", IsCorrect: true},
						{Text: "Delete dengan kompleksitas O(n)", IsCorrect: false},
						{Text: "Traversal dengan kompleksitas O(1)", IsCorrect: false},
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