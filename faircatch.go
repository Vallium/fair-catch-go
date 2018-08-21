package faircatch

type TCR struct {
	Try     func()
	Catch   func(Exception)
	Recover func()
}

type Exception interface{}

func Throw(up Exception) {
	panic(up)
}

func (tcr TCR) Do() {
	if tcr.Recover != nil {
		defer tcr.Recover()
	}
	if tcr.Catch != nil {
		defer func() {
			if r := recover(); r != nil {
				tcr.Catch(r)
			}
		}()
	}
	tcr.Try()
}
