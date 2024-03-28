package apiserver

import (
	conf "postcard/apiserver/config"
	"postcard/apiserver/internal/app"
	"postcard/apiserver/internal/option"
)

func NewAPIServer() *app.App {
	opts := option.NewOptions()
	application := app.NewApp("API Server",
		conf.ProjectName,
		app.WithOptions(opts),
		app.WithRunFunc(run(opts)),
	)

	return application
}

func run(opts *option.Options) app.RunFunc {
	return func(basename string) error {
		return nil
	}
}
