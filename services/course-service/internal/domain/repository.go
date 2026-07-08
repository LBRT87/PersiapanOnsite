package domain

import "context"

type CourseRepository interface {
	FindByID(c context.Context, id string) (*Course, error)
	Create(c context.Context, course *Course) error
}
type ModuleRepository interface{
	FindByID (c context.Context, id string) (Module, error)
	Create(c context.Context, module *Module) error
}
type ComponentRepository interface{
	FindByID(c context.Context, id string) (*Component, error)
	Create(c context.Context, component *Component) error
}
type EnrollmentRepository interface{
	FindByID(c context.Context, id string) (*Enrollment, error)
	FindByUserCourse(c context.Context, userId, courseId string) (bool, error)
	Create(c context.Context, enrollment *Enrollment) error
}