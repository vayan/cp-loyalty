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

var _ = Describe("App", func() {
	var request *http.Request
	var recorder *httptest.ResponseRecorder
	var user User
	var a App

	BeforeEach(func() {
		recorder = httptest.NewRecorder()
		a = App{}
		a.Initialize("test.db")
		a.DB.AutoMigrate(&User{})
	})

	AfterEach(func() {
		a.DB.DropTable(&User{})
	})

	Describe("GET /users/:id", func() {
		BeforeEach(func() {
			user = User{LoyaltyPoint: 0}
			a.DB.Create(&user)
			route := fmt.Sprintf("/users/%d", user.ID)
			request, _ = http.NewRequest("GET", route, nil)
		})

		It("returns a status code of 200", func() {
			a.Router.ServeHTTP(recorder, request)
			Expect(recorder.Code).To(Equal(200))
		})

		It("returns a json serialized user", func() {
			marshal_user, _ := json.Marshal(user)
			a.Router.ServeHTTP(recorder, request)
			Expect(recorder.Body.String()).To(Equal(string(marshal_user)))
		})
	})

	Describe("POST /users", func() {
		Context("with valid JSON", func() {
			BeforeEach(func() {
				user = User{LoyaltyPoint: 256}
				body, _ := json.Marshal(user)
				request, _ = http.NewRequest("POST", "/users", bytes.NewReader(body))
			})

			It("returns a status code of 200", func() {
				a.Router.ServeHTTP(recorder, request)
				Expect(recorder.Code).To(Equal(200))
			})

			It("returns the newly created user", func() {
				a.Router.ServeHTTP(recorder, request)
				Expect(recorder.Body.String()).To(ContainSubstring("\"loyalty_point\":256"))
			})
		})

		Context("with invalid JSON", func() {
			BeforeEach(func() {
				user = User{LoyaltyPoint: -1}
				body, _ := json.Marshal(user)
				request, _ = http.NewRequest("POST", "/users", bytes.NewReader(body))
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
