package jwtauth

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestJwtAuth(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "JwtAuth Suite")
}

var _ = Describe("JwtAuth", func() {
	var (
		jwtAuth *JwtAuth
		email   string
	)

	BeforeEach(func() {
		jwtAuth = NewJwtAuth()
		email = "test@example.com"
	})

	Describe("GenerateToken", func() {
		It("should generate a token without error", func() {
			token, err := jwtAuth.GenerateToken(email)
			Expect(err).NotTo(HaveOccurred())
			Expect(token).NotTo(BeEmpty())
		})

		It("should generate a token with the correct claims", func() {
			token, _ := jwtAuth.GenerateToken(email)
			claims, _ := jwtAuth.ValidateToken(token)

			Expect(claims.Email).To(Equal(email))
			Expect(claims.Issuer).To(Equal(jwtAuth.Issuer))
		})
	})
})
