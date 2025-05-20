package config

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName                     string
	AppKey                      string
	BaseUrl                     string
	FrontEndUrl                 string
	Environtment                string
	SmtpHost                    string
	SmtpPort                    int
	SmtpSender                  string
	SmtpPassword                string
	XenditApiKey                string
	MidtransServerKey           string
	RajaOngkirKey               string
	DirPath                     string
	DatabaseURL                 string
	DatabaseUsername            string
	DatabasePassword            string
	DatabaseHost                string
	DatabasePort                string
	DatabaseName                string
	DatabasePlannerName         string
	PathDB                      string
	CacheURL                    string
	CachePassword               string
	LoggerLevel                 string
	ContextTimeout              int
	Port                        string
	GoogleClientID              string
	GoogleClientSecret          string
	POSFrontendUrl              string
	BOFrontendUrl               string
	EnableCronJob               bool
	EnableConcurrent            bool
	EnableCSRF                  bool
	EnableDatabaseAutomigration bool
	EnableSaas                  bool
	EnableAPIKey                bool
	APIKey                      string
	IsDesktop                   bool
	GoldAPIUrl                  string
	// GoldAPIKey                  string
	OpenExchangeRatesUrl        string
	APIGeolocationAPIKey        string
	RunLocalDatabaseVia         string
	DesktopUserFullname         string
	DesktopUserEmail            string
	DesktopUserPassword         string
	DesktopUserPhone            string
	DesktopUserCompanyName      string
	DesktopOwnerAccountLimit    int
	DesktopEmployeeAccountLimit int
	DesktopBranchLimit          int
	DesktopUserBranchName       string
	ComputerUserName            string
	ComputerName                string
	ComputerPath                string
	ComputerAppData             string
	IcanDelivEmail              string
	IcanDelivPassword           string
}

