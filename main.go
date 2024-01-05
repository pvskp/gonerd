package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const FONT_FAMILIES_ENDPOINT = "https://github.com/ryanoasis/nerd-fonts/tree/master/patched-fonts"
const FONT_DOWNLOAD_ENDPOINT = "https://github.com/ryanoasis/nerd-fonts/raw/HEAD/patched-fonts/"

func printSlice(s []string) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

// DownloadFont downloads a font from the Nerd Fonts GitHub repo
func DownloadFont(fontRelativeName string) {
	if !strings.HasSuffix(fontRelativeName, ".otf") && !strings.HasSuffix(fontRelativeName, ".ttf") {
		log.Fatalf("Invalid font file extension: %s", fontRelativeName)
	}

	fontPath := fmt.Sprintf("%s/%s", FONT_DOWNLOAD_ENDPOINT, fontRelativeName)

	resp, err := http.Get(fontPath)
	if err != nil {
		log.Printf("Error fetching URL: %v\n", err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Printf("Erro: status code %d\n", resp.StatusCode)
	}

	defer resp.Body.Close()

	fmt.Println("Creating file...")
	fmt.Println(fontRelativeName[strings.LastIndex(fontRelativeName, "/")+1:])
	out, err := os.Create(fontRelativeName[strings.LastIndex(fontRelativeName, "/")+1:])
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("File downloaded successfully")
}

func FetchDirInfo(dirName string) []string {
	fontInfoPath := fmt.Sprintf("%s/%s", FONT_FAMILIES_ENDPOINT, dirName)

	resp, err := http.Get(fontInfoPath)
	if err != nil {
		log.Fatalf("Error fetching URL: %v", err)
	}
	defer resp.Body.Close()

	ghResponse := FetchFromGitHub(resp.Body)

	var validRows []string

	for _, folder := range ghResponse.Payload.Tree.Items {
		if folder.ContentType == "directory" ||
			strings.HasSuffix(folder.Name, ".otf") ||
			strings.HasSuffix(folder.Name, ".ttf") {
			validRows = append(validRows, folder.Name)
		}
	}

	return validRows
}

func FetchFromGitHub(r io.Reader) *GitHubRepoResponse {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		log.Fatalf("Error loading HTTP response body: %v", err)
	}

	var jsonData string
	doc.Find("script[type='application/json']").Each(func(i int, s *goquery.Selection) {
		if dataTarget, exists := s.Attr("data-target"); exists && dataTarget == "react-app.embeddedData" {
			jsonData = s.Text()
		}
	})

	if jsonData == "" {
		log.Fatalf("No JSON data found")
	}

	ghResponse := GitHubRepoResponse{}
	if err := json.Unmarshal([]byte(jsonData), &ghResponse); err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}
	return &ghResponse
}

func GetFontFamilies() []string {
	resp, err := http.Get("https://github.com/ryanoasis/nerd-fonts/tree/master/patched-fonts")
	if err != nil {
		log.Fatalf("Error fetching font names: %v", err)
	}
	defer resp.Body.Close()

	ghResponse := FetchFromGitHub(resp.Body)

	var fontNames []string

	for _, font := range ghResponse.Payload.Tree.Items {
		fontNames = append(fontNames, font.Name)
	}

	return fontNames
}

func main() {
	DownloadFont("3270/Regular/3270NerdFont-Regular.ttf")
}
