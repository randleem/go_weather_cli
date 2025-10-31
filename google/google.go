package google

type WeatherResponse struct {
	CurrentTime string `json:"currentTime"`
	TimeZone    struct {
		ID string `json:"id"`
	} `json:"timeZone"`
	IsDaytime        bool `json:"isDaytime"`
	WeatherCondition struct {
		IconBaseURI string `json:"iconBaseUri"`
		Description struct {
			Text         string `json:"text"`
			LanguageCode string `json:"languageCode"`
		} `json:"description"`
		Type string `json:"type"`
	} `json:"weatherCondition"`
	Temperature struct {
		Degrees float64 `json:"degrees"`
		Unit    string  `json:"unit"`
	} `json:"temperature"`
	FeelsLikeTemperature struct {
		Degrees float64 `json:"degrees"`
		Unit    string  `json:"unit"`
	} `json:"feelsLikeTemperature"`
	DewPoint struct {
		Degrees float64 `json:"degrees"`
		Unit    string  `json:"unit"`
	} `json:"dewPoint"`
	HeatIndex struct {
		Degrees float64 `json:"degrees"`
		Unit    string  `json:"unit"`
	} `json:"heatIndex"`
	WindChill struct {
		Degrees float64 `json:"degrees"`
		Unit    string  `json:"unit"`
	} `json:"windChill"`
	RelativeHumidity float64 `json:"relativeHumidity"`
	UVIndex          float64 `json:"uvIndex"`
	Precipitation    struct {
		Probability struct {
			Percent float64 `json:"percent"`
			Type    string  `json:"type"`
		} `json:"probability"`
		QPF struct {
			Quantity float64 `json:"quantity"`
			Unit     string  `json:"unit"`
		} `json:"qpf"`
	} `json:"precipitation"`
	ThunderstormProbability float64 `json:"thunderstormProbability"`
	AirPressure             struct {
		MeanSeaLevelMillibars float64 `json:"meanSeaLevelMillibars"`
	} `json:"airPressure"`
	Wind struct {
		Direction struct {
			Degrees  float64 `json:"degrees"`
			Cardinal string  `json:"cardinal"`
		} `json:"direction"`
		Speed struct {
			Value float64 `json:"value"`
			Unit  string  `json:"unit"`
		} `json:"speed"`
		Gust struct {
			Value float64 `json:"value"`
			Unit  string  `json:"unit"`
		} `json:"gust"`
	} `json:"wind"`
	Visibility struct {
		Distance float64 `json:"distance"`
		Unit     string  `json:"unit"`
	} `json:"visibility"`
	CloudCover               float64 `json:"cloudCover"`
	CurrentConditionsHistory struct {
		TemperatureChange struct {
			Degrees float64 `json:"degrees"`
			Unit    string  `json:"unit"`
		} `json:"temperatureChange"`
		MaxTemperature struct {
			Degrees float64 `json:"degrees"`
			Unit    string  `json:"unit"`
		} `json:"maxTemperature"`
		MinTemperature struct {
			Degrees float64 `json:"degrees"`
			Unit    string  `json:"unit"`
		} `json:"minTemperature"`
		QPF struct {
			Quantity float64 `json:"quantity"`
			Unit     string  `json:"unit"`
		} `json:"qpf"`
	} `json:"currentConditionsHistory"`
}

type GeocodeResponse struct {
	Results []struct {
		AddressComponents []struct {
			LongName  string   `json:"long_name"`
			ShortName string   `json:"short_name"`
			Types     []string `json:"types"`
		} `json:"address_components"`
		FormattedAddress string `json:"formatted_address"`
		Geometry         struct {
			Bounds struct {
				Northeast struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"northeast"`
				Southwest struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"southwest"`
			} `json:"bounds"`
			Location struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"location"`
			LocationType string `json:"location_type"`
			Viewport     struct {
				Northeast struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"northeast"`
				Southwest struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"southwest"`
			} `json:"viewport"`
		} `json:"geometry"`
		PlaceID  string `json:"place_id"`
		PlusCode *struct {
			CompoundCode string `json:"compound_code"`
			GlobalCode   string `json:"global_code"`
		} `json:"plus_code,omitempty"`
		Types []string `json:"types"`
	} `json:"results"`
	Status string `json:"status"`
}
