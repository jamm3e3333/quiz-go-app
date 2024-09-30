package model

import (
	"sync"
)

type Submission struct {
	CorrectAnswers int32
	TotalQuestions int32
}

type SubmissionModel struct {
	submissions []Submission
	sync.RWMutex
}

func NewSubmissionModel() *SubmissionModel {
	return &SubmissionModel{
		submissions: []Submission{},
	}
}

func (s *SubmissionModel) SaveSubmission(submission Submission) {
	s.Lock()
	defer s.Unlock()

	s.submissions = append(s.submissions, submission)
}

func (s *SubmissionModel) Submissions() []Submission {
	s.RLock()
	defer s.RUnlock()

	return s.submissions
}
