package main

import "fmt"

type Worker struct {
	Name             string
	IsMale           bool
	Age              int
	Nationality      string
	Spec             string
	ProgressRate     float64
	WorkHourDuration float64
	Organization
}

type Organization struct {
	CompanyName  string
	Campus       string
	NumOfFellows int
	Pay          float64
	Location     string
	OpeningHour  float64
}

// One method for Worker to access both Worker and Organization fields
func (w Worker) ComposeWorkerArticle() {
	// Dynamic pronouns and titles
	pronoun := "he"
	title := "Mr"

	if !w.IsMale {
		pronoun = "she"
		title = "Mrs"
	}

	// Work closing time
	closing := w.OpeningHour + w.WorkHourDuration -12.0

	fmt.Printf(
		`This is %s.%s, %s is %d years old and hails from %s.
%s specializes in %s and currently works at %s located in %s.

The organization operates from the %s campus and has %d fellows.
%s resumes work every day at %.2fa.m and closes by %.2fp.m after working for %.2f hours.

%s earns %.2f and maintains a progress rate of %.2f%%.

`,
		title,
		w.Name,
		pronoun,
		w.Age,
		w.Nationality,
		pronoun,
		w.Spec,
		w.CompanyName,
		w.Location,
		w.Campus,
		w.NumOfFellows,
		pronoun,
		w.OpeningHour,
		closing,
		w.WorkHourDuration,
		pronoun,
		w.Pay,
		w.ProgressRate,
	)
}

func main() {
	worker := Worker{
		Name:             "Michael",
		IsMale:           true,
		Age:              24,
		Nationality:      "Nigeria",
		Spec:             "Backend Engineering",
		ProgressRate:     91.5,
		WorkHourDuration: 8,

		Organization: Organization{
			CompanyName:  "Learn2Earn",
			Campus:       "Yaba",
			NumOfFellows: 150,
			Pay:          2500.75,
			Location:     "Lagos",
			OpeningHour:  8.00,
		},
	}

	worker.ComposeWorkerArticle()
}