func LoadConfig() (config *Config) {

	if err := godotenv.Load(RootPath() + `/.env`); err != nil {
		fmt.Println(err)
	}

	appName := os.Getenv("APP_NAME")
	appKey := os.Getenv("APP_KEY")
	baseurl := os.Getenv("BASE_URL")
	frontendurl := os.Getenv("FRONT_END_URL")
	environment := strings.ToUpper(os.Getenv("ENVIRONMENT"))
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))
	smtpSender := os.Getenv("SMTP_SENDER")
	smtpPassword := os.Getenv("SMTP_PASSWORD")
	xenditApiKey := os.Getenv("XENDIT_API_KEY")
	midtransServerKey := os.Getenv("MIDTRANS_SERVER_KEY")
	rajaOngkirKey := os.Getenv("RAJA_ONGKIR_KEY")
	dirPath := os.Getenv("DIR_PATH")
	databaseURL := os.Getenv("DATABASE_URL")
	databaseUsername := os.Getenv("DATABASE_USERNAME")
	databasePassword := os.Getenv("DATABASE_PASSWORD")
	databaseHost := os.Getenv("DATABASE_HOST")
	databasePort := os.Getenv("DATABASE_PORT")
	databaseName := os.Getenv("DATABASE_NAME")
	databasePlannerName := os.Getenv("DATABASE_PLANNER_NAME")
	PathDB := os.Getenv("PATH_DB")
	cacheURL := os.Getenv("CACHE_URL")
	cachePassword := os.Getenv("CACHE_PASSWORD")
	loggerLevel := os.Getenv("LOGGER_LEVEL")
	contextTimeout, _ := strconv.Atoi(os.Getenv("CONTEXT_TIMEOUT"))
	port := os.Getenv("PORT")
	googleClientID := os.Getenv("GOOGLE_CLIENT_ID")
	googleClientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	posFrontendUrl := os.Getenv("POS_FRONT_END_URL")
	boFrontendUrl := os.Getenv("BO_FRONT_END_URL")
	enableCronJob, _ := strconv.ParseBool(os.Getenv("ENABLE_CRONJOB"))
	enableConcurrent, _ := strconv.ParseBool(os.Getenv("ENABLE_CONCURRENT"))
	enableCSRF, _ := strconv.ParseBool(os.Getenv("ENABLE_CSRF"))
	enableDatabaseAutomigration, _ := strconv.ParseBool(os.Getenv("ENABLE_DATABASE_AUTOMIGRATION"))
	enableSaas, _ := strconv.ParseBool(os.Getenv("ENABLE_SAAS"))
	enableApiKey, _ := strconv.ParseBool(os.Getenv("ENABLE_API_KEY"))
	goldAPIUrl := os.Getenv("GOLDAPI_URL")
	// goldAPIKey := os.Getenv("GOLDAPI_KEY")
	openExchangeRatesUrl := os.Getenv("OPENEXCHAGERATES_URL")
	apiGeolocationAPIKey := os.Getenv("APIGEOLOCATION_API_KEY")
	runLocalDatabaseVia := strings.ToUpper(os.Getenv("RUN_LOCAL_DATABASE_VIA"))
	desktopUserFullname := os.Getenv("DESKTOP_USER_FULLNAME")
	desktopUserEmail := os.Getenv("DESKTOP_USER_EMAIL")
	desktopUserPassword := os.Getenv("DESKTOP_USER_PASSWORD")
	desktopUserPhone := os.Getenv("DESKTOP_USER_PHONE")
	desktopUserCompanyName := os.Getenv("DESKTOP_USER_COMPANY_NAME")
	desktopOwnerAccountLimit, _ := strconv.Atoi(os.Getenv("DEKSTOP_OWNER_ACCOUNT_LIMIT"))
	desktopEmployeeAccountLimit, _ := strconv.Atoi(os.Getenv("DESKTOP_EMPLOYEE_ACCOUNT_LIMIT"))
	desktopBranchLimit, _ := strconv.Atoi(os.Getenv("DESKTOP_BRANCH_LIMIT"))
	desktopUserBranchName := os.Getenv("DESKTOP_USER_BRANCH_NAME")
	computerUsername := os.Getenv("USERNAME")
	computerName, _ := os.Hostname()
	computerPath := os.Getenv("PATH")
	computerAppData := filepath.Join(os.Getenv("APPDATA"), appName)
	apiKey := os.Getenv("API_KEY")
	icanDelivEmail := os.Getenv("ICAN_DELIV_EMAIL")
	icanDelivPassword := os.Getenv("ICAN_DELIV_PASSWORD")

	var isDesktop bool
	if environment == "DESKTOP" {
		isDesktop = true
		enableSaas = false
		enableDatabaseAutomigration = true
	}

	return &Config{
		AppName:                     appName,
		AppKey:                      appKey,
		BaseUrl:                     baseurl,
		FrontEndUrl:                 frontendurl,
		Environtment:                environment,
		SmtpHost:                    smtpHost,
		SmtpPort:                    smtpPort,
		SmtpSender:                  smtpSender,
		SmtpPassword:                smtpPassword,
		XenditApiKey:                xenditApiKey,
		MidtransServerKey:           midtransServerKey,
		RajaOngkirKey:               rajaOngkirKey,
		DirPath:                     dirPath,
		DatabaseURL:                 databaseURL,
		DatabaseUsername:            databaseUsername,
		DatabasePassword:            databasePassword,
		DatabaseHost:                databaseHost,
		DatabasePort:                databasePort,
		DatabaseName:                databaseName,
		DatabasePlannerName:         databasePlannerName,
		PathDB:                      PathDB,
		CacheURL:                    cacheURL,
		CachePassword:               cachePassword,
		LoggerLevel:                 loggerLevel,
		ContextTimeout:              contextTimeout,
		Port:                        port,
		GoogleClientID:              googleClientID,
		GoogleClientSecret:          googleClientSecret,
		POSFrontendUrl:              posFrontendUrl,
		BOFrontendUrl:               boFrontendUrl,
		EnableCronJob:               enableCronJob,
		EnableConcurrent:            enableConcurrent,
		EnableCSRF:                  enableCSRF,
		EnableDatabaseAutomigration: enableDatabaseAutomigration,
		EnableSaas:                  enableSaas,
		IsDesktop:                   isDesktop,
		GoldAPIUrl:                  goldAPIUrl,
		// GoldAPIKey:                  goldAPIKey,
		OpenExchangeRatesUrl:        openExchangeRatesUrl,
		APIGeolocationAPIKey:        apiGeolocationAPIKey,
		RunLocalDatabaseVia:         runLocalDatabaseVia,
		DesktopUserFullname:         desktopUserFullname,
		DesktopUserEmail:            desktopUserEmail,
		DesktopUserPassword:         desktopUserPassword,
		DesktopUserPhone:            desktopUserPhone,
		DesktopUserCompanyName:      desktopUserCompanyName,
		DesktopOwnerAccountLimit:    desktopOwnerAccountLimit,
		DesktopEmployeeAccountLimit: desktopEmployeeAccountLimit,
		DesktopBranchLimit:          desktopBranchLimit,
		DesktopUserBranchName:       desktopUserBranchName,
		ComputerUserName:            computerUsername,
		ComputerName:                computerName,
		ComputerPath:                computerPath,
		ComputerAppData:             computerAppData,
		EnableAPIKey:                enableApiKey,
		APIKey:                      apiKey,
		IcanDelivEmail:              icanDelivEmail,
		IcanDelivPassword:           icanDelivPassword,
	}
}

func RootPath() string {
	projectDirName := os.Getenv("DIR_NAME")
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))
	return string(rootPath)
}

func WriteLogForDesktop(filename, message, category string) {
	if LoadConfig().IsDesktop {
		file, err := os.OpenFile(filepath.Join(LoadConfig().ComputerAppData, filename), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("Failed to open file. Error:", err)
			return
		}
		defer file.Close()

		writer := bufio.NewWriter(file)

		_, err = writer.WriteString("[" + time.Now().Format("2024-04-25 15:04:05") + "] [" + strings.ToUpper(category) + "] " + message + "\n")
		if err != nil {
			fmt.Println("Failed to write into file. Error:", err)
			return
		}

		err = writer.Flush()
		if err != nil {
			fmt.Println("Failed to save file. Error:", err)
			return
		}

		fmt.Println("Success to write into file.")
	}
}
