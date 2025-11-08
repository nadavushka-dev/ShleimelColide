package config

type Config struct {
	ScreenWidth   int
	ScreenHeight  int
	FrameOX       int
	FrameOY       int
	TicksPerFrame int
}

func GetDefaultConfig() Config {
	return Config{
		ScreenWidth:   640,
		ScreenHeight:  480,
		FrameOX:       0,
		FrameOY:       32,
		TicksPerFrame: 10,
	}
}
