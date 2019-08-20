package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"os"
	"strconv"

	"github.com/howeyc/gopass"
	"github.com/tredoe/osutil/user/crypt/sha512_crypt"
)

func main() {

	pwd := getPass()
	salt := getSalt()
	hasher := sha512_crypt.New()
	hash, err := hasher.Generate(pwd, salt)
	check(err)
	err = ioutil.WriteFile("passwd.hash", []byte(hash), 0644)
	check(err)
	fmt.Println("Password hash saved to passwd.hash")
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
	if len(os.Args) < 2 {
		randsalt := strconv.Itoa(rand.Intn(math.MaxInt32))
		buffer.WriteString("$6$")
		buffer.WriteString(randsalt)
	} else {
		buffer.WriteString("$6$")
		buffer.WriteString(os.Args[1])
	}
	defer buffer.Reset()
	return []byte(buffer.String())
}
func check(e error) {
	if e != nil {
		panic(e)
	}
}
