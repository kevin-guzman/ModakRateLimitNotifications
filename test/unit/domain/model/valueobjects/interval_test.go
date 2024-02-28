package valueobjects_test

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"modak_golang/internal/domain/model/valueobjects"
)

var _ = Describe("IntervalType Tests", func() {
	Context("Testing IsUnderTimeRage", func() {
		It("Should return true because it is under time range", func() {
			Expect(valueobjects.Minute.IsUnderTimeRage(time.Now().Add(-time.Minute))).To(Equal(true))
			Expect(valueobjects.Hour.IsUnderTimeRage(time.Now().Add(-time.Hour))).To(Equal(true))
			Expect(valueobjects.Day.IsUnderTimeRage(time.Now().Add(-time.Hour * 24))).To(Equal(true))
		})

		It("Should return false because it is not under time range", func() {
			Expect(valueobjects.Minute.IsUnderTimeRage(time.Now())).To(Equal(false))
			Expect(valueobjects.Hour.IsUnderTimeRage(time.Now())).To(Equal(false))
			Expect(valueobjects.Day.IsUnderTimeRage(time.Now())).To(Equal(false))
		})
	})
})
