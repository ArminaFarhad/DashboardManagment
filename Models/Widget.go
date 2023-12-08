package Models

type Widget struct {
	Id         int
	Name       string
	WidgetType int
}

const (
	Calender = iota
)

type GetOrAddWidget struct {
	Name       string
	WidgetType int
}
