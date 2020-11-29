package main

import (
	"flag"
	"github.com/avct/uasurfer"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/yaml.v2"
)

func main() {
	path := flag.String("config", "", "Path for the config file")
	flag.Parse()

	config := parseConfig(*path)

	router := httprouter.New()
	router.GET("/:direction", handler(config))

	log.Println("Server is starting at port 80!")
	if err := http.ListenAndServe(":80", router); err != nil {
		log.Fatal(err)
	}
}

func handler(config *Config) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		direction := p.ByName("direction")
		log.Println("Redirecting for:", direction)

		userAgent := uasurfer.Parse(r.UserAgent())
		log.Println(r.UserAgent(), "=>", userAgent.OS.Name.String())

		switch userAgent.OS.Name {
		case uasurfer.OSiOS:
			log.Println("Redirecting to", config.Directions[direction].IOS)
			http.Redirect(w, r, config.Directions[direction].IOS, 301)

		case uasurfer.OSAndroid:
			log.Println("Redirecting to", config.Directions[direction].Android)
			http.Redirect(w, r, config.Directions[direction].Android, 301)

		default:
			log.Println("No possible direction found!")
			if _, err := w.Write([]byte("No possible direction found!")); err != nil {
				log.Println(err.Error())
			}
		}
	}
}

type Config struct {
	Directions map[string]struct {
		IOS     string `yaml:"ios"`
		Android string `yaml:"android"`
	} `yaml:"directions"`
}

func parseConfig(path string) *Config {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	var c = &Config{}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	defer func() {
		directions := make([]string, 0, len(c.Directions))
		for k := range c.Directions {
			directions = append(directions, k)
		}
		log.Printf("Loaded config for: %s", strings.Join(directions, ", "))
	}()

	return c
}
