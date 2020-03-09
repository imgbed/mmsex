package tgbot

type Options struct {
	alive int64
}

type OptionHandleFunc func(options *Options)

func SetAliveTime(alive int64)  OptionHandleFunc{
	return func(options *Options) {
		options.alive = alive
	}
}