package main

import (
	"encoding/xml"
	"fmt"
	"github.com/coldruze/scraping/data_types"
	"github.com/gocolly/colly"
)

var news []data_types.Item // Срез, в котором хранятся данные о новостях.

func main() {
	// Определение констант для типов данных.
	const (
		XMLDataType = "XML"
		TXTDataType = "TXT"
	)

	c := colly.NewCollector(
		colly.AllowedDomains("nfuunit.ru"), // Разрешенные домены для парсинга.
		colly.Async(true),                  // Асинхронный режим работы.
	)

	url := "https://nfuunit.ru/news" // URL страницы с новостями.

	getData(c, url) // Получение данных о новостях.

	// Проверка наличия данных о новостях.
	if news == nil {
		fmt.Println("Произошла ошибка, либо нет новостей")
		return
	}

	printNews(XMLDataType) // Вывод данных о новостях в формате XML.
}

// getData осуществляет сбор данных о новостях с указанного URL-адреса.
func getData(c *colly.Collector, url string) {
	c.OnHTML(".news-item", func(e *colly.HTMLElement) {
		temp := data_types.Item{}
		temp.StoryUrl = "https://nfuunit.ru/" + e.ChildAttr("a", "href") // Получение URL новости.
		temp.Title = e.ChildAttr("a", "title")                           // Получение заголовка новости.
		news = append(news, temp)                                        // Добавление новости в срез.
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String(), "\n") // Вывод информации о текущей странице.
	})

	if err := c.Visit(url); err != nil {
		fmt.Println(err) // Обработка ошибки при парсинге страницы.
	}

	c.Wait() // Ожидание окончания работы всех запросов.
}

// printNews выводит данные о новостях в указанном формате (XML или TXT).
func printNews(dataType string) {
	switch dataType {
	case "XML":
		newsXml := data_types.DataXML{
			Items: news,
		}

		xmlData, err := xml.MarshalIndent(newsXml, "", "\t") // Преобразование данных в XML.
		if err != nil {
			fmt.Println("Ошибка при создании XML:", err)
			return
		}

		fmt.Println(string(xmlData)) // Вывод данных в формате XML.

	case "TXT":
		for _, n := range news {
			fmt.Printf("Новость: %s\nСсылка: %s\n", n.Title, n.StoryUrl) // Вывод данных в текстовом формате.
			fmt.Println("")
		}
	}
}
