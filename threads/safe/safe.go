package safe

import "log"

func GoSafe(cf func()) {
	go RunSafe(cf)
}

func RunSafe(cf func()) {
	defer Recovers()
	cf()
}

func Recovers(cfs ...func()) {

	for _, f := range cfs {
		f()
	}
	if err := recover(); err != nil {
		log.Println(err)
	}

}
