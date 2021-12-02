package threads

type ISafe interface {
	// Run
	Run(func())
	// SafeRun
	SafeRun(func())
	// Wait
	Wait()
}

var (
	ServiceMux ISafe = new(mux)

	//ServiceGpx ISafe = new(groupx)
)
