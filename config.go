package pikselkapcio

const (
	//TextGenerationRandom is used for pseudorandom alphanumeric string for code (default value)
	TextGenerationRandom = 0
	//TextGenerationCustomWords is used to randomly pick code from provided custom words list
	TextGenerationCustomWords = 1
	//ColorPairsRotationRandom is used for random selection of color pair for each character in code image (default value)
	ColorPairsRotationRandom = 0
	//ColorPairsRotationSequence is used for cycling through available color pairs in sequence for each character in code image
	ColorPairsRotationSequence = 1
)

//Config strucure for whole package
type Config struct {
	Scale               int
	TextGenerationMode  int8
	RandomTextLength    int8
	CustomWords         []string
	ColorHexStringPairs []HexStringPair
	ColorPairsRotation  int8
}

func getDefaultConfig() Config {
	return Config{
		Scale:               5,
		RandomTextLength:    4,
		CustomWords:         getDefaultCustomWords(),
		ColorHexStringPairs: getDefaultHexStringPairs(),
	}
}

func mergeConfig(customConfig Config) Config {
	config := getDefaultConfig()

	if customConfig.Scale != 0 {
		config.Scale = customConfig.Scale
	}

	if customConfig.TextGenerationMode != TextGenerationRandom {
		config.TextGenerationMode = customConfig.TextGenerationMode
	}

	if customConfig.RandomTextLength > 0 && customConfig.RandomTextLength < 37 {
		config.RandomTextLength = customConfig.RandomTextLength
	}

	if len(customConfig.CustomWords) > 0 {
		config.CustomWords = customConfig.CustomWords
	}

	if len(customConfig.ColorHexStringPairs) > 0 {
		config.ColorHexStringPairs = customConfig.ColorHexStringPairs
	}

	if customConfig.ColorPairsRotation != ColorPairsRotationRandom {
		config.ColorPairsRotation = customConfig.ColorPairsRotation
	}

	return config
}

func getDefaultHexStringPairs() []HexStringPair {
	hexStringPairs := []HexStringPair{
		{
			BackgroundColor: "CCCCCC",
			ForegroundColor: "888888",
		},
		{
			BackgroundColor: "888888",
			ForegroundColor: "CCCCCC",
		},
		{
			BackgroundColor: "00CC00",
			ForegroundColor: "97EA97",
		},
		{
			BackgroundColor: "97EA97",
			ForegroundColor: "00CC00",
		},
		{
			BackgroundColor: "9797EA",
			ForegroundColor: "5C5CDE",
		},
		{
			BackgroundColor: "5C5CDE",
			ForegroundColor: "9797EA",
		},
		{
			BackgroundColor: "FF8800",
			ForegroundColor: "FFCE97",
		},
		{
			BackgroundColor: "FFCE97",
			ForegroundColor: "FF8800",
		},
	}

	return hexStringPairs
}

func getDefaultCustomWords() []string {
	return []string{
		"angry",
		"capitol",
		"cappuccino",
		"coyote",
		"czomo",
		"dubi",
		"electra",
		"login",
		"moustache",
		"pterodakl",
		"smacznego",
		"wacor",
		"wymarzony",
	}
}
