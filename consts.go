package binance_connector

import (
	"fmt"
	"os"
)

const Name = "binance-connector-go"

const Version = "0.5.2"

func SetProxy(address string) error {
	if err := os.Setenv("HTTP_PROXY", address); err != nil {
		return fmt.Errorf("failed to set http proxy, err: %s", err.Error())
	}

	if err := os.Setenv("HTTPS_PROXY", address); err != nil {
		return fmt.Errorf("failed to set https proxy, err: %s", err.Error())
	}

	return nil
}
