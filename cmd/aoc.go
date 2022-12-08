package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

const (
	aocUrl  = "https://adventofcode.com"
	success = "That's the right answer!"
)

var Session string

var levels = map[string]string{
	"A": "1",
	"B": "2",
}

type InputRequest struct {
	Year, Day int
	Path      string
}

type SubmissionRequest struct {
	Year, Day int
	Level     string
}

func closeOrPanic(body io.Closer) {
	if err := body.Close(); err != nil {
		panic(err)
	}
}

func GetInput(r *InputRequest) error {
	if _, err := os.Stat(r.Path); err == nil {
		log.Printf("File already exists: %s", r.Path)
		return nil
	}

	resource := fmt.Sprintf("%s/%d/day/%d/input", aocUrl, r.Year, r.Day)

	req, err := http.NewRequest(http.MethodGet, resource, nil)
	if err != nil {
		return err
	}

	req.AddCookie(&http.Cookie{Name: "session", Value: Session})

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to get resource %s: %s", resource, res.Status)
	}
	defer closeOrPanic(res.Body)

	f, err := os.Create(r.Path)
	if err != nil {
		return err
	}
	defer closeOrPanic(f)

	if _, err = f.ReadFrom(res.Body); err != nil {
		return err
	}

	return nil
}

func SubmitAnswer(ans string, r *SubmissionRequest) error {
	formData := url.Values{
		"level":  {levels[r.Level]},
		"answer": {ans},
	}

	resource := fmt.Sprintf("%s/%d/day/%d/answer", aocUrl, r.Year, r.Day)

	req, err := http.NewRequest(http.MethodPost, resource, strings.NewReader(formData.Encode()))
	if err != nil {
		return err
	}

	req.AddCookie(&http.Cookie{Name: "session", Value: Session})
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to post data %s to resource %s: %s", formData, resource, res.Status)
	}
	defer closeOrPanic(res.Body)

	bts, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	txt := string(bts)

	if !strings.Contains(txt, success) {
		return fmt.Errorf(txt)
	}

	log.Println(success)
	return nil
}

func CreateFromTemplate(day int) error {
	name := fmt.Sprintf("day%02d", day)
	if err := os.Mkdir(name, 0777); err != nil {
		return err
	}

	// test file
	content, err := os.ReadFile("template/template_test.go")
	if err != nil {
		return err
	}
	if err = os.WriteFile(name+"/"+name+"_test.go", content, 0644); err != nil {
		return err
	}

	// source file
	content, err = os.ReadFile("template/template.go")
	if err != nil {
		return err
	}
	content = bytes.ReplaceAll(content, []byte("-1_234"), []byte(strconv.Itoa(day)))
	if err = os.WriteFile(name+"/"+name+".go", content, 0644); err != nil {
		return err
	}

	return nil
}
