package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func get_host_name() string { // get host name from /etc/hostname
	host_file, err := os.Open("/etc/hostname")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(0)
	}
	host_info := make([]byte, 1024)
	host_file.Read(host_info)
	host_file.Close()
	host_file_str := strings.Replace(string(host_info), "\x00", "", -1)
	host_file_str = strings.Replace(host_file_str, "\n", "", -1)
	return strings.TrimSpace(string(host_file_str))
}

func get_cpu_info(value string) string { // get cpu info from /proc/cpuinfo
	// get cpu_info
	cpu_file, err := os.Open("/proc/cpuinfo")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(0)
	}
	// read the cpu_info
	cpu_info := make([]byte, 1024)
	cpu_file.Read(cpu_info)
	cpu_file.Close()
	// split the cpu_info by \n into list
	cpu_list := strings.Split(string(cpu_info), "\n")
	// make the cpu_list into variable with key and value
	cpu_tuples := make(map[string]string)
	for _, v := range cpu_list {
		if strings.Contains(v, ":") {
			// parsing
			kv := strings.Split(v, ":")
			kv[0] = strings.TrimSpace(kv[0])
			kv[1] = strings.TrimSpace(kv[1])
			cpu_tuples[kv[0]] = kv[1]
		}
	}
	return cpu_tuples[value]
}

func get_cpu_stat(
	value string, value2 int) string { // get cpu stat from /proc/stat
	// get cpu_stat
	stat_file, err := os.Open("/proc/stat")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(0)
	}
	// read the stat_file
	stat_info := make([]byte, 1024)
	stat_file.Read(stat_info)
	stat_file.Close()
	// split the stat_info by \n into list
	stat_list := strings.Split(string(stat_info), "\n")
	// make the stat_list into variable with key and value
	stat_tuples := make(map[string]string)
	for _, v := range stat_list {
		// parsing
		kv := strings.Split(v, " ")
		k1 := strings.TrimSpace(kv[0])
		k2 := strings.Split(strings.TrimSpace(strings.Join(kv[1:], " ")), " ")
		stat_tuples[k1] = k2[value2]
	}
	return stat_tuples[value]
}

func get_mem_info(value string) string { //get memory info from /proc/meminfo
	// get mem_info
	mem_file, err := os.Open("/proc/meminfo")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(0)
	}
	// read the mem_info
	mem_info := make([]byte, 1024)
	mem_file.Read(mem_info)
	mem_file.Close()
	// split the mem_info by \n into list
	mem_list := strings.Split(string(mem_info), "\n")
	// make the mem_list into variable with key and value
	mem_tuples := make(map[string]string)
	for _, v := range mem_list {
		if strings.Contains(v, ":") {
			// parsing
			kv := strings.Split(v, ":")
			kv[0] = strings.TrimSpace(kv[0])
			kv[1] = strings.TrimSpace(kv[1])
			mem_tuples[strings.Replace(kv[0], " ", "", -1)] = strings.Replace(
				kv[1], " ", "", -1)
		}
	}
	return strings.TrimSpace(strings.Replace(mem_tuples[value], "kB", "", -1))
}

func get_vga_series() string { // get vga series from lspci
	cmd := exec.Command("lspci")
	// pipe the output of lspci to a variable
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(0)
	}
	// find line that contains "VGA compatible controller"
	vga_info := regexp.MustCompile(`\[.*?/`).FindAll(regexp.MustCompile(`VGA.*`).Find(out), -1)
	// return strings after "[" and before "]"
	t1 := strings.TrimSpace(strings.Trim(string(strings.Trim(string(vga_info[0]), "[")), "/"))
	t2 := strings.TrimSpace(strings.Trim(string(strings.Trim(string(vga_info[1]), "[")), "/"))
	return t1 + " " + t2
}

func get_monitor_size() string { // get monitor size
	cmd := exec.Command("xdpyinfo")
	// pipe the output of xdpyinfo to a variable
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(0)
	}
	// find line that contains "dimensions:"
	monitor_info := regexp.MustCompile(`dimensions:.*?x.*? `).FindAll(out, -1)
	monitor_info_trim := strings.TrimPrefix(
		string(monitor_info[0]),
		"dimensions: ")
	// return strings after ":" and before "x"
	return strings.TrimSpace(monitor_info_trim)
}

func get_distro() string { // get distro from /etc/os-release
	distro, err := os.Open("/etc/os-release")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(0)
	}
	distro_info := make([]byte, 1024)
	distro.Read(distro_info)
	distro.Close()
	distro_list := strings.Split(string(distro_info), "\n")
	distro_tuples := make(map[string]string)
	for _, v := range distro_list {
		if strings.Contains(v, "=") {
			kv := strings.Split(v, "=")
			kv[0] = strings.TrimSpace(kv[0])
			kv[1] = strings.TrimSpace(kv[1])
			distro_tuples[kv[0]] = kv[1]
		}
	}
	return strings.Trim(distro_tuples["PRETTY_NAME"], "\"")
}

func get_computer_brand() string {
	// get computer brand from /sys/devices/virtual/dmi/id/
	vendor, err1 := os.Open("/sys/devices/virtual/dmi/id/sys_vendor")
	name, err2 := os.Open("/sys/devices/virtual/dmi/id/product_name")
	if err1 != nil || err2 != nil {
		fmt.Println("Error:", err1, err2)
		os.Exit(0)
	}
	vendor_info := make([]byte, 1024)
	name_info := make([]byte, 1024)
	vendor.Read(vendor_info)
	name.Read(name_info)
	vendor.Close()
	name.Close()

	return strings.Replace(
		string(vendor_info), "\n", "", -1) + " " + strings.Replace(string(name_info), "\n", "", -1)
}

func get_env(env string) string {
	// get os enviroment
	out := os.Getenv(env)
	// get shell name
	_out := strings.Split(out, "/")

	return _out[len(_out)-1]
}

func get_kernel() string { // get kernel version from /proc/version
	kernel_file, err := os.Open("/proc/version")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(0)
	}
	kernel_info := make([]byte, 1024)
	kernel_file.Read(kernel_info)
	kernel_file.Close()
	kernel_list := strings.Split(string(kernel_info), " ")
	return strings.TrimSpace(kernel_list[2])
}

func get_teminal_color_palette() string {
	color1 := "\x1b[0;29m  \x1b[0m"
	color2 := "\x1b[0;31m  \x1b[0m"
	color3 := "\x1b[0;32m  \x1b[0m"
	color4 := "\x1b[0;33m  \x1b[0m"
	color5 := "\x1b[0;34m  \x1b[0m"
	color6 := "\x1b[0;35m  \x1b[0m"
	color7 := "\x1b[0;36m  \x1b[0m"
	color8 := "\x1b[0;37m  \x1b[0m"

	return color1 + " " + color2 + " " + color3 + " " + color4 + " " + color5 + " " + color6 + " " + color7 + " " + color8
}
