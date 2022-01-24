package search

import (
	"io/fs"
	"strings"

	"github.com/manifoldco/promptui"
)

type Searcher struct {
	Items []string
}

func NewFileSearcher(files []fs.FileInfo) *Searcher {
	items := []string{}
	for _, f := range files {
		if f.Name() != ".git" {
			items = append(items, f.Name())
		}
	}

	return &Searcher{items}
}

func (s *Searcher) SearchPrompt(label string) *promptui.Select {
	return &promptui.Select{
		Label:             label,
		Items:             s.Items,
		StartInSearchMode: true,
		Searcher:          s.Search,
		Size:              10,
	}
}

func (s *Searcher) Search(input string, index int) bool {
	item := strings.Replace(strings.ToLower(s.Items[index]), " ", "", -1)
	input = strings.Replace(strings.ToLower(input), " ", "", -1)

	return strings.Contains(item, input)
}
