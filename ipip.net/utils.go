package ipip_net

import (
	"encoding/binary"
	"net"
	"strings"
	"time"
)

func makeLocation(data string) *Location {
	location := new(Location)

	mapping := []*string{
		&location.Country, &location.Province, &location.City, &location.Unit,
		&location.ISP,
		&location.Latitude, &location.Longitude,
		&location.TimeZoneCode, &location.TimeZoneUTC,
		&location.GB2260Code, &location.CallingCode, &location.ISO3166Code, &location.ContinentCode,
	}

	for index, field := range strings.Split(data, "\t") {
		*mapping[index] = field
	}

	return location
}

func ip2int(ip net.IP) uint32 {
	return binary.BigEndian.Uint32(ip.To4())
}

func padding(data []byte, length int) []byte {
	payload := make([]byte, length)
	copy(payload[0:len(data)], data)
	return payload
}

func resolvePublishDate(version string) (date time.Time) {
	layout := "2006010215"
	loc, _ := time.LoadLocation("Asia/Shanghai")
	date, _ = time.ParseInLocation(layout, version, loc)
	return
}
