package main_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	. "github.com/vayan/itw-cp-loyalty"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Ride", func() {
	var request *http.Request
	var recorder *httptest.ResponseRecorder
	var ride Ride
	var baseUser = User{LoyaltyPoint: 0}
	var a App

	BeforeEach(func() {
		recorder = httptest.NewRecorder()
		a = App{}
		a.Initialize("test.db")
		a.DB.Create(&baseUser)
		baseUser.SetBaseRank(a.DB)
		baseUser.Save(a.DB)
	})

	AfterEach(func() {
		a.DB.DropTable(&Ride{})
		a.DB.DropTable(&User{})
	})

	Describe("POST /rides", func() {
		Context("with valid JSON", func() {
			BeforeEach(func() {
				route := fmt.Sprintf("/users/%d/rides", baseUser.ID)
				ride = Ride{Price: 33, User: baseUser}
				body, _ := json.Marshal(ride)
				request, _ = http.NewRequest("POST", route, bytes.NewReader(body))
			})

			It("returns a status code of 200", func() {
				a.Router.ServeHTTP(recorder, request)
				Expect(recorder.Code).To(Equal(200))
			})

			It("returns the newly created ride", func() {
				a.Router.ServeHTTP(recorder, request)
				Expect(recorder.Body.String()).To(ContainSubstring("\"price\":33"))
			})

			It("raises the user loyalty points", func() {
				var user User
				a.Router.ServeHTTP(recorder, request)
				a.DB.First(&user, baseUser.ID)
				Expect(user.LoyaltyPoint).To(Equal(33))
			})

			It("raises the user loyalty ranks", func() {
				for i := 1; i < 5; i++ {
					a.DB.Create(&Ride{Price: i * 11, User: baseUser})
				}
				user := FetchUser(baseUser.ID, a.DB)
				Expect(user.LoyaltyRank.Name).To(Equal("bronze"))
				a.Router.ServeHTTP(recorder, request)
				user = FetchUser(baseUser.ID, a.DB)
				Expect(user.LoyaltyRank.Name).To(Equal("silver"))
			})

			It("raises the user loyalty points according to his rank", func() {
				for i := 1; i < 6; i++ {
					a.DB.Create(&Ride{Price: 1, User: baseUser})
				}
				a.Router.ServeHTTP(recorder, request)
				a.DB.First(&baseUser, baseUser.ID)
				Expect(baseUser.LoyaltyPoint).To(Equal(99))
			})
		})

		Context("with invalid JSON", func() {
			BeforeEach(func() {
				route := fmt.Sprintf("/users/%d/rides", baseUser.ID)
				ride = Ride{Price: 0, User: baseUser}
				body, _ := json.Marshal(ride)
				request, _ = http.NewRequest("POST", route, bytes.NewReader(body))
			})

			It("returns a status code of 400", func() {
				a.Router.ServeHTTP(recorder, request)
				Expect(recorder.Code).To(Equal(400))
			})

			It("returns an error", func() {
				a.Router.ServeHTTP(recorder, request)
				Expect(recorder.Body.String()).To(ContainSubstring("\"error\":\"Invalid Params\""))
			})
		})
	})
})
