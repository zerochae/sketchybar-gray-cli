package tui

import (
	"fmt"
	"sort"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/zerochae/gsbar/internal/colors"
	"github.com/zerochae/gsbar/internal/config"
	"github.com/zerochae/gsbar/internal/sketchybar"
)

type screen int

const (
	menuScreen screen = iota
	listScreen
	setScreen
	getScreen
	reloadScreen
	resultScreen
)

type model struct {
	screen       screen
	cursor       int
	menuItems    []string
	configs      map[string]string
	configKeys   []string
	inputMode    string
	keyInput     string
	valueInput   string
	result       string
	err          error
	quitting     bool
}

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color(colors.Magenta)).
			MarginBottom(1)

	selectedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(colors.Sky)).
			Bold(true)

	normalStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(colors.Pearl))

	errorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(colors.Red)).
			Bold(true)

	successStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(colors.Green)).
			Bold(true)

	accentStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(colors.Fuchsia)).
			Bold(true)

	dimStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(colors.Gray))

	containerStyle = lipgloss.NewStyle().
			Padding(1, 2)
)

func initialModel() model {
	return model{
		screen: menuScreen,
		cursor: 0,
		menuItems: []string{
			"Config List",
			"Config Set",
			"Config Get",
			"Reload Sketchybar",
			"Quit",
		},
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			// 입력 모드에서는 q를 일반 문자로 처리
			if (m.screen == setScreen && m.inputMode != "") || m.screen == getScreen {
				if msg.String() == "q" {
					if m.screen == setScreen {
						if m.inputMode == "key" {
							m.keyInput += msg.String()
						} else if m.inputMode == "value" {
							m.valueInput += msg.String()
						}
					} else if m.screen == getScreen {
						m.keyInput += msg.String()
					}
				} else {
					// ctrl+c는 항상 종료
					m.quitting = true
					return m, tea.Quit
				}
			} else if m.screen == menuScreen || m.screen == resultScreen {
				m.quitting = true
				return m, tea.Quit
			} else {
				m.screen = menuScreen
				m.cursor = 0
				m.result = ""
				m.err = nil
			}
			return m, nil

		case "esc":
			m.screen = menuScreen
			m.cursor = 0
			m.result = ""
			m.err = nil
			return m, nil

		case "up", "k":
			// 입력 모드에서는 네비게이션 무시
			if (m.screen == setScreen && m.inputMode != "") || m.screen == getScreen {
				if m.screen == setScreen {
					if m.inputMode == "key" {
						m.keyInput += msg.String()
					} else if m.inputMode == "value" {
						m.valueInput += msg.String()
					}
				} else if m.screen == getScreen {
					m.keyInput += msg.String()
				}
			} else if m.screen == menuScreen && m.cursor > 0 {
				m.cursor--
			} else if m.screen == listScreen && m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			// 입력 모드에서는 네비게이션 무시
			if (m.screen == setScreen && m.inputMode != "") || m.screen == getScreen {
				if m.screen == setScreen {
					if m.inputMode == "key" {
						m.keyInput += msg.String()
					} else if m.inputMode == "value" {
						m.valueInput += msg.String()
					}
				} else if m.screen == getScreen {
					m.keyInput += msg.String()
				}
			} else if m.screen == menuScreen && m.cursor < len(m.menuItems)-1 {
				m.cursor++
			} else if m.screen == listScreen && m.cursor < len(m.configKeys)-1 {
				m.cursor++
			}

		case "enter":
			if m.screen == menuScreen {
				return m.handleMenuSelection()
			} else if m.screen == setScreen {
				if m.inputMode == "" {
					m.inputMode = "key"
				} else if m.inputMode == "key" {
					m.inputMode = "value"
				} else if m.inputMode == "value" {
					return m.handleSetConfig()
				}
			} else if m.screen == getScreen {
				return m.handleGetConfig()
			} else if m.screen == resultScreen {
				m.screen = menuScreen
				m.cursor = 0
				m.result = ""
				m.err = nil
			}

		case "backspace":
			if m.screen == setScreen {
				if m.inputMode == "key" && len(m.keyInput) > 0 {
					m.keyInput = m.keyInput[:len(m.keyInput)-1]
				} else if m.inputMode == "value" && len(m.valueInput) > 0 {
					m.valueInput = m.valueInput[:len(m.valueInput)-1]
				}
			} else if m.screen == getScreen && len(m.keyInput) > 0 {
				m.keyInput = m.keyInput[:len(m.keyInput)-1]
			}

		default:
			if m.screen == setScreen {
				if m.inputMode == "key" {
					m.keyInput += msg.String()
				} else if m.inputMode == "value" {
					m.valueInput += msg.String()
				}
			} else if m.screen == getScreen {
				m.keyInput += msg.String()
			}
		}
	}

	return m, nil
}

