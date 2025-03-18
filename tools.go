package main

import (
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

//##############################################################################
//#                               check related                                #
//##############################################################################

func check_requirement() { //check if the requirement is met
	// check chafa
	_, err := exec.LookPath("chafa")
	if err != nil {
		fmt.Println("Error: chafa is not installed")
		os.Exit(0)
	}
}

func check_os() { // check if the os is linux
	if runtime.GOOS != "linux" {
		fmt.Println("This is not a Linux system")
		os.Exit(0)
	}
}

//##############################################################################
//#                              display related                               #
//##############################################################################

func clear_screen() { // clear the screen but only the last output
	os.Stdout.Write([]byte("\033[0;0H"))
}

func clear_screen_full() { // clear the screen
	os.Stdout.Write([]byte("\033[H\033[2J"))
}

func display_image(
	location string,
	width int,
	height int) []byte { // display image using chafa
	// convert width and height to string
	width_str := strconv.Itoa(width)
	height_str := strconv.Itoa(height)
	// display an image using chafa
	cmd := exec.Command(
		"chafa", location, fmt.Sprintf("--size=%sx%s", width_str, height_str))
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(0)
	}
	return out
}

func bold(value string) string { // bold text
	return fmt.Sprintf("\033[1m%s\033[0m", value)
}

func templete(_len int) string { // template for print
	var _templete string
	for i := 0; i < _len; i++ {
		_templete += " "
	}
	return _templete
}
func coloring(text string, _hex string) string {
	_hex = strings.Trim(_hex, "#")
	var (
		r string
		g string
		b string
	)
	if len(_hex) == 6 {
		r = _hex[0:2]
		g = _hex[2:4]
		b = _hex[4:6]
	} else {
		return text
	}
	rgb, err := hex.DecodeString(r + g + b)
	if err != nil {
		return text
	}
	return fmt.Sprintf(
		"\x1b[38;2;%d;%d;%dm%s\x1b[0m", rgb[0], rgb[1], rgb[2], text)
}
func make_header(text string, _len int) string {
	var foo string
	for i := 0; i < _len; i++ {
		if i == int((_len/2)-(len(text)/2)) {
			foo += text
		}
		foo += "─"
		if i == _len-len(text) {
			break
		}
	}
	return foo
}
func color_templete(txt string, value int) string { // color templete
	switch value {
	case 1:
		return fmt.Sprintf("\x1b[0;29m%s\x1b[0m", txt)
	case 2:
		return fmt.Sprintf("\033[0;31m%s\033[0m", txt)
	case 3:
		return fmt.Sprintf("\033[0;32m%s\033[0m", txt)
	case 4:
		return fmt.Sprintf("\033[0;33m%s\033[0m", txt)
	case 5:
		return fmt.Sprintf("\033[0;34m%s\033[0m", txt)
	case 6:
		return fmt.Sprintf("\033[0;35m%s\033[0m", txt)
	case 7:
		return fmt.Sprintf("\033[0;36m%s\033[0m", txt)
	case 8:
		return fmt.Sprintf("\033[0;37m%s\033[0m", txt)
	}
	return txt
}

//##############################################################################
//#                             convertion related                             #
//##############################################################################

func kb_to_gb(kb_str string) string { // convert kb to gb
	kb_str = strings.Replace(kb_str, "kB", "", -1)
	kb, err := strconv.Atoi(kb_str)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(0)
	}
	gb := float64(kb) / 1024 / 1024
	return fmt.Sprintf("%.2f", gb)
}

func percent(value string, value2 string) string { // calculate percent
	// convert value and value2 to float64
	value_float, err := strconv.ParseFloat(value, 64)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(0)
	}
	value2_float, err := strconv.ParseFloat(value2, 64)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(0)
	}
	// calculate percent
	percent := value_float / value2_float * 100
	return fmt.Sprintf("%.2f%%", percent)
}

func str2int(str string) int { // convert string to int
	i, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(0)
	}
	return i
}

func int2str(value int) string { // convert int to string
	return strconv.Itoa(value)
}

//##############################################################################
//#                                open related                                #
//##############################################################################

func pwd() string { // open file and read
	get_pwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(0)
	}
	return get_pwd
}

func copy_file(src string, dst string) error { // copy file
	// open file
	src_file, err := os.Open(src)
	if err != nil {
		return err
	}
	defer src_file.Close()
	// open file
	dst_file, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer dst_file.Close()
	// copy file
	_, err = io.Copy(dst_file, src_file)
	if err != nil {
		return err
	}
	return nil
}
