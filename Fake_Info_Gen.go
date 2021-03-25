package main

import (
	"fmt"
	"github.com/Pallinder/go-randomdata"
	"github.com/fatih/color"
	"time"
)

func main() {
	Fake_Info()
}
func Fake_Info() {
	redirect := "www.instagram.com/6o9s"
	color.Red("███████╗░█████╗░██╗░░██╗███████╗  ██╗███╗░░██╗███████╗░█████╗░  ░██████╗░███████╗███╗░░██╗\n")
	color.Red("██╔════╝██╔══██╗██║░██╔╝██╔════╝  ██║████╗░██║██╔════╝██╔══██╗  ██╔════╝░██╔════╝████╗░██║\n")
	color.Red("█████╗░░███████║█████═╝░█████╗░░  ██║██╔██╗██║█████╗░░██║░░██║  ██║░░██╗░█████╗░░██╔██╗██║\n")
	color.Red("██╔══╝░░██╔══██║██╔═██╗░██╔══╝░░  ██║██║╚████║██╔══╝░░██║░░██║  ██║░░╚██╗██╔══╝░░██║╚████║\n")
	color.Red("██║░░░░░██║░░██║██║░╚██╗███████╗  ██║██║░╚███║██║░░░░░╚█████╔╝  ╚██████╔╝███████╗██║░╚███║\n")
	color.Red("╚═╝░░░░░╚═╝░░╚═╝╚═╝░░╚═╝╚══════╝  ╚═╝╚═╝░░╚══╝╚═╝      ╚════╝░  ░╚═════╝░╚══════╝╚═╝░░╚══╝\n")
	color.Red("\n")
	color.Red("██████╗░██╗░░░██╗  ░██████╗███╗░░██╗░░░░░██╗██╗░░██╗           ░█████╗░░█████╗░░█████╗░░██████╗\n")
	color.Red("██╔══██╗╚██╗░██╔╝  ██╔════╝████╗░██║░░░░░██║╚██╗██╔╝           ██╔═══╝░██╔══██╗██╔══██╗██╔════╝\n")
	color.Red("██████╦╝░╚████╔╝░  ╚█████╗░██╔██╗██║░░░░░██║░╚███╔╝░  █████╗   ██████╗░██║░░██║╚██████║╚█████╗░\n")
	color.Red("██╔══██╗░░╚██╔╝░░  ░╚═══██╗██║╚████║██╗░░██║░██╔██╗░  ╚════╝   ██╔══██╗██║░░██║░╚═══██║░╚═══██╗\n")
	color.Red("██████╦╝░░░██║░░░  ██████╔╝██║░╚███║╚█████╔╝██╔╝╚██╗           ╚█████╔╝╚█████╔╝░█████╔╝██████╔╝\n")
	color.Red("╚═════╝░░░░╚═╝░░░  ╚═════╝░╚═╝░░╚══╝░╚════╝░╚═╝░░╚═╝           ░╚════╝░░╚════╝░░╚════╝░╚═════╝░\n")
	color.Red("\n")
	color.Red("\n")
	print("Press Any key to Start...")
	fmt.Scanln()
	print("\nName : ", randomdata.FirstName(randomdata.Male), " ")
	print(randomdata.LastName(), "\n")
	time.Sleep(900000000 * 2)
	print("\nAddress : ", randomdata.Address(), "\n")
	time.Sleep(900000000 * 2)
	print("\nDate : ", randomdata.FullDateInRange(), "\n")
	time.Sleep(900000000 * 2)
	print("\nIp : ", randomdata.IpV4Address(), "\n")
	time.Sleep(900000000 * 2)
	print("\nUser Agent : ", randomdata.UserAgentString(), "\n")
	time.Sleep(500000000)
	print("\nProgrammed")
	time.Sleep(500000000)
	print(" By")
	time.Sleep(500000000)
	print(" SNjX , iG : ", redirect)
	time.Sleep(500000000)
	print("\nGood ")
	time.Sleep(900000000)
	print(" Bye (: ")
	time.Sleep(500000000 * 10)
	print("\n")
	print("ress Any Key To Exit")
	fmt.Scanln()
}
