package creational

type WidgetSetting struct{}

type Widget struct {
	widgetType string
	setting    WidgetSetting
}

func NewWidget(widgetType string, setting WidgetSetting) Widget {
	return Widget{
		widgetType: widgetType,
		setting:    setting,
	}
}

func (w *Widget) clone() Widget {
	return NewWidget(w.widgetType, w.setting)
}

func (w *Widget) customize(customSetting WidgetSetting) {
	w.setting = customSetting
}
