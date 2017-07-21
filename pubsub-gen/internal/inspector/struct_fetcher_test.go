package inspector_test

import (
	"go/parser"
	"go/token"
	"testing"

	"github.com/apoydence/onpar"
	. "github.com/apoydence/onpar/expect"
	. "github.com/apoydence/onpar/matchers"
	"github.com/apoydence/pubsub/pubsub-gen/internal/inspector"
)

type TSF struct {
	*testing.T
	f inspector.StructFetcher
}

func TestStructFetcher(t *testing.T) {
	t.Parallel()
	o := onpar.New()
	defer o.Run(t)

	o.BeforeEach(func(t *testing.T) TSF {
		return TSF{
			T: t,
			f: inspector.NewStructFetcher(),
		}
	})

	o.Group("primitive fields", func() {
		o.Spec("it parses and returns a single struct", func(t TSF) {
			src := `
package p
type x struct {
	i string
	j int
}
`
			fset := token.NewFileSet()
			n, err := parser.ParseFile(fset, "src.go", src, 0)
			Expect(t, err == nil).To(BeTrue())

			s, err := t.f.Parse(n)
			Expect(t, err == nil).To(BeTrue())
			Expect(t, s).To(HaveLen(1))
			Expect(t, s[0].Name).To(Equal("x"))
			Expect(t, s[0].Fields).To(HaveLen(2))

			Expect(t, s[0].Fields[0].Name).To(Equal("i"))
			Expect(t, s[0].Fields[0].Type).To(Equal("string"))

			Expect(t, s[0].Fields[1].Name).To(Equal("j"))
			Expect(t, s[0].Fields[1].Type).To(Equal("int"))
		})

		o.Spec("it parses multiple structs", func(t TSF) {
			src := `
package p
type x struct {
	i string
	j int
}

type y struct {
	i string
	j int
}
`
			fset := token.NewFileSet()
			n, err := parser.ParseFile(fset, "src.go", src, 0)
			Expect(t, err == nil).To(BeTrue())

			s, err := t.f.Parse(n)
			Expect(t, err == nil).To(BeTrue())

			Expect(t, s).To(HaveLen(2))
			Expect(t, s[0].Name).To(Equal("x"))
			Expect(t, s[1].Name).To(Equal("y"))
		})
	})
}