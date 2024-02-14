package prbf2

import "os"

type StateTransition func(*PRBF2)

var stateTransition = map[State]map[State]StateTransition{
	StateStopping: {
		StateStopped: stop,
	},
	StateStarting: {
		StateRunning: start,
	},
	StateAdopting: {
		StateRunning: adopt,
	},
	StateRunning: {
		StateExited: onExit,
	},
	StateRestarting: {
		StateRunning: restart,
	},
	StateExited: {
		StateRunning: restartOnExit,
	},
	StateCleaningError: {
		StateStopped: cleanError,
	},
}

var actionTransitions = map[State]map[Action][2]State{
	StateStopped: {
		ActionStart: [2]State{StateStarting, StateRunning},
		ActionAdopt: [2]State{StateAdopting, StateRunning},
	},
	StateRunning: {
		ActionStop:    [2]State{StateStopping, StateStopped},
		ActionRestart: [2]State{StateRestarting, StateRunning},
	},
	StateErrored: {
		ActionStop: [2]State{StateCleaningError, StateStopped},
	},
}

func stop(p *PRBF2) {
	sanityStop(p)

	p.currentState = StateStopped
}

func start(p *PRBF2) {
	proc, err := startProcess(p.path)
	if err != nil {
		p.err = err
		p.currentState = StateErrored
		return
	}

	p.proc = proc
	p.watchProcess()
	p.currentState = StateRunning
}

func adopt(p *PRBF2) {
	sanityStop(p)

	p.proc = p.toAdopt
	p.toAdopt = nil

	p.watchProcess()
	p.currentState = StateRunning
}

func restart(p *PRBF2) {
	sanityStop(p)

	var proc *os.Process
	var err error

	for i := 0; i < 3; i++ {
		proc, err = startProcess(p.path)
		if err == nil {
			break
		}
	}

	if err != nil {
		p.err = err
		p.currentState = StateErrored
		return
	}

	p.proc = proc
	p.watchProcess()
	p.currentState = StateRunning
}

func onExit(p *PRBF2) {
	sanityStop(p)

	err := p.restartCtx.inc()
	if err != nil {
		p.err = err
		p.currentState = StateErrored
		return
	}

	p.currentState = StateExited
}

func restartOnExit(p *PRBF2) {
	sanityStop(p)

	start(p)
}

func cleanError(p *PRBF2) {
	p.err = nil
	p.currentState = StateStopped
}

func sanityStop(p *PRBF2) {
	if p.watchProcessCancel != nil {
		p.watchProcessCancel()
	}

	if p.proc != nil {
		_ = stopProcess(p.proc)
	}

	p.proc = nil
}
