package erratum

func Use(opener ResourceOpener, input string) (err error) {
	var r Resource

	for {
		r, err = opener()
		if err != nil {
			if _, ok := err.(TransientError); ok {
				continue // try again
			}
			return err // different error
		}
		break
	}

	defer func() {
		cerr := r.Close()
		if err == nil {
			err = cerr
		}
	}()

	defer func() {
		if rec := recover(); rec != nil {
			switch e := rec.(type) {
			case FrobError:
				r.Defrob(e.defrobTag)
				err = e.inner
			case error:
				err = e
			default:
				panic(rec) // unexpected panic
			}
		}
	}()

	r.Frob(input)
	return nil
}
