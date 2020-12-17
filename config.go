package pikselkapcio

const (
	//TextGenerationRandom should be used for pseudorandom alphanumeric string for code (it is a default value)
	TextGenerationRandom = 1
	//TextGenerationCustomWords should be used to generate code from provided custom words list
	TextGenerationCustomWords = 2
	//DefaultSessionKey is default key which is used to store code in users session
	DefaultSessionKey = "_wl_kapcio"
)

//Config strucure for whole package
type Config struct {
	Scale               int
	TextGenerationMode  int
	RandomTextLength    int
	CustomWords         []string
	ColorHexStringPairs []HexStringPair
	SessionKey          string
}

func getDefaultConfig() Config {
	return Config{
		Scale:               5,
		TextGenerationMode:  TextGenerationRandom,
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

	if customConfig.TextGenerationMode != 0 {
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

	if customConfig.SessionKey != "" {
		config.SessionKey = customConfig.SessionKey
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
