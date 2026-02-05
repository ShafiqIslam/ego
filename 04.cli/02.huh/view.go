package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) completedView() string {
	s := m.styles

	title, role := m.getRole()
	title = s.Highlight.Render(title)

	var b strings.Builder
	fmt.Fprintf(&b, "Congratulations, you’re Charm’s newest\n%s!\n\n", title)
	fmt.Fprintf(&b, "Your job description is as follows:\n\n%s\n\nPlease proceed to HR immediately.", role)
	return s.Status.Margin(0, 1).Padding(1, 2).Width(48).Render(b.String()) + "\n\n"
}

func (m Model) applicationView() string {
	s := m.styles

	header := m.headerView()

	v := strings.TrimSuffix(m.form.View(), "\n\n")
	form := m.lg.NewStyle().Margin(1, 0).Render(v)
	status := m.statusView(form)
	body := lipgloss.JoinHorizontal(lipgloss.Left, form, status)

	footer := m.footerView()

	return s.Base.Render(header + "\n" + body + "\n\n" + footer)
}

func (m Model) headerView() string {
	if len(m.form.Errors()) > 0 {
		return m.appErrorBoundaryView(m.errorView())
	}

	return m.appBoundaryView("Charm Employment Application")
}

func (m Model) statusView(form string) string {
	s := m.styles

	var (
		buildInfo      = "(None)"
		role           string
		jobDescription string
		level          string
	)

	if m.form.GetString("level") != "" {
		level = "Level: " + m.form.GetString("level")
		role, jobDescription = m.getRole()
		role = "\n\n" + s.StatusHeader.Render("Projected Role") + "\n" + role
		jobDescription = "\n\n" + s.StatusHeader.Render("Duties") + "\n" + jobDescription
	}
	if m.form.GetString("class") != "" {
		var class string
		if m.form.GetString("class") != "" {
			class = "Class: " + m.form.GetString("class")
		}

		buildInfo = fmt.Sprintf("%s\n%s", class, level)
	}

	const statusWidth = 28
	statusMarginLeft := m.width - statusWidth - lipgloss.Width(form) - s.Status.GetMarginRight()
	return s.Status.
		Height(lipgloss.Height(form)).
		Width(statusWidth).
		MarginLeft(statusMarginLeft).
		Render(s.StatusHeader.Render("Current Build") + "\n" +
			buildInfo +
			role +
			jobDescription)
}

func (m Model) footerView() string {
	if len(m.form.Errors()) > 0 {
		return m.appErrorBoundaryView("")
	}

	return m.appBoundaryView(m.form.Help().ShortHelpView(m.form.KeyBinds()))
}

func (m Model) errorView() string {
	var s strings.Builder
	for _, err := range m.form.Errors() {
		s.WriteString(err.Error())
	}
	return s.String()
}

func (m Model) appBoundaryView(text string) string {
	return lipgloss.PlaceHorizontal(
		m.width,
		lipgloss.Left,
		m.styles.HeaderText.Render(text),
		lipgloss.WithWhitespaceChars("/"),
		lipgloss.WithWhitespaceForeground(indigo),
	)
}

func (m Model) appErrorBoundaryView(text string) string {
	return lipgloss.PlaceHorizontal(
		m.width,
		lipgloss.Left,
		m.styles.ErrorHeaderText.Render(text),
		lipgloss.WithWhitespaceChars("/"),
		lipgloss.WithWhitespaceForeground(red),
	)
}
