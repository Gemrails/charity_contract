package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
)

type CharityUser struct {
	DonationName string `json:"donationName"`
	ALLMoney     int32  `json:"allMoney"`
	LeftMoney    int32  `json:"leftMoney"`
	DealNumbers  int    `json:"dealNumbers"`
}

type ExecArgs struct {
	Args []string
}

func DonationUser(username, money string) {
	//cUser := &CharityUser{}
	cexec := &ExecArgs{}
	cexec.Args = append(cexec.Args, "donation")
	cexec.Args = append(cexec.Args, username)
	cexec.Args = append(cexec.Args, money)
	cexecBytes, err := json.Marshal(cexec)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println("asdfadfafadfa")
	f, err := RunCommand(string(cexecBytes[:]))
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Print(f)

	/*
		iMoney, err := strconv.ParseInt(money, 10, 64)
		if err != nil {
			return cUser, err
		}
		cUser = &CharityUser{
			DonationName: username + f,
			ALLMoney:     int32(iMoney),
			LeftMoney:    int32(iMoney),
			DealNumbers:  0,
		}
		return cUser, nil
	*/
}

func RunCommand(arg string) (string, error) {
	//cmd := exec.Command("bash", "1.sh", "peer", "chaincode", "invoke", "-n", "charity", "-c", arg, "-C", "myc")
	cmd := exec.Command("bash", "1.sh", arg)
	fmt.Println(cmd)
	//f, err := exec.Command("echo", arg).Output()
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	f, err := cmd.Output()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		fmt.Println("99999")
		fmt.Println(err)
		return "", err
	}
	fmt.Println("12313131321313231")
	fmt.Println(string(f[:]))
	for i := 0; i < len(f); i++ {
		if f[i] == 0 {
			fmt.Print("fadf")
			return string(f[0:i]), nil
		}
	}
	fmt.Println("end")
	return string(f), nil
}

func main() {
	DonationUser("aa", "20000")
}
