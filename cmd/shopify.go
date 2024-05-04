package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var Format string

type ShopifyCSVColumn int

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

	DEFAULT_MAXIMUM_TO_GENERATE int = 10000
)

var shopifyColumns = [17]ShopifyCSVColumn{
	FIRST_NAME, LAST_NAME, EMAIL,
	ACCEPTS_EMAIL_MARKETING, DEFAULT_ADDRESS_COMPANY,
	DEFAULT_ADDRESS_ADDRESS_1, DEFAULT_ADDRESS_ADDRESS_2,
	DEFAULT_ADDRESS_CITY, DEFAULT_ADDRESS_PROVINCE_CODE,
	DEFAULT_ADDRESS_ZIP, DEFAULT_ADDRESS_PHONE,
	PHONE, ACCEPTS_SMS_MARKETING, TAGS, NOTE, TAX_EXEMPT,
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

func generateCSV() [DEFAULT_MAXIMUM_TO_GENERATE][len(shopifyColumns)]string {
	var records [DEFAULT_MAXIMUM_TO_GENERATE][len(shopifyColumns)]string

	for i := 1; i <= DEFAULT_MAXIMUM_TO_GENERATE; i++ {
		var record [len(shopifyColumns)]string
		for j := 0; j < len(shopifyColumns); j++ {
			record[j] = csvContent(ShopifyCSVColumn(j))
		}
		records[i-1] = record
		fmt.Print(record)
	}

	return records
}

func csvContent(column ShopifyCSVColumn) string {
	var content string

	switch column {
	case FIRST_NAME:
	case LAST_NAME:
		content = generateName(column)
	case EMAIL:
		content = generateEmail(column)
	case ACCEPTS_EMAIL_MARKETING:
	case ACCEPTS_SMS_MARKETING:
	case TAX_EXEMPT:
		content = generateYesNo(column)
	case DEFAULT_ADDRESS_COMPANY:
		content = generateAddressCompany(column)
	case DEFAULT_ADDRESS_ADDRESS_1:
		content = generateAddress1(column)
	case DEFAULT_ADDRESS_ADDRESS_2:
		content = generateAddress2(column)
	case DEFAULT_ADDRESS_CITY:
		content = generateAddressCity(column)
	case DEFAULT_ADDRESS_PROVINCE_CODE:
		content = generateDefaultProvinceCode(column)
	case DEFAULT_ADDRESS_COUNTRY_CODE:
		content = generateAddressCountryCode(column)
	case DEFAULT_ADDRESS_ZIP:
		content = generateDefaultAddressZip(column)
	case DEFAULT_ADDRESS_PHONE:
	case PHONE:
		content = generatePhone(column)
	case TAGS:
		content = generateTags(column)
	case NOTE:
		content = generateNote(column)
	default:
		panic("Unsupported column!")
	}

	return content
}

func generateName(column ShopifyCSVColumn) string {
	return "Test" + strconv.Itoa(int(column))
}

func generateEmail(column ShopifyCSVColumn) string {
	return "Test" + strconv.Itoa(int(column)) + "@lunaris.jp"
}

func generateYesNo(column ShopifyCSVColumn) string {
	return "no"
}

func generateAddressCompany(column ShopifyCSVColumn) string {
	return ""
}

func generateAddress1(column ShopifyCSVColumn) string {
	return strconv.Itoa(int(column)) + " Fake Street"
}

func generateAddress2(column ShopifyCSVColumn) string {
	return ""
}

func generateAddressCity(column ShopifyCSVColumn) string {
	return "Nakano"
}

func generateDefaultProvinceCode(column ShopifyCSVColumn) string {
	return "Tokyo"
}

func generateAddressCountryCode(column ShopifyCSVColumn) string {
	return "JP"
}

func generateDefaultAddressZip(column ShopifyCSVColumn) string {
	return "164-0012"
}

func generatePhone(column ShopifyCSVColumn) string {
	return "+81 (555) 5555 5555"
}

func generateTags(column ShopifyCSVColumn) string {
	return ""
}

func generateNote(column ShopifyCSVColumn) string {
	return ""
}
