package end2end_test

import (
	"github.com/apoydence/pubsub"
	"github.com/apoydence/pubsub/pubsub-gen/internal/end2end"
)

type StructTraverser struct{}

func NewStructTraverser() StructTraverser { return StructTraverser{} }

func (s StructTraverser) Traverse(data interface{}) pubsub.Paths {
	return s._I(data)
}

func (s StructTraverser) done(data interface{}) pubsub.Paths {
	return pubsub.Paths(func(idx int, data interface{}) (path interface{}, nextTraverser pubsub.TreeTraverser, ok bool) {
		return nil, nil, false
	})
}

func (s StructTraverser) _I(data interface{}) pubsub.Paths {

	return pubsub.Paths(func(idx int, data interface{}) (path interface{}, nextTraverser pubsub.TreeTraverser, ok bool) {
		switch idx {
		case 0:
			return nil, pubsub.TreeTraverser(s._J), true
		case 1:
			return data.(*end2end.X).I, pubsub.TreeTraverser(s._J), true
		default:
			return nil, nil, false
		}
	})
}

func (s StructTraverser) _J(data interface{}) pubsub.Paths {

	return pubsub.Paths(func(idx int, data interface{}) (path interface{}, nextTraverser pubsub.TreeTraverser, ok bool) {
		switch idx {
		case 0:
			return nil,
				pubsub.TreeTraverser(func(data interface{}) pubsub.Paths {
					return s.__Y1_Y2_M
				}), true
		case 1:
			return data.(*end2end.X).J,
				pubsub.TreeTraverser(func(data interface{}) pubsub.Paths {
					return s.__Y1_Y2_M
				}), true
		default:
			return nil, nil, false
		}
	})
}

func (s StructTraverser) __Y1_Y2_M(idx int, data interface{}) (path interface{}, nextTraverser pubsub.TreeTraverser, ok bool) {
	switch idx {

	case 0:

		return "Y1", pubsub.TreeTraverser(s._Y1_I), true

	case 1:

		if data.(*end2end.X).Y2 == nil {
			return nil, pubsub.TreeTraverser(s.done), true
		}

		return "Y2", pubsub.TreeTraverser(s._Y2_I), true

	case 2:
		switch data.(*end2end.X).M.(type) {
		case end2end.M1:
			return "M1", s._M_M1_A, true

		case end2end.M2:
			return "M2", s._M_M2_A, true

		default:
			return nil, pubsub.TreeTraverser(s.done), true
		}

	default:
		return nil, nil, false
	}
}

func (s StructTraverser) _Y1(data interface{}) pubsub.Paths {

	return pubsub.Paths(func(idx int, data interface{}) (path interface{}, nextTraverser pubsub.TreeTraverser, ok bool) {
		switch idx {
		case 0:
			return "Y1", pubsub.TreeTraverser(s._Y1_I), true
		default:
			return nil, nil, false
		}
	})
}

func (s StructTraverser) _Y1_I(data interface{}) pubsub.Paths {

	return pubsub.Paths(func(idx int, data interface{}) (path interface{}, nextTraverser pubsub.TreeTraverser, ok bool) {
		switch idx {
		case 0:
			return nil, pubsub.TreeTraverser(s._Y1_J), true
		case 1:
			return data.(*end2end.X).Y1.I, pubsub.TreeTraverser(s._Y1_J), true
		default:
			return nil, nil, false
		}
	})
}

func (s StructTraverser) _Y1_J(data interface{}) pubsub.Paths {

	return pubsub.Paths(func(idx int, data interface{}) (path interface{}, nextTraverser pubsub.TreeTraverser, ok bool) {
		switch idx {
		case 0:
			return nil, pubsub.TreeTraverser(s.done), true
		case 1:
			return data.(*end2end.X).Y1.J, pubsub.TreeTraverser(s.done), true
		default:
			return nil, nil, false
		}
	})
}

func (s StructTraverser) _Y2(data interface{}) pubsub.Paths {

	if data.(*end2end.X).Y2 == nil {
		return pubsub.Paths(func(idx int, data interface{}) (path interface{}, nextTraverser pubsub.TreeTraverser, ok bool) {
			switch idx {
			case 0:
				return nil, pubsub.TreeTraverser(s.done), true
			default:
				return nil, nil, false
			}
		})
	}

	return pubsub.Paths(func(idx int, data interface{}) (path interface{}, nextTraverser pubsub.TreeTraverser, ok bool) {
		switch idx {
		case 0:
			return "Y2", pubsub.TreeTraverser(s._Y2_I), true
		default:
			return nil, nil, false
		}
	})
}

func (s StructTraverser) _Y2_I(data interface{}) pubsub.Paths {

	return pubsub.Paths(func(idx int, data interface{}) (path interface{}, nextTraverser pubsub.TreeTraverser, ok bool) {
		switch idx {
		case 0:
			return nil, pubsub.TreeTraverser(s._Y2_J), true
		case 1:
			return data.(*end2end.X).Y2.I, pubsub.TreeTraverser(s._Y2_J), true
		default:
			return nil, nil, false
		}
	})
}

func (s StructTraverser) _Y2_J(data interface{}) pubsub.Paths {

	return pubsub.Paths(func(idx int, data interface{}) (path interface{}, nextTraverser pubsub.TreeTraverser, ok bool) {
		switch idx {
		case 0:
			return nil, pubsub.TreeTraverser(s.done), true
		case 1:
			return data.(*end2end.X).Y2.J, pubsub.TreeTraverser(s.done), true
		default:
			return nil, nil, false
		}
	})
}

