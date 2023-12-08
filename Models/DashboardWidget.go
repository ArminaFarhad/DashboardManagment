package main

type DashboardWidget struct {
	Id          int
	DashboardId int
	WidgetId    int
	position	int
}

type AddDashboardWidget struct {
	DashboardId int
	WidgetId    int
	position	int
}
