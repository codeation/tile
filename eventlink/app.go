package eventlink

type appFrame struct {
	Control
	Framer
}

// AppWithFrame combines an AppFrame with a specified Framer
func AppWithFrame(app Control, framer Framer) *appFrame {
	return &appFrame{
		Control: app,
		Framer:  framer,
	}
}

type linkApp struct {
	Linker
	AppFramer
}

// AppWithLink combines an AppFrame with the specified Linker
func AppWithLink(app AppFramer, linker Linker) *linkApp {
	return &linkApp{
		Linker:    linker,
		AppFramer: app,
	}
}
