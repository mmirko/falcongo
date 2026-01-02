package main

import (
	"github.com/mmirko/falcongo/pkg/p/go9p/p"
	"github.com/mmirko/falcongo/pkg/p/go9p/p/srv"
	"flag"
	"fmt"
	"github.com/tarm/serial"
	"log"
	"strconv"
	"strings"
	"time"
)

const (
	INITTIME  = 5000
	CICLETIME = 1000
)

type falcstatus struct {
	Ready    bool
	Temp00   int
	Hum00    int
	Temp01   int
	Hum01    int
	Dc00     float64
	Ac00     float64
	Light00  float64
	Switch00 bool
	Switch01 bool
	Switch02 bool
	Switch03 bool
	Switch04 bool
	Switch05 bool
	Switch06 bool
	Switch07 bool
}

type Temp00 struct {
	srv.File
}

type Hum00 struct {
	srv.File
}

type Temp01 struct {
	srv.File
}

type Hum01 struct {
	srv.File
}

type Ac00 struct {
	srv.File
}

type Dc00 struct {
	srv.File
}

type Light00 struct {
	srv.File
}

type Switch00 struct {
	srv.File
}

type Switch01 struct {
	srv.File
}

type Switch02 struct {
	srv.File
}

type Switch03 struct {
	srv.File
}

type Switch04 struct {
	srv.File
}

type Switch05 struct {
	srv.File
}

type Switch06 struct {
	srv.File
}

type Switch07 struct {
	srv.File
}

var q chan string

var c *serial.Config
var s *serial.Port
var st *falcstatus

var addr = flag.String("addr", ":9998", "network address")
var debug = flag.Bool("d", false, "print debug messages")

var serv *srv.Fsrv
var root *srv.File

func (ctl *Temp00) Read(fid *srv.FFid, buf []byte, offset uint64) (int, error) {
	if offset > uint64(0) {
		return 0, nil
	}

	name := strconv.Itoa(st.Temp00) + "\n"

	b := []byte(name)

	copy(buf, b)
	return len(b), nil
}

func (ctl *Temp00) Write(fid *srv.FFid, data []byte, offset uint64) (int, error) {
	return len(data), nil
}

func (ctl *Temp01) Read(fid *srv.FFid, buf []byte, offset uint64) (int, error) {
	if offset > uint64(0) {
		return 0, nil
	}

	name := strconv.Itoa(st.Temp01) + "\n"

	b := []byte(name)

	copy(buf, b)
	return len(b), nil
}

func (ctl *Temp01) Write(fid *srv.FFid, data []byte, offset uint64) (int, error) {
	return len(data), nil
}

func (ctl *Hum00) Read(fid *srv.FFid, buf []byte, offset uint64) (int, error) {
	if offset > uint64(0) {
		return 0, nil
	}

	name := strconv.Itoa(st.Hum00) + "\n"

	b := []byte(name)

	copy(buf, b)
	return len(b), nil
}

func (ctl *Hum00) Write(fid *srv.FFid, data []byte, offset uint64) (int, error) {
	return len(data), nil
}

func (ctl *Hum01) Read(fid *srv.FFid, buf []byte, offset uint64) (int, error) {
	if offset > uint64(0) {
		return 0, nil
	}

	name := strconv.Itoa(st.Hum01) + "\n"

	b := []byte(name)

	copy(buf, b)
	return len(b), nil
}

func (ctl *Hum01) Write(fid *srv.FFid, data []byte, offset uint64) (int, error) {
	return len(data), nil
}

func (ctl *Ac00) Read(fid *srv.FFid, buf []byte, offset uint64) (int, error) {
	if offset > uint64(0) {
		return 0, nil
	}

	name := strconv.FormatFloat(st.Ac00, 'f', 6, 64) + "\n"

	b := []byte(name)

	copy(buf, b)
	return len(b), nil
}

func (ctl *Ac00) Write(fid *srv.FFid, data []byte, offset uint64) (int, error) {
	return len(data), nil
}

func (ctl *Dc00) Read(fid *srv.FFid, buf []byte, offset uint64) (int, error) {
	if offset > uint64(0) {
		return 0, nil
	}

	name := strconv.FormatFloat(st.Dc00, 'f', 6, 64) + "\n"

	b := []byte(name)

	copy(buf, b)
	return len(b), nil
}

func (ctl *Dc00) Write(fid *srv.FFid, data []byte, offset uint64) (int, error) {
	return len(data), nil
}

func (ctl *Light00) Read(fid *srv.FFid, buf []byte, offset uint64) (int, error) {
	if offset > uint64(0) {
		return 0, nil
	}

	name := strconv.FormatFloat(st.Light00, 'f', 6, 64) + "\n"

	b := []byte(name)

	copy(buf, b)
	return len(b), nil
}

