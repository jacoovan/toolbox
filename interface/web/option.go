package web

type Option interface {
	apply(c *ginImp)
}

type optionFunc func(c *ginImp)

func (fn optionFunc) apply(c *ginImp) {
	fn(c)
}
