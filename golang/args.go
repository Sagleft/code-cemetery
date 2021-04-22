package main

import (
	"flag"
)

type apiKeys struct {
	apikeyPublic string
	apikeySecret string
}

type launchParams struct {
	apikeys apiKeys
	isdebug bool
}

func parseLaunchParams() launchParams {
	flag.Parse()
	apikeysArr := apiKeys{*argsApikeyPublic, *argsApikeySecret}
	return launchParams{
		apikeys: apikeysArr,
		isdebug: *argsIsDebug,
	}
}
