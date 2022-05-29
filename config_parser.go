package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

var (
	home_env = os.Getenv("HOME")
	local_image = pwd()+"/sample.jpg"
	remote_image = home_env+"/.config/gofetch/sample.jpg"

)

func deploy_sample_image() {
	copy_file(local_image, remote_image)
}
func check_config_file() string{
	file, err := os.Open(home_env+"/.config/gofetch/config.json")
	pre_config := "# config file for gofetch\n"+
	"image_location: "+home_env+"/.config/gofetch/sample.jpg"
	if err != nil {
		// create config file
		file, err := os.Create(home_env+"/.config/gofetch/config.json")
		// init config file
		if err != nil {
			// create folder
			os.Mkdir(home_env+"/.config/gofetch", 0755)
			// create config file
			file, err := os.Create(home_env+"/.config/gofetch/config.json")
			if err != nil {
				fmt.Println("Error:", err)
				os.Exit(0)
			}
			file.WriteString(pre_config)
			deploy_sample_image()

		}
		file.WriteString(pre_config)
		deploy_sample_image()
	}
	data := make([]byte, 1024)
	file.Read(data)
	defer file.Close()
	data_str := strings.Replace(string(data), "\x00", "", -1)
	return data_str
}
func get_image_location() string {
	config_file := check_config_file()
	re := regexp.MustCompile(`image_location: (.*)`)
	image_location := strings.Replace(re.FindString(config_file), "image_location: ", "", -1)
	image_location = strings.Replace(image_location, "\n", "", -1)
	image_location = strings.TrimSpace(image_location)
	return image_location
}