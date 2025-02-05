# GoScrape

GoScrape is a Golang learning project designed to scrape websites and output the data in JSON format. This project demonstrates web scraping techniques for both static and dynamic websites using different libraries.

## Versions

- **v1**: Scrapes **static websites** using the [Colly](https://github.com/gocolly/colly) library.
- **v2**: Targets **dynamic websites** using the [Rod](https://github.com/go-rod/rod) library. *This version is still in progress.*

## Build Instructions

To build the GoScrape project, run the following command:

```bash
go build -o goscrape ./cmd/goscrape/main.go
```

## Run Instructions

After building the project, you can run it using:

```bash
./goscrape {your_website_page}
```

Replace `{your_website_page}` with the URL of the website you want to scrape.

## Output

The scraped data will be saved in a `.json` file in the project directory.

## Dependencies

- [Colly](https://github.com/gocolly/colly) for static web scraping (v1)
- [Rod](https://github.com/go-rod/rod) for dynamic web scraping (v2, ongoing)

## License

This project is for learning purposes and is open-source under the MIT License.

---

Happy scraping! 🚀

