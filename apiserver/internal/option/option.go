package option

import "postcard/apiserver/internal/pkg/http"

type Options struct {
	WebOptions *http.WebOption
	//BotOptions *
}

func NewOptions() *Options {
	return &Options{
		WebOptions: http.NewWebOptions(),
	}
}
