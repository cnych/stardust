package cache

type IsNotFoundErrorPred func(err error) bool

var (
	notFoundErrorPreds []IsNotFoundErrorPred = make([]IsNotFoundErrorPred, 0, 4)
)

func IsNotFoundErr(err error) bool {
	for _, pred := range notFoundErrorPreds {
		if pred(err) {
			return true
		}
	}
	return false
}

func RegNotFoundErrPred(preds ...IsNotFoundErrorPred) {
	for _, pred := range preds {
		if pred != nil {
			notFoundErrorPreds = append(notFoundErrorPreds, pred)
		}
	}
}

func AsNotFoundErrPred(err error) IsNotFoundErrorPred {
	return func(err1 error) bool {
		return err1 == err
	}
}
