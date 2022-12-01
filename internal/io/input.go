package io

import (
	"errors"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strconv"
)

const AOC_BASE_URL = "https://adventofcode.com"

func getSessionToken() (string, bool) {
	return os.LookupEnv("AOC_SESSION")
}

func getExternalInput(day int) ([]byte, bool) {
	session, ok := getSessionToken()

	if !ok {
		return make([]byte, 0), false
	}

	cookieJar, _ := cookiejar.New(nil)
	cookie := &http.Cookie{
		Name:   "session",
		Value:  session,
		MaxAge: 300,
	}
	u, _ := url.Parse(AOC_BASE_URL)
	cookieJar.SetCookies(u, []*http.Cookie{cookie})

	client := http.Client{Jar: cookieJar}

	response, err := client.Get(AOC_BASE_URL + "/2022/day/" + strconv.Itoa(day+1) + "/input")

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	if response.StatusCode == 404 {
		return make([]byte, 0), false
	}

	body, _ := io.ReadAll(response.Body)
	return body, true
}

func getCachedInput(day int) ([]byte, error) {
	data, err := os.ReadFile(getInputAssetName(day))

	if err != nil {
		return nil, err
	}
	return data, nil
}

func cacheInput(day int, data []byte) error {
	return os.WriteFile(getInputAssetName(day), data, 0644)
}

func GetInput(day int) ([]byte, error) {
	data, err := getCachedInput(day)

	if err == nil {
		return data, nil
	}

	data, ok := getExternalInput(day)

	if ok {
		err := cacheInput(day, data)
		if err != nil {
			return nil, err
		}
		return data, nil
	}

	return nil, errors.New("Failed to get input for day " + strconv.Itoa(day))
}