func (s StructTraverser) _M_M1(data interface{}) pubsub.Paths {

	return pubsub.Paths(func(idx int, data interface{}) (path interface{}, nextTraverser pubsub.TreeTraverser, ok bool) {
		switch idx {
		case 0:
			return "M1", pubsub.TreeTraverser(s._M_M1_A), true
		default:
			return nil, nil, false
		}
	})
}

func (s StructTraverser) _M_M1_A(data interface{}) pubsub.Paths {

	return pubsub.Paths(func(idx int, data interface{}) (path interface{}, nextTraverser pubsub.TreeTraverser, ok bool) {
		switch idx {
		case 0:
			return nil, pubsub.TreeTraverser(s.done), true
		case 1:
			return data.(*end2end.X).M.(end2end.M1).A, pubsub.TreeTraverser(s.done), true
		default:
			return nil, nil, false
		}
	})
}

func (s StructTraverser) _M_M2(data interface{}) pubsub.Paths {

	return pubsub.Paths(func(idx int, data interface{}) (path interface{}, nextTraverser pubsub.TreeTraverser, ok bool) {
		switch idx {
		case 0:
			return "M2", pubsub.TreeTraverser(s._M_M2_A), true
		default:
			return nil, nil, false
		}
	})
}

func (s StructTraverser) _M_M2_A(data interface{}) pubsub.Paths {

	return pubsub.Paths(func(idx int, data interface{}) (path interface{}, nextTraverser pubsub.TreeTraverser, ok bool) {
		switch idx {
		case 0:
			return nil, pubsub.TreeTraverser(s._M_M2_B), true
		case 1:
			return data.(*end2end.X).M.(end2end.M2).A, pubsub.TreeTraverser(s._M_M2_B), true
		default:
			return nil, nil, false
		}
	})
}

func (s StructTraverser) _M_M2_B(data interface{}) pubsub.Paths {

	return pubsub.Paths(func(idx int, data interface{}) (path interface{}, nextTraverser pubsub.TreeTraverser, ok bool) {
		switch idx {
		case 0:
			return nil, pubsub.TreeTraverser(s.done), true
		case 1:
			return data.(*end2end.X).M.(end2end.M2).B, pubsub.TreeTraverser(s.done), true
		default:
			return nil, nil, false
		}
	})
}

type XFilter struct {
	I    *int
	J    *string
	Y1   *YFilter
	Y2   *YFilter
	M_M1 *M1Filter
	M_M2 *M2Filter
}

type YFilter struct {
	I *int
	J *string
}

type M1Filter struct {
	A *int
}

type M2Filter struct {
	A *int
	B *int
}

func (g StructTraverser) CreatePath(f *XFilter) []interface{} {
	if f == nil {
		return nil
	}
	var path []interface{}

	var count int
	if f.Y1 != nil {
		count++
	}

	if f.Y2 != nil {
		count++
	}

	if f.M_M1 != nil {
		count++
	}

	if f.M_M2 != nil {
		count++
	}

	if count > 1 {
		panic("Only one field can be set")
	}

	if f.I != nil {
		path = append(path, *f.I)
	} else {
		path = append(path, nil)
	}

	if f.J != nil {
		path = append(path, *f.J)
	} else {
		path = append(path, nil)
	}

	path = append(path, g.createPath_Y1(f.Y1)...)

	path = append(path, g.createPath_Y2(f.Y2)...)

	path = append(path, g.createPath_M_M1(f.M_M1)...)

	path = append(path, g.createPath_M_M2(f.M_M2)...)

	for i := len(path) - 1; i >= 1; i-- {
		if path[i] != nil {
			break
		}
		path = path[:i]
	}

	return path
}

func (g StructTraverser) createPath_Y1(f *YFilter) []interface{} {
	if f == nil {
		return nil
	}
	var path []interface{}

	path = append(path, "Y1")

	var count int
	if count > 1 {
		panic("Only one field can be set")
	}

	if f.I != nil {
		path = append(path, *f.I)
	} else {
		path = append(path, nil)
	}

	if f.J != nil {
		path = append(path, *f.J)
	} else {
		path = append(path, nil)
	}

	return path
}

func (g StructTraverser) createPath_Y2(f *YFilter) []interface{} {
	if f == nil {
		return nil
	}
	var path []interface{}

	path = append(path, "Y2")

	var count int
	if count > 1 {
		panic("Only one field can be set")
	}

	if f.I != nil {
		path = append(path, *f.I)
	} else {
		path = append(path, nil)
	}

	if f.J != nil {
		path = append(path, *f.J)
	} else {
		path = append(path, nil)
	}

	return path
}

func (g StructTraverser) createPath_M_M1(f *M1Filter) []interface{} {
	if f == nil {
		return nil
	}
	var path []interface{}

	path = append(path, "M1")

	var count int
	if count > 1 {
		panic("Only one field can be set")
	}

	if f.A != nil {
		path = append(path, *f.A)
	} else {
		path = append(path, nil)
	}

	return path
}

func (g StructTraverser) createPath_M_M2(f *M2Filter) []interface{} {
	if f == nil {
		return nil
	}
	var path []interface{}

	path = append(path, "M2")

	var count int
	if count > 1 {
		panic("Only one field can be set")
	}

	if f.A != nil {
		path = append(path, *f.A)
	} else {
		path = append(path, nil)
	}

	if f.B != nil {
		path = append(path, *f.B)
	} else {
		path = append(path, nil)
	}

	return path
}
