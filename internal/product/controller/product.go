package controller

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/oreshkindev/profilss.ru-backend/common"
	"github.com/oreshkindev/profilss.ru-backend/internal/product/entity"
	"github.com/xuri/excelize/v2"
)

type ProductController struct {
	usecase entity.ProductUsecase
}

var (
	destination = os.Getenv("REMOTE_PATH")
)

func NewProductController(usecase entity.ProductUsecase) *ProductController {
	return &ProductController{
		usecase: usecase,
	}
}

func (controller *ProductController) Create(w http.ResponseWriter, r *http.Request) {
	entity := &entity.Product{}

	if err := render.DecodeJSON(r.Body, entity); err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	result, err := controller.usecase.Create(entity)
	if err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, result)
}

func (controller *ProductController) Find(w http.ResponseWriter, r *http.Request) {
	result, err := controller.usecase.Find()
	if err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, result)
}

func (controller *ProductController) First(w http.ResponseWriter, r *http.Request) {
	// get id from request
	id := chi.URLParam(r, "id")

	result, err := controller.usecase.First(id)
	if err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, result)
}

func (controller *ProductController) Update(w http.ResponseWriter, r *http.Request) {

	entry := &entity.Product{}

	log.Println(entry.Published)

	if err := render.DecodeJSON(r.Body, entry); err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	result, err := controller.usecase.Update(entry)
	if err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}
	log.Println(result.Published)

	render.JSON(w, r, result)
}

func (controller *ProductController) Delete(w http.ResponseWriter, r *http.Request) {
	// get id from request
	id := chi.URLParam(r, "id")

	err := controller.usecase.Delete(id)
	if err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, nil)
}

