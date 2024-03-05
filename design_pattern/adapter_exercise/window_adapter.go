package main

type WindowAdapter struct {
	Com Windows
}

func (w WindowAdapter) InsertIntoLightningPort() {
	w.Com.insertIntoUSBPort()
}

func NewWindowAdapter() WindowAdapter {
	return WindowAdapter{Com: Windows{}}
}
