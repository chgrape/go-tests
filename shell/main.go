package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var cmds = map[string]func([]string){}

func init() {
	cmds["exit"] = exit
	cmds["echo"] = echo
	cmds["type"] = _type
	cmds["pwd"] = pwd
	cmds["cd"] = cd
}

func exit(input []string) {
	if strings.Join(input, "") == "0" {
		os.Exit(0)
	} else {
		os.Exit(1)
	}
}

func echo(input []string) {
	if len(input) == 0 {
		fmt.Println("echo works")
		return
	}
	fmt.Println(strings.Join(input, " "))
}

func _type(input []string) {
	if len(input) != 1 {
		return
	}

	cmd := strings.Join(input, "")
	full_path, _ := exec.LookPath(cmd)

	_, ok := cmds[cmd]

	if ok {
		fmt.Println(cmd + " is a shell builtin")
		return
	}

	if full_path != "" {
		fmt.Println(cmd + " is " + full_path)
	} else {
		fmt.Println(cmd + ": not found")

	}

}

func _exec(input []string) {

	full_path, _ := exec.LookPath(input[0])

	if full_path != "" {
		output, err := exec.Command(input[0], input[1:]...).CombinedOutput()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error executing command:", err)
			return
		}
		fmt.Print(string(output))
	} else {
		fmt.Fprintf(os.Stderr, "%s: command not found\n", input[0])
	}
}

func pwd(input []string) {
	path, _ := os.Getwd()

	fmt.Println(path)
}

func cd(input []string) {
	if len(input) > 1 {
		fmt.Println("Too many arguments")
	}

	path := strings.Join(input, "")
	cur := path

	if string(path[:2]) == ".." {
		path_arr := strings.Split(path, "/")
		for _, seg := range path_arr {
			if seg == "" {
				continue
			}
			// ../
			// ../..
			// ../as/../asd
			// ../as/sa
			cur_dir, err := os.Getwd()
			if err != nil {
				os.Exit(1)
			}
			path = filepath.Dir(cur_dir)
			os.Chdir(path)
		}
		return
	} else if string(path[0]) == "." {
		//relative path
		// ./pat/to/dir
		// /bin/pat/to/dir
		cur_dir, err := os.Getwd()
		if err != nil {
			os.Exit(1)
		}
		path = cur_dir + string(input[0][1:])
	}
	ev, _ := os.Stat(path)
	if ev == nil {
		fmt.Println("cd: " + cur + ": No such file or directory")
		return
	}

	os.Chdir(path)

}

func main() {
	for {
		writer := bufio.NewWriter(os.Stdout)
		writer.WriteString("$ ")
		writer.Flush()

		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			panic(err)
		}
		line := strings.TrimSpace(command)
		args := strings.Split(line, " ")
		cmd, flags := args[0], args[1:]

		run, ok := cmds[cmd]

		if ok {
			run(flags)
		} else {
			_exec(args)
		}

	}
}
