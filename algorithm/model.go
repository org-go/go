package algorithm

import (
	"math/rand"
)

type RandomGenerator struct {
	*rand.Rand
}

// Model is the interface for all models. Any model in this
// package should implement it.
type Model interface {
	SetParams(params Params)
	GetParams() Params
	GetParamsGrid() ParamsGrid
	Clear()
	Invalid() bool
}

// BaseModel model must be included by every recommendation model. Hyper-parameters,
// ID sets, random generator and fitting options are managed the BaseModel model.
type BaseModel struct {
	Params    Params          // Hyper-parameters
	rng       RandomGenerator // Random generator
	randState int64           // Random seed
}

func NewRandomGenerator(seed int64) RandomGenerator {
	return RandomGenerator{rand.New(rand.NewSource(int64(seed)))}
}

// SetParams sets hyper-parameters for the BaseModel model.
func (model *BaseModel) SetParams(params Params) {
	model.Params = params
	model.randState = model.Params.GetInt64(RandomState, 0)
	model.rng = NewRandomGenerator(model.randState)
}

// GetParams returns all hyper-parameters.
func (model *BaseModel) GetParams() Params {
	return model.Params
}

func (model *BaseModel) GetRandomGenerator() RandomGenerator {
	return model.rng
}
