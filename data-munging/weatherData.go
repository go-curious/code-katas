package datamunging

import (
	"log"
	"strconv"
	"strings"
)

type WeatherRecord struct {
	Day     int
	MaxTemp int
	MinTemp int
}

func parseWeatherDataFile(path string) []WeatherRecord {
	return []WeatherRecord{}
}

func (wr WeatherRecord) getTemperatureSpread() int {
	return wr.MaxTemp - wr.MinTemp
}

func ParseWeatherRecord(dr string) WeatherRecord {
	fields := strings.Fields(dr)

	day, err := strconv.Atoi(fields[0])
	if err != nil {
		log.Fatal(err)
	}

	return WeatherRecord{
		Day:     day,
		MaxTemp: parseTemperature(fields[1]),
		MinTemp: parseTemperature(fields[2]),
	}
}

func parseTemperature(tmpStr string) int {
	trimmed := strings.Trim(tmpStr, "*")
	temp, err := strconv.Atoi(trimmed)
	if err != nil {
		log.Fatal(err)
	}
	return temp
}
