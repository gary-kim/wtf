package git

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func (widget *Widget) display() {
	repoData := widget.currentData()
	if repoData == nil {
		widget.Redraw(widget.CommonSettings().Title, " Git repo data is unavailable ", false)
		return
	}

	title := fmt.Sprintf("%s - [green]%s[white]", widget.CommonSettings().Title, repoData.Repository)

	_, _, width, _ := widget.View.GetRect()
	str := widget.settings.common.SigilStr(len(widget.GitRepos), widget.Idx, width) + "\n"
	str += " [red]Branch[white]\n"
	str += fmt.Sprintf(" %s", repoData.Branch)
	str += "\n"
	str += widget.formatChanges(repoData.ChangedFiles)
	str += "\n"
	str += widget.formatCommits(repoData.Commits)

	widget.Redraw(title, str, false)
}

func (widget *Widget) formatChanges(data []string) string {
	str := " [red]Changed Files[white]\n"

	if len(data) == 1 {
		str += " [grey]none[white]\n"
	} else {
		for _, line := range data {
			str += widget.formatChange(line)
		}
	}

	return str
}

func (widget *Widget) formatChange(line string) string {
	if len(line) == 0 {
		return ""
	}

	line = strings.TrimSpace(line)
	firstChar, _ := utf8.DecodeRuneInString(line)

	// Revisit this and kill the ugly duplication
	switch firstChar {
	case 'A':
		line = strings.Replace(line, "A", "[green]A[white]", 1)
	case 'D':
		line = strings.Replace(line, "D", "[red]D[white]", 1)
	case 'M':
		line = strings.Replace(line, "M", "[yellow]M[white]", 1)
	case 'R':
		line = strings.Replace(line, "R", "[purple]R[white]", 1)
	}

	return fmt.Sprintf(" %s\n", strings.Replace(line, "\"", "", -1))
}

func (widget *Widget) formatCommits(data []string) string {
	str := " [red]Recent Commits[white]\n"

	for _, line := range data {
		str += widget.formatCommit(line)
	}

	return str
}

func (widget *Widget) formatCommit(line string) string {
	return fmt.Sprintf(" %s\n", strings.Replace(line, "\"", "", -1))
}
