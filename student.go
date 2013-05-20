package main

type Student struct {
	Name string
	Resume string
	Skills string
	Major string
	Minor string
	Interests string
	Email string
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
