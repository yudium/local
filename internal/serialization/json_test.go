package serialization_test

import (
	"encoding/json"

	"github.com/go-seidon/local/internal/serialization"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Json Serialization Package", func() {
	Context("NewJsonSerializer function", func() {
		When("function is called", func() {
			It("should return result", func() {
				res := serialization.NewJsonSerializer()

				Expect(res).ToNot(BeNil())
			})
		})
	})

	Context("Encode function", func() {
		var (
			serializer serialization.Serializer
		)

		BeforeEach(func() {
			serializer = serialization.NewJsonSerializer()
		})

		When("success encode data", func() {
			It("should return result", func() {
				d := struct{}{}
				res, err := serializer.Encode(d)

				expected, _ := json.Marshal(d)
				Expect(err).To(BeNil())
				Expect(res).To(Equal(expected))
			})
		})

		When("failed encode data", func() {
			It("should return error", func() {
				d := make(chan int)
				res, err := serializer.Encode(d)

				Expect(err).ToNot(BeNil())
				Expect(res).To(BeNil())
			})
		})

		When("data is nil", func() {
			It("should return result", func() {
				res, err := serializer.Encode(nil)

				expected, _ := json.Marshal(nil)
				Expect(err).To(BeNil())
				Expect(res).To(Equal(expected))
			})
		})
	})

	Context("Decode function", func() {
		type Data struct {
			Val string `json:"val"`
		}

		var (
			d          []byte
			serializer serialization.Serializer
		)

		BeforeEach(func() {
			serializer = serialization.NewJsonSerializer()
			d = []byte(`{"val":"ok"}`)
		})

		When("success decode data", func() {
			It("should return result", func() {
				var res Data
				err := serializer.Decode(d, &res)

				Expect(err).To(BeNil())
				Expect(res.Val).To(Equal("ok"))
			})
		})

		When("failed decode data", func() {
			It("should return error", func() {
				var res Data
				d = []byte{}
				err := serializer.Decode(d, &res)

				Expect(err).ToNot(BeNil())
				Expect(res.Val).To(Equal(""))
			})
		})

		When("data is nil", func() {
			It("should return result", func() {
				var res Data
				err := serializer.Decode(nil, &res)

				Expect(err.Error()).To(Equal("unexpected end of JSON input"))
				Expect(res.Val).To(Equal(""))
			})
		})
	})
})