func (ctl *Light00) Write(fid *srv.FFid, data []byte, offset uint64) (int, error) {
	return len(data), nil
}

func (ctl *Switch00) Read(fid *srv.FFid, buf []byte, offset uint64) (int, error) {
	if offset > uint64(0) {
		return 0, nil
	}

	var name string

	if st.Switch00 == true {
		name = "true\n"
	} else {
		name = "false\n"
	}

	b := []byte(name)

	copy(buf, b)
	return len(b), nil
}

func (ctl *Switch00) Write(fid *srv.FFid, data []byte, offset uint64) (int, error) {
	if offset > uint64(0) {
		return 0, nil
	}

	if len(string(data)) >= 2 {
		if string(data)[0:2] == "on" {
			q <- "on switch00"
			q <- "get switch00"
		} else if string(data)[0:2] == "of" {
			q <- "off switch00"
			q <- "get switch00"
		} else if string(data)[0:2] == "sw" {
			q <- "sw switch00"
			q <- "get switch00"
		}
	}
	return len(data), nil
}

func (ctl *Switch01) Read(fid *srv.FFid, buf []byte, offset uint64) (int, error) {
	if offset > uint64(0) {
		return 0, nil
	}

	var name string

	if st.Switch01 == true {
		name = "true\n"
	} else {
		name = "false\n"
	}

	b := []byte(name)

	copy(buf, b)
	return len(b), nil
}

func (ctl *Switch01) Write(fid *srv.FFid, data []byte, offset uint64) (int, error) {
	if offset > uint64(0) {
		return 0, nil
	}

	if len(string(data)) >= 2 {
		if string(data)[0:2] == "on" {
			q <- "on switch01"
			q <- "get switch01"
		} else if string(data)[0:2] == "of" {
			q <- "off switch01"
			q <- "get switch01"
		} else if string(data)[0:2] == "sw" {
			q <- "sw switch01"
			q <- "get switch01"
		}
	}
	return len(data), nil
}

func (ctl *Switch02) Read(fid *srv.FFid, buf []byte, offset uint64) (int, error) {
	if offset > uint64(0) {
		return 0, nil
	}

	var name string

	if st.Switch02 == true {
		name = "true\n"
	} else {
		name = "false\n"
	}

	b := []byte(name)

	copy(buf, b)
	return len(b), nil
}

func (ctl *Switch02) Write(fid *srv.FFid, data []byte, offset uint64) (int, error) {
	if offset > uint64(0) {
		return 0, nil
	}

	if len(string(data)) >= 2 {
		if string(data)[0:2] == "on" {
			q <- "on switch02"
			q <- "get switch02"
		} else if string(data)[0:2] == "of" {
			q <- "off switch02"
			q <- "get switch02"
		} else if string(data)[0:2] == "sw" {
			q <- "sw switch02"
			q <- "get switch02"
		}
	}
	return len(data), nil
}

func (ctl *Switch03) Read(fid *srv.FFid, buf []byte, offset uint64) (int, error) {
	if offset > uint64(0) {
		return 0, nil
	}

	var name string

	if st.Switch03 == true {
		name = "true\n"
	} else {
		name = "false\n"
	}

	b := []byte(name)

	copy(buf, b)
	return len(b), nil
}

func (ctl *Switch03) Write(fid *srv.FFid, data []byte, offset uint64) (int, error) {
	if offset > uint64(0) {
		return 0, nil
	}

	if len(string(data)) >= 2 {
		if string(data)[0:2] == "on" {
			q <- "on switch03"
			q <- "get switch03"
		} else if string(data)[0:2] == "of" {
			q <- "off switch03"
			q <- "get switch03"
		} else if string(data)[0:2] == "sw" {
			q <- "sw switch03"
			q <- "get switch03"
		}
	}
	return len(data), nil
}

func (ctl *Switch04) Read(fid *srv.FFid, buf []byte, offset uint64) (int, error) {
	if offset > uint64(0) {
		return 0, nil
	}

	var name string

	if st.Switch04 == true {
		name = "true\n"
	} else {
		name = "false\n"
	}

	b := []byte(name)

	copy(buf, b)
	return len(b), nil
}

func (ctl *Switch04) Write(fid *srv.FFid, data []byte, offset uint64) (int, error) {
	if offset > uint64(0) {
		return 0, nil
	}

	if len(string(data)) >= 2 {
		if string(data)[0:2] == "on" {
			q <- "on switch04"
			q <- "get switch04"
		} else if string(data)[0:2] == "of" {
			q <- "off switch04"
			q <- "get switch04"
		} else if string(data)[0:2] == "sw" {
			q <- "sw switch04"
			q <- "get switch04"
		}
	}
	return len(data), nil
}

