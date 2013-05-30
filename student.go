package main

import (
)

type Skill struct {
	Name string
	Value string
}

type Student struct {
	UUID string
	FirstName string
	LastName string
	Resume string
	Skills []Skill
	Major string
	Minors []string
	Interests []string
	Email string
	Projects []string

	//Users auth token
	token string

	//users password hash
	password string
}

func NewStudent(uuid, fname, lname, resume, major, email string, minors, interests []string, skills []Skill) *Student {
	s := new(Student)
	s.UUID = uuid
	s.FirstName = fname
	s.LastName = lname
	s.Major = major
	s.Minors = minors
	s.Resume = resume
	s.Skills = skills
	s.Interests = interests
	s.Email = email

	return s
}

func (s *Student) Equal(o *Student) bool {
	return s.FirstName == o.FirstName &&
			s.LastName == o.LastName &&
			s.Major == o.Major &&
			s.Resume == o.Resume &&
			s.Email == o.Email &&
			s.UUID == o.UUID &&
			CompLists(s.Minors, o.Minors) &&
			CompLists(s.Interests, o.Interests)
}

//Update student with the non-blank values in 'us'
//return true if changes were made
func (stu *Student) Update(us *Student) bool {
	if stu.UUID != us.UUID {
		panic("students must have same UUID to update!")
	}
	change := false
	if us.Resume != "" {
		stu.Resume = us.Resume
		change = true
	}
	if us.Major != "" {
		stu.Major = us.Major
		change = true
	}
	if us.Minors != nil {
		stu.Minors = us.Minors
		change = true
	}
	if us.Interests != nil {
		stu.Interests = us.Interests
		change = true
	}
	if us.Email != "" {
		stu.Email = us.Email
		change = true
	}
	if us.Skills != nil {
		stu.Skills = us.Skills
		change = true
	}
	return change
}

func CompLists(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}
