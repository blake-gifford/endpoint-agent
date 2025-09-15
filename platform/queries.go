package platform

var queryInstalledSoftwareFallbacks = []string{
	`SELECT name, version FROM deb_packages;`,
	`SELECT name, bundle_version as version FROM apps;`,
}

var querySystemInfo = `
		SELECT 
			hostname, 
			computer_name as name, 
			hardware_model as model, 
			hardware_serial as serial, 
			hardware_vendor as manufacturer 
		FROM 
			system_info;
	`
