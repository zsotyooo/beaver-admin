package googleauth

import (
	"context"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"google.golang.org/api/idtoken"
)

func TestGoogleAuth(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoogleAuth Suite")
}

type MockIDTokenWrapper struct {
	ValidateFunc func(ctx context.Context, rawToken string, audience string) (*idtoken.Payload, error)
}

func (w MockIDTokenWrapper) Validate(ctx context.Context, rawToken string, audience string) (*idtoken.Payload, error) {
	return w.ValidateFunc(ctx, rawToken, audience)
}

var _ = Describe("GoogleAuth", func() {
	Describe("VerifyToken", func() {
		var (
			token string
			email string
			name  string
		)

		BeforeEach(func() {
			token = "IGNORED"
			email = "test@example.com"
			name = "Test User"
			TokenValidator = MockIDTokenWrapper{
				ValidateFunc: func(ctx context.Context, rawToken string, audience string) (*idtoken.Payload, error) {
					return &idtoken.Payload{
						Claims: map[string]interface{}{
							"email": email,
							"name":  name,
						},
					}, nil
				},
			}
		})

		It("should verify the token without error", func() {
			_, err := VerifyToken(token)
			Expect(err).NotTo(HaveOccurred())
		})

		It("should return a GoogleUser with the correct email", func() {
			user, _ := VerifyToken(token)
			Expect(user.Email).To(Equal(email))
			Expect(user.Name).To(Equal(name))
		})
	})
})
