package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	//for {
	//fmt.Println("hello")
	//fmt.Println("there")
	if len(os.Args) > 1 {
		fmt.Println("you're stupid", os.Args[1])
	}
	stringo := "hello"
	//byto := byte(stringo)
	os.Stdout.WriteString(stringo)
	fmt.Println("!=!=")
	//consolereader := bufio.NewScanner(os.Stdin)

	//input, err := consolereader.ReadString('\n') // this will prompt the user for input

	//if err != nil {
	//     fmt.Println(err)
	//     os.Exit(1)
	// }

	//fmt.Print(input)
	//fmt.Println(input)
	//return

	//}
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		var stdin []string
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			fmt.Println("part", stdin)
			stdin = append(stdin, scanner.Text())
			//stdin = append(stdin, byte(64))
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("stdin = %s\n", stdin)
		fmt.Println(len(stdin))
	}
	/*else {
	    fmt.Println("Enter your name")

	    var name string
	    fmt.Scanf("%s", &name)
	    fmt.Printf("name = %s\n", name)
	}*/

}
