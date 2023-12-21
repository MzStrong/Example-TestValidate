package unit

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
	"github.com/tanapon395/sa-66-example/entity"
)

func TestAll(t *testing.T) {

	g := NewGomegaWithT(t)

	t.Run(`test all check`, func(t *testing.T) {
		member := entity.Member{
			PhoneNumber: "0123456789",
			Password:    "",
			Url:         "https://www.linkedin.com/company/ilink/",
		}

		ok, err := govalidator.ValidateStruct(member)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})
}

func TestPhoneNumber(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run(`PhoneNumber length is not 10 digits.`, func(t *testing.T) {
		member := entity.Member{
			PhoneNumber: "012345678",
			Password:    "",
			Url:         "https://www.linkedin.com/company/ilink/",
		}

		ok, err := govalidator.ValidateStruct(member)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("PhoneNumber length is not 10 digits."))

	})
}

func TestUrl(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run(`Url is invalid`, func(t *testing.T) {
		member := entity.Member{
			PhoneNumber: "0123456789",
			Password:    "",
			Url:         ".linkedin.com/company/ilink/",
		}

		ok, err := govalidator.ValidateStruct(member)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("Url is invalid"))

	})
}