func (m model) handleMenuSelection() (tea.Model, tea.Cmd) {
	switch m.cursor {
	case 0:
		m.screen = listScreen
		m.cursor = 0
		m.loadConfigs()
	case 1:
		m.screen = setScreen
		m.inputMode = ""
		m.keyInput = ""
		m.valueInput = ""
	case 2:
		m.screen = getScreen
		m.keyInput = ""
	case 3:
		m.screen = reloadScreen
		return m.handleReload()
	case 4:
		m.quitting = true
		return m, tea.Quit
	}
	return m, nil
}

func (m *model) loadConfigs() {
	userCfg := config.NewUser()
	if userCfg != nil {
		userCfg.Load()
		m.configs = userCfg.List()
		m.configKeys = make([]string, 0, len(m.configs))
		for k := range m.configs {
			m.configKeys = append(m.configKeys, k)
		}
		sort.Strings(m.configKeys)
	}
}

func (m model) handleSetConfig() (tea.Model, tea.Cmd) {
	userCfg := config.NewUser()
	var err error

	if userCfg == nil {
		err = fmt.Errorf("failed to get user config path")
	} else {
		if err = userCfg.Load(); err == nil {
			userCfg.Set(m.keyInput, m.valueInput)
			err = userCfg.Save()
		}
	}

	m.screen = resultScreen
	if err != nil {
		m.err = err
	} else {
		m.result = fmt.Sprintf("Set %s = %s", m.keyInput, m.valueInput)
	}
	m.keyInput = ""
	m.valueInput = ""
	m.inputMode = ""

	return m, nil
}

func (m model) handleGetConfig() (tea.Model, tea.Cmd) {
	value, err := config.GetValueCascade(m.keyInput)

	m.screen = resultScreen
	if err != nil {
		m.err = err
	} else {
		m.result = fmt.Sprintf("%s = %s", m.keyInput, value)
	}
	m.keyInput = ""

	return m, nil
}

func (m model) handleReload() (tea.Model, tea.Cmd) {
	err := sketchybar.ReloadSketchybar()

	m.screen = resultScreen
	if err != nil {
		m.err = fmt.Errorf("failed to reload sketchybar: %w", err)
	} else {
		m.result = "Sketchybar reloaded successfully"
	}

	return m, nil
}

func (m model) View() string {
	if m.quitting {
		return ""
	}

	var content string
	switch m.screen {
	case menuScreen:
		content = m.viewMenu()
	case listScreen:
		content = m.viewList()
	case setScreen:
		content = m.viewSet()
	case getScreen:
		content = m.viewGet()
	case resultScreen:
		content = m.viewResult()
	}

	return containerStyle.Render(content)
}

func (m model) viewMenu() string {
	s := titleStyle.Render("┌─┐┬─┐┌─┐┬ ┬   ┌─┐┬┌─┌─┐┌┬┐┌─┐┬ ┬┬ ┬   ┌┐ ┌─┐┬─┐\n│ ┬├┬┘├─┤└┬┘───└─┐├┴┐├┤  │ │  ├─┤└┬┘───├┴┐├─┤├┬┘\n└─┘┴└─┴ ┴ ┴    └─┘┴ ┴└─┘ ┴ └─┘┴ ┴ ┴    └─┘┴ ┴┴└─")
	s += "\n\n"

	for i, item := range m.menuItems {
		if m.cursor == i {
			cursor := accentStyle.Render("❯")
			s += cursor + " " + selectedStyle.Render(item) + "\n"
		} else {
			s += "  " + normalStyle.Render(item) + "\n"
		}
	}

	s += "\n" + dimStyle.Render("[↑/↓] Navigate  [Enter] Select  [q] Quit")
	return s
}

