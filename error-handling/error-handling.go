// Package erratum is the implementaion of the error-handling exercise
package erratum

const testVersion = 2

// Use is the implementaion of the error-handling exercise
// see README.md for further information
func Use(o ResourceOpener, input string) (result error) {
	resource, err := open(o)
	if err != nil {
		return err
	}
	defer resource.Close()

	defer func() {
		if recover := recover(); recover != nil {
			if asFrob, ok := recover.(FrobError); ok {
				resource.Defrob(asFrob.defrobTag)
			}
			if asError, ok := recover.(error); ok {
				result = asError
			}
		}
	}()
	resource.Frob(input)
	return result
}

func open(o ResourceOpener) (res Resource, err error) {
	for {
		res, err = o()

		if _, ok := err.(TransientError); ok {
			continue
		} else {
			break
		}
	}
	return
}
