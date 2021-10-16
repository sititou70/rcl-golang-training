package ftp

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path"
	"strconv"
	"strings"
)

func HandleFTP(c net.Conn) error {
	defer c.Close()

	// ready
	_, err := io.WriteString(c, FTPReplies[220])
	if err != nil {
		return err
	}

	// determine color
	remoteAddr := c.RemoteAddr().String()
	colorPrefix := []string{"31", "32", "33", "34", "35", "36", "37", "90", "91", "92", "93", "94", "95", "96", "97", "97;41", "97;42", "97;43", "97;44", "97;45", "97;46", "30;47", "97;100", "97;101", "30;102", "30;103", "97;104", "30;105", "30;106", "30;107"}[md5.Sum([]byte(remoteAddr))[0]%30]

	// create session state
	cwd, err := os.Getwd()
	state := FTPSessionState{
		CWD:  cwd,
		Type: "A",
	}
	if err != nil {
		return err
	}

	// handle commands
	s := bufio.NewScanner(c)
	for s.Scan() {
		log.Printf("\033[%sm[%s] %s\033[0m\n", colorPrefix, c.RemoteAddr().String(), s.Text())

		args := strings.Split(s.Text(), " ")
		switch strings.ToUpper(args[0]) {
		case "USER":
			err = handleUSER(c, args, &state)
		case "PASS":
			err = handlePASS(c, args, &state)
		case "SYST":
			err = handleSYST(c, args, &state)
		case "PORT":
			err = handlePORT(c, args, &state)
		case "LIST":
			err = handleLIST(c, args, &state)
		case "CWD":
			err = handleCWD(c, args, &state)
		case "PWD":
			err = handlePWD(c, args, &state)
		case "TYPE":
			err = handleTYPE(c, args, &state)
		case "RETR":
			err = handleRETR(c, args, &state)
		case "QUIT":
			return nil
		default:
			err = handleUnknownCommand(c, args, &state)
		}

		if err != nil {
			fmt.Println(err)
			_, err := io.WriteString(c, FTPReplies[550])
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func handleUSER(c net.Conn, args []string, state *FTPSessionState) error {
	if len(args) != 2 {
		_, err := io.WriteString(c, FTPReplies[500])
		if err != nil {
			return err
		}
	}

	_, err := io.WriteString(c, fmt.Sprintf(FTPReplies[331], args[1]))
	if err != nil {
		return err
	}
	state.User = args[1]

	return nil
}

func handlePASS(c net.Conn, args []string, state *FTPSessionState) error {
	if len(args) != 2 {
		_, err := io.WriteString(c, FTPReplies[500])
		if err != nil {
			return err
		}
	}

	_, err := io.WriteString(c, fmt.Sprintf(FTPReplies[230], state.User))
	if err != nil {
		return err
	}

	return nil
}

func handleSYST(c net.Conn, args []string, state *FTPSessionState) error {
	_, err := io.WriteString(c, FTPReplies[215])
	if err != nil {
		return err
	}

	return nil
}

func handlePORT(c net.Conn, args []string, state *FTPSessionState) error {
	if len(args) != 2 {
		_, err := io.WriteString(c, FTPReplies[500])
		if err != nil {
			return err
		}
	}

	nums := strings.Split(args[1], ",")
	if len(nums) != 6 {
		_, err := io.WriteString(c, FTPReplies[500])
		if err != nil {
			return err
		}
	}

	address := strings.Join(nums[0:4], ".")
	p1, err := strconv.Atoi(nums[4])
	if err != nil {
		return err
	}
	p2, err := strconv.Atoi(nums[5])
	if err != nil {
		return err
	}

	state.LatestAddress = address + ":" + strconv.Itoa(p1*256+p2)

	_, err = io.WriteString(c, fmt.Sprintf(FTPReplies[200], "PORT"))
	if err != nil {
		return err
	}

	return nil
}

func handleLIST(c net.Conn, args []string, state *FTPSessionState) error {
	_, err := io.WriteString(c, FTPReplies[150])
	if err != nil {
		return err
	}

	dir, err := listDir(state.CWD)
	if err != nil {
		return err
	}
	err = writeFTPData(state.LatestAddress, dir)
	if err != nil {
		return err
	}

	_, err = io.WriteString(c, FTPReplies[250])
	if err != nil {
		return err
	}

	return nil
}

func handleCWD(c net.Conn, args []string, state *FTPSessionState) error {
	if len(args) != 2 {
		_, err := io.WriteString(c, FTPReplies[500])
		if err != nil {
			return err
		}
	}

	newPath := path.Join(state.CWD, args[1])
	_, err := os.Stat(newPath)
	if err != nil {
		return err
	}

	state.CWD = newPath

	_, err = io.WriteString(c, FTPReplies[250])
	if err != nil {
		return err
	}

	return nil
}

func handlePWD(c net.Conn, args []string, state *FTPSessionState) error {
	_, err := io.WriteString(c, fmt.Sprintf("257 \"%s\" is current directory.\r\n", state.CWD))
	if err != nil {
		return err
	}

	return nil
}

func handleTYPE(c net.Conn, args []string, state *FTPSessionState) error {
	if len(args) != 2 {
		_, err := io.WriteString(c, FTPReplies[500])
		if err != nil {
			return err
		}
	}

	state.Type = args[1]

	_, err := io.WriteString(c, fmt.Sprintf(FTPReplies[200], fmt.Sprintf("Type set to %s,", state.Type)))
	if err != nil {
		return err
	}

	return nil
}

func handleRETR(c net.Conn, args []string, state *FTPSessionState) error {
	if len(args) != 2 {
		_, err := io.WriteString(c, FTPReplies[500])
		if err != nil {
			return err
		}
	}

	targetPath := path.Join(state.CWD, args[1])
	_, err := os.Stat(targetPath)
	if err != nil {
		return err
	}

	f, err := os.Open(targetPath)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.WriteString(c, FTPReplies[150])
	if err != nil {
		return err
	}

	switch state.Type {
	case "A":
		b := new(bytes.Buffer)
		s := bufio.NewScanner(f)
		for s.Scan() {
			fmt.Fprintf(b, "%s\r\n", s.Text())
		}
		err = writeFTPData(state.LatestAddress, b)
	case "I":
		err = writeFTPData(state.LatestAddress, f)
	default:
		return fmt.Errorf("unknown transfer type")
	}
	if err != nil {
		return err
	}

	_, err = io.WriteString(c, FTPReplies[250])
	if err != nil {
		return err
	}

	return nil
}

func handleUnknownCommand(c net.Conn, args []string, state *FTPSessionState) error {
	_, err := io.WriteString(c, FTPReplies[502])
	if err != nil {
		return err
	}

	return nil
}
