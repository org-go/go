package model

type ICloneable interface {
	clone() ICloneable
}

/**
 * NewCloneable
 * @Description:
 * @return *Cloneable
 */
func NewCloneable() *Cloneable {
	return &Cloneable{obs: make(map[string]ICloneable)}
}

/**
 * set
 * @Description:
 * @receiver c
 * @param name
 * @param cloneable
 */
func (c *Cloneable) set(name string, cloneable ICloneable) {
	c.obs[name] = cloneable
}

/**
 * get
 * @Description:
 * @receiver c
 * @param name
 * @return ICloneable
 */
func (c *Cloneable) get(name string) ICloneable {
	return c.obs[name]
}
