package main

import (
	"fmt"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

type state int

const (
	statusNormal state = iota
	stateDone
)

type Model struct {
	state  state
	lg     *lipgloss.Renderer
	styles *Styles
	form   *huh.Form
	width  int
}

func NewModel() Model {
	m := Model{width: maxWidth}
	m.lg = lipgloss.DefaultRenderer()
	m.styles = NewStyles(m.lg)

	m.form = huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Key("class").
				Options(huh.NewOptions("Warrior", "Mage", "Rogue")...).
				Title("Choose your class").
				Description("This will determine your department"),

			huh.NewSelect[string]().
				Key("level").
				Options(huh.NewOptions("1", "20", "9999")...).
				Title("Choose your level").
				Description("This will determine your benefits package"),

			huh.NewConfirm().
				Key("done").
				Title("All done?").
				Validate(func(v bool) error {
					if !v {
						return fmt.Errorf("Welp, finish up then")
					}
					return nil
				}).
				Affirmative("Yep").
				Negative("Wait, no"),
		),
	).
		WithWidth(45).
		WithShowHelp(false).
		WithShowErrors(false)
	return m
}

func (m Model) getRole() (string, string) {
	level := m.form.GetString("level")
	switch m.form.GetString("class") {
	case "Warrior":
		switch level {
		case "1":
			return "Tank Intern", "Assists with tank-related activities. Paid position."
		case "9999":
			return "Tank Manager", "Manages tanks and tank-related activities."
		default:
			return "Tank", "General tank. Does damage, takes damage. Responsible for tanking."
		}
	case "Mage":
		switch level {
		case "1":
			return "DPS Associate", "Finds DPS deals and passes them on to DPS Manager."
		case "9999":
			return "DPS Operating Officer", "Oversees all DPS activities."
		default:
			return "DPS", "Does damage and ideally does not take damage. Logs hours in JIRA."
		}
	case "Rogue":
		switch level {
		case "1":
			return "Stealth Junior Designer", "Designs rogue-like activities. Reports to Stealth Lead."
		case "9999":
			return "Stealth Lead", "Lead designer for all things stealth. Some travel required."
		default:
			return "Sneaky Person", "Sneaks around and does sneaky things. Reports to Stealth Lead."
		}
	default:
		return "", ""
	}
}
