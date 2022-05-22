package app

type Category struct {
	Name  string `mapstructure:"name"`
	Tools []Tool `mapstructure:"tools"`
}

type Tool struct {
	Name        string `mapstructure:"name"`
	Description string `mapstructure:"description"`
	URL         string `mapstructure:"url"`
}
