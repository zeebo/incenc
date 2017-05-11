package main

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"errors"
	"fmt"
	"incenc"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"gopkg.in/spacemonkeygo/monkit.v2"
)

var mon = monkit.Package()

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	monkit.Stats(func(name string, val float64) {
		if strings.Contains(name, " times ") {
			fmt.Println(name, "\t", time.Duration(val*float64(time.Second)))
		} else {
			fmt.Println(name, "\t", val)
		}
	})
	fmt.Println(incenc.Total)
}

func run() (err error) {
	switch os.Args[1] {
	case "compress", "c":
		return compress()
	case "decompress", "d":
		return decompress()
	}
	return errors.New("unknown command. usage: incenc compress|decompress")
}

func compress() (err error) {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return err
	}
	func() {
		defer mon.TaskNamed("gzip")(nil)(nil)
		gzip.NewWriter(ioutil.Discard).Write(data)
	}()

	e := incenc.NewEncoder()
	scanner := bufio.NewScanner(bytes.NewReader(data))
	for scanner.Scan() {
		e.Add(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	_, err = os.Stdout.Write(e.Bytes())
	return err
}

func decompress() (err error) {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return err
	}
	// fmt.Println(hex.Dump(data))

	d := incenc.NewDecoder(data)

	d.Iterate(nil, func(line []byte) {
		os.Stdout.Write(line)
		os.Stdout.WriteString("\n")
	})
	return nil
}
