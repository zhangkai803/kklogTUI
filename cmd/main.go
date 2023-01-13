package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
)

type State int
type Choices []string
type logFinishMsg struct {err error}

var (
	StateChooseNul State = 0
	StateChooseEnv State = 1
	StateChooseDep State = 2
	StateChoosePod State = 3
	StateChooseNsp State = 4
	StateChooseTyp State = 5
	StateDispLog   State = 6

	DeploymentPodsMap map[string]Choices = map[string]Choices{
		"wk-tag-manage": {
			"wk-tag-manage",
		},
		"wk-miniprogram-cms": {
			"wk-miniprogram-cms",
			"wk-miniprogram-cms-async-task",
		},
	}
	EnvNamespacesMap map[string]Choices = map[string]Choices{
		"dev": {"sit", "dev1"},
		"prod": {"production", "iprod"},
	}
	ChoicesPodTypes 	Choices = Choices{"api", "script"}
	ChoicesDeployments	Choices = Choices{}
	ChoicesEnv 			Choices = Choices{}
	ChoicesNul 			Choices = Choices{}
)

func init() {
	tea.LogToFile("./tea.debug.log", "debug")
	log.Println("-------------------------------------------------------------")

	for k := range DeploymentPodsMap {
		ChoicesDeployments = append(ChoicesDeployments, k)
	}
	for k := range EnvNamespacesMap {
		ChoicesEnv = append(ChoicesEnv, k)
	}

	// log.Println(ChoicesEnv)
	// log.Println(ChoicesDeployments)
	// log.Println(ChoicesPodTypes)
	// log.Panicln(ChoicesEnv)
	// log.Panicln(ChoicesEnv)
}

type model struct {
	state	 			State

	selectedEnv 		string
	selectedNs  		string
	selectedDeployment 	string
	selectedPod 		string
	selectedPodType 	string

	choices  			Choices			// items on the to-do list
    cursor   			int             // which to-do list item our cursor is pointing at
	selectedIdx 		int
}

func initialModel() model {
	return model{
		state:    StateChooseEnv,
		choices: Choices{},
		selectedIdx: -1,
		cursor:   0,
	}
}

func (m model) dispalyLog() tea.Cmd {
	cmd := exec.Command(
		"kklog",
		"-d",
		m.selectedDeployment,
		"-e",
		m.selectedEnv,
		"-n",
		m.selectedPod,
		"-ns",
		m.selectedNs,
		"-t",
		m.selectedPodType,
		"-l",
		"500",
		"2>",
		"/tmp/kklog_grep_buf_`date +%s`",
		"|",
		"tail",
		"-f",
		"/tmp/kklog_grep_buf_`date +%s`",
	)
	fmt.Printf("cmd: %v\n", cmd.String())

	return tea.ExecProcess(cmd, func(err error) tea.Msg {
		return logFinishMsg{err: err}
	})
}

func (m model) Init() tea.Cmd {
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	log.Printf("Current State: %d", m.state)
	switch m.state {
	case StateChooseEnv:
		m.choices = ChoicesEnv
	case StateChooseNsp:
		m.choices = EnvNamespacesMap[m.selectedEnv]
	case StateChooseDep:
		m.choices = ChoicesDeployments
	case StateChoosePod:
		m.choices = DeploymentPodsMap[m.selectedDeployment]
	case StateChooseTyp:
		m.choices = ChoicesPodTypes
	}

    switch msg := msg.(type) {

    // Is it a key press?
    case tea.KeyMsg:

        // Cool, what was the actual key pressed?
        switch msg.String() {

        // These keys should exit the program.
        case "ctrl+c", "q":
            return m, tea.Quit

        // The "up" and "k" keys move the cursor up
        case "up", "k":
            if m.cursor > 0 {
                m.cursor--
            }

        // The "down" and "j" keys move the cursor down
        case "down", "j":
            if m.cursor < len(m.choices)-1 {
                m.cursor++
            }

        // The "enter" key and the spacebar (a literal space) toggle
        // the selected state for the item that the cursor is pointing at.
        case "enter", " ":
			switch m.state {
			case StateChooseEnv:
				m.selectedEnv = ChoicesEnv[m.cursor]
				log.Printf("m.selectedEnv: %v\n", m.selectedEnv)
				m.state = StateChooseNsp
			case StateChooseNsp:
				m.selectedNs = EnvNamespacesMap[m.selectedEnv][m.cursor]
				log.Printf("m.selectedNs: %v\n", m.selectedNs)
				m.state = StateChooseDep
			case StateChooseDep:
				m.selectedDeployment = ChoicesDeployments[m.cursor]
				log.Printf("m.selectedDeployment: %v\n", m.selectedDeployment)
				m.state = StateChoosePod
			case StateChoosePod:
				m.selectedPod = DeploymentPodsMap[m.selectedDeployment][m.cursor]
				log.Printf("m.selectedPod: %v\n", m.selectedPod)
				m.state = StateChooseTyp
			case StateChooseTyp:
				m.selectedPodType = ChoicesPodTypes[m.cursor]
				log.Printf("m.selectedPodType: %v\n", m.selectedPodType)
				m.state = StateDispLog
			case StateDispLog:
				fmt.Printf(
					"Your choices is :\n\n[%s] [%s] [%s] [%s] [%s]\n\n",
					m.selectedEnv,
					m.selectedNs,
					m.selectedDeployment,
					m.selectedPod,
					m.selectedPodType,
				)
				return m, m.dispalyLog()
			}

        }
	case logFinishMsg:
		if msg.err != nil {
			log.Printf("err: %s", msg.err.Error())
		}
		return m, tea.Quit
    }

    // Return the updated model to the Bubble Tea runtime for processing.
    // Note that we're not returning a command.
    return m, nil
}

func (m model) View() string {
    // The header
    s := ""

	switch m.state {
	case StateChooseEnv:
		m.choices = ChoicesEnv
		s = "Choose Env:\n\n"
	case StateChooseNsp:
		m.choices = EnvNamespacesMap[m.selectedEnv]
		s = "Choose Namespace:\n\n"
	case StateChooseDep:
		m.choices = ChoicesDeployments
		s = "Choose Deployment:\n\n"
	case StateChoosePod:
		m.choices = DeploymentPodsMap[m.selectedDeployment]
		s = "Choose Pod:\n\n"
	case StateChooseTyp:
		m.choices = ChoicesPodTypes
		s = "Choose Pod Type:\n\n"
	case StateDispLog:
		m.choices = ChoicesNul
		s = fmt.Sprintf(
			"Your choices is :\n\n[%s] [%s] [%s] [%s] [%s]",
			m.selectedEnv,
			m.selectedNs,
			m.selectedDeployment,
			m.selectedPod,
			m.selectedPodType,
		)
	}

	for i := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		s += " ["
		s += cursor
		s += "] "
		s += m.choices[i]
		s += "\n"

		m.state = StateChooseNsp
	}

	s += "\n\n"

    // The footer
    s += "\nPress q to quit.\n"

    // Send the UI for rendering
    return s
}

func main() {
	p := tea.NewProgram(initialModel())
    if _, err := p.Run(); err != nil {
        log.Printf("Alas, there's been an error: %v", err)
        os.Exit(1)
    }
}