func (ctl *Switch05) Read(fid *srv.FFid, buf []byte, offset uint64) (int, error) {
	if offset > uint64(0) {
		return 0, nil
	}

	var name string

	if st.Switch05 == true {
		name = "true\n"
	} else {
		name = "false\n"
	}

	b := []byte(name)

	copy(buf, b)
	return len(b), nil
}

func (ctl *Switch05) Write(fid *srv.FFid, data []byte, offset uint64) (int, error) {
	if offset > uint64(0) {
		return 0, nil
	}

	if len(string(data)) >= 2 {
		if string(data)[0:2] == "on" {
			q <- "on switch05"
			q <- "get switch05"
		} else if string(data)[0:2] == "of" {
			q <- "off switch05"
			q <- "get switch05"
		} else if string(data)[0:2] == "sw" {
			q <- "sw switch05"
			q <- "get switch05"
		}
	}
	return len(data), nil
}

func (ctl *Switch06) Read(fid *srv.FFid, buf []byte, offset uint64) (int, error) {
	if offset > uint64(0) {
		return 0, nil
	}

	var name string

	if st.Switch06 == true {
		name = "true\n"
	} else {
		name = "false\n"
	}

	b := []byte(name)

	copy(buf, b)
	return len(b), nil
}

func (ctl *Switch06) Write(fid *srv.FFid, data []byte, offset uint64) (int, error) {
	if offset > uint64(0) {
		return 0, nil
	}

	if len(string(data)) >= 2 {
		if string(data)[0:2] == "on" {
			q <- "on switch06"
			q <- "get switch06"
		} else if string(data)[0:2] == "of" {
			q <- "off switch06"
			q <- "get switch06"
		} else if string(data)[0:2] == "sw" {
			q <- "sw switch06"
			q <- "get switch06"
		}
	}
	return len(data), nil
}

func (ctl *Switch07) Read(fid *srv.FFid, buf []byte, offset uint64) (int, error) {
	if offset > uint64(0) {
		return 0, nil
	}

	var name string

	if st.Switch07 == true {
		name = "true\n"
	} else {
		name = "false\n"
	}

	b := []byte(name)

	copy(buf, b)
	return len(b), nil
}

func (ctl *Switch07) Write(fid *srv.FFid, data []byte, offset uint64) (int, error) {
	if offset > uint64(0) {
		return 0, nil
	}

	if len(string(data)) >= 2 {
		if string(data)[0:2] == "on" {
			q <- "on switch07"
			q <- "get switch07"
		} else if string(data)[0:2] == "of" {
			q <- "off switch07"
			q <- "get switch07"
		} else if string(data)[0:2] == "sw" {
			q <- "sw switch07"
			q <- "get switch07"
		}
	}
	return len(data), nil
}

func responses(respc chan string) {
	for {
		resp := <-respc
		//fmt.Println("resp: ", resp)
		if resp == "Ready" {
			st.Ready = true
		} else {
			keyval := strings.Split(resp, ":")
			if len(keyval) == 2 {
				// TODO Status check
				switch keyval[0] {
				case "temp00":
					st.Temp00, _ = strconv.Atoi(keyval[1])
				case "hum00":
					st.Hum00, _ = strconv.Atoi(keyval[1])
				case "temp01":
					st.Temp01, _ = strconv.Atoi(keyval[1])
				case "hum01":
					st.Hum01, _ = strconv.Atoi(keyval[1])
				case "dc00":
					st.Dc00, _ = strconv.ParseFloat(keyval[1], 64)
				case "ac00":
					st.Ac00, _ = strconv.ParseFloat(keyval[1], 64)
				case "light00":
					st.Light00, _ = strconv.ParseFloat(keyval[1], 64)
				case "switch00":
					if keyval[1] == "0" {
						st.Switch00 = false
					} else {
						st.Switch00 = true
					}
				case "switch01":
					if keyval[1] == "0" {
						st.Switch01 = false
					} else {
						st.Switch01 = true
					}
				case "switch02":
					if keyval[1] == "0" {
						st.Switch02 = false
					} else {
						st.Switch02 = true
					}
				case "switch03":
					if keyval[1] == "0" {
						st.Switch03 = false
					} else {
						st.Switch03 = true
					}
				case "switch04":
					if keyval[1] == "0" {
						st.Switch04 = false
					} else {
						st.Switch04 = true
					}
				case "switch05":
					if keyval[1] == "0" {
						st.Switch05 = false
					} else {
						st.Switch05 = true
					}
				case "switch06":
					if keyval[1] == "0" {
						st.Switch06 = false
					} else {
						st.Switch06 = true
					}
				case "switch07":
					if keyval[1] == "0" {
						st.Switch07 = false
					} else {
						st.Switch07 = true
					}
				}
			}
		}
		//fmt.Println(resp, st)
	}
}

