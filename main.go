package main

import (
	"fmt"
)

func main() {
	//################################################################################
	//#                                 preparation                                  #            
	//################################################################################

	// checking
	check_os()
	check_requirement()

	// bug fixing (temporary)
	bogus_for_bug_fixing := get_image_location()
	fmt.Println(bogus_for_bug_fixing)

	// print the image
	clear_screen_full()
	image_loc := get_image_location()
	fmt.Printf("%s", display_image(
		image_loc,
		35,
		14,
	))
	clear_screen()
	
	//################################################################################
	//#                              prepare variables                               #                  
	//################################################################################
	
	// Hardware Information
	computer_brand 	:= get_computer_brand()
	cpu_info 		:= get_cpu_info("model name")
	memory_total 	:= kb_to_gb(get_mem_info("MemTotal"))
	memory_used 	:= kb_to_gb(
		int2str(
			str2int(get_mem_info("MemTotal")) - 
			str2int(get_mem_info("MemFree")) - 
			str2int(get_mem_info("Buffers")) -
			str2int(get_mem_info("Cached")),
		),
	)
	memory_percent 	:= percent(memory_used, memory_total)
	vga_series 		:= get_vga_series()
	monitor_size 	:= get_monitor_size()
	
	// Software Information
	host_name 		:= get_host_name()
	distro 			:= get_distro()
	kernel_version	:= get_kernel()
	shell_name 		:= get_env("SHELL")
	session			:= get_env("DESKTOP_SESSION")

	// other
	template 		:= templete(37)
	header 			:= make_header(" "+host_name+" ", 39)

	// icon
	icon_computer_brand 		:= color_templete("", 7)
	icon_cpu_info 				:= color_templete("", 7)
	icon_memory_total 			:= color_templete("塞", 7)
	icon_vga_series 			:= color_templete("﬙", 7)
	icon_monitor_size 			:= color_templete("", 7)
	icon_distro 				:= color_templete("", 7)
	icon_kernel_version 		:= color_templete("", 7)
	icon_shell_name 			:= color_templete("", 7)
	icon_session 				:= color_templete("", 7)

	//################################################################################
	//#                              print to terminal                               #                  
	//################################################################################
	
	fmt.Print(template);fmt.Println(bold(coloring(" "+header+" ", "#c3c7d1")))
	fmt.Print(template);fmt.Println(coloring(bold(	"┌───────── Hardware Information ─────────┐"), "828997"))
	fmt.Print(template);fmt.Printf(					" %s  %s\n", 			icon_computer_brand, computer_brand)
	fmt.Print(template);fmt.Printf(					" %s  %s\n", 			icon_cpu_info, cpu_info)
	fmt.Print(template);fmt.Printf(					" %s %s / %s (%s)\n", 	icon_memory_total, memory_used, memory_total, memory_percent)
	fmt.Print(template);fmt.Printf(					" %s  %s\n", 			icon_vga_series, vga_series)
	fmt.Print(template);fmt.Printf(					" %s  %s\n", 			icon_monitor_size, monitor_size)
	fmt.Print(template);fmt.Println(coloring(bold(	"├──────── Software Information ──────────┤"), "828997"))
	fmt.Print(template);fmt.Printf(					" %s  %s\n", 			icon_distro, distro)
	fmt.Print(template);fmt.Printf(					" %s  %s\n", 			icon_kernel_version, kernel_version)
	fmt.Print(template);fmt.Printf(					" %s  %s\n", 			icon_shell_name, shell_name)
	fmt.Print(template);fmt.Printf(					" %s  %s\n", 			icon_session, session)
	fmt.Print(template);fmt.Println(coloring(bold(	"└────────────────────────────────────────┘"), "828997"))
	fmt.Print(template);fmt.Println(				"      "+get_teminal_color_palette())
}
