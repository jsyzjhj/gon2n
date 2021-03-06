package cmd

const (
	keyPrefix               = "edge."
	configFileDefault       = ""
	serverHostPortKey       = keyPrefix + "serverHostPort"
	configFileKey           = keyPrefix + "configFile"
	allowP2PKey             = keyPrefix + "allowP2P"
	allowRoutingKey         = keyPrefix + "allowRouting"
	communityNameKey        = keyPrefix + "communityName"
	disablePMTUDiscoveryKey = keyPrefix + "disablePMTUDiscovery"
	disableMulticastKey     = keyPrefix + "disableMulticast"
	dynamicIPModeKey        = keyPrefix + "dynamicIPMode"
	encryptionKeyKey        = keyPrefix + "encryptionKey"
	localPortKey            = keyPrefix + "localPort"
	managementPortKey       = keyPrefix + "managementPort"
	registerIntervalKey     = keyPrefix + "registerInterval"
	registerTTLKey          = keyPrefix + "registerTTL"
	supernodeHostPortKey    = keyPrefix + "supernodeHostPort"
	typeOfServiceKey        = keyPrefix + "typeOfService"
	encryptionMethodKey     = keyPrefix + "encryptionMethod"
	deviceNameKey           = keyPrefix + "deviceName"
	addressModeKey          = keyPrefix + "addressMode"
	deviceIPKey             = keyPrefix + "deviceIP"
	deviceNetmaskKey        = keyPrefix + "deviceNetmask"
	deviceMACAddressKey     = keyPrefix + "deviceMACAddress"
	mtuKey                  = keyPrefix + "MTU"
)

var (
	serverHostPortFlag string
	configFileFlag     string
)
