package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Trains App")

	// Widgets
	label := widget.NewLabel("Приложение \"Электрички\" в Великий Новгород")

	btn1 := widget.NewButton("Узнать расписание на 02.05.2022", func() {
		text_rasp := widget.NewLabel("Расписание электричек СПБ - Великий Новгород на 02.05.2022")
		train1 := widget.NewCard(
			"7001 Санкт-Петербург — Великий Новгород",
			"\"Ласточка\"	СЗППК",
			widget.NewLabel("Время: 07:26 - 10:17 (2 ч 51 мин)"),
		)
		train3 := widget.NewCard(
			"7005 Санкт-Петербург (Московский вокзал) — Великий Новгород",
			"\"Ласточка\"	СЗППК",
			widget.NewLabel("Время: 11:14 - 13:52 (2 ч 38 мин)"),
		)
		train4 := widget.NewCard(
			"7003 Санкт-Петербург (Московский вокзал) — Великий Новгород",
			"\"Ласточка\"	СЗППК",
			widget.NewLabel("Время: 19:32 - 23:09 (3 ч 37 мин)"),
		)
		rasp_obratno_btn := widget.NewButton("Расписание обратно (Великий Новгород - СПБ) на 02.05.2022", func() {
			textObr_rasp := widget.NewLabel("Расписание электричек Великий Новгород - СПБ на 02.05.2022")
			trainObr1 := widget.NewCard(
				"7002 Великий Новгород — Санкт-Петербург (Московский вокзал)",
				"\"Ласточка\"	СЗППК",
				widget.NewLabel("Время: 06:10 - 09:44 (3 ч 34 мин)"),
			)
			trainObr2 := widget.NewCard(
				"7004 Великий Новгород — Санкт-Петербург (Московский вокзал)",
				"\"Ласточка\"	СЗППК",
				widget.NewLabel("Время: 18:06 - 20:55 (2 ч 49 мин)"),
			)
			trainObr3 := widget.NewCard(
				"7006 Великий Новгород — Санкт-Петербург (Московский вокзал)",
				"\"Ласточка\"	СЗППК",
				widget.NewLabel("Время: 20:08 - 23:30 (3 ч 22 мин)"),
			)
			w.SetContent(container.NewVBox(
				textObr_rasp,
				trainObr1,
				trainObr2,
				trainObr3,
			))
		})
		w.SetContent(container.NewVBox(
			text_rasp,
			train1,
			train3,
			train4,
			rasp_obratno_btn,
		))
	})

	btn2 := widget.NewButton("Узнать расписание на 07.05.2022", func() {
		text_rasp := widget.NewLabel("Расписание электричек СПБ - Великий Новгород на 07.05.2022")
		train1 := widget.NewCard(
			"7001 Санкт-Петербург — Великий Новгород",
			"\"Ласточка\"	СЗППК",
			widget.NewLabel("Время: 07:26 - 10:17 (2 ч 51 мин)"),
		)
		train2 := widget.NewCard(
			"6685 Санкт-Петербург (Витебский вокзал) — Великий Новгород",
			"состав 2-3 вагона	СЗППК",
			widget.NewLabel("Время: 07:48 - 11:49 (4 ч 1 мин)"),
		)
		train3 := widget.NewCard(
			"7005 Санкт-Петербург (Московский вокзал) — Великий Новгород",
			"\"Ласточка\"	СЗППК",
			widget.NewLabel("Время: 11:14 - 13:52 (2 ч 38 мин)"),
		)
		train4 := widget.NewCard(
			"7003 Санкт-Петербург (Московский вокзал) — Великий Новгород",
			"\"Ласточка\"	СЗППК",
			widget.NewLabel("Время: 19:32 - 23:09 (3 ч 37 мин)"),
		)
		rasp_obratno_btn := widget.NewButton("Расписание обратно (Великий Новгород - СПБ) на 07.05.2022", func() {
			textObr_rasp := widget.NewLabel("Расписание электричек Великий Новгород - СПБ на 07.05.2022")
			trainObr1 := widget.NewCard(
				"7002 Великий Новгород — Санкт-Петербург (Московский вокзал)",
				"\"Ласточка\"	СЗППК",
				widget.NewLabel("Время: 06:10 - 09:44 (3 ч 34 мин)"),
			)
			trainObr2 := widget.NewCard(
				"7004 Великий Новгород — Санкт-Петербург (Московский вокзал)",
				"\"Ласточка\"	СЗППК",
				widget.NewLabel("Время: 18:06 - 20:55 (2 ч 49 мин)"),
			)
			trainObr3 := widget.NewCard(
				"7006 Великий Новгород — Санкт-Петербург (Московский вокзал)",
				"\"Ласточка\"	СЗППК",
				widget.NewLabel("Время: 20:08 - 23:30 (3 ч 22 мин)"),
			)
			w.SetContent(container.NewVBox(
				textObr_rasp,
				trainObr1,
				trainObr2,
				trainObr3,
			))
		})
		w.SetContent(container.NewVBox(
			text_rasp,
			train1,
			train2,
			train3,
			train4,
			rasp_obratno_btn,
		))
	})

	btn3 := widget.NewButton("Узнать расписание на 08.05.2022", func() {
		text_rasp := widget.NewLabel("Расписание электричек СПБ - Великий Новгород на 08.05.2022")
		train1 := widget.NewCard(
			"7001 Санкт-Петербург — Великий Новгород",
			"\"Ласточка\"	СЗППК",
			widget.NewLabel("Время: 07:26 - 10:17 (2 ч 51 мин)"),
		)
		train2 := widget.NewCard(
			"6685 Санкт-Петербург (Витебский вокзал) — Великий Новгород",
			"состав 2-3 вагона	СЗППК",
			widget.NewLabel("Время: 07:48 - 11:49 (4 ч 1 мин)"),
		)
		train3 := widget.NewCard(
			"7005 Санкт-Петербург (Московский вокзал) — Великий Новгород",
			"\"Ласточка\"	СЗППК",
			widget.NewLabel("Время: 11:14 - 13:52 (2 ч 38 мин)"),
		)
		train4 := widget.NewCard(
			"7003 Санкт-Петербург (Московский вокзал) — Великий Новгород",
			"\"Ласточка\"	СЗППК",
			widget.NewLabel("Время: 19:32 - 23:09 (3 ч 37 мин)"),
		)
		rasp_obratno_btn := widget.NewButton("Расписание обратно (Великий Новгород - СПБ) на 08.05.2022", func() {
			textObr_rasp := widget.NewLabel("Расписание электричек Великий Новгород - СПБ на 08.05.2022")
			trainObr1 := widget.NewCard(
				"7002 Великий Новгород — Санкт-Петербург (Московский вокзал)",
				"\"Ласточка\"	СЗППК",
				widget.NewLabel("Время: 06:12 - 09:51 (3 ч 39 мин)"),
			)
			trainObr2 := widget.NewCard(
				"7004 Великий Новгород — Санкт-Петербург (Московский вокзал)",
				"\"Ласточка\"	СЗППК",
				widget.NewLabel("Время: 18:06 - 20:55 (2 ч 49 мин)"),
			)
			trainObr3 := widget.NewCard(
				"7006 Великий Новгород — Санкт-Петербург (Московский вокзал)",
				"\"Ласточка\"	СЗППК",
				widget.NewLabel("Время: 20:08 - 23:30 (3 ч 22 мин)"),
			)
			w.SetContent(container.NewVBox(
				textObr_rasp,
				trainObr1,
				trainObr2,
				trainObr3,
			))
		})
		w.SetContent(container.NewVBox(
			text_rasp,
			train1,
			train2,
			train3,
			train4,
			rasp_obratno_btn,
		))
	})

	w.SetContent(container.NewVBox(
		label,
		btn1,
		btn2,
		btn3,
	))

	w.ShowAndRun()
}
