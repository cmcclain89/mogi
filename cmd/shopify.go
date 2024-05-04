package cmd

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var Format string

type ShopifyCSVColumn uint8

const (
	FIRST_NAME                    ShopifyCSVColumn = 0
	LAST_NAME                     ShopifyCSVColumn = 1
	EMAIL                         ShopifyCSVColumn = 2
	ACCEPTS_EMAIL_MARKETING       ShopifyCSVColumn = 3
	DEFAULT_ADDRESS_COMPANY       ShopifyCSVColumn = 4
	DEFAULT_ADDRESS_ADDRESS_1     ShopifyCSVColumn = 5
	DEFAULT_ADDRESS_ADDRESS_2     ShopifyCSVColumn = 6
	DEFAULT_ADDRESS_CITY          ShopifyCSVColumn = 7
	DEFAULT_ADDRESS_PROVINCE_CODE ShopifyCSVColumn = 8
	DEFAULT_ADDRESS_COUNTRY_CODE  ShopifyCSVColumn = 9
	DEFAULT_ADDRESS_ZIP           ShopifyCSVColumn = 10
	DEFAULT_ADDRESS_PHONE         ShopifyCSVColumn = 11
	PHONE                         ShopifyCSVColumn = 12
	ACCEPTS_SMS_MARKETING         ShopifyCSVColumn = 13
	TAGS                          ShopifyCSVColumn = 14
	NOTE                          ShopifyCSVColumn = 15
	TAX_EXEMPT                    ShopifyCSVColumn = 16

	DEFAULT_MAXIMUM_TO_GENERATE int = 100000
)

var headers_map = map[ShopifyCSVColumn]string{
	FIRST_NAME:                    "First Name",
	LAST_NAME:                     "Last Name",
	EMAIL:                         "Email",
	ACCEPTS_EMAIL_MARKETING:       "Accepts Email Marketing",
	DEFAULT_ADDRESS_COMPANY:       "Default Address Company",
	DEFAULT_ADDRESS_ADDRESS_1:     "Default Address Address1",
	DEFAULT_ADDRESS_ADDRESS_2:     "Default Address Address2",
	DEFAULT_ADDRESS_CITY:          "Default Address City",
	DEFAULT_ADDRESS_PROVINCE_CODE: "Default Address Province Code",
	DEFAULT_ADDRESS_COUNTRY_CODE:  "Default Address Country Code",
	DEFAULT_ADDRESS_ZIP:           "Default Address Zip",
	DEFAULT_ADDRESS_PHONE:         "Default Address Phone",
	PHONE:                         "Phone",
	ACCEPTS_SMS_MARKETING:         "Accepts SMS Marketing",
	TAGS:                          "Tags",
	NOTE:                          "Note",
	TAX_EXEMPT:                    "Tax Exempt",
}

func init() {
	shopifyCmd.PersistentFlags().StringVarP(&Format, "format", "f", "", "Format of the mock data")
	rootCmd.AddCommand(shopifyCmd)
}

var shopifyCmd = &cobra.Command{
	Use:   "shopify",
	Short: "Generate mock data for Shopify",
	Long: `Generate mock data for Shopify. 
Planned support is just CSV for now.`,
	Run: func(cmd *cobra.Command, args []string) {
		generateCSV()
		fmt.Println("neat")
	},
}

func generateCSV() {
	records := make([][]string, DEFAULT_MAXIMUM_TO_GENERATE+1)

	headers := make([]string, len(headers_map))
	for i := 0; i < len(headers_map); i++ {
		headers[i] = headers_map[ShopifyCSVColumn(i)]
	}
	records[0] = headers

	for i := 1; i <= DEFAULT_MAXIMUM_TO_GENERATE; i++ {
		record := make([]string, len(headers_map))
		for j := 0; j < len(headers_map); j++ {
			record[j] = csvContent(ShopifyCSVColumn(j), i)
		}
		records[i] = record
	}

	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	csvFile, _ := os.Create(dirname + "/test.csv")
	defer csvFile.Close()

	w := csv.NewWriter(csvFile)
	w.WriteAll(records)

	if err := w.Error(); err != nil {
		log.Fatalln("error writing csv:", err)
	}
}

func csvContent(column ShopifyCSVColumn, record int) string {
	var content string

	switch column {
	case FIRST_NAME:
		fallthrough
	case LAST_NAME:
		content = generateName(record)
	case EMAIL:
		content = generateEmail(record)
	case ACCEPTS_EMAIL_MARKETING:
		fallthrough
	case ACCEPTS_SMS_MARKETING:
		fallthrough
	case TAX_EXEMPT:
		content = generateYesNo(record)
	case DEFAULT_ADDRESS_COMPANY:
		content = generateAddressCompany()
	case DEFAULT_ADDRESS_ADDRESS_1:
		content = generateAddress1(record)
	case DEFAULT_ADDRESS_ADDRESS_2:
		content = generateAddress2()
	case DEFAULT_ADDRESS_CITY:
		content = generateAddressCity()
	case DEFAULT_ADDRESS_PROVINCE_CODE:
		content = generateDefaultProvinceCode()
	case DEFAULT_ADDRESS_COUNTRY_CODE:
		content = generateAddressCountryCode()
	case DEFAULT_ADDRESS_ZIP:
		content = generateDefaultAddressZip()
	case DEFAULT_ADDRESS_PHONE:
		fallthrough
	case PHONE:
		content = generatePhone()
	case TAGS:
		content = generateTags()
	case NOTE:
		content = generateNote()
	default:
		panic("Unsupported column!")
	}

	return content
}

func generateName(record int) string {
	return "Test" + strconv.Itoa(record)
}

func generateEmail(record int) string {
	return "Test" + strconv.Itoa(record) + "@lunaris.jp"
}

func generateYesNo(record int) string {
	if record%2 == 0 {
		return "yes"
	} else {
		return "no"
	}
}

func generateAddressCompany() string {
	return ""
}

func generateAddress1(record int) string {
	return strconv.Itoa(record%1000+1) + " Fake Street"
}

func generateAddress2() string {
	return ""
}

func generateAddressCity() string {
	return "Nakano"
}

func generateDefaultProvinceCode() string {
	return "Tokyo"
}

func generateAddressCountryCode() string {
	return "JA"
}

func generateDefaultAddressZip() string {
	return "164-0012"
}

func generatePhone() string {
	return "+81 (555) 5555 5555"
}

func generateTags() string {
	return ""
}

func generateNote() string {
	return ""
}
