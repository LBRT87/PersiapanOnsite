package domain

import "context"

type CourseRepository interface {
	FindByID(c context.Context, id string) (*Course, error)
	Create(c context.Context, course *Course)
}
type ModuleRepository interface{
	FindByID (c context.Context, id string) (*Module, error)
}
type ComponentRepository interface{
	FindByID(c context.Context, id string) (*Component, error)
}
type EnrollmentRepository interface{
	FindByID(c context.Context, id string) (*Enrollment, error)
	FindByUserCourse(c context.Context, userId, lecturerId string) (*Enrollment, error)
}