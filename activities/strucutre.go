package activities

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"reflect"
	"runtime"
	"strconv"
	"strings"

	"github.com/omar0ali/make_request_cli/models"
)

func Draw(menu []models.MenuItem, handleClick func(models.MenuItem) bool) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		for i := 0; i < len(menu); i++ {
			if i == 0 {
				fmt.Println(strings.Repeat("*", 50))
				fmt.Printf("* Type %v to %v\tMAKE_REQUEST_CLI\t *\n", menu[i].ID, menu[i].Title)
				fmt.Println(strings.Repeat("*", 50))
				continue
			}
			fmt.Printf("%v: %v\n", menu[i].ID, menu[i].Title)
		}
		fmt.Print("ENTER: ")

		if !scanner.Scan() {
			fmt.Println("Error reading input:", scanner.Err())
			return
		}
		var item models.MenuItem
		choice, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
		if choice < 0 || choice > len(menu) {
			fmt.Println("Try another choice.")
			continue
		} else {
			item = menu[choice]
		}
		if err != nil {
			fmt.Printf("Error: %v\n", err.Error())
			return
		}
		if endSignal := handleClick(item); endSignal {
			break
		}
	}
}

func Input(title string) (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("%v: ", title)
	if !scanner.Scan() {
		fmt.Println("Error reading input:", scanner.Err())
		return "", scanner.Err()
	}
	return strings.TrimSpace(scanner.Text()), nil
}

func ConvertToInt(value string) (int, error) {
	return strconv.Atoi(value)
}

func MakeSelection(lists interface{}) (int, error) {
	v := reflect.ValueOf(lists)
	if v.Kind() != reflect.Slice {
		return -1, fmt.Errorf("MakeSelection: expected a slice, got %v", v.Kind())
	}

	if v.Len() < 1 {
		return -1, fmt.Errorf("there is nothing to select")
	}

	var builder strings.Builder
	builder.WriteString("Select one of the following:\n\n")
	for i := 0; i < v.Len(); i++ {
		builder.WriteString(fmt.Sprintf("%d: %v\n", i, v.Index(i)))
	}
	Display(builder.String())

	choiceStr, err := Input("To cancel type any letter.\nENTER (i.e 0)")
	if err != nil {
		return -1, err
	}
	choice, err := ConvertToInt(choiceStr)
	if err != nil {
		return -1, err
	}

	if choice < 0 || choice >= v.Len() {
		return -1, fmt.Errorf("MakeSelection: choice out of range")
	}

	return choice, nil
}

func CreateDialogYesNo(title string, ifTrue func() error) error {
	ClearScreen()
	fmt.Println(strings.Repeat("*", 50))
	dialog := models.CreateDialog(title)
	fmt.Println("* ", dialog.GetQuestion())
	choice, err := Input("* 1: Yes\t2: Cancel\n* Type 1 or 2")
	if err != nil {
		return err
	}
	var selection bool = false
	if choice == "1" {
		selection = true
	}
	dialog.SetChoice(selection)
	dialog.OnClick(func() {
		ifTrue()
	})
	ClearScreen()
	return nil
}

const (
	GET    = "GET"
	POST   = "POST"
	DELETE = "DELETE"
)

func HttpRequest(method, url string, jsonData []byte) ([]byte, error) {
	var resp *http.Response
	var err error

	// Perform HTTP request based on method
	switch method {
	case GET:
		resp, err = http.Get(url)
	case POST:
		req, err := http.NewRequest(POST, url, bytes.NewBuffer([]byte(fmt.Sprintf(`%v`, string(jsonData)))))
		if err != nil {
			return nil, err
		}
		req.Header.Set("Content-Type", "application/json") // Set content type header
		client := &http.Client{}
		resp, err = client.Do(req)
		if err != nil {
			return nil, err
		}
	case DELETE:
		req, err := http.NewRequest(DELETE, url, nil) // Create DELETE request
		if err != nil {
			return nil, err
		}
		client := &http.Client{}
		resp, err = client.Do(req) // Send DELETE request
		if err != nil {
			return nil, err
		}
	}
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read and return response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func ClearScreen() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		fmt.Print("\033[H\033[2J") // ANSI escape code for clearing the screen
	}
}

func Display(data string) {
	var builder strings.Builder
	builder.WriteString(strings.Repeat("#", 50))
	builder.WriteString(fmt.Sprintf("\n\n%v\n\n", data))
	builder.WriteString(strings.Repeat("#", 50))
	fmt.Println(builder.String())
}
