package liner

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

type viState struct {
	viEnabled     bool
	viPrompt      bool
	viMode        vimode
	viNormalStyle style
	viInsertStyle style
	lastline      []rune
	debug         log.Logger
}

type vimode int

const (
	ViNormal vimode = iota
	ViInsert
	ViReplace
)

// enable or disable the debug log for vi functions
func (s *viState) InitializeLogger(state bool) error {
	if !state {
		s.debug.SetOutput(io.Discard)
		return nil
	}
	dir, _ := os.UserHomeDir()
	logfile, err := os.OpenFile(filepath.Join(dir, "liner.log"), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0655)
	if err != nil {
		return err
	}
	s.debug.SetOutput(logfile)
	s.debug.Println("log intialized")
	return nil
}

// enable or disable vi keybindings
func (s *viState) EnableViMode(state bool) {
	s.InitializeLogger(false) // debug log is off by default
	s.debug.Printf("vimode set to %v\n", state)
	s.viEnabled = state
}

// allow prompt to change visually based on current viMode
func (s *viState) EnableViPrompt(state bool) {
	s.debug.Printf("viprompt set to %v\n", state)
	s.viPrompt = state
}

// explicitly set the vi mode (normal, insert, replace)
func (s *viState) SetViMode(state vimode) {
	s.debug.Printf("vimode set to %v\n", state)
	s.viMode = state
}

// style used when in vi normal mode
func (s *viState) SetViNormalStyle(style style) {
	s.debug.Printf("viNormal style set to %v\n", style)
	s.viNormalStyle = style
}

// style used when in vi insert mode
func (s *viState) SetViInsertStyle(style style) {
	s.debug.Printf("viInsert style set to %v\n", style)
	s.viInsertStyle = style
}

// explicitly set style by printing the associated escape code
func (s *viState) setStyle(style style) {
	code, ok := styleCode[style]
	if !ok {
		return
	}
	fmt.Printf(code)
}

// call this function to automatically print the escape code assigned to the current vi mode
func (s *viState) toggleStyle() {
	if !s.viEnabled || !s.viPrompt {
		return
	}
	if s.viMode == ViInsert {
		s.setStyle(s.viInsertStyle)
	}
	if s.viMode == ViNormal {
		s.setStyle(s.viNormalStyle)
	}
}

// put the line into normal mode unconditionally
func (s *State) enterViNormal(prompt []rune, line []rune, pos int) error {
	s.debug.Printf("enter normal mode\n")
	s.viMode = ViNormal
	s.needRefresh = true
	err := s.refresh(prompt, line, pos)
	if err != nil {
		return err
	}
	return nil
}

// put the line into insert mode unconditionally
func (s *State) enterViInsert(prompt []rune, line []rune, pos int) error {
	s.debug.Printf("enter insert mode\n")
	s.viMode = ViInsert
	s.needRefresh = true
	err := s.refresh(prompt, line, pos)
	if err != nil {
		return err
	}
	return nil
}

// put the line into replace mode
func (s *State) enterViReplace() {
	s.debug.Printf("enter replace mode\n")
	s.viMode = ViReplace
}

type style int

const (
	Default style = iota
	Bold
	Dim
	Italic
	Under
	Blink
	Reverse
	Invisible
	Strikethrough
)

var styleCode = map[style]string{
	Default:       "\x1b[0m",
	Bold:          "\x1b[1m",
	Dim:           "\x1b[2m",
	Italic:        "\x1b[3m",
	Under:         "\x1b[4m",
	Blink:         "\x1b[6m",
	Reverse:       "\x1b[7m",
	Invisible:     "\x1b[8m",
	Strikethrough: "\x1b[9m",
}
