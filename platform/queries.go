package platform

var (
	queryInstalledSoftwareFallbacks = []string{
		`SELECT name, bundle_version as version FROM apps WHERE name IS NOT NULL AND bundle_version IS NOT NULL AND name != '' AND bundle_version != '';`,
		`SELECT name, version FROM chrome_extensions WHERE name IS NOT NULL AND version IS NOT NULL AND name != '' AND version != '';`,
		`SELECT name, version FROM firefox_addons WHERE name IS NOT NULL AND version IS NOT NULL AND name != '' AND version != '';`,
		`SELECT name, version FROM ie_extensions WHERE name IS NOT NULL AND version IS NOT NULL AND name != '' AND version != '';`,
		`SELECT name, version FROM deb_packages WHERE name IS NOT NULL AND version IS NOT NULL AND name != '' AND version != '';`,
		`SELECT name, version FROM rpm_packages WHERE name IS NOT NULL AND version IS NOT NULL AND name != '' AND version != '';`,
		`SELECT name, version FROM yum_packages WHERE name IS NOT NULL AND version IS NOT NULL AND name != '' AND version != '';`,
		`SELECT name, version FROM programs WHERE name IS NOT NULL AND version IS NOT NULL AND name != '' AND version != '';`,
	}

	querySystemInfo = `
		SELECT 
			hostname, 
			computer_name as name, 
			hardware_model as model, 
			hardware_serial as serial, 
			hardware_vendor as manufacturer 
		FROM 
			system_info;
	`
)
