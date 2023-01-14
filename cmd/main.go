package main

import (
	"fmt"
	"kklogTUI/constant"
	"kklogTUI/dto"
	"kklogTUI/utils"
	"log"
	"os"
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
)

var (
	Deployments 		[]*dto.Deployment = []*dto.Deployment{}
	DeploymentPodsMap   map[string][]*dto.Pod = map[string][]*dto.Pod{}
)

func init() {
	tea.LogToFile("./tea.debug.log", "debug")
	log.Println("-------------------------------------------------------------")

	deploymentSet := utils.NewSet[*dto.Deployment]()

	for i := range constant.Pods {
		pod := constant.Pods[i]
		deploymentSet.Add(pod.Deployment)
		DeploymentPodsMap[pod.Deployment.Name] = append(DeploymentPodsMap[pod.Deployment.Name], pod)
	}

	for i := range deploymentSet.Elems() {
		Deployments = append(Deployments, i)
	}
}

type model struct {
	selectedEnv 		*dto.Env
	selectedDeployment 	*dto.Deployment
	selectedPod 		*dto.Pod
	selectedNs  		*dto.Namespace

	state	 			dto.State
	choices  			dto.Choices
    cursor   			int
}

func initialModel() model {
	return model{
		state:    		constant.StateChooseEnv,
		choices: 		dto.Choices{},
		cursor:   		0,
	}
}

func (m *model) clearChoices() {
	m.choices = constant.ChoicesEmpty
}

func (m *model) setChoicesEnv() {
	m.clearChoices()
	for _, v := range constant.Envs {
		m.choices = append(m.choices, v.String())
	}
}

func (m *model) setChoicesDep() {
	m.clearChoices()
	for i := range Deployments {
		m.choices = append(m.choices, Deployments[i].String())
	}
}

func (m *model) setChoicesPod() {
	m.clearChoices()
	for i := range DeploymentPodsMap[m.selectedDeployment.Name] {
		m.choices = append(m.choices, DeploymentPodsMap[m.selectedDeployment.Name][i].String())
	}
}

func (m *model) setChoicesNsp() {
	m.clearChoices()
	for i := range constant.DevNsSlice {
		m.choices = append(m.choices, constant.DevNsSlice[i].String())
	}
}

func (m model) dispalyLog() tea.Cmd {
	cmd := exec.Command(
		"kklog",
		"-d",
		m.selectedDeployment.Name,
		"-e",
		m.selectedEnv.Name,
		"-n",
		m.selectedPod.Name,
		"-ns",
		m.selectedNs.Name,
		"-t",
		string(m.selectedPod.Type),
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
		return dto.LogFinishMsg{Err: err}
	})
}

func (m model) Init() tea.Cmd {
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	log.Printf("Current State: %d", m.state)
	switch m.state {
	case constant.StateChooseEnv:
		m.setChoicesEnv()
	case constant.StateChooseDep:
		m.setChoicesDep()
	case constant.StateChoosePod:
		m.setChoicesPod()
	case constant.StateChooseNsp:
		m.setChoicesNsp()
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
			case constant.StateChooseEnv:
				m.selectedEnv = constant.Envs[m.cursor]
				m.state = constant.StateChooseDep
				log.Printf("m.selectedEnv: %v\n", m.selectedEnv)
			case constant.StateChooseDep:
				m.selectedDeployment = Deployments[m.cursor]
				m.state = constant.StateChoosePod
				log.Printf("m.selectedDeployment: %v\n", m.selectedDeployment)
			case constant.StateChoosePod:
				m.selectedPod = DeploymentPodsMap[m.selectedDeployment.Name][m.cursor]
				if m.selectedEnv.IsProd() {
					m.state = constant.StateDispLog
					m.selectedNs = m.selectedPod.Deployment.ProdNamespace
				} else {
					m.state = constant.StateChooseNsp
				}
				log.Printf("m.selectedPod: %v\n", m.selectedPod)
			case constant.StateChooseNsp:
				m.selectedNs = constant.DevNsSlice[m.cursor]
				m.state = constant.StateDispLog
				log.Printf("m.selectedNs: %v\n", m.selectedNs)
			case constant.StateDispLog:
				fmt.Printf(
					"Your choices is :\n\n[%s] [%s] [%s] [%s] [%s]\n\n",
					m.selectedEnv.String(),
					m.selectedNs.String(),
					m.selectedDeployment.String(),
					m.selectedPod.String(),
					m.selectedPod.Type,
				)
				return m, m.dispalyLog()
			}
			m.cursor = 0
        }
	case dto.LogFinishMsg:
		if msg.Err != nil {
			log.Printf("err: %s", msg.Err.Error())
		}
		return m, tea.Quit
    }

    return m, nil
}

func (m model) View() string {
    // The header
    s := ""

	switch m.state {
	case constant.StateChooseEnv:
		m.setChoicesEnv()
		s = "Choose Env:\n\n"
	case constant.StateChooseDep:
		m.setChoicesDep()
		s = "Choose Deployment:\n\n"
	case constant.StateChoosePod:
		m.setChoicesPod()
		s = "Choose Pod:\n\n"
	case constant.StateChooseNsp:
		m.setChoicesNsp()
		s = "Choose Namespace:\n\n"
	case constant.StateDispLog:
		m.clearChoices()
		s = fmt.Sprintf(
			"Your choices is :\n\n[%s] [%s] [%s] [%s] [%s]",
			m.selectedEnv.String(),
			m.selectedNs.String(),
			m.selectedDeployment.String(),
			m.selectedPod.String(),
			m.selectedPod.Type,
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

		m.state = constant.StateChooseNsp
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
