package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"

	"github.com/howeyc/gopass"
	"github.com/tredoe/osutil/user/crypt/sha512_crypt"
)

func main() {
	toStdout := flag.Bool("stdout", false, "Output to stdout instead of file")
	filename := flag.String("filename", "passwd.hash", "Filename to write password hash")
	flag.Parse()
	if flag.NArg() != 0 {
		checkSalt()
	}
	pwd := getPass()
	salt := getSalt()
	hasher := sha512_crypt.New()
	hash, err := hasher.Generate(pwd, salt)
	check(err)
	if *toStdout == true {
		fmt.Println(hash)
		os.Exit(0)
	} else {
		err = ioutil.WriteFile(*filename, []byte(hash), 0644)
		check(err)
		fmt.Println("Password hash saved to", *filename)
	}
}
func getPass() []byte {
	var pwd []byte
	var pwdConf []byte
	var err error
	fmt.Print("Enter password: ")
	pwd, err = gopass.GetPasswd()
	check(err)
	fmt.Print("Confirm password: ")
	pwdConf, err = gopass.GetPasswd()
	check(err)
	if string(pwd) != string(pwdConf) {
		fmt.Println("Password mismatch! Exiting...")
		os.Exit(10)
	}
	return pwd
}
func getSalt() []byte {
	var buffer bytes.Buffer
	const randChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789."
	if flag.NArg() == 0 {
		randsalt := make([]byte, 16)
		for i := range randsalt {
			randsalt[i] = randChars[rand.Intn(len(randChars))]
		}
		buffer.WriteString("$6$")
		for i := range randsalt {
			buffer.WriteByte(randsalt[i])
		}
	} else {
		buffer.WriteString("$6$")
		buffer.WriteString(flag.Arg(0))
	}
	defer buffer.Reset()
	return []byte(buffer.String())
}

func checkSalt() {
	if len(os.Args[1]) < 8 || len(os.Args[1]) > 16 {
		fmt.Println("Provided salt might be 8 to 16 characters")
		os.Exit(10)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
