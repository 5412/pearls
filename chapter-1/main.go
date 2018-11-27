package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"unsafe"
)

func main() {
	rest := make([]uint64, 10e7/64+1)
	restt  := make([]int, 0,  10e7)
	inputFile, err := os.Open("input-file")
	defer func() {
		err := inputFile.Close()
		if err != nil {
			log.Println(err)
		}
	}()
	if err != nil {
		log.Println(err)
	}

	stat, err := inputFile.Stat()

	if err != nil {
		log.Println(err)
	}

	size := stat.Size()

	log.Println("input-file's size is", size)

	buf := bufio.NewReader(inputFile)
	outFile, err := os.OpenFile("out-file", 1, os.ModeAppend)
	bufw := bufio.NewWriter(outFile)
	defer func() {
		err := outFile.Close()
		if err != nil {
			log.Println(err)
		}
	}()
	if err != nil {
		log.Println(err)
	}

	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)

		if err != nil {
			log.Println(err)
			if err == io.EOF {
				log.Println("File read Ok")
				break
			} else {
				log.Println("read file error")
				break
			}
		}

		num, err := strconv.Atoi(line)

		if num > 0 {
			index := num / 64
			realnum := num % 64
			if realnum == 0 {
				index--
				realnum = 64
			}

			rest[index] = rest[index] | uint64(1<<uint64(64-realnum))
			//restt = append(restt, num)
		}
	}

	log.Println(unsafe.Sizeof(rest))
	log.Println(unsafe.Sizeof(restt))
	//sort.Ints(restt)

	//for _,v := range restt {
	//	_, err := outFile.WriteString(strconv.Itoa(v) + "\n")
	//	if err != nil {
	//		log.Println(err)
	//	}
	//}

	for i, v := range rest {
		for k := 63; k >= 0; k-- {
			//fmt.Printf("%b \n", v)
			if 1<<uint64(k)&v > 0 {
				num := uint64(i)*64 + uint64(64-k)
				//log.Println(int(num))
				_, err := bufw.WriteString(strconv.Itoa(int(num)) + "\n")
				if err != nil {
					log.Println(err)
				}
			}
		}
	}
}