func (controller *ProductController) DumpExcel(w http.ResponseWriter, r *http.Request) {

	products, err := controller.usecase.Find()
	if err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	f, err := excelize.OpenFile(filepath.Join(destination, "/doc/", "Профиль-С прайс-лист.xlsx"))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	// Используем вспомогательную функцию для вставки шапки таблицы
	insertHeader := func(j *int) {
		for _, row := range [][]interface{}{
			{"Типоразмер", "Толщина стенки", "Вес 1 п/м", "Цена за 1 п/м c НДС", "Цена за 1 тонну с НДС"},
		} {
			cell, err := excelize.CoordinatesToCellName(1, *j)
			if err != nil {
				fmt.Println(err)
				return
			}
			f.SetSheetRow("Страница 1", cell, &row)
			*j++
		}
	}

	var (
		headerStyle, boldStyle, defaultStyle, sizeStyle int
	)

	if headerStyle, err = f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{Vertical: "center", Horizontal: "center", Indent: 1},
		Font:      &excelize.Font{Size: 14},
		Border: []excelize.Border{
			{Type: "left", Color: "#000000", Style: 1},
			{Type: "right", Color: "#000000", Style: 1},
			{Type: "top", Color: "#000000", Style: 1},
			{Type: "bottom", Color: "#000000", Style: 1},
		},
	}); err != nil {
		fmt.Println(err)
	}

	if boldStyle, err = f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{Vertical: "center", Horizontal: "center", Indent: 1},
		Font:      &excelize.Font{Bold: true, Size: 14},
	}); err != nil {
		fmt.Println(err)
	}

	if sizeStyle, err = f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{Vertical: "center", Horizontal: "center", Indent: 1},
		Font:      &excelize.Font{Size: 24},
	}); err != nil {
		fmt.Println(err)
	}

	if defaultStyle, err = f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{Vertical: "center", Horizontal: "center", Indent: 1},
		Font:      &excelize.Font{Size: 14},
	}); err != nil {
		fmt.Println(err)
	}

	type Group struct {
		Name string
	}

	library := make(map[Group][]entity.Product)

	for _, p := range products {
		category := Group{p.Category.Name}
		library[category] = append(library[category], p)
	}

	// Начинаем с первой строки
	j := 1

	for category, products := range library {
		// Пропускаем строку для отступа
		j++

		// Выполняем объединение ячеек
		if err := f.MergeCell("Страница 1", "A"+strconv.Itoa(j), "E"+strconv.Itoa(j)); err != nil {
			fmt.Println(err)
		}
		if err := f.SetRowHeight("Страница 1", j, 30); err != nil {
			fmt.Println(err)
		}
		if err = f.SetCellStyle("Страница 1", "A"+strconv.Itoa(j), "B"+strconv.Itoa(j), defaultStyle); err != nil {
			fmt.Println(err)
		}

		// Вставляем название категории в колонку A
		if err = f.SetCellValue("Страница 1", "A"+strconv.Itoa(j), products[0].Category.SubCategory.Name); err != nil {
			fmt.Println(err)
		}

		// Вставляем название категории и дату только для первой категории
		if j == 3 {
			// Выполняем объединение ячеек
			if err := f.MergeCell("Страница 1", "A"+strconv.Itoa(j), "E"+strconv.Itoa(j)); err != nil {
				fmt.Println(err)
			}
			if err := f.SetRowHeight("Страница 1", j, 30); err != nil {
				fmt.Println(err)
			}
			if err = f.SetCellStyle("Страница 1", "A"+strconv.Itoa(j), "B"+strconv.Itoa(j), boldStyle); err != nil {
				fmt.Println(err)
			}

			// Здесь мы вставляем список ISO в поле "ГОСТ" без дублирования
			isoSet := make(map[string]struct{})

			for _, p := range products {
				for _, iso := range p.Category.Iso {
					isoSet[iso.Name] = struct{}{}
				}
			}

			isoList := ""
			for isoName := range isoSet {
				isoList += isoName + ", "
			}
			// Удаляем последнюю запятую
			isoList = strings.TrimSuffix(isoList, ", ")

			if err = f.SetCellValue("Страница 1", "A"+strconv.Itoa(j), isoList); err != nil {
				fmt.Println(err)
				return
			}

			// Вставляем название категории в колонку A
			if err = f.SetCellValue("Страница 1", "A"+strconv.Itoa(j), category.Name+" "+isoList); err != nil {
				fmt.Println(err)
			}
			// Пропускаем строку для отступа
			j++

		} else {
			// Пропускаем строку для отступа
			j++

			// Выполняем объединение ячеек
			if err := f.MergeCell("Страница 1", "A"+strconv.Itoa(j), "E"+strconv.Itoa(j)); err != nil {
				fmt.Println(err)
			}
			if err := f.SetRowHeight("Страница 1", j, 30); err != nil {
				fmt.Println(err)
			}
			if err = f.SetCellStyle("Страница 1", "A"+strconv.Itoa(j), "B"+strconv.Itoa(j), boldStyle); err != nil {
				fmt.Println(err)
			}

			// Здесь мы вставляем список ISO в поле "ГОСТ" без дублирования
			isoSet := make(map[string]struct{})

			for _, p := range products {
				for _, iso := range p.Category.Iso {
					isoSet[iso.Name] = struct{}{}
				}
			}

			isoList := ""
			for isoName := range isoSet {
				isoList += isoName + ", "
			}
			// Удаляем последнюю запятую
			isoList = strings.TrimSuffix(isoList, ", ")

			if err = f.SetCellValue("Страница 1", "A"+strconv.Itoa(j), isoList); err != nil {
				fmt.Println(err)
				return
			}

			// Вставляем название категории в колонку A
			if err = f.SetCellValue("Страница 1", "A"+strconv.Itoa(j), category.Name+" "+isoList); err != nil {
				fmt.Println(err)
			}
		}

		if err = f.SetRowHeight("Страница 1", j, 25); err != nil {
			fmt.Println(err)
		}
		j++ // Переходим к следующей строке для шапки таблицы

		if err = f.SetCellStyle("Страница 1", "A"+strconv.Itoa(j), "E"+strconv.Itoa(j), headerStyle); err != nil {
			fmt.Println(err)
		}

		if err := f.SetRowHeight("Страница 1", j, 40); err != nil {
			fmt.Println(err)
		}

		// Вставляем шапку таблицы
		insertHeader(&j)

		var currentSizeType string
		var mergeStart int

		for _, product := range products {
			// Вставляем данные о каждом товаре
			dataRow := []interface{}{
				product.Characteristic.Size,
				product.Characteristic.Thickness,
				product.Characteristic.Weight,
				product.Characteristic.Price,
				product.Characteristic.MaxPrice,
			}
			cell, err := excelize.CoordinatesToCellName(1, j)
			if err != nil {
				fmt.Println(err)
				return
			}
			if err = f.SetCellStyle("Страница 1", "A"+strconv.Itoa(j), "E"+strconv.Itoa(j), headerStyle); err != nil {
				fmt.Println(err)
			}
			if err = f.SetSheetRow("Страница 1", cell, &dataRow); err != nil {
				fmt.Println(err)
				return
			}

			// Check if the current product's size type is the same as the last one.
			if product.Characteristic.Size == currentSizeType {
				// If so, update the last row to be merged (this will be 'j').
			} else {
				// If it's not, merge the previous range if it's more than one row.
				if mergeStart > 0 && j-mergeStart > 1 {
					if err = f.SetCellStyle("Страница 1", "A"+strconv.Itoa(mergeStart), "A"+strconv.Itoa(j-1), defaultStyle); err != nil {
						fmt.Println(err)
					}
					if err := f.MergeCell("Страница 1", "A"+strconv.Itoa(mergeStart), "A"+strconv.Itoa(j-1)); err != nil {
						fmt.Println(err)
						return
					}
				}
				// Then, update the current size type and start a new merge range.
				currentSizeType = product.Characteristic.Size
				mergeStart = j
			}

			j++ // Переходим к следующей строке для данных товара
		}

		// Merge the final range after the loop if needed.
		if mergeStart > 0 && j-mergeStart > 1 {
			if err = f.SetCellStyle("Страница 1", "A"+strconv.Itoa(mergeStart), "A"+strconv.Itoa(j-1), defaultStyle); err != nil {
				fmt.Println(err)
			}
			if err := f.MergeCell("Страница 1", "A"+strconv.Itoa(mergeStart), "A"+strconv.Itoa(j-1)); err != nil {
				fmt.Println(err)
				return
			}
		}
	}

	j += 2 // Переходим к следующей строке для данных товара

	// Выполняем объединение ячеек
	if err := f.MergeCell("Страница 1", "A"+strconv.Itoa(j), "E"+strconv.Itoa(j)); err != nil {
		fmt.Println(err)
	}
	if err := f.SetRowHeight("Страница 1", j, 30); err != nil {
		fmt.Println(err)
	}
	if err = f.SetCellStyle("Страница 1", "A"+strconv.Itoa(j), "B"+strconv.Itoa(j), defaultStyle); err != nil {
		fmt.Println(err)
	}

	// Вставляем подвал таблицы
	if err = f.SetCellValue("Страница 1", "A"+strconv.Itoa(j), "Осуществляем доставку и порезку трубы!"); err != nil {
		fmt.Println(err)
	}

	j++ // Переходим к следующей строке для данных товара

	// Выполняем объединение ячеек
	if err := f.MergeCell("Страница 1", "A"+strconv.Itoa(j), "E"+strconv.Itoa(j)); err != nil {
		fmt.Println(err)
	}
	if err := f.SetRowHeight("Страница 1", j, 30); err != nil {
		fmt.Println(err)
	}
	if err = f.SetCellStyle("Страница 1", "A"+strconv.Itoa(j), "B"+strconv.Itoa(j), defaultStyle); err != nil {
		fmt.Println(err)
	}

	// Вставляем подвал таблицы
	if err = f.SetCellValue("Страница 1", "A"+strconv.Itoa(j), "e-mail: kng@profilss.ru"); err != nil {
		fmt.Println(err)
	}

	j++ // Переходим к следующей строке для данных товара

	// Выполняем объединение ячеек
	if err := f.MergeCell("Страница 1", "A"+strconv.Itoa(j), "E"+strconv.Itoa(j)); err != nil {
		fmt.Println(err)
	}
	if err := f.SetRowHeight("Страница 1", j, 30); err != nil {
		fmt.Println(err)
	}
	if err = f.SetCellStyle("Страница 1", "A"+strconv.Itoa(j), "B"+strconv.Itoa(j), defaultStyle); err != nil {
		fmt.Println(err)
	}

	// Вставляем подвал таблицы
	if err = f.SetCellValue("Страница 1", "A"+strconv.Itoa(j), "По всем интересующим вас вопросам обращайтесь"); err != nil {
		fmt.Println(err)
	}

	j++ // Переходим к следующей строке для данных товара

	// Выполняем объединение ячеек
	if err := f.MergeCell("Страница 1", "A"+strconv.Itoa(j), "E"+strconv.Itoa(j)); err != nil {
		fmt.Println(err)
	}
	if err := f.SetRowHeight("Страница 1", j, 30); err != nil {
		fmt.Println(err)
	}
	if err = f.SetCellStyle("Страница 1", "A"+strconv.Itoa(j), "B"+strconv.Itoa(j), sizeStyle); err != nil {
		fmt.Println(err)
	}

	// Вставляем подвал таблицы
	if err = f.SetCellValue("Страница 1", "A"+strconv.Itoa(j), "Курмышева Наталья"); err != nil {
		fmt.Println(err)
	}

	j++ // Переходим к следующей строке для данных товара

	// Выполняем объединение ячеек
	if err := f.MergeCell("Страница 1", "A"+strconv.Itoa(j), "E"+strconv.Itoa(j)); err != nil {
		fmt.Println(err)
	}
	if err := f.SetRowHeight("Страница 1", j, 30); err != nil {
		fmt.Println(err)
	}
	if err = f.SetCellStyle("Страница 1", "A"+strconv.Itoa(j), "B"+strconv.Itoa(j), defaultStyle); err != nil {
		fmt.Println(err)
	}

	// Вставляем подвал таблицы
	if err = f.SetCellValue("Страница 1", "A"+strconv.Itoa(j), "8-902-380-20-11"); err != nil {
		fmt.Println(err)
	}

	j++ // Переходим к следующей строке для данных товара

	// Save spreadsheet by the given path.
	if err := f.SaveAs(filepath.Join(destination, "/doc/", "Профиль-С прайс-лист 2024.xlsx")); err != nil {
		fmt.Println(err)
	}

	render.JSON(w, r, nil)
}