func (m model) viewList() string {
	s := titleStyle.Render("Configuration List")
	s += "\n\n"

	if len(m.configKeys) == 0 {
		s += dimStyle.Render("No user configuration found")
	} else {
		keyStyle := lipgloss.NewStyle().Foreground(lipgloss.Color(colors.Aqua))
		valueStyle := lipgloss.NewStyle().Foreground(lipgloss.Color(colors.Lime))

		for i, key := range m.configKeys {
			keyPart := keyStyle.Render(key)
			valuePart := valueStyle.Render(m.configs[key])
			line := keyPart + " " + dimStyle.Render("=") + " " + valuePart

			if i == m.cursor {
				cursor := accentStyle.Render("❯")
				s += cursor + " " + line + "\n"
			} else {
				s += "  " + line + "\n"
			}
		}
	}

	s += "\n" + dimStyle.Render("[↑/↓] Navigate  [esc] Back  [q] Quit")
	return s
}

func (m model) viewSet() string {
	s := titleStyle.Render("Set Configuration")
	s += "\n\n"

	labelStyle := lipgloss.NewStyle().Foreground(lipgloss.Color(colors.Aqua)).Bold(true)
	cursorStyle := lipgloss.NewStyle().Foreground(lipgloss.Color(colors.Magenta))

	if m.inputMode == "" {
		s += dimStyle.Render("Press [Enter] to start") + "\n"
	} else {
		s += labelStyle.Render("Key: ")
		if m.inputMode == "key" {
			s += selectedStyle.Render(m.keyInput) + cursorStyle.Render("█")
		} else {
			s += normalStyle.Render(m.keyInput)
		}
		s += "\n"

		s += labelStyle.Render("Value: ")
		if m.inputMode == "value" {
			s += selectedStyle.Render(m.valueInput) + cursorStyle.Render("█")
		} else {
			s += normalStyle.Render(m.valueInput)
		}
		s += "\n"
	}

	s += "\n"
	if m.inputMode == "" {
		s += dimStyle.Render("[Enter] Start  [esc] Back")
	} else if m.inputMode == "key" {
		s += dimStyle.Render("[Enter] Next  [esc] Back")
	} else if m.inputMode == "value" {
		s += dimStyle.Render("[Enter] Save  [esc] Back")
	}

	return s
}

func (m model) viewGet() string {
	s := titleStyle.Render("Get Configuration")
	s += "\n\n"

	labelStyle := lipgloss.NewStyle().Foreground(lipgloss.Color(colors.Aqua)).Bold(true)
	cursorStyle := lipgloss.NewStyle().Foreground(lipgloss.Color(colors.Magenta))

	s += labelStyle.Render("Key: ") + selectedStyle.Render(m.keyInput) + cursorStyle.Render("█")
	s += "\n\n"

	s += dimStyle.Render("[Enter] Get  [esc] Back  [q] Quit")
	return s
}

func (m model) viewResult() string {
	s := titleStyle.Render("Result")
	s += "\n\n"

	if m.err != nil {
		iconStyle := lipgloss.NewStyle().Foreground(lipgloss.Color(colors.Red)).Bold(true)
		s += iconStyle.Render("✗") + " " + errorStyle.Render(m.err.Error())
	} else {
		iconStyle := lipgloss.NewStyle().Foreground(lipgloss.Color(colors.Green)).Bold(true)
		s += iconStyle.Render("✓") + " " + successStyle.Render(m.result)
	}

	s += "\n\n"
	s += dimStyle.Render("[Enter] Continue  [q] Quit")
	return s
}

func Run() error {
	p := tea.NewProgram(initialModel())
	_, err := p.Run()
	return err
}
