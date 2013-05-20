package main

import (
)

type Student struct {
	UUID string
	Name string
	Resume string
	Skills string
	Major string
	Minor string
	Interests string
	Email string

	//Users auth token
	token string
}

func NewStudent(name, resume, skills, major, minor, interests, email string) *Student {
	s := new(Student)
	s.Name = name
	s.Major = major
	s.Minor = minor
	s.Resume = resume
	s.Skills = skills
	s.Interests = interests
	s.Email = email
	return s
}

func (s *Student) Equal(o *Student) bool {
	return s.Name == o.Name &&
			s.Major == o.Major &&
			s.Minor == o.Minor &&
			s.Resume == o.Resume &&
			s.Skills == o.Skills &&
			s.Interests == o.Interests &&
			s.Email == o.Email &&
			s.UUID == o.UUID
}

//Update student with the non-blank values in 'us'
//return true if changes were made
func (stu *Student) Update(us *Student) bool {
	change := false
	if us.Resume != "" {
		stu.Resume = us.Resume
		change = true
	}
	if us.Major != "" {
		stu.Major = us.Major
		change = true
	}
	if us.Minor != "" {
		stu.Minor = us.Minor
		change = true
	}
	if us.Interests != "" {
		stu.Interests = us.Interests
		change = true
	}
	if us.Email != "" {
		stu.Email = us.Email
		change = true
	}
	if us.Skills != "" {
		stu.Skills = us.Skills
		change = true
	}
	return change
}
