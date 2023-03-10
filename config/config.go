package config

var Conf appConfig

type appConfig struct {
	URL string `json:"url"`
}

func LoadConfig() error {
	// file, err := os.ReadFile("config.json")
	// fmt.Println(file)
	// if err != nil {
	// 	return err
	// }

	// if err != nil {
	// 	return err
	// }
	// json.Unmarshal(file, &Conf)
	Conf.URL = "www.joyreactor.cc/"

	return nil
}
