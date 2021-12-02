package model

type (
	// Broker notify broker list, call sub action, please implement IBroker before use Broker
	Broker struct {
		obs  []IBroker
		call chan interface{}
	}

	// Cloneable	use Cloneable before implement ICloneable
	Cloneable struct {
		obs map[string]ICloneable
	}

	// Chunk  use Chunk batch handler data
	Chunk struct{}

	// Operator use Operator before implement  IOperatorFactory
	Operator struct {
		a int
		b int
	}

	// Facade use Facade before implement IFacade with self struct
	Facade struct {
		/*a FacadeImlA
		b FacadeImlB*/
		obs []IFacadeImplement
	}
)
