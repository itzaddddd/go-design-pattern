package behavior

import "fmt"

// momento
type EditorMomento struct {
	state string
}

func NewEditoMomento() *EditorMomento {
	return &EditorMomento{}
}

func (e *EditorMomento) getState() string {
	return e.state
}

// originator
type Editor struct {
	content string
}

func NewEditor() *Editor {
	return &Editor{}
}

func (e *Editor) setContent(content string) {
	e.content = content
}

func (e *Editor) getContent() string {
	return e.content
}

func (e *Editor) save() *EditorMomento {
	return &EditorMomento{
		state: e.content,
	}
}

func (e *Editor) restore(momento EditorMomento) {
	e.content = momento.getState()
}

// caretaker
type History struct {
	momentos []EditorMomento
}

func NewHistory() *History {
	return &History{}
}

func (h *History) addMomento(momento EditorMomento) {
	h.momentos = append(h.momentos, momento)
}

func (h *History) getMomento(index int) EditorMomento {
	return h.momentos[index]
}

func RunMomento() {
	editor := NewEditor()
	history := NewHistory()

	editor.setContent("First")
	history.addMomento(*editor.save())

	editor.setContent("Second")
	history.addMomento(*editor.save())

	editor.setContent("Third")
	history.addMomento(*editor.save())

	fmt.Printf("editor conent: %s\n", editor.getContent())

	// undo to second
	editor.restore(history.getMomento(1))
	fmt.Printf("editor conent: %s\n", editor.getContent())

}
