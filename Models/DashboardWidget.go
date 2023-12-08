package Models

type DashboardWidget struct {
	Id          int
	DashboardId int
	WidgetId    int
	Position    int
}

type AddDashboardWidget struct {
	DashboardId int
	WidgetId    int
	Position    int
}