func queries() {
	time.Sleep(INITTIME * time.Millisecond)
	command_cicle := []string{"get temp00", "get hum00", "get temp01", "get hum01", "get dc00", "get ac00", "get light00", "get switch00", "get switch01", "get switch02", "get switch03", "get switch04", "get switch05", "get switch06", "get switch07"}
	i := 0
	for {
		q <- command_cicle[i]

		time.Sleep(CICLETIME * time.Millisecond)

		i++
		if i == len(command_cicle) {
			i = 0
		}
	}
}

func writer() {
	for {
		command := <-q
		//fmt.Println("comm: ", command)
		_, err := s.Write([]byte(command + "\n\r"))
		if err != nil {
			log.Fatal(err)
		}

	}
}

func read_serial(respc chan string) {
	var n int
	var err error

	buf := make([]byte, 128)
	partial := ""
	for {
		n, err = s.Read(buf)
		if err != nil {
			//log.Fatal(err)
		}

		for _, ch := range buf[:n] {
			if ch == 13 {
				respc <- partial
				partial = ""
			} else if ch >= 32 && ch <= 126 {
				partial = partial + string(ch)
			}
		}
	}
}

func init() {
	var err error

	st = new(falcstatus)

	c = &serial.Config{Name: "/dev/ttyUSB0", Baud: 9600}
	s, err = serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	var err error
	var temp00 *Temp00
	var hum00 *Hum00
	var temp01 *Temp01
	var hum01 *Hum01
	var ac00 *Ac00
	var dc00 *Dc00
	var light00 *Light00
	var switch00 *Switch00
	var switch01 *Switch01
	var switch02 *Switch02
	var switch03 *Switch03
	var switch04 *Switch04
	var switch05 *Switch05
	var switch06 *Switch06
	var switch07 *Switch07

	flag.Parse()

	respc := make(chan string)
	q = make(chan string)

	go responses(respc)

	go queries()

	go read_serial(respc)

	go writer()

	root = new(srv.File)
	err = root.Add(nil, "/", nil, nil, p.DMDIR|0777, nil)
	if err != nil {
		goto error
	}

	temp00 = new(Temp00)
	err = temp00.Add(root, "temp00", nil, nil, 0444, temp00)
	if err != nil {
		goto error
	}

	hum00 = new(Hum00)
	err = hum00.Add(root, "hum00", nil, nil, 0444, hum00)
	if err != nil {
		goto error
	}

	temp01 = new(Temp01)
	err = temp01.Add(root, "temp01", nil, nil, 0444, temp01)
	if err != nil {
		goto error
	}

	hum01 = new(Hum01)
	err = hum01.Add(root, "hum01", nil, nil, 0444, hum01)
	if err != nil {
		goto error
	}

	ac00 = new(Ac00)
	err = ac00.Add(root, "ac00", nil, nil, 0444, ac00)
	if err != nil {
		goto error
	}

	dc00 = new(Dc00)
	err = dc00.Add(root, "dc00", nil, nil, 0444, dc00)
	if err != nil {
		goto error
	}

	light00 = new(Light00)
	err = light00.Add(root, "light00", nil, nil, 0444, light00)
	if err != nil {
		goto error
	}

	switch00 = new(Switch00)
	err = switch00.Add(root, "switch00", nil, nil, 0777, switch00)
	if err != nil {
		goto error
	}

	switch01 = new(Switch01)
	err = switch01.Add(root, "switch01", nil, nil, 0777, switch01)
	if err != nil {
		goto error
	}

	switch02 = new(Switch02)
	err = switch02.Add(root, "switch02", nil, nil, 0777, switch02)
	if err != nil {
		goto error
	}

	switch03 = new(Switch03)
	err = switch03.Add(root, "switch03", nil, nil, 0777, switch03)
	if err != nil {
		goto error
	}

	switch04 = new(Switch04)
	err = switch04.Add(root, "switch04", nil, nil, 0777, switch04)
	if err != nil {
		goto error
	}

	switch05 = new(Switch05)
	err = switch05.Add(root, "switch05", nil, nil, 0777, switch05)
	if err != nil {
		goto error
	}

	switch06 = new(Switch06)
	err = switch06.Add(root, "switch06", nil, nil, 0777, switch06)
	if err != nil {
		goto error
	}

	switch07 = new(Switch07)
	err = switch07.Add(root, "switch07", nil, nil, 0777, switch07)
	if err != nil {
		goto error
	}

	serv = srv.NewFileSrv(root)
	serv.Dotu = false

	if *debug {
		serv.Debuglevel = 1
	}

	serv.Start(serv)
	err = serv.StartNetListener("tcp", *addr)
	if err != nil {
		goto error
	}
	return

error:
	log.Println(fmt.Sprintf("Error: %s", err))
}
