package ui


import (

ui "github.com/andlabs/ui"

)

func Load(DefaultInterfaces map[string]map[string]interface{}) {


if _, ok := DefaultInterfaces["ui"]; !ok {
   DefaultInterfaces["ui"] = map[string]interface{}{}
}

DefaultInterfaces["ui"]["align"] = func(isPtr bool) interface{} {
	if isPtr {
		return new(ui.Align)
	}					
	return *new(ui.Align)
}
DefaultInterfaces["ui"]["area"] = func(isPtr bool) interface{} {
	if isPtr {
		return new(ui.Area)
	}					
	return *new(ui.Area)
}
DefaultInterfaces["ui"]["areaDrawParams"] = func(isPtr bool) interface{} {
	if isPtr {
		return new(ui.AreaDrawParams)
	}					
	return *new(ui.AreaDrawParams)
}
DefaultInterfaces["ui"]["areaHandler"] = func(isPtr bool) interface{} {
	if isPtr {
		return new(ui.AreaHandler)
	}					
	return *new(ui.AreaHandler)
}
DefaultInterfaces["ui"]["areaKeyEvent"] = func(isPtr bool) interface{} {
	if isPtr {
		return new(ui.AreaKeyEvent)
	}					
	return *new(ui.AreaKeyEvent)
}
DefaultInterfaces["ui"]["areaMouseEvent"] = func(isPtr bool) interface{} {
	if isPtr {
		return new(ui.AreaMouseEvent)
	}					
	return *new(ui.AreaMouseEvent)
}
DefaultInterfaces["ui"]["at"] = func(isPtr bool) interface{} {
	if isPtr {
		return new(ui.At)
	}					
	return *new(ui.At)
}
DefaultInterfaces["ui"]["attribute"] = func(isPtr bool) interface{} {
	if isPtr {
		return new(ui.Attribute)
	}					
	return *new(ui.Attribute)
}
DefaultInterfaces["ui"]["attributedString"] = func(isPtr bool) interface{} {
	if isPtr {
		return new(ui.AttributedString)
	}					
	return *new(ui.AttributedString)
}
DefaultInterfaces["ui"]["box"] = func(isPtr bool) interface{} {
	if isPtr {
		return new(ui.Box)
	}					
	return *new(ui.Box)
}
DefaultInterfaces["ui"]["button"] = func(isPtr bool) interface{} {
	if isPtr {
		return new(ui.Button)
	}					
	return *new(ui.Button)
}
DefaultInterfaces["ui"]["checkbox"] = func(isPtr bool) interface{} {
	if isPtr {
		return new(ui.Checkbox)
	}					
	return *new(ui.Checkbox)
}
DefaultInterfaces["ui"]["colorButton"] = func(isPtr bool) interface{} {
	if isPtr {
		return new(ui.ColorButton)
	}					
	return *new(ui.ColorButton)
}
DefaultInterfaces["ui"]["combobox"] = func(isPtr bool) interface{} {
	if isPtr {
		return new(ui.Combobox)
	}					
	return *new(ui.Combobox)
}
DefaultInterfaces["ui"]["control"] = func(isPtr bool) interface{} {
	if isPtr {
		return new(ui.Control)
	}					
	return *new(ui.Control)
}
DefaultInterfaces["ui"]["controlBase"] = func(isPtr bool) interface{} {
	if isPtr {
		return new(ui.ControlBase)
	}					
	return *new(ui.ControlBase)
}
DefaultInterfaces["ui"]["dateTimePicker"] = func(isPtr bool) interface{} {
	if isPtr {
		return new(ui.DateTimePicker)
	}					
	return *new(ui.DateTimePicker)
}
DefaultInterfaces["ui"]["drawBrush"] = func(isPtr bool) interface{} {
	if isPtr {
		return new(ui.DrawBrush)
	}					
	return *new(ui.DrawBrush)
}
DefaultInterfaces["ui"]["drawBrushType"] = func(isPtr bool) interface{} {
	if isPtr {
		return new(ui.DrawBrushType)
	}					
	return *new(ui.DrawBrushType)
}
DefaultInterfaces["ui"]["drawContext"] = func(isPtr bool) interface{} {
	if isPtr {
		return new(ui.DrawContext)
	}					
	return *new(ui.DrawContext)
}
DefaultInterfaces["ui"]["drawFillMode"] = func(isPtr bool) interface{} {
	if isPtr {
		return new(ui.DrawFillMode)
	}					
	return *new(ui.DrawFillMode)
}
DefaultInterfaces["ui"]["drawGradientStop"] = func(isPtr bool) interface{} {
	if isPtr {
		return new(ui.DrawGradientStop)
	}					
	return *new(ui.DrawGradientStop)
}
DefaultInterfaces["ui"]["drawLineCap"] = func(isPtr bool) interface{} {
	if isPtr {
		return new(ui.DrawLineCap)
	}					
	return *new(ui.DrawLineCap)
}
DefaultInterfaces["ui"]["drawLineJoin"] = func(isPtr bool) interface{} {
	if isPtr {
		return new(ui.DrawLineJoin)
	}					
	return *new(ui.DrawLineJoin)
}
DefaultInterfaces["ui"]["drawMatrix"] = func(isPtr bool) interface{} {
	if isPtr {
		return new(ui.DrawMatrix)
	}					
	return *new(ui.DrawMatrix)
}
DefaultInterfaces["ui"]["drawPath"] = func(isPtr bool) interface{} {
	if isPtr {
		return new(ui.DrawPath)
	}					
	return *new(ui.DrawPath)
}
DefaultInterfaces["ui"]["drawStrokeParams"] = func(isPtr bool) interface{} {
	if isPtr {
		return new(ui.DrawStrokeParams)
	}					
	return *new(ui.DrawStrokeParams)
}
DefaultInterfaces["ui"]["drawTextAlign"] = func(isPtr bool) interface{} {
	if isPtr {
		return new(ui.DrawTextAlign)
	}					
	return *new(ui.DrawTextAlign)
}
DefaultInterfaces["ui"]["drawTextLayout"] = func(isPtr bool) interface{} {
	if isPtr {
		return new(ui.DrawTextLayout)
	}					
	return *new(ui.DrawTextLayout)
}
DefaultInterfaces["ui"]["drawTextLayoutParams"] = func(isPtr bool) interface{} {
	if isPtr {
		return new(ui.DrawTextLayoutParams)
	}					
	return *new(ui.DrawTextLayoutParams)
}
DefaultInterfaces["ui"]["editableCombobox"] = func(isPtr bool) interface{} {
	if isPtr {
		return new(ui.EditableCombobox)
	}					
	return *new(ui.EditableCombobox)
}
DefaultInterfaces["ui"]["entry"] = func(isPtr bool) interface{} {
	if isPtr {
		return new(ui.Entry)
	}					
	return *new(ui.Entry)
}
DefaultInterfaces["ui"]["extKey"] = func(isPtr bool) interface{} {
	if isPtr {
		return new(ui.ExtKey)
	}					
	return *new(ui.ExtKey)
}
DefaultInterfaces["ui"]["fontButton"] = func(isPtr bool) interface{} {
	if isPtr {
		return new(ui.FontButton)
	}					
	return *new(ui.FontButton)
}
DefaultInterfaces["ui"]["fontDescriptor"] = func(isPtr bool) interface{} {
	if isPtr {
		return new(ui.FontDescriptor)
	}					
	return *new(ui.FontDescriptor)
}
DefaultInterfaces["ui"]["form"] = func(isPtr bool) interface{} {
	if isPtr {
		return new(ui.Form)
	}					
	return *new(ui.Form)
}
DefaultInterfaces["ui"]["grid"] = func(isPtr bool) interface{} {
	if isPtr {
		return new(ui.Grid)
	}					
	return *new(ui.Grid)
}
DefaultInterfaces["ui"]["group"] = func(isPtr bool) interface{} {
	if isPtr {
		return new(ui.Group)
	}					
	return *new(ui.Group)
}
DefaultInterfaces["ui"]["image"] = func(isPtr bool) interface{} {
	if isPtr {
		return new(ui.Image)
	}					
	return *new(ui.Image)
}
DefaultInterfaces["ui"]["label"] = func(isPtr bool) interface{} {
	if isPtr {
		return new(ui.Label)
	}					
	return *new(ui.Label)
}
DefaultInterfaces["ui"]["modifiers"] = func(isPtr bool) interface{} {
	if isPtr {
		return new(ui.Modifiers)
	}					
	return *new(ui.Modifiers)
}
DefaultInterfaces["ui"]["multilineEntry"] = func(isPtr bool) interface{} {
	if isPtr {
		return new(ui.MultilineEntry)
	}					
	return *new(ui.MultilineEntry)
}
DefaultInterfaces["ui"]["openTypeFeatures"] = func(isPtr bool) interface{} {
	if isPtr {
		return new(ui.OpenTypeFeatures)
	}					
	return *new(ui.OpenTypeFeatures)
}
DefaultInterfaces["ui"]["openTypeTag"] = func(isPtr bool) interface{} {
	if isPtr {
		return new(ui.OpenTypeTag)
	}					
	return *new(ui.OpenTypeTag)
}
DefaultInterfaces["ui"]["progressBar"] = func(isPtr bool) interface{} {
	if isPtr {
		return new(ui.ProgressBar)
	}					
	return *new(ui.ProgressBar)
}
DefaultInterfaces["ui"]["radioButtons"] = func(isPtr bool) interface{} {
	if isPtr {
		return new(ui.RadioButtons)
	}					
	return *new(ui.RadioButtons)
}
DefaultInterfaces["ui"]["separator"] = func(isPtr bool) interface{} {
	if isPtr {
		return new(ui.Separator)
	}					
	return *new(ui.Separator)
}
DefaultInterfaces["ui"]["slider"] = func(isPtr bool) interface{} {
	if isPtr {
		return new(ui.Slider)
	}					
	return *new(ui.Slider)
}
DefaultInterfaces["ui"]["spinbox"] = func(isPtr bool) interface{} {
	if isPtr {
		return new(ui.Spinbox)
	}					
	return *new(ui.Spinbox)
}
DefaultInterfaces["ui"]["tab"] = func(isPtr bool) interface{} {
	if isPtr {
		return new(ui.Tab)
	}					
	return *new(ui.Tab)
}
DefaultInterfaces["ui"]["table"] = func(isPtr bool) interface{} {
	if isPtr {
		return new(ui.Table)
	}					
	return *new(ui.Table)
}
DefaultInterfaces["ui"]["tableColor"] = func(isPtr bool) interface{} {
	if isPtr {
		return new(ui.TableColor)
	}					
	return *new(ui.TableColor)
}
DefaultInterfaces["ui"]["tableImage"] = func(isPtr bool) interface{} {
	if isPtr {
		return new(ui.TableImage)
	}					
	return *new(ui.TableImage)
}
DefaultInterfaces["ui"]["tableInt"] = func(isPtr bool) interface{} {
	if isPtr {
		return new(ui.TableInt)
	}					
	return *new(ui.TableInt)
}
DefaultInterfaces["ui"]["tableModel"] = func(isPtr bool) interface{} {
	if isPtr {
		return new(ui.TableModel)
	}					
	return *new(ui.TableModel)
}
DefaultInterfaces["ui"]["tableModelHandler"] = func(isPtr bool) interface{} {
	if isPtr {
		return new(ui.TableModelHandler)
	}					
	return *new(ui.TableModelHandler)
}
DefaultInterfaces["ui"]["tableParams"] = func(isPtr bool) interface{} {
	if isPtr {
		return new(ui.TableParams)
	}					
	return *new(ui.TableParams)
}
DefaultInterfaces["ui"]["tableString"] = func(isPtr bool) interface{} {
	if isPtr {
		return new(ui.TableString)
	}					
	return *new(ui.TableString)
}
DefaultInterfaces["ui"]["tableTextColumnOptionalParams"] = func(isPtr bool) interface{} {
	if isPtr {
		return new(ui.TableTextColumnOptionalParams)
	}					
	return *new(ui.TableTextColumnOptionalParams)
}
DefaultInterfaces["ui"]["tableValue"] = func(isPtr bool) interface{} {
	if isPtr {
		return new(ui.TableValue)
	}					
	return *new(ui.TableValue)
}
DefaultInterfaces["ui"]["textBackground"] = func(isPtr bool) interface{} {
	if isPtr {
		return new(ui.TextBackground)
	}					
	return *new(ui.TextBackground)
}
DefaultInterfaces["ui"]["textColor"] = func(isPtr bool) interface{} {
	if isPtr {
		return new(ui.TextColor)
	}					
	return *new(ui.TextColor)
}
DefaultInterfaces["ui"]["textFamily"] = func(isPtr bool) interface{} {
	if isPtr {
		return new(ui.TextFamily)
	}					
	return *new(ui.TextFamily)
}
DefaultInterfaces["ui"]["textItalic"] = func(isPtr bool) interface{} {
	if isPtr {
		return new(ui.TextItalic)
	}					
	return *new(ui.TextItalic)
}
DefaultInterfaces["ui"]["textSize"] = func(isPtr bool) interface{} {
	if isPtr {
		return new(ui.TextSize)
	}					
	return *new(ui.TextSize)
}
DefaultInterfaces["ui"]["textStretch"] = func(isPtr bool) interface{} {
	if isPtr {
		return new(ui.TextStretch)
	}					
	return *new(ui.TextStretch)
}
DefaultInterfaces["ui"]["textWeight"] = func(isPtr bool) interface{} {
	if isPtr {
		return new(ui.TextWeight)
	}					
	return *new(ui.TextWeight)
}
DefaultInterfaces["ui"]["underline"] = func(isPtr bool) interface{} {
	if isPtr {
		return new(ui.Underline)
	}					
	return *new(ui.Underline)
}
DefaultInterfaces["ui"]["underlineColor"] = func(isPtr bool) interface{} {
	if isPtr {
		return new(ui.UnderlineColor)
	}					
	return *new(ui.UnderlineColor)
}
DefaultInterfaces["ui"]["underlineColorCustom"] = func(isPtr bool) interface{} {
	if isPtr {
		return new(ui.UnderlineColorCustom)
	}					
	return *new(ui.UnderlineColorCustom)
}
DefaultInterfaces["ui"]["window"] = func(isPtr bool) interface{} {
	if isPtr {
		return new(ui.Window)
	}					
	return *new(ui.Window)
}
DefaultInterfaces["ui"]["LibuiFreeText"] = ui.LibuiFreeText
DefaultInterfaces["ui"]["Main"] = ui.Main
DefaultInterfaces["ui"]["MsgBox"] = ui.MsgBox
DefaultInterfaces["ui"]["MsgBoxError"] = ui.MsgBoxError
DefaultInterfaces["ui"]["OnShouldQuit"] = ui.OnShouldQuit
DefaultInterfaces["ui"]["OpenFile"] = ui.OpenFile
DefaultInterfaces["ui"]["QueueMain"] = ui.QueueMain
DefaultInterfaces["ui"]["Quit"] = ui.Quit
DefaultInterfaces["ui"]["SaveFile"] = ui.SaveFile
DefaultInterfaces["ui"]["NewArea"] = ui.NewArea
DefaultInterfaces["ui"]["NewScrollingArea"] = ui.NewScrollingArea
DefaultInterfaces["ui"]["NewAttributedString"] = ui.NewAttributedString
DefaultInterfaces["ui"]["NewHorizontalBox"] = ui.NewHorizontalBox
DefaultInterfaces["ui"]["NewVerticalBox"] = ui.NewVerticalBox
DefaultInterfaces["ui"]["NewButton"] = ui.NewButton
DefaultInterfaces["ui"]["NewCheckbox"] = ui.NewCheckbox
DefaultInterfaces["ui"]["NewColorButton"] = ui.NewColorButton
DefaultInterfaces["ui"]["NewCombobox"] = ui.NewCombobox
DefaultInterfaces["ui"]["ControlFromLibui"] = ui.ControlFromLibui
DefaultInterfaces["ui"]["NewControlBase"] = ui.NewControlBase
DefaultInterfaces["ui"]["NewDatePicker"] = ui.NewDatePicker
DefaultInterfaces["ui"]["NewDateTimePicker"] = ui.NewDateTimePicker
DefaultInterfaces["ui"]["NewTimePicker"] = ui.NewTimePicker
DefaultInterfaces["ui"]["DrawNewMatrix"] = ui.DrawNewMatrix
DefaultInterfaces["ui"]["DrawNewPath"] = ui.DrawNewPath
DefaultInterfaces["ui"]["DrawNewTextLayout"] = ui.DrawNewTextLayout
DefaultInterfaces["ui"]["NewEditableCombobox"] = ui.NewEditableCombobox
DefaultInterfaces["ui"]["NewEntry"] = ui.NewEntry
DefaultInterfaces["ui"]["NewPasswordEntry"] = ui.NewPasswordEntry
DefaultInterfaces["ui"]["NewSearchEntry"] = ui.NewSearchEntry
DefaultInterfaces["ui"]["NewFontButton"] = ui.NewFontButton
DefaultInterfaces["ui"]["NewForm"] = ui.NewForm
DefaultInterfaces["ui"]["NewGrid"] = ui.NewGrid
DefaultInterfaces["ui"]["NewGroup"] = ui.NewGroup
DefaultInterfaces["ui"]["NewImage"] = ui.NewImage
DefaultInterfaces["ui"]["NewLabel"] = ui.NewLabel
DefaultInterfaces["ui"]["NewMultilineEntry"] = ui.NewMultilineEntry
DefaultInterfaces["ui"]["NewNonWrappingMultilineEntry"] = ui.NewNonWrappingMultilineEntry
DefaultInterfaces["ui"]["ToOpenTypeTag"] = ui.ToOpenTypeTag
DefaultInterfaces["ui"]["NewProgressBar"] = ui.NewProgressBar
DefaultInterfaces["ui"]["NewRadioButtons"] = ui.NewRadioButtons
DefaultInterfaces["ui"]["NewHorizontalSeparator"] = ui.NewHorizontalSeparator
DefaultInterfaces["ui"]["NewVerticalSeparator"] = ui.NewVerticalSeparator
DefaultInterfaces["ui"]["NewSlider"] = ui.NewSlider
DefaultInterfaces["ui"]["NewSpinbox"] = ui.NewSpinbox
DefaultInterfaces["ui"]["NewTab"] = ui.NewTab
DefaultInterfaces["ui"]["NewTable"] = ui.NewTable
DefaultInterfaces["ui"]["NewTableModel"] = ui.NewTableModel
DefaultInterfaces["ui"]["NewWindow"] = ui.NewWindow

}
