package interfaces

import "github.com/IsFariza/ap2-Migrations/doctor-service/internal/model"

type DoctorPublisher interface {
	PublishDoctorCreated(doc *model.Doctor) error
}
